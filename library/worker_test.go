// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestLibrary_Worker_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		worker *Worker
		want   *Worker
	}{
		{
			worker: testWorker(),
			want:   testWorker(),
		},
		{
			worker: new(Worker),
			want:   new(Worker),
		},
	}

	// run tests
	for _, test := range tests {
		if test.worker.GetID() != test.want.GetID() {
			t.Errorf("GetID is %v, want %v", test.worker.GetID(), test.want.GetID())
		}

		if test.worker.GetHostname() != test.want.GetHostname() {
			t.Errorf("GetHostname is %v, want %v", test.worker.GetHostname(), test.want.GetHostname())
		}

		if test.worker.GetAddress() != test.want.GetAddress() {
			t.Errorf("GetURL is %v, want %v", test.worker.GetAddress(), test.want.GetAddress())
		}

		if test.worker.GetActive() != test.want.GetActive() {
			t.Errorf("GetActive is %v, want %v", test.worker.GetActive(), test.want.GetActive())
		}

		if test.worker.GetLastCheckedIn() != test.want.GetLastCheckedIn() {
			t.Errorf("GetLastCheckedIn is %v, want %v", test.worker.GetLastCheckedIn(), test.want.GetLastCheckedIn())
		}
	}
}

func TestLibrary_Worker_Setters(t *testing.T) {
	// setup types
	var w *Worker

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
			want: new(Worker),
		},
	}

	// run tests
	for _, test := range tests {
		test.repo.SetID(test.want.GetID())
		test.repo.SetHostname(test.want.GetHostname())
		test.repo.SetAddress(test.want.GetAddress())
		test.repo.SetActive(test.want.GetActive())
		test.repo.SetLastCheckedIn(test.want.GetLastCheckedIn())

		if test.repo.GetID() != test.want.GetID() {
			t.Errorf("SetID is %v, want %v", test.repo.GetID(), test.want.GetID())
		}

		if test.repo.GetHostname() != test.want.GetHostname() {
			t.Errorf("SetHostname is %v, want %v", test.repo.GetHostname(), test.want.GetHostname())
		}

		if test.repo.GetAddress() != test.want.GetAddress() {
			t.Errorf("SetAddress is %v, want %v", test.repo.GetAddress(), test.want.GetAddress())
		}

		if test.repo.GetActive() != test.want.GetActive() {
			t.Errorf("SetActive is %v, want %v", test.repo.GetActive(), test.want.GetActive())
		}

		if test.repo.GetLastCheckedIn() != test.want.GetLastCheckedIn() {
			t.Errorf("SetLastCheckedIn is %v, want %v", test.repo.GetLastCheckedIn(), test.want.GetLastCheckedIn())
		}
	}
}

func TestLibrary_Worker_String(t *testing.T) {
	// setup types
	w := testWorker()

	want := fmt.Sprintf(`{
  ID: %d,
  Hostname: %s,
  Address: %s,
  Active: %t,
  LastCheckedIn: %v,
}`,
		w.GetID(),
		w.GetHostname(),
		w.GetAddress(),
		w.GetActive(),
		w.GetLastCheckedIn(),
	)

	// run test
	got := w.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

// testWorker is a test helper function to create a Worker
// type with all fields set to a fake value.
func testWorker() *Worker {
	w := new(Worker)

	w.SetID(1)
	w.SetHostname("worker_0")
	w.SetAddress("http://localhost:8080")
	w.SetActive(true)
	w.SetLastCheckedIn(time.Time{})

	return w
}
