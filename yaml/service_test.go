// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"
	"github.com/goccy/go-yaml"
)

func TestYaml_ServiceSlice_ToPipeline(t *testing.T) {
	// setup tests
	tests := []struct {
		services *ServiceSlice
		want     *pipeline.ContainerSlice
	}{
		{
			services: &ServiceSlice{
				{
					Entrypoint:  []string{"/usr/local/bin/docker-entrypoint.sh"},
					Environment: map[string]string{"FOO": "bar"},
					Image:       "postgres:12-alpine",
					Name:        "postgres",
					Ports:       []string{"5432:5432"},
					Pull:        "not_present",
				},
			},
			want: &pipeline.ContainerSlice{
				{
					Detach:      true,
					Entrypoint:  []string{"/usr/local/bin/docker-entrypoint.sh"},
					Environment: map[string]string{"FOO": "bar"},
					Image:       "postgres:12-alpine",
					Name:        "postgres",
					Ports:       []string{"5432:5432"},
					Pull:        "not_present",
				},
			},
		},
	}

	// run tests
	for _, test := range tests {
		got := test.services.ToPipeline()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ToPipeline is %v, want %v", got, test.want)
		}
	}
}

func TestYaml_ServiceSlice_UnmarshalYAML(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		file    string
		want    *ServiceSlice
	}{
		{
			failure: false,
			file:    "testdata/service.yml",
			want: &ServiceSlice{
				{
					Environment: raw.StringSliceMap{
						"POSTGRES_DB": "foo",
					},
					Image: "postgres:latest",
					Name:  "postgres",
					Ports: []string{"5432:5432"},
					Pull:  "not_present",
				},
				{
					Environment: raw.StringSliceMap{
						"MYSQL_DATABASE": "foo",
					},
					Image: "mysql:latest",
					Name:  "mysql",
					Ports: []string{"3061:3061"},
					Pull:  "not_present",
				},
			},
		},
		{
			failure: true,
			file:    "testdata/invalid.yml",
			want:    nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := new(ServiceSlice)

		b, err := ioutil.ReadFile(test.file)
		if err != nil {
			t.Errorf("unable to read file: %v", err)
		}

		err = yaml.Unmarshal(b, got)

		if test.failure {
			if err == nil {
				t.Errorf("UnmarshalYAML should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("UnmarshalYAML returned err: %v", err)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("UnmarshalYAML is %v, want %v", got, test.want)
		}
	}
}

func TestServiceSlice_Validate(t *testing.T) {
	//setup types
	tests := []struct {
		name    string
		file    string
		wantErr bool
	}{
		{
			name:    "success: minimal service block",
			file:    "testdata/service/validate/minimal.yml",
			wantErr: false,
		},
		{
			name:    "failure: missing name yaml tag",
			file:    "testdata/service/validate/missing_name.yml",
			wantErr: true,
		},
		{
			name:    "failure: missing image yaml tag",
			file:    "testdata/service/validate/missing_image.yml",
			wantErr: true,
		},
		{
			name:    "failure: bad image tag data",
			file:    "testdata/service/validate/bad_image.yml",
			wantErr: true,
		},
	}

	// run tests
	for _, test := range tests {
		b := new(Build)

		pipeline, err := ioutil.ReadFile(test.file)
		if err != nil {
			t.Errorf("Reading file for Validate returned err: %v", err)
		}

		err = yaml.Unmarshal(pipeline, b)

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}

		t.Run(test.name, func(t *testing.T) {
			if err := b.Services.Validate(pipeline); (err != nil) != test.wantErr {
				t.Errorf("Validate is %v, want %v", err, test.wantErr)
			}
		})
	}
}
