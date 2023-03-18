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
			initStep: testStepInitStep(),
			want:     testStepInitStep(),
		},
		{
			initStep: testServiceInitStep(),
			want:     testServiceInitStep(),
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

		if test.initStep.GetStepID() != test.want.GetStepID() {
			t.Errorf("GetStepID is %v, want %v", test.initStep.GetStepID(), test.want.GetStepID())
		}

		if test.initStep.GetServiceID() != test.want.GetServiceID() {
			t.Errorf("GetServiceID is %v, want %v", test.initStep.GetServiceID(), test.want.GetServiceID())
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
			initStep: testStepInitStep(),
			want:     testStepInitStep(),
		},
		{
			initStep: testServiceInitStep(),
			want:     testServiceInitStep(),
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
		test.initStep.SetStepID(test.want.GetStepID())
		test.initStep.SetServiceID(test.want.GetServiceID())
		test.initStep.SetNumber(test.want.GetNumber())
		test.initStep.SetReporter(test.want.GetReporter())
		test.initStep.SetName(test.want.GetName())

		if test.initStep.GetID() != test.want.GetID() {
			t.Errorf("SetID is %v, want %v", test.initStep.GetID(), test.want.GetID())
		}

		if test.initStep.GetRepoID() != test.want.GetRepoID() {
			t.Errorf("SetRepoID is %v, want %v", test.initStep.GetRepoID(), test.want.GetRepoID())
		}

		if test.initStep.GetBuildID() != test.want.GetBuildID() {
			t.Errorf("SetBuildID is %v, want %v", test.initStep.GetBuildID(), test.want.GetBuildID())
		}

		if test.initStep.GetStepID() != test.want.GetStepID() {
			t.Errorf("SetStepID is %v, want %v", test.initStep.GetStepID(), test.want.GetStepID())
		}

		if test.initStep.GetServiceID() != test.want.GetServiceID() {
			t.Errorf("SetServiceID is %v, want %v", test.initStep.GetServiceID(), test.want.GetServiceID())
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
	}
}

func TestLibrary_InitStep_String(t *testing.T) {
	// setup types
	i := testInitStep()

	want := fmt.Sprintf(`{
  ID: %d,
  RepoID: %d,
  BuildID: %d,
  StepID: %d,
  ServiceID: %d,
  Number: %d,
  Reporter: %s,
  Name: %s,
}`,
		i.GetID(),
		i.GetRepoID(),
		i.GetBuildID(),
		i.GetStepID(),
		i.GetServiceID(),
		i.GetNumber(),
		i.GetReporter(),
		i.GetName(),
	)

	// run test
	got := i.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

func TestLibrary_InitStepFromPipelineInitStep(t *testing.T) {
	// setup types
	i := testInitStep()
	i.SetReporter("Foobar Runtime")

	// modify fields that aren't set
	i.ID = nil
	i.RepoID = nil
	i.BuildID = nil
	i.StepID = nil
	i.ServiceID = nil

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
			},
			want: i,
		},
	}

	// run tests
	for _, test := range tests {
		got := InitStepFromPipelineInitStep(test.initStep)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("InitStepFromPipelineInitStep for %s is %v, want %v", test.name, got, test.want)
		}
	}
}

func TestLibrary_InitStepLogFromBuild(t *testing.T) {
	// setup types
	i := testInitStep()

	// modify fields that aren't set
	i.ID = nil
	i.StepID = nil
	i.ServiceID = nil
	i.Number = nil
	i.Reporter = nil
	i.Name = nil

	tests := []struct {
		name    string
		build   *Build
		want    *InitStep
		wantLog *Log
	}{
		{
			name:    "nil Build",
			build:   nil,
			want:    &InitStep{},
			wantLog: &Log{},
		},
		{
			name:    "empty Build",
			build:   new(Build),
			want:    &InitStep{},
			wantLog: &Log{},
		},
		{
			name: "populated Build",
			build: &Build{
				ID:     i.BuildID,
				RepoID: i.RepoID,
			},
			want: i,
			wantLog: &Log{
				RepoID:  i.RepoID,
				BuildID: i.BuildID,
			},
		},
	}

	// run tests
	for _, test := range tests {
		got, gotLog := InitStepLogFromBuild(test.build)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("InitStepLogFromBuild for %s InitStep is %v, want %v", test.name, got, test.want)
		}

		if !reflect.DeepEqual(gotLog, test.wantLog) {
			t.Errorf("InitStepLogFromBuild for %s Log is %v, want %v", test.name, gotLog, test.wantLog)
		}
	}
}

