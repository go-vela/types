// SPDX-License-Identifier: Apache-2.0

package actions

import (
	"reflect"
	"testing"

	"github.com/go-vela/types/constants"
)

func TestLibrary_Delete_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		actions *Delete
		want    *Delete
	}{
		{
			actions: testDelete(),
			want:    testDelete(),
		},
		{
			actions: new(Delete),
			want:    new(Delete),
		},
	}

	// run tests
	for _, test := range tests {
		if test.actions.GetBranch() != test.want.GetBranch() {
			t.Errorf("GetBranch is %v, want %v", test.actions.GetBranch(), test.want.GetBranch())
		}

		if test.actions.GetTag() != test.want.GetTag() {
			t.Errorf("GetTag is %v, want %v", test.actions.GetTag(), test.want.GetTag())
		}
	}
}

func TestLibrary_Delete_Setters(t *testing.T) {
	// setup types
	var a *Delete

	// setup tests
	tests := []struct {
		actions *Delete
		want    *Delete
	}{
		{
			actions: testDelete(),
			want:    testDelete(),
		},
		{
			actions: a,
			want:    new(Delete),
		},
	}

	// run tests
	for _, test := range tests {
		test.actions.SetBranch(test.want.GetBranch())
		test.actions.SetTag(test.want.GetTag())

		if test.actions.GetBranch() != test.want.GetBranch() {
			t.Errorf("SetBranch is %v, want %v", test.actions.GetBranch(), test.want.GetBranch())
		}

		if test.actions.GetTag() != test.want.GetTag() {
			t.Errorf("SetTag is %v, want %v", test.actions.GetTag(), test.want.GetTag())
		}
	}
}

func TestLibrary_Delete_FromMask(t *testing.T) {
	// setup types
	mask := testMask()

	want := testDelete()

	// run test
	got := new(Delete).FromMask(mask)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FromMask is %v, want %v", got, want)
	}
}

func TestLibrary_Delete_ToMask(t *testing.T) {
	// setup types
	actions := testDelete()

	want := int64(constants.AllowDeleteBranch | constants.AllowDeleteTag)

	// run test
	got := actions.ToMask()

	if want != got {
		t.Errorf("ToMask is %v, want %v", got, want)
	}
}

func testDelete() *Delete {
	delete := new(Delete)
	delete.SetBranch(true)
	delete.SetTag(true)

	return delete
}
