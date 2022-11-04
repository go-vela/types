// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package types

import (
	"strings"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
)

var (
	skipDeleteEventMsg              = "tag/branch delete event"
	skipUnsupportedReleaseActionMsg = "unsupported release action"
	skipDirectiveMsg                = "skip ci directive found in commit title/message"
)

// Webhook defines a struct that is used to return
// the required data when processing webhook event
// a for a source provider event.
type Webhook struct {
	Comment  string
	PRNumber int
	TagName  string
	Hook     *library.Hook
	Repo     *library.Repo
	Build    *library.Build
}

// ShouldSkip uses the build information
// associated with the given hook to determine
// whether the hook should be skipped.
func (w *Webhook) ShouldSkip() (bool, string) {
	// push, tag, or release event
	if strings.EqualFold(constants.EventPush, w.Build.GetEvent()) ||
		strings.EqualFold(constants.EventTag, w.Build.GetEvent()) ||
		strings.EqualFold(constants.EventRelease, w.Build.GetEvent()) {
		// the head commit will return null in the hook
		// payload from the scm when the event is
		// associated with a branch/tag delete
		if !strings.EqualFold(constants.EventRelease, w.Build.GetEvent()) &&
			len(w.Build.GetCommit()) == 0 {
			return true, skipDeleteEventMsg
		}

		// only release events with the action of "released" should be processed
		if strings.EqualFold(constants.EventRelease, w.Build.GetEvent()) &&
			!strings.EqualFold(constants.ActionReleased, w.Build.GetEventAction()) {
			return true, skipUnsupportedReleaseActionMsg
		}

		// check for skip ci directive in message or title
		if hasSkipDirective(w.Build.GetMessage()) ||
			hasSkipDirective(w.Build.GetTitle()) {
			return true, skipDirectiveMsg
		}
	}

	return false, ""
}

// hasSkipDirective is a small helper function
// to check a string for a number of patterns
// that signal to vela that the hook should
// be skipped from processing.
func hasSkipDirective(s string) bool {
	sl := strings.ToLower(s)

	switch {
	case strings.Contains(sl, "[skip ci]"),
		strings.Contains(sl, "[ci skip]"),
		strings.Contains(sl, "[skip vela]"),
		strings.Contains(sl, "[vela skip]"),
		strings.Contains(sl, "***no_ci***"):
		return true
	default:
		return false
	}
}
