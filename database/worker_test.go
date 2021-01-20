// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"reflect"
	"testing"

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
		LastCheckedIn: sql.NullInt64{Int64: 0, Valid: false},
		BuildLimit:    sql.NullInt64{Int64: 0, Valid: false},
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
	want.SetRoutes([]string{"vela"})
	want.SetActive(true)
	want.SetLastCheckedIn(1563474077)
	want.SetBuildLimit(2)

	// run test
	got := testWorker().ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Worker_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		worker  *Worker
	}{
		{
			failure: false,
			worker:  testWorker(),
		},
		{ // no Hostname set for worker
			failure: true,
			worker: &Worker{
				ID:            sql.NullInt64{Int64: 1, Valid: true},
				Address:       sql.NullString{String: "http://localhost:8080", Valid: true},
				Active:        sql.NullBool{Bool: true, Valid: true},
				LastCheckedIn: sql.NullInt64{Int64: 1563474077, Valid: true},
			},
		},
		{ // no Address set for worker
			failure: true,
			worker: &Worker{
				ID:            sql.NullInt64{Int64: 1, Valid: true},
				Hostname:      sql.NullString{String: "worker_0", Valid: true},
				Active:        sql.NullBool{Bool: true, Valid: true},
				LastCheckedIn: sql.NullInt64{Int64: 1563474077, Valid: true},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.worker.Validate()

		if test.failure {
			if err == nil {
				t.Errorf("Validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}

func TestDatabase_WorkerFromLibrary(t *testing.T) {
	// setup types
	w := new(library.Worker)

	w.SetID(1)
	w.SetHostname("worker_0")
	w.SetAddress("http://localhost:8080")
	w.SetRoutes([]string{"vela"})
	w.SetActive(true)
	w.SetLastCheckedIn(1563474077)
	w.SetBuildLimit(2)

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
		Routes:        []string{"vela"},
		Active:        sql.NullBool{Bool: true, Valid: true},
		LastCheckedIn: sql.NullInt64{Int64: 1563474077, Valid: true},
		BuildLimit:    sql.NullInt64{Int64: 2, Valid: true},
	}
}
