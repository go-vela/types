// SPDX-License-Identifier: Apache-2.0

package library

import (
	"reflect"
	"testing"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library/actions"
)

func TestLibrary_Events_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		events *Events
		want   *Events
	}{
		{
			events: testEvents(),
			want:   testEvents(),
		},
		{
			events: new(Events),
			want:   new(Events),
		},
	}

	// run tests
	for _, test := range tests {
		if !reflect.DeepEqual(test.events.GetPush(), test.want.GetPush()) {
			t.Errorf("GetPush is %v, want %v", test.events.GetPush(), test.want.GetPush())
		}

		if !reflect.DeepEqual(test.events.GetPullRequest(), test.want.GetPullRequest()) {
			t.Errorf("GetPullRequest is %v, want %v", test.events.GetPush(), test.want.GetPush())
		}

		if !reflect.DeepEqual(test.events.GetDeployment(), test.want.GetDeployment()) {
			t.Errorf("GetDeployment is %v, want %v", test.events.GetPush(), test.want.GetPush())
		}

		if !reflect.DeepEqual(test.events.GetComment(), test.want.GetComment()) {
			t.Errorf("GetComment is %v, want %v", test.events.GetPush(), test.want.GetPush())
		}
	}
}

func TestLibrary_Events_Setters(t *testing.T) {
	// setup types
	var e *Events

	// setup tests
	tests := []struct {
		events *Events
		want   *Events
	}{
		{
			events: testEvents(),
			want:   testEvents(),
		},
		{
			events: e,
			want:   new(Events),
		},
	}

	// run tests
	for _, test := range tests {
		test.events.SetPush(test.want.GetPush())
		test.events.SetPullRequest(test.want.GetPullRequest())
		test.events.SetDeployment(test.want.GetDeployment())
		test.events.SetComment(test.want.GetComment())

		if !reflect.DeepEqual(test.events.GetPush(), test.want.GetPush()) {
			t.Errorf("SetPush is %v, want %v", test.events.GetPush(), test.want.GetPush())
		}

		if !reflect.DeepEqual(test.events.GetPullRequest(), test.want.GetPullRequest()) {
			t.Errorf("SetPullRequest is %v, want %v", test.events.GetPullRequest(), test.want.GetPullRequest())
		}

		if !reflect.DeepEqual(test.events.GetDeployment(), test.want.GetDeployment()) {
			t.Errorf("SetDeployment is %v, want %v", test.events.GetDeployment(), test.want.GetDeployment())
		}

		if !reflect.DeepEqual(test.events.GetComment(), test.want.GetComment()) {
			t.Errorf("SetComment is %v, want %v", test.events.GetComment(), test.want.GetComment())
		}
	}
}

func TestLibrary_Events_List(t *testing.T) {
	// setup types
	e := testEvents()

	want := []string{"push", "pull_request:opened", "pull_request:synchronize", "tag"}

	// run test
	got := e.List()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("List is %v, want %v", got, want)
	}
}

func TestLibrary_Events_NewEventsFromMask(t *testing.T) {
	// setup mask
	mask := int64(
		constants.AllowPushBranch |
			constants.AllowPushTag |
			constants.AllowPullOpen |
			constants.AllowPullSync |
			constants.AllowPullReopen,
	)

	want := testEvents()

	// run test
	got := NewEventsFromMask(mask)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewEventsFromMask is %v, want %v", got, want)
	}
}

func testEvents() *Events {
	e := new(Events)

	pr := new(actions.Pull)
	pr.SetOpened(true)
	pr.SetSynchronize(true)
	pr.SetEdited(false)
	pr.SetReopened(true)

	push := new(actions.Push)
	push.SetBranch(true)
	push.SetTag(true)

	deploy := new(actions.Deploy)
	deploy.SetCreated(false)

	comment := new(actions.Comment)
	comment.SetCreated(false)
	comment.SetEdited(false)

	e.SetPush(push)
	e.SetPullRequest(pr)
	e.SetDeployment(deploy)
	e.SetComment(comment)

	return e
}
