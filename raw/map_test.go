// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package raw

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/goccy/go-yaml"
)

func TestRaw_StringSliceMap_UnmarshalJSON(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		file    string
		want    *StringSliceMap
	}{
		{
			failure: false,
			file:    "testdata/string_map.json",
			want:    &StringSliceMap{"foo": "bar"},
		},
		{
			failure: false,
			file:    "testdata/slice_map.json",
			want:    &StringSliceMap{"foo": "bar"},
		},
		{
			failure: false,
			file:    "testdata/map.json",
			want:    &StringSliceMap{"foo": "bar"},
		},
		{
			failure: false,
			file:    "",
			want:    new(StringSliceMap),
		},
		{
			failure: true,
			file:    "testdata/invalid.json",
			want:    nil,
		},
	}

	// run tests
	for _, test := range tests {
		var (
			err error

			b   = []byte{}
			got = new(StringSliceMap)
		)

		if len(test.file) > 0 {
			b, err = ioutil.ReadFile(test.file)
			if err != nil {
				t.Errorf("unable to read %s file: %v", test.file, err)
			}
		}

		err = got.UnmarshalJSON(b)

		if test.failure {
			if err == nil {
				t.Errorf("UnmarshalJSON should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("UnmarshalJSON returned err: %v", err)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("UnmarshalJSON is %v, want %v", got, test.want)
		}
	}
}

func TestRaw_StringSliceMap_UnmarshalYAML(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		file    string
		want    *StringSliceMap
	}{
		{
			failure: false,
			file:    "testdata/string_map.yml",
			want:    &StringSliceMap{"foo": "bar"},
		},
		{
			failure: false,
			file:    "testdata/slice_map.yml",
			want:    &StringSliceMap{"foo": "bar"},
		},
		{
			failure: false,
			file:    "testdata/map.yml",
			want:    &StringSliceMap{"foo": "bar"},
		},
		{
			failure: true,
			file:    "testdata/invalid.yml",
			want:    nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := new(StringSliceMap)

		b, err := ioutil.ReadFile(test.file)
		if err != nil {
			t.Errorf("unable to read %s file: %v", test.file, err)
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
