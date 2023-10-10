// SPDX-License-Identifier: Apache-2.0

package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
)

func TestDatabase_BuildQueue_ToLibrary(t *testing.T) {
	// setup types
	want := new(library.BuildQueue)

	want.SetNumber(1)
	want.SetStatus("running")
	want.SetCreated(1563474076)
	want.SetFullName("github/octocat")

	// run test
	got := testBuildQueue().ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_BuildQueueFromLibrary(t *testing.T) {
	// setup types
	b := new(library.BuildQueue)

	b.SetNumber(1)
	b.SetStatus("running")
	b.SetCreated(1563474076)
	b.SetFullName("github/octocat")

	want := testBuildQueue()

	// run test
	got := BuildQueueFromLibrary(b)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BuildQueueFromLibrary is %v, want %v", got, want)
	}
}

// testBuildQueue is a test helper function to create a BuildQueue
// type with all fields set to a fake value.
func testBuildQueue() *BuildQueue {
	return &BuildQueue{
		Number:   sql.NullInt32{Int32: 1, Valid: true},
		Status:   sql.NullString{String: "running", Valid: true},
		Created:  sql.NullInt64{Int64: 1563474076, Valid: true},
		FullName: sql.NullString{String: "github/octocat", Valid: true},
	}
}
