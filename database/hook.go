// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"errors"

	"github.com/go-vela/types/library"
)

var (
	// ErrEmptyHookNumber defines the error type when a
	// Hook type has an empty Number field provided.
	ErrEmptyHookNumber = errors.New("empty webhook number provided")

	// ErrEmptyHookRepoID defines the error type when a
	// Hook type has an empty RepoID field provided.
	ErrEmptyHookRepoID = errors.New("empty webhook repo_id provided")

	// ErrEmptyHookSourceID defines the error type when a
	// Hook type has an empty SourceID field provided.
	ErrEmptyHookSourceID = errors.New("empty webhook source_id provided")
)

// Hook is the database representation of a webhook for a repo.
type Hook struct {
	ID       sql.NullInt64  `sql:"id"`
	RepoID   sql.NullInt64  `sql:"repo_id"`
	BuildID  sql.NullInt64  `sql:"build_id"`
	Number   sql.NullInt32  `sql:"number"`
	SourceID sql.NullString `sql:"source_id"`
	Created  sql.NullInt64  `sql:"created"`
	Host     sql.NullString `sql:"host"`
	Event    sql.NullString `sql:"event"`
	Branch   sql.NullString `sql:"branch"`
	Error    sql.NullString `sql:"error"`
	Status   sql.NullString `sql:"status"`
	Link     sql.NullString `sql:"link"`
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the Hook type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (w *Hook) Nullify() *Hook {
	if w == nil {
		return nil
	}

	// check if the ID field should be false
	if w.ID.Int64 == 0 {
		w.ID.Valid = false
	}

	// check if the RepoID field should be false
	if w.RepoID.Int64 == 0 {
		w.RepoID.Valid = false
	}

	// check if the BuildID field should be false
	if w.BuildID.Int64 == 0 {
		w.BuildID.Valid = false
	}

	// check if the Number field should be false
	if w.Number.Int32 == 0 {
		w.Number.Valid = false
	}

	// check if the SourceID field should be false
	if len(w.SourceID.String) == 0 {
		w.SourceID.Valid = false
	}

	// check if the Created field should be false
	if w.Created.Int64 == 0 {
		w.Created.Valid = false
	}

	// check if the Host field should be false
	if len(w.Host.String) == 0 {
		w.Host.Valid = false
	}

	// check if the Event field should be false
	if len(w.Event.String) == 0 {
		w.Event.Valid = false
	}

	// check if the Branch field should be false
	if len(w.Branch.String) == 0 {
		w.Branch.Valid = false
	}

	// check if the Error field should be false
	if len(w.Error.String) == 0 {
		w.Error.Valid = false
	}

	// check if the Status field should be false
	if len(w.Status.String) == 0 {
		w.Status.Valid = false
	}

	// check if the Link field should be false
	if len(w.Link.String) == 0 {
		w.Link.Valid = false
	}

	return w
}

// ToLibrary converts the Hook type
// to a library Hook type.
func (w *Hook) ToLibrary() *library.Hook {
	n := int(w.Number.Int32)
	return &library.Hook{
		ID:       &w.ID.Int64,
		RepoID:   &w.RepoID.Int64,
		BuildID:  &w.BuildID.Int64,
		Number:   &n,
		SourceID: &w.SourceID.String,
		Created:  &w.Created.Int64,
		Host:     &w.Host.String,
		Event:    &w.Event.String,
		Branch:   &w.Branch.String,
		Error:    &w.Error.String,
		Status:   &w.Status.String,
		Link:     &w.Link.String,
	}
}

// Validate verifies the necessary fields for
// the Hook type are populated correctly.
func (w *Hook) Validate() error {
	// verify the RepoID field is populated
	if w.RepoID.Int64 <= 0 {
		return ErrEmptyHookRepoID
	}

	// verify the Number field is populated
	if w.Number.Int32 <= 0 {
		return ErrEmptyHookNumber
	}

	// verify the SourceID field is populated
	if len(w.SourceID.String) <= 0 {
		return ErrEmptyHookSourceID
	}

	return nil
}

// HookFromLibrary converts the Hook type
// to a library Hook type.
func HookFromLibrary(w *library.Hook) *Hook {
	webhook := &Hook{
		ID:       sql.NullInt64{Int64: w.GetID(), Valid: true},
		RepoID:   sql.NullInt64{Int64: w.GetRepoID(), Valid: true},
		BuildID:  sql.NullInt64{Int64: w.GetBuildID(), Valid: true},
		Number:   sql.NullInt32{Int32: int32(w.GetNumber()), Valid: true},
		SourceID: sql.NullString{String: w.GetSourceID(), Valid: true},
		Created:  sql.NullInt64{Int64: w.GetCreated(), Valid: true},
		Host:     sql.NullString{String: w.GetHost(), Valid: true},
		Event:    sql.NullString{String: w.GetEvent(), Valid: true},
		Branch:   sql.NullString{String: w.GetBranch(), Valid: true},
		Error:    sql.NullString{String: w.GetError(), Valid: true},
		Status:   sql.NullString{String: w.GetStatus(), Valid: true},
		Link:     sql.NullString{String: w.GetLink(), Valid: true},
	}

	return webhook.Nullify()
}
