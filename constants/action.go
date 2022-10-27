// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// Build and repo events.
const (
	// ActionOpened defines the action for opening pull requests.
	ActionOpened = "opened"

	// ActionCreated defines the action for creating issue comments or releases.
	ActionCreated = "created"

	// ActionEdited defines the action for the editing of pull requests, issue comments, or releases.
	ActionEdited = "edited"

	// ActionPublished defines the action for the deletion of releases.
	ActionDeleted = "deleted"

	// ActionRenamed defines the action for renaming a repository.
	ActionRenamed = "renamed"

	// ActionSynchronize defines the action for the synchronizing of pull requests.
	ActionSynchronize = "synchronize"

	// ActionPublished defines the action for the publishing of releases.
	ActionPublished = "published"

	// ActionPublished defines the action for the unpublishing of releases.
	ActionUnpublished = "unpublished"

	// ActionPrereleased defines the action for the creation of a prerelease.
	ActionPrereleased = "prereleased"

	// ActionReleased defines the action for the creation of a release.
	ActionReleased = "released"
)
