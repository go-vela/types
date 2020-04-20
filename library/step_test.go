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
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	s := &Step{
		ID:           &num64,
		BuildID:      &num64,
		RepoID:       &num64,
		Number:       &num,
		Name:         &str,
		Image:        &str,
		Stage:        &str,
		Status:       &str,
		Error:        &str,
		ExitCode:     &num,
		Created:      &num64,
		Started:      &num64,
		Finished:     &num64,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
	}
	wantID := num64
	wantBuildID := num64
	wantRepoID := num64
	wantNumber := num
	wantName := str
	wantImage := str
	wantStage := str
	wantStatus := str
	wantError := str
	wantExitCode := num
	wantCreated := num64
	wantStarted := num64
	wantFinished := num64
	wantHost := str
	wantRuntime := str
	wantDistribution := str

	// run test
	gotID := s.GetID()
	gotBuildID := s.GetBuildID()
	gotRepoID := s.GetRepoID()
	gotNumber := s.GetNumber()
	gotName := s.GetName()
	gotImage := s.GetImage()
	gotStage := s.GetStage()
	gotStatus := s.GetStatus()
	gotError := s.GetError()
	gotExitCode := s.GetExitCode()
	gotCreated := s.GetCreated()
	gotStarted := s.GetStarted()
	gotFinished := s.GetFinished()
	gotHost := s.GetHost()
	gotRuntime := s.GetRuntime()
	gotDistribution := s.GetDistribution()

	if gotID != wantID {
		t.Errorf("GetID is %v, want %v", gotID, wantID)
	}
	if gotBuildID != wantBuildID {
		t.Errorf("GetBuildID is %v, want %v", gotBuildID, wantBuildID)
	}
	if gotRepoID != wantRepoID {
		t.Errorf("GetRepoID is %v, want %v", gotRepoID, wantRepoID)
	}
	if gotNumber != wantNumber {
		t.Errorf("GetNumber is %v, want %v", gotNumber, wantNumber)
	}
	if gotName != wantName {
		t.Errorf("GetName is %v, want %v", gotName, wantName)
	}
	if gotImage != wantImage {
		t.Errorf("GetImage is %v, want %v", gotImage, wantImage)
	}
	if gotStage != wantStage {
		t.Errorf("GetStage is %v, want %v", gotStage, wantStage)
	}
	if gotStatus != wantStatus {
		t.Errorf("GetStatus is %v, want %v", gotStatus, wantStatus)
	}
	if gotError != wantError {
		t.Errorf("GetError is %v, want %v", gotError, wantError)
	}
	if gotExitCode != wantExitCode {
		t.Errorf("GetExitCode is %v, want %v", gotExitCode, wantExitCode)
	}
	if gotCreated != wantCreated {
		t.Errorf("GetCreated is %v, want %v", gotCreated, wantCreated)
	}
	if gotStarted != wantStarted {
		t.Errorf("GetStarted is %v, want %v", gotStarted, wantStarted)
	}
	if gotFinished != wantFinished {
		t.Errorf("GetFinished is %v, want %v", gotFinished, wantFinished)
	}
	if gotHost != wantHost {
		t.Errorf("GetHost is %v, want %v", gotHost, wantHost)
	}
	if gotRuntime != wantRuntime {
		t.Errorf("GetRuntime is %v, want %v", gotRuntime, wantRuntime)
	}
	if gotDistribution != wantDistribution {
		t.Errorf("GetDistribution is %v, want %v", gotDistribution, wantDistribution)
	}
}

