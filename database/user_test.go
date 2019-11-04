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

func TestDatabase_User_ToLibrary(t *testing.T) {
	// setup types
	booL := false
	num64 := int64(1)
	str := "foo"
	want := &library.User{
		ID:     &num64,
		Name:   &str,
		Token:  &str,
		Hash:   &str,
		Active: &booL,
		Admin:  &booL,
	}
	u := &User{
		ID:     sql.NullInt64{Int64: num64, Valid: true},
		Name:   sql.NullString{String: str, Valid: true},
		Token:  sql.NullString{String: str, Valid: true},
		Hash:   sql.NullString{String: str, Valid: true},
		Active: sql.NullBool{Bool: booL, Valid: true},
		Admin:  sql.NullBool{Bool: booL, Valid: true},
	}

	// run test
	got := u.ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_User_UserFromLibrary(t *testing.T) {
	// setup types
	booL := false
	num64 := int64(1)
	str := "foo"
	want := &User{
		ID:     sql.NullInt64{Int64: num64, Valid: true},
		Name:   sql.NullString{String: str, Valid: true},
		Token:  sql.NullString{String: str, Valid: true},
		Hash:   sql.NullString{String: str, Valid: true},
		Active: sql.NullBool{Bool: booL, Valid: true},
		Admin:  sql.NullBool{Bool: booL, Valid: true},
	}
	u := &library.User{
		ID:     &num64,
		Name:   &str,
		Token:  &str,
		Hash:   &str,
		Active: &booL,
		Admin:  &booL,
	}

	// run test
	got := UserFromLibrary(u)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_User_Validate(t *testing.T) {
	// setup types
	u := &User{
		ID:    sql.NullInt64{Int64: 1, Valid: true},
		Name:  sql.NullString{String: "foo", Valid: true},
		Token: sql.NullString{String: "bar", Valid: true},
		Hash:  sql.NullString{String: "baz", Valid: true},
	}

	// run test
	err := u.Validate()

	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestDatabase_User_Validate_NoName(t *testing.T) {
	// setup types
	u := &User{
		ID:    sql.NullInt64{Int64: 1, Valid: true},
		Token: sql.NullString{String: "bar", Valid: true},
	}
	// run test
	err := u.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_User_Validate_NoToken(t *testing.T) {
	// setup types
	u := &User{
		ID:   sql.NullInt64{Int64: 1, Valid: true},
		Name: sql.NullString{String: "foo", Valid: true},
	}
	// run test
	err := u.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_User_Validate_NoHash(t *testing.T) {
	// setup types
	u := &User{
		ID:    sql.NullInt64{Int64: 1, Valid: true},
		Name:  sql.NullString{String: "foo", Valid: true},
		Token: sql.NullString{String: "bar", Valid: true},
	}
	// run test
	err := u.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_User_Validate_NameInvalid(t *testing.T) {
	// setup types
	u := &User{
		ID:    sql.NullInt64{Int64: 1, Valid: true},
		Name:  sql.NullString{String: "!@#$%^&*()", Valid: true},
		Token: sql.NullString{String: "bar", Valid: true},
		Hash:  sql.NullString{String: "baz", Valid: true},
	}

	// run test
	err := u.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
