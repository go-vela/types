// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"reflect"
	"testing"
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
