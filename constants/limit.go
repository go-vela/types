// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// Limits and constraints.
const (
	// BuildTimeoutMin defines the minimum value in minutes for repo build timeout.
	BuildTimeoutMin = 1

	// BuildTimeoutMax defines the maximum value in minutes for repo build timeout.
	BuildTimeoutMax = 90

	// BuildTimeoutDefault defines the default value in minutes for repo build timeout.
	BuildTimeoutDefault = 30

	// FavoritesMaxSize defines the maximum size in characters for user favorites.
	FavoritesMaxSize = 5000
)
