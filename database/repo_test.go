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

func TestDatabase_Repo_ToLibrary(t *testing.T) {
	// setup types
	booL := false
	num := 1
	num64 := int64(num)
	str := "foo"
	want := &library.Repo{
		ID:          &num64,
		UserID:      &num64,
		Org:         &str,
		Name:        &str,
		FullName:    &str,
		Link:        &str,
		Clone:       &str,
		Branch:      &str,
		Timeout:     &num64,
		Visibility:  &str,
		Private:     &booL,
		Trusted:     &booL,
		Active:      &booL,
		AllowPull:   &booL,
		AllowPush:   &booL,
		AllowDeploy: &booL,
		AllowTag:    &booL,
	}
	r := &Repo{
		ID:          sql.NullInt64{Int64: num64, Valid: true},
		UserID:      sql.NullInt64{Int64: num64, Valid: true},
		Org:         sql.NullString{String: str, Valid: true},
		Name:        sql.NullString{String: str, Valid: true},
		FullName:    sql.NullString{String: str, Valid: true},
		Link:        sql.NullString{String: str, Valid: true},
		Clone:       sql.NullString{String: str, Valid: true},
		Branch:      sql.NullString{String: str, Valid: true},
		Timeout:     sql.NullInt64{Int64: num64, Valid: true},
		Visibility:  sql.NullString{String: str, Valid: true},
		Private:     sql.NullBool{Bool: booL, Valid: true},
		Trusted:     sql.NullBool{Bool: booL, Valid: true},
		Active:      sql.NullBool{Bool: booL, Valid: true},
		AllowPull:   sql.NullBool{Bool: booL, Valid: true},
		AllowPush:   sql.NullBool{Bool: booL, Valid: true},
		AllowDeploy: sql.NullBool{Bool: booL, Valid: true},
		AllowTag:    sql.NullBool{Bool: booL, Valid: true},
	}

	// run test
	got := r.ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Repo_RepoFromLibrary(t *testing.T) {
	// setup types
	booL := false
	num := 1
	num64 := int64(num)
	str := "foo"
	want := &Repo{
		ID:          sql.NullInt64{Int64: num64, Valid: true},
		UserID:      sql.NullInt64{Int64: num64, Valid: true},
		Org:         sql.NullString{String: str, Valid: true},
		Name:        sql.NullString{String: str, Valid: true},
		FullName:    sql.NullString{String: str, Valid: true},
		Link:        sql.NullString{String: str, Valid: true},
		Clone:       sql.NullString{String: str, Valid: true},
		Branch:      sql.NullString{String: str, Valid: true},
		Timeout:     sql.NullInt64{Int64: num64, Valid: true},
		Visibility:  sql.NullString{String: str, Valid: true},
		Private:     sql.NullBool{Bool: booL, Valid: true},
		Trusted:     sql.NullBool{Bool: booL, Valid: true},
		Active:      sql.NullBool{Bool: booL, Valid: true},
		AllowPull:   sql.NullBool{Bool: booL, Valid: true},
		AllowPush:   sql.NullBool{Bool: booL, Valid: true},
		AllowDeploy: sql.NullBool{Bool: booL, Valid: true},
		AllowTag:    sql.NullBool{Bool: booL, Valid: true},
	}
	r := &library.Repo{
		ID:          &num64,
		UserID:      &num64,
		Org:         &str,
		Name:        &str,
		FullName:    &str,
		Link:        &str,
		Clone:       &str,
		Branch:      &str,
		Timeout:     &num64,
		Visibility:  &str,
		Private:     &booL,
		Trusted:     &booL,
		Active:      &booL,
		AllowPull:   &booL,
		AllowPush:   &booL,
		AllowDeploy: &booL,
		AllowTag:    &booL,
	}

	// run test
	got := RepoFromLibrary(r)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Repo_Validate(t *testing.T) {
	// setup types
	r := &Repo{
		ID:       sql.NullInt64{Int64: 1, Valid: true},
		UserID:   sql.NullInt64{Int64: 1, Valid: true},
		Org:      sql.NullString{String: "foo", Valid: true},
		Name:     sql.NullString{String: "bar", Valid: true},
		FullName: sql.NullString{String: "foo/bar", Valid: true},
	}

	// run test
	err := r.Validate()

	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestDatabase_Repo_Validate_NoUserID(t *testing.T) {
	// setup types
	r := &Repo{
		ID:       sql.NullInt64{Int64: 1, Valid: true},
		Org:      sql.NullString{String: "foo", Valid: true},
		Name:     sql.NullString{String: "bar", Valid: true},
		FullName: sql.NullString{String: "foo/bar", Valid: true},
	}

	// run test
	err := r.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Repo_Validate_NoOrg(t *testing.T) {
	// setup types
	r := &Repo{
		ID:       sql.NullInt64{Int64: 1, Valid: true},
		UserID:   sql.NullInt64{Int64: 1, Valid: true},
		Name:     sql.NullString{String: "bar", Valid: true},
		FullName: sql.NullString{String: "foo/bar", Valid: true},
	}

	// run test
	err := r.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Repo_Validate_NoName(t *testing.T) {
	// setup types
	r := &Repo{
		ID:       sql.NullInt64{Int64: 1, Valid: true},
		UserID:   sql.NullInt64{Int64: 1, Valid: true},
		Org:      sql.NullString{String: "foo", Valid: true},
		FullName: sql.NullString{String: "foo/bar", Valid: true},
	}
	// run test
	err := r.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Repo_Validate_NoFullName(t *testing.T) {
	// setup types
	r := &Repo{
		ID:     sql.NullInt64{Int64: 1, Valid: true},
		UserID: sql.NullInt64{Int64: 1, Valid: true},
		Org:    sql.NullString{String: "foo", Valid: true},
		Name:   sql.NullString{String: "bar", Valid: true},
	}
	// run test
	err := r.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
