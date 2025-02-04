// SPDX-License-Identifier: Apache-2.0

package database

import (
	"database/sql"
	"errors"

	"github.com/adhocore/gronx"

	"github.com/go-vela/types/library"
)

var (
	// ErrEmptyScheduleEntry defines the error type when a Schedule type has an empty Entry field provided.
	ErrEmptyScheduleEntry = errors.New("empty schedule entry provided")

	// ErrEmptyScheduleName defines the error type when a Schedule type has an empty Name field provided.
	ErrEmptyScheduleName = errors.New("empty schedule name provided")

	// ErrEmptyScheduleRepoID defines the error type when a Schedule type has an empty RepoID field provided.
	ErrEmptyScheduleRepoID = errors.New("empty schedule repo_id provided")

	// ErrInvalidScheduleEntry defines the error type when a Schedule type has an invalid Entry field provided.
	ErrInvalidScheduleEntry = errors.New("invalid schedule entry provided")
)

// Schedule is the database representation of a schedule for a repo.
//
// Deprecated: use Schedule from github.com/go-vela/server/database/types instead.
type Schedule struct {
	ID          sql.NullInt64  `sql:"id"`
	RepoID      sql.NullInt64  `sql:"repo_id"`
	Active      sql.NullBool   `sql:"active"`
	Name        sql.NullString `sql:"name"`
	Entry       sql.NullString `sql:"entry"`
	CreatedAt   sql.NullInt64  `sql:"created_at"`
	CreatedBy   sql.NullString `sql:"created_by"`
	UpdatedAt   sql.NullInt64  `sql:"updated_at"`
	UpdatedBy   sql.NullString `sql:"updated_by"`
	ScheduledAt sql.NullInt64  `sql:"scheduled_at"`
	Branch      sql.NullString `sql:"branch"`
}

// ScheduleFromLibrary converts the library.Schedule type to a database Schedule type.
func ScheduleFromLibrary(s *library.Schedule) *Schedule {
	schedule := &Schedule{
		ID:          sql.NullInt64{Int64: s.GetID(), Valid: true},
		RepoID:      sql.NullInt64{Int64: s.GetRepoID(), Valid: true},
		Active:      sql.NullBool{Bool: s.GetActive(), Valid: true},
		Name:        sql.NullString{String: s.GetName(), Valid: true},
		Entry:       sql.NullString{String: s.GetEntry(), Valid: true},
		CreatedAt:   sql.NullInt64{Int64: s.GetCreatedAt(), Valid: true},
		CreatedBy:   sql.NullString{String: s.GetCreatedBy(), Valid: true},
		UpdatedAt:   sql.NullInt64{Int64: s.GetUpdatedAt(), Valid: true},
		UpdatedBy:   sql.NullString{String: s.GetUpdatedBy(), Valid: true},
		ScheduledAt: sql.NullInt64{Int64: s.GetScheduledAt(), Valid: true},
		Branch:      sql.NullString{String: s.GetBranch(), Valid: true},
	}

	return schedule.Nullify()
}

// Nullify ensures the valid flag for the sql.Null types are properly set.
//
// When a field within the Schedule type is the zero value for the field, the
// valid flag is set to false causing it to be NULL in the database.
func (s *Schedule) Nullify() *Schedule {
	if s == nil {
		return nil
	}

	// check if the ID field should be valid
	s.ID.Valid = s.ID.Int64 != 0
	// check if the RepoID field should be valid
	s.RepoID.Valid = s.RepoID.Int64 != 0
	// check if the ID field should be valid
	s.Active.Valid = s.RepoID.Int64 != 0
	// check if the Name field should be valid
	s.Name.Valid = len(s.Name.String) != 0
	// check if the Entry field should be valid
	s.Entry.Valid = len(s.Entry.String) != 0
	// check if the CreatedAt field should be valid
	s.CreatedAt.Valid = s.CreatedAt.Int64 != 0
	// check if the CreatedBy field should be valid
	s.CreatedBy.Valid = len(s.CreatedBy.String) != 0
	// check if the UpdatedAt field should be valid
	s.UpdatedAt.Valid = s.UpdatedAt.Int64 != 0
	// check if the UpdatedBy field should be valid
	s.UpdatedBy.Valid = len(s.UpdatedBy.String) != 0
	// check if the ScheduledAt field should be valid
	s.ScheduledAt.Valid = s.ScheduledAt.Int64 != 0
	// check if the Branch field should be valid
	s.Branch.Valid = len(s.Branch.String) != 0

	return s
}

// ToLibrary converts the Schedule type to a library.Schedule type.
func (s *Schedule) ToLibrary() *library.Schedule {
	return &library.Schedule{
		ID:          &s.ID.Int64,
		RepoID:      &s.RepoID.Int64,
		Active:      &s.Active.Bool,
		Name:        &s.Name.String,
		Entry:       &s.Entry.String,
		CreatedAt:   &s.CreatedAt.Int64,
		CreatedBy:   &s.CreatedBy.String,
		UpdatedAt:   &s.UpdatedAt.Int64,
		UpdatedBy:   &s.UpdatedBy.String,
		ScheduledAt: &s.ScheduledAt.Int64,
		Branch:      &s.Branch.String,
	}
}

// Validate verifies the necessary fields for the Schedule type are populated correctly.
func (s *Schedule) Validate() error {
	// verify the RepoID field is populated
	if s.RepoID.Int64 <= 0 {
		return ErrEmptyScheduleRepoID
	}

	// verify the Name field is populated
	if len(s.Name.String) <= 0 {
		return ErrEmptyScheduleName
	}

	// verify the Entry field is populated
	if len(s.Entry.String) <= 0 {
		return ErrEmptyScheduleEntry
	}

	gron := gronx.New()
	if !gron.IsValid(s.Entry.String) {
		return ErrInvalidScheduleEntry
	}

	// ensure that all Schedule string fields that can be returned as JSON are sanitized to avoid unsafe HTML content
	s.Name = sql.NullString{String: sanitize(s.Name.String), Valid: s.Name.Valid}
	s.Entry = sql.NullString{String: sanitize(s.Entry.String), Valid: s.Entry.Valid}

	return nil
}
