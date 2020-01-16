// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

type (
	// TemplateSlice is the yaml representation
	// of the templates block for a pipeline.
	TemplateSlice []*Template

	// Template is the yaml representation of a template
	// from the templates block for a pipeline.
	Template struct {
		Name   string `yaml:"name,omitempty"`
		Source string `yaml:"source,omitempty"`
		Type   string `yaml:"type,omitempty"`
	}

	// StepTemplate is the yaml representation of the
	// template block for a step in a pipeline.
	StepTemplate struct {
		Name      string                 `yaml:"name,omitempty"`
		Variables map[string]interface{} `yaml:"vars,omitempty"`
	}
)

// UnmarshalYAML implements the Unmarshaler interface for the TemplateSlice type.
func (t *TemplateSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// template slice we try unmarshaling to
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
