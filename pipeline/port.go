// SPDX-License-Identifier: Apache-2.0

package pipeline

type (
	// PortSlice is the pipeline representation
	// of the ports for a step in a pipeline.
	//
	// Deprecated: use PortSlice from github.com/go-vela/server/compiler/types/pipeline instead.
	PortSlice []*Port

	// Port is the pipeline representation
	// of a port for a step in a pipeline.
	//
	// Deprecated: use Port from github.com/go-vela/server/compiler/types/pipeline instead.
	Port struct {
		Port     int    `json:"port,omitempty"     yaml:"port,omitempty"`
		Host     int    `json:"host,omitempty"     yaml:"host,omitempty"`
		Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	}
)
