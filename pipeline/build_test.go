// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"reflect"
	"testing"

	"github.com/go-vela/types/constants"
)

func TestPipeline_Build_Purge_Stages(t *testing.T) {
	// setup types
	p := &Build{
		Services: ContainerSlice{
			&Container{
				Image:  "postgres:latest",
				Name:   "postgres",
				Number: 1,
			}},
		Worker: Worker{
			Flavor:   "16cpu8gb",
			Platform: "gcp",
		},
		Stages: StageSlice{
			&Stage{
				Name:  "install",
				Needs: []string{"clone"},
				Steps: ContainerSlice{
					&Container{
						Commands: []string{"./gradlew downloadDependencies"},
						Image:    "openjdk:latest",
						Name:     "install",
						Number:   1,
						Pull:     true,
					},
				},
			},
			&Stage{
				Name:  "test",
				Needs: []string{"install"},
				Steps: ContainerSlice{
					&Container{
						Commands: []string{"./gradlew check"},
						Image:    "openjdk:latest",
						Name:     "test",
						Number:   2,
						Pull:     true,
						Ruleset: Ruleset{
							If: Rules{
								Event: []string{"push"},
							},
							Operator: "and",
						},
					},
				},
			},
		},
	}

	r := &RuleData{
		Branch: "master",
		Event:  "pull_request",
		Path:   []string{},
		Repo:   "foo/bar",
		Status: "success",
		Tag:    "refs/heads/master",
	}

	want := &Build{
		Services: ContainerSlice{
			&Container{
				Image:  "postgres:latest",
				Name:   "postgres",
				Number: 1,
			}},
		Worker: Worker{
			Flavor:   "16cpu8gb",
			Platform: "gcp",
		},
		Stages: StageSlice{
			&Stage{
				Name:  "install",
				Needs: []string{"clone"},
				Steps: ContainerSlice{
					&Container{
						Commands: []string{"./gradlew downloadDependencies"},
						Image:    "openjdk:latest",
						Name:     "install",
						Number:   1,
						Pull:     true,
					},
				},
			},
		},
	}

	// run test
	got := p.Purge(r)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Purge is %v, want %v", got, want)
	}
}

func TestPipeline_Build_Purge_Steps(t *testing.T) {
	// setup types
	p := &Build{
		Services: ContainerSlice{
			&Container{
				Image:  "postgres:latest",
				Name:   "postgres",
				Number: 1,
			}},
		Worker: Worker{
			Flavor:   "16cpu8gb",
			Platform: "gcp",
		},
		Steps: ContainerSlice{
			&Container{
				Commands: []string{"./gradlew downloadDependencies"},
				Image:    "openjdk:latest",
				Name:     "install",
				Number:   1,
				Pull:     true,
			},
			&Container{
				Commands: []string{"./gradlew check"},
				Image:    "openjdk:latest",
				Name:     "test",
				Number:   2,
				Pull:     true,
				Ruleset: Ruleset{
					If: Rules{
						Event: []string{"push"},
					},
					Operator: "and",
				},
			},
		},
	}

	r := &RuleData{
		Branch: "master",
		Event:  "pull_request",
		Path:   []string{},
		Repo:   "foo/bar",
		Status: "success",
		Tag:    "refs/heads/master",
	}

	want := &Build{
		Services: ContainerSlice{
			&Container{
				Image:  "postgres:latest",
				Name:   "postgres",
				Number: 1,
			}},
		Worker: Worker{
			Flavor:   "16cpu8gb",
			Platform: "gcp",
		},
		Steps: ContainerSlice{
			&Container{
				Commands: []string{"./gradlew downloadDependencies"},
				Image:    "openjdk:latest",
				Name:     "install",
				Number:   1,
				Pull:     true,
			},
		},
	}

	// run test
	got := p.Purge(r)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Purge is %v, want %v", got, want)
	}
}

func TestPipeline_Build_Purge_Invalid(t *testing.T) {
	// setup types
	p := &Build{
		Stages: StageSlice{
			&Stage{
				Name:  "install",
				Needs: []string{"clone"},
				Steps: ContainerSlice{
					&Container{
						Commands: []string{"./gradlew downloadDependencies"},
						Image:    "openjdk:latest",
						Name:     "install",
						Number:   1,
						Pull:     true,
					},
				},
			},
		},
		Steps: ContainerSlice{
			&Container{
				Commands: []string{"./gradlew downloadDependencies"},
				Image:    "openjdk:latest",
				Name:     "install",
				Number:   1,
				Pull:     true,
			},
		},
	}

	r := &RuleData{
		Branch: "master",
		Event:  "pull_request",
		Path:   []string{},
		Repo:   "foo/bar",
		Status: "success",
		Tag:    "refs/heads/master",
	}

	// run test
	got := p.Purge(r)

	if got != nil {
		t.Errorf("Purge is %v, want nil", got)
	}
}

