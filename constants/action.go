// SPDX-License-Identifier: Apache-2.0

package constants

// Build and repo events.
const (
	// ActionOpened defines the action for opening pull requests.
	ActionOpened = "opened"

	// ActionCreated defines the action for creating issue comments.
	ActionCreated = "created"

	// ActionEdited defines the action for the editing of pull requests or issue comments.
	ActionEdited = "edited"

	// ActionRenamed defines the action for renaming a repository.
	ActionRenamed = "renamed"

	// ActionReopened defines the action for re-opening a pull request (or issue).
	ActionReopened = "reopened"

	// ActionSynchronize defines the action for the synchronizing of pull requests.
	ActionSynchronize = "synchronize"

	// ActionLabeled defines the action for the labelin of pull requests.
	ActionLabeled = "labeled"

	// ActionTransferred defines the action for transferring repository ownership.
	ActionTransferred = "transferred"

	// ActionBranch defines the action for deleting a branch.
	ActionBranch = "branch"

	// ActionTag defines the action for deleting a tag.
	ActionTag = "tag"
)
