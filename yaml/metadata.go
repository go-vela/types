// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import "github.com/go-vela/types/pipeline"

// Metadata is the yaml representation of
// the metadata block for a pipeline.
type Metadata struct {
	Template bool `yaml:"template,omitempty"`
}

// ToPipeline converts the Metadata type
// to a pipeline Metadata type.
func (m *Metadata) ToPipeline() *pipeline.Metadata {
	return &pipeline.Metadata{
		Template: m.Template,
	}
}
