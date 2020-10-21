// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"io/ioutil"
	"reflect"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestBuild_TemplateSlice_UnmarshalYAML(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		file    string
		want    *TemplateSlice
	}{
		{
			failure: false,
			file:    "testdata/template.yml",
			want: &TemplateSlice{
				{
					Name:   "docker_build",
					Source: "github.com/go-vela/atlas/stable/docker_create",
					Type:   "github",
				},
				{
					Name:   "docker_build",
					Source: "github.com/go-vela/atlas/stable/docker_build",
					Format: "go",
					Type:   "github",
				},
				{
					Name:   "docker_publish",
					Source: "github.com/go-vela/atlas/stable/docker_publish",
					Format: "starlark",
					Type:   "github",
				},
			},
		},
		{
			failure: true,
			file:    "testdata/invalid.yml",
			want:    nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := new(TemplateSlice)

		b, err := ioutil.ReadFile(test.file)
		if err != nil {
			t.Errorf("unable to read file: %v", err)
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
