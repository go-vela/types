// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"errors"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"
)

type (
	// SecretSlice is the yaml representation
	// of the secrets block for a pipeline.
	SecretSlice []*Secret

	// Secret is the yaml representation of a secret
	// from the secrets block for a pipeline.
	// nolint:lll // jsonschema will cause long lines
	Secret struct {
		Name   string `yaml:"name,omitempty"   jsonschema:"required,minLength=1,description=Name of secret to reference in the pipeline.\nReference: https://go-vela.github.io/docs/concepts/pipeline/secrets/"`
		Key    string `yaml:"key,omitempty"    jsonschema:"minLength=1,description=Path to secret to fetch from storage backend.\nReference: https://go-vela.github.io/docs/concepts/pipeline/secrets/key/"`
		Engine string `yaml:"engine,omitempty" jsonschema:"enum=native,enum=vault,default=native,description=Name of storage backend to fetch secret from.\nReference: https://go-vela.github.io/docs/concepts/pipeline/secrets/engine/"`
		Type   string `yaml:"type,omitempty"   jsonschema:"enum=repo,enum=org,enum=shared,default=repo,description=Type of secret to fetch from storage backend.\nReference: https://go-vela.github.io/docs/concepts/pipeline/secrets/type/"`
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
		})
	}

	return secretSlice
}

// UnmarshalYAML implements the Unmarshaler interface for the SecretSlice type.
func (s *SecretSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// secret slice we try unmarshaling to
	secretSlice := new([]*Secret)

	// attempt to unmarshal as a secret slice type
	err := unmarshal(secretSlice)
	if err != nil {
		return err
	}

	// iterate through each element in the secret slice
	for _, secret := range *secretSlice {
		// implicitly set `key` field if empty
		if len(secret.Key) == 0 {
			secret.Key = secret.Name
		}

		// implicitly set `engine` field if empty
		if len(secret.Engine) == 0 {
			secret.Engine = constants.DriverNative
		}

		// implicitly set `type` field if empty
		if len(secret.Type) == 0 {
			secret.Type = constants.SecretRepo
		}
	}

	// overwrite existing SecretSlice
	*s = *secretSlice

	return nil
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
	// string slice we try unmarshaling to
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

	// step secret slice we try unmarshaling to
	secrets := new([]*StepSecret)

	// attempt to unmarshal as a step secret slice type
	err = unmarshal(secrets)
	if err == nil {

		// overwrite existing StepSecretSlice
		*s = StepSecretSlice(*secrets)
		return nil
	}

	return errors.New("Failed to unmarshal StepSecretSlice")
}
