// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
)

func TestDatabase_Log_Compress(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		log     *Log
	}{
		{
			failure: false,
			log:     testLog(),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.log.Compress()

		if test.failure {
			if err == nil {
				t.Errorf("Compress should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Compress returned err: %v", err)
		}
	}
}

func TestDatabase_Log_Decompress(t *testing.T) {
	// setup types
	l := testLog()
	err := l.Compress()
	if err != nil {
		t.Errorf("unable to compress log: %v", err)
	}

	// setup tests
	tests := []struct {
		failure bool
		log     *Log
	}{
		{
			failure: false,
			log:     l,
		},
		{
			failure: true,
			log:     testLog(),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.log.Decompress()

		if test.failure {
			if err == nil {
				t.Errorf("Decompress should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Decompress returned err: %v", err)
		}
	}
}

func TestDatabase_Log_Nullify(t *testing.T) {
	// setup types
	var l *Log

	want := &Log{
		ID:        sql.NullInt64{Int64: 0, Valid: false},
		BuildID:   sql.NullInt64{Int64: 0, Valid: false},
		RepoID:    sql.NullInt64{Int64: 0, Valid: false},
		ServiceID: sql.NullInt64{Int64: 0, Valid: false},
		StepID:    sql.NullInt64{Int64: 0, Valid: false},
	}

	// setup tests
	tests := []struct {
		log  *Log
		want *Log
	}{
		{
			log:  testLog(),
			want: testLog(),
		},
		{
			log:  l,
			want: nil,
		},
		{
			log:  new(Log),
			want: want,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.log.Nullify()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Nullify is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_Log_ToLibrary(t *testing.T) {
	// setup types
	want := new(library.Log)

	want.SetID(1)
	want.SetServiceID(1)
	want.SetStepID(1)
	want.SetBuildID(1)
	want.SetRepoID(1)
	want.SetData([]byte("foo"))

	// run test
	got := testLog().ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Log_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		log     *Log
	}{
		{
			failure: false,
			log:     testLog(),
		},
		{ // no service_id or step_id set for log
			failure: true,
			log: &Log{
				ID:      sql.NullInt64{Int64: 1, Valid: true},
				BuildID: sql.NullInt64{Int64: 1, Valid: true},
				RepoID:  sql.NullInt64{Int64: 1, Valid: true},
			},
		},
		{ // no build_id set for log
			failure: true,
			log: &Log{
				ID:        sql.NullInt64{Int64: 1, Valid: true},
				RepoID:    sql.NullInt64{Int64: 1, Valid: true},
				ServiceID: sql.NullInt64{Int64: 1, Valid: true},
				StepID:    sql.NullInt64{Int64: 1, Valid: true},
			},
		},
		{ // no repo_id set for log
			failure: true,
			log: &Log{
				ID:        sql.NullInt64{Int64: 1, Valid: true},
				BuildID:   sql.NullInt64{Int64: 1, Valid: true},
				ServiceID: sql.NullInt64{Int64: 1, Valid: true},
				StepID:    sql.NullInt64{Int64: 1, Valid: true},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.log.Validate()

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

func TestDatabase_LogFromLibrary(t *testing.T) {
	// setup types
	l := new(library.Log)

	l.SetID(1)
	l.SetServiceID(1)
	l.SetStepID(1)
	l.SetBuildID(1)
	l.SetRepoID(1)
	l.SetData([]byte("foo"))

	want := testLog()

	// run test
	got := LogFromLibrary(l)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LogFromLibrary is %v, want %v", got, want)
	}
}

// testLog is a test helper function to create a Log
// type with all fields set to a fake value.
func testLog() *Log {
	return &Log{
		ID:        sql.NullInt64{Int64: 1, Valid: true},
		BuildID:   sql.NullInt64{Int64: 1, Valid: true},
		RepoID:    sql.NullInt64{Int64: 1, Valid: true},
		ServiceID: sql.NullInt64{Int64: 1, Valid: true},
		StepID:    sql.NullInt64{Int64: 1, Valid: true},
		Data:      []byte("foo"),
	}
}
