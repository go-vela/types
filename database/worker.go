// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"

	"github.com/go-vela/types/library"
)

// Worker is the database representation of a worker.
type Worker struct {
	ID            sql.NullInt64  `sql:"id"`
	Hostname      sql.NullString `sql:"hostname"`
	Path          sql.NullString `sql:"path"`
	Online        sql.NullBool   `sql:"online"`
	LastCheckedIn sql.NullTime   `sql:"last_checked_in"`
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

	// check if the URL field should be false
	if len(w.Path.String) == 0 {
		w.Hostname.Valid = false
	}

	return w
}

// ToLibrary converts the Worker type
// to a library Worker type.
func (w *Worker) ToLibrary() *library.Worker {
	worker := new(library.Worker)

	worker.SetID(w.ID.Int64)
	worker.SetHostname(w.Hostname.String)
	worker.SetPath(w.Path.String)
	worker.SetOnline(w.Online.Bool)
	worker.SetLastCheckedIn(w.LastCheckedIn.Time)
	return worker
}

// WorkerFromLibrary converts the library worker type
// to a database worker type.
func WorkerFromLibrary(w *library.Worker) *Worker {
	worker := &Worker{
		ID:            sql.NullInt64{Int64: w.GetID(), Valid: true},
		Hostname:      sql.NullString{String: w.GetHostname(), Valid: true},
		Path:          sql.NullString{String: w.GetPath(), Valid: true},
		Online:        sql.NullBool{Bool: w.GetOnline(), Valid: true},
		LastCheckedIn: sql.NullTime{Time: w.GetLastCheckedIn(), Valid: true},
	}

	return worker.Nullify()
}
