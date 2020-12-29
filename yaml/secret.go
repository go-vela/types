// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"
	"github.com/goccy/go-yaml"
)

// nolint:lll // jsonschema will cause long lines
type (
	// SecretSlice is the yaml representation
	// of the secrets block for a pipeline.
	SecretSlice []*Secret

	// Secret is the yaml representation of a secret
	// from the secrets block for a pipeline.
	Secret struct {
		Name   string `yaml:"name,omitempty"   json:"name,omitempty" jsonschema:"required,minLength=1,description=Name of secret to reference in the pipeline.\nReference: https://go-vela.github.io/docs/concepts/pipeline/secrets/"`
		Key    string `yaml:"key,omitempty"    json:"key,omitempty" jsonschema:"minLength=1,description=Path to secret to fetch from storage backend.\nReference: https://go-vela.github.io/docs/concepts/pipeline/secrets/key/"`
		Engine string `yaml:"engine,omitempty" json:"engine,omitempty" jsonschema:"enum=native,enum=vault,default=native,description=Name of storage backend to fetch secret from.\nReference: https://go-vela.github.io/docs/concepts/pipeline/secrets/engine/"`
		Type   string `yaml:"type,omitempty"   json:"type,omitempty" jsonschema:"enum=repo,enum=org,enum=shared,default=repo,description=Type of secret to fetch from storage backend.\nReference: https://go-vela.github.io/docs/concepts/pipeline/secrets/type/"`
		Origin Origin `yaml:"origin,omitempty" json:"origin,omitempty" jsonschema:"description=Declaration to pull secrets from non-internal secret providers.\nReference: https://go-vela.github.io/docs/concepts/pipeline/secrets/origin/"`
	}

	// Origin is the yaml representation of a method
	// for looking up secrets with a secret plugin.
	Origin struct {
		Environment raw.StringSliceMap     `yaml:"environment,omitempty" json:"environment,omitempty" jsonschema:"description=Variables to inject into the container environment.\nReference: coming soon"`
		Image       string                 `yaml:"image,omitempty"       json:"image,omitempty" jsonschema:"required,minLength=1,description=Docker image to use to create the ephemeral container.\nReference: "`
		Name        string                 `yaml:"name,omitempty"        json:"name,omitempty" jsonschema:"required,minLength=1,description=Unique name for the secret origin."`
		Parameters  map[string]interface{} `yaml:"parameters,omitempty"  json:"parameters,omitempty" jsonschema:"description=Extra configuration variables for the secret plugin.\nReference: coming soon"`
		Secrets     StepSecretSlice        `yaml:"secrets,omitempty"     json:"secrets,omitempty" jsonschema:"description=Secrets to inject that are necessary to retrieve the secrets.\nReference: coming soon"`
		Pull        string                 `yaml:"pull,omitempty"        json:"pull,omitempty" jsonschema:"enum=always,enum=not_present,enum=on_start,enum=never,default=not_present,description=Declaration to configure if and when the Docker image is pulled.\nReference: coming soon"`
		Ruleset     Ruleset                `yaml:"ruleset,omitempty"     json:"ruleset,omitempty" jsonschema:"description=Conditions to limit the execution of the container.\nReference: coming soon"`
	}
)

// ToPipeline converts the SecretSlice type
// to a pipeline SecretSlice type.
func (s *SecretSlice) ToPipeline() *pipeline.SecretSlice {
	// secret slice we want to return
	secretSlice := new(pipeline.SecretSlice)

	// iterate through each element in the secret slice
	for _, secret := range *s {
		// append the element to the pipeline secret slice
		*secretSlice = append(*secretSlice, &pipeline.Secret{
			Name:   secret.Name,
			Key:    secret.Key,
			Engine: secret.Engine,
			Type:   secret.Type,
			Origin: secret.Origin.ToPipeline(),
		})
	}

	return secretSlice
}

