// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"errors"
	"fmt"
	"strings"
)

// Build is the yaml representation of a build for a pipeline.
// nolint:lll // jsonschema will cause long lines
type Build struct {
	Version   string        `yaml:"version,omitempty"   json:"version,omitempty"  jsonschema:"required,minLength=1,description=Provide syntax version used to evaluate the pipeline.\nReference: https://go-vela.github.io/docs/concepts/pipeline/version/"`
	Metadata  Metadata      `yaml:"metadata,omitempty"  json:"metadata,omitempty" jsonschema:"description=Pass extra information.\nReference: https://go-vela.github.io/docs/concepts/pipeline/metadata/"`
	Worker    Worker        `yaml:"worker,omitempty"    json:"worker,omitempty" jsonschema:"description=Limit the pipeline to certain types of workers.\nReference: coming soon"`
	Secrets   SecretSlice   `yaml:"secrets,omitempty"   json:"secrets,omitempty" jsonschema:"description=Provide sensitive information.\nReference: https://go-vela.github.io/docs/concepts/pipeline/secrets/"`
	Services  ServiceSlice  `yaml:"services,omitempty"  json:"services,omitempty" jsonschema:"description=Provide detached (headless) execution instructions.\nReference: https://go-vela.github.io/docs/concepts/pipeline/services/"`
	Stages    StageSlice    `yaml:"stages,omitempty"    json:"stages,omitempty" jsonschema:"oneof_required=stages,description=Provide parallel execution instructions.\nReference: https://go-vela.github.io/docs/concepts/pipeline/stages/"`
	Steps     StepSlice     `yaml:"steps,omitempty"     json:"steps,omitempty" jsonschema:"oneof_required=steps,description=Provide sequential execution instructions.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/"`
	Templates TemplateSlice `yaml:"templates,omitempty" json:"templates,omitempty" jsonschema:"description=Provide the name of templates to expand.\nReference: https://go-vela.github.io/docs/concepts/pipeline/templates/"`
}

// Validate lints if the build configuration is valid.
func (b *Build) Validate(pipeline []byte) error {
	invalid := errors.New("invalid pipeline found")

	// check a version is provided
	if len(b.Version) == 0 {
		invalid = fmt.Errorf("%w: %s", invalid, "no version provided")
	}

	// check that stages or steps are provided
	if len(b.Stages) == 0 && len(b.Steps) == 0 {
		invalid = fmt.Errorf("%w: %s", invalid, "no stages or steps provided")
	}

	// check that stages and steps aren't provided
	if len(b.Stages) > 0 && len(b.Steps) > 0 {
		invalid = fmt.Errorf("%w: %s", invalid, "stages and steps provided")
	}

	// validate the services block provided
	err := b.Services.Validate(pipeline)
	if err != nil {
		invalid = fmt.Errorf("%v: %w", invalid, err)
	}

	// validate the stages block provided
	if len(b.Stages) > 0 {
		err = b.Stages.Validate(pipeline)
		if err != nil {
			invalid = fmt.Errorf("%v: %w", invalid, err)
		}
	}

	// validate the steps block provided
	err = b.Steps.Validate(pipeline)
	if err != nil {
		invalid = fmt.Errorf("%v: %w", invalid, err)
	}

	if !strings.EqualFold(invalid.Error(), "invalid pipeline found") {
		return invalid
	}

	return nil
}
