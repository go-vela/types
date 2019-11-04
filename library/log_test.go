// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
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
		ID:      &num64,
		StepID:  &num64,
		BuildID: &num64,
		RepoID:  &num64,
		Data:    &bytes,
	}
	wantID := num64
	wantStepID := num64
	wantBuildID := num64
	wantRepoID := num64
	wantData := bytes

	// run test
	gotID := l.GetID()
	gotStepID := l.GetStepID()
	gotBuildID := l.GetBuildID()
	gotRepoID := l.GetRepoID()
	gotData := l.GetData()

	if gotID != wantID {
		t.Errorf("GetID is %v, want %v", gotID, wantID)
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
	l := &Log{}

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
	l := &Log{}

	wantID := num64
	wantStepID := num64
	wantBuildID := num64
	wantRepoID := num64
	wantData := bytes

	// run test
	l.SetID(wantID)
	l.SetStepID(wantStepID)
	l.SetBuildID(wantBuildID)
	l.SetRepoID(wantRepoID)
	l.SetData(wantData)

	if *l.ID != wantID {
		t.Errorf("GetID is %v, want %v", *l.ID, wantID)
	}
	if *l.StepID != wantStepID {
		t.Errorf("GetStepID is %v, want %v", *l.StepID, wantStepID)
	}
	if *l.BuildID != wantBuildID {
		t.Errorf("GetBuildID is %v, want %v", *l.BuildID, wantBuildID)
	}
	if *l.RepoID != wantRepoID {
		t.Errorf("GetRepoID is %v, want %v", *l.RepoID, wantRepoID)
	}
	if !reflect.DeepEqual(*l.Data, wantData) {
		t.Errorf("GetData is %v, want %v", *l.Data, wantData)
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
