// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import "strings"

type (
	// ContainerSlice is the pipeline representation
	// of the Containers block for a pipeline.
	ContainerSlice []*Container

	// Container is the pipeline representation
	// of a Container in a pipeline.
	Container struct {
		ID          string            `json:"id,omitempty"`
		Commands    []string          `json:"commands,omitempty"`
		Detach      bool              `json:"detach,omitempty"`
		Directory   string            `json:"directory,omitempty"`
		Entrypoint  []string          `json:"entrypoint,omitempty"`
		Environment map[string]string `json:"environment,omitempty"`
		ExitCode    int               `json:"exit_code,omitempty"`
		Image       string            `json:"image,omitempty"`
		Name        string            `json:"name,omitempty"`
		Needs       []string          `json:"needs,omitempty"`
		Networks    []string          `json:"networks,omitempty"`
		Number      int               `json:"number,omitempty"`
		Ports       []string          `json:"ports,omitempty"`
		Privileged  bool              `json:"privileged,omitempty"`
		Pull        bool              `json:"pull,omitempty"`
		Ruleset     Ruleset           `json:"ruleset,omitempty"`
		Secrets     StepSecretSlice   `json:"secrets,omitempty"`
		Ulimits     UlimitSlice       `json:"ulimits,omitempty"`
		Volumes     VolumeSlice       `json:"volumes,omitempty"`
	}
)

// Purge removes the Containers that have a ruleset
// that do not match the provided ruledata.
func (c *ContainerSlice) Purge(r *RuleData) *ContainerSlice {
	counter := 1
	containers := new(ContainerSlice)

	// iterate through each Container in the pipeline
	for _, container := range *c {

		// verify ruleset matches
		if container.Ruleset.Match(r) {
			// overwrite the Container number with the Container counter
			container.Number = counter

			// increment Container counter
			counter = counter + 1

			// append the Container to the new slice of Containers
			*containers = append(*containers, container)
		}
	}

	// return the new slice of Containers
	return containers
}

// Sanitize cleans the ID for the container which
// allows it to be safely executed on a worker.
func (c *Container) Sanitize() {
	if strings.Contains(c.ID, " ") {
		c.ID = strings.ReplaceAll(c.ID, " ", "-")
	}
}
