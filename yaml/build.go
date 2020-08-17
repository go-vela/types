// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

// Build is the yaml representation of a build for a pipeline.
// nolint:lll // jsonschema will cause long lines
type Build struct {
	Version   string        `yaml:"version,omitempty" json:"version,omitempty" jsonschema:"minLength=1,description=Provide syntax version used to evaluate the pipeline.\nReference: https://go-vela.github.io/docs/concepts/pipeline/version/"`
	Metadata  Metadata      `yaml:"metadata,omitempty" json:"metadata,omitempty" jsonschema:"description=Pass extra information.\nReference: https://go-vela.github.io/docs/concepts/pipeline/metadata/"`
	Worker    Worker        `yaml:"worker,omitempty" json:"worker,omitempty" jsonschema:"description=Limit the pipeline to certain types of workers.\nReference: coming soon"`
	Secrets   SecretSlice   `yaml:"secrets,omitempty" json:"secrets,omitempty" jsonschema:"description=Provide sensitive information.\nReference: https://go-vela.github.io/docs/concepts/pipeline/secrets/"`
	Services  ServiceSlice  `yaml:"services,omitempty" json:"services,omitempty" jsonschema:"description=Provide detached (headless) execution instructions.\nReference: https://go-vela.github.io/docs/concepts/pipeline/services/"`
	Stages    StageSlice    `yaml:"stages,omitempty" json:"stages,omitempty" jsonschema:"description=Provide parallel execution instructions.\nReference: https://go-vela.github.io/docs/concepts/pipeline/stages/"`
	Steps     StepSlice     `yaml:"steps,omitempty" json:"steps,omitempty" jsonschema:"description=Provide sequential execution instructions.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/"`
	Templates TemplateSlice `yaml:"templates,omitempty" json:"templates,omitempty" jsonschema:"description=Provide the name of templates to expand.\nReference: https://go-vela.github.io/docs/concepts/pipeline/templates/"`
}
