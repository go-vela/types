// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import "github.com/go-vela/types/pipeline"

// Metadata is the yaml representation of
// the metadata block for a pipeline.
// nolint:lll // jsonschema will cause long lines
type Metadata struct {
	Template bool  `yaml:"template,omitempty" json:"template,omitempty" jsonschema:"description=Enables compiling the pipeline as a template.\nReference: https://go-vela.github.io/docs/concepts/pipeline/metadata/"`
	Clone    *bool `yaml:"clone,omitempty" json:"clone,omitempty" jsonschema:"default=true,description=Enables injecting the default clone process.\nReference: https://go-vela.github.io/docs/concepts/pipeline/metadata/"`
}

// ToPipeline converts the Metadata type
// to a pipeline Metadata type.
func (m *Metadata) ToPipeline() *pipeline.Metadata {
	var clone bool
	if m.Clone == nil {
		clone = true
	} else {
		clone = *m.Clone
	}
	return &pipeline.Metadata{
		Template: m.Template,
		Clone:    clone,
	}
}
