// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

type (
	// VolumeSlice is the pipeline representation of
	// the volumes block for a step in a pipeline.
	VolumeSlice []*Volume

	// Volume is the pipeline representation of a volume
	// from a volumes block for a step in a pipeline.
	Volume struct {
		Source      string `json:"source,omitempty"`
		Destination string `json:"destination,omitempty"`
		AccessMode  string `json:"access_mode,omitempty"`
	}
)
