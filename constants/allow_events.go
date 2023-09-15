// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// Allowed repo events.
const (
	AllowPushBranch = 1 << iota // 00000001 = 1
	AllowPullOpen               // 00000010 = 2
	AllowPullEdit               // 00000100 = 4
	AllowPullSync               // ...
	AllowPushTag
	AllowDeployCreate
	AllowCommentCreate
	AllowCommentEdit
)
