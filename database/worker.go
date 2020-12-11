// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"errors"

	"github.com/go-vela/types/library"
	"github.com/lib/pq"
)

var (
	// ErrEmptyWorkerHost defines the error type when a
	// Worker type has an empty Host field provided.
	ErrEmptyWorkerHost = errors.New("empty worker address provided")

	// ErrEmptyWorkerAddress defines the error type when a
	// Worker type has an empty Address field provided.
	ErrEmptyWorkerAddress = errors.New("empty worker address provided")
)

// Worker is the database representation of a worker.
type Worker struct {
	ID            sql.NullInt64  `sql:"id"`
	Hostname      sql.NullString `sql:"hostname"`
	Address       sql.NullString `sql:"address"`
	Routes        pq.StringArray `sql:"routes"`
	Active        sql.NullBool   `sql:"active"`
	LastCheckedIn sql.NullInt64  `sql:"last_checked_in"`
	BuildLimit    sql.NullInt64  `sql:"build_limit"`
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the Build type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (w *Worker) Nullify() *Worker {
	if w == nil {
		return nil
	}

	// check if the ID field should be false
	if w.ID.Int64 == 0 {
		w.ID.Valid = false
	}

	// check if the Hostname field should be false
	if len(w.Hostname.String) == 0 {
		w.Hostname.Valid = false
	}

	// check if the Address field should be false
	if len(w.Address.String) == 0 {
		w.Hostname.Valid = false
	}

	// check if the LastCheckedIn field should be false
	if w.LastCheckedIn.Int64 == 0 {
		w.LastCheckedIn.Valid = false
	}

	if w.BuildLimit.Int64 == 0 {
		w.BuildLimit.Valid = false
	}

	return w
}

// ToLibrary converts the Worker type
// to a library Worker type.
func (w *Worker) ToLibrary() *library.Worker {
	worker := new(library.Worker)

	worker.SetID(w.ID.Int64)
	worker.SetHostname(w.Hostname.String)
	worker.SetAddress(w.Address.String)
	worker.SetRoutes(w.Routes)
	worker.SetActive(w.Active.Bool)
	worker.SetLastCheckedIn(w.LastCheckedIn.Int64)
	worker.SetBuildLimit(w.BuildLimit.Int64)
	return worker
}

// Validate verifies the necessary fields for
// the Worker type are populated correctly.
func (w *Worker) Validate() error {
	// verify the Host field is populated
	if len(w.Hostname.String) == 0 {
		return ErrEmptyWorkerHost
	}

	// verify the Address field is populated
	if len(w.Address.String) == 0 {
		return ErrEmptyWorkerAddress
	}

	return nil
}

// WorkerFromLibrary converts the library worker type
// to a database worker type.
func WorkerFromLibrary(w *library.Worker) *Worker {
	worker := &Worker{
		ID:            sql.NullInt64{Int64: w.GetID(), Valid: true},
		Hostname:      sql.NullString{String: w.GetHostname(), Valid: true},
		Address:       sql.NullString{String: w.GetAddress(), Valid: true},
		Routes:        w.GetRoutes(),
		Active:        sql.NullBool{Bool: w.GetActive(), Valid: true},
		LastCheckedIn: sql.NullInt64{Int64: w.GetLastCheckedIn(), Valid: true},
		BuildLimit:    sql.NullInt64{Int64: w.GetBuildLimit(), Valid: true},
	}

	return worker.Nullify()
}
