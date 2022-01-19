// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/go-vela/types/pipeline"

	"github.com/buildkite/yaml"
)

func TestYaml_Metadata_ToPipeline(t *testing.T) {
	tBool := true
	fBool := false
	// setup tests
	tests := []struct {
		metadata *Metadata
		want     *pipeline.Metadata
	}{
		{
			metadata: &Metadata{
				Template:    false,
				Clone:       &fBool,
				Environment: []string{"steps", "services", "secrets"},
			},
			want: &pipeline.Metadata{
				Template:    false,
				Clone:       false,
				Environment: []string{"steps", "services", "secrets"},
			},
		},
		{
			metadata: &Metadata{
				Template:    false,
				Clone:       &tBool,
				Environment: []string{"steps", "services"},
			},
			want: &pipeline.Metadata{
				Template:    false,
				Clone:       true,
				Environment: []string{"steps", "services"},
			},
		},
		{
			metadata: &Metadata{
				Template:    false,
				Clone:       nil,
				Environment: []string{"steps"},
			},
			want: &pipeline.Metadata{
				Template:    false,
				Clone:       true,
				Environment: []string{"steps"},
			},
		},
	}

	// run tests
	for _, test := range tests {
		got := test.metadata.ToPipeline()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ToPipeline is %v, want %v", got, test.want)
		}
	}
}

func TestYaml_Metadata_HasEnvironment(t *testing.T) {
	// setup tests
	tests := []struct {
		metadata  *Metadata
		container string
		want      bool
	}{
		{
			metadata: &Metadata{
				Environment: []string{"steps", "services", "secrets"},
			},
			container: "steps",
			want:      true,
		},
		{
			metadata: &Metadata{
				Environment: []string{"services", "secrets"},
			},
			container: "services",
			want:      true,
		},
		{
			metadata: &Metadata{
				Environment: []string{"steps", "services", "secrets"},
			},
			container: "notacontainer",
			want:      false,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.metadata.HasEnvironment(test.container)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ToPipeline is %v, want %v", got, test.want)
		}
	}
}

func TestYaml_Metadata_UnmarshalYAML(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		file    string
		want    *Metadata
	}{
		{
			failure: false,
			file:    "testdata/metadata.yml",
			want: &Metadata{
				Template:    false,
				Clone:       nil,
				Environment: []string{"steps", "services", "secrets"},
			},
		},
		{
			file: "testdata/metadata_env.yml",
			want: &Metadata{
				Template:    false,
				Clone:       nil,
				Environment: []string{"steps"},
			},
		},
	}

	// run tests
	for _, test := range tests {
		got := new(Metadata)

		b, err := ioutil.ReadFile(test.file)
		if err != nil {
			t.Errorf("Reading file for UnmarshalYAML returned err: %v", err)
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
