// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
)

func TestDatabase_InitStep_Nullify(t *testing.T) {
	// setup types
	var s *InitStep

	want := &InitStep{
		ID:       sql.NullInt64{Int64: 0, Valid: false},
		RepoID:   sql.NullInt64{Int64: 0, Valid: false},
		BuildID:  sql.NullInt64{Int64: 0, Valid: false},
		Number:   sql.NullInt32{Int32: 0, Valid: false},
		Reporter: sql.NullString{String: "", Valid: false},
		Name:     sql.NullString{String: "", Valid: false},
		Mimetype: sql.NullString{String: "", Valid: false},
	}

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
			want:     nil,
		},
		{
			initStep: new(InitStep),
			want:     want,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.initStep.Nullify()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Nullify is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_InitStep_ToLibrary(t *testing.T) {
	// setup types
	want := new(library.InitStep)

	want.SetID(1)
	want.SetRepoID(1)
	want.SetBuildID(1)
	want.SetNumber(1)
	want.SetReporter("Foobar Runtime")
	want.SetName("foobar")
	want.SetMimetype("text/plain")

	// run test
	got := testInitStep().ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_InitStep_Validate(t *testing.T) {
	// setup types
	tests := []struct {
		failure  bool
		initStep *InitStep
	}{
		{
			failure:  false,
			initStep: testInitStep(),
		},
		{ // no repo_id set for InitStep
			failure: true,
			initStep: &InitStep{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				BuildID:  sql.NullInt64{Int64: 1, Valid: true},
				Number:   sql.NullInt32{Int32: 1, Valid: true},
				Reporter: sql.NullString{String: "Foobar Runtime", Valid: true},
				Name:     sql.NullString{String: "foobar", Valid: true},
			},
		},
		{ // no build_id set for InitStep
			failure: true,
			initStep: &InitStep{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				RepoID:   sql.NullInt64{Int64: 1, Valid: true},
				Number:   sql.NullInt32{Int32: 1, Valid: true},
				Reporter: sql.NullString{String: "Foobar Runtime", Valid: true},
				Name:     sql.NullString{String: "foobar", Valid: true},
			},
		},
		{ // no number set for InitStep
			failure: true,
			initStep: &InitStep{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				RepoID:   sql.NullInt64{Int64: 1, Valid: true},
				BuildID:  sql.NullInt64{Int64: 1, Valid: true},
				Reporter: sql.NullString{String: "Foobar Runtime", Valid: true},
				Name:     sql.NullString{String: "foobar", Valid: true},
			},
		},
		{ // no reporter set for InitStep
			failure: true,
			initStep: &InitStep{
				ID:      sql.NullInt64{Int64: 1, Valid: true},
				RepoID:  sql.NullInt64{Int64: 1, Valid: true},
				BuildID: sql.NullInt64{Int64: 1, Valid: true},
				Number:  sql.NullInt32{Int32: 1, Valid: true},
				Name:    sql.NullString{String: "foobar", Valid: true},
			},
		},
		{ // no name set for InitStep
			failure: true,
			initStep: &InitStep{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				RepoID:   sql.NullInt64{Int64: 1, Valid: true},
				BuildID:  sql.NullInt64{Int64: 1, Valid: true},
				Number:   sql.NullInt32{Int32: 1, Valid: true},
				Reporter: sql.NullString{String: "Foobar Runtime", Valid: true},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.initStep.Validate()

		if test.failure {
			if err == nil {
				t.Errorf("Validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}

func TestDatabase_InitStepFromLibrary(t *testing.T) {
	// setup types
	s := new(library.InitStep)

	s.SetID(1)
	s.SetRepoID(1)
	s.SetBuildID(1)
	s.SetNumber(1)
	s.SetReporter("Foobar Runtime")
	s.SetName("foobar")
	s.SetMimetype("text/plain")

	want := testInitStep()

	// run test
	got := InitStepFromLibrary(s)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("InitStepFromLibrary is %v, want %v", got, want)
	}
}

// testInitStep is a test helper function to create a InitStep
// type with all fields set to a fake value.
func testInitStep() *InitStep {
	return &InitStep{
		ID:       sql.NullInt64{Int64: 1, Valid: true},
		RepoID:   sql.NullInt64{Int64: 1, Valid: true},
		BuildID:  sql.NullInt64{Int64: 1, Valid: true},
		Number:   sql.NullInt32{Int32: 1, Valid: true},
		Reporter: sql.NullString{String: "Foobar Runtime", Valid: true},
		Name:     sql.NullString{String: "foobar", Valid: true},
		Mimetype: sql.NullString{String: "text/plain", Valid: true},
	}
}
