// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"

	"github.com/buildkite/yaml"
)

func TestYaml_StageSlice_ToPipeline(t *testing.T) {
	// setup types
	num := int64(1)
	str := "foo"
	slice := []string{"foo"}
	mapp := map[string]string{"foo": "bar"}

	want := &pipeline.StageSlice{
		&pipeline.Stage{
			Name:  str,
			Needs: slice,
			Steps: pipeline.ContainerSlice{
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
			},
		},
	}

	s := &StageSlice{
		&Stage{
			Name:  str,
			Needs: slice,
			Steps: StepSlice{
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
			},
		},
	}

	// run test
	got := s.ToPipeline()

	// WARNING: hack to compare stages
	//
	// Channel values can only be compared for equality.
	// Two channel values are considered equal if they
	// originated from the same make call meaning they
	// refer to the same channel value in memory.
	for i, stage := range *got {
		tmp := *want

		tmp[i].Done = stage.Done
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToPipeline is %v, want %v", got, want)
	}
}

func TestYaml_StageSlice_UnmarshalYAML(t *testing.T) {
	// setup types
	want := &StageSlice{
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
				},
			},
		},
		&Stage{
			Name:  "test",
			Needs: raw.StringSlice{"dependencies"},
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
				},
			},
		},
		&Stage{
			Name:  "build",
			Needs: raw.StringSlice{"dependencies"},
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
				},
			},
		},
	}
	got := new(StageSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/stage.yml")
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

func TestYaml_StageSlice_UnmarshalYAML_Bad(t *testing.T) {
	// run test
	err := yaml.Unmarshal([]byte("- foo"), new(StageSlice))

	if err == nil {
		t.Errorf("UnmarshalYAML should have returned err")
	}
}

func TestYaml_StageSlice_UnmarshalYAML_Invalid(t *testing.T) {
	// run test
	b, err := ioutil.ReadFile("testdata/invalid.yml")
	if err != nil {
		t.Errorf("Reading file for UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, new(StageSlice))

	if err == nil {
		t.Errorf("UnmarshalYAML should have returned err")
	}
}
