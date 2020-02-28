// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"reflect"
	"testing"

	"github.com/go-vela/types/constants"
)

func TestPipeline_ContainerSlice_Purge(t *testing.T) {
	// setup types
	s := ContainerSlice{
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
	}

	r := &RuleData{
		Branch: "master",
		Event:  "pull_request",
		Path:   []string{},
		Repo:   "foo/bar",
		Status: "success",
		Tag:    "refs/heads/master",
	}

	want := &ContainerSlice{
		&Container{
			Commands: []string{"./gradlew downloadDependencies"},
			Image:    "openjdk:latest",
			Name:     "install",
			Number:   1,
			Pull:     true,
		},
	}

	// run test
	got := s.Purge(r)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Purge is %v, want %v", got, want)
	}
}

func TestPipeline_ContainerSlice_Sanitize_Docker(t *testing.T) {
	// setup types
	c := &ContainerSlice{
		{
			ID:       "step_foo_bar_1_echo foo",
			Commands: []string{"echo foo"},
			Image:    "alpine:latest",
			Name:     "echo foo",
			Number:   1,
			Pull:     true,
		},
	}

	want := &ContainerSlice{
		{
			ID:       "step_foo_bar_1_echo-foo",
			Commands: []string{"echo foo"},
			Image:    "alpine:latest",
			Name:     "echo foo",
			Number:   1,
			Pull:     true,
		},
	}

	// run test
	got := c.Sanitize(constants.DriverDocker)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sanitize is %v, want %v", got, want)
	}
}

func TestPipeline_ContainerSlice_Sanitize_Kubernetes(t *testing.T) {
	// setup types
	c := &ContainerSlice{
		{
			ID:       "step_foo_bar_1_echo foo",
			Commands: []string{"echo foo"},
			Image:    "alpine:latest",
			Name:     "echo foo",
			Number:   1,
			Pull:     true,
		},
		{
			ID:       "step_foo_bar_1_echo_bar",
			Commands: []string{"echo bar"},
			Image:    "alpine:latest",
			Name:     "echo_bar",
			Number:   2,
			Pull:     true,
		},
		{
			ID:       "step_foo_bar_1_echo.baz",
			Commands: []string{"echo baz"},
			Image:    "alpine:latest",
			Name:     "echo.baz",
			Number:   3,
			Pull:     true,
		},
	}

	want := &ContainerSlice{
		{
			ID:       "step-foo-bar-1-echo-foo",
			Commands: []string{"echo foo"},
			Image:    "alpine:latest",
			Name:     "echo foo",
			Number:   1,
			Pull:     true,
		},
		{
			ID:       "step-foo-bar-1-echo-bar",
			Commands: []string{"echo bar"},
			Image:    "alpine:latest",
			Name:     "echo_bar",
			Number:   2,
			Pull:     true,
		},
		{
			ID:       "step-foo-bar-1-echo-baz",
			Commands: []string{"echo baz"},
			Image:    "alpine:latest",
			Name:     "echo.baz",
			Number:   3,
			Pull:     true,
		},
	}

	// run test
	got := c.Sanitize(constants.DriverKubernetes)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sanitize is %v, want %v", got, want)
	}
}

func TestPipeline_ContainerSlice_Sanitize_NoDriver(t *testing.T) {
	// setup types
	c := &ContainerSlice{}

	// run test
	got := c.Sanitize("")

	if got != nil {
		t.Errorf("Sanitize is %v, want nil", got)
	}
}

func TestPipeline_Container_Sanitize_Stage(t *testing.T) {
	// setup types
	got := &Container{
		ID:       "github_octocat_1_install deps_install",
		Commands: []string{"./gradlew downloadDependencies"},
		Image:    "openjdk:latest",
		Name:     "install",
		Number:   1,
		Pull:     true,
	}

	want := &Container{
		ID:       "github_octocat_1_install-deps_install",
		Commands: []string{"./gradlew downloadDependencies"},
		Image:    "openjdk:latest",
		Name:     "install",
		Number:   1,
		Pull:     true,
	}

	// run test
	got.Sanitize()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sanitize is %v, want %v", got, want)
	}
}

func TestPipeline_Container_Sanitize_Step(t *testing.T) {
	// setup types
	got := &Container{
		ID:       "github_octocat_1_install deps",
		Commands: []string{"./gradlew downloadDependencies"},
		Image:    "openjdk:latest",
		Name:     "install deps",
		Number:   1,
		Pull:     true,
	}

	want := &Container{
		ID:       "github_octocat_1_install-deps",
		Commands: []string{"./gradlew downloadDependencies"},
		Image:    "openjdk:latest",
		Name:     "install deps",
		Number:   1,
		Pull:     true,
	}

	// run test
	got.Sanitize()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sanitize is %v, want %v", got, want)
	}
}
