// SPDX-License-Identifier: Apache-2.0

package pipeline

// Metadata is the pipeline representation of the metadata block for a pipeline.
//
// swagger:model PipelineMetadata
type Metadata struct {
	Template    bool     `json:"template,omitempty" yaml:"template,omitempty"`
	Clone       bool     `json:"clone,omitempty" yaml:"clone,omitempty"`
	Environment []string `json:"environment,omitempty" yaml:"environment,omitempty"`
}
