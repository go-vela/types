// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"github.com/go-vela/types/library"
	"github.com/go-vela/types/raw"
)

// Build is the yaml representation of a build for a pipeline.
// nolint:lll // jsonschema will cause long lines
type Build struct {
	Version     string             `yaml:"version,omitempty"   json:"version,omitempty"  jsonschema:"required,minLength=1,description=Provide syntax version used to evaluate the pipeline.\nReference: https://go-vela.github.io/docs/reference/yaml/version/"`
	Metadata    Metadata           `yaml:"metadata,omitempty"  json:"metadata,omitempty" jsonschema:"description=Pass extra information.\nReference: https://go-vela.github.io/docs/reference/yaml/metadata/"`
	Environment raw.StringSliceMap `yaml:"environment,omitempty" json:"environment,omitempty" jsonschema:"description=Provide global environment variables injected into the container environment.\nReference: https://go-vela.github.io/docs/reference/yaml/steps/#the-environment-tag"`
	Worker      Worker             `yaml:"worker,omitempty"    json:"worker,omitempty" jsonschema:"description=Limit the pipeline to certain types of workers.\nReference: https://go-vela.github.io/docs/reference/yaml/worker/"`
	Secrets     SecretSlice        `yaml:"secrets,omitempty"   json:"secrets,omitempty" jsonschema:"description=Provide sensitive information.\nReference: https://go-vela.github.io/docs/reference/yaml/secrets/"`
	Services    ServiceSlice       `yaml:"services,omitempty"  json:"services,omitempty" jsonschema:"description=Provide detached (headless) execution instructions.\nReference: https://go-vela.github.io/docs/reference/yaml/services/"`
	Stages      StageSlice         `yaml:"stages,omitempty"    json:"stages,omitempty" jsonschema:"oneof_required=stages,description=Provide parallel execution instructions.\nReference: https://go-vela.github.io/docs/reference/yaml/stages/"`
	Steps       StepSlice          `yaml:"steps,omitempty"     json:"steps,omitempty" jsonschema:"oneof_required=steps,description=Provide sequential execution instructions.\nReference: https://go-vela.github.io/docs/reference/yaml/steps/"`
	Templates   TemplateSlice      `yaml:"templates,omitempty" json:"templates,omitempty" jsonschema:"description=Provide the name of templates to expand.\nReference: https://go-vela.github.io/docs/reference/yaml/templates/"`
}

// ToLibrary converts the Build type to a library Pipeline type.
func (b *Build) ToLibrary() *library.Pipeline {
	pipeline := new(library.Pipeline)

	pipeline.SetFlavor(b.Worker.Flavor)
	pipeline.SetPlatform(b.Worker.Platform)
	pipeline.SetVersion(b.Version)
	pipeline.SetServices(len(b.Services) > 0)
	pipeline.SetStages(len(b.Stages) > 0)
	pipeline.SetSteps(len(b.Steps) > 0)
	pipeline.SetTemplates(len(b.Templates) > 0)

	return pipeline
}
