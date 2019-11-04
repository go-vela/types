// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
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

func TestYaml_UlimitSlice_ToPipeline(t *testing.T) {
	// setup types
	num := int64(1)
	str := "foo"
	want := &pipeline.UlimitSlice{
		&pipeline.Ulimit{
			Name: str,
			Soft: num,
			Hard: num,
		},
	}

	u := &UlimitSlice{
		&Ulimit{
			Name: str,
			Soft: num,
			Hard: num,
		},
	}

	// run test
	got := u.ToPipeline()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToPipeline is %v, want %v", got, want)
	}
}

func TestYaml_UlimitSlice_UnmarshalYAML_Slice(t *testing.T) {
	// setup types
	want := &UlimitSlice{
		&Ulimit{
			Name: "foo",
			Soft: 1024,
			Hard: 1024,
		},
		&Ulimit{
			Name: "bar",
			Soft: 1024,
			Hard: 2048,
		},
	}
	got := new(UlimitSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/ulimit_slice.yml")
	if err != nil {
		t.Errorf("Reading file for UlimitSlice UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err != nil {
		t.Errorf("UlimitSlice UnmarshalYAML returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UlimitSlice UnmarshalYAML is %v, want %v", got, want)
	}
}

func TestYaml_UlimitSlice_UnmarshalYAML_String(t *testing.T) {
	// setup types
	want := &UlimitSlice{
		&Ulimit{
			Name: "foo",
			Soft: 1024,
			Hard: 1024,
		},
		&Ulimit{
			Name: "bar",
			Soft: 1024,
			Hard: 2048,
		},
	}
	got := new(UlimitSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/ulimit_string.yml")
	if err != nil {
		t.Errorf("Reading file for UlimitSlice UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err != nil {
		t.Errorf("UlimitSlice UnmarshalYAML returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UlimitSlice UnmarshalYAML is %v, want %v", got, want)
	}
}

func TestYaml_UlimitSlice_UnmarshalYAML_Slice_Error(t *testing.T) {
	// setup types
	got := new(UlimitSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/invalid.yml")
	if err != nil {
		t.Errorf("Reading file for VolumeSlice UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err == nil {
		t.Errorf("UlimitSlice UnmarshalYAML should have returned err")
	}
}

func TestYaml_UlimitSlice_UnmarshalYAML_String_EqualError(t *testing.T) {
	// setup types
	got := new(UlimitSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/ulimit_equal_error.yml")
	if err != nil {
		t.Errorf("Reading file for UlimitSlice UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err == nil {
		t.Errorf("UlimitSlice UnmarshalYAML should have returned err")
	}
}

func TestYaml_UlimitSlice_UnmarshalYAML_String_ColonError(t *testing.T) {
	// setup types
	got := new(UlimitSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/ulimit_colon_error.yml")
	if err != nil {
		t.Errorf("Reading file for UlimitSlice UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err == nil {
		t.Errorf("UlimitSlice UnmarshalYAML should have returned err")
	}
}

func TestYaml_UlimitSlice_UnmarshalYAML_String_SoftLimitError(t *testing.T) {
	// setup types
	got := new(UlimitSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/ulimit_softlimit_error.yml")
	if err != nil {
		t.Errorf("Reading file for UlimitSlice UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err == nil {
		t.Errorf("UlimitSlice UnmarshalYAML should have returned err")
	}
}

func TestYaml_UlimitSlice_UnmarshalYAML_String_HardLimit1Error(t *testing.T) {
	// setup types
	got := new(UlimitSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/ulimit_hardlimit1_error.yml")
	if err != nil {
		t.Errorf("Reading file for UlimitSlice UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err == nil {
		t.Errorf("UlimitSlice UnmarshalYAML should have returned err")
	}
}

func TestYaml_UlimitSlice_UnmarshalYAML_String_HardLimit2Error(t *testing.T) {
	// setup types
	got := new(UlimitSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/ulimit_hardlimit2_error.yml")
	if err != nil {
		t.Errorf("Reading file for UlimitSlice UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err == nil {
		t.Errorf("UlimitSlice UnmarshalYAML should have returned err")
	}
}
