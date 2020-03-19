// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"reflect"
	"testing"

	"github.com/go-vela/types/constants"
)

func TestPipeline_StageSlice_Purge(t *testing.T) {
	// setup types
	done := make(chan error)

	s := StageSlice{
		&Stage{
			Done:  done,
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
			Done:  done,
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
	}

	r := &RuleData{
		Branch: "master",
		Event:  "pull_request",
		Path:   []string{},
		Repo:   "foo/bar",
		Status: "success",
		Tag:    "refs/heads/master",
	}

	want := &StageSlice{
		&Stage{
			Done:  done,
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
	}

	// run test
	got := s.Purge(r)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Purge is %v, want %v", got, want)
	}
}

func TestPipeline_StageSlice_Sanitize_Docker(t *testing.T) {
	// setup types
	done := make(chan error)

	s := &StageSlice{
		{
			Done: done,
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
	}

	want := &StageSlice{
		{
			Done: done,
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
	}

	// run test
	got := s.Sanitize(constants.DriverDocker)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sanitize is %v, want %v", got, want)
	}
}

func TestPipeline_StageSlice_Sanitize_Kubernetes(t *testing.T) {
	// setup types
	done := make(chan error)

	s := &StageSlice{
		{
			Done: done,
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
				{
					ID:       "foo_bar_1_test_echo_bar",
					Commands: []string{"echo bar"},
					Image:    "alpine:latest",
					Name:     "echo_bar",
					Number:   2,
					Pull:     true,
				},
				{
					ID:       "foo_bar_1_test_echo.baz",
					Commands: []string{"echo baz"},
					Image:    "alpine:latest",
					Name:     "echo.baz",
					Number:   3,
					Pull:     true,
				},
			},
		},
	}

	want := &StageSlice{
		{
			Done: done,
			Name: "test",
			Steps: ContainerSlice{
				{
					ID:       "foo-bar-1-test-echo-foo",
					Commands: []string{"echo foo"},
					Image:    "alpine:latest",
					Name:     "echo foo",
					Number:   1,
					Pull:     true,
				},
				{
					ID:       "foo-bar-1-test-echo-bar",
					Commands: []string{"echo bar"},
					Image:    "alpine:latest",
					Name:     "echo_bar",
					Number:   2,
					Pull:     true,
				},
				{
					ID:       "foo-bar-1-test-echo-baz",
					Commands: []string{"echo baz"},
					Image:    "alpine:latest",
					Name:     "echo.baz",
					Number:   3,
					Pull:     true,
				},
			},
		},
	}

	// run test
	got := s.Sanitize(constants.DriverKubernetes)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sanitize is %v, want %v", got, want)
	}
}

func TestPipeline_StageSlice_Sanitize_NoDriver(t *testing.T) {
	// setup types
	s := &StageSlice{}

	// run test
	got := s.Sanitize("")

	if got != nil {
		t.Errorf("Sanitize is %v, want nil", got)
	}
}
