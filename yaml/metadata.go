// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"github.com/go-vela/types/pipeline"
)

//version: "1"
//
//metadata:
//render_inline: true # defaults to false
//
//templates:
//- name: go
//	source: github.com/octocat/hello-world/.vela/build.yml
//	format: go
//	type: github
//	vars:
//	image: golang:latest

type (
	// Metadata is the yaml representation of
	// the metadata block for a pipeline.
	Metadata struct {
		Template     bool     `yaml:"template,omitempty" json:"template,omitempty" jsonschema:"description=Enables compiling the pipeline as a template.\nReference: https://go-vela.github.io/docs/reference/yaml/metadata/#the-template-tag"`
		RenderInline bool     `yaml:"render_inline,omitempty" json:"render_inline,omitempty" jsonschema:"description=Enables inline compiling for the pipeline templates.\nReference: TODO"`
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

// UnmarshalYAML implements the Unmarshaler interface for the Metadata type.
func (m *Metadata) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// metadata we try unmarshalling to
	metadata := new(struct {
		Template     bool
		RenderInline bool `yaml:"render_inline,omitempty" json:"render_inline,omitempty"`
		Clone        *bool
		Environment  []string
	})

	// attempt to unmarshal as a metadata type
	err := unmarshal(metadata)
	if err != nil {
		return err
	}

	if len(metadata.Environment) == 0 {
		metadata.Environment = []string{"steps", "services", "secrets"}
	}

	// overwrite existing metadata environment details
	m.Template = metadata.Template
	m.RenderInline = metadata.RenderInline
	m.Clone = metadata.Clone
	m.Environment = metadata.Environment

	return nil
}
