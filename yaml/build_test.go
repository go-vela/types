// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/buildkite/yaml"
	"github.com/go-vela/types/raw"
)

func TestYaml_Build_UnmarshalYAML(t *testing.T) {
	// setup tests
	tests := []struct {
		file string
		want *Build
	}{
		{
			file: "testdata/build.yml",
			want: &Build{
				Version: "1",
				Metadata: Metadata{
					Template:    false,
					Clone:       nil,
					Environment: []string{"steps", "services", "secrets"},
				},
				Environment: raw.StringSliceMap{
					"HELLO": "Hello, Global Message",
				},
				Worker: Worker{
					Flavor:   "16cpu8gb",
					Platform: "gcp",
				},
				Services: ServiceSlice{
					{
						Ports: []string{"5432:5432"},
						Environment: raw.StringSliceMap{
							"POSTGRES_DB": "foo",
						},
						Name:  "postgres",
						Image: "postgres:latest",
						Pull:  "not_present",
					},
				},
				Steps: StepSlice{
					{
						Commands: raw.StringSlice{"./gradlew downloadDependencies"},
						Environment: raw.StringSliceMap{
							"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
							"GRADLE_USER_HOME": ".gradle",
						},
						Image: "openjdk:latest",
						Name:  "install",
						Pull:  "always",
						Ruleset: Ruleset{
							If:       Rules{Event: []string{"push", "pull_request:opened", "pull_request:synchronized", "pull_request:edited"}},
							Matcher:  "filepath",
							Operator: "and",
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
					{
						Commands: raw.StringSlice{"./gradlew check"},
						Environment: raw.StringSliceMap{
							"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
							"GRADLE_USER_HOME": ".gradle",
						},
						Name:  "test",
						Image: "openjdk:latest",
						Pull:  "always",
						Ruleset: Ruleset{
							If:       Rules{Event: []string{"push", "pull_request:opened", "pull_request:synchronized"}},
							Matcher:  "filepath",
							Operator: "and",
						},
						Volumes: VolumeSlice{
							{
								Source:      "/foo",
								Destination: "/bar",
								AccessMode:  "ro",
							},
						},
						Ulimits: UlimitSlice{
							{
								Name: "foo",
								Soft: 1024,
								Hard: 2048,
							},
						},
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
						Ruleset: Ruleset{
							If:       Rules{Event: []string{"push", "pull_request:opened", "pull_request:synchronized"}},
							Matcher:  "filepath",
							Operator: "and",
						},
						Volumes: VolumeSlice{
							{
								Source:      "/foo",
								Destination: "/bar",
								AccessMode:  "ro",
							},
						},
						Ulimits: UlimitSlice{
							{
								Name: "foo",
								Soft: 1024,
								Hard: 2048,
							},
						},
					},
					{
						Name: "docker_build",
						Parameters: map[string]interface{}{
							"dry_run":  true,
							"registry": "index.docker.io",
							"repo":     "github/octocat",
							"tags":     []interface{}{"latest", "dev"},
						},
						Image: "plugins/docker:18.09",
						Pull:  "always",
						Ruleset: Ruleset{
							If:       Rules{Event: []string{"push", "pull_request:opened", "pull_request:synchronized"}},
							Matcher:  "filepath",
							Operator: "and",
						},
					},
					{
						Name: "docker_publish",
						Parameters: map[string]interface{}{
							"registry": "index.docker.io",
							"repo":     "github/octocat",
							"tags":     []interface{}{"latest", "dev"},
						},
						Image: "plugins/docker:18.09",
						Pull:  "always",
						Ruleset: Ruleset{
							If:       Rules{Branch: []string{"master"}, Event: []string{"push"}},
							Matcher:  "filepath",
							Operator: "and",
						},
						Secrets: StepSecretSlice{
							{
								Source: "docker_username",
								Target: "plugin_username",
							},
							{
								Source: "docker_password",
								Target: "plugin_password",
							},
						},
					},
				},
				Secrets: SecretSlice{
					{
						Name:   "docker_username",
						Key:    "org/repo/docker/username",
						Engine: "native",
						Type:   "repo",
					},
					{
						Name:   "docker_password",
						Key:    "org/repo/docker/password",
						Engine: "vault",
						Type:   "repo",
					},
					{
						Name:   "docker_username",
						Key:    "org/docker/username",
						Engine: "native",
						Type:   "org",
					},
					{
						Name:   "docker_password",
						Key:    "org/docker/password",
						Engine: "vault",
						Type:   "org",
					},
					{
						Name:   "docker_username",
						Key:    "org/team/docker/username",
						Engine: "native",
						Type:   "shared",
					},
					{
						Name:   "docker_password",
						Key:    "org/team/docker/password",
						Engine: "vault",
						Type:   "shared",
					},
				},
				Templates: TemplateSlice{
					{
						Name:   "docker_publish",
						Source: "github.com/go-vela/atlas/stable/docker_publish",
						Type:   "github",
					},
				},
			},
		},
		{
			file: "testdata/build_anchor_stage.yml",
			want: &Build{
				Version: "1",
				Metadata: Metadata{
					Template:    false,
					Clone:       nil,
					Environment: []string{"steps", "services", "secrets"},
				},
				Stages: StageSlice{
					{
						Name:  "dependencies",
						Needs: []string{"clone"},
						Steps: StepSlice{
							{
								Commands: raw.StringSlice{"./gradlew downloadDependencies"},
								Environment: raw.StringSliceMap{
									"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
									"GRADLE_USER_HOME": ".gradle",
								},
								Image: "openjdk:latest",
								Name:  "install",
								Pull:  "always",
								Ruleset: Ruleset{
									If:       Rules{Event: []string{"push", "pull_request:opened", "pull_request:synchronized"}},
									Matcher:  "filepath",
									Operator: "and",
								},
								Volumes: VolumeSlice{
									{
										Source:      "/foo",
										Destination: "/bar",
										AccessMode:  "ro",
									},
								},
								Ulimits: UlimitSlice{
									{
										Name: "foo",
										Soft: 1024,
										Hard: 2048,
									},
								},
							},
						},
					},
					{
						Name:  "test",
						Needs: []string{"dependencies", "clone"},
						Steps: StepSlice{
							{
								Commands: raw.StringSlice{"./gradlew check"},
								Environment: raw.StringSliceMap{
									"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
									"GRADLE_USER_HOME": ".gradle",
								},
								Name:  "test",
								Image: "openjdk:latest",
								Pull:  "always",
								Ruleset: Ruleset{
									If:       Rules{Event: []string{"push", "pull_request:opened", "pull_request:synchronized"}},
									Matcher:  "filepath",
									Operator: "and",
								},
								Volumes: VolumeSlice{
									{
										Source:      "/foo",
										Destination: "/bar",
										AccessMode:  "ro",
									},
								},
								Ulimits: UlimitSlice{
									{
										Name: "foo",
										Soft: 1024,
										Hard: 2048,
									},
								},
							},
						},
					},
					{
						Name:  "build",
						Needs: []string{"dependencies", "clone"},
						Steps: StepSlice{
							{
								Commands: raw.StringSlice{"./gradlew build"},
								Environment: raw.StringSliceMap{
									"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
									"GRADLE_USER_HOME": ".gradle",
								},
								Name:  "build",
								Image: "openjdk:latest",
								Pull:  "always",
								Ruleset: Ruleset{
									If:       Rules{Event: []string{"push", "pull_request:opened", "pull_request:synchronized"}},
									Matcher:  "filepath",
									Operator: "and",
								},
								Volumes: VolumeSlice{
									{
										Source:      "/foo",
										Destination: "/bar",
										AccessMode:  "ro",
									},
								},
								Ulimits: UlimitSlice{
									{
										Name: "foo",
										Soft: 1024,
										Hard: 2048,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			file: "testdata/build_anchor_step.yml",
			want: &Build{
				Version: "1",
				Metadata: Metadata{
					Template:    false,
					Clone:       nil,
					Environment: []string{"steps", "services", "secrets"},
				},
				Steps: StepSlice{
					{
						Commands: raw.StringSlice{"./gradlew downloadDependencies"},
						Environment: raw.StringSliceMap{
							"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
							"GRADLE_USER_HOME": ".gradle",
						},
						Image: "openjdk:latest",
						Name:  "install",
						Pull:  "always",
						Ruleset: Ruleset{
							If:       Rules{Event: []string{"push", "pull_request:opened", "pull_request:synchronized"}},
							Matcher:  "filepath",
							Operator: "and",
						},
						Volumes: VolumeSlice{
							{
								Source:      "/foo",
								Destination: "/bar",
								AccessMode:  "ro",
							},
						},
						Ulimits: UlimitSlice{
							{
								Name: "foo",
								Soft: 1024,
								Hard: 2048,
							},
						},
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
						Ruleset: Ruleset{
							If:       Rules{Event: []string{"push", "pull_request:opened", "pull_request:synchronized"}},
							Matcher:  "filepath",
							Operator: "and",
						},
						Volumes: VolumeSlice{
							{
								Source:      "/foo",
								Destination: "/bar",
								AccessMode:  "ro",
							},
						},
						Ulimits: UlimitSlice{
							{
								Name: "foo",
								Soft: 1024,
								Hard: 2048,
							},
						},
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
						Ruleset: Ruleset{
							If:       Rules{Event: []string{"push", "pull_request:opened", "pull_request:synchronized"}},
							Matcher:  "filepath",
							Operator: "and",
						},
						Volumes: VolumeSlice{
							{
								Source:      "/foo",
								Destination: "/bar",
								AccessMode:  "ro",
							},
						},
						Ulimits: UlimitSlice{
							{
								Name: "foo",
								Soft: 1024,
								Hard: 2048,
							},
						},
					},
				},
			},
		},
	}

	// run tests
	for _, test := range tests {
		got := new(Build)

		b, err := ioutil.ReadFile(test.file)
		if err != nil {
			t.Errorf("Reading file for UnmarshalYAML returned err: %v", err)
		}

		err = yaml.Unmarshal(b, got)

		if err != nil {
			t.Errorf("UnmarshalYAML returned err: %v", err)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("UnmarshalYAML is %v, want %v", got, test.want)
		}
	}
}
