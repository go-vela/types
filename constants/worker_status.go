// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// Worker statuses.
const (
	// WorkerStatusIdle defines the status for a worker
	// where worker RunningBuildIDs.length = 0.
	WorkerStatusIdle = "idle"

	// WorkerStatusAvailable defines the status type for a worker in an available state,
	// where worker RunningBuildIDs.length > 0 and < worker BuildLimit.
	WorkerStatusAvailable = "available"

	// WorkerStatusBusy defines the status type for a worker in an unavailable state,
	// where worker BuildLimit == worker RunningBuildIDs.length.
	WorkerStatusBusy = "busy"

	// WorkerStatusError defines the status for a worker in an error state.
	WorkerStatusError = "error"
)
