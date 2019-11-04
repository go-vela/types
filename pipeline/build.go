// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

// Build is the pipeline representation of a build for a pipeline.
type Build struct {
	ID       string         `json:"id,omitempty"`
	Version  string         `json:"version,omitempty"`
	Metadata Metadata       `json:"metadata,omitempty"`
	Secrets  SecretSlice    `json:"secrets,omitempty"`
	Services ContainerSlice `json:"services,omitempty"`
	Stages   StageSlice     `json:"stages,omitempty"`
	Steps    ContainerSlice `json:"steps,omitempty"`
}

// Purge removes the steps, in every stage, that contain a ruleset
// that do not match the provided ruledata. If all steps from a
// stage are removed, then the entire stage is removed from the
// pipeline. If no stages are provided in the pipeline, then the
// function will remove the steps that have a ruleset that do not
// match the provided ruledata. If both stages and steps are
// provided, then an empty pipeline is returned.
func (b *Build) Purge(r *RuleData) *Build {
	// return an empty pipeline if both stages and steps are provided
	if len(b.Stages) > 0 && len(b.Steps) > 0 {
		return nil
	}

	// purge stages pipeline if stages are provided
	if len(b.Stages) > 0 {
		b.Stages = *b.Stages.Purge(r)
	}

	// purge steps pipeline if steps are provided
	if len(b.Steps) > 0 {
		b.Steps = *b.Steps.Purge(r)
	}

	// purge services pipeline if services are provided
	if len(b.Services) > 0 {
		b.Services = *b.Services.Purge(r)
	}

	// return the purged pipeline
	return b
}
