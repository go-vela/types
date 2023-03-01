// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

type (
	// InitSlice is the pipeline representation
	// of the Inits block for a pipeline.
	//
	// swagger:model PipelineInitSlice
	//
	// swagger:model PipelineInitSlice
	InitSlice []*Init

	// Init is the pipeline representation of an Init entry in a pipeline.
	//
	// An Init allows logs to be associated with something other than a container.
	//
	// swagger:model PipelineInit
	Init struct {
		ID       string `json:"id,omitempty"         yaml:"id,omitempty"`
		Name     string `json:"name,omitempty"       yaml:"name,omitempty"`
		Number   int    `json:"number,omitempty"     yaml:"number,omitempty"`
		Reporter string `json:"reporter,omitempty"   yaml:"reporter,omitempty"`
		Mimetype string `json:"mimetype,omitempty"   yaml:"mimetype,omitempty"`
	}
)
