// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
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
			&Webhook{Build: testPushBuild("testing [SKIP CI]", "", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing", "wip [ci skip]", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing [skip VELA]", "", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing", "wip [vela skip]", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing ***NO_CI*** ok", "nothing", true)},
			true,
			skipDirectiveMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing ok", "nothing", false)},
			true,
			skipDeleteEventMsg,
		},
		{
			&Webhook{Build: testPushBuild("testing ok", "nothing", true)},
			false,
			"",
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

func testPushBuild(message, title string, hasCommit bool) *library.Build {
	b := new(library.Build)

	b.SetEvent(constants.EventPush)

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
