// SPDX-License-Identifier: Apache-2.0

package constants

// Allowed repo events. NOTE: these can NOT change order.
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
