// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// Secret types.
const (
	// SecretOrg defines the secret type for a secret scoped to a specific org.
	SecretOrg = "org"

	// SecretRepo defines the secret type for a secret scoped to a specific repo.
	SecretRepo = "repo"

	// SecretShared defines the secret type for a secret shared across the installation.
	SecretShared = "shared"

	// SecretMask defines the secret mask to be used in place of secret values returned to users.
	SecretMask = "[secure]"

	// SecretLogMask defines the secret mask to be used when distributing logs that contain secrets.
	SecretLogMask = "***"

	// SecretRestrictedCharacters defines the set of characters that a secret name cannot contain.
	SecretRestrictedCharacters = "=\x00"
)
