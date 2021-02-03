// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"

	"github.com/goccy/go-yaml"
)

func TestYaml_StepSlice_ToPipeline(t *testing.T) {
	// setup tests
	tests := []struct {
		steps *StepSlice
		want  *pipeline.ContainerSlice
	}{
		{
			steps: &StepSlice{
				{
					Commands:    []string{"echo hello"},
					Detach:      false,
					Entrypoint:  []string{"/bin/sh"},
					Environment: map[string]string{"FOO": "bar"},
					Image:       "alpine:latest",
					Name:        "echo",
					Privileged:  false,
					Pull:        "not_present",
					Ruleset: Ruleset{
						If: Rules{
							Branch:  []string{"master"},
							Comment: []string{"test comment"},
							Event:   []string{"push"},
							Path:    []string{"foo.txt"},
							Repo:    []string{"github/octocat"},
							Status:  []string{"success"},
							Tag:     []string{"v0.1.0"},
							Target:  []string{"production"},
						},
						Unless: Rules{
							Branch:  []string{"master"},
							Comment: []string{"real comment"},
							Event:   []string{"pull_request"},
							Path:    []string{"bar.txt"},
							Repo:    []string{"github/octocat"},
							Status:  []string{"failure"},
							Tag:     []string{"v0.2.0"},
							Target:  []string{"production"},
						},
						Operator: "and",
						Continue: false,
					},
					Secrets: StepSecretSlice{
						{
							Source: "docker_username",
							Target: "plugin_username",
						},
					},
					Ulimits: UlimitSlice{
						{
							Name: "foo",
							Soft: 1024,
							Hard: 2048,
						},
					},
					Volumes: VolumeSlice{
						{
							Source:      "/foo",
							Destination: "/bar",
							AccessMode:  "ro",
						},
					},
				},
			},
			want: &pipeline.ContainerSlice{
				{
					Commands:    []string{"echo hello"},
					Detach:      false,
					Entrypoint:  []string{"/bin/sh"},
					Environment: map[string]string{"FOO": "bar"},
					Image:       "alpine:latest",
					Name:        "echo",
					Privileged:  false,
					Pull:        "not_present",
					Ruleset: pipeline.Ruleset{
						If: pipeline.Rules{
							Branch:  []string{"master"},
							Comment: []string{"test comment"},
							Event:   []string{"push"},
							Path:    []string{"foo.txt"},
							Repo:    []string{"github/octocat"},
							Status:  []string{"success"},
							Tag:     []string{"v0.1.0"},
							Target:  []string{"production"},
						},
						Unless: pipeline.Rules{
							Branch:  []string{"master"},
							Comment: []string{"real comment"},
							Event:   []string{"pull_request"},
							Path:    []string{"bar.txt"},
							Repo:    []string{"github/octocat"},
							Status:  []string{"failure"},
							Tag:     []string{"v0.2.0"},
							Target:  []string{"production"},
						},
						Operator: "and",
						Continue: false,
					},
					Secrets: pipeline.StepSecretSlice{
						{
							Source: "docker_username",
							Target: "plugin_username",
						},
					},
					Ulimits: pipeline.UlimitSlice{
						{
							Name: "foo",
							Soft: 1024,
							Hard: 2048,
						},
					},
					Volumes: pipeline.VolumeSlice{
						{
							Source:      "/foo",
							Destination: "/bar",
							AccessMode:  "ro",
						},
					},
				},
			},
		},
	}

	// run tests
	for _, test := range tests {
		got := test.steps.ToPipeline()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ToPipeline is %v, want %v", got, test.want)
		}
	}
}

func TestYaml_StepSlice_UnmarshalYAML(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		file    string
		want    *StepSlice
	}{
		{
			failure: false,
			file:    "testdata/step.yml",
			want: &StepSlice{
				{
					Commands: raw.StringSlice{"./gradlew downloadDependencies"},
					Environment: raw.StringSliceMap{
						"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
						"GRADLE_USER_HOME": ".gradle",
					},
					Name:  "install",
					Image: "openjdk:latest",
					Pull:  "always",
				},
				{
					Commands: raw.StringSlice{"./gradlew check"},
					Environment: raw.StringSliceMap{
						"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
						"GRADLE_USER_HOME": ".gradle",
					},
					Name:  "test",
					Image: "openjdk:latest",
					Pull:  "always",
				},
				{
					Commands: raw.StringSlice{"./gradlew build"},
					Environment: raw.StringSliceMap{
						"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
						"GRADLE_USER_HOME": ".gradle",
					},
					Name:  "build",
					Image: "openjdk:latest",
					Pull:  "always",
				},
				{
					Name:  "docker_build",
					Image: "plugins/docker:18.09",
					Pull:  "always",
					Parameters: map[string]interface{}{
						"registry": "index.docker.io",
						"repo":     "github/octocat",
						"tags":     []interface{}{"latest", "dev"},
					},
				},
				{
					Name: "templated_publish",
					Pull: "not_present",
					Template: StepTemplate{
						Name: "docker_publish",
						Variables: map[string]interface{}{
							"registry": "index.docker.io",
							"repo":     "github/octocat",
							"tags":     []interface{}{"latest", "dev"},
						},
					},
				},
			},
		},
		{
			failure: true,
			file:    "testdata/invalid.yml",
			want:    nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := new(StepSlice)

		b, err := ioutil.ReadFile(test.file)
		if err != nil {
			t.Errorf("unable to read file: %v", err)
		}

		err = yaml.Unmarshal(b, got)

		if test.failure {
			if err == nil {
				t.Errorf("UnmarshalYAML should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("UnmarshalYAML returned err: %v", err)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("UnmarshalYAML is %v, want %v", got, test.want)
		}
	}
}

func TestStepSlice_Validate(t *testing.T) {
	//setup types
	tests := []struct {
		name    string
		file    string
		wantErr bool
	}{
		{
			name:    "success: minimal step block",
			file:    "testdata/step/validate/minimal.yml",
			wantErr: false,
		},
		{
			name:    "failure: missing commands, environment, parameters, secrets or template yaml tag",
			file:    "testdata/step/validate/missing.yml",
			wantErr: true,
		},
		{
			name:    "failure: missing name yaml tag",
			file:    "testdata/step/validate/missing_name.yml",
			wantErr: true,
		},
		{
			name:    "failure: missing image yaml tag",
			file:    "testdata/step/validate/missing_image.yml",
			wantErr: true,
		},
		{
			name:    "failure: bad image tag data",
			file:    "testdata/step/validate/bad_image.yml",
			wantErr: true,
		},
		{
			name:    "success: step with variable image",
			file:    "testdata/step/validate/variable_image.yml",
			wantErr: false,
		},
	}

	// run tests
	for _, test := range tests {
		b := new(Build)

		pipeline, err := ioutil.ReadFile(test.file)
		if err != nil {
			t.Errorf("Reading file for Validate returned err: %v", err)
		}

		err = yaml.Unmarshal(pipeline, b)

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}

		t.Run(test.name, func(t *testing.T) {
			if err := b.Steps.Validate(pipeline); (err != nil) != test.wantErr {
				t.Errorf("Validate is %v, want %v", err, test.wantErr)
			}
		})
	}
}
