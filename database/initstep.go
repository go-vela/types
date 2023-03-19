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
	// ErrEmptyInitStepRepoID defines the error type when an
	// InitStep type has an empty RepoID field provided.
	ErrEmptyInitStepRepoID = errors.New("empty init step repo_id provided")

	// ErrEmptyInitStepBuildID defines the error type when an
	// InitStep type has an empty BuildID field provided.
	ErrEmptyInitStepBuildID = errors.New("empty init step build_id provided")

	// ErrEmptyInitStepNumber defines the error type when an
	// InitStep type has an empty Number field provided.
	ErrEmptyInitStepNumber = errors.New("empty init step number provided")

	// ErrEmptyInitStepReporter defines the error type when an
	// InitStep type has an empty Reporter field provided.
	ErrEmptyInitStepReporter = errors.New("empty init step reporter provided")

	// ErrEmptyInitStepName defines the error type when an
	// InitStep type has an empty Name field provided.
	ErrEmptyInitStepName = errors.New("empty init step name provided")
)

// InitStep is the database representation of an init step in a build.
type InitStep struct {
	ID        sql.NullInt64  `sql:"id"`
	RepoID    sql.NullInt64  `sql:"repo_id"`
	BuildID   sql.NullInt64  `sql:"build_id"`
	StepID    sql.NullInt64  `sql:"step_id"`
	ServiceID sql.NullInt64  `sql:"service_id"`
	Number    sql.NullInt32  `sql:"number"`
	Reporter  sql.NullString `sql:"reporter"`
	Name      sql.NullString `sql:"name"`
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the InitStep type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (i *InitStep) Nullify() *InitStep {
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

	// check if the StepID field should be false
	if i.StepID.Int64 == 0 {
		i.StepID.Valid = false
	}

	// check if the ServiceID field should be false
	if i.ServiceID.Int64 == 0 {
		i.ServiceID.Valid = false
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

	return i
}

// ToLibrary converts the InitStep type
// to a library InitStep type.
func (i *InitStep) ToLibrary() *library.InitStep {
	initStep := new(library.InitStep)

	initStep.SetID(i.ID.Int64)
	initStep.SetRepoID(i.RepoID.Int64)
	initStep.SetBuildID(i.BuildID.Int64)
	if i.StepID.Valid {
		initStep.SetStepID(i.StepID.Int64)
	}
	if i.ServiceID.Valid {
		initStep.SetServiceID(i.ServiceID.Int64)
	}
	initStep.SetNumber(int(i.Number.Int32))
	initStep.SetReporter(i.Reporter.String)
	initStep.SetName(i.Name.String)

	return initStep
}

// Validate verifies the necessary fields for
// the InitStep type are populated correctly.
func (i *InitStep) Validate() error {
	// verify the RepoID field is populated
	if i.RepoID.Int64 <= 0 {
		return ErrEmptyInitStepRepoID
	}

	// verify the BuildID field is populated
	if i.BuildID.Int64 <= 0 {
		return ErrEmptyInitStepBuildID
	}

	// verify the Number field is populated
	if i.Number.Int32 <= 0 {
		return ErrEmptyInitStepNumber
	}

	// verify the Reporter field is populated
	if len(i.Reporter.String) == 0 {
		return ErrEmptyInitStepReporter
	}

	// verify the Name field is populated
	if len(i.Name.String) == 0 {
		return ErrEmptyInitStepName
	}

	// ensure that all InitStep string fields
	// that can be returned as JSON are sanitized
	// to avoid unsafe HTML content
	i.Name = sql.NullString{String: sanitize(i.Name.String), Valid: i.Name.Valid}
	i.Reporter = sql.NullString{String: sanitize(i.Reporter.String), Valid: i.Reporter.Valid}

	return nil
}

// InitStepFromLibrary converts the library InitStep type
// to a database InitStep type.
func InitStepFromLibrary(i *library.InitStep) *InitStep {
	initStep := &InitStep{
		ID:        sql.NullInt64{Int64: i.GetID(), Valid: true},
		RepoID:    sql.NullInt64{Int64: i.GetRepoID(), Valid: true},
		BuildID:   sql.NullInt64{Int64: i.GetBuildID(), Valid: true},
		StepID:    sql.NullInt64{Int64: i.GetStepID(), Valid: true},
		ServiceID: sql.NullInt64{Int64: i.GetServiceID(), Valid: true},
		Number:    sql.NullInt32{Int32: int32(i.GetNumber()), Valid: true},
		Reporter:  sql.NullString{String: i.GetReporter(), Valid: true},
		Name:      sql.NullString{String: i.GetName(), Valid: true},
	}

	return initStep.Nullify()
}
