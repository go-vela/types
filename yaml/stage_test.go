// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/go-vela/types/pipeline"

	"github.com/goccy/go-yaml"
)

func TestYaml_StageSlice_ToPipeline(t *testing.T) {
	// setup tests
	tests := []struct {
		stages *StageSlice
		want   *pipeline.StageSlice
	}{
		{
			stages: &StageSlice{
				{
					Name:  "echo",
					Needs: []string{"clone"},
					Steps: StepSlice{
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
				},
			},
			want: &pipeline.StageSlice{
				{
					Name:  "echo",
					Needs: []string{"clone"},
					Steps: pipeline.ContainerSlice{
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
			},
		},
	}

	// run tests
	for _, test := range tests {
		got := test.stages.ToPipeline()

		// WARNING: hack to compare stages
		//
		// Channel values can only be compared for equality.
		// Two channel values are considered equal if they
		// originated from the same make call meaning they
		// refer to the same channel value in memory.
		for i, stage := range *got {
			tmp := *test.want

			tmp[i].Done = stage.Done
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ToPipeline is %v, want %v", got, test.want)
		}
	}
}

func TestYaml_StageSlice_UnmarshalYAML(t *testing.T) {
	// setup types
	var (
		b   []byte
		err error
	)

	// setup tests
	tests := []struct {
		failure bool
		file    string
		want    *StageSlice
	}{
		{
			failure: false,
			file:    "testdata/stage.yml",
			want: &StageSlice{
				{
					Name:  "dependencies",
					Needs: []string{"clone"},
					Steps: StepSlice{
						{
							Commands: []string{"./gradlew downloadDependencies"},
							Environment: map[string]string{
								"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
								"GRADLE_USER_HOME": ".gradle",
							},
							Image: "openjdk:latest",
							Name:  "install",
							Pull:  "always",
						},
					},
				},
				{
					Name:  "test",
					Needs: []string{"dependencies"},
					Steps: StepSlice{
						{
							Commands: []string{"./gradlew check"},
							Environment: map[string]string{
								"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
								"GRADLE_USER_HOME": ".gradle",
							},
							Name:  "test",
							Image: "openjdk:latest",
							Pull:  "always",
						},
					},
				},
				{
					Name:  "build",
					Needs: []string{"dependencies"},
					Steps: StepSlice{
						{
							Commands: []string{"./gradlew build"},
							Environment: map[string]string{
								"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
								"GRADLE_USER_HOME": ".gradle",
							},
							Name:  "build",
							Image: "openjdk:latest",
							Pull:  "always",
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
		{
			failure: true,
			file:    "",
			want:    nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := new(StageSlice)

		if len(test.file) > 0 {
			b, err = ioutil.ReadFile(test.file)
			if err != nil {
				t.Errorf("unable to read file: %v", err)
			}
		} else {
			b = []byte("- foo")
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

func TestStageSlice_Validate(t *testing.T) {
	//setup types
	tests := []struct {
		name    string
		file    string
		wantErr bool
	}{
		{
			name:    "success: minimal stage block",
			file:    "testdata/stage/validate/minimal.yml",
			wantErr: false,
		},
		{
			name:    "failure: missing commands, environment, parameters, secrets or template yaml tag",
			file:    "testdata/stage/validate/missing.yml",
			wantErr: true,
		},
		{
			name:    "failure: missing name yaml tag",
			file:    "testdata/stage/validate/missing_name.yml",
			wantErr: true,
		},
		{
			name:    "failure: missing image yaml tag",
			file:    "testdata/stage/validate/missing_image.yml",
			wantErr: true,
		},
		{
			name:    "failure: bad image tag data",
			file:    "testdata/stage/validate/bad_image.yml",
			wantErr: true,
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
			if err := b.Stages.Validate(pipeline); (err != nil) != test.wantErr {
				t.Errorf("Validate is %v, want %v", err, test.wantErr)
			}
		})
	}
}
