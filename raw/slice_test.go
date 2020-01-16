// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package raw

import (
	"io/ioutil"
	"reflect"
	"testing"

	"encoding/json"

	yaml "gopkg.in/yaml.v2"
)

func TestRaw_StringSlice_UnmarshalJSON_String(t *testing.T) {
	// setup types
	want := &StringSlice{"foo"}
	got := new(StringSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/string.json")
	if err != nil {
		t.Errorf("Reading file for UnmarshalYAML returned err: %v", err)
	}

	err = json.Unmarshal(b, got)
	if err != nil {
		t.Errorf("UnmarshalJSON returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalJSON is %v, want %v", got, want)
	}
}

func TestRaw_StringSlice_UnmarshalJSON_Slice(t *testing.T) {
	// setup types
	want := &StringSlice{"foo", "bar"}
	got := new(StringSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/slice.json")
	if err != nil {
		t.Errorf("Reading file for UnmarshalYAML returned err: %v", err)
	}

	err = json.Unmarshal(b, got)
	if err != nil {
		t.Errorf("UnmarshalJSON returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalJSON is %v, want %v", got, want)
	}
}

func TestRaw_StringSlice_UnmarshalJSON_Empty(t *testing.T) {
	// setup types
	want := new(StringSlice)
	got := new(StringSlice)

	// run test
	err := got.UnmarshalJSON([]byte(""))
	if err != nil {
		t.Errorf("UnmarshalJSON returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalJSON is %v, want %v", got, want)
	}
}

func TestRaw_StringSlice_UnmarshalJSON_Invalid(t *testing.T) {
	// setup types
	want := new(StringSlice)
	got := new(StringSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/invalid.json")
	if err != nil {
		t.Errorf("Reading file for UnmarshalYAML returned err: %v", err)
	}

	err = json.Unmarshal(b, got)
	if err == nil {
		t.Errorf("UnmarshalJSON should have returned err")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalJSON is %v, want %v", got, want)
	}
}

func TestRaw_StringSlice_UnmarshalYAML_String(t *testing.T) {
	// setup types
	want := &StringSlice{"foo"}
	got := new(StringSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/string.yml")
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

func TestRaw_StringSlice_UnmarshalYAML_Slice(t *testing.T) {
	// setup types
	want := &StringSlice{"foo", "bar"}
	got := new(StringSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/slice.yml")
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

func TestRaw_StringSlice_UnmarshalYAML_Invalid(t *testing.T) {
	// setup types
	want := new(StringSlice)
	got := new(StringSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/invalid.yml")
	if err != nil {
		t.Errorf("Reading file for UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)
	if err == nil {
		t.Errorf("UnmarshalYAML should have returned err")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalYAML is %v, want %v", got, want)
	}
}
