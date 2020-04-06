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

func TestYaml_Ruleset_ToPipeline(t *testing.T) {
	// setup types
	str := "foo"
	slice := []string{"foo"}
	want := &pipeline.Ruleset{
		If: pipeline.Rules{
			Branch:  slice,
			Comment: slice,
			Event:   slice,
			Path:    slice,
			Repo:    slice,
			Status:  slice,
			Tag:     slice,
		},
		Unless: pipeline.Rules{
			Branch:  slice,
			Comment: slice,
			Event:   slice,
			Path:    slice,
			Repo:    slice,
			Status:  slice,
			Tag:     slice,
		},
		Operator: str,
		Continue: false,
	}

	r := &Ruleset{
		If: Rules{
			Branch:  slice,
			Comment: slice,
			Event:   slice,
			Path:    slice,
			Repo:    slice,
			Status:  slice,
			Tag:     slice,
		},
		Unless: Rules{
			Branch:  slice,
			Comment: slice,
			Event:   slice,
			Path:    slice,
			Repo:    slice,
			Status:  slice,
			Tag:     slice,
		},
		Operator: str,
		Continue: false,
	}

	// run test
	got := r.ToPipeline()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToPipeline is %v, want %v", got, want)
	}
}

func TestYaml_Ruleset_UnmarshalYAML_Simple(t *testing.T) {
	// setup types
	want := &Ruleset{
		If: Rules{
			Branch:  []string{"master"},
			Event:   []string{"push"},
			Path:    []string{"foo.txt", "/foo/bar.txt"},
			Comment: []string{"ok to test", "rerun"},
		},
		Operator: "and",
		Continue: true,
	}
	got := new(Ruleset)

	// run test
	b, err := ioutil.ReadFile("testdata/ruleset_simple.yml")
	if err != nil {
		t.Errorf("Reading file for Ruleset UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err != nil {
		t.Errorf("Ruleset UnmarshalYAML returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ruleset UnmarshalYAML is %v, want %v", got, want)
	}
}

func TestYaml_Ruleset_UnmarshalYAML_Advanced(t *testing.T) {
	// setup types
	want := &Ruleset{
		If: Rules{
			Branch: []string{"master"},
			Event:  []string{"push"},
		},
		Unless: Rules{
			Event: []string{"deployment", "pull_request"},
			Path:  []string{"foo.txt", "/foo/bar.txt"},
		},
		Operator: "or",
		Continue: true,
	}
	got := new(Ruleset)

	// run test
	b, err := ioutil.ReadFile("testdata/ruleset_advanced.yml")
	if err != nil {
		t.Errorf("Reading file for Ruleset UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err != nil {
		t.Errorf("Ruleset UnmarshalYAML returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ruleset UnmarshalYAML is %v, want %v", got, want)
	}
}

func TestYaml_Rules_ToPipeline(t *testing.T) {
	// setup types
	slice := []string{"foo"}
	want := &pipeline.Rules{
		Branch:  slice,
		Comment: slice,
		Event:   slice,
		Path:    slice,
		Repo:    slice,
		Status:  slice,
		Tag:     slice,
	}

	r := &Rules{
		Branch:  slice,
		Comment: slice,
		Event:   slice,
		Path:    slice,
		Repo:    slice,
		Status:  slice,
		Tag:     slice,
	}

	// run test
	got := r.ToPipeline()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToPipeline is %v, want %v", got, want)
	}
}

func TestYaml_Rules_UnmarshalYAML(t *testing.T) {
	// setup types
	want := &Rules{
		Branch:  []string{"master"},
		Event:   []string{"push"},
		Path:    []string{"foo.txt", "/foo/bar.txt"},
		Comment: []string{"ok to test", "rerun"},
	}
	got := new(Rules)

	// run test
	b, err := ioutil.ReadFile("testdata/ruleset_simple.yml")
	if err != nil {
		t.Errorf("Reading file for Rules UnmarshalYAML returned err: %v", err)
	}

	err = yaml.Unmarshal(b, got)

	if err != nil {
		t.Errorf("Rules UnmarshalYAML returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Rules UnmarshalYAML is %v, want %v", got, want)
	}
}

func TestYaml_Rules_UnmarshalYAML_Error(t *testing.T) {
	// setup types
	r := new(Rules)

	// run test
	err := yaml.Unmarshal([]byte("!@#$%^&*()"), r)

	if err == nil {
		t.Errorf("Rules UnmarshalYAML should have returned err")
	}
}
