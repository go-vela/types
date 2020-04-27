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
	containers := testContainers()
	*containers = (*containers)[:len(*containers)-1]

	// setup tests
	tests := []struct {
		containers *ContainerSlice
		want       *ContainerSlice
	}{
		{
			containers: testContainers(),
			want:       containers,
		},
		{
			containers: new(ContainerSlice),
			want:       new(ContainerSlice),
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

		got := test.containers.Purge(r)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Purge is %v, want %v", got, test.want)
		}
	}
}

func TestPipeline_ContainerSlice_Sanitize(t *testing.T) {
	// setup types
	containers := testContainers()
	(*containers)[0].ID = "step_github-octocat._1_init"
	(*containers)[1].ID = "step_github-octocat._1_clone"
	(*containers)[2].ID = "step_github-octocat._1_echo"

	kubeContainers := testContainers()
	(*kubeContainers)[0].ID = "step-github-octocat--1-init"
	(*kubeContainers)[1].ID = "step-github-octocat--1-clone"
	(*kubeContainers)[2].ID = "step-github-octocat--1-echo"

	// setup tests
	tests := []struct {
		driver     string
		containers *ContainerSlice
		want       *ContainerSlice
	}{
		{
			driver:     constants.DriverDocker,
			containers: testContainers(),
			want:       containers,
		},
		{
			driver:     constants.DriverKubernetes,
			containers: testContainers(),
			want:       kubeContainers,
		},
		{
			driver:     constants.DriverDocker,
			containers: new(ContainerSlice),
			want:       new(ContainerSlice),
		},
		{
			driver:     constants.DriverKubernetes,
			containers: new(ContainerSlice),
			want:       new(ContainerSlice),
		},
		{
			driver:     "foo",
			containers: new(ContainerSlice),
			want:       nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.containers.Sanitize(test.driver)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Sanitize is %v, want %v", got, test.want)
		}
	}
}

func testContainers() *ContainerSlice {
	return &ContainerSlice{
		{
			ID:          "step_github octocat._1_init",
			Directory:   "/home/github/octocat",
			Environment: map[string]string{"FOO": "bar"},
			Image:       "#init",
			Name:        "init",
			Number:      1,
			Pull:        true,
		},
		{
			ID:          "step_github octocat._1_clone",
			Directory:   "/home/github/octocat",
			Environment: map[string]string{"FOO": "bar"},
			Image:       "target/vela-git:v0.3.0",
			Name:        "clone",
			Number:      2,
			Pull:        true,
		},
		{
			ID:          "step_github octocat._1_echo",
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
	}
}
