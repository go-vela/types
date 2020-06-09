// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_Step_Environment(t *testing.T) {
	// setup types
	want := map[string]string{
		"VELA_STEP_CREATED":      "1563474076",
		"VELA_STEP_DISTRIBUTION": "linux",
		"VELA_STEP_EXIT_CODE":    "0",
		"VELA_STEP_FINISHED":     "1563474079",
		"VELA_STEP_HOST":         "example.company.com",
		"VELA_STEP_IMAGE":        "target/vela-git:v0.3.0",
		"VELA_STEP_NAME":         "clone",
		"VELA_STEP_NUMBER":       "1",
		"VELA_STEP_RUNTIME":      "docker",
		"VELA_STEP_STAGE":        "",
		"VELA_STEP_STARTED":      "1563474078",
		"VELA_STEP_STATUS":       "running",
	}

	// run test
	got := testStep().Environment()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Environment is %v, want %v", got, want)
	}
}

func TestStep_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		step *Step
		want *Step
	}{
		{
			step: testStep(),
			want: testStep(),
		},
		{
			step: new(Step),
			want: new(Step),
		},
	}

	// run tests
	for _, test := range tests {
		if test.step.GetID() != test.want.GetID() {
			t.Errorf("GetID is %v, want %v", test.step.GetID(), test.want.GetID())
		}

		if test.step.GetBuildID() != test.want.GetBuildID() {
			t.Errorf("GetBuildID is %v, want %v", test.step.GetBuildID(), test.want.GetBuildID())
		}

		if test.step.GetRepoID() != test.want.GetRepoID() {
			t.Errorf("GetRepoID is %v, want %v", test.step.GetRepoID(), test.want.GetRepoID())
		}

		if test.step.GetNumber() != test.want.GetNumber() {
			t.Errorf("GetNumber is %v, want %v", test.step.GetNumber(), test.want.GetNumber())
		}

		if test.step.GetName() != test.want.GetName() {
			t.Errorf("GetName is %v, want %v", test.step.GetName(), test.want.GetName())
		}

		if test.step.GetImage() != test.want.GetImage() {
			t.Errorf("GetImage is %v, want %v", test.step.GetImage(), test.want.GetImage())
		}

		if test.step.GetStage() != test.want.GetStage() {
			t.Errorf("GetStage is %v, want %v", test.step.GetStage(), test.want.GetStage())
		}

		if test.step.GetStatus() != test.want.GetStatus() {
			t.Errorf("GetStatus is %v, want %v", test.step.GetStatus(), test.want.GetStatus())
		}

		if test.step.GetError() != test.want.GetError() {
			t.Errorf("GetError is %v, want %v", test.step.GetError(), test.want.GetError())
		}

		if test.step.GetExitCode() != test.want.GetExitCode() {
			t.Errorf("GetExitCode is %v, want %v", test.step.GetExitCode(), test.want.GetExitCode())
		}

		if test.step.GetCreated() != test.want.GetCreated() {
			t.Errorf("GetCreated is %v, want %v", test.step.GetCreated(), test.want.GetCreated())
		}

		if test.step.GetStarted() != test.want.GetStarted() {
			t.Errorf("GetStarted is %v, want %v", test.step.GetStarted(), test.want.GetStarted())
		}

		if test.step.GetFinished() != test.want.GetFinished() {
			t.Errorf("GetFinished is %v, want %v", test.step.GetFinished(), test.want.GetFinished())
		}

		if test.step.GetHost() != test.want.GetHost() {
			t.Errorf("GetHost is %v, want %v", test.step.GetHost(), test.want.GetHost())
		}

		if test.step.GetRuntime() != test.want.GetRuntime() {
			t.Errorf("GetRuntime is %v, want %v", test.step.GetRuntime(), test.want.GetRuntime())
		}

		if test.step.GetDistribution() != test.want.GetDistribution() {
			t.Errorf("GetDistribution is %v, want %v", test.step.GetDistribution(), test.want.GetDistribution())
		}
	}
}

