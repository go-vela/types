// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"reflect"
	"testing"
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
