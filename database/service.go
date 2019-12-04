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
	// ErrEmptyServiceBuildID defines the error type when a
	// Service type has an empty BuildID field provided.
	ErrEmptyServiceBuildID = errors.New("empty service build_id provided")

	// ErrEmptyServiceName defines the error type when a
	// Service type has an empty Name field provided.
	ErrEmptyServiceName = errors.New("empty service name provided")

	// ErrEmptyServiceImage defines the error type when a
	// Service type has an empty Image field provided.
	ErrEmptyServiceImage = errors.New("empty service image provided")

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
	Image    sql.NullString `sql:"image"`
	Status   sql.NullString `sql:"status"`
	Error    sql.NullString `sql:"error"`
	ExitCode sql.NullInt32  `sql:"exit_code"`
	Created  sql.NullInt64  `sql:"created"`
	Started  sql.NullInt64  `sql:"started"`
	Finished sql.NullInt64  `sql:"finished"`
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the Service type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (s *Service) Nullify() *Service {
	if s == nil {
		return nil
	}

	// check if the ID field should be false
	if s.ID.Int64 == 0 {
		s.ID.Valid = false
	}

	// check if the BuildID field should be false
	if s.BuildID.Int64 == 0 {
		s.BuildID.Valid = false
	}

	// check if the RepoID field should be false
	if s.RepoID.Int64 == 0 {
		s.RepoID.Valid = false
	}

	// check if the Number field should be false
	if s.Number.Int32 == 0 {
		s.Number.Valid = false
	}

	// check if the Name field should be false
	if len(s.Name.String) == 0 {
		s.Name.Valid = false
	}

	// check if the Image field should be false
	if len(s.Image.String) == 0 {
		s.Image.Valid = false
	}

	// check if the Status field should be false
	if len(s.Status.String) == 0 {
		s.Status.Valid = false
	}

	// check if the Error field should be false
	if len(s.Error.String) == 0 {
		s.Error.Valid = false
	}

	// check if the ExitCode field should be false
	if s.ExitCode.Int32 == 0 {
		s.ExitCode.Valid = false
	}

	// check if Created field should be false
	if s.Created.Int64 == 0 {
		s.Created.Valid = false
	}

	// check if Started field should be false
	if s.Started.Int64 == 0 {
		s.Started.Valid = false
	}

	// check if Finished field should be false
	if s.Finished.Int64 == 0 {
		s.Finished.Valid = false
	}

	return s
}

// ToLibrary converts the Service type
// to a library Service type.
func (s *Service) ToLibrary() *library.Service {
	service := new(library.Service)

	service.SetID(s.ID.Int64)
	service.SetBuildID(s.BuildID.Int64)
	service.SetRepoID(s.RepoID.Int64)
	service.SetNumber(int(s.Number.Int32))
	service.SetName(s.Name.String)
	service.SetImage(s.Image.String)
	service.SetStatus(s.Status.String)
	service.SetError(s.Error.String)
	service.SetExitCode(int(s.ExitCode.Int32))
	service.SetCreated(s.Created.Int64)
	service.SetStarted(s.Started.Int64)
	service.SetFinished(s.Finished.Int64)

	return service
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

	// verify the Image field is populated
	if len(s.Image.String) == 0 {
		return ErrEmptyServiceImage
	}

	return nil
}

// ServiceFromLibrary converts the library Service type
// to a database Service type.
func ServiceFromLibrary(s *library.Service) *Service {
	service := &Service{
		ID:       sql.NullInt64{Int64: s.GetID(), Valid: true},
		BuildID:  sql.NullInt64{Int64: s.GetBuildID(), Valid: true},
		RepoID:   sql.NullInt64{Int64: s.GetRepoID(), Valid: true},
		Number:   sql.NullInt32{Int32: int32(s.GetNumber()), Valid: true},
		Name:     sql.NullString{String: s.GetName(), Valid: true},
		Image:    sql.NullString{String: s.GetImage(), Valid: true},
		Status:   sql.NullString{String: s.GetStatus(), Valid: true},
		Error:    sql.NullString{String: s.GetError(), Valid: true},
		ExitCode: sql.NullInt32{Int32: int32(s.GetExitCode()), Valid: true},
		Created:  sql.NullInt64{Int64: s.GetCreated(), Valid: true},
		Started:  sql.NullInt64{Int64: s.GetStarted(), Valid: true},
		Finished: sql.NullInt64{Int64: s.GetFinished(), Valid: true},
	}

	return service.Nullify()
}
