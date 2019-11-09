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

func TestDatabase_Secret_Nullify(t *testing.T) {
	// setup types
	s := &Secret{
		ID:     sql.NullInt64{Int64: 0, Valid: true},
		Org:    sql.NullString{String: "", Valid: true},
		Repo:   sql.NullString{String: "", Valid: true},
		Team:   sql.NullString{String: "", Valid: true},
		Name:   sql.NullString{String: "", Valid: true},
		Value:  sql.NullString{String: "", Valid: true},
		Type:   sql.NullString{String: "", Valid: true},
		Images: []string{},
		Events: []string{},
	}
	want := &Secret{
		ID:     sql.NullInt64{Int64: 0, Valid: false},
		Org:    sql.NullString{String: "", Valid: false},
		Repo:   sql.NullString{String: "", Valid: false},
		Team:   sql.NullString{String: "", Valid: false},
		Name:   sql.NullString{String: "", Valid: false},
		Value:  sql.NullString{String: "", Valid: false},
		Type:   sql.NullString{String: "", Valid: false},
		Images: []string{},
		Events: []string{},
	}

	// run test
	got := s.Nullify()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Nullify is %v, want %v", got, want)
	}
}

func TestDatabase_Secret_Nullify_Empty(t *testing.T) {
	// setup types
	s := &Secret{}
	s = nil

	// run test
	got := s.Nullify()

	if got != nil {
		t.Errorf("Nullify is %v, want nil", got)
	}
}

func TestDatabase_Secret_ToLibrary(t *testing.T) {
	// setup types
	num64 := int64(1)
	str := "foo"
	arr := []string{"foo", "bar"}
	want := &library.Secret{
		ID:     &num64,
		Org:    &str,
		Repo:   &str,
		Team:   &str,
		Name:   &str,
		Value:  &str,
		Type:   &str,
		Images: &arr,
		Events: &arr,
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
	}

	// run test
	got := s.ToLibrary()

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

func TestDatabase_SecretFromLibrary(t *testing.T) {
	// setup types
	num64 := int64(1)
	str := "foo"
	arr := []string{"foo", "bar"}
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
	}

	s := &library.Secret{
		ID:     &num64,
		Org:    &str,
		Repo:   &str,
		Team:   &str,
		Name:   &str,
		Value:  &str,
		Type:   &str,
		Images: &arr,
		Events: &arr,
	}

	// run test
	got := SecretFromLibrary(s)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SecretFromLibrary is %v, want %v", got, want)
	}
}
