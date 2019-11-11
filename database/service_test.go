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

func TestDatabase_Service_Nullify(t *testing.T) {
	// setup types
	s := &Service{
		ID:       sql.NullInt64{Int64: 0, Valid: true},
		BuildID:  sql.NullInt64{Int64: 0, Valid: true},
		RepoID:   sql.NullInt64{Int64: 0, Valid: true},
		Number:   sql.NullInt32{Int32: 0, Valid: true},
		Name:     sql.NullString{String: "", Valid: true},
		Status:   sql.NullString{String: "", Valid: true},
		Error:    sql.NullString{String: "", Valid: true},
		ExitCode: sql.NullInt32{Int32: 0, Valid: true},
		Created:  sql.NullInt64{Int64: 0, Valid: true},
		Started:  sql.NullInt64{Int64: 0, Valid: true},
		Finished: sql.NullInt64{Int64: 0, Valid: true},
	}
	want := &Service{
		ID:       sql.NullInt64{Int64: 0, Valid: false},
		BuildID:  sql.NullInt64{Int64: 0, Valid: false},
		RepoID:   sql.NullInt64{Int64: 0, Valid: false},
		Number:   sql.NullInt32{Int32: 0, Valid: false},
		Name:     sql.NullString{String: "", Valid: false},
		Status:   sql.NullString{String: "", Valid: false},
		Error:    sql.NullString{String: "", Valid: false},
		ExitCode: sql.NullInt32{Int32: 0, Valid: false},
		Created:  sql.NullInt64{Int64: 0, Valid: false},
		Started:  sql.NullInt64{Int64: 0, Valid: false},
		Finished: sql.NullInt64{Int64: 0, Valid: false},
	}

	// run test
	got := s.Nullify()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Nullify is %v, want %v", got, want)
	}
}

func TestDatabase_Service_Nullify_Empty(t *testing.T) {
	// setup types
	s := &Service{}
	s = nil

	// run test
	got := s.Nullify()

	if got != nil {
		t.Errorf("Nullify is %v, want nil", got)
	}
}

func TestDatabase_Service_ToLibrary(t *testing.T) {
	// setup types
	num := 1
	sqlNum := sql.NullInt32{Int32: 1, Valid: true}
	num64 := int64(num)
	str := "foo"
	want := &library.Service{
		ID:       &num64,
		BuildID:  &num64,
		RepoID:   &num64,
		Number:   &num,
		Name:     &str,
		Status:   &str,
		Error:    &str,
		ExitCode: &num,
		Created:  &num64,
		Started:  &num64,
		Finished: &num64,
	}
	s := &Service{
		ID:       sql.NullInt64{Int64: num64, Valid: true},
		BuildID:  sql.NullInt64{Int64: num64, Valid: true},
		RepoID:   sql.NullInt64{Int64: num64, Valid: true},
		Number:   sqlNum,
		Name:     sql.NullString{String: str, Valid: true},
		Status:   sql.NullString{String: str, Valid: true},
		Error:    sql.NullString{String: str, Valid: true},
		ExitCode: sqlNum,
		Created:  sql.NullInt64{Int64: num64, Valid: true},
		Started:  sql.NullInt64{Int64: num64, Valid: true},
		Finished: sql.NullInt64{Int64: num64, Valid: true},
	}

	// run test
	got := s.ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Service_Validate(t *testing.T) {
	// setup types
	s := &Service{
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

func TestDatabase_Service_Validate_NoBuildID(t *testing.T) {
	// setup types
	s := &Service{
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

func TestDatabase_Service_Validate_NoRepoID(t *testing.T) {
	// setup types
	s := &Service{
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

func TestDatabase_Service_Validate_NoNumber(t *testing.T) {
	// setup types
	s := &Service{
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

func TestDatabase_Service_Validate_NoName(t *testing.T) {
	// setup types
	s := &Service{
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

func TestDatabase_ServiceFromLibrary(t *testing.T) {
	// setup types
	num := 1
	sqlNum := sql.NullInt32{Int32: 1, Valid: true}
	num64 := int64(num)
	str := "foo"
	want := &Service{
		ID:       sql.NullInt64{Int64: num64, Valid: true},
		BuildID:  sql.NullInt64{Int64: num64, Valid: true},
		RepoID:   sql.NullInt64{Int64: num64, Valid: true},
		Number:   sqlNum,
		Name:     sql.NullString{String: str, Valid: true},
		Status:   sql.NullString{String: str, Valid: true},
		Error:    sql.NullString{String: str, Valid: true},
		ExitCode: sqlNum,
		Created:  sql.NullInt64{Int64: num64, Valid: true},
		Started:  sql.NullInt64{Int64: num64, Valid: true},
		Finished: sql.NullInt64{Int64: num64, Valid: true},
	}
	s := &library.Service{
		ID:       &num64,
		BuildID:  &num64,
		RepoID:   &num64,
		Number:   &num,
		Name:     &str,
		Status:   &str,
		Error:    &str,
		ExitCode: &num,
		Created:  &num64,
		Started:  &num64,
		Finished: &num64,
	}

	// run test
	got := ServiceFromLibrary(s)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ServiceFromLibrary is %v, want %v", got, want)
	}
}
