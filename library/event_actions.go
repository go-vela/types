// SPDX-License-Identifier: Apache-2.0

package library

import "github.com/go-vela/types/constants"

// PushActions is the library representation of the various actions associated
// with the push event webhook from the SCM.
type PushActions struct {
	Branch *bool `json:"branch"`
	Tag    *bool `json:"tag"`
}

// PullActions is the library representation of the various actions associated
// with the pull_request event webhook from the SCM.
type PullActions struct {
	Opened      *bool `json:"opened"`
	Edited      *bool `json:"edited"`
	Synchronize *bool `json:"synchronize"`
}

// DeployActions is the library representation of the various actions associated
// with the deploy event webhook from the SCM.
type DeployActions struct {
	Created *bool `json:"created"`
}

// CommentActions is the library representation of the various actions associated
// with the comment event webhook from the SCM.
type CommentActions struct {
	Created *bool `json:"created"`
	Edited  *bool `json:"edited"`
}

// **
// PUSH ACTIONS
// **

// FromMask returns the PushActions type resulting from the provided integer mask.
func (a *PushActions) FromMask(mask int64) *PushActions {
	a.SetBranch(mask&constants.AllowPushBranch > 0)
	a.SetTag(mask&constants.AllowPushTag > 0)

	return a
}

// ToMask returns the integer mask of the values for the PushActions set.
func (a *PushActions) ToMask() int64 {
	mask := int64(0)

	if a.GetBranch() {
		mask = mask | constants.AllowPushBranch
	}

	if a.GetTag() {
		mask = mask | constants.AllowPushTag
	}

	return mask
}

// GetBranch returns the Branch field from the provided PushActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *PushActions) GetBranch() bool {
	// return zero value if PushActions type or Branch field is nil
	if a == nil || a.Branch == nil {
		return false
	}

	return *a.Branch
}

// GetTag returns the Tag field from the provided PushActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *PushActions) GetTag() bool {
	// return zero value if PushActions type or Tag field is nil
	if a == nil || a.Tag == nil {
		return false
	}

	return *a.Tag
}

// SetBranch sets the PushActions Branch field.
//
// When the provided PushActions type is nil, it
// will set nothing and immediately return.
func (a *PushActions) SetBranch(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Branch = &v
}

// SetTag sets the PushActions Tag field.
//
// When the provided PushActions type is nil, it
// will set nothing and immediately return.
func (a *PushActions) SetTag(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Tag = &v
}

// **
// PULL ACTIONS
// **

// FromMask returns the PullActions type resulting from the provided integer mask.
func (a *PullActions) FromMask(mask int64) *PullActions {
	a.SetOpened(mask&constants.AllowPullOpen > 0)
	a.SetSynchronize(mask&constants.AllowPullSync > 0)
	a.SetEdited(mask&constants.AllowPullEdit > 0)

	return a
}

// ToMask returns the integer mask of the values for the PullActions set.
func (a *PullActions) ToMask() int64 {
	mask := int64(0)

	if a.GetOpened() {
		mask = mask | constants.AllowPullOpen
	}

	if a.GetSynchronize() {
		mask = mask | constants.AllowPullSync
	}

	if a.GetEdited() {
		mask = mask | constants.AllowPullEdit
	}

	return mask
}

// GetOpened returns the Opened field from the provided PullActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *PullActions) GetOpened() bool {
	// return zero value if PullActions type or Opened field is nil
	if a == nil || a.Opened == nil {
		return false
	}

	return *a.Opened
}

// GetSynchronize returns the Synchronize field from the provided PullActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *PullActions) GetSynchronize() bool {
	// return zero value if PullActions type or Synchronize field is nil
	if a == nil || a.Synchronize == nil {
		return false
	}

	return *a.Synchronize
}

// GetEdited returns the Edited field from the provided PullActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *PullActions) GetEdited() bool {
	// return zero value if PullActions type or Edited field is nil
	if a == nil || a.Edited == nil {
		return false
	}

	return *a.Edited
}

// SetOpened sets the PullActions Opened field.
//
// When the provided PullActions type is nil, it
// will set nothing and immediately return.
func (a *PullActions) SetOpened(v bool) {
	// return if PullActions type is nil
	if a == nil {
		return
	}

	a.Opened = &v
}

// SetSynchronize sets the PullActions Synchronize field.
//
// When the provided PullActions type is nil, it
// will set nothing and immediately return.
func (a *PullActions) SetSynchronize(v bool) {
	// return if PullActions type is nil
	if a == nil {
		return
	}

	a.Synchronize = &v
}

// SetEdited sets the PullActions Edited field.
//
// When the provided PullActions type is nil, it
// will set nothing and immediately return.
func (a *PullActions) SetEdited(v bool) {
	// return if PullActions type is nil
	if a == nil {
		return
	}

	a.Edited = &v
}

// **
// DEPLOY ACTIONS
// **

// FromMask returns the DeployActions type resulting from the provided integer mask.
func (a *DeployActions) FromMask(mask int64) *DeployActions {
	a.SetCreated(mask&constants.AllowDeployCreate > 0)

	return a
}

// ToMask returns the integer mask of the values for the DeployActions set.
func (a *DeployActions) ToMask() int64 {
	mask := int64(0)

	if a.GetCreated() {
		mask = mask | constants.AllowDeployCreate
	}

	return mask
}

// GetCreated returns the Created field from the provided DeployActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *DeployActions) GetCreated() bool {
	// return zero value if DeployActions type or Created field is nil
	if a == nil || a.Created == nil {
		return false
	}

	return *a.Created
}

// SetCreated sets the DeployActions Created field.
//
// When the provided DeployActions type is nil, it
// will set nothing and immediately return.
func (a *DeployActions) SetCreated(v bool) {
	// return if DeployActions type is nil
	if a == nil {
		return
	}

	a.Created = &v
}

// **
// COMMENT ACTIONS
// **

// FromMask returns the CommentActions type resulting from the provided integer mask.
func (a *CommentActions) FromMask(mask int64) *CommentActions {
	a.SetCreated(mask&constants.AllowCommentCreate > 0)
	a.SetEdited(mask&constants.AllowCommentEdit > 0)

	return a
}

// ToMask returns the integer mask of the values for the CommentActions set.
func (a *CommentActions) ToMask() int64 {
	mask := int64(0)

	if a.GetCreated() {
		mask = mask | constants.AllowCommentCreate
	}

	if a.GetEdited() {
		mask = mask | constants.AllowCommentEdit
	}

	return mask
}

// GetCreated returns the Created field from the provided CommentActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *CommentActions) GetCreated() bool {
	// return zero value if Events type or Created field is nil
	if a == nil || a.Created == nil {
		return false
	}

	return *a.Created
}

// GetEdited returns the Edited field from the provided CommentActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *CommentActions) GetEdited() bool {
	// return zero value if Events type or Edited field is nil
	if a == nil || a.Edited == nil {
		return false
	}

	return *a.Edited
}

// SetCreated sets the CommentActions Created field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (a *CommentActions) SetCreated(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Created = &v
}

// SetEdited sets the CommentActions Edited field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (a *CommentActions) SetEdited(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Edited = &v
}
