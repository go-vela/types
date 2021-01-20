// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"errors"
	"fmt"
	"strings"

	"github.com/docker/distribution/reference"
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"
	"github.com/goccy/go-yaml"
)

type (
	// ServiceSlice is the yaml representation
	// of the Services block for a pipeline.
	ServiceSlice []*Service

	// Service is the yaml representation
	// of a Service in a pipeline.
	// nolint:lll // jsonschema will cause long lines
	Service struct {
		Image       string             `yaml:"image,omitempty"       json:"image,omitempty" jsonschema:"required,minLength=1,description=Docker image used to create ephemeral container.\nReference: https://go-vela.github.io/docs/concepts/pipeline/services/image/"`
		Name        string             `yaml:"name,omitempty"        json:"name,omitempty" jsonschema:"required,minLength=1,description=Unique identifier for the container in the pipeline.\nReference: https://go-vela.github.io/docs/concepts/pipeline/services/"`
		Entrypoint  raw.StringSlice    `yaml:"entrypoint,omitempty"  json:"entrypoint,omitempty" jsonschema:"description=Commands to execute inside the container.\nReference: https://go-vela.github.io/docs/concepts/pipeline/services/entrypoint/"`
		Environment raw.StringSliceMap `yaml:"environment,omitempty" json:"environment,omitempty" jsonschema:"description=Variables to inject into the container environment.\nReference: https://go-vela.github.io/docs/concepts/pipeline/services/environment/"`
		Ports       raw.StringSlice    `yaml:"ports,omitempty"       json:"ports,omitempty" jsonschema:"description=List of ports to map for the container in the pipeline.\nReference: https://go-vela.github.io/docs/concepts/pipeline/services/ports/"`
		Pull        string             `yaml:"pull,omitempty"        json:"pull,omitempty" jsonschema:"enum=always,enum=not_present,enum=on_start,enum=never,default=not_present,description=Declaration to configure if and when the Docker image is pulled.\nReference: https://go-vela.github.io/docs/concepts/pipeline/services/pull/"`
	}
)

// ToPipeline converts the ServiceSlice type
// to a pipeline ContainerSlice type.
func (s *ServiceSlice) ToPipeline() *pipeline.ContainerSlice {
	// service slice we want to return
	serviceSlice := new(pipeline.ContainerSlice)

	// iterate through each element in the service slice
	for _, service := range *s {
		// append the element to the pipeline container slice
		*serviceSlice = append(*serviceSlice, &pipeline.Container{
			Detach:      true,
			Image:       service.Image,
			Name:        service.Name,
			Entrypoint:  service.Entrypoint,
			Environment: service.Environment,
			Ports:       service.Ports,
			Pull:        service.Pull,
		})
	}

	return serviceSlice
}

// UnmarshalYAML implements the Unmarshaler interface for the ServiceSlice type.
// nolint:dupl // accepting duplicative code that exists in step.go as well
func (s *ServiceSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// service slice we try unmarshalling to
	serviceSlice := new([]*Service)

	// attempt to unmarshal as a service slice type
	err := unmarshal(serviceSlice)
	if err != nil {
		return err
	}

	// iterate through each element in the service slice
	for _, service := range *serviceSlice {
		// implicitly set `pull` field if empty
		if len(service.Pull) == 0 {
			service.Pull = constants.PullNotPresent
		}

		// TODO: remove this in a future release
		//
		// handle true deprecated pull policy
		//
		// a `true` pull policy equates to `always`
		if strings.EqualFold(service.Pull, "true") {
			service.Pull = constants.PullAlways
		}

		// TODO: remove this in a future release
		//
		// handle false deprecated pull policy
		//
		// a `false` pull policy equates to `not_present`
		if strings.EqualFold(service.Pull, "false") {
			service.Pull = constants.PullNotPresent
		}
	}

	// overwrite existing ServiceSlice
	*s = ServiceSlice(*serviceSlice)

	return nil
}

// Validate lints if the services configuration is valid.
func (s *ServiceSlice) Validate(pipeline []byte) error {
	invalid := errors.New("invalid service block found")

	// iterate through each service and linting yaml tags
	for i, service := range *s {
		// check required fields
		if len(service.Name) == 0 {
			path, err := yaml.PathString(fmt.Sprintf("$.services[%d]", i))
			if err != nil {
				return err
			}
			source, err := path.AnnotateSource(pipeline, true)
			if err != nil {
				return err
			}

			invalid = fmt.Errorf("%w: %s", invalid,
				fmt.Sprintf("no name provided for service:\n%s\n ", string(source)))
		}

		if len(service.Image) == 0 {
			path, err := yaml.PathString(fmt.Sprintf("$.services[%d]", i))
			if err != nil {
				return err
			}
			source, err := path.AnnotateSource(pipeline, true)
			if err != nil {
				return err
			}

			invalid = fmt.Errorf("%w: %s", invalid,
				fmt.Errorf("no image provided for service:\n%s\n ", string(source)))
		} else {
			// parse the image provided into a
			// named, fully qualified reference
			//
			// https://pkg.go.dev/github.com/docker/distribution/reference?tab=doc#ParseAnyReference
			_, err := reference.ParseAnyReference(service.Image)
			if err != nil {
				// output error with YAML source
				path, err := yaml.PathString(fmt.Sprintf("$.services[%d].image", i))
				if err != nil {
					return err
				}
				source, err := path.AnnotateSource(pipeline, true)
				if err != nil {
					return err
				}

				invalid = fmt.Errorf("%w: %s", invalid,
					fmt.Errorf("invalid image value %s:\n%s\n ", service.Image, string(source)))
			}
		}
	}

	// check if only default error exists
	if !strings.EqualFold(invalid.Error(), "invalid service block found") {
		return invalid
	}

	return nil
}
