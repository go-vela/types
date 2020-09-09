// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/go-vela/types/library"
)

func TestDatabase_Worker_Nullify(t *testing.T) {
	// setup types
	var w *Worker

	want := &Worker{
		ID:            sql.NullInt64{Int64: 0, Valid: false},
		Hostname:      sql.NullString{String: "", Valid: false},
		Address:       sql.NullString{String: "", Valid: false},
		Active:        sql.NullBool{Bool: false, Valid: false},
		LastCheckedIn: sql.NullTime{Time: time.Time{}, Valid: false},
	}

	// setup tests
	tests := []struct {
		repo *Worker
		want *Worker
	}{
		{
			repo: testWorker(),
			want: testWorker(),
		},
		{
			repo: w,
			want: nil,
		},
		{
			repo: new(Worker),
			want: want,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.repo.Nullify()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Nullify is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_Worker_ToLibrary(t *testing.T) {
	// setup types
	want := new(library.Worker)

	want.SetID(1)
	want.SetHostname("worker_0")
	want.SetAddress("http://localhost:8080")
	want.SetActive(true)
	want.SetLastCheckedIn(time.Time{})

	// run test
	got := testWorker().ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_WorkerFromLibrary(t *testing.T) {
	// setup types
	w := new(library.Worker)

	w.SetID(1)
	w.SetHostname("worker_0")
	w.SetAddress("http://localhost:8080")
	w.SetActive(true)
	w.SetLastCheckedIn(time.Time{})

	want := testWorker()

	// run test
	got := WorkerFromLibrary(w)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("WorkerFromLibrary is %v, want %v", got, want)
	}
}

// testWorker is a test helper function to create a Worker
// type with all fields set to a fake value.
func testWorker() *Worker {
	return &Worker{
		ID:            sql.NullInt64{Int64: 1, Valid: true},
		Hostname:      sql.NullString{String: "worker_0", Valid: true},
		Address:       sql.NullString{String: "http://localhost:8080", Valid: true},
		Active:        sql.NullBool{Bool: true, Valid: true},
		LastCheckedIn: sql.NullTime{Time: time.Time{}, Valid: true},
	}
}
