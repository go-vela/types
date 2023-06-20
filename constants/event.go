// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// Build and repo events.
const (
	// EventComment defines the event type for comments added to a pull request.
	EventComment = "comment"

	// EventDeploy defines the event type for build and repo deployment events.
	EventDeploy = "deployment"

	// EventPull defines the event type for build and repo pull_request events.
	EventPull = "pull_request"

	// EventPullReview defines the event type for build and repo pull_request_review events.
	EventPullReview = "pull_request_review"

	// EventPush defines the event type for build and repo push events.
	EventPush = "push"

	// EventRepository defines the general event type for repo management.
	EventRepository = "repository"

	// EventSchedule defines the event type for build and repo schedule events.
	EventSchedule = "schedule"

	// EventTag defines the event type for build and repo tag events.
	EventTag = "tag"
)
