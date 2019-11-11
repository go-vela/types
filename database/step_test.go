// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
)

func TestDatabase_Step_Nullify(t *testing.T) {
	// setup types
	s := &Step{
		ID:           sql.NullInt64{Int64: 0, Valid: true},
		BuildID:      sql.NullInt64{Int64: 0, Valid: true},
		RepoID:       sql.NullInt64{Int64: 0, Valid: true},
		Number:       sql.NullInt32{Int32: 0, Valid: true},
		Name:         sql.NullString{String: "", Valid: true},
		Stage:        sql.NullString{String: "", Valid: true},
		Status:       sql.NullString{String: "", Valid: true},
		Error:        sql.NullString{String: "", Valid: true},
		ExitCode:     sql.NullInt32{Int32: 0, Valid: true},
		Created:      sql.NullInt64{Int64: 0, Valid: true},
		Started:      sql.NullInt64{Int64: 0, Valid: true},
		Finished:     sql.NullInt64{Int64: 0, Valid: true},
		Host:         sql.NullString{String: "", Valid: true},
		Runtime:      sql.NullString{String: "", Valid: true},
		Distribution: sql.NullString{String: "", Valid: true},
	}
	want := &Step{
		ID:           sql.NullInt64{Int64: 0, Valid: false},
		BuildID:      sql.NullInt64{Int64: 0, Valid: false},
		RepoID:       sql.NullInt64{Int64: 0, Valid: false},
		Number:       sql.NullInt32{Int32: 0, Valid: false},
		Name:         sql.NullString{String: "", Valid: false},
		Stage:        sql.NullString{String: "", Valid: false},
		Status:       sql.NullString{String: "", Valid: false},
		Error:        sql.NullString{String: "", Valid: false},
		ExitCode:     sql.NullInt32{Int32: 0, Valid: false},
		Created:      sql.NullInt64{Int64: 0, Valid: false},
		Started:      sql.NullInt64{Int64: 0, Valid: false},
		Finished:     sql.NullInt64{Int64: 0, Valid: false},
		Host:         sql.NullString{String: "", Valid: false},
		Runtime:      sql.NullString{String: "", Valid: false},
		Distribution: sql.NullString{String: "", Valid: false},
	}

	// run test
	got := s.Nullify()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Nullify is %v, want %v", got, want)
	}
}

func TestDatabase_Step_Nullify_Empty(t *testing.T) {
	// setup types
	s := &Step{}
	s = nil

	// run test
	got := s.Nullify()

	if got != nil {
		t.Errorf("Nullify is %v, want nil", got)
	}
}

func TestDatabase_Step_ToLibrary(t *testing.T) {
	// setup types
	num := 1
	sqlNum := sql.NullInt32{Int32: 1, Valid: true}
	num64 := int64(num)
	str := "foo"
	want := &library.Step{
		ID:           &num64,
		BuildID:      &num64,
		RepoID:       &num64,
		Number:       &num,
		Name:         &str,
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
	s := &Step{
		ID:           sql.NullInt64{Int64: num64, Valid: true},
		BuildID:      sql.NullInt64{Int64: num64, Valid: true},
		RepoID:       sql.NullInt64{Int64: num64, Valid: true},
		Number:       sqlNum,
		Name:         sql.NullString{String: str, Valid: true},
		Stage:        sql.NullString{String: str, Valid: true},
		Status:       sql.NullString{String: str, Valid: true},
		Error:        sql.NullString{String: str, Valid: true},
		ExitCode:     sqlNum,
		Created:      sql.NullInt64{Int64: num64, Valid: true},
		Started:      sql.NullInt64{Int64: num64, Valid: true},
		Finished:     sql.NullInt64{Int64: num64, Valid: true},
		Host:         sql.NullString{String: str, Valid: true},
		Runtime:      sql.NullString{String: str, Valid: true},
		Distribution: sql.NullString{String: str, Valid: true},
	}

	// run test
	got := s.ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Step_Validate(t *testing.T) {
	// setup types
	s := &Step{
		ID:      sql.NullInt64{Int64: 1, Valid: true},
		BuildID: sql.NullInt64{Int64: 1, Valid: true},
		RepoID:  sql.NullInt64{Int64: 1, Valid: true},
		Number:  sql.NullInt32{Int32: 1, Valid: true},
		Name:    sql.NullString{String: "foo", Valid: true},
	}

	// run test
	err := s.Validate()

	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestDatabase_Step_Validate_NoBuildID(t *testing.T) {
	// setup types
	s := &Step{
		ID:     sql.NullInt64{Int64: 1, Valid: true},
		RepoID: sql.NullInt64{Int64: 1, Valid: true},
		Number: sql.NullInt32{Int32: 1, Valid: true},
		Name:   sql.NullString{String: "foo", Valid: true},
	}

	// run test
	err := s.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Step_Validate_NoRepoID(t *testing.T) {
	// setup types
	s := &Step{
		ID:      sql.NullInt64{Int64: 1, Valid: true},
		BuildID: sql.NullInt64{Int64: 1, Valid: true},
		Number:  sql.NullInt32{Int32: 1, Valid: true},
		Name:    sql.NullString{String: "foo", Valid: true},
	}
	// run test
	err := s.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Step_Validate_NoNumber(t *testing.T) {
	// setup types
	s := &Step{
		ID:      sql.NullInt64{Int64: 1, Valid: true},
		BuildID: sql.NullInt64{Int64: 1, Valid: true},
		RepoID:  sql.NullInt64{Int64: 1, Valid: true},
		Name:    sql.NullString{String: "foo", Valid: true},
	}
	// run test
	err := s.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Step_Validate_NoName(t *testing.T) {
	// setup types
	s := &Step{
		ID:      sql.NullInt64{Int64: 1, Valid: true},
		BuildID: sql.NullInt64{Int64: 1, Valid: true},
		RepoID:  sql.NullInt64{Int64: 1, Valid: true},
		Number:  sql.NullInt32{Int32: 1, Valid: true},
	}
	// run test
	err := s.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_StepFromLibrary(t *testing.T) {
	// setup types
	num := 1
	sqlNum := sql.NullInt32{Int32: 1, Valid: true}
	num64 := int64(num)
	str := "foo"
	want := &Step{
		ID:           sql.NullInt64{Int64: num64, Valid: true},
		BuildID:      sql.NullInt64{Int64: num64, Valid: true},
		RepoID:       sql.NullInt64{Int64: num64, Valid: true},
		Number:       sqlNum,
		Name:         sql.NullString{String: str, Valid: true},
		Stage:        sql.NullString{String: str, Valid: true},
		Status:       sql.NullString{String: str, Valid: true},
		Error:        sql.NullString{String: str, Valid: true},
		ExitCode:     sqlNum,
		Created:      sql.NullInt64{Int64: num64, Valid: true},
		Started:      sql.NullInt64{Int64: num64, Valid: true},
		Finished:     sql.NullInt64{Int64: num64, Valid: true},
		Host:         sql.NullString{String: str, Valid: true},
		Runtime:      sql.NullString{String: str, Valid: true},
		Distribution: sql.NullString{String: str, Valid: true},
	}
	s := &library.Step{
		ID:           &num64,
		BuildID:      &num64,
		RepoID:       &num64,
		Number:       &num,
		Name:         &str,
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

	// run test
	got := StepFromLibrary(s)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("StepFromLibrary is %v, want %v", got, want)
	}
}
