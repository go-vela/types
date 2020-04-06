// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"github.com/go-vela/types/constants"
)

type (
	// StageSlice is the pipeline representation
	// of the stages block for a pipeline.
	StageSlice []*Stage

	// Stage is the pipeline representation
	// of a stage in a pipeline.
	Stage struct {
		Done  chan error     `json:"-"               yaml:"-"`
		Name  string         `json:"name,omitempty"  yaml:"name,omitempty"`
		Needs []string       `json:"needs,omitempty" yaml:"needs,omitempty"`
		Steps ContainerSlice `json:"steps,omitempty" yaml:"steps,omitempty"`
	}
)

// Purge removes the steps, from the stages, that have
// a ruleset that do not match the provided ruledata.
// If all steps from a stage are removed, then the
// entire stage is removed from the pipeline.
func (s *StageSlice) Purge(r *RuleData) *StageSlice {
	counter := 1
	stages := new(StageSlice)

	// iterate through each stage for the pipeline
	for _, stage := range *s {
		containers := new(ContainerSlice)

		// iterate through each step for the stage in the pipeline
		for _, step := range stage.Steps {

			// verify ruleset matches
			if step.Ruleset.Match(r) {
				// overwrite the step number with the step counter
				step.Number = counter

				// increment step counter
				counter = counter + 1

				// append the step to the new slice of containers
				*containers = append(*containers, step)
			}
		}

		// no steps for the stage so we continue processing to the next stage
		if len(*containers) == 0 {
			continue
		}

		// overwrite the steps for the stage with the new slice of steps
		stage.Steps = *containers

		// append the stage to the new slice of stages
		*stages = append(*stages, stage)
	}

	// return the new slice of stages
	return stages
}

// Sanitize cleans the fields for every step in each stage so they
// can be safely executed on the worker. The fields are sanitized
// based off of the provided runtime driver which is setup on every
// worker. Currently, this function supports the following runtimes:
//
//   * Docker
//   * Kubernetes
func (s *StageSlice) Sanitize(driver string) *StageSlice {
	stages := new(StageSlice)

	switch driver {
	// sanitize container for Docker
	case constants.DriverDocker:
		for _, stage := range *s {
			stage.Steps.Sanitize(driver)

			*stages = append(*stages, stage)
		}

		return stages
	// sanitize container for Kubernetes
	case constants.DriverKubernetes:
		for _, stage := range *s {
			stage.Steps.Sanitize(driver)

			*stages = append(*stages, stage)
		}

		return stages
	// unrecognized driver
	default:
		// log here?
		return nil
	}
}
