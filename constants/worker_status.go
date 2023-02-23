// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// Worker statuses.
const (
	// WorkerStatusUnregistered defines the status type for a new worker
	// that has not registered yet, or needs to be re-registered.
	WorkerStatusUnregistered = "unregistered"

	// WorkerStatusAvailable defines the status type for a worker in an available state,
	// where worker RunningBuildIDs.length < worker BuildLimit.
	WorkerStatusAvailable = "available"

	// WorkerStatusBusy defines the status type for a worker in an unavailable state,
	// where worker BuildLimit == worker RunningBuildIDs.length.
	WorkerStatusBusy = "busy"

	// WorkerStatusMaintenance defines the status for a worker
	// marked in maintenance mode to prevent it from being pushed a build.
	WorkerStatusMaintenance = "maintenance"

	// WorkerStatusError defines the status for a worker in an error state.
	WorkerStatusError = "error"
)
