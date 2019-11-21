// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

// Build is the yaml representation of a build for a pipeline.
type Build struct {
	Version   string        `yaml:"version,omitempty"`
	Metadata  Metadata      `yaml:"metadata,omitempty"`
	Worker    Worker        `yaml:"worker,omitempty"`
	Secrets   SecretSlice   `yaml:"secrets,omitempty"`
	Services  ServiceSlice  `yaml:"services,omitempty"`
	Stages    StageSlice    `yaml:"stages,omitempty"`
	Steps     StepSlice     `yaml:"steps,omitempty"`
	Templates TemplateSlice `yaml:"templates,omitempty"`
}
