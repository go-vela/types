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

func TestYaml_SecretSlice_ToPipeline(t *testing.T) {
	// setup types
	str := "foo"
	want := &pipeline.SecretSlice{
		&pipeline.Secret{
			Name:   str,
			Key:    str,
			Engine: str,
			Type:   str,
		},
	}

	s := &SecretSlice{
		&Secret{
			Name:   str,
			Key:    str,
			Engine: str,
			Type:   str,
		},
	}

	// run test
	got := s.ToPipeline()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToPipeline is %v, want %v", got, want)
	}
}

func TestYaml_SecretSlice_UnmarshalYAML(t *testing.T) {
	// setup types
	want := &SecretSlice{
		&Secret{
			Name:   "foo",
			Key:    "bar",
			Engine: "native",
			Type:   "repo",
		},
		&Secret{
			Name:   "noKey",
			Key:    "noKey",
			Engine: "native",
			Type:   "repo",
		},
		&Secret{
			Name:   "noType",
			Key:    "bar",
			Engine: "native",
			Type:   "repo",
		},
		&Secret{
			Name:   "noEngine",
			Key:    "bar",
			Engine: "native",
			Type:   "repo",
		},
		&Secret{
			Name:   "noKeyEngineAndType",
			Key:    "noKeyEngineAndType",
			Engine: "native",
			Type:   "repo",
		},
	}
	got := new(SecretSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/secret.yml")
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

func TestYaml_SecretSlice_UnmarshalYAML_Invalid(t *testing.T) {
	// run test
	b, err := ioutil.ReadFile("testdata/invalid.yml")
	if err != nil {
		t.Errorf("Reading file for UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, new(SecretSlice))

	if err == nil {
		t.Errorf("UnmarshalYAML should have returned err")
	}
}

func TestYaml_StepSecretSlice_ToPipeline(t *testing.T) {
	// setup types
	str := "foo"
	want := &pipeline.StepSecretSlice{
		&pipeline.StepSecret{
			Source: str,
			Target: str,
		},
	}

	s := &StepSecretSlice{
		&StepSecret{
			Source: str,
			Target: str,
		},
	}

	// run test
	got := s.ToPipeline()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToPipeline is %v, want %v", got, want)
	}
}

func TestYaml_StepSecretSlice_UnmarshalYAML_Slice(t *testing.T) {
	// setup types
	want := &StepSecretSlice{
		&StepSecret{
			Source: "foo",
			Target: "bar",
		},
		&StepSecret{
			Source: "hello",
			Target: "world",
		},
	}
	got := new(StepSecretSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/step_secret_slice.yml")
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

func TestYaml_StepSecretSlice_UnmarshalYAML_String(t *testing.T) {
	// setup types
	want := &StepSecretSlice{
		&StepSecret{
			Source: "foo",
			Target: "foo",
		},
		&StepSecret{
			Source: "hello",
			Target: "hello",
		},
	}
	got := new(StepSecretSlice)

	// run test
	b, err := ioutil.ReadFile("testdata/step_secret_string.yml")
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

func TestYaml_StepSecretSlice_UnmarshalYAML_Invalid(t *testing.T) {
	// run test
	b, err := ioutil.ReadFile("testdata/invalid.yml")
	if err != nil {
		t.Errorf("Reading file for UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, new(StepSecretSlice))

	if err == nil {
		t.Errorf("UnmarshalYAML should have returned err")
	}
}
