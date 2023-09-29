// SPDX-License-Identifier: Apache-2.0

package database

import (
	"database/sql"

	"github.com/go-vela/types/library"
)

// BuildQueue is the database representation of the builds in the queue.
type BuildQueue struct {
	Status   sql.NullString `sql:"status"`
	Number   sql.NullInt32  `sql:"number"`
	Created  sql.NullInt64  `sql:"created"`
	FullName sql.NullString `sql:"full_name"`
}

// ToLibrary converts the BuildQueue type
// to a library BuildQueue type.
func (b *BuildQueue) ToLibrary() *library.BuildQueue {
	buildQueue := new(library.BuildQueue)

	buildQueue.SetStatus(b.Status.String)
	buildQueue.SetNumber(b.Number.Int32)
	buildQueue.SetCreated(b.Created.Int64)
	buildQueue.SetFullName(b.FullName.String)

	return buildQueue
}

// BuildQueueFromLibrary converts the library BuildQueue type
// to a database build queue type.
func BuildQueueFromLibrary(b *library.BuildQueue) *BuildQueue {
	buildQueue := &BuildQueue{
		Status:   sql.NullString{String: b.GetStatus(), Valid: true},
		Number:   sql.NullInt32{Int32: b.GetNumber(), Valid: true},
		Created:  sql.NullInt64{Int64: b.GetCreated(), Valid: true},
		FullName: sql.NullString{String: b.GetFullName(), Valid: true},
	}

	return buildQueue
}