// UnmarshalYAML implements the Unmarshaler interface for the SecretSlice type.
func (s *SecretSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// secret slice we try unmarshalling to
	secretSlice := new([]*Secret)

	// attempt to unmarshal as a secret slice type
	err := unmarshal(secretSlice)
	if err != nil {
		return err
	}

	// iterate through each element in the secret slice
	for _, secret := range *secretSlice {
		// implicitly set `key` field if empty
		if secret.Origin.Empty() && len(secret.Key) == 0 {
			secret.Key = secret.Name
		}

		// implicitly set `engine` field if empty
		if secret.Origin.Empty() && len(secret.Engine) == 0 {
			secret.Engine = constants.DriverNative
		}

		// implicitly set `type` field if empty
		if secret.Origin.Empty() && len(secret.Type) == 0 {
			secret.Type = constants.SecretRepo
		}

		// implicitly set `pull` field if empty
		if !secret.Origin.Empty() && len(secret.Origin.Pull) == 0 {
			secret.Origin.Pull = constants.PullNotPresent
		}

		// TODO: remove this in a future release
		//
		// handle true deprecated pull policy
		//
		// a `true` pull policy equates to `always`
		if !secret.Origin.Empty() && strings.EqualFold(secret.Origin.Pull, "true") {
			secret.Origin.Pull = constants.PullAlways
		}

		// TODO: remove this in a future release
		//
		// handle false deprecated pull policy
		//
		// a `false` pull policy equates to `not_present`
		if !secret.Origin.Empty() && strings.EqualFold(secret.Origin.Pull, "false") {
			secret.Origin.Pull = constants.PullNotPresent
		}
	}

	// overwrite existing SecretSlice
	*s = *secretSlice

	return nil
}

// Empty returns true if the provided origin is empty.
func (o *Origin) Empty() bool {
	// return true if every origin field is empty
	if o.Environment == nil &&
		len(o.Image) == 0 &&
		len(o.Name) == 0 &&
		o.Parameters == nil &&
		len(o.Secrets) == 0 &&
		len(o.Pull) == 0 {
		return true
	}

	return false
}

// ToPipeline converts the Origin type
// to a pipeline Container type.
func (o *Origin) ToPipeline() *pipeline.Container {
	return &pipeline.Container{
		Environment: o.Environment,
		Image:       o.Image,
		Name:        o.Name,
		Pull:        o.Pull,
		Ruleset:     *o.Ruleset.ToPipeline(),
		Secrets:     *o.Secrets.ToPipeline(),
	}
}

type (
	// StepSecretSlice is the yaml representation of
	// the secrets block for a step in a pipeline.
	StepSecretSlice []*StepSecret

	// StepSecret is the yaml representation of a secret
	// from a secrets block for a step in a pipeline.
	StepSecret struct {
		Source string `yaml:"source,omitempty"`
		Target string `yaml:"target,omitempty"`
	}
)

// ToPipeline converts the StepSecretSlice type
// to a pipeline StepSecretSlice type.
func (s *StepSecretSlice) ToPipeline() *pipeline.StepSecretSlice {
	// step secret slice we want to return
	secretSlice := new(pipeline.StepSecretSlice)

	// iterate through each element in the step secret slice
	for _, secret := range *s {
		// append the element to the pipeline step secret slice
		*secretSlice = append(*secretSlice, &pipeline.StepSecret{
			Source: secret.Source,
			Target: secret.Target,
		})
	}

	return secretSlice
}

// UnmarshalYAML implements the Unmarshaler interface for the StepSecretSlice type.
func (s *StepSecretSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// string slice we try unmarshalling to
	stringSlice := new(raw.StringSlice)

	// attempt to unmarshal as a string slice type
	err := unmarshal(stringSlice)
	if err == nil {
		// iterate through each element in the string slice
		for _, secret := range *stringSlice {
			// append the element to the step secret slice
			*s = append(*s, &StepSecret{
				Source: secret,
				Target: secret,
			})
		}

		return nil
	}

	// step secret slice we try unmarshalling to
	secrets := new([]*StepSecret)

	// attempt to unmarshal as a step secret slice type
	err = unmarshal(secrets)
	if err == nil {
		// check for secret source and target
		for _, secret := range *secrets {
			if len(secret.Source) == 0 || len(secret.Target) == 0 {
				return fmt.Errorf("no secret source or target found")
			}
		}

		// overwrite existing StepSecretSlice
		*s = StepSecretSlice(*secrets)
		return nil
	}

	return errors.New("failed to unmarshal StepSecretSlice")
}

// Validate lints if the steps configuration is valid.
func (s *SecretSlice) Validate(pipeline []byte) error {
	invalid := errors.New("invalid step block found")

	// iterate through each step and linting yaml tags
	for i, secret := range *s {
		// check required name field
		if len(secret.Name) == 0 {
			path, err := yaml.PathString(fmt.Sprintf("$.secrets[%d]", i))
			if err != nil {
				return err
			}
			source, err := path.AnnotateSource(pipeline, true)
			if err != nil {
				return err
			}

			// nolint:cSpell // ignore line length
			invalid = fmt.Errorf("%w: %s", invalid, fmt.Sprintf("no name provided for secret:\n%s\n ", string(source)))
		}

		// validate secret by type
		switch secret.Type {
		case constants.SecretRepo:
			err := secret.validateRepo(pipeline, i)
			if err != nil {
				invalid = fmt.Errorf("%w: %v", invalid, err)
			}
		case constants.SecretOrg:
			err := secret.validateOrg(pipeline, i)
			if err != nil {
				invalid = fmt.Errorf("%w: %v", invalid, err)
			}
		}
	}

	// check if only default error exists
	if !strings.EqualFold(invalid.Error(), "invalid step block found") {
		return invalid
	}

	return nil
}

