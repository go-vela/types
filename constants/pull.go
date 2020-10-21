// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// service and step pull policies
const (
	// PullAlways defines the pull policy type for
	// a service or step to always pull an image.
	PullAlways = "always"

	// PullNotPresent defines the pull policy type for
	// a service or step to only pull an image if it doesn't exist.
	PullNotPresent = "not_present"

	// PullOnStart defines the pull policy type for
	// a service or step to only pull an image before the container starts.
	PullOnStart = "on_start"

	// PullNever defines the pull policy type for
	// a service or step to never pull an image.
	PullNever = "never"
)
