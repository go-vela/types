// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"fmt"

	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"

	"github.com/buildkite/yaml"
)

type (
	// StageSlice is the yaml representation
	// of the stages block for a pipeline.
	StageSlice []*Stage

	// Stage is the yaml representation
	// of a stage in a pipeline.
	Stage struct {
		Name  string          `yaml:"name,omitempty"`
		Needs raw.StringSlice `yaml:"needs,omitempty"`
		Steps StepSlice       `yaml:"steps,omitempty"`
	}
)

// ToPipeline converts the StageSlice type
// to a pipeline StageSlice type.
func (s *StageSlice) ToPipeline() *pipeline.StageSlice {
	// stage slice we want to return
	stageSlice := new(pipeline.StageSlice)

	// iterate through each element in the stage slice
	for _, stage := range *s {
		// append the element to the pipeline stage slice
		*stageSlice = append(*stageSlice, &pipeline.Stage{
			Name:  stage.Name,
			Needs: stage.Needs,
			Steps: *stage.Steps.ToPipeline(),
		})
	}

	return stageSlice
}

// UnmarshalYAML implements the Unmarshaler interface for the StageSlice type.
func (s *StageSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// map slice we try unmarshaling to
	mapSlice := new(yaml.MapSlice)

	// attempt to unmarshal as a map slice type
	err := unmarshal(mapSlice)
	if err != nil {
		return err
	}

	// iterate through each element in the map slice
	for _, v := range *mapSlice {
		// stage we try unmarshaling to
		stage := new(Stage)

		// marshal interface value from ordered map
		out, _ := yaml.Marshal(v.Value)

		// unmarshal interface value as stage
		err = yaml.Unmarshal(out, stage)
		if err != nil {
			return err
		}

		// implicitly set stage `name` if empty
		if len(stage.Name) == 0 {
			stage.Name = fmt.Sprintf("%v", v.Key)
		}

		// implicitly set the stage `needs` if empty
		if len(stage.Needs) == 0 {
			stage.Needs = []string{"clone"}
		}

		// append stage to stage slice
		*s = append(*s, stage)
	}

	return nil
}