func TestLibrary_InitStepLogFromStep(t *testing.T) {
	// setup types
	i := testStepInitStep()

	// modify fields that aren't set
	i.ID = nil
	i.ServiceID = nil
	i.Number = nil
	i.Reporter = nil
	i.Name = nil

	tests := []struct {
		name    string
		step    *Step
		want    *InitStep
		wantLog *Log
	}{
		{
			name:    "nil Step",
			step:    nil,
			want:    &InitStep{},
			wantLog: &Log{},
		},
		{
			name:    "empty Step",
			step:    new(Step),
			want:    &InitStep{},
			wantLog: &Log{},
		},
		{
			name: "populated Step",
			step: &Step{
				ID:      i.StepID,
				RepoID:  i.RepoID,
				BuildID: i.BuildID,
			},
			want: i,
			wantLog: &Log{
				RepoID:  i.RepoID,
				BuildID: i.BuildID,
			},
		},
	}

	// run tests
	for _, test := range tests {
		got, gotLog := InitStepLogFromStep(test.step)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("InitStepLogFromBuild for %s InitStep is %v, want %v", test.name, got, test.want)
		}

		if !reflect.DeepEqual(gotLog, test.wantLog) {
			t.Errorf("InitStepLogFromBuild for %s Log is %v, want %v", test.name, gotLog, test.wantLog)
		}
	}
}

func TestLibrary_InitStepLogFromService(t *testing.T) {
	// setup types
	i := testServiceInitStep()

	// modify fields that aren't set
	i.ID = nil
	i.StepID = nil
	i.Number = nil
	i.Reporter = nil
	i.Name = nil

	tests := []struct {
		name    string
		service *Service
		want    *InitStep
		wantLog *Log
	}{
		{
			name:    "nil Service",
			service: nil,
			want:    &InitStep{},
			wantLog: &Log{},
		},
		{
			name:    "empty Service",
			service: new(Service),
			want:    &InitStep{},
			wantLog: &Log{},
		},
		{
			name: "populated Service",
			service: &Service{
				ID:      i.ServiceID,
				RepoID:  i.RepoID,
				BuildID: i.BuildID,
			},
			want: i,
			wantLog: &Log{
				RepoID:  i.RepoID,
				BuildID: i.BuildID,
			},
		},
	}

	// run tests
	for _, test := range tests {
		got, gotLog := InitStepLogFromService(test.service)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("InitStepLogFromBuild for %s InitStep is %v, want %v", test.name, got, test.want)
		}

		if !reflect.DeepEqual(gotLog, test.wantLog) {
			t.Errorf("InitStepLogFromBuild for %s Log is %v, want %v", test.name, gotLog, test.wantLog)
		}
	}
}

// testInitStep is a test helper function to create a InitStep
// type for a Build with all fields set to a fake value.
func testInitStep() *InitStep {
	i := new(InitStep)

	i.SetID(1)
	i.SetRepoID(1)
	i.SetBuildID(1)
	i.SetNumber(1)
	i.SetReporter("Kubernetes Runtime")
	i.SetName("clone")

	return i
}

// testStepInitStep is a test helper function to create a InitStep
// type for a Step with all fields set to a fake value.
func testStepInitStep() *InitStep {
	i := testInitStep()
	i.SetStepID(1)

	return i
}

// testServiceInitStep is a test helper function to create a InitStep
// type for a Service with all fields set to a fake value.
func testServiceInitStep() *InitStep {
	i := testInitStep()
	i.SetServiceID(1)

	return i
}