func TestStep_Getters_Empty(t *testing.T) {
	// setup types
	s := new(Step)

	// run test
	gotID := s.GetID()
	gotBuildID := s.GetBuildID()
	gotRepoID := s.GetRepoID()
	gotNumber := s.GetNumber()
	gotName := s.GetName()
	gotImage := s.GetImage()
	gotStage := s.GetStage()
	gotStatus := s.GetStatus()
	gotError := s.GetError()
	gotExitCode := s.GetExitCode()
	gotCreated := s.GetCreated()
	gotStarted := s.GetStarted()
	gotFinished := s.GetFinished()
	gotHost := s.GetHost()
	gotRuntime := s.GetRuntime()
	gotDistribution := s.GetDistribution()

	if gotID != 0 {
		t.Errorf("GetID is %v, want 0", gotID)
	}
	if gotBuildID != 0 {
		t.Errorf("GetBuildID is %v, want 0", gotBuildID)
	}
	if gotRepoID != 0 {
		t.Errorf("GetRepoID is %v, want 0", gotRepoID)
	}
	if gotNumber != 0 {
		t.Errorf("GetNumber is %v, want 0", gotNumber)
	}
	if gotName != "" {
		t.Errorf("GetName is %v, want \"\"", gotName)
	}
	if gotImage != "" {
		t.Errorf("GetImage is %v, want \"\"", gotImage)
	}
	if gotStage != "" {
		t.Errorf("GetStage is %v, want \"\"", gotStage)
	}
	if gotStatus != "" {
		t.Errorf("GetStatus is %v, want \"\"", gotStatus)
	}
	if gotError != "" {
		t.Errorf("GetError is %v, want \"\"", gotError)
	}
	if gotExitCode != 0 {
		t.Errorf("GetExitCode is %v, want 0", gotExitCode)
	}
	if gotCreated != 0 {
		t.Errorf("GetCreated is %v, want 0", gotCreated)
	}
	if gotStarted != 0 {
		t.Errorf("GetStarted is %v, want 0", gotStarted)
	}
	if gotFinished != 0 {
		t.Errorf("GetFinished is %v, want 0", gotFinished)
	}
	if gotHost != "" {
		t.Errorf("GetHost is %v, want \"\"", gotHost)
	}
	if gotRuntime != "" {
		t.Errorf("GetRuntime is %v, want \"\"", gotRuntime)
	}
	if gotDistribution != "" {
		t.Errorf("GetDistribution is %v, want \"\"", gotDistribution)
	}
}

func TestLibrary_Step_Setters(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	s := new(Step)

	wantID := num64
	wantBuildID := num64
	wantRepoID := num64
	wantNumber := num
	wantName := str
	wantImage := str
	wantStage := str
	wantStatus := str
	wantError := str
	wantExitCode := num
	wantCreated := num64
	wantStarted := num64
	wantFinished := num64
	wantHost := str
	wantRuntime := str
	wantDistribution := str

	// run test
	s.SetID(wantID)
	s.SetBuildID(wantBuildID)
	s.SetRepoID(wantRepoID)
	s.SetNumber(wantNumber)
	s.SetName(wantName)
	s.SetImage(wantImage)
	s.SetStage(wantStage)
	s.SetStatus(wantStatus)
	s.SetError(wantError)
	s.SetExitCode(wantExitCode)
	s.SetCreated(wantCreated)
	s.SetStarted(wantStarted)
	s.SetFinished(wantFinished)
	s.SetHost(wantHost)
	s.SetRuntime(wantRuntime)
	s.SetDistribution(wantDistribution)

	if s.GetID() != wantID {
		t.Errorf("SetID is %v, want %v", s.GetID(), wantID)
	}
	if s.GetBuildID() != wantBuildID {
		t.Errorf("SetBuildID is %v, want %v", s.GetBuildID(), wantBuildID)
	}
	if s.GetRepoID() != wantRepoID {
		t.Errorf("SetRepoID is %v, want %v", s.GetRepoID(), wantRepoID)
	}
	if s.GetNumber() != wantNumber {
		t.Errorf("SetNumber is %v, want %v", s.GetNumber(), wantNumber)
	}
	if s.GetName() != wantName {
		t.Errorf("SetName is %v, want %v", s.GetName(), wantName)
	}
	if s.GetImage() != wantImage {
		t.Errorf("SetImage is %v, want %v", s.GetImage(), wantImage)
	}
	if s.GetStage() != wantStage {
		t.Errorf("SetStage is %v, want %v", s.GetStage(), wantStage)
	}
	if s.GetStatus() != wantStatus {
		t.Errorf("SetStatus is %v, want %v", s.GetStatus(), wantStatus)
	}
	if s.GetError() != wantError {
		t.Errorf("SetError is %v, want %v", s.GetError(), wantError)
	}
	if s.GetExitCode() != wantExitCode {
		t.Errorf("SetExitCode is %v, want %v", s.GetExitCode(), wantExitCode)
	}
	if s.GetCreated() != wantCreated {
		t.Errorf("SetCreated is %v, want %v", s.GetCreated(), wantCreated)
	}
	if s.GetStarted() != wantStarted {
		t.Errorf("SetStarted is %v, want %v", s.GetStarted(), wantStarted)
	}
	if s.GetFinished() != wantFinished {
		t.Errorf("SetFinished is %v, want %v", s.GetFinished(), wantFinished)
	}
	if s.GetHost() != wantHost {
		t.Errorf("SetHost is %v, want %v", s.GetHost(), wantHost)
	}
	if s.GetRuntime() != wantRuntime {
		t.Errorf("SetRuntime is %v, want %v", s.GetRuntime(), wantRuntime)
	}
	if s.GetDistribution() != wantDistribution {
		t.Errorf("SetDistribution is %v, want %v", s.GetDistribution(), wantDistribution)
	}
}

