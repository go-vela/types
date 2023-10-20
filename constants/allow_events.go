// SPDX-License-Identifier: Apache-2.0

package constants

// Allowed repo events. NOTE: these can NOT change order.
const (
	AllowPushBranch = 1 << iota // 00000001 = 1
	AllowPushTag                // 00000010 = 2
	AllowPullOpen               // 00000010 = 4
	AllowPullEdit               // ...
	AllowPullSync
	_ // AllowPullLabel - Not Implemented
	_ // AllowPullReviewRequest - Not Implemented
	_ // AllowPullClosed - Not Implemented
	AllowDeployCreate
	AllowCommentCreate
	AllowCommentEdit
)
