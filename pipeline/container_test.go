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

func TestPipeline_Container_Execute(t *testing.T) {
	// setup types
	containers := testContainers()
	*containers = (*containers)[:len(*containers)-1]

	// setup tests
	tests := []struct {
		container *Container
		ruleData  *RuleData
		want      bool
	}{
		{ // empty container with build success
			container: &Container{
				Name:     "empty",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "push",
				Repo:   "foo/bar",
				Status: "success",
			},
			want: true,
		},
		{ // empty container with build failure
			container: &Container{
				Name:     "empty",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "push",
				Repo:   "foo/bar",
				Status: "failure",
			},
			want: false,
		},
		{ // status success container with build success
			container: &Container{
				Name:     "status success",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
				Ruleset: Ruleset{
					If: Rules{
						Status: []string{constants.StatusSuccess},
					},
				},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "push",
				Repo:   "foo/bar",
				Status: "success",
			},
			want: true,
		},
		{ // status success container with build failure
			container: &Container{
				Name:     "status success",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
				Ruleset: Ruleset{
					If: Rules{
						Status: []string{constants.StatusSuccess},
					},
				},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "push",
				Repo:   "foo/bar",
				Status: "failure",
			},
			want: false,
		},
		{ // status/failure success container with build failure
			container: &Container{
				Name:     "status/failure",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
				Ruleset: Ruleset{
					If: Rules{
						Status: []string{constants.StatusSuccess, constants.StatusFailure},
					},
				},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "push",
				Repo:   "foo/bar",
				Status: "success",
			},
			want: true,
		},
		{ // status/failure success container with build failure
			container: &Container{
				Name:     "status/failure",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
				Ruleset: Ruleset{
					If: Rules{
						Status: []string{constants.StatusSuccess, constants.StatusFailure},
					},
				},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "push",
				Repo:   "foo/bar",
				Status: "failure",
			},
			want: true,
		},
		{ // no status container with build success
			container: &Container{
				Name:     "branch/event/status",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
				Ruleset: Ruleset{
					If: Rules{
						Branch: []string{"master"},
						Event:  []string{constants.EventPush},
					},
				},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "push",
				Repo:   "foo/bar",
				Status: "failure",
			},
			want: false,
		},
		{ // branch/event/status container with build success
			container: &Container{
				Name:     "branch/event/status",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
				Ruleset: Ruleset{
					If: Rules{
						Branch: []string{"master"},
						Event:  []string{constants.EventPush},
						Status: []string{constants.StatusSuccess},
					},
				},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "push",
				Repo:   "foo/bar",
				Status: "success",
			},
			want: true,
		},
		{ // branch/event/status container with build failure
			container: &Container{
				Name:     "branch/event/status",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
				Ruleset: Ruleset{
					If: Rules{
						Branch: []string{"master"},
						Event:  []string{constants.EventPush},
						Status: []string{constants.StatusSuccess},
					},
				},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "push",
				Repo:   "foo/bar",
				Status: "failure",
			},
			want: false,
		},
		{ // branch/event/status container with build failure
			container: &Container{
				Name:     "branch/event/status",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
				Ruleset: Ruleset{
					If: Rules{
						Branch: []string{"master"},
						Event:  []string{constants.EventPush},
						Status: []string{constants.StatusSuccess},
					},
					Operator: "or",
				},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "push",
				Repo:   "foo/bar",
				Status: "failure",
			},
			want: true,
		},
		{ // tag/event/status container with build success
			container: &Container{
				Name:     "tag/event/status",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
				Ruleset: Ruleset{
					If: Rules{
						Tag:    []string{"v0.1.0"},
						Event:  []string{constants.EventTag},
						Status: []string{constants.StatusSuccess},
					},
				},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "tag",
				Repo:   "foo/bar",
				Status: "success",
				Tag:    "v*",
			},
			want: true,
		},
		{ // status unless success container with build success
			container: &Container{
				Name:     "status unless",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
				Ruleset: Ruleset{
					Unless: Rules{
						Status: []string{constants.StatusSuccess},
					},
				},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "push",
				Repo:   "foo/bar",
				Status: "success",
			},
			want: false,
		},
		{ // status unless success container with build success
			container: &Container{
				Name:     "status unless",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
				Ruleset: Ruleset{
					Unless: Rules{
						Status: []string{constants.StatusSuccess},
					},
				},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "push",
				Repo:   "foo/bar",
				Status: "failure",
			},
			want: true,
		},
		{ // status unless success container with build success
			container: &Container{
				Name:     "status unless",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
				Ruleset: Ruleset{
					Unless: Rules{
						Branch: []string{"master"},
						Event:  []string{constants.EventPush},
						Status: []string{constants.StatusSuccess},
					},
				},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "push",
				Repo:   "foo/bar",
				Status: "success",
			},
			want: false,
		},
		{ // status unless success container with build success
			container: &Container{
				Name:     "status unless",
				Image:    "alpine:latest",
				Commands: []string{"echo \"Hey Vela\""},
				Ruleset: Ruleset{
					Unless: Rules{
						Branch: []string{"dev"},
						Event:  []string{constants.EventPush},
						Status: []string{constants.StatusSuccess},
					},
				},
			},
			ruleData: &RuleData{
				Branch: "master",
				Event:  "pull_request",
				Repo:   "foo/bar",
				Status: "failure",
			},
			want: true,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.container.Execute(test.ruleData)

		if got != test.want {
			t.Errorf("Container Execute %s is %v, want %v", test.container.Name, got, test.want)
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
