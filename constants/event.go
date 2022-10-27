// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// Build and repo events.
const (
	// EventPush defines the event type for build and repo push events.
	EventPush = "push"

	// EventPull defines the event type for build and repo pull_request events.
	EventPull = "pull_request"

	// EventPullFork defines the event type for build and repo pull_request_fork events.
	EventPullFork = "pull_request_fork"

	// EventTag defines the event type for build and repo tag events.
	EventTag = "tag"

	// EventDeploy defines the event type for build and repo deployment events.
	EventDeploy = "deployment"

	// EventComment defines the event type for comments added to a pull request.
	EventComment = "comment"

	// EventRepository defines the general event type for repo management.
	EventRepository = "repository"
)
