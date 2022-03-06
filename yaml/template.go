// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"github.com/go-vela/types/library"
)

// nolint:lll // jsonschema will cause long lines
type (
	// TemplateSlice is the yaml representation
	// of the templates block for a pipeline.
	TemplateSlice []*Template

	// Template is the yaml representation of a template
	// from the templates block for a pipeline.
	Template struct {
		Name      string                 `yaml:"name,omitempty"   json:"name,omitempty"  jsonschema:"required,minLength=1,description=Unique identifier for the template.\nReference: https://go-vela.github.io/docs/reference/yaml/templates/#the-name-tag"`
		Source    string                 `yaml:"source,omitempty" json:"source,omitempty" jsonschema:"required,minLength=1,description=Path to template in remote system.\nReference: https://go-vela.github.io/docs/reference/yaml/templates/#the-source-tag"`
		Format    string                 `yaml:"format,omitempty" json:"format,omitempty" jsonschema:"enum=starlark,enum=golang,enum=go,default=go,minLength=1,description=language used within the template file \nReference: https://go-vela.github.io/docs/reference/yaml/templates/#the-format-tag"`
		Type      string                 `yaml:"type,omitempty"   json:"type,omitempty" jsonschema:"minLength=1,example=github,description=Type of template provided from the remote system.\nReference: https://go-vela.github.io/docs/reference/yaml/templates/#the-type-tag"`
		Variables map[string]interface{} `yaml:"vars,omitempty"   json:"vars,omitempty" jsonschema:"description=Variables injected into the template.\nReference: TODO"`
	}

	// StepTemplate is the yaml representation of the
	// template block for a step in a pipeline.
	StepTemplate struct {
		Name      string                 `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"required,minLength=1,description=Unique identifier for the template.\nReference: https://go-vela.github.io/docs/reference/yaml/steps/#the-template-tag"`
		Variables map[string]interface{} `yaml:"vars,omitempty" json:"vars,omitempty" jsonschema:"description=Variables injected into the template.\nReference: https://go-vela.github.io/docs/reference/yaml/steps/#the-template-tag"`
	}
)

// UnmarshalYAML implements the Unmarshaler interface for the TemplateSlice type.
func (t *TemplateSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// template slice we try unmarshalling to
	templateSlice := new([]*Template)

	// attempt to unmarshal as a template slice type
	err := unmarshal(templateSlice)
	if err != nil {
		return err
	}

	// overwrite existing TemplateSlice
	*t = TemplateSlice(*templateSlice)

	return nil
}

// ToLibrary converts the Template type
// to a library Template type.
func (t *Template) ToLibrary() *library.Template {
	template := new(library.Template)

	template.SetName(t.Name)
	template.SetSource(t.Source)
	template.SetType(t.Type)

	return template
}

// TemplateFromLibrary converts the library Template type
// to a yaml Template type.
func TemplateFromLibrary(t *library.Template) *Template {
	template := &Template{
		Name:   t.GetName(),
		Source: t.GetSource(),
		Type:   t.GetType(),
	}

	return template
}

// Map helper function that creates a map of templates from a slice of templates.
func (t *TemplateSlice) Map() map[string]*Template {
	m := make(map[string]*Template)

	if t == nil {
		return m
	}

	for _, tmpl := range *t {
		m[tmpl.Name] = tmpl
	}

	return m
}
