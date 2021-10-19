// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// Build and repo events.
const (
	// EventPush defines the event type for build and repo push events.
	EventPush = "push"

	// EventPull defines the event type for build and repo pull_request events.
	EventPull = "pull_request"

	// EventTag defines the event type for build and repo tag events.
	EventTag = "tag"

	// EventDeploy defines the event type for build and repo deployment events.
	EventDeploy = "deployment"

	// EventComment defines the event type for comments added to a pull request.
	EventComment = "comment"

	// EventRepository defines the event type for repository actions such as delete and rename.
	EventRepository = "repository"
)
