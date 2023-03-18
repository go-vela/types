// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

type (
	// InitStepSlice is the pipeline representation
	// of the InitSteps block for a pipeline.
	//
	// swagger:model PipelineInitSlice
	//
	// swagger:model PipelineInitSlice
	InitStepSlice []*InitStep

	// InitStep is the pipeline representation of an init step in a pipeline.
	//
	// An InitStep allows logs to be associated with something other than a container.
	//
	// swagger:model PipelineInit
	InitStep struct {
		ID       string `json:"id,omitempty"         yaml:"id,omitempty"`
		Number   int    `json:"number,omitempty"     yaml:"number,omitempty"`
		Reporter string `json:"reporter,omitempty"   yaml:"reporter,omitempty"`
		Name     string `json:"name,omitempty"       yaml:"name,omitempty"`
	}
)
