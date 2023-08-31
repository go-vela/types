// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"testing"
)

func TestLibrary_Queue_Registration_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		qR   *WorkerRegistration
		want *WorkerRegistration
	}{
		{
			qR:   testQueueRegistration(),
			want: testQueueRegistration(),
		},
		{
			qR:   new(WorkerRegistration),
			want: new(WorkerRegistration),
		},
	}

	// run tests
	for _, test := range tests {
		if test.qR.GetRegistrationToken() != test.want.GetRegistrationToken() {
			t.Errorf("GetRegistrationToken is %v, want %v", test.qR.GetRegistrationToken(), test.want.GetRegistrationToken())
		}

		if test.qR.GetQueueAddress() != test.want.GetQueueAddress() {
			t.Errorf("GetQueueAddress is %v, want %v", test.qR.GetQueueAddress(), test.want.GetQueueAddress())
		}

		if test.qR.GetPublicKey() != test.want.GetPublicKey() {
			t.Errorf("GetPublicKey is %v, want %v", test.qR.GetPublicKey(), test.want.GetPublicKey())
		}
	}
}

func TestLibrary_QueueRegistration_Setters(t *testing.T) {
	// setup types
	var w *WorkerRegistration

	// setup tests
	tests := []struct {
		qR   *WorkerRegistration
		want *WorkerRegistration
	}{
		{
			qR:   testQueueRegistration(),
			want: testQueueRegistration(),
		},
		{
			qR:   w,
			want: new(WorkerRegistration),
		},
	}

	// run tests
	for _, test := range tests {
		test.qR.SetRegistrationToken(test.want.GetRegistrationToken())
		test.qR.SetQueueAddress(test.want.GetQueueAddress())
		test.qR.SetPublicKey(test.want.GetPublicKey())

		if test.qR.GetRegistrationToken() != test.want.GetRegistrationToken() {
			t.Errorf("GetRegistrationToken is %v, want %v", test.qR.GetRegistrationToken(), test.want.GetRegistrationToken())
		}

		if test.qR.GetQueueAddress() != test.want.GetQueueAddress() {
			t.Errorf("GetQueueAddress is %v, want %v", test.qR.GetQueueAddress(), test.want.GetQueueAddress())
		}

		if test.qR.GetPublicKey() != test.want.GetPublicKey() {
			t.Errorf("GetPublicKey is %v, want %v", test.qR.GetPublicKey(), test.want.GetPublicKey())
		}
	}
}

// testWorker is a test helper function to create a Worker
// type with all fields set to a fake value.
func testQueueRegistration() *WorkerRegistration {
	w := new(WorkerRegistration)
	w.SetRegistrationToken("1234356")
	w.SetPublicKey("http://localhost:8080")
	w.SetPublicKey("worker_0")

	return w
}
