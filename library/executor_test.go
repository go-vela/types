// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_Executor_Getters(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	e := &Executor{
		ID:           &num64,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
	}
	wantID := num64
	wantHost := str
	wantRuntime := str
	wantDistribution := str

	// run test
	gotID := e.GetID()
	gotHost := e.GetHost()
	gotRuntime := e.GetRuntime()
	gotDistribution := e.GetDistribution()

	if gotID != wantID {
		t.Errorf("GetID is %v, want %v", gotID, wantID)
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

func TestLibrary_Executor_Getters_Empty(t *testing.T) {
	// setup types
	e := &Executor{}

	// run test
	gotID := e.GetID()
	gotHost := e.GetHost()
	gotRuntime := e.GetRuntime()
	gotDistribution := e.GetDistribution()

	if gotID != 0 {
		t.Errorf("GetID is %v, want 0", gotID)
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

func TestLibrary_Executor_Setters(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	e := &Executor{}

	wantID := num64
	wantHost := str
	wantRuntime := str
	wantDistribution := str

	// Run tests
	e.SetID(wantID)
	e.SetHost(wantHost)
	e.SetRuntime(wantRuntime)
	e.SetDistribution(wantDistribution)

	if e.GetID() != wantID {
		t.Errorf("SetID is %v, want %v", e.GetID(), wantID)
	}
	if e.GetHost() != wantHost {
		t.Errorf("SetHost is %v, want %v", e.GetHost(), wantHost)
	}
	if e.GetRuntime() != wantRuntime {
		t.Errorf("SetRuntime is %v, want %v", e.GetRuntime(), wantRuntime)
	}
	if e.GetDistribution() != wantDistribution {
		t.Errorf("SetDistribution is %v, want %v", e.GetDistribution(), wantDistribution)
	}
}

func TestLibrary_Executor_Setters_Empty(t *testing.T) {
	// setup types
	e := &Executor{}
	e = nil

	// Run tests
	e.SetID(0)
	e.SetHost("")
	e.SetRuntime("")
	e.SetDistribution("")

	if e.GetID() != 0 {
		t.Errorf("SetID is %v, want 0", e.GetID())
	}
	if e.GetHost() != "" {
		t.Errorf("SetHost is %v, want \"\"", e.GetHost())
	}
	if e.GetRuntime() != "" {
		t.Errorf("SetRuntime is %v, want \"\"", e.GetRuntime())
	}
	if e.GetDistribution() != "" {
		t.Errorf("SetDistribution is %v, want \"\"", e.GetDistribution())
	}
}

func TestLibrary_Executor_String(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	e := &Executor{
		ID:           &num64,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
	}
	want := fmt.Sprintf("%+v", *e)

	// run test
	got := e.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}
