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
	stages := testStages()
	*stages = (*stages)[:len(*stages)-1]

	// setup tests
	tests := []struct {
		stages *StageSlice
		want   *StageSlice
	}{
		{
			stages: testStages(),
			want:   stages,
		},
		{
			stages: new(StageSlice),
			want:   new(StageSlice),
		},
	}

	// run tests
	for _, test := range tests {
		r := &RuleData{
			Branch: "master",
			Event:  "pull_request",
			Path:   []string{},
			Repo:   "foo/bar",
			Tag:    "refs/heads/master",
		}

		got := test.stages.Purge(r)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Purge is %v, want %v", got, test.want)
		}
	}
}

func TestPipeline_StageSlice_Sanitize(t *testing.T) {
	// setup types
	stages := testStages()
	(*stages)[0].Steps[0].ID = "github-octocat._1_init_init"
	(*stages)[1].Steps[0].ID = "github-octocat._1_clone_clone"
	(*stages)[2].Steps[0].ID = "github-octocat._1_echo_echo"

	kubeStages := testStages()
	(*kubeStages)[0].Steps[0].ID = "github-octocat--1-init-init"
	(*kubeStages)[1].Steps[0].ID = "github-octocat--1-clone-clone"
	(*kubeStages)[2].Steps[0].ID = "github-octocat--1-echo-echo"

	// setup tests
	tests := []struct {
		driver string
		stages *StageSlice
		want   *StageSlice
	}{
		{
			driver: constants.DriverDocker,
			stages: testStages(),
			want:   stages,
		},
		{
			driver: constants.DriverKubernetes,
			stages: testStages(),
			want:   kubeStages,
		},
		{
			driver: constants.DriverDocker,
			stages: new(StageSlice),
			want:   new(StageSlice),
		},
		{
			driver: constants.DriverKubernetes,
			stages: new(StageSlice),
			want:   new(StageSlice),
		},
		{
			driver: "foo",
			stages: new(StageSlice),
			want:   nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.stages.Sanitize(test.driver)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Sanitize is %v, want %v", got, test.want)
		}
	}
}

func testStages() *StageSlice {
	return &StageSlice{
		{
			Name: "init",
			Steps: ContainerSlice{
				{
					ID:          "github octocat._1_init_init",
					Directory:   "/home/github/octocat",
					Environment: map[string]string{"FOO": "bar"},
					Image:       "#init",
					Name:        "init",
					Number:      1,
					Pull:        true,
				},
			},
		},
		{
			Name:  "clone",
			Needs: []string{"init"},
			Steps: ContainerSlice{
				{
					ID:          "github octocat._1_clone_clone",
					Directory:   "/home/github/octocat",
					Environment: map[string]string{"FOO": "bar"},
					Image:       "target/vela-git:v0.3.0",
					Name:        "clone",
					Number:      2,
					Pull:        true,
				},
			},
		},
		{
			Name:  "echo",
			Needs: []string{"clone"},
			Steps: ContainerSlice{
				{
					ID:          "github octocat._1_echo_echo",
					Commands:    []string{"echo hello"},
					Directory:   "/home/github/octocat",
					Environment: map[string]string{"FOO": "bar"},
					Image:       "alpine:latest",
					Name:        "echo",
					Number:      3,
					Pull:        true,
					Ruleset: Ruleset{
						If:       Rules{Event: []string{"push"}},
						Operator: "and",
					},
				},
			},
		},
	}
}
