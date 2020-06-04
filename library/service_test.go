// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_Service_Environment(t *testing.T) {
	// setup types
	want := map[string]string{
		"VELA_SERVICE_CREATED":      "1563474076",
		"VELA_SERVICE_DISTRIBUTION": "linux",
		"VELA_SERVICE_EXIT_CODE":    "0",
		"VELA_SERVICE_FINISHED":     "1563474079",
		"VELA_SERVICE_HOST":         "example.company.com",
		"VELA_SERVICE_IMAGE":        "postgres:12-alpine",
		"VELA_SERVICE_NAME":         "postgres",
		"VELA_SERVICE_NUMBER":       "1",
		"VELA_SERVICE_RUNTIME":      "docker",
		"VELA_SERVICE_STARTED":      "1563474078",
		"VELA_SERVICE_STATUS":       "running",
	}

	// run test
	got := testService().Environment()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Environment is %v, want %v", got, want)
	}
}

func TestService_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		service *Service
		want    *Service
	}{
		{
			service: testService(),
			want:    testService(),
		},
		{
			service: new(Service),
			want:    new(Service),
		},
	}

	// run tests
	for _, test := range tests {
		if test.service.GetID() != test.want.GetID() {
			t.Errorf("GetID is %v, want %v", test.service.GetID(), test.want.GetID())
		}

		if test.service.GetBuildID() != test.want.GetBuildID() {
			t.Errorf("GetBuildID is %v, want %v", test.service.GetBuildID(), test.want.GetBuildID())
		}

		if test.service.GetRepoID() != test.want.GetRepoID() {
			t.Errorf("GetRepoID is %v, want %v", test.service.GetRepoID(), test.want.GetRepoID())
		}

		if test.service.GetNumber() != test.want.GetNumber() {
			t.Errorf("GetNumber is %v, want %v", test.service.GetNumber(), test.want.GetNumber())
		}

		if test.service.GetName() != test.want.GetName() {
			t.Errorf("GetName is %v, want %v", test.service.GetName(), test.want.GetName())
		}

		if test.service.GetImage() != test.want.GetImage() {
			t.Errorf("GetImage is %v, want %v", test.service.GetImage(), test.want.GetImage())
		}

		if test.service.GetStatus() != test.want.GetStatus() {
			t.Errorf("GetStatus is %v, want %v", test.service.GetStatus(), test.want.GetStatus())
		}

		if test.service.GetError() != test.want.GetError() {
			t.Errorf("GetError is %v, want %v", test.service.GetError(), test.want.GetError())
		}

		if test.service.GetExitCode() != test.want.GetExitCode() {
			t.Errorf("GetExitCode is %v, want %v", test.service.GetExitCode(), test.want.GetExitCode())
		}

		if test.service.GetCreated() != test.want.GetCreated() {
			t.Errorf("GetCreated is %v, want %v", test.service.GetCreated(), test.want.GetCreated())
		}

		if test.service.GetStarted() != test.want.GetStarted() {
			t.Errorf("GetStarted is %v, want %v", test.service.GetStarted(), test.want.GetStarted())
		}

		if test.service.GetFinished() != test.want.GetFinished() {
			t.Errorf("GetFinished is %v, want %v", test.service.GetFinished(), test.want.GetFinished())
		}

		if test.service.GetHost() != test.want.GetHost() {
			t.Errorf("GetHost is %v, want %v", test.service.GetHost(), test.want.GetHost())
		}

		if test.service.GetRuntime() != test.want.GetRuntime() {
			t.Errorf("GetRuntime is %v, want %v", test.service.GetRuntime(), test.want.GetRuntime())
		}

		if test.service.GetDistribution() != test.want.GetDistribution() {
			t.Errorf("GetDistribution is %v, want %v", test.service.GetDistribution(), test.want.GetDistribution())
		}
	}
}

