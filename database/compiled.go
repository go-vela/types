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
	// ErrEmptyCompiledRepoID defines the error type when a
	// Compiled type has an empty RepoID field provided.
	ErrEmptyCompiledBuildID = errors.New("empty compiled build_id provided")
)

// Compiled is the database representation of a Compiled.
type Compiled struct {
	ID      sql.NullInt64 `sql:"id"`
	BuildID sql.NullInt64 `sql:"build_id"`
	Data    []byte        `sql:"data"`
}

// Compress will manipulate the existing data for the
// Compiled by compressing that data. This produces
// a significantly smaller amount of data that is
// stored in the system.
func (c *Compiled) Compress(level int) error {
	// compress the database Compiled data
	data, err := compress(level, c.Data)
	if err != nil {
		return err
	}

	// overwrite database Compiled data with compressed Compiled data
	c.Data = data

	return nil
}

// Decompress will manipulate the existing data for the
// Compiled by decompressing that data. This allows us
// to have a significantly smaller amount of data that
// is stored in the system.
func (c *Compiled) Decompress() error {
	// decompress the database Compiled data
	data, err := decompress(c.Data)
	if err != nil {
		return err
	}

	// overwrite compressed Compiled data with decompressed Compiled data
	c.Data = data

	return nil
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the Compiled type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (c *Compiled) Nullify() *Compiled {
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

// ToLibrary converts the Compiled type
// to a library Compiled type.
func (c *Compiled) ToLibrary() *library.Compiled {
	compiled := new(library.Compiled)

	compiled.SetID(c.ID.Int64)
	compiled.SetBuildID(c.BuildID.Int64)
	compiled.SetData(c.Data)

	return compiled
}

// Validate verifies the necessary fields for
// the Compiled type are populated correctly.
func (c *Compiled) Validate() error {
	// verify the RepoID field is populated
	if c.BuildID.Int64 <= 0 {
		return ErrEmptyCompiledBuildID
	}

	return nil
}

// CompiledFromLibrary converts the library Compiled type
// to a database Compiled type.
func CompiledFromLibrary(c *library.Compiled) *Compiled {
	compiled := &Compiled{
		ID:      sql.NullInt64{Int64: c.GetID(), Valid: true},
		BuildID: sql.NullInt64{Int64: c.GetBuildID(), Valid: true},
		Data:    c.GetData(),
	}

	return compiled.Nullify()
}
