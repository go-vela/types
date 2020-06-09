// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

type (
	// VolumeSlice is the pipeline representation of
	// the volumes block for a step in a pipeline.
	//
	// swagger:model PipelineVolumeSlice
	VolumeSlice []*Volume

	// Volume is the pipeline representation of a volume
	// from a volumes block for a step in a pipeline.
	//
	// swagger:model PipelineVolume
	Volume struct {
		Source      string `json:"source,omitempty"      yaml:"source,omitempty"`
		Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`
		AccessMode  string `json:"access_mode,omitempty" yaml:"access_mode,omitempty"`
	}
)
