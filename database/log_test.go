// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
)

func TestDatabase_Log_Nullify(t *testing.T) {
	// setup types
	l := &Log{
		ID:        sql.NullInt64{Int64: 0, Valid: true},
		BuildID:   sql.NullInt64{Int64: 0, Valid: true},
		RepoID:    sql.NullInt64{Int64: 0, Valid: true},
		ServiceID: sql.NullInt64{Int64: 0, Valid: true},
		StepID:    sql.NullInt64{Int64: 0, Valid: true},
		Data:      []byte{},
	}
	want := &Log{
		ID:        sql.NullInt64{Int64: 0, Valid: false},
		BuildID:   sql.NullInt64{Int64: 0, Valid: false},
		RepoID:    sql.NullInt64{Int64: 0, Valid: false},
		ServiceID: sql.NullInt64{Int64: 0, Valid: false},
		StepID:    sql.NullInt64{Int64: 0, Valid: false},
		Data:      []byte{},
	}

	// run test
	got := l.Nullify()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Nullify is %v, want %v", got, want)
	}
}

func TestDatabase_Log_Nullify_Empty(t *testing.T) {
	// setup types
	var l *Log

	// run test
	got := l.Nullify()

	if got != nil {
		t.Errorf("Nullify is %v, want nil", got)
	}
}

func TestDatabase_Log_ToLibrary(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	bytes := []byte("foo")
	want := &library.Log{
		ID:        &num64,
		BuildID:   &num64,
		RepoID:    &num64,
		ServiceID: &num64,
		StepID:    &num64,
		Data:      &bytes,
	}
	l := &Log{
		ID:        sql.NullInt64{Int64: num64, Valid: true},
		BuildID:   sql.NullInt64{Int64: num64, Valid: true},
		RepoID:    sql.NullInt64{Int64: num64, Valid: true},
		ServiceID: sql.NullInt64{Int64: num64, Valid: true},
		StepID:    sql.NullInt64{Int64: num64, Valid: true},
		Data:      bytes,
	}

	// run test
	got := l.ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Log_Validate(t *testing.T) {
	// setup types
	l := &Log{
		ID:        sql.NullInt64{Int64: 1, Valid: true},
		BuildID:   sql.NullInt64{Int64: 1, Valid: true},
		RepoID:    sql.NullInt64{Int64: 1, Valid: true},
		ServiceID: sql.NullInt64{Int64: 1, Valid: true},
		StepID:    sql.NullInt64{Int64: 1, Valid: true},
	}

	// run test
	err := l.Validate()

	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestDatabase_Log_Validate_NoStepID(t *testing.T) {
	// setup types
	l := &Log{
		ID:      sql.NullInt64{Int64: 1, Valid: true},
		BuildID: sql.NullInt64{Int64: 1, Valid: true},
		RepoID:  sql.NullInt64{Int64: 1, Valid: true},
	}

	// run test
	err := l.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Log_Validate_NoBuildID(t *testing.T) {
	// setup types
	l := &Log{
		ID:     sql.NullInt64{Int64: 1, Valid: true},
		RepoID: sql.NullInt64{Int64: 1, Valid: true},
		StepID: sql.NullInt64{Int64: 1, Valid: true},
	}

	// run test
	err := l.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Log_Validate_NoRepoID(t *testing.T) {
	// setup types
	l := &Log{
		ID:      sql.NullInt64{Int64: 1, Valid: true},
		BuildID: sql.NullInt64{Int64: 1, Valid: true},
		StepID:  sql.NullInt64{Int64: 1, Valid: true},
	}

	// run test
	err := l.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_LogFromLibrary(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	bytes := []byte("foo")
	want := &Log{
		ID:        sql.NullInt64{Int64: num64, Valid: true},
		BuildID:   sql.NullInt64{Int64: num64, Valid: true},
		RepoID:    sql.NullInt64{Int64: num64, Valid: true},
		ServiceID: sql.NullInt64{Int64: num64, Valid: true},
		StepID:    sql.NullInt64{Int64: num64, Valid: true},
		Data:      bytes,
	}

	l := &library.Log{
		ID:        &num64,
		BuildID:   &num64,
		RepoID:    &num64,
		ServiceID: &num64,
		StepID:    &num64,
		Data:      &bytes,
	}

	// run test
	got := LogFromLibrary(l)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LogFromLibrary is %v, want %v", got, want)
	}
}
