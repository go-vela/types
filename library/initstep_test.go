// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-vela/types/pipeline"
)

func TestLibrary_InitStep_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		initStep *InitStep
		want     *InitStep
	}{
		{
			initStep: testInitStep(),
			want:     testInitStep(),
		},
		{
			initStep: new(InitStep),
			want:     new(InitStep),
		},
	}

	// run tests
	for _, test := range tests {
		if test.initStep.GetID() != test.want.GetID() {
			t.Errorf("GetID is %v, want %v", test.initStep.GetID(), test.want.GetID())
		}

		if test.initStep.GetRepoID() != test.want.GetRepoID() {
			t.Errorf("GetRepoID is %v, want %v", test.initStep.GetRepoID(), test.want.GetRepoID())
		}

		if test.initStep.GetBuildID() != test.want.GetBuildID() {
			t.Errorf("GetBuildID is %v, want %v", test.initStep.GetBuildID(), test.want.GetBuildID())
		}

		if test.initStep.GetNumber() != test.want.GetNumber() {
			t.Errorf("GetNumber is %v, want %v", test.initStep.GetNumber(), test.want.GetNumber())
		}

		if test.initStep.GetReporter() != test.want.GetReporter() {
			t.Errorf("GetReporter is %v, want %v", test.initStep.GetReporter(), test.want.GetReporter())
		}

		if test.initStep.GetName() != test.want.GetName() {
			t.Errorf("GetName is %v, want %v", test.initStep.GetName(), test.want.GetName())
		}

		if test.initStep.GetMimetype() != test.want.GetMimetype() {
			t.Errorf("GetMimetype is %v, want %v", test.initStep.GetMimetype(), test.want.GetMimetype())
		}
	}
}

func TestLibrary_InitStep_Setters(t *testing.T) {
	// setup types
	var s *InitStep

	// setup tests
	tests := []struct {
		initStep *InitStep
		want     *InitStep
	}{
		{
			initStep: testInitStep(),
			want:     testInitStep(),
		},
		{
			initStep: s,
			want:     new(InitStep),
		},
	}

	// run tests
	for _, test := range tests {
		test.initStep.SetID(test.want.GetID())
		test.initStep.SetRepoID(test.want.GetRepoID())
		test.initStep.SetBuildID(test.want.GetBuildID())
		test.initStep.SetNumber(test.want.GetNumber())
		test.initStep.SetReporter(test.want.GetReporter())
		test.initStep.SetName(test.want.GetName())
		test.initStep.SetMimetype(test.want.GetMimetype())

		if test.initStep.GetID() != test.want.GetID() {
			t.Errorf("SetID is %v, want %v", test.initStep.GetID(), test.want.GetID())
		}

		if test.initStep.GetRepoID() != test.want.GetRepoID() {
			t.Errorf("SetRepoID is %v, want %v", test.initStep.GetRepoID(), test.want.GetRepoID())
		}

		if test.initStep.GetBuildID() != test.want.GetBuildID() {
			t.Errorf("SetBuildID is %v, want %v", test.initStep.GetBuildID(), test.want.GetBuildID())
		}

		if test.initStep.GetNumber() != test.want.GetNumber() {
			t.Errorf("SetNumber is %v, want %v", test.initStep.GetNumber(), test.want.GetNumber())
		}

		if test.initStep.GetReporter() != test.want.GetReporter() {
			t.Errorf("SetReporter is %v, want %v", test.initStep.GetReporter(), test.want.GetReporter())
		}

		if test.initStep.GetName() != test.want.GetName() {
			t.Errorf("SetName is %v, want %v", test.initStep.GetName(), test.want.GetName())
		}

		if test.initStep.GetMimetype() != test.want.GetMimetype() {
			t.Errorf("SetMimetype is %v, want %v", test.initStep.GetMimetype(), test.want.GetMimetype())
		}
	}
}

func TestLibrary_InitStep_String(t *testing.T) {
	// setup types
	i := testInitStep()

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

func TestLibrary_InitStepFromBuildInitStep(t *testing.T) {
	// setup types
	i := testInitStep()
	i.SetReporter("Foobar Runtime")

	// modify fields that aren't set
	i.ID = nil
	i.RepoID = nil
	i.BuildID = nil

	tests := []struct {
		name     string
		initStep *pipeline.InitStep
		want     *InitStep
	}{
		{
			name:     "nil InitStep",
			initStep: nil,
			want:     &InitStep{},
		},
		{
			name:     "empty InitStep",
			initStep: new(pipeline.InitStep),
			want:     &InitStep{},
		},
		{
			name: "populated InitStep",
			initStep: &pipeline.InitStep{
				Number:   i.GetNumber(),
				Reporter: i.GetReporter(),
				Name:     i.GetName(),
				Mimetype: i.GetMimetype(),
			},
			want: i,
		},
	}

	// run tests
	for _, test := range tests {
		got := InitStepFromBuildInitStep(test.initStep)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("InitStepFromBuildInitStep for %s is %v, want %v", test.name, got, test.want)
		}
	}
}

// testInitStep is a test helper function to create a InitStep
// type with all fields set to a fake value.
func testInitStep() *InitStep {
	i := new(InitStep)

	i.SetID(1)
	i.SetRepoID(1)
	i.SetBuildID(1)
	i.SetNumber(1)
	i.SetReporter("Kubernetes Runtime")
	i.SetName("clone")
	i.SetMimetype("text/plain")

	return i
}
