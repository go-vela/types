// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"io/ioutil"
	"reflect"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestBuild_TemplateSlice_UnmarshalYAML(t *testing.T) {
	// setup types
	want := &TemplateSlice{
		&Template{
			Name:   "docker_build",
			Source: "github.com/go-vela/atlas/stable/docker_build",
			Type:   "github",
		},
		&Template{
			Name:   "docker_publish",
			Source: "github.com/go-vela/atlas/stable/docker_publish",
			Type:   "github",
		},
	}
	got := new(TemplateSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/template.yml")
	if err != nil {
		t.Errorf("Reading file for UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err != nil {
		t.Errorf("UnmarshalYAML returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalYAML is %v, want %v", got, want)
	}
}

func TestBuild_TemplateSlice_UnmarshalYAML_Invalid(t *testing.T) {
	// run test
	b, err := ioutil.ReadFile("testdata/invalid.yml")
	if err != nil {
		t.Errorf("Reading file for UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, new(TemplateSlice))

	if err == nil {
		t.Errorf("UnmarshalYAML should have returned err")
	}
}
