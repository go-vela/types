// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"errors"

	"github.com/go-vela/types/library"
)

var (
	// ErrEmptyInitBuildID defines the error type when a
	// Init type has an empty BuildID field provided.
	ErrEmptyInitBuildID = errors.New("empty init build_id provided")

	// ErrEmptyInitNumber defines the error type when a
	// Init type has an empty Number field provided.
	ErrEmptyInitNumber = errors.New("empty init number provided")

	// ErrEmptyInitReporter defines the error type when a
	// Init type has an empty Reporter field provided.
	ErrEmptyInitReporter = errors.New("empty init reporter provided")

	// ErrEmptyInitName defines the error type when a
	// Init type has an empty Name field provided.
	ErrEmptyInitName = errors.New("empty init name provided")

	// ErrEmptyInitRepoID defines the error type when a
	// Init type has an empty RepoID field provided.
	ErrEmptyInitRepoID = errors.New("empty init repo_id provided")
)

// Init is the database representation of a init in a build.
type Init struct {
	ID       sql.NullInt64  `sql:"id"`
	RepoID   sql.NullInt64  `sql:"repo_id"`
	BuildID  sql.NullInt64  `sql:"build_id"`
	Number   sql.NullInt32  `sql:"number"`
	Reporter sql.NullString `sql:"reporter"`
	Name     sql.NullString `sql:"name"`
	Mimetype sql.NullString `sql:"mimetype"`
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the Init type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (i *Init) Nullify() *Init {
	if i == nil {
		return nil
	}

	// check if the ID field should be false
	if i.ID.Int64 == 0 {
		i.ID.Valid = false
	}

	// check if the RepoID field should be false
	if i.RepoID.Int64 == 0 {
		i.RepoID.Valid = false
	}

	// check if the BuildID field should be false
	if i.BuildID.Int64 == 0 {
		i.BuildID.Valid = false
	}

	// check if the Number field should be false
	if i.Number.Int32 == 0 {
		i.Number.Valid = false
	}

	// check if the Reporter field should be false
	if len(i.Reporter.String) == 0 {
		i.Reporter.Valid = false
	}

	// check if the Name field should be false
	if len(i.Name.String) == 0 {
		i.Name.Valid = false
	}

	// check if the Mimetype field should be false
	if len(i.Mimetype.String) == 0 {
		i.Mimetype.Valid = false
	}

	return i
}

// ToLibrary converts the Init type
// to a library Init type.
func (i *Init) ToLibrary() *library.Init {
	init := new(library.Init)

	init.SetID(i.ID.Int64)
	init.SetRepoID(i.RepoID.Int64)
	init.SetBuildID(i.BuildID.Int64)
	init.SetNumber(int(i.Number.Int32))
	init.SetReporter(i.Reporter.String)
	init.SetName(i.Name.String)
	init.SetMimetype(i.Mimetype.String)

	return init
}

// Validate verifies the necessary fields for
// the Init type are populated correctly.
func (i *Init) Validate() error {
	// verify the RepoID field is populated
	if i.RepoID.Int64 <= 0 {
		return ErrEmptyInitRepoID
	}

	// verify the BuildID field is populated
	if i.BuildID.Int64 <= 0 {
		return ErrEmptyInitBuildID
	}

	// verify the Number field is populated
	if i.Number.Int32 <= 0 {
		return ErrEmptyInitNumber
	}

	// verify the Reporter field is populated
	if len(i.Reporter.String) == 0 {
		return ErrEmptyInitReporter
	}

	// verify the Name field is populated
	if len(i.Name.String) == 0 {
		return ErrEmptyInitName
	}

	// ensure that all Init string fields
	// that can be returned as JSON are sanitized
	// to avoid unsafe HTML content
	i.Name = sql.NullString{String: sanitize(i.Name.String), Valid: i.Name.Valid}
	i.Reporter = sql.NullString{String: sanitize(i.Reporter.String), Valid: i.Reporter.Valid}
	i.Mimetype = sql.NullString{String: sanitize(i.Mimetype.String), Valid: i.Mimetype.Valid}

	return nil
}

// InitFromLibrary converts the library Init type
// to a database Init type.
func InitFromLibrary(i *library.Init) *Init {
	init := &Init{
		ID:       sql.NullInt64{Int64: i.GetID(), Valid: true},
		RepoID:   sql.NullInt64{Int64: i.GetRepoID(), Valid: true},
		BuildID:  sql.NullInt64{Int64: i.GetBuildID(), Valid: true},
		Number:   sql.NullInt32{Int32: int32(i.GetNumber()), Valid: true},
		Reporter: sql.NullString{String: i.GetReporter(), Valid: true},
		Name:     sql.NullString{String: i.GetName(), Valid: true},
		Mimetype: sql.NullString{String: i.GetMimetype(), Valid: true},
	}

	return init.Nullify()
}
