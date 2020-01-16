// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"
)

type (
	// StepSlice is the yaml representation
	// of the steps block for a pipeline.
	StepSlice []*Step

	// Step is the yaml representation of a step
	// from the steps block for a pipeline.
	Step struct {
		Commands    raw.StringSlice        `yaml:"commands,omitempty"`
		Detach      bool                   `yaml:"detach,omitempty"`
		Entrypoint  raw.StringSlice        `yaml:"entrypoint,omitempty"`
		Environment raw.StringSliceMap     `yaml:"environment,omitempty"`
		Image       string                 `yaml:"image,omitempty"`
		Name        string                 `yaml:"name,omitempty"`
		Parameters  map[string]interface{} `yaml:"parameters,omitempty"`
		Privileged  bool                   `yaml:"privileged,omitempty"`
		Pull        bool                   `yaml:"pull,omitempty"`
		Ruleset     Ruleset                `yaml:"ruleset,omitempty"`
		Secrets     StepSecretSlice        `yaml:"secrets,omitempty"`
		Template    StepTemplate           `yaml:"template,omitempty"`
		Ulimits     UlimitSlice            `yaml:"ulimits,omitempty"`
		Volumes     VolumeSlice            `yaml:"volumes,omitempty"`
	}
)

// ToPipeline converts the StepSlice type
// to a pipeline ContainerSlice type.
func (s *StepSlice) ToPipeline() *pipeline.ContainerSlice {
	// step slice we want to return
	stepSlice := new(pipeline.ContainerSlice)

	// iterate through each element in the step slice
	for _, step := range *s {
		// append the element to the pipeline container slice
		*stepSlice = append(*stepSlice, &pipeline.Container{
			Commands:    step.Commands,
			Detach:      step.Detach,
			Entrypoint:  step.Entrypoint,
			Environment: step.Environment,
			Image:       step.Image,
			Name:        step.Name,
			Privileged:  step.Privileged,
			Pull:        step.Pull,
			Ruleset:     *step.Ruleset.ToPipeline(),
			Secrets:     *step.Secrets.ToPipeline(),
			Ulimits:     *step.Ulimits.ToPipeline(),
			Volumes:     *step.Volumes.ToPipeline(),
		})
	}

	return stepSlice
}

// UnmarshalYAML implements the Unmarshaler interface for the StepSlice type.
func (s *StepSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// step slice we try unmarshaling to
	stepSlice := new([]*Step)

	// attempt to unmarshal as a step slice type
	err := unmarshal(stepSlice)
	if err != nil {
		return err
	}

	// overwrite existing StepSlice
	*s = StepSlice(*stepSlice)

	return nil
}
