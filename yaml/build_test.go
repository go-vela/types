// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/go-vela/types/raw"

	"github.com/buildkite/yaml"
)

func TestYaml_Build_UnmarshalYAML(t *testing.T) {
	// setup types
	want := &Build{
		Version: "1",
		Metadata: Metadata{
			Template: false,
		},
		Worker: Worker{
			Flavor:   "16cpu8gb",
			Platform: "gcp",
		},
		Services: ServiceSlice{
			&Service{
				Ports: []string{"5432:5432"},
				Environment: raw.StringSliceMap{
					"POSTGRES_DB": "foo",
				},
				Name:  "postgres",
				Image: "postgres:latest",
			},
		},
		Steps: StepSlice{
			&Step{
				Commands: raw.StringSlice{"./gradlew downloadDependencies"},
				Environment: raw.StringSliceMap{
					"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
					"GRADLE_USER_HOME": ".gradle",
				},
				Image: "openjdk:latest",
				Name:  "install",
				Pull:  true,
				Ruleset: Ruleset{
					If:       Rules{Event: []string{"push", "pull_request"}},
					Operator: "and",
				},
				Volumes: VolumeSlice{
					&Volume{
						Source:      "/foo",
						Destination: "/bar",
						AccessMode:  "ro",
					},
				},
				Ulimits: UlimitSlice{
					&Ulimit{
						Name: "foo",
						Soft: 1024,
						Hard: 2048,
					},
				},
			},
			&Step{
				Commands: raw.StringSlice{"./gradlew check"},
				Environment: raw.StringSliceMap{
					"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
					"GRADLE_USER_HOME": ".gradle",
				},
				Name:  "test",
				Image: "openjdk:latest",
				Pull:  true,
				Ruleset: Ruleset{
					If:       Rules{Event: []string{"push", "pull_request"}},
					Operator: "and",
				},
				Volumes: VolumeSlice{
					&Volume{
						Source:      "/foo",
						Destination: "/bar",
						AccessMode:  "ro",
					},
				},
				Ulimits: UlimitSlice{
					&Ulimit{
						Name: "foo",
						Soft: 1024,
						Hard: 2048,
					},
				},
			},
			&Step{
				Commands: raw.StringSlice{"./gradlew build"},
				Environment: raw.StringSliceMap{
					"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
					"GRADLE_USER_HOME": ".gradle",
				},
				Name:  "build",
				Image: "openjdk:latest",
				Pull:  true,
				Ruleset: Ruleset{
					If:       Rules{Event: []string{"push", "pull_request"}},
					Operator: "and",
				},
				Volumes: VolumeSlice{
					&Volume{
						Source:      "/foo",
						Destination: "/bar",
						AccessMode:  "ro",
					},
				},
				Ulimits: UlimitSlice{
					&Ulimit{
						Name: "foo",
						Soft: 1024,
						Hard: 2048,
					},
				},
			},
			&Step{
				Name: "docker_build",
				Parameters: map[string]interface{}{
					"dry_run":  true,
					"registry": "index.docker.io",
					"repo":     "github/octocat",
					"tags":     []interface{}{"latest", "dev"},
				},
				Image: "plugins/docker:18.09",
				Pull:  true,
				Ruleset: Ruleset{
					If:       Rules{Event: []string{"push", "pull_request"}},
					Operator: "and",
				},
			},
			&Step{
				Name: "docker_publish",
				Parameters: map[string]interface{}{
					"registry": "index.docker.io",
					"repo":     "github/octocat",
					"tags":     []interface{}{"latest", "dev"},
				},
				Image: "plugins/docker:18.09",
				Pull:  true,
				Ruleset: Ruleset{
					If:       Rules{Branch: []string{"master"}, Event: []string{"push"}},
					Operator: "and",
				},
				Secrets: StepSecretSlice{
					&StepSecret{
						Source: "docker_username",
						Target: "plugin_username",
					},
					&StepSecret{
						Source: "docker_password",
						Target: "plugin_password",
					},
				},
			},
		},
		Secrets: SecretSlice{
			&Secret{
				Name:   "docker_username",
				Key:    "org/repo/docker/username",
				Engine: "native",
				Type:   "repo",
			},
			&Secret{
				Name:   "docker_password",
				Key:    "org/repo/docker/password",
				Engine: "vault",
				Type:   "repo",
			},
			&Secret{
				Name:   "docker_username",
				Key:    "org/docker/username",
				Engine: "native",
				Type:   "org",
			},
			&Secret{
				Name:   "docker_password",
				Key:    "org/docker/password",
				Engine: "vault",
				Type:   "org",
			},
			&Secret{
				Name:   "docker_username",
				Key:    "org/team/docker/username",
				Engine: "native",
				Type:   "shared",
			},
			&Secret{
				Name:   "docker_password",
				Key:    "org/team/docker/password",
				Engine: "vault",
				Type:   "shared",
			},
		},
		Templates: TemplateSlice{
			&Template{
				Name:   "docker_publish",
				Source: "github.com/go-vela/atlas/stable/docker_publish",
				Type:   "github",
			},
		},
	}
	got := new(Build)

	// run test
	b, err := ioutil.ReadFile("testdata/build.yml")
	if err != nil {
		t.Errorf("Reading file for UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err != nil {
		t.Errorf("UnmarshalYAML returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalYAML is %v, want %v", got, want)
	}
}

func TestYaml_Build_UnmarshalYAML_AnchorStage(t *testing.T) {
	// setup types
	want := &Build{
		Version: "1",
		Metadata: Metadata{
			Template: false,
		},
		Stages: StageSlice{
			&Stage{
				Name:  "dependencies",
				Needs: []string{"clone"},
				Steps: StepSlice{
					&Step{
						Commands: raw.StringSlice{"./gradlew downloadDependencies"},
						Environment: raw.StringSliceMap{
							"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
							"GRADLE_USER_HOME": ".gradle",
						},
						Image: "openjdk:latest",
						Name:  "install",
						Pull:  true,
						Ruleset: Ruleset{
							If:       Rules{Event: []string{"push", "pull_request"}},
							Operator: "and",
						},
						Volumes: VolumeSlice{
							&Volume{
								Source:      "/foo",
								Destination: "/bar",
								AccessMode:  "ro",
							},
						},
						Ulimits: UlimitSlice{
							&Ulimit{
								Name: "foo",
								Soft: 1024,
								Hard: 2048,
							},
						},
					},
				},
			},
			&Stage{
				Name:  "test",
				Needs: []string{"dependencies"},
				Steps: StepSlice{
					&Step{
						Commands: raw.StringSlice{"./gradlew check"},
						Environment: raw.StringSliceMap{
							"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
							"GRADLE_USER_HOME": ".gradle",
						},
						Name:  "test",
						Image: "openjdk:latest",
						Pull:  true,
						Ruleset: Ruleset{
							If:       Rules{Event: []string{"push", "pull_request"}},
							Operator: "and",
						},
						Volumes: VolumeSlice{
							&Volume{
								Source:      "/foo",
								Destination: "/bar",
								AccessMode:  "ro",
							},
						},
						Ulimits: UlimitSlice{
							&Ulimit{
								Name: "foo",
								Soft: 1024,
								Hard: 2048,
							},
						},
					},
				},
			},
			&Stage{
				Name:  "build",
				Needs: []string{"dependencies"},
				Steps: StepSlice{
					&Step{
						Commands: raw.StringSlice{"./gradlew build"},
						Environment: raw.StringSliceMap{
							"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
							"GRADLE_USER_HOME": ".gradle",
						},
						Name:  "build",
						Image: "openjdk:latest",
						Pull:  true,
						Ruleset: Ruleset{
							If:       Rules{Event: []string{"push", "pull_request"}},
							Operator: "and",
						},
						Volumes: VolumeSlice{
							&Volume{
								Source:      "/foo",
								Destination: "/bar",
								AccessMode:  "ro",
							},
						},
						Ulimits: UlimitSlice{
							&Ulimit{
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
	got := new(Build)

	// run test
	b, err := ioutil.ReadFile("testdata/build_anchor_stage.yml")
	if err != nil {
		t.Errorf("Reading file for UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err != nil {
		t.Errorf("UnmarshalYAML returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalYAML is %v, want %v", got, want)
	}
}

func TestYaml_Build_UnmarshalYAML_AnchorSteps(t *testing.T) {
	// setup types
	want := &Build{
		Version: "1",
		Metadata: Metadata{
			Template: false,
		},
		Steps: StepSlice{
			&Step{
				Commands: raw.StringSlice{"./gradlew downloadDependencies"},
				Environment: raw.StringSliceMap{
					"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
					"GRADLE_USER_HOME": ".gradle",
				},
				Image: "openjdk:latest",
				Name:  "install",
				Pull:  true,
				Ruleset: Ruleset{
					If:       Rules{Event: []string{"push", "pull_request"}},
					Operator: "and",
				},
				Volumes: VolumeSlice{
					&Volume{
						Source:      "/foo",
						Destination: "/bar",
						AccessMode:  "ro",
					},
				},
				Ulimits: UlimitSlice{
					&Ulimit{
						Name: "foo",
						Soft: 1024,
						Hard: 2048,
					},
				},
			},
			&Step{
				Commands: raw.StringSlice{"./gradlew check"},
				Environment: raw.StringSliceMap{
					"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
					"GRADLE_USER_HOME": ".gradle",
				},
				Name:  "test",
				Image: "openjdk:latest",
				Pull:  true,
				Ruleset: Ruleset{
					If:       Rules{Event: []string{"push", "pull_request"}},
					Operator: "and",
				},
				Volumes: VolumeSlice{
					&Volume{
						Source:      "/foo",
						Destination: "/bar",
						AccessMode:  "ro",
					},
				},
				Ulimits: UlimitSlice{
					&Ulimit{
						Name: "foo",
						Soft: 1024,
						Hard: 2048,
					},
				},
			},
			&Step{
				Commands: raw.StringSlice{"./gradlew build"},
				Environment: raw.StringSliceMap{
					"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
					"GRADLE_USER_HOME": ".gradle",
				},
				Name:  "build",
				Image: "openjdk:latest",
				Pull:  true,
				Ruleset: Ruleset{
					If:       Rules{Event: []string{"push", "pull_request"}},
					Operator: "and",
				},
				Volumes: VolumeSlice{
					&Volume{
						Source:      "/foo",
						Destination: "/bar",
						AccessMode:  "ro",
					},
				},
				Ulimits: UlimitSlice{
					&Ulimit{
						Name: "foo",
						Soft: 1024,
						Hard: 2048,
					},
				},
			},
		},
	}
	got := new(Build)

	// run test
	b, err := ioutil.ReadFile("testdata/build_anchor_step.yml")
	if err != nil {
		t.Errorf("Reading file for UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err != nil {
		t.Errorf("UnmarshalYAML returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalYAML is %v, want %v", got, want)
	}
}
