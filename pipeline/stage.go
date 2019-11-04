// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

type (
	// StageSlice is the pipeline representation
	// of the stages block for a pipeline.
	StageSlice []*Stage

	// Stage is the pipeline representation
	// of a stage in a pipeline.
	Stage struct {
		Name  string         `json:"name,omitempty"`
		Needs []string       `json:"needs,omitempty"`
		Steps ContainerSlice `json:"steps,omitempty"`
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
