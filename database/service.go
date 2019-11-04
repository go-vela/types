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
	// ErrEmptyServiceBuildID defines the error type when a
	// Service type has an empty BuildID field provided.
	ErrEmptyServiceBuildID = errors.New("empty service build_id provided")

	// ErrEmptyServiceName defines the error type when a
	// Service type has an empty Name field provided.
	ErrEmptyServiceName = errors.New("empty service name provided")

	// ErrEmptyServiceNumber defines the error type when a
	// Service type has an empty Number field provided.
	ErrEmptyServiceNumber = errors.New("empty service number provided")

	// ErrEmptyServiceRepoID defines the error type when a
	// Service type has an empty RepoID field provided.
	ErrEmptyServiceRepoID = errors.New("empty service repo_id provided")
)

// Service is the database representation of a service in a build.
type Service struct {
	ID       sql.NullInt64  `sql:"id"`
	BuildID  sql.NullInt64  `sql:"build_id"`
	RepoID   sql.NullInt64  `sql:"repo_id"`
	Number   sql.NullInt32  `sql:"number"`
	Name     sql.NullString `sql:"name"`
	Status   sql.NullString `sql:"status"`
	Error    sql.NullString `sql:"error"`
	ExitCode sql.NullInt32  `sql:"exit_code"`
	Created  sql.NullInt64  `sql:"created"`
	Started  sql.NullInt64  `sql:"started"`
	Finished sql.NullInt64  `sql:"finished"`
}

// ToLibrary converts the Service type
// to a library Service type.
func (s *Service) ToLibrary() *library.Service {
	n := int(s.Number.Int32)
	e := int(s.ExitCode.Int32)
	return &library.Service{
		ID:       &s.ID.Int64,
		BuildID:  &s.BuildID.Int64,
		RepoID:   &s.RepoID.Int64,
		Number:   &n,
		Name:     &s.Name.String,
		Status:   &s.Status.String,
		Error:    &s.Error.String,
		ExitCode: &e,
		Created:  &s.Created.Int64,
		Started:  &s.Started.Int64,
		Finished: &s.Finished.Int64,
	}
}

// Nullify is a helper function to overwrite fields in the
// service to ensure the valid flag is properly set for a sqlnull type.
func (s *Service) nullify() {
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
}

// ServiceFromLibrary converts the library Service type
// to a database Service type.
func ServiceFromLibrary(s *library.Service) *Service {
	entry := &Service{
		ID:       sql.NullInt64{Int64: s.GetID(), Valid: true},
		BuildID:  sql.NullInt64{Int64: s.GetBuildID(), Valid: true},
		RepoID:   sql.NullInt64{Int64: s.GetRepoID(), Valid: true},
		Number:   sql.NullInt32{Int32: int32(s.GetNumber()), Valid: true},
		Name:     sql.NullString{String: s.GetName(), Valid: true},
		Status:   sql.NullString{String: s.GetStatus(), Valid: true},
		Error:    sql.NullString{String: s.GetError(), Valid: true},
		ExitCode: sql.NullInt32{Int32: int32(s.GetExitCode()), Valid: true},
		Created:  sql.NullInt64{Int64: s.GetCreated(), Valid: true},
		Started:  sql.NullInt64{Int64: s.GetStarted(), Valid: true},
		Finished: sql.NullInt64{Int64: s.GetFinished(), Valid: true},
	}

	entry.nullify()

	return entry
}

// Validate verifies the necessary fields for
// the Service type are populated correctly.
func (s *Service) Validate() error {
	// verify the BuildID field is populated
	if s.BuildID.Int64 <= 0 {
		return ErrEmptyServiceBuildID
	}

	// verify the RepoID field is populated
	if s.RepoID.Int64 <= 0 {
		return ErrEmptyServiceRepoID
	}

	// verify the Number field is populated
	if s.Number.Int32 <= 0 {
		return ErrEmptyServiceNumber
	}

	// verify the Name field is populated
	if len(s.Name.String) == 0 {
		return ErrEmptyServiceName
	}

	return nil
}
