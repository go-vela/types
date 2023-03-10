// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"os"
	"reflect"
	"testing"

	"github.com/buildkite/yaml"
	"github.com/go-vela/types/pipeline"
)

func TestYaml_Ruleset_ToPipeline(t *testing.T) {
	// setup tests
	tests := []struct {
		ruleset *Ruleset
		want    *pipeline.Ruleset
	}{
		{
			ruleset: &Ruleset{
				If: Rules{
					Branch:  []string{"master"},
					Comment: []string{"test comment"},
					Event:   []string{"push"},
					Path:    []string{"foo.txt"},
					Repo:    []string{"github/octocat"},
					Status:  []string{"success"},
					Tag:     []string{"v0.1.0"},
					Target:  []string{"production"},
				},
				Unless: Rules{
					Branch:  []string{"master"},
					Comment: []string{"real comment"},
					Event:   []string{"pull_request"},
					Path:    []string{"bar.txt"},
					Repo:    []string{"github/octocat"},
					Status:  []string{"failure"},
					Tag:     []string{"v0.2.0"},
					Target:  []string{"production"},
				},
				Matcher:  "filepath",
				Operator: "and",
				Continue: false,
			},
			want: &pipeline.Ruleset{
				If: pipeline.Rules{
					Branch:  []string{"master"},
					Comment: []string{"test comment"},
					Event:   []string{"push"},
					Path:    []string{"foo.txt"},
					Repo:    []string{"github/octocat"},
					Status:  []string{"success"},
					Tag:     []string{"v0.1.0"},
					Target:  []string{"production"},
				},
				Unless: pipeline.Rules{
					Branch:  []string{"master"},
					Comment: []string{"real comment"},
					Event:   []string{"pull_request"},
					Path:    []string{"bar.txt"},
					Repo:    []string{"github/octocat"},
					Status:  []string{"failure"},
					Tag:     []string{"v0.2.0"},
					Target:  []string{"production"},
				},
				Matcher:  "filepath",
				Operator: "and",
				Continue: false,
			},
		},
	}

	// run tests
	for _, test := range tests {
		got := test.ruleset.ToPipeline()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ToPipeline is %v, want %v", got, test.want)
		}
	}
}

func TestYaml_Ruleset_UnmarshalYAML(t *testing.T) {
	// setup tests
	tests := []struct {
		file string
		want *Ruleset
	}{
		{
			file: "testdata/ruleset_simple.yml",
			want: &Ruleset{
				If: Rules{
					Branch:  []string{"master"},
					Comment: []string{"test comment"},
					Event:   []string{"push"},
					Path:    []string{"foo.txt"},
					Repo:    []string{"github/octocat"},
					Status:  []string{"success"},
					Tag:     []string{"v0.1.0"},
					Target:  []string{"production"},
				},
				Matcher:  "filepath",
				Operator: "and",
				Continue: true,
			},
		},
		{
			file: "testdata/ruleset_advanced.yml",
			want: &Ruleset{
				If: Rules{
					Branch: []string{"master"},
					Event:  []string{"push"},
					Tag:    []string{"^refs/tags/(\\d+\\.)+\\d+$"},
				},
				Unless: Rules{
					Event: []string{"deployment", "pull_request:opened", "pull_request:synchronize", "comment:created", "comment:edited"},
					Path:  []string{"foo.txt", "/foo/bar.txt"},
				},
				Matcher:  "regexp",
				Operator: "or",
				Continue: true,
			},
		},
		{
			file: "testdata/ruleset_regex.yml",
			want: &Ruleset{
				If: Rules{
					Branch: []string{"master"},
					Event:  []string{"tag"},
					Tag:    []string{"^refs/tags/(\\d+\\.)+\\d+$"},
				},
				Operator: "and",
				Matcher:  "regex",
			},
		},
		{
			file: "testdata/ruleset_release.yml",
			want: &Ruleset{
				If: Rules{
					Event: []string{"release:released"},
				},
				Matcher:  "filepath",
				Operator: "and",
				Continue: true,
			},
		},
	}

	// run tests
	for _, test := range tests {
		got := new(Ruleset)

		b, err := os.ReadFile(test.file)
		if err != nil {
			t.Errorf("unable to read file: %v", err)
		}

		err = yaml.Unmarshal(b, got)

		if err != nil {
			t.Errorf("UnmarshalYAML returned err: %v", err)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("UnmarshalYAML is %v, want %v", got, test.want)
		}
	}
}

func TestYaml_Rules_ToPipeline(t *testing.T) {
	// setup tests
	tests := []struct {
		rules *Rules
		want  *pipeline.Rules
	}{
		{
			rules: &Rules{
				Branch:  []string{"master"},
				Comment: []string{"test comment"},
				Event:   []string{"push"},
				Path:    []string{"foo.txt"},
				Repo:    []string{"github/octocat"},
				Status:  []string{"success"},
				Tag:     []string{"v0.1.0"},
				Target:  []string{"production"},
			},
			want: &pipeline.Rules{
				Branch:  []string{"master"},
				Comment: []string{"test comment"},
				Event:   []string{"push"},
				Path:    []string{"foo.txt"},
				Repo:    []string{"github/octocat"},
				Status:  []string{"success"},
				Tag:     []string{"v0.1.0"},
				Target:  []string{"production"},
			},
		},
	}

	// run tests
	for _, test := range tests {
		got := test.rules.ToPipeline()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ToPipeline is %v, want %v", got, test.want)
		}
	}
}

func TestYaml_Rules_UnmarshalYAML(t *testing.T) {
	// setup types
	var (
		b   []byte
		err error
	)

	// setup tests
	tests := []struct {
		failure bool
		file    string
		want    *Rules
	}{
		{
			failure: false,
			file:    "testdata/ruleset_simple.yml",
			want: &Rules{
				Branch:  []string{"master"},
				Comment: []string{"test comment"},
				Event:   []string{"push"},
				Path:    []string{"foo.txt"},
				Repo:    []string{"github/octocat"},
				Status:  []string{"success"},
				Tag:     []string{"v0.1.0"},
				Target:  []string{"production"},
			},
		},
		{
			failure: true,
			file:    "",
			want:    nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := new(Rules)

		if len(test.file) > 0 {
			b, err = os.ReadFile(test.file)
			if err != nil {
				t.Errorf("unable to read file: %v", err)
			}
		} else {
			b = []byte("``")
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
