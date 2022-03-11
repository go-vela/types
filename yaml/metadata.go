// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"github.com/go-vela/types/pipeline"
)

type (
	// Metadata is the yaml representation of
	// the metadata block for a pipeline.
	Metadata struct {
		Template     bool     `yaml:"template,omitempty" json:"template,omitempty" jsonschema:"description=Enables compiling the pipeline as a template.\nReference: https://go-vela.github.io/docs/reference/yaml/metadata/#the-template-tag"`
		RenderInline bool     `yaml:"render_inline,omitempty" json:"render_inline,omitempty" jsonschema:"description=Enables inline compiling for the pipeline templates.\nReference: https://go-vela.github.io/docs/reference/yaml/metadata/#the-render-inline-tag"`
		Clone        *bool    `yaml:"clone,omitempty" json:"clone,omitempty" jsonschema:"default=true,description=Enables injecting the default clone process.\nReference: https://go-vela.github.io/docs/reference/yaml/metadata/#the-clone-tag"`
		Environment  []string `yaml:"environment,omitempty" json:"environment,omitempty" jsonschema:"description=Controls which containers processes can have global env injected.\nReference: https://go-vela.github.io/docs/reference/yaml/metadata/#the-environment-tag"`
	}
)

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
		Template:    m.Template,
		Clone:       clone,
		Environment: m.Environment,
	}
}

// HasEnvironment checks if the container type
// is contained within the environment list.
func (m *Metadata) HasEnvironment(container string) bool {
	for _, e := range m.Environment {
		if e == container {
			return true
		}
	}

	return false
}