func TestLibrary_Step_Setters_Empty(t *testing.T) {
	// setup types
	var s *Step

	// run test
	s.SetID(0)
	s.SetBuildID(0)
	s.SetRepoID(0)
	s.SetNumber(0)
	s.SetName("")
	s.SetImage("")
	s.SetStage("")
	s.SetStatus("")
	s.SetError("")
	s.SetExitCode(0)
	s.SetCreated(0)
	s.SetStarted(0)
	s.SetFinished(0)
	s.SetHost("")
	s.SetRuntime("")
	s.SetDistribution("")

	if s.GetID() != 0 {
		t.Errorf("SetID is %v, want 0", s.GetID())
	}
	if s.GetBuildID() != 0 {
		t.Errorf("SetBuildID is %v, want 0", s.GetBuildID())
	}
	if s.GetRepoID() != 0 {
		t.Errorf("SetRepoID is %v, want 0", s.GetRepoID())
	}
	if s.GetNumber() != 0 {
		t.Errorf("SetNumber is %v, want 0", s.GetNumber())
	}
	if s.GetName() != "" {
		t.Errorf("SetName is %v, want \"\"", s.GetName())
	}
	if s.GetImage() != "" {
		t.Errorf("SetImage is %v, want \"\"", s.GetImage())
	}
	if s.GetStage() != "" {
		t.Errorf("SetStage is %v, want \"\"", s.GetStage())
	}
	if s.GetStatus() != "" {
		t.Errorf("SetStatus is %v, want \"\"", s.GetStatus())
	}
	if s.GetError() != "" {
		t.Errorf("SetError is %v, want \"\"", s.GetError())
	}
	if s.GetExitCode() != 0 {
		t.Errorf("SetExitCode is %v, want 0", s.GetExitCode())
	}
	if s.GetCreated() != 0 {
		t.Errorf("SetCreated is %v, want 0", s.GetCreated())
	}
	if s.GetStarted() != 0 {
		t.Errorf("SetStarted is %v, want 0", s.GetStarted())
	}
	if s.GetFinished() != 0 {
		t.Errorf("SetFinished is %v, want 0", s.GetFinished())
	}
	if s.GetHost() != "" {
		t.Errorf("SetHost is %v, want \"\"", s.GetHost())
	}
	if s.GetRuntime() != "" {
		t.Errorf("SetRuntime is %v, want \"\"", s.GetRuntime())
	}
	if s.GetDistribution() != "" {
		t.Errorf("SetDistribution is %v, want \"\"", s.GetDistribution())
	}
}

func TestStep_String(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	s := &Step{
		ID:           &num64,
		BuildID:      &num64,
		RepoID:       &num64,
		Number:       &num,
		Name:         &str,
		Image:        &str,
		Stage:        &str,
		Status:       &str,
		Error:        &str,
		ExitCode:     &num,
		Created:      &num64,
		Started:      &num64,
		Finished:     &num64,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
	}
	want := fmt.Sprintf("%+v", *s)

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
