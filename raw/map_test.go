// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package raw

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestRaw_StringSliceMap_UnmarshalJSON_String(t *testing.T) {
	// setup types
	want := &StringSliceMap{"foo": "bar"}
	got := new(StringSliceMap)

	// run test
	b, err := ioutil.ReadFile("testdata/string_map.json")
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

func TestRaw_StringSliceMap_UnmarshalJSON_Slice(t *testing.T) {
	// setup types
	want := &StringSliceMap{"foo": "bar"}
	got := new(StringSliceMap)

	// run test
	b, err := ioutil.ReadFile("testdata/slice_map.json")
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

func TestRaw_StringSliceMap_UnmarshalJSON_Map(t *testing.T) {
	// setup types
	want := &StringSliceMap{"foo": "bar"}
	got := new(StringSliceMap)

	// run test
	b, err := ioutil.ReadFile("testdata/map.json")
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

func TestRaw_StringSliceMap_UnmarshalJSON_Empty(t *testing.T) {
	// setup types
	want := new(StringSliceMap)
	got := new(StringSliceMap)

	// run test
	err := got.UnmarshalJSON([]byte(""))
	if err != nil {
		t.Errorf("UnmarshalJSON returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalJSON is %v, want %v", got, want)
	}
}

func TestRaw_StringSliceMap_UnmarshalJSON_Invalid(t *testing.T) {
	// setup types
	want := new(StringSliceMap)
	got := new(StringSliceMap)

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

func TestRaw_StringSliceMap_UnmarshalYAML_String(t *testing.T) {
	// setup types
	want := &StringSliceMap{"foo": "bar"}
	got := new(StringSliceMap)

	// run test
	b, err := ioutil.ReadFile("testdata/string_map.yml")
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

func TestRaw_StringSliceMap_UnmarshalYAML_Slice(t *testing.T) {
	// setup types
	want := &StringSliceMap{"foo": "bar"}
	got := new(StringSliceMap)

	// run test
	b, err := ioutil.ReadFile("testdata/slice_map.yml")
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

func TestRaw_StringSliceMap_UnmarshalYAML_Map(t *testing.T) {
	// setup types
	want := &StringSliceMap{"foo": "bar"}
	got := new(StringSliceMap)

	// run test
	b, err := ioutil.ReadFile("testdata/map.yml")
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

func TestRaw_StringSliceMap_UnmarshalYAML_Invalid(t *testing.T) {
	// setup types
	want := new(StringSliceMap)
	got := new(StringSliceMap)

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
