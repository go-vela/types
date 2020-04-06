// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

type (
	// SecretSlice is the pipeline representation
	// of the secrets block for a pipeline.
	SecretSlice []*Secret

	// Secret is the pipeline representation of a
	// secret from the secrets block for a pipeline.
	Secret struct {
		Name   string `json:"name,omitempty"   yaml:"name,omitempty"`
		Value  string `json:"value,omitempty"  yaml:"value,omitempty"`
		Key    string `json:"key,omitempty"    yaml:"key,omitempty"`
		Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`
		Type   string `json:"type,omitempty"   yaml:"type,omitempty"`
	}

	// StepSecretSlice is the pipeline representation
	// of the secrets block for a step in a pipeline.
	StepSecretSlice []*StepSecret

	// StepSecret is the pipeline representation of a secret
	// from a secrets block for a step in a pipeline.
	StepSecret struct {
		Source string `json:"source,omitempty" yaml:"source,omitempty"`
		Target string `json:"target,omitempty" yaml:"target,omitempty"`
	}
)
