// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"strings"

	"github.com/go-vela/types/constants"
)

// Build is the pipeline representation of a build for a pipeline.
//
// swagger:model PipelineBuild
type Build struct {
	ID       string         `json:"id,omitempty"       yaml:"id,omitempty"`
	Version  string         `json:"version,omitempty"  yaml:"version,omitempty"`
	Metadata Metadata       `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Worker   Worker         `json:"worker,omitempty"   yaml:"worker,omitempty"`
	Secrets  SecretSlice    `json:"secrets,omitempty"  yaml:"secrets,omitempty"`
	Services ContainerSlice `json:"services,omitempty" yaml:"services,omitempty"`
	Stages   StageSlice     `json:"stages,omitempty"   yaml:"stages,omitempty"`
	Steps    ContainerSlice `json:"steps,omitempty"    yaml:"steps,omitempty"`
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

// Sanitize cleans the fields for every step in each stage so they
// can be safely executed on the worker. If no stages are provided
// in the pipeline, then the function will sanitize the fields for
// every step in the pipeline. The fields are sanitized based off
// of the provided runtime driver which is setup on every worker.
// Currently, this function supports the following runtimes:
//
//   * Docker
//   * Kubernetes
func (b *Build) Sanitize(driver string) *Build {
	// return an empty pipeline if both stages and steps are provided
	if len(b.Stages) > 0 && len(b.Steps) > 0 {
		return nil
	}

	// sanitize stages pipeline if they are provided
	if len(b.Stages) > 0 {
		b.Stages = *b.Stages.Sanitize(driver)
	}

	// sanitize steps pipeline if they are provided
	if len(b.Steps) > 0 {
		b.Steps = *b.Steps.Sanitize(driver)
	}

	// sanitize services pipeline if they are provided
	if len(b.Services) > 0 {
		b.Services = *b.Services.Sanitize(driver)
	}

	// sanitize secret plugins pipeline if they are provided
	for i, secret := range b.Secrets {
		if secret.Origin == nil {
			continue
		}

		b.Secrets[i].Origin = secret.Origin.Sanitize(driver)
	}

	switch driver {
	// sanitize pipeline for Docker
	case constants.DriverDocker:
		if strings.Contains(b.ID, " ") {
			b.ID = strings.ReplaceAll(b.ID, " ", "-")
		}
	// sanitize pipeline for Kubernetes
	case constants.DriverKubernetes:
		if strings.Contains(b.ID, " ") {
			b.ID = strings.ReplaceAll(b.ID, " ", "-")
		}

		if strings.Contains(b.ID, "_") {
			b.ID = strings.ReplaceAll(b.ID, "_", "-")
		}

		if strings.Contains(b.ID, ".") {
			b.ID = strings.ReplaceAll(b.ID, ".", "-")
		}
	}

	// return the purged pipeline
	return b
}
