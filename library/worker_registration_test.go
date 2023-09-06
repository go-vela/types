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
		wR   *WorkerRegistration
		want *WorkerRegistration
	}{
		{
			wR:   testWorkerRegistration(),
			want: testWorkerRegistration(),
		},
		{
			wR:   new(WorkerRegistration),
			want: new(WorkerRegistration),
		},
	}

	// run tests
	for _, test := range tests {
		if test.wR.GetRegistrationToken() != test.want.GetRegistrationToken() {
			t.Errorf("GetRegistrationToken is %v, want %v", test.wR.GetRegistrationToken(), test.want.GetRegistrationToken())
		}

		if test.wR.GetQueueAddress() != test.want.GetQueueAddress() {
			t.Errorf("GetQueueAddress is %v, want %v", test.wR.GetQueueAddress(), test.want.GetQueueAddress())
		}

		if test.wR.GetPublicKey() != test.want.GetPublicKey() {
			t.Errorf("GetPublicKey is %v, want %v", test.wR.GetPublicKey(), test.want.GetPublicKey())
		}
	}
}

func TestLibrary_QueueRegistration_Setters(t *testing.T) {
	// setup types
	var w *WorkerRegistration

	// setup tests
	tests := []struct {
		wR   *WorkerRegistration
		want *WorkerRegistration
	}{
		{
			wR:   testWorkerRegistration(),
			want: testWorkerRegistration(),
		},
		{
			wR:   w,
			want: new(WorkerRegistration),
		},
	}

	// run tests
	for _, test := range tests {
		test.wR.SetRegistrationToken(test.want.GetRegistrationToken())
		test.wR.SetQueueAddress(test.want.GetQueueAddress())
		test.wR.SetPublicKey(test.want.GetPublicKey())

		if test.wR.GetRegistrationToken() != test.want.GetRegistrationToken() {
			t.Errorf("GetRegistrationToken is %v, want %v", test.wR.GetRegistrationToken(), test.want.GetRegistrationToken())
		}

		if test.wR.GetQueueAddress() != test.want.GetQueueAddress() {
			t.Errorf("GetQueueAddress is %v, want %v", test.wR.GetQueueAddress(), test.want.GetQueueAddress())
		}

		if test.wR.GetPublicKey() != test.want.GetPublicKey() {
			t.Errorf("GetPublicKey is %v, want %v", test.wR.GetPublicKey(), test.want.GetPublicKey())
		}
	}
}

// testWorkerRegistration is a test helper function to register a Worker
// type with all fields set to a fake value.
func testWorkerRegistration() *WorkerRegistration {
	w := new(WorkerRegistration)
	w.SetRegistrationToken("1234356")
	w.SetQueueAddress("http://localhost:8080")
	w.SetPublicKey("isfnw1234")

	return w
}
