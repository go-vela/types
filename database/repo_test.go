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

func TestDatabase_Repo_Nullify(t *testing.T) {
	// setup types
	r := &Repo{
		ID:           sql.NullInt64{Int64: 0, Valid: true},
		UserID:       sql.NullInt64{Int64: 0, Valid: true},
		Hash:         sql.NullString{String: "", Valid: true},
		Org:          sql.NullString{String: "", Valid: true},
		Name:         sql.NullString{String: "", Valid: true},
		FullName:     sql.NullString{String: "", Valid: true},
		Link:         sql.NullString{String: "", Valid: true},
		Clone:        sql.NullString{String: "", Valid: true},
		Branch:       sql.NullString{String: "", Valid: true},
		Timeout:      sql.NullInt64{Int64: 0, Valid: true},
		Visibility:   sql.NullString{String: "", Valid: true},
		Private:      sql.NullBool{Bool: false, Valid: true},
		Trusted:      sql.NullBool{Bool: false, Valid: true},
		Active:       sql.NullBool{Bool: false, Valid: true},
		AllowPull:    sql.NullBool{Bool: false, Valid: true},
		AllowPush:    sql.NullBool{Bool: false, Valid: true},
		AllowDeploy:  sql.NullBool{Bool: false, Valid: true},
		AllowTag:     sql.NullBool{Bool: false, Valid: true},
		AllowComment: sql.NullBool{Bool: false, Valid: true},
	}
	want := &Repo{
		ID:           sql.NullInt64{Int64: 0, Valid: false},
		UserID:       sql.NullInt64{Int64: 0, Valid: false},
		Hash:         sql.NullString{String: "", Valid: false},
		Org:          sql.NullString{String: "", Valid: false},
		Name:         sql.NullString{String: "", Valid: false},
		FullName:     sql.NullString{String: "", Valid: false},
		Link:         sql.NullString{String: "", Valid: false},
		Clone:        sql.NullString{String: "", Valid: false},
		Branch:       sql.NullString{String: "", Valid: false},
		Timeout:      sql.NullInt64{Int64: 0, Valid: false},
		Visibility:   sql.NullString{String: "", Valid: false},
		Private:      sql.NullBool{Bool: false, Valid: true},
		Trusted:      sql.NullBool{Bool: false, Valid: true},
		Active:       sql.NullBool{Bool: false, Valid: true},
		AllowPull:    sql.NullBool{Bool: false, Valid: true},
		AllowPush:    sql.NullBool{Bool: false, Valid: true},
		AllowDeploy:  sql.NullBool{Bool: false, Valid: true},
		AllowTag:     sql.NullBool{Bool: false, Valid: true},
		AllowComment: sql.NullBool{Bool: false, Valid: true},
	}

	// run test
	got := r.Nullify()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Nullify is %v, want %v", got, want)
	}
}

func TestDatabase_Repo_Nullify_Empty(t *testing.T) {
	// setup types
	var r *Repo

	// run test
	got := r.Nullify()

	if got != nil {
		t.Errorf("Nullify is %v, want nil", got)
	}
}

