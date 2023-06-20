// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// Allowed repo events.
const (
	AllowPush   = 1 << iota // 00000001 = 1
	AllowPROpen             // 00000010 = 2
	AllowPREdit             // 00000100 = 4
	AllowPRSync             // ...
	AllowPRLabel
	AllowPRReviewRequest
	AllowTag
	AllowDeploy
	AllowCommentCreate
	AllowCommentEdit
	AllowReviewSubmit
	AllowReviewEdit
	AllowSchedule
)