func TestPipeline_Build_Sanitize_Stages(t *testing.T) {
	// setup types
	p := &Build{
		ID: "foo bar_1",
		Stages: StageSlice{
			{
				Name: "test",
				Steps: ContainerSlice{
					{
						ID:       "foo_bar_1_test_echo foo",
						Commands: []string{"echo foo"},
						Image:    "alpine:latest",
						Name:     "echo foo",
						Number:   1,
						Pull:     true,
					},
				},
			},
		},
	}

	want := &Build{
		ID: "foo-bar_1",
		Stages: StageSlice{
			{
				Name: "test",
				Steps: ContainerSlice{
					{
						ID:       "foo_bar_1_test_echo-foo",
						Commands: []string{"echo foo"},
						Image:    "alpine:latest",
						Name:     "echo foo",
						Number:   1,
						Pull:     true,
					},
				},
			},
		},
	}

	// run test
	got := p.Sanitize(constants.DriverDocker)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sanitize is %v, want %v", got, want)
	}
}

func TestPipeline_Build_Sanitize_Steps(t *testing.T) {
	// setup types
	p := &Build{
		ID: "foo bar_1",
		Steps: ContainerSlice{
			{
				ID:       "step_foo_bar_1_echo foo",
				Commands: []string{"echo foo"},
				Image:    "alpine:latest",
				Name:     "echo foo",
				Number:   1,
				Pull:     true,
			},
		},
	}

	want := &Build{
		ID: "foo-bar_1",
		Steps: ContainerSlice{
			{
				ID:       "step_foo_bar_1_echo-foo",
				Commands: []string{"echo foo"},
				Image:    "alpine:latest",
				Name:     "echo foo",
				Number:   1,
				Pull:     true,
			},
		},
	}

	// run test
	got := p.Sanitize(constants.DriverDocker)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sanitize is %v, want %v", got, want)
	}
}

func TestPipeline_Build_Sanitize_StagesAndSteps(t *testing.T) {
	// setup types
	p := &Build{
		ID: "foo bar_1",
		Stages: StageSlice{
			{
				Name: "test",
				Steps: ContainerSlice{
					{
						ID:       "foo_bar_1_test_echo foo",
						Commands: []string{"echo foo"},
						Image:    "alpine:latest",
						Name:     "echo foo",
						Number:   1,
						Pull:     true,
					},
				},
			},
		},
		Steps: ContainerSlice{
			{
				ID:       "step_foo_bar_1_echo foo",
				Commands: []string{"echo foo"},
				Image:    "alpine:latest",
				Name:     "echo foo",
				Number:   1,
				Pull:     true,
			},
		},
	}

	// run test
	got := p.Sanitize(constants.DriverDocker)

	if got != nil {
		t.Errorf("Sanitize is %v, want nil", got)
	}
}

func TestPipeline_Build_Sanitize_Docker(t *testing.T) {
	// setup types
	p := &Build{
		ID: "foo bar_1",
		Services: ContainerSlice{
			{
				ID:     "service_foo bar_1_postgres",
				Image:  "postgres:latest",
				Name:   "postgres",
				Number: 1,
			},
		},
	}

	want := &Build{
		ID: "foo-bar_1",
		Services: ContainerSlice{
			{
				ID:     "service_foo-bar_1_postgres",
				Image:  "postgres:latest",
				Name:   "postgres",
				Number: 1,
			},
		},
	}

	// run test
	got := p.Sanitize(constants.DriverDocker)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sanitize is %v, want %v", got, want)
	}
}

func TestPipeline_Build_Sanitize_Kubernetes(t *testing.T) {
	// setup types
	p := &Build{
		ID: "foo bar_1.",
		Services: ContainerSlice{
			{
				ID:     "service_foo bar_1_postgres",
				Image:  "postgres:latest",
				Name:   "postgres",
				Number: 1,
			},
		},
	}

	want := &Build{
		ID: "foo-bar-1-",
		Services: ContainerSlice{
			{
				ID:     "service-foo-bar-1-postgres",
				Image:  "postgres:latest",
				Name:   "postgres",
				Number: 1,
			},
		},
	}

	// run test
	got := p.Sanitize(constants.DriverKubernetes)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sanitize is %v, want %v", got, want)
	}
}
