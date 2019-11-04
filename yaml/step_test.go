// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"

	yaml "gopkg.in/yaml.v2"
)

func TestYaml_StepSlice_ToPipeline(t *testing.T) {
	// setup types
	num := int64(1)
	str := "foo"
	slice := []string{"foo"}
	mapp := map[string]string{"foo": "bar"}
	want := &pipeline.ContainerSlice{
		&pipeline.Container{
			Commands:    slice,
			Detach:      false,
			Entrypoint:  slice,
			Environment: mapp,
			Image:       str,
			Name:        str,
			Privileged:  false,
			Pull:        false,
			Ruleset: pipeline.Ruleset{
				If: pipeline.Rules{
					Branch: slice,
					Event:  slice,
					Path:   slice,
					Repo:   slice,
					Status: slice,
					Tag:    slice,
				},
				Unless: pipeline.Rules{
					Branch: slice,
					Event:  slice,
					Path:   slice,
					Repo:   slice,
					Status: slice,
					Tag:    slice,
				},
				Operator: str,
				Continue: false,
			},
			Secrets: pipeline.StepSecretSlice{
				&pipeline.StepSecret{
					Source: str,
					Target: str,
				},
			},
			Ulimits: pipeline.UlimitSlice{
				&pipeline.Ulimit{
					Name: str,
					Soft: num,
					Hard: num,
				},
			},
			Volumes: pipeline.VolumeSlice{
				&pipeline.Volume{
					Source:      str,
					Destination: str,
					AccessMode:  str,
				},
			},
		},
	}

	s := &StepSlice{
		&Step{
			Commands:    slice,
			Detach:      false,
			Entrypoint:  slice,
			Environment: mapp,
			Image:       str,
			Name:        str,
			Privileged:  false,
			Pull:        false,
			Ruleset: Ruleset{
				If: Rules{
					Branch: slice,
					Event:  slice,
					Path:   slice,
					Repo:   slice,
					Status: slice,
					Tag:    slice,
				},
				Unless: Rules{
					Branch: slice,
					Event:  slice,
					Path:   slice,
					Repo:   slice,
					Status: slice,
					Tag:    slice,
				},
				Operator: str,
				Continue: false,
			},
			Secrets: StepSecretSlice{
				&StepSecret{
					Source: str,
					Target: str,
				},
			},
			Ulimits: UlimitSlice{
				&Ulimit{
					Name: str,
					Soft: num,
					Hard: num,
				},
			},
			Volumes: VolumeSlice{
				&Volume{
					Source:      str,
					Destination: str,
					AccessMode:  str,
				},
			},
		},
	}

	// run test
	got := s.ToPipeline()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToPipeline is %v, want %v", got, want)
	}
}

func TestYaml_StepSlice_UnmarshalYAML(t *testing.T) {
	// setup types
	want := &StepSlice{
		&Step{
			Commands: raw.StringSlice{"./gradlew downloadDependencies"},
			Environment: raw.StringSliceMap{
				"GRADLE_OPTS":      "-Dorg.gradle.daemon=false -Dorg.gradle.workers.max=1 -Dorg.gradle.parallel=false",
				"GRADLE_USER_HOME": ".gradle",
			},
			Name:  "install",
			Image: "openjdk:latest",
			Pull:  true,
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
		},
		&Step{
			Name:  "docker_build",
			Image: "plugins/docker:18.09",
			Pull:  true,
			Parameters: map[string]interface{}{
				"registry": "index.docker.io",
				"repo":     "github/octocat",
				"tags":     []interface{}{"latest", "dev"},
			},
		},
		&Step{
			Name: "templated_publish",
			Template: StepTemplate{
				Name: "docker_publish",
				Variables: map[string]interface{}{
					"registry": "index.docker.io",
					"repo":     "github/octocat",
					"tags":     []interface{}{"latest", "dev"},
				},
			},
		},
	}
	got := new(StepSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/step.yml")
	if err != nil {
		t.Errorf("Reading file for Ruleset UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err != nil {
		t.Errorf("YamlSlice UnmarshalYAML returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("YamlSlice UnmarshalYAML is %v, want %v", got, want)
	}
}

func TestYaml_StepSlice_UnmarshalYAML_Invalid(t *testing.T) {
	// run test
	b, err := ioutil.ReadFile("testdata/invalid.yml")
	if err != nil {
		t.Errorf("Reading file for YamlSlice UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, new(StepSlice))

	if err == nil {
		t.Errorf("YamlSlice UnmarshalYAML should have returned err")
	}
}
