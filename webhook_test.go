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
		wantString string
		wantBool   bool
	}{
		{
			&Webhook{Build: testPushBuild("testing [SKIP CI]", "", constants.EventPush, true)},
			skipDirectiveMsg,
			true,
		},
		{
			&Webhook{Build: testPushBuild("testing", "wip [ci skip]", constants.EventPush, true)},
			skipDirectiveMsg,
			true,
		},
		{
			&Webhook{Build: testPushBuild("testing [skip VELA]", "", constants.EventPush, true)},
			skipDirectiveMsg,
			true,
		},
		{
			&Webhook{Build: testPushBuild("testing", "wip [vela skip]", constants.EventPush, true)},
			skipDirectiveMsg,
			true,
		},
		{
			&Webhook{Build: testPushBuild("testing ***NO_CI*** ok", "nothing", constants.EventPush, true)},
			skipDirectiveMsg,
			true,
		},
		{
			&Webhook{Build: testPushBuild("testing ok", "nothing", constants.EventPush, false)},
			skipDeleteEventMsg,
			true,
		},
		{
			&Webhook{Build: testPushBuild("testing ok", "nothing", constants.EventPush, true)},
			"",
			false,
		},
		{
			&Webhook{Build: testPushBuild("testing [SKIP CI]", "", constants.EventTag, true)},
			skipDirectiveMsg,
			true,
		},
		{
			&Webhook{Build: testPushBuild("testing", "wip [ci skip]", constants.EventTag, true)},
			skipDirectiveMsg,
			true,
		},
		{
			&Webhook{Build: testPushBuild("testing [skip VELA]", "", constants.EventTag, true)},
			skipDirectiveMsg,
			true,
		},
		{
			&Webhook{Build: testPushBuild("testing", "wip [vela skip]", constants.EventTag, true)},
			skipDirectiveMsg,
			true,
		},
		{
			&Webhook{Build: testPushBuild("testing ***NO_CI*** ok", "nothing", constants.EventTag, true)},
			skipDirectiveMsg,
			true,
		},
		{
			&Webhook{Build: testPushBuild("testing ok", "nothing", constants.EventTag, false)},
			skipDeleteEventMsg,
			true,
		},
		{
			&Webhook{Build: testPushBuild("testing ok", "nothing", constants.EventTag, true)},
			"",
			false,
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

func testPushBuild(message, title, event string, hasCommit bool) *library.Build {
	b := new(library.Build)

	b.SetEvent(event)

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
