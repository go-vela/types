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
	// setup types
	m := &Metadata{
		Template: false,
	}

	want := &pipeline.Metadata{
		Template: false,
	}

	// run test
	got := m.ToPipeline()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToPipeline is %v, want %v", got, want)
	}
}

func TestYaml_Metadata_UnmarshalYAML(t *testing.T) {
	// setup types
	want := &Metadata{
		Template: false,
	}

	got := new(Metadata)

	// run test
	b, err := ioutil.ReadFile("testdata/metadata.yml")
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
