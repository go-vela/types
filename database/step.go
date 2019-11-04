// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/go-vela/types/library"
)

var (
	// ErrEmptyStepBuildID defines the error type when a
	// Step type has an empty BuildID field provided.
	ErrEmptyStepBuildID = errors.New("empty step build_id provided")

	// ErrEmptyStepName defines the error type when a
	// Step type has an empty Name field provided.
	ErrEmptyStepName = errors.New("empty step name provided")

	// ErrEmptyStepNumber defines the error type when a
	// Step type has an empty Number field provided.
	ErrEmptyStepNumber = errors.New("empty step number provided")

	// ErrEmptyStepRepoID defines the error type when a
	// Step type has an empty RepoID field provided.
	ErrEmptyStepRepoID = errors.New("empty step repo_id provided")
)

// Step is the database representation of a step in a build.
type Step struct {
	ID           sql.NullInt64  `sql:"id"`
	BuildID      sql.NullInt64  `sql:"build_id"`
	RepoID       sql.NullInt64  `sql:"repo_id"`
	Number       sql.NullInt32  `sql:"number"`
	Name         sql.NullString `sql:"name"`
	Stage        sql.NullString `sql:"stage"`
	Status       sql.NullString `sql:"status"`
	Error        sql.NullString `sql:"error"`
	ExitCode     sql.NullInt32  `sql:"exit_code"`
	Created      sql.NullInt64  `sql:"created"`
	Started      sql.NullInt64  `sql:"started"`
	Finished     sql.NullInt64  `sql:"finished"`
	Host         sql.NullString `sql:"host"`
	Runtime      sql.NullString `sql:"runtime"`
	Distribution sql.NullString `sql:"distribution"`
}

// ToLibrary converts the Step type
// to a library Step type.
func (s *Step) ToLibrary() *library.Step {
	n := int(s.Number.Int32)
	e := int(s.ExitCode.Int32)
	return &library.Step{
		ID:           &s.ID.Int64,
		BuildID:      &s.BuildID.Int64,
		RepoID:       &s.RepoID.Int64,
		Number:       &n,
		Name:         &s.Name.String,
		Stage:        &s.Stage.String,
		Status:       &s.Status.String,
		Error:        &s.Error.String,
		ExitCode:     &e,
		Created:      &s.Created.Int64,
		Started:      &s.Started.Int64,
		Finished:     &s.Finished.Int64,
		Host:         &s.Host.String,
		Runtime:      &s.Runtime.String,
		Distribution: &s.Distribution.String,
	}
}

// Nullify is a helper function to overwrite fields in the
// step to ensure the valid flag is properly set for a sqlnull type.
func (s *Step) nullify() {
	// check if the ID should be false
	if s.ID.Int64 == 0 {
		s.ID.Valid = false
	}

	// check if the BuildID should be false
	if s.BuildID.Int64 == 0 {
		s.BuildID.Valid = false
	}

	// check if the RepoID should be false
	if s.RepoID.Int64 == 0 {
		s.RepoID.Valid = false
	}

	// check if the Number should be false
	if s.Number.Int32 == 0 {
		s.Number.Valid = false
	}

	// check if the Name should be false
	if strings.EqualFold(s.Name.String, "") {
		s.Name.Valid = false
	}

	// check if the Stage should be false
	if strings.EqualFold(s.Stage.String, "") {
		s.Stage.Valid = false
	}

	// check if the Status should be false
	if strings.EqualFold(s.Status.String, "") {
		s.Status.Valid = false
	}

	// check if the Error should be false
	if strings.EqualFold(s.Error.String, "") {
		s.Error.Valid = false
	}

	// check if the ExitCode should be false
	if s.ExitCode.Int32 == 0 {
		s.ExitCode.Valid = false
	}

	// check if Created should be false
	if s.Created.Int64 == 0 {
		s.Created.Valid = false
	}

	// check if Started should be false
	if s.Started.Int64 == 0 {
		s.Started.Valid = false
	}

	// check if Finished should be false
	if s.Finished.Int64 == 0 {
		s.Finished.Valid = false
	}

	// check if the Host should be false
	if strings.EqualFold(s.Host.String, "") {
		s.Host.Valid = false
	}

	// check if the Runtime should be false
	if strings.EqualFold(s.Runtime.String, "") {
		s.Runtime.Valid = false
	}

	// check if the Distrobution should be false
	if strings.EqualFold(s.Distribution.String, "") {
		s.Distribution.Valid = false
	}
}

// StepFromLibrary converts the library Step type
// to a database Step type.
func StepFromLibrary(s *library.Step) *Step {
	entry := &Step{
		ID:           sql.NullInt64{Int64: s.GetID(), Valid: true},
		BuildID:      sql.NullInt64{Int64: s.GetBuildID(), Valid: true},
		RepoID:       sql.NullInt64{Int64: s.GetRepoID(), Valid: true},
		Number:       sql.NullInt32{Int32: int32(s.GetNumber()), Valid: true},
		Name:         sql.NullString{String: s.GetName(), Valid: true},
		Stage:        sql.NullString{String: s.GetStage(), Valid: true},
		Status:       sql.NullString{String: s.GetStatus(), Valid: true},
		Error:        sql.NullString{String: s.GetError(), Valid: true},
		ExitCode:     sql.NullInt32{Int32: int32(s.GetExitCode()), Valid: true},
		Created:      sql.NullInt64{Int64: s.GetCreated(), Valid: true},
		Started:      sql.NullInt64{Int64: s.GetStarted(), Valid: true},
		Finished:     sql.NullInt64{Int64: s.GetFinished(), Valid: true},
		Host:         sql.NullString{String: s.GetHost(), Valid: true},
		Runtime:      sql.NullString{String: s.GetRuntime(), Valid: true},
		Distribution: sql.NullString{String: s.GetDistribution(), Valid: true},
	}

	entry.nullify()

	return entry
}

// Validate verifies the necessary fields for
// the Step type are populated correctly.
func (s *Step) Validate() error {
	// verify the BuildID field is populated
	if s.BuildID.Int64 <= 0 {
		return ErrEmptyStepBuildID
	}

	// verify the RepoID field is populated
	if s.RepoID.Int64 <= 0 {
		return ErrEmptyStepRepoID
	}

	// verify the Number field is populated
	if s.Number.Int32 <= 0 {
		return ErrEmptyStepNumber
	}

	// verify the Name field is populated
	if len(s.Name.String) == 0 {
		return ErrEmptyStepName
	}

	return nil
}
