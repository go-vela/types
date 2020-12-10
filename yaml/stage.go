// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"

	"github.com/goccy/go-yaml"
)

type (
	// StageSlice is the yaml representation
	// of the stages block for a pipeline.
	StageSlice []*Stage

	// Stage is the yaml representation
	// of a stage in a pipeline.
	// nolint:lll // jsonschema will cause long lines
	Stage struct {
		Name  string          `yaml:"name,omitempty"  json:"name,omitempty"  jsonschema:"minLength=1,description=Unique identifier for the stage in the pipeline.\nReference: https://go-vela.github.io/docs/concepts/pipeline/stages/"`
		Needs raw.StringSlice `yaml:"needs,omitempty" json:"needs,omitempty" jsonschema:"description=Stages that must complete before starting the current one.\nReference: https://go-vela.github.io/docs/concepts/pipeline/stages/needs/"`
		Steps StepSlice       `yaml:"steps,omitempty" json:"steps,omitempty" jsonschema:"required,description=Sequential execution instructions for the stage.\nReference: https://go-vela.github.io/docs/concepts/pipeline/stages/steps/"`
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
			Done:  make(chan error, 1),
			Name:  stage.Name,
			Needs: stage.Needs,
			Steps: *stage.Steps.ToPipeline(),
		})
	}

	return stageSlice
}

// UnmarshalYAML implements the Unmarshaler interface for the StageSlice type.
func (s *StageSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// map slice we try unmarshalling to
	mapSlice := new(yaml.MapSlice)

	// attempt to unmarshal as a map slice type
	err := unmarshal(mapSlice)
	if err != nil {
		return err
	}

	// iterate through each element in the map slice
	for _, v := range *mapSlice {
		// stage we try unmarshalling to
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

// Validate lints if the stages configuration is valid.
func (s *StageSlice) Validate(pipeline []byte) error {
	invalid := errors.New("invalid stage block found")

	// iterate through each stage and linting yaml tags
	for _, stage := range *s {
		// check required fields
		if len(stage.Name) == 0 {
			path, err := yaml.PathString("$.stages")
			if err != nil {
				return err
			}
			source, err := path.AnnotateSource(pipeline, true)
			if err != nil {
				return err
			}

			// nolint:cSpell // ignore line length
			invalid = fmt.Errorf("%w: %s", invalid, fmt.Sprintf("no name provided for stage:\n%s\n ", string(source)))
		}

		// extract the raw yaml steps from pipeline
		path, err := yaml.PathString(fmt.Sprintf("$.stages.%s", stage.Name))
		if err != nil {
			return err
		}

		stage := &Stage{}

		// read steps within stage
		err = path.Read(strings.NewReader(string(pipeline)), stage)
		if err != nil {
			return err
		}

		// convert the steps slice into bytes
		raw, err := yaml.Marshal(stage)
		if err != nil {
			return err
		}

		// check steps within that stage
		err = stage.Steps.Validate(raw)
		if err != nil {
			invalid = fmt.Errorf("%w: %v", invalid, err)
		}
	}

	// check if only default error exists
	if !strings.EqualFold(invalid.Error(), "invalid stage block found") {
		return invalid
	}

	return nil
}
