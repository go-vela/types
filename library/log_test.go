// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_Log_Getters(t *testing.T) {
	// setup types
	num64 := int64(1)
	bytes := []byte("foo")
	l := &Log{
		ID:        &num64,
		ServiceID: &num64,
		StepID:    &num64,
		BuildID:   &num64,
		RepoID:    &num64,
		Data:      &bytes,
	}
	wantID := num64
	wantServiceID := num64
	wantStepID := num64
	wantBuildID := num64
	wantRepoID := num64
	wantData := bytes

	// run test
	gotID := l.GetID()
	gotServiceID := l.GetServiceID()
	gotStepID := l.GetStepID()
	gotBuildID := l.GetBuildID()
	gotRepoID := l.GetRepoID()
	gotData := l.GetData()

	if gotID != wantID {
		t.Errorf("GetID is %v, want %v", gotID, wantID)
	}
	if gotServiceID != wantServiceID {
		t.Errorf("GetServiceID is %v, want %v", gotServiceID, wantServiceID)
	}
	if gotStepID != wantStepID {
		t.Errorf("GetStepID is %v, want %v", gotStepID, wantStepID)
	}
	if gotBuildID != wantBuildID {
		t.Errorf("GetBuildID is %v, want %v", gotBuildID, wantBuildID)
	}
	if gotRepoID != wantRepoID {
		t.Errorf("GetRepoID is %v, want %v", gotRepoID, wantRepoID)
	}
	if !reflect.DeepEqual(gotData, wantData) {
		t.Errorf("GetData is %v, want %v", gotData, wantData)
	}
}

func TestLibrary_Log_Getters_Empty(t *testing.T) {
	// setup types
	l := new(Log)

	// run test
	gotID := l.GetID()
	gotStepID := l.GetStepID()
	gotBuildID := l.GetBuildID()
	gotRepoID := l.GetRepoID()
	gotData := l.GetData()

	if gotID != 0 {
		t.Errorf("GetID is %v, want 0", gotID)
	}
	if gotStepID != 0 {
		t.Errorf("GetStepID is %v, want 0", gotStepID)
	}
	if gotBuildID != 0 {
		t.Errorf("GetBuildID is %v, want 0", gotBuildID)
	}
	if gotRepoID != 0 {
		t.Errorf("GetRepoID is %v, want 0", gotRepoID)
	}
	if !reflect.DeepEqual(gotData, []byte{}) {
		t.Errorf("GetData is %v, want []byte{}", gotData)
	}
}

func TestLibrary_Log_Setters(t *testing.T) {
	// setup types
	num64 := int64(1)
	bytes := []byte("foo")
	l := new(Log)

	wantID := num64
	wantServiceID := num64
	wantStepID := num64
	wantBuildID := num64
	wantRepoID := num64
	wantData := bytes

	// run test
	l.SetID(wantID)
	l.SetServiceID(wantServiceID)
	l.SetStepID(wantStepID)
	l.SetBuildID(wantBuildID)
	l.SetRepoID(wantRepoID)
	l.SetData(wantData)

	if l.GetID() != wantID {
		t.Errorf("SetID is %v, want %v", l.GetID(), wantID)
	}
	if l.GetServiceID() != wantServiceID {
		t.Errorf("SetServiceID is %v, want %v", l.GetServiceID(), wantServiceID)
	}
	if l.GetStepID() != wantStepID {
		t.Errorf("SetStepID is %v, want %v", l.GetStepID(), wantStepID)
	}
	if l.GetBuildID() != wantBuildID {
		t.Errorf("SetBuildID is %v, want %v", l.GetBuildID(), wantBuildID)
	}
	if l.GetRepoID() != wantRepoID {
		t.Errorf("SetRepoID is %v, want %v", l.GetRepoID(), wantRepoID)
	}
	if !reflect.DeepEqual(l.GetData(), wantData) {
		t.Errorf("SetData is %v, want %v", l.GetData(), wantData)
	}
}

func TestLibrary_Log_Setters_Empty(t *testing.T) {
	// setup types
	var l *Log

	// run test
	l.SetID(0)
	l.SetServiceID(0)
	l.SetStepID(0)
	l.SetBuildID(0)
	l.SetRepoID(0)
	l.SetData([]byte{})

	if l.GetID() != 0 {
		t.Errorf("SetID is %v, want 0", l.GetID())
	}
	if l.GetServiceID() != 0 {
		t.Errorf("SetServiceID is %v, want 0", l.GetServiceID())
	}
	if l.GetStepID() != 0 {
		t.Errorf("SetStepID is %v, want 0", l.GetStepID())
	}
	if l.GetBuildID() != 0 {
		t.Errorf("SetBuildID is %v, want 0", l.GetBuildID())
	}
	if l.GetRepoID() != 0 {
		t.Errorf("SetRepoID is %v, want 0", l.GetRepoID())
	}
	if !reflect.DeepEqual(l.GetData(), []byte{}) {
		t.Errorf("SetData is %v, want []byte{}", l.GetData())
	}
}

func TestLibrary_Log_String(t *testing.T) {
	// setup types
	num64 := int64(1)
	bytes := []byte("foo")
	l := &Log{
		ID:      &num64,
		StepID:  &num64,
		BuildID: &num64,
		RepoID:  &num64,
		Data:    &bytes,
	}
	want := fmt.Sprintf("%+v", *l)

	// run test
	got := l.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}
