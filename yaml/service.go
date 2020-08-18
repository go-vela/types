// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"
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
		Pull        bool               `yaml:"pull,omitempty"        json:"pull,omitempty" jsonschema:"default=false,description=Enable pulling the latest version of the image.\nReference: https://go-vela.github.io/docs/concepts/pipeline/services/pull/"`
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
func (s *ServiceSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// service slice we try unmarshalling to
	serviceSlice := new([]*Service)

	// attempt to unmarshal as a service slice type
	err := unmarshal(serviceSlice)
	if err != nil {
		return err
	}

	// overwrite existing ServiceSlice
	*s = ServiceSlice(*serviceSlice)

	return nil
}
