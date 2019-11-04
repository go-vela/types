// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

type (
	// UlimitSlice is the pipeline representation of
	// the ulimits block for a step in a pipeline.
	UlimitSlice []*Ulimit

	// Ulimit is the pipeline representation of a ulimit
	// from the ulimits block for a step in a pipeline.
	Ulimit struct {
		Name string `json:"name,omitempty"`
		Soft int64  `json:"soft,omitempty"`
		Hard int64  `json:"hard,omitempty"`
	}
)
