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
	// ErrEmptyBuildItineraryBuildID defines the error type when a
	// BuildItinerary type has an empty BuildID field provided.
	ErrEmptyBuildItineraryBuildID = errors.New("empty build_itinerary build_id provided")
)

// BuildItinerary is the database representation of a BuildItinerary.
type BuildItinerary struct {
	ID      sql.NullInt64 `sql:"id"`
	BuildID sql.NullInt64 `sql:"build_id"`
	Data    []byte        `sql:"data"`
}

// Compress will manipulate the existing data for the
// BuildItinerary by compressing that data. This produces
// a significantly smaller amount of data that is
// stored in the system.
func (c *BuildItinerary) Compress(level int) error {
	// compress the database BuildItinerary data
	data, err := compress(level, c.Data)
	if err != nil {
		return err
	}

	// overwrite database BuildItinerary data with compressed BuildItinerary data
	c.Data = data

	return nil
}

// Decompress will manipulate the existing data for the
// BuildItinerary by decompressing that data. This allows us
// to have a significantly smaller amount of data that
// is stored in the system.
func (c *BuildItinerary) Decompress() error {
	// decompress the database BuildItinerary data
	data, err := decompress(c.Data)
	if err != nil {
		return err
	}

	// overwrite compressed BuildItinerary data with decompressed BuildItinerary data
	c.Data = data

	return nil
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the BuildItinerary type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (c *BuildItinerary) Nullify() *BuildItinerary {
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

// ToLibrary converts the BuildItinerary type
// to a library BuildItinerary type.
func (c *BuildItinerary) ToLibrary() *library.BuildItinerary {
	buildItinerary := new(library.BuildItinerary)

	buildItinerary.SetID(c.ID.Int64)
	buildItinerary.SetBuildID(c.BuildID.Int64)
	buildItinerary.SetData(c.Data)

	return buildItinerary
}

// Validate verifies the necessary fields for
// the BuildItinerary type are populated correctly.
func (c *BuildItinerary) Validate() error {
	// verify the BuildID field is populated
	if c.BuildID.Int64 <= 0 {
		return ErrEmptyBuildItineraryBuildID
	}

	return nil
}

// BuildItineraryFromLibrary converts the library BuildItinerary type
// to a database BuildItinerary type.
func BuildItineraryFromLibrary(c *library.BuildItinerary) *BuildItinerary {
	buildItinerary := &BuildItinerary{
		ID:      sql.NullInt64{Int64: c.GetID(), Valid: true},
		BuildID: sql.NullInt64{Int64: c.GetBuildID(), Valid: true},
		Data:    c.GetData(),
	}

	return buildItinerary.Nullify()
}
