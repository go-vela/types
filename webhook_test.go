// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package types

import (
	"testing"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
)

func TestWebhook_ShouldSkip(t *testing.T) {
	// set up tests
	tests := []struct {
		hook       *Webhook
		wantBool   bool
		wantString string
	}{
		{
			&Webhook{Build: testPushBuild("testing [SKIP CI]", "", constants.EventPush, "", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing", "wip [ci skip]", constants.EventPush, "", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing [skip VELA]", "", constants.EventPush, "", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing", "wip [vela skip]", constants.EventPush, "", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing ***NO_CI*** ok", "nothing", constants.EventPush, "", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing ok", "nothing", constants.EventPush, "", false)},
			true,
			skipDeleteEventMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing ok", "nothing", constants.EventPush, "", true)},
			false,
			"",
		},
		{
			&Webhook{Build: testPushBuild("testing [SKIP CI]", "", constants.EventTag, "", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing", "wip [ci skip]", constants.EventTag, "", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing [skip VELA]", "", constants.EventTag, "", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing", "wip [vela skip]", constants.EventTag, "", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing ***NO_CI*** ok", "nothing", constants.EventTag, "", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing ok", "nothing", constants.EventTag, "", false)},
			true,
			skipDeleteEventMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing ok", "nothing", constants.EventTag, "", true)},
			false,
			"",
		},
		{
			&Webhook{Build: testPushBuild("testing unsupported release action", "nothing", constants.EventRelease, "created", true)},
			true,
			skipUnsupportedReleaseActionMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing unsupported release action", "nothing", constants.EventRelease, "edited", true)},
			true,
			skipUnsupportedReleaseActionMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing unsupported release action", "nothing", constants.EventRelease, "deleted", true)},
			true,
			skipUnsupportedReleaseActionMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing unsupported release action", "nothing", constants.EventRelease, "published", true)},
			true,
			skipUnsupportedReleaseActionMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing unsupported release action", "nothing", constants.EventRelease, "unpublished", true)},
			true,
			skipUnsupportedReleaseActionMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing unsupported release action", "nothing", constants.EventRelease, "prereleased", true)},
			true,
			skipUnsupportedReleaseActionMsg,
		},
	}

	// run tests
	for _, test := range tests {
		gotBool, gotString := test.hook.ShouldSkip()

		if gotString != test.wantString {
			t.Errorf("returned an error, wanted %s, but got %s", test.wantString, gotString)
		}

		if gotBool != test.wantBool {
			t.Errorf("returned an error, wanted %v, but got %v", test.wantBool, gotBool)
		}
	}
}

func testPushBuild(message, title, event string, eventAction string, hasCommit bool) *library.Build {
	b := new(library.Build)

	b.SetEvent(event)
	b.SetEventAction(eventAction)

	if len(message) > 0 {
		b.SetMessage(message)
	}

	if len(title) > 0 {
		b.SetTitle(title)
	}

	if hasCommit {
		b.SetCommit("deadbeef")
	}

	return b
}
