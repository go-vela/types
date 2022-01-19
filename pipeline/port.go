// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

type (
	// PortSlice is the pipeline representation
	// of the ports for a step in a pipeline.
	PortSlice []*Port

	// Port is the pipeline representation
	// of a port for a step in a pipeline.
	//
	// swagger:model PipelinePort
	Port struct {
		Port     int    `json:"port,omitempty"     yaml:"port,omitempty"`
		Host     int    `json:"host,omitempty"     yaml:"host,omitempty"`
		Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	}
)
