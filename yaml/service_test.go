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
	// setup types
	str := "foo"
	mapp := map[string]string{"foo": "bar"}
	slice := []string{"8000:8000"}
	want := &pipeline.ContainerSlice{
		&pipeline.Container{
			Ports:       slice,
			Entrypoint:  slice,
			Environment: mapp,
			Image:       str,
			Name:        str,
		},
	}

	s := &ServiceSlice{
		&Service{
			Ports:       slice,
			Entrypoint:  slice,
			Environment: mapp,
			Image:       str,
			Name:        str,
		},
	}

	// run test
	got := s.ToPipeline()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToPipeline is %v, want %v", got, want)
	}
}

func TestYaml_ServiceSlice_UnmarshalYAML(t *testing.T) {
	// setup types
	want := &ServiceSlice{
		&Service{
			Ports: []string{"5432:5432"},
			Environment: raw.StringSliceMap{
				"POSTGRES_DB": "foo",
			},
			Name:  "postgres",
			Image: "postgres:latest",
		},
		&Service{
			Ports: []string{"3061:3061"},
			Environment: raw.StringSliceMap{
				"MYSQL_DATABASE": "foo",
			},
			Name:  "mysql",
			Image: "mysql:latest",
		},
	}
	got := new(ServiceSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/service.yml")
	if err != nil {
		t.Errorf("Reading file for Ruleset UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err != nil {
		t.Errorf("YamlSlice UnmarshalYAML returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("YamlSlice UnmarshalYAML is %v, want %v", got, want)
	}
}

func TestYaml_ServiceSlice_UnmarshalYAML_Invalid(t *testing.T) {
	// run test
	b, err := ioutil.ReadFile("testdata/invalid.yml")
	if err != nil {
		t.Errorf("Reading file for YamlSlice UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, new(ServiceSlice))

	if err == nil {
		t.Errorf("YamlSlice UnmarshalYAML should have returned err")
	}
}