func TestDatabase_Repo_ToLibrary(t *testing.T) {
	// setup types
	booL := false
	num := 1
	num64 := int64(num)
	str := "foo"
	want := &library.Repo{
		ID:           &num64,
		UserID:       &num64,
		Hash:         &str,
		Org:          &str,
		Name:         &str,
		FullName:     &str,
		Link:         &str,
		Clone:        &str,
		Branch:       &str,
		Timeout:      &num64,
		Visibility:   &str,
		Private:      &booL,
		Trusted:      &booL,
		Active:       &booL,
		AllowPull:    &booL,
		AllowPush:    &booL,
		AllowDeploy:  &booL,
		AllowTag:     &booL,
		AllowComment: &booL,
	}
	r := &Repo{
		ID:           sql.NullInt64{Int64: num64, Valid: true},
		UserID:       sql.NullInt64{Int64: num64, Valid: true},
		Hash:         sql.NullString{String: str, Valid: true},
		Org:          sql.NullString{String: str, Valid: true},
		Name:         sql.NullString{String: str, Valid: true},
		FullName:     sql.NullString{String: str, Valid: true},
		Link:         sql.NullString{String: str, Valid: true},
		Clone:        sql.NullString{String: str, Valid: true},
		Branch:       sql.NullString{String: str, Valid: true},
		Timeout:      sql.NullInt64{Int64: num64, Valid: true},
		Visibility:   sql.NullString{String: str, Valid: true},
		Private:      sql.NullBool{Bool: booL, Valid: true},
		Trusted:      sql.NullBool{Bool: booL, Valid: true},
		Active:       sql.NullBool{Bool: booL, Valid: true},
		AllowPull:    sql.NullBool{Bool: booL, Valid: true},
		AllowPush:    sql.NullBool{Bool: booL, Valid: true},
		AllowDeploy:  sql.NullBool{Bool: booL, Valid: true},
		AllowTag:     sql.NullBool{Bool: booL, Valid: true},
		AllowComment: sql.NullBool{Bool: booL, Valid: true},
	}

	// run test
	got := r.ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Repo_Validate(t *testing.T) {
	// setup types
	r := &Repo{
		ID:         sql.NullInt64{Int64: 1, Valid: true},
		UserID:     sql.NullInt64{Int64: 1, Valid: true},
		Hash:       sql.NullString{String: "baz", Valid: true},
		Org:        sql.NullString{String: "foo", Valid: true},
		Name:       sql.NullString{String: "bar", Valid: true},
		FullName:   sql.NullString{String: "foo/bar", Valid: true},
		Visibility: sql.NullString{String: "public", Valid: true},
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
		ID:         sql.NullInt64{Int64: 1, Valid: true},
		Org:        sql.NullString{String: "foo", Valid: true},
		Hash:       sql.NullString{String: "baz", Valid: true},
		Name:       sql.NullString{String: "bar", Valid: true},
		FullName:   sql.NullString{String: "foo/bar", Valid: true},
		Visibility: sql.NullString{String: "public", Valid: true},
	}

	// run test
	err := r.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Repo_Validate_NoHash(t *testing.T) {
	// setup types
	r := &Repo{
		ID:         sql.NullInt64{Int64: 1, Valid: true},
		UserID:     sql.NullInt64{Int64: 1, Valid: true},
		Org:        sql.NullString{String: "foo", Valid: true},
		Name:       sql.NullString{String: "bar", Valid: true},
		FullName:   sql.NullString{String: "foo/bar", Valid: true},
		Visibility: sql.NullString{String: "public", Valid: true},
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
		ID:         sql.NullInt64{Int64: 1, Valid: true},
		UserID:     sql.NullInt64{Int64: 1, Valid: true},
		Hash:       sql.NullString{String: "baz", Valid: true},
		Name:       sql.NullString{String: "bar", Valid: true},
		FullName:   sql.NullString{String: "foo/bar", Valid: true},
		Visibility: sql.NullString{String: "public", Valid: true},
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
		ID:         sql.NullInt64{Int64: 1, Valid: true},
		UserID:     sql.NullInt64{Int64: 1, Valid: true},
		Hash:       sql.NullString{String: "baz", Valid: true},
		Org:        sql.NullString{String: "foo", Valid: true},
		FullName:   sql.NullString{String: "foo/bar", Valid: true},
		Visibility: sql.NullString{String: "public", Valid: true},
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
		ID:         sql.NullInt64{Int64: 1, Valid: true},
		UserID:     sql.NullInt64{Int64: 1, Valid: true},
		Hash:       sql.NullString{String: "baz", Valid: true},
		Org:        sql.NullString{String: "foo", Valid: true},
		Name:       sql.NullString{String: "bar", Valid: true},
		Visibility: sql.NullString{String: "public", Valid: true},
	}
	// run test
	err := r.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Repo_Validate_NoVisibility(t *testing.T) {
	// setup types
	r := &Repo{
		ID:       sql.NullInt64{Int64: 1, Valid: true},
		UserID:   sql.NullInt64{Int64: 1, Valid: true},
		Hash:     sql.NullString{String: "baz", Valid: true},
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

func TestDatabase_RepoFromLibrary(t *testing.T) {
	// setup types
	booL := false
	num := 1
	num64 := int64(num)
	str := "foo"
	want := &Repo{
		ID:           sql.NullInt64{Int64: num64, Valid: true},
		UserID:       sql.NullInt64{Int64: num64, Valid: true},
		Hash:         sql.NullString{String: str, Valid: true},
		Org:          sql.NullString{String: str, Valid: true},
		Name:         sql.NullString{String: str, Valid: true},
		FullName:     sql.NullString{String: str, Valid: true},
		Link:         sql.NullString{String: str, Valid: true},
		Clone:        sql.NullString{String: str, Valid: true},
		Branch:       sql.NullString{String: str, Valid: true},
		Timeout:      sql.NullInt64{Int64: num64, Valid: true},
		Visibility:   sql.NullString{String: str, Valid: true},
		Private:      sql.NullBool{Bool: booL, Valid: true},
		Trusted:      sql.NullBool{Bool: booL, Valid: true},
		Active:       sql.NullBool{Bool: booL, Valid: true},
		AllowPull:    sql.NullBool{Bool: booL, Valid: true},
		AllowPush:    sql.NullBool{Bool: booL, Valid: true},
		AllowDeploy:  sql.NullBool{Bool: booL, Valid: true},
		AllowTag:     sql.NullBool{Bool: booL, Valid: true},
		AllowComment: sql.NullBool{Bool: booL, Valid: true},
	}
	r := &library.Repo{
		ID:           &num64,
		UserID:       &num64,
		Hash:         &str,
		Org:          &str,
		Name:         &str,
		FullName:     &str,
		Link:         &str,
		Clone:        &str,
		Branch:       &str,
		Timeout:      &num64,
		Visibility:   &str,
		Private:      &booL,
		Trusted:      &booL,
		Active:       &booL,
		AllowPull:    &booL,
		AllowPush:    &booL,
		AllowDeploy:  &booL,
		AllowTag:     &booL,
		AllowComment: &booL,
	}

	// run test
	got := RepoFromLibrary(r)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RepoFromLibrary is %v, want %v", got, want)
	}
}
