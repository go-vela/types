// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// Repo visibility types.
const (
	// VisibilityPublic defines the visibility type for allowing any
	// users in Vela to access their repo regardless of the access
	// defined in the source control system.
	VisibilityPublic = "public"

	// VisibilityPrivate defines the visibility type for only allowing
	// users in Vela with pre-defined access in the source control
	// system to access their repo.
	VisibilityPrivate = "private"
)
