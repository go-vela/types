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

func TestDatabase_Init_Nullify(t *testing.T) {
	// setup types
	var s *Init

	want := &Init{
		ID:       sql.NullInt64{Int64: 0, Valid: false},
		BuildID:  sql.NullInt64{Int64: 0, Valid: false},
		RepoID:   sql.NullInt64{Int64: 0, Valid: false},
		Number:   sql.NullInt32{Int32: 0, Valid: false},
		Name:     sql.NullString{String: "", Valid: false},
		Reporter: sql.NullString{String: "", Valid: false},
		Mimetype: sql.NullString{String: "", Valid: false},
	}

	// setup tests
	tests := []struct {
		init *Init
		want *Init
	}{
		{
			init: testInit(),
			want: testInit(),
		},
		{
			init: s,
			want: nil,
		},
		{
			init: new(Init),
			want: want,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.init.Nullify()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Nullify is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_Init_ToLibrary(t *testing.T) {
	// setup types
	want := new(library.Init)

	want.SetID(1)
	want.SetBuildID(1)
	want.SetRepoID(1)
	want.SetNumber(1)
	want.SetName("foobar")
	want.SetMimetype("text/plain")
	want.SetReporter("Foobar Runtime")

	// run test
	got := testInit().ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Init_Validate(t *testing.T) {
	// setup types
	tests := []struct {
		failure bool
		init    *Init
	}{
		{
			failure: false,
			init:    testInit(),
		},
		{ // no build_id set for init
			failure: true,
			init: &Init{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				RepoID:   sql.NullInt64{Int64: 1, Valid: true},
				Number:   sql.NullInt32{Int32: 1, Valid: true},
				Name:     sql.NullString{String: "foobar", Valid: true},
				Reporter: sql.NullString{String: "Foobar Runtime", Valid: true},
			},
		},
		{ // no repo_id set for init
			failure: true,
			init: &Init{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				BuildID:  sql.NullInt64{Int64: 1, Valid: true},
				Number:   sql.NullInt32{Int32: 1, Valid: true},
				Name:     sql.NullString{String: "foobar", Valid: true},
				Reporter: sql.NullString{String: "Foobar Runtime", Valid: true},
			},
		},
		{ // no name set for init
			failure: true,
			init: &Init{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				BuildID:  sql.NullInt64{Int64: 1, Valid: true},
				RepoID:   sql.NullInt64{Int64: 1, Valid: true},
				Number:   sql.NullInt32{Int32: 1, Valid: true},
				Reporter: sql.NullString{String: "Foobar Runtime", Valid: true},
			},
		},
		{ // no number set for init
			failure: true,
			init: &Init{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				BuildID:  sql.NullInt64{Int64: 1, Valid: true},
				RepoID:   sql.NullInt64{Int64: 1, Valid: true},
				Name:     sql.NullString{String: "foobar", Valid: true},
				Reporter: sql.NullString{String: "Foobar Runtime", Valid: true},
			},
		},
		{ // no reporter set for init
			failure: true,
			init: &Init{
				ID:      sql.NullInt64{Int64: 1, Valid: true},
				BuildID: sql.NullInt64{Int64: 1, Valid: true},
				RepoID:  sql.NullInt64{Int64: 1, Valid: true},
				Number:  sql.NullInt32{Int32: 1, Valid: true},
				Name:    sql.NullString{String: "foobar", Valid: true},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.init.Validate()

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

func TestDatabase_InitFromLibrary(t *testing.T) {
	// setup types
	s := new(library.Init)

	s.SetID(1)
	s.SetBuildID(1)
	s.SetRepoID(1)
	s.SetNumber(1)
	s.SetName("foobar")
	s.SetMimetype("text/plain")
	s.SetReporter("Foobar Runtime")

	want := testInit()

	// run test
	got := InitFromLibrary(s)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("InitFromLibrary is %v, want %v", got, want)
	}
}

// testInit is a test helper function to create a Init
// type with all fields set to a fake value.
func testInit() *Init {
	return &Init{
		ID:       sql.NullInt64{Int64: 1, Valid: true},
		BuildID:  sql.NullInt64{Int64: 1, Valid: true},
		RepoID:   sql.NullInt64{Int64: 1, Valid: true},
		Number:   sql.NullInt32{Int32: 1, Valid: true},
		Name:     sql.NullString{String: "foobar", Valid: true},
		Mimetype: sql.NullString{String: "text/plain", Valid: true},
		Reporter: sql.NullString{String: "Foobar Runtime", Valid: true},
	}
}
