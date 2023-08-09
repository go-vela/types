// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"errors"

	"github.com/go-vela/types/library"
)

var (
	// ErrEmptyBuildExecutableBuildID defines the error type when a
	// BuildExecutable type has an empty BuildID field provided.
	ErrEmptyBuildExecutableBuildID = errors.New("empty build_executable build_id provided")
)

// BuildExecutable is the database representation of a BuildExecutable.
type BuildExecutable struct {
	ID      sql.NullInt64 `sql:"id"`
	BuildID sql.NullInt64 `sql:"build_id"`
	Data    []byte        `sql:"data"`
}

// Compress will manipulate the existing data for the
// BuildExecutable by compressing that data. This produces
// a significantly smaller amount of data that is
// stored in the system.
func (c *BuildExecutable) Compress(level int) error {
	// compress the database BuildExecutable data
	data, err := compress(level, c.Data)
	if err != nil {
		return err
	}

	// overwrite database BuildExecutable data with compressed BuildExecutable data
	c.Data = data

	return nil
}

// Decompress will manipulate the existing data for the
// BuildExecutable by decompressing that data. This allows us
// to have a significantly smaller amount of data that
// is stored in the system.
func (c *BuildExecutable) Decompress() error {
	// decompress the database BuildExecutable data
	data, err := decompress(c.Data)
	if err != nil {
		return err
	}

	// overwrite compressed BuildExecutable data with decompressed BuildExecutable data
	c.Data = data

	return nil
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the BuildExecutable type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (c *BuildExecutable) Nullify() *BuildExecutable {
	if c == nil {
		return nil
	}

	// check if the ID field should be false
	if c.ID.Int64 == 0 {
		c.ID.Valid = false
	}

	// check if the BuildID field should be false
	if c.BuildID.Int64 == 0 {
		c.BuildID.Valid = false
	}

	return c
}

// ToLibrary converts the BuildExecutable type
// to a library BuildExecutable type.
func (c *BuildExecutable) ToLibrary() *library.BuildExecutable {
	buildExecutable := new(library.BuildExecutable)

	buildExecutable.SetID(c.ID.Int64)
	buildExecutable.SetBuildID(c.BuildID.Int64)
	buildExecutable.SetData(c.Data)

	return buildExecutable
}

// Validate verifies the necessary fields for
// the BuildExecutable type are populated correctly.
func (c *BuildExecutable) Validate() error {
	// verify the BuildID field is populated
	if c.BuildID.Int64 <= 0 {
		return ErrEmptyBuildExecutableBuildID
	}

	return nil
}

// BuildExecutableFromLibrary converts the library BuildExecutable type
// to a database BuildExecutable type.
func BuildExecutableFromLibrary(c *library.BuildExecutable) *BuildExecutable {
	buildExecutable := &BuildExecutable{
		ID:      sql.NullInt64{Int64: c.GetID(), Valid: true},
		BuildID: sql.NullInt64{Int64: c.GetBuildID(), Valid: true},
		Data:    c.GetData(),
	}

	return buildExecutable.Nullify()
}
