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

func TestYaml_Metadata_ToPipeline(t *testing.T) {
	// setup tests
	tests := []struct {
		metadata *Metadata
		want     *pipeline.Metadata
	}{
		{
			metadata: &Metadata{
				Template: false,
			},
			want: &pipeline.Metadata{
				Template: false,
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
