// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package version

import (
	"fmt"
	"reflect"
	"testing"
)

func TestVersion_Version_Meta(t *testing.T) {
	// setup types
	v := &Version{
		Canonical:  "v1.2.3",
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Metadata: Metadata{
			Architecture:    "amd64",
			BuildDate:       "1970-1-1T00:00:00Z",
			Compiler:        "gc",
			GitCommit:       "abcdef123456789",
			GoVersion:       "1.16.0",
			OperatingSystem: "llinux",
		},
	}

	want := v.Metadata.String()

	// run test
	got := v.Meta()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

func TestVersion_Version_Semantic(t *testing.T) {
	// setup types
	v := &Version{
		Canonical:  "v1.2.3",
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Metadata: Metadata{
			Architecture:    "amd64",
			BuildDate:       "1970-1-1T00:00:00Z",
			Compiler:        "gc",
			GitCommit:       "abcdef123456789",
			GoVersion:       "1.16.0",
			OperatingSystem: "llinux",
		},
	}

	want := v.Canonical

	// run test
	got := v.Semantic()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

func TestVersion_Version_String(t *testing.T) {
	// setup types
	v := &Version{
		Canonical:  "v1.2.3",
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Metadata: Metadata{
			Architecture:    "amd64",
			BuildDate:       "1970-1-1T00:00:00Z",
			Compiler:        "gc",
			GitCommit:       "abcdef123456789",
			GoVersion:       "1.16.0",
			OperatingSystem: "llinux",
		},
	}

	want := fmt.Sprintf(
		versionFormat,
		v.Canonical,
		v.Major,
		v.Minor,
		v.Patch,
		v.PreRelease,
		v.Metadata.Architecture,
		v.Metadata.BuildDate,
		v.Metadata.Compiler,
		v.Metadata.GitCommit,
		v.Metadata.GoVersion,
		v.Metadata.OperatingSystem,
	)

	// run test
	got := v.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}
