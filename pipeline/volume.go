// SPDX-License-Identifier: Apache-2.0

package pipeline

type (
	// VolumeSlice is the pipeline representation of
	// the volumes block for a step in a pipeline.
	//
	// Deprecated: use VolumeSlice from github.com/go-vela/server/compiler/types/pipeline instead.
	VolumeSlice []*Volume

	// Volume is the pipeline representation of a volume
	// from a volumes block for a step in a pipeline.
	//
	// Deprecated: use Volume from github.com/go-vela/server/compiler/types/pipeline instead.
	Volume struct {
		Source      string `json:"source,omitempty"      yaml:"source,omitempty"`
		Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`
		AccessMode  string `json:"access_mode,omitempty" yaml:"access_mode,omitempty"`
	}
)
