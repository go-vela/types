// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"reflect"
	"testing"

	"github.com/go-vela/types/constants"
)

func TestLibrary_PushActions_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		actions *PushActions
		want    *PushActions
	}{
		{
			actions: testPushActions(),
			want:    testPushActions(),
		},
		{
			actions: new(PushActions),
			want:    new(PushActions),
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

func TestLibrary_PushActions_Setters(t *testing.T) {
	// setup types
	var a *PushActions

	// setup tests
	tests := []struct {
		actions *PushActions
		want    *PushActions
	}{
		{
			actions: testPushActions(),
			want:    testPushActions(),
		},
		{
			actions: a,
			want:    new(PushActions),
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

func TestLibrary_PushActions_FromMask(t *testing.T) {
	// setup types
	mask := testMask()

	want := testPushActions()

	// run test
	got := new(PushActions).FromMask(mask)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FromMask is %v, want %v", got, want)
	}
}

func TestLibrary_PushActions_ToMask(t *testing.T) {
	// setup types
	actions := testPushActions()

	want := int64(constants.AllowPushBranch | constants.AllowPushTag)

	// run test
	got := actions.ToMask()

	if want != got {
		t.Errorf("ToMask is %v, want %v", got, want)
	}
}

func TestLibrary_PullActions_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		actions *PullActions
		want    *PullActions
	}{
		{
			actions: testPullActions(),
			want:    testPullActions(),
		},
		{
			actions: new(PullActions),
			want:    new(PullActions),
		},
	}

	// run tests
	for _, test := range tests {
		if test.actions.GetOpened() != test.want.GetOpened() {
			t.Errorf("GetOpened is %v, want %v", test.actions.GetOpened(), test.want.GetOpened())
		}

		if test.actions.GetSynchronize() != test.want.GetSynchronize() {
			t.Errorf("GetSynchronize is %v, want %v", test.actions.GetSynchronize(), test.want.GetSynchronize())
		}

		if test.actions.GetEdited() != test.want.GetEdited() {
			t.Errorf("GetEdited is %v, want %v", test.actions.GetEdited(), test.want.GetEdited())
		}
	}
}

func TestLibrary_PullActions_Setters(t *testing.T) {
	// setup types
	var a *PullActions

	// setup tests
	tests := []struct {
		actions *PullActions
		want    *PullActions
	}{
		{
			actions: testPullActions(),
			want:    testPullActions(),
		},
		{
			actions: a,
			want:    new(PullActions),
		},
	}

	// run tests
	for _, test := range tests {
		test.actions.SetOpened(test.want.GetOpened())
		test.actions.SetSynchronize(test.want.GetSynchronize())
		test.actions.SetEdited(test.want.GetEdited())

		if test.actions.GetOpened() != test.want.GetOpened() {
			t.Errorf("SetOpened is %v, want %v", test.actions.GetOpened(), test.want.GetOpened())
		}

		if test.actions.GetSynchronize() != test.want.GetSynchronize() {
			t.Errorf("SetSynchronize is %v, want %v", test.actions.GetSynchronize(), test.want.GetSynchronize())
		}

		if test.actions.GetEdited() != test.want.GetEdited() {
			t.Errorf("SetEdited is %v, want %v", test.actions.GetEdited(), test.want.GetEdited())
		}
	}
}

func TestLibrary_PullActions_FromMask(t *testing.T) {
	// setup types
	mask := testMask()

	want := testPullActions()

	// run test
	got := new(PullActions).FromMask(mask)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FromMask is %v, want %v", got, want)
	}
}

func TestLibrary_PullActions_ToMask(t *testing.T) {
	// setup types
	actions := testPullActions()

	want := int64(constants.AllowPullOpen | constants.AllowPullSync)

	// run test
	got := actions.ToMask()

	if want != got {
		t.Errorf("ToMask is %v, want %v", got, want)
	}
}

func TestLibrary_DeployActions_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		actions *DeployActions
		want    *DeployActions
	}{
		{
			actions: testDeployActions(),
			want:    testDeployActions(),
		},
		{
			actions: new(DeployActions),
			want:    new(DeployActions),
		},
	}

	// run tests
	for _, test := range tests {
		if test.actions.GetCreated() != test.want.GetCreated() {
			t.Errorf("GetCreated is %v, want %v", test.actions.GetCreated(), test.want.GetCreated())
		}
	}
}

