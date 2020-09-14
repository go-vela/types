// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"strings"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"
)

type (
	// StepSlice is the yaml representation
	// of the steps block for a pipeline.
	StepSlice []*Step

	// Step is the yaml representation of a step
	// from the steps block for a pipeline.
	// nolint:lll // jsonschema will cause long lines
	Step struct {
		Commands    raw.StringSlice        `yaml:"commands,omitempty"    json:"commands,omitempty" jsonschema:"description=Execution instructions to run inside the container.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/commands/"`
		Detach      bool                   `yaml:"detach,omitempty"      json:"detach,omitempty" jsonschema:"description=Run the container in a detached (headless) state.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/detach/"`
		Entrypoint  raw.StringSlice        `yaml:"entrypoint,omitempty"  json:"entrypoint,omitempty" jsonschema:"description=Command to execute inside the container.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/entrypoint/"`
		Environment raw.StringSliceMap     `yaml:"environment,omitempty" json:"environment,omitempty" jsonschema:"description=Provide environment variables injected into the container environment.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/environment/"`
		Image       string                 `yaml:"image,omitempty"       json:"image,omitempty" jsonschema:"oneof_required=image,minLength=1,description=Docker image to use to create the ephemeral container.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/image/"`
		Name        string                 `yaml:"name,omitempty"        json:"name,omitempty" jsonschema:"required,minLength=1,description=Unique name for the step."`
		Parameters  map[string]interface{} `yaml:"parameters,omitempty"  json:"parameters,omitempty" jsonschema:"description=Extra configuration variables for a plugin.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/parameters/"`
		Privileged  bool                   `yaml:"privileged,omitempty"  json:"privileged,omitempty" jsonschema:"description=Run the container with extra privileges.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/privileged/"`
		Pull        string                 `yaml:"pull,omitempty"        json:"pull,omitempty" jsonschema:"description=Automatically upgrade to the latest version of the image.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/pull/"`
		Ruleset     Ruleset                `yaml:"ruleset,omitempty"     json:"ruleset,omitempty" jsonschema:"description=Conditions to limit the execution of the container.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/ruleset/"`
		Secrets     StepSecretSlice        `yaml:"secrets,omitempty"     json:"secrets,omitempty" jsonschema:"description=Sensitive variables injected into the container environment.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/secrets/"`
		Template    StepTemplate           `yaml:"template,omitempty"    json:"template,omitempty" jsonschema:"oneof_required=template,description=Name of template to expand in the pipeline.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/template/"`
		Ulimits     UlimitSlice            `yaml:"ulimits,omitempty"     json:"ulimits,omitempty" jsonschema:"description=Set the user limits for the container.\nReference: coming soon"`
		Volumes     VolumeSlice            `yaml:"volumes,omitempty"     json:"volumes,omitempty" jsonschema:"description=Mount volumes for the container.\nReference: coming soon"`
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
	// step slice we try unmarshalling to
	stepSlice := new([]*Step)

	// attempt to unmarshal as a step slice type
	err := unmarshal(stepSlice)
	if err != nil {
		return err
	}

	// iterate through each element in the step slice
	for _, step := range *stepSlice {
		// implicitly set `pull` field if empty
		if len(step.Pull) == 0 {
			step.Pull = constants.PullNotPresent
		}

		// TODO: remove this in a future release
		//
		// handle true deprecated pull policy
		//
		// a `true` pull policy equates to `always`
		if strings.EqualFold(step.Pull, "true") {
			step.Pull = constants.PullAlways
		}

		// TODO: remove this in a future release
		//
		// handle false deprecated pull policy
		//
		// a `false` pull policy equates to `not_present`
		if strings.EqualFold(step.Pull, "false") {
			step.Pull = constants.PullNotPresent
		}
	}

	// overwrite existing StepSlice
	*s = StepSlice(*stepSlice)

	return nil
}