func TestLibrary_Service_Setters(t *testing.T) {
	// setup types
	var s *Service

	// setup tests
	tests := []struct {
		service *Service
		want    *Service
	}{
		{
			service: testService(),
			want:    testService(),
		},
		{
			service: s,
			want:    new(Service),
		},
	}

	// run tests
	for _, test := range tests {
		test.service.SetID(test.service.GetID())
		test.service.SetBuildID(test.service.GetBuildID())
		test.service.SetRepoID(test.service.GetRepoID())
		test.service.SetNumber(test.service.GetNumber())
		test.service.SetName(test.service.GetName())
		test.service.SetImage(test.service.GetImage())
		test.service.SetStatus(test.service.GetStatus())
		test.service.SetError(test.service.GetError())
		test.service.SetExitCode(test.service.GetExitCode())
		test.service.SetCreated(test.service.GetCreated())
		test.service.SetStarted(test.service.GetStarted())
		test.service.SetFinished(test.service.GetFinished())
		test.service.SetHost(test.service.GetHost())
		test.service.SetRuntime(test.service.GetRuntime())
		test.service.SetDistribution(test.service.GetDistribution())

		if test.service.GetID() != test.service.GetID() {
			t.Errorf("SetID is %v, want %v", test.service.GetID(), test.service.GetID())
		}

		if test.service.GetBuildID() != test.service.GetBuildID() {
			t.Errorf("SetBuildID is %v, want %v", test.service.GetBuildID(), test.service.GetBuildID())
		}

		if test.service.GetRepoID() != test.service.GetRepoID() {
			t.Errorf("SetRepoID is %v, want %v", test.service.GetRepoID(), test.service.GetRepoID())
		}

		if test.service.GetNumber() != test.service.GetNumber() {
			t.Errorf("SetNumber is %v, want %v", test.service.GetNumber(), test.service.GetNumber())
		}

		if test.service.GetName() != test.service.GetName() {
			t.Errorf("SetName is %v, want %v", test.service.GetName(), test.service.GetName())
		}

		if test.service.GetImage() != test.service.GetImage() {
			t.Errorf("SetImage is %v, want %v", test.service.GetImage(), test.service.GetImage())
		}

		if test.service.GetStatus() != test.service.GetStatus() {
			t.Errorf("SetStatus is %v, want %v", test.service.GetStatus(), test.service.GetStatus())
		}

		if test.service.GetError() != test.service.GetError() {
			t.Errorf("SetError is %v, want %v", test.service.GetError(), test.service.GetError())
		}

		if test.service.GetExitCode() != test.service.GetExitCode() {
			t.Errorf("SetExitCode is %v, want %v", test.service.GetExitCode(), test.service.GetExitCode())
		}

		if test.service.GetCreated() != test.service.GetCreated() {
			t.Errorf("SetCreated is %v, want %v", test.service.GetCreated(), test.service.GetCreated())
		}

		if test.service.GetStarted() != test.service.GetStarted() {
			t.Errorf("SetStarted is %v, want %v", test.service.GetStarted(), test.service.GetStarted())
		}

		if test.service.GetFinished() != test.service.GetFinished() {
			t.Errorf("SetFinished is %v, want %v", test.service.GetFinished(), test.service.GetFinished())
		}

		if test.service.GetHost() != test.service.GetHost() {
			t.Errorf("SetHost is %v, want %v", test.service.GetHost(), test.service.GetHost())
		}

		if test.service.GetRuntime() != test.service.GetRuntime() {
			t.Errorf("SetRuntime is %v, want %v", test.service.GetRuntime(), test.service.GetRuntime())
		}

		if test.service.GetDistribution() != test.service.GetDistribution() {
			t.Errorf("SetDistribution is %v, want %v", test.service.GetDistribution(), test.service.GetDistribution())
		}
	}
}

func TestService_String(t *testing.T) {
	// setup types
	s := testService()

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
		s.GetStarted(),
		s.GetStatus(),
	)

	// run test
	got := s.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

// testService is a test helper function to create a Service
// type with all fields set to a fake value.
func testService() *Service {
	s := new(Service)

	s.SetID(1)
	s.SetBuildID(1)
	s.SetRepoID(1)
	s.SetNumber(1)
	s.SetName("postgres")
	s.SetImage("postgres:12-alpine")
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
