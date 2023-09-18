// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"testing"
)

func TestLibrary_QueueRegistration_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		qR   *QueueRegistration
		want *QueueRegistration
	}{
		{
			qR:   testQueueRegistration(),
			want: testQueueRegistration(),
		},
		{
			qR:   new(QueueRegistration),
			want: new(QueueRegistration),

		},
	}

	// run tests
	for _, test := range tests {
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
	var w *QueueRegistration

	// setup tests
	tests := []struct {
		qR   *QueueRegistration
		want *QueueRegistration
	}{
		{
			qR:   testQueueRegistration(),
			want: testQueueRegistration(),
		},
		{
			qR:   w,
			want: new(QueueRegistration),
		},
	}

	// run tests
	for _, test := range tests {
		test.qR.SetQueueAddress(test.want.GetQueueAddress())
		test.qR.SetPublicKey(test.want.GetPublicKey())

		if test.qR.GetQueueAddress() != test.want.GetQueueAddress() {
			t.Errorf("GetQueueAddress is %v, want %v", test.qR.GetQueueAddress(), test.want.GetQueueAddress())
		}

		if test.qR.GetPublicKey() != test.want.GetPublicKey() {
			t.Errorf("GetPublicKey is %v, want %v", test.qR.GetPublicKey(), test.want.GetPublicKey())
		}
	}
}

// testQueueRegistration is a test helper function to register a QueueRegistration
// type with all fields set to a fake value.
func testQueueRegistration() *QueueRegistration {
	w := new(QueueRegistration)
	w.SetQueueAddress("http://localhost:8080")
	w.SetPublicKey("CuS+EQAzofbk3tVFS3bt5f2tIb4YiJJC4nVMFQYQElg=")

	return w
}
