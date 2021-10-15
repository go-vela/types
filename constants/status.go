// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// Build and step statuses.
const (
	// StatusError defines the status type for build and step error statuses.
	StatusError = "error"

	// StatusFailure defines the status type for build and step failure statuses.
	StatusFailure = "failure"

	// StatusKilled defines the status type for build and step killed statuses.
	StatusKilled = "killed"

	// StatusCanceled defines the status type for build and step canceled statuses.
	StatusCanceled = "canceled"

	// StatusPending defines the status type for build and step pending statuses.
	StatusPending = "pending"

	// StatusRunning defines the status type for build and step running statuses.
	StatusRunning = "running"

	// StatusSuccess defines the status type for build and step success statuses.
	StatusSuccess = "success"

	// StatusSkipped defines the status type for build and step skipped statuses.
	StatusSkipped = "skipped"
)