// validateRepo is a helper function to lint secrets of type "repo".
//
// this function is used to check the fields of secret with the explicit
// definition yaml style.
func (s *Secret) validateRepo(pipeline []byte, i int) error {
	var invalid error

	// check if the engine is not a "native" or "vault"
	if len(s.Engine) != 0 {
		if !strings.EqualFold(s.Engine, constants.DriverNative) &&
			!strings.EqualFold(s.Engine, constants.DriverVault) {
			path, err := yaml.PathString(fmt.Sprintf("$.secrets[%d].engine", i))
			if err != nil {
				return err
			}
			source, err := path.AnnotateSource(pipeline, true)
			if err != nil {
				return err
			}

			// nolint:cSpell // ignore line length
			invalid = fmt.Errorf("%w: %s", invalid, fmt.Sprintf("invalid engine provided for secret:\n%s\n ", string(source)))
		}
	}

	// check if a key was provided for explicit definition
	// when the key == name than we have an implicit definition
	if len(s.Key) != 0 && s.Key != s.Name {
		match, err := regexp.MatchString(`.+\/.+\/.+`, s.Key)
		if err != nil {
			return err
		}

		if !match {
			path, err := yaml.PathString(fmt.Sprintf("$.secrets[%d].key", i))
			if err != nil {
				return err
			}
			source, err := path.AnnotateSource(pipeline, true)
			if err != nil {
				return err
			}

			// nolint:cSpell // ignore line length
			invalid = fmt.Errorf("%w: %s", invalid, fmt.Sprintf("invalid key provided for secret:\n%s\n ", string(source)))
		}
	}

	return invalid
}

// validateOrg is a helper function to lint secrets of type "org".
func (s *Secret) validateOrg(pipeline []byte, i int) error {
	var invalid error

	// check if the engine is not a "native" or "vault"
	if len(s.Engine) == 0 {
		path, err := yaml.PathString(fmt.Sprintf("$.secrets[%d]", i))
		if err != nil {
			return err
		}
		source, err := path.AnnotateSource(pipeline, true)
		if err != nil {
			return err
		}

		// nolint:cSpell // ignore line length
		invalid = fmt.Errorf("%w: %s", invalid, fmt.Sprintf("no engine provided for secret:\n%s\n ", string(source)))
	} else {
		if !strings.EqualFold(s.Engine, constants.DriverNative) &&
			!strings.EqualFold(s.Engine, constants.DriverVault) {
			path, err := yaml.PathString(fmt.Sprintf("$.secrets[%d].engine", i))
			if err != nil {
				return err
			}
			source, err := path.AnnotateSource(pipeline, true)
			if err != nil {
				return err
			}

			// nolint:cSpell // ignore line length
			invalid = fmt.Errorf("%w: %s", invalid, fmt.Sprintf("invalid engine provided for secret:\n%s\n ", string(source)))
		}
	}

	// check if a key was provided
	if len(s.Key) == 0 {
		path, err := yaml.PathString(fmt.Sprintf("$.secrets[%d]", i))
		if err != nil {
			return err
		}
		source, err := path.AnnotateSource(pipeline, true)
		if err != nil {
			return err
		}

		// nolint:cSpell // ignore line length
		invalid = fmt.Errorf("%w: %s", invalid, fmt.Sprintf("no key provided for secret:\n%s\n ", string(source)))
	} else {
		match, err := regexp.MatchString(`.+\/.+`, s.Key)
		if err != nil {
			return err
		}

		if !match {
			// nolint:gomnd // accepting magic number
			if len(strings.SplitN(s.Key, "/", 2)) != 2 {
				path, err := yaml.PathString(fmt.Sprintf("$.secrets[%d].key", i))
				if err != nil {
					return err
				}
				source, err := path.AnnotateSource(pipeline, true)
				if err != nil {
					return err
				}

				// nolint:cSpell // ignore line length
				invalid = fmt.Errorf("%w: %s", invalid, fmt.Sprintf("invalid key provided for secret:\n%s\n ", string(source)))
			}
		}
	}

	return invalid
}
