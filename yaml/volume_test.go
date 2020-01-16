// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/go-vela/types/pipeline"
	yaml "gopkg.in/yaml.v2"
)

func TestYaml_VolumeSlice_ToPipeline(t *testing.T) {
	// setup types
	str := "foo"
	want := &pipeline.VolumeSlice{
		&pipeline.Volume{
			Source:      str,
			Destination: str,
			AccessMode:  str,
		},
	}

	v := &VolumeSlice{
		&Volume{
			Source:      str,
			Destination: str,
			AccessMode:  str,
		},
	}

	// run test
	got := v.ToPipeline()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToPipeline is %v, want %v", got, want)
	}
}

func TestYaml_VolumeSlice_UnmarshalYAML_Slice(t *testing.T) {
	// setup types
	want := &VolumeSlice{
		&Volume{
			Source:      "/foo",
			Destination: "/foo",
			AccessMode:  "ro",
		},
		&Volume{
			Source:      "/foo",
			Destination: "/bar",
			AccessMode:  "ro",
		},
		&Volume{
			Source:      "/foo",
			Destination: "/foobar",
			AccessMode:  "ro",
		},
	}
	got := new(VolumeSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/volume_slice.yml")
	if err != nil {
		t.Errorf("Reading file for VolumeSlice UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err != nil {
		t.Errorf("VolumeSlice UnmarshalYAML returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("VolumeSlice UnmarshalYAML is %v, want %v", got, want)
	}
}

func TestYaml_VolumeSlice_UnmarshalYAML_String(t *testing.T) {
	// setup types
	want := &VolumeSlice{
		&Volume{
			Source:      "/foo",
			Destination: "/foo",
			AccessMode:  "ro",
		},
		&Volume{
			Source:      "/foo",
			Destination: "/bar",
			AccessMode:  "ro",
		},
		&Volume{
			Source:      "/foo",
			Destination: "/foobar",
			AccessMode:  "ro",
		},
	}
	got := new(VolumeSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/volume_string.yml")
	if err != nil {
		t.Errorf("Reading file for VolumeSlice UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err != nil {
		t.Errorf("VolumeSlice UnmarshalYAML returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("VolumeSlice UnmarshalYAML is %v, want %v", got, want)
	}
}

func TestYaml_VolumeSlice_UnmarshalYAML_Slice_Error(t *testing.T) {
	// setup types
	got := new(VolumeSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/invalid.yml")
	if err != nil {
		t.Errorf("Reading file for VolumeSlice UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err == nil {
		t.Errorf("VolumeSlice UnmarshalYAML should have returned err")
	}
}

func TestYaml_VolumeSlice_UnmarshalYAML_String_Error(t *testing.T) {
	// setup types
	got := new(VolumeSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/volume_error.yml")
	if err != nil {
		t.Errorf("Reading file for VolumeSlice UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err == nil {
		t.Errorf("VolumeSlice UnmarshalYAML should have returned err")
	}
}
