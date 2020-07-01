// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

// Metadata is the yaml representation of the metadata block for a pipeline.
//
// swagger:model PipelineMetadata
type Metadata struct {
	Template bool `json:"template,omitempty" yaml:"template,omitempty"`
}
