// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"errors"
	"fmt"
	"strings"
)

type (
	// SecretSlice is the pipeline representation
	// of the secrets block for a pipeline.
	//
	// swagger:model PipelineSecretSlice
	SecretSlice []*Secret

	// Secret is the pipeline representation of a
	// secret from the secrets block for a pipeline.
	//
	// swagger:model PipelineSecret
	Secret struct {
		Name   string     `json:"name,omitempty"   yaml:"name,omitempty"`
		Value  string     `json:"value,omitempty"  yaml:"value,omitempty"`
		Key    string     `json:"key,omitempty"    yaml:"key,omitempty"`
		Engine string     `json:"engine,omitempty" yaml:"engine,omitempty"`
		Type   string     `json:"type,omitempty"   yaml:"type,omitempty"`
		Origin *Container `json:"origin,omitempty" yaml:"origin,omitempty"`
	}

	// StepSecretSlice is the pipeline representation
	// of the secrets block for a step in a pipeline.
	//
	// swagger:model PipelineStepSecretSlice
	StepSecretSlice []*StepSecret

	// StepSecret is the pipeline representation of a secret
	// from a secrets block for a step in a pipeline.
	//
	// swagger:model PipelineStepSecret
	StepSecret struct {
		Source string `json:"source,omitempty" yaml:"source,omitempty"`
		Target string `json:"target,omitempty" yaml:"target,omitempty"`
	}
)

var (
	// ErrInvalidOrg defines the error type when the
	// org in key does not equal the name of the organization.
	ErrInvalidOrg = errors.New("invalid organization in key")
	// ErrInvalidRepo defines the error type when the
	// repo in key does not equal the name of the repository.
	ErrInvalidRepo = errors.New("invalid repository in key")
	// ErrInvalidShared defines the error type when the
	// org in key does not equal the name of the team.
	ErrInvalidShared = errors.New("invalid team in key")
	// ErrInvalidPath defines the error type when the
	// path provided for a type (org, repo, shared) is invalid.
	ErrInvalidPath = errors.New("invalid secret path")
)

// Purge removes the secrets that have a ruleset
// that do not match the provided ruledata.
func (s *SecretSlice) Purge(r *RuleData) *SecretSlice {
	counter := 1
	secrets := new(SecretSlice)

	// iterate through each Secret in the pipeline
	for _, secret := range *s {
		if secret.Origin.Empty() {
			// append the secret to the new slice of secrets
			*secrets = append(*secrets, secret)

			continue
		}

		// verify ruleset matches
		if secret.Origin.Ruleset.Match(r) {
			// overwrite the Container number with the Container counter
			secret.Origin.Number = counter

			// increment Container counter
			counter = counter + 1

			// append the secret to the new slice of secrets
			*secrets = append(*secrets, secret)
		}
	}

	return secrets
}

// ValidOrg returns true when the secret is valid for a given
// organization.
func (s *Secret) ValidOrg(org string) error {
	path := s.Key

	// check if a path was provided
	if !strings.Contains(path, "/") {
		return fmt.Errorf("%s: %s ", ErrInvalidPath, path)
	}

	// split the full path into parts
	parts := strings.SplitN(path, "/", 2)

	// secret is invalid
	if len(parts) != 2 {
		return fmt.Errorf("%s: %s ", ErrInvalidPath, path)
	}

	// check if the org provided matches what we expect
	if !strings.EqualFold(parts[0], org) {
		return fmt.Errorf("%s: %s ", ErrInvalidOrg, org)
	}

	return nil
}

// ValidRepo returns an error when the secret is valid for a given
// organization and repository.
func (s *Secret) ValidRepo(org, repo string) error {
	path := s.Key

	// check if a path was provided
	if !strings.Contains(path, "/") {
		return fmt.Errorf("%s: %s ", ErrInvalidPath, path)
	}

	// split the full path into parts
	parts := strings.SplitN(path, "/", 3)

	// secret is invalid
	if len(parts) != 3 {
		return fmt.Errorf("%s: %s ", ErrInvalidPath, path)
	}

	// check if the org provided matches what we expect
	if !strings.EqualFold(parts[0], org) {
		return fmt.Errorf("%s: %s ", ErrInvalidOrg, org)
	}

	// check if the repo provided matches what we expect
	if !strings.EqualFold(parts[1], repo) {
		return fmt.Errorf("%s: %s ", ErrInvalidRepo, repo)
	}

	return nil
}

// ValidShared returns true when the secret is valid for a given
// organization and team.
func (s *Secret) ValidShared(org string) error {
	path := s.Key

	// check if a path was provided
	if !strings.Contains(path, "/") {
		return fmt.Errorf("%s: %s ", ErrInvalidPath, path)
	}

	// split the full path into parts
	parts := strings.SplitN(path, "/", 3)

	// secret is invalid
	if len(parts) != 3 {
		return fmt.Errorf("%s: %s ", ErrInvalidPath, path)
	}

	// check if the org provided is not empty
	if !strings.EqualFold(parts[0], org) {
		return fmt.Errorf("%s: %s ", ErrInvalidOrg, org)
	}

	return nil
}
