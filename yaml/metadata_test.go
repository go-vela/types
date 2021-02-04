// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
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
				Template: false,
				Clone:    &fBool,
			},
			want: &pipeline.Metadata{
				Template: false,
				Clone:    false,
			},
		},
		{
			metadata: &Metadata{
				Template: false,
				Clone:    &tBool,
			},
			want: &pipeline.Metadata{
				Template: false,
				Clone:    true,
			},
		},
		{
			metadata: &Metadata{
				Template: false,
				Clone:    nil,
			},
			want: &pipeline.Metadata{
				Template: false,
				Clone:    true,
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

func TestYaml_Metadata_UnmarshalYAML(t *testing.T) {
	// setup tests
	tests := []struct {
		file string
		want *Metadata
	}{
		{
			file: "testdata/metadata.yml",
			want: &Metadata{
				Template: false,
				Clone:    nil,
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
