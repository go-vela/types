// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

// nolint:lll // jsonschema will cause long lines
type (
	// TemplateSlice is the yaml representation
	// of the templates block for a pipeline.
	TemplateSlice []*Template

	// Template is the yaml representation of a template
	// from the templates block for a pipeline.
	Template struct {
		Name   string `yaml:"name,omitempty"   json:"name,omitempty"  jsonschema:"required,minLength=1,description=Unique identifier for the template.\nReference: https://go-vela.github.io/docs/concepts/pipeline/templates/"`
		Source string `yaml:"source,omitempty" json:"source,omitempty" jsonschema:"required,minLength=1,description=Path to template in remote system.\nReference: https://go-vela.github.io/docs/concepts/pipeline/templates/source/"`
		Type   string `yaml:"type,omitempty"   json:"type,omitempty" jsonschema:"minLength=1,description=Type of template provided from the remote system.\nReference: https://go-vela.github.io/docs/concepts/pipeline/templates/type/,example=github"`
	}

	// StepTemplate is the yaml representation of the
	// template block for a step in a pipeline.
	StepTemplate struct {
		Name      string                 `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"required,minLength=1,description=Unique identifier for the template.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/template/#fields"`
		Variables map[string]interface{} `yaml:"vars,omitempty" json:"vars,omitempty" jsonschema:"description=Variables injected into the template.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/template/#fields"`
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

// Map helper function that creates a map of templates from a slice of templates.
func (t *TemplateSlice) Map(templates TemplateSlice) map[string]*Template {
	m := make(map[string]*Template)

	for _, tmpl := range templates {
		m[tmpl.Name] = tmpl
	}

	return m
}
