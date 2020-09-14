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

	yaml "gopkg.in/yaml.v2"
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