func TestLibrary_Step_Setters(t *testing.T) {
	// setup types
	var s *Step

	// setup tests
	tests := []struct {
		step *Step
		want *Step
	}{
		{
			step: testStep(),
			want: testStep(),
		},
		{
			step: s,
			want: new(Step),
		},
	}

	// run tests
	for _, test := range tests {
		test.step.SetID(test.want.GetID())
		test.step.SetBuildID(test.want.GetBuildID())
		test.step.SetRepoID(test.want.GetRepoID())
		test.step.SetNumber(test.want.GetNumber())
		test.step.SetName(test.want.GetName())
		test.step.SetImage(test.want.GetImage())
		test.step.SetStage(test.want.GetStage())
		test.step.SetStatus(test.want.GetStatus())
		test.step.SetError(test.want.GetError())
		test.step.SetExitCode(test.want.GetExitCode())
		test.step.SetCreated(test.want.GetCreated())
		test.step.SetStarted(test.want.GetStarted())
		test.step.SetFinished(test.want.GetFinished())
		test.step.SetHost(test.want.GetHost())
		test.step.SetRuntime(test.want.GetRuntime())
		test.step.SetDistribution(test.want.GetDistribution())

		if test.step.GetID() != test.want.GetID() {
			t.Errorf("SetID is %v, want %v", test.step.GetID(), test.want.GetID())
		}

		if test.step.GetBuildID() != test.want.GetBuildID() {
			t.Errorf("SetBuildID is %v, want %v", test.step.GetBuildID(), test.want.GetBuildID())
		}

		if test.step.GetRepoID() != test.want.GetRepoID() {
			t.Errorf("SetRepoID is %v, want %v", test.step.GetRepoID(), test.want.GetRepoID())
		}

		if test.step.GetNumber() != test.want.GetNumber() {
			t.Errorf("SetNumber is %v, want %v", test.step.GetNumber(), test.want.GetNumber())
		}

		if test.step.GetName() != test.want.GetName() {
			t.Errorf("SetName is %v, want %v", test.step.GetName(), test.want.GetName())
		}

		if test.step.GetImage() != test.want.GetImage() {
			t.Errorf("SetImage is %v, want %v", test.step.GetImage(), test.want.GetImage())
		}

		if test.step.GetStage() != test.want.GetStage() {
			t.Errorf("SetStage is %v, want %v", test.step.GetStage(), test.want.GetStage())
		}

		if test.step.GetStatus() != test.want.GetStatus() {
			t.Errorf("SetStatus is %v, want %v", test.step.GetStatus(), test.want.GetStatus())
		}

		if test.step.GetError() != test.want.GetError() {
			t.Errorf("SetError is %v, want %v", test.step.GetError(), test.want.GetError())
		}

		if test.step.GetExitCode() != test.want.GetExitCode() {
			t.Errorf("SetExitCode is %v, want %v", test.step.GetExitCode(), test.want.GetExitCode())
		}

		if test.step.GetCreated() != test.want.GetCreated() {
			t.Errorf("SetCreated is %v, want %v", test.step.GetCreated(), test.want.GetCreated())
		}

		if test.step.GetStarted() != test.want.GetStarted() {
			t.Errorf("SetStarted is %v, want %v", test.step.GetStarted(), test.want.GetStarted())
		}

		if test.step.GetFinished() != test.want.GetFinished() {
			t.Errorf("SetFinished is %v, want %v", test.step.GetFinished(), test.want.GetFinished())
		}

		if test.step.GetHost() != test.want.GetHost() {
			t.Errorf("SetHost is %v, want %v", test.step.GetHost(), test.want.GetHost())
		}

		if test.step.GetRuntime() != test.want.GetRuntime() {
			t.Errorf("SetRuntime is %v, want %v", test.step.GetRuntime(), test.want.GetRuntime())
		}

		if test.step.GetDistribution() != test.want.GetDistribution() {
			t.Errorf("SetDistribution is %v, want %v", test.step.GetDistribution(), test.want.GetDistribution())
		}
	}
}

func TestStep_String(t *testing.T) {
	// setup types
	s := testStep()

	want := fmt.Sprintf(`{
  BuildID: %d,
  Created: %d,
  Distribution: %s,
  Error: %s,
  ExitCode: %d,
  Finished: %d,
  Host: %s,
  ID: %d,
  Image: %s,
  Name: %s,
  Number: %d,
  RepoID: %d,
  Runtime: %s,
  Stage: %s,
  Started: %d,
  Status: %s,
}`,
		s.GetBuildID(),
		s.GetCreated(),
		s.GetDistribution(),
		s.GetError(),
		s.GetExitCode(),
		s.GetFinished(),
		s.GetHost(),
		s.GetID(),
		s.GetImage(),
		s.GetName(),
		s.GetNumber(),
		s.GetRepoID(),
		s.GetRuntime(),
		s.GetStage(),
		s.GetStarted(),
		s.GetStatus(),
	)

	// run test
	got := s.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

// testStep is a test helper function to create a Step
// type with all fields set to a fake value.
func testStep() *Step {
	s := new(Step)

	s.SetID(1)
	s.SetBuildID(1)
	s.SetRepoID(1)
	s.SetNumber(1)
	s.SetName("clone")
	s.SetImage("target/vela-git:v0.3.0")
	s.SetStatus("running")
	s.SetExitCode(0)
	s.SetCreated(1563474076)
	s.SetStarted(1563474078)
	s.SetFinished(1563474079)
	s.SetHost("example.company.com")
	s.SetRuntime("docker")
	s.SetDistribution("linux")

	return s
}
