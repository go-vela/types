// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/go-vela/types/library"
)

func TestDatabase_Hook_Nullify(t *testing.T) {
	// setup types
	h := &Hook{
		ID:       sql.NullInt64{Int64: 0, Valid: true},
		RepoID:   sql.NullInt64{Int64: 0, Valid: true},
		BuildID:  sql.NullInt64{Int64: 0, Valid: true},
		SourceID: sql.NullString{String: "", Valid: true},
		Created:  sql.NullInt64{Int64: 0, Valid: true},
		Host:     sql.NullString{String: "", Valid: true},
		Event:    sql.NullString{String: "", Valid: true},
		Branch:   sql.NullString{String: "", Valid: true},
		Error:    sql.NullString{String: "", Valid: true},
		Status:   sql.NullString{String: "", Valid: true},
		Link:     sql.NullString{String: "", Valid: true},
	}
	want := &Hook{
		ID:       sql.NullInt64{Int64: 0, Valid: false},
		RepoID:   sql.NullInt64{Int64: 0, Valid: false},
		BuildID:  sql.NullInt64{Int64: 0, Valid: false},
		SourceID: sql.NullString{String: "", Valid: false},
		Created:  sql.NullInt64{Int64: 0, Valid: false},
		Host:     sql.NullString{String: "", Valid: false},
		Event:    sql.NullString{String: "", Valid: false},
		Branch:   sql.NullString{String: "", Valid: false},
		Error:    sql.NullString{String: "", Valid: false},
		Status:   sql.NullString{String: "", Valid: false},
		Link:     sql.NullString{String: "", Valid: false},
	}

	// run test
	got := h.Nullify()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Nullify is %v, want %v", got, want)
	}
}

func TestDatabase_Hook_Nullify_Empty(t *testing.T) {
	// setup types
	h := &Hook{}
	h = nil

	// run test
	got := h.Nullify()

	if got != nil {
		t.Errorf("Nullify is %v, want nil", got)
	}
}

func TestDatabase_Hook_ToLibrary(t *testing.T) {
	// setup types
	want := new(library.Hook)
	want.SetID(1)
	want.SetRepoID(1)
	want.SetBuildID(1)
	want.SetSourceID("c8da1302-07d6-11ea-882f-4893bca275b8")
	want.SetCreated(time.Now().UTC().Unix())
	want.SetHost("github.com")
	want.SetEvent("push")
	want.SetBranch("master")
	want.SetError("")
	want.SetStatus("success")
	want.SetLink("https://github.com/github/octocat/settings/hooks/1")

	h := &Hook{
		ID:       sql.NullInt64{Int64: 1, Valid: true},
		RepoID:   sql.NullInt64{Int64: 1, Valid: true},
		BuildID:  sql.NullInt64{Int64: 1, Valid: true},
		SourceID: sql.NullString{String: "c8da1302-07d6-11ea-882f-4893bca275b8", Valid: true},
		Created:  sql.NullInt64{Int64: time.Now().UTC().Unix(), Valid: true},
		Host:     sql.NullString{String: "github.com", Valid: true},
		Event:    sql.NullString{String: "push", Valid: true},
		Branch:   sql.NullString{String: "master", Valid: true},
		Error:    sql.NullString{String: "", Valid: true},
		Status:   sql.NullString{String: "success", Valid: true},
		Link:     sql.NullString{String: "https://github.com/github/octocat/settings/hooks/1", Valid: true},
	}

	// run test
	got := h.ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Hook_Validate(t *testing.T) {
	// setup types
	h := &Hook{
		ID:       sql.NullInt64{Int64: 1, Valid: true},
		RepoID:   sql.NullInt64{Int64: 1, Valid: true},
		SourceID: sql.NullString{String: "c8da1302-07d6-11ea-882f-4893bca275b8", Valid: true},
	}

	// run test
	err := h.Validate()

	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestDatabase_Hook_Validate_NoRepoID(t *testing.T) {
	// setup types
	h := &Hook{
		ID:       sql.NullInt64{Int64: 1, Valid: true},
		SourceID: sql.NullString{String: "c8da1302-07d6-11ea-882f-4893bca275b8", Valid: true},
	}

	// run test
	err := h.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Hook_Validate_NoSourceID(t *testing.T) {
	// setup types
	h := &Hook{
		ID:     sql.NullInt64{Int64: 1, Valid: true},
		RepoID: sql.NullInt64{Int64: 1, Valid: true},
	}

	// run test
	err := h.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_HookFromLibrary(t *testing.T) {
	// setup types
	want := &Hook{
		ID:       sql.NullInt64{Int64: 1, Valid: true},
		RepoID:   sql.NullInt64{Int64: 1, Valid: true},
		BuildID:  sql.NullInt64{Int64: 1, Valid: true},
		SourceID: sql.NullString{String: "c8da1302-07d6-11ea-882f-4893bca275b8", Valid: true},
		Created:  sql.NullInt64{Int64: time.Now().UTC().Unix(), Valid: true},
		Host:     sql.NullString{String: "github.com", Valid: true},
		Event:    sql.NullString{String: "push", Valid: true},
		Branch:   sql.NullString{String: "master", Valid: true},
		Error:    sql.NullString{String: "", Valid: false},
		Status:   sql.NullString{String: "success", Valid: true},
		Link:     sql.NullString{String: "https://github.com/github/octocat/settings/hooks/1", Valid: true},
	}

	h := new(library.Hook)
	h.SetID(1)
	h.SetRepoID(1)
	h.SetBuildID(1)
	h.SetSourceID("c8da1302-07d6-11ea-882f-4893bca275b8")
	h.SetCreated(time.Now().UTC().Unix())
	h.SetHost("github.com")
	h.SetEvent("push")
	h.SetBranch("master")
	h.SetError("")
	h.SetStatus("success")
	h.SetLink("https://github.com/github/octocat/settings/hooks/1")

	// run test
	got := HookFromLibrary(h)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("HookFromLibrary is %v, want %v", got, want)
	}
}
