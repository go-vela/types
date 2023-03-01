// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"github.com/go-vela/types/pipeline"
	"reflect"
	"testing"
)

func TestLibrary_Init_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		step *Init
		want *Init
	}{
		{
			step: testInit(),
			want: testInit(),
		},
		{
			step: new(Init),
			want: new(Init),
		},
	}

	// run tests
	for _, test := range tests {
		if test.step.GetID() != test.want.GetID() {
			t.Errorf("GetID is %v, want %v", test.step.GetID(), test.want.GetID())
		}

		if test.step.GetRepoID() != test.want.GetRepoID() {
			t.Errorf("GetRepoID is %v, want %v", test.step.GetRepoID(), test.want.GetRepoID())
		}

		if test.step.GetBuildID() != test.want.GetBuildID() {
			t.Errorf("GetBuildID is %v, want %v", test.step.GetBuildID(), test.want.GetBuildID())
		}

		if test.step.GetNumber() != test.want.GetNumber() {
			t.Errorf("GetNumber is %v, want %v", test.step.GetNumber(), test.want.GetNumber())
		}

		if test.step.GetReporter() != test.want.GetReporter() {
			t.Errorf("GetReporter is %v, want %v", test.step.GetReporter(), test.want.GetReporter())
		}

		if test.step.GetName() != test.want.GetName() {
			t.Errorf("GetName is %v, want %v", test.step.GetName(), test.want.GetName())
		}

		if test.step.GetMimetype() != test.want.GetMimetype() {
			t.Errorf("GetMimetype is %v, want %v", test.step.GetMimetype(), test.want.GetMimetype())
		}
	}
}

func TestLibrary_Init_Setters(t *testing.T) {
	// setup types
	var s *Init

	// setup tests
	tests := []struct {
		step *Init
		want *Init
	}{
		{
			step: testInit(),
			want: testInit(),
		},
		{
			step: s,
			want: new(Init),
		},
	}

	// run tests
	for _, test := range tests {
		test.step.SetID(test.want.GetID())
		test.step.SetRepoID(test.want.GetRepoID())
		test.step.SetBuildID(test.want.GetBuildID())
		test.step.SetNumber(test.want.GetNumber())
		test.step.SetReporter(test.want.GetReporter())
		test.step.SetName(test.want.GetName())
		test.step.SetMimetype(test.want.GetMimetype())

		if test.step.GetID() != test.want.GetID() {
			t.Errorf("SetID is %v, want %v", test.step.GetID(), test.want.GetID())
		}

		if test.step.GetRepoID() != test.want.GetRepoID() {
			t.Errorf("SetRepoID is %v, want %v", test.step.GetRepoID(), test.want.GetRepoID())
		}

		if test.step.GetBuildID() != test.want.GetBuildID() {
			t.Errorf("SetBuildID is %v, want %v", test.step.GetBuildID(), test.want.GetBuildID())
		}

		if test.step.GetNumber() != test.want.GetNumber() {
			t.Errorf("SetNumber is %v, want %v", test.step.GetNumber(), test.want.GetNumber())
		}

		if test.step.GetReporter() != test.want.GetReporter() {
			t.Errorf("SetReporter is %v, want %v", test.step.GetReporter(), test.want.GetReporter())
		}

		if test.step.GetName() != test.want.GetName() {
			t.Errorf("SetName is %v, want %v", test.step.GetName(), test.want.GetName())
		}

		if test.step.GetMimetype() != test.want.GetMimetype() {
			t.Errorf("SetMimetype is %v, want %v", test.step.GetMimetype(), test.want.GetMimetype())
		}
	}
}

func TestLibrary_Init_String(t *testing.T) {
	// setup types
	i := testInit()

	want := fmt.Sprintf(`{
  BuildID: %d,
  ID: %d,
  Mimetype: %s,
  Name: %s,
  Number: %d,
  RepoID: %d,
  Reporter: %s,
}`,
		i.GetBuildID(),
		i.GetID(),
		i.GetMimetype(),
		i.GetName(),
		i.GetNumber(),
		i.GetRepoID(),
		i.GetReporter(),
	)

	// run test
	got := i.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

func TestLibrary_InitFromBuildInit(t *testing.T) {
	// setup types
	s := testInit()
	s.SetReporter("Foobar Runtime")

	// modify fields that aren't set
	s.ID = nil
	s.RepoID = nil
	s.BuildID = nil

	tests := []struct {
		name string
		init *pipeline.Init
		want *Init
	}{
		{
			name: "nil init",
			init: nil,
			want: &Init{},
		},
		{
			name: "empty init",
			init: new(pipeline.Init),
			want: &Init{},
		},
		{
			name: "populated init",
			init: &pipeline.Init{
				Number:   s.GetNumber(),
				Reporter: s.GetReporter(),
				Name:     s.GetName(),
				Mimetype: s.GetMimetype(),
			},
			want: s,
		},
	}

	// run tests
	for _, test := range tests {
		got := InitFromBuildInit(test.init)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("InitFromBuildInit for %s is %v, want %v", test.name, got, test.want)
		}
	}
}

// testInit is a test helper function to create a Init
// type with all fields set to a fake value.
func testInit() *Init {
	s := new(Init)

	s.SetID(1)
	s.SetRepoID(1)
	s.SetBuildID(1)
	s.SetNumber(1)
	s.SetReporter("Kubernetes Runtime")
	s.SetName("clone")
	s.SetMimetype("text/plain")

	return s
}