func TestLibrary_DeployActions_Setters(t *testing.T) {
	// setup types
	var a *DeployActions

	// setup tests
	tests := []struct {
		actions *DeployActions
		want    *DeployActions
	}{
		{
			actions: testDeployActions(),
			want:    testDeployActions(),
		},
		{
			actions: a,
			want:    new(DeployActions),
		},
	}

	// run tests
	for _, test := range tests {
		test.actions.SetCreated(test.want.GetCreated())

		if test.actions.GetCreated() != test.want.GetCreated() {
			t.Errorf("SetCreated is %v, want %v", test.actions.GetCreated(), test.want.GetCreated())
		}
	}
}

func TestLibrary_DeployActions_FromMask(t *testing.T) {
	// setup types
	mask := testMask()

	want := testDeployActions()

	// run test
	got := new(DeployActions).FromMask(mask)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FromMask is %v, want %v", got, want)
	}
}

func TestLibrary_DeployActions_ToMask(t *testing.T) {
	// setup types
	actions := testDeployActions()

	want := int64(constants.AllowDeployCreate)

	// run test
	got := actions.ToMask()

	if want != got {
		t.Errorf("ToMask is %v, want %v", got, want)
	}
}

func TestLibrary_CommentActions_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		actions *CommentActions
		want    *CommentActions
	}{
		{
			actions: testCommentActions(),
			want:    testCommentActions(),
		},
		{
			actions: new(CommentActions),
			want:    new(CommentActions),
		},
	}

	// run tests
	for _, test := range tests {
		if test.actions.GetCreated() != test.want.GetCreated() {
			t.Errorf("GetCreated is %v, want %v", test.actions.GetCreated(), test.want.GetCreated())
		}

		if test.actions.GetEdited() != test.want.GetEdited() {
			t.Errorf("GetEdited is %v, want %v", test.actions.GetEdited(), test.want.GetEdited())
		}
	}
}

func TestLibrary_CommentActions_Setters(t *testing.T) {
	// setup types
	var a *CommentActions

	// setup tests
	tests := []struct {
		actions *CommentActions
		want    *CommentActions
	}{
		{
			actions: testCommentActions(),
			want:    testCommentActions(),
		},
		{
			actions: a,
			want:    new(CommentActions),
		},
	}

	// run tests
	for _, test := range tests {
		test.actions.SetCreated(test.want.GetCreated())
		test.actions.SetEdited(test.want.GetEdited())

		if test.actions.GetCreated() != test.want.GetCreated() {
			t.Errorf("SetCreated is %v, want %v", test.actions.GetCreated(), test.want.GetCreated())
		}

		if test.actions.GetEdited() != test.want.GetEdited() {
			t.Errorf("SetEdited is %v, want %v", test.actions.GetEdited(), test.want.GetEdited())
		}
	}
}

func TestLibrary_CommentActions_FromMask(t *testing.T) {
	// setup types
	mask := testMask()

	want := testCommentActions()

	// run test
	got := new(CommentActions).FromMask(mask)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FromMask is %v, want %v", got, want)
	}
}

func TestLibrary_CommentActions_ToMask(t *testing.T) {
	// setup types
	actions := testCommentActions()

	want := int64(constants.AllowCommentCreate)

	// run test
	got := actions.ToMask()

	if want != got {
		t.Errorf("ToMask is %v, want %v", got, want)
	}
}

func testPushActions() *PushActions {
	push := new(PushActions)
	push.SetBranch(true)
	push.SetTag(true)

	return push
}

func testPullActions() *PullActions {
	pr := new(PullActions)
	pr.SetOpened(true)
	pr.SetSynchronize(true)
	pr.SetEdited(false)

	return pr
}

func testDeployActions() *DeployActions {
	deploy := new(DeployActions)
	deploy.SetCreated(true)

	return deploy
}

func testCommentActions() *CommentActions {
	comment := new(CommentActions)
	comment.SetCreated(true)
	comment.SetEdited(false)

	return comment
}

func testMask() int64 {
	return int64(constants.AllowPushBranch | constants.AllowPushTag | constants.AllowPullOpen | constants.AllowPullSync | constants.AllowDeployCreate | constants.AllowCommentCreate)
}
