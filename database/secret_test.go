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

func TestDatabase_Secret_ToLibrary(t *testing.T) {
	// setup types
	num64 := int64(1)
	str := "foo"
	arr := []string{"foo", "bar"}
	booL := false
	want := &library.Secret{
		ID:           &num64,
		Org:          &str,
		Repo:         &str,
		Team:         &str,
		Name:         &str,
		Value:        &str,
		Type:         &str,
		Images:       &arr,
		Events:       &arr,
		AllowCommand: &booL,
	}
	s := &Secret{
		ID:     sql.NullInt64{Int64: num64, Valid: true},
		Org:    sql.NullString{String: str, Valid: true},
		Repo:   sql.NullString{String: str, Valid: true},
		Team:   sql.NullString{String: str, Valid: true},
		Name:   sql.NullString{String: str, Valid: true},
		Value:  sql.NullString{String: str, Valid: true},
		Type:   sql.NullString{String: str, Valid: true},
		Images: arr,
		Events: arr,
		AllowCommand: sql.NullBool{Bool: booL, Valid: true},
	}

	// run test
	got := s.ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Secret_SecretFromLibrary(t *testing.T) {
	// setup types
	num64 := int64(1)
	str := "foo"
	arr := []string{"foo", "bar"}
	booL := false
	want := &Secret{
		ID:     sql.NullInt64{Int64: num64, Valid: true},
		Org:    sql.NullString{String: str, Valid: true},
		Repo:   sql.NullString{String: str, Valid: true},
		Team:   sql.NullString{String: str, Valid: true},
		Name:   sql.NullString{String: str, Valid: true},
		Value:  sql.NullString{String: str, Valid: true},
		Type:   sql.NullString{String: str, Valid: true},
		Images: arr,
		Events: arr,
		AllowCommand: sql.NullBool{Bool: booL, Valid:true},
	}

	s := &library.Secret{
		ID:           &num64,
		Org:          &str,
		Repo:         &str,
		Team:         &str,
		Name:         &str,
		Value:        &str,
		Type:         &str,
		Images:       &arr,
		Events:       &arr,
		AllowCommand: &booL,
	}

	// run test
	got := SecretFromLibrary(s)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}
func TestDatabase_Secret_Validate(t *testing.T) {
	// setup types
	s := &Secret{
		ID:    sql.NullInt64{Int64: 1, Valid: true},
		Org:   sql.NullString{String: "foo", Valid: true},
		Repo:  sql.NullString{String: "bar", Valid: true},
		Team:  sql.NullString{String: "foobar", Valid: true},
		Name:  sql.NullString{String: "foo", Valid: true},
		Value: sql.NullString{String: "bar", Valid: true},
		Type:  sql.NullString{String: "repo", Valid: true},
	}

	// run test
	err := s.Validate()

	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestDatabase_Secret_Validate_NoOrg(t *testing.T) {
	// setup types
	s := &Secret{
		ID:    sql.NullInt64{Int64: 1, Valid: true},
		Repo:  sql.NullString{String: "bar", Valid: true},
		Team:  sql.NullString{String: "foobar", Valid: true},
		Name:  sql.NullString{String: "foo", Valid: true},
		Value: sql.NullString{String: "bar", Valid: true},
		Type:  sql.NullString{String: "repo", Valid: true},
	}

	// run test
	err := s.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Secret_Validate_NoType(t *testing.T) {
	// setup types
	s := &Secret{
		ID:    sql.NullInt64{Int64: 1, Valid: true},
		Org:   sql.NullString{String: "foo", Valid: true},
		Repo:  sql.NullString{String: "bar", Valid: true},
		Team:  sql.NullString{String: "foobar", Valid: true},
		Name:  sql.NullString{String: "foo", Valid: true},
		Value: sql.NullString{String: "bar", Valid: true},
	}

	// run test
	err := s.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Secret_Validate_NoRepo(t *testing.T) {
	// setup types
	s := &Secret{
		ID:    sql.NullInt64{Int64: 1, Valid: true},
		Org:   sql.NullString{String: "foo", Valid: true},
		Team:  sql.NullString{String: "foobar", Valid: true},
		Name:  sql.NullString{String: "foo", Valid: true},
		Value: sql.NullString{String: "bar", Valid: true},
		Type:  sql.NullString{String: "repo", Valid: true},
	}

	// run test
	err := s.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Secret_Validate_NoTeam(t *testing.T) {
	// setup types
	s := &Secret{
		ID:    sql.NullInt64{Int64: 1, Valid: true},
		Org:   sql.NullString{String: "foo", Valid: true},
		Repo:  sql.NullString{String: "bar", Valid: true},
		Name:  sql.NullString{String: "foo", Valid: true},
		Value: sql.NullString{String: "bar", Valid: true},
		Type:  sql.NullString{String: "shared", Valid: true},
	}

	// run test
	err := s.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Secret_Validate_NoName(t *testing.T) {
	// setup types
	s := &Secret{
		ID:    sql.NullInt64{Int64: 1, Valid: true},
		Org:   sql.NullString{String: "foo", Valid: true},
		Repo:  sql.NullString{String: "bar", Valid: true},
		Team:  sql.NullString{String: "foobar", Valid: true},
		Value: sql.NullString{String: "bar", Valid: true},
		Type:  sql.NullString{String: "repo", Valid: true},
	}

	// run test
	err := s.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Secret_Validate_NoValue(t *testing.T) {
	// setup types
	s := &Secret{
		ID:   sql.NullInt64{Int64: 1, Valid: true},
		Org:  sql.NullString{String: "foo", Valid: true},
		Repo: sql.NullString{String: "bar", Valid: true},
		Team: sql.NullString{String: "foobar", Valid: true},
		Name: sql.NullString{String: "foo", Valid: true},
		Type: sql.NullString{String: "repo", Valid: true},
	}

	// run test
	err := s.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
