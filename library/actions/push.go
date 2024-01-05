// SPDX-License-Identifier: Apache-2.0
//
//nolint:dupl // similar code to comment.go
package actions

import "github.com/go-vela/types/constants"

// Push is the library representation of the various actions associated
// with the push event webhook from the SCM.
type Push struct {
	Branch *bool `json:"branch"`
	Tag    *bool `json:"tag"`
}

// FromMask returns the Push type resulting from the provided integer mask.
func (a *Push) FromMask(mask int64) *Push {
	a.SetBranch(mask&constants.AllowPushBranch > 0)
	a.SetTag(mask&constants.AllowPushTag > 0)

	return a
}

// ToMask returns the integer mask of the values for the Push set.
func (a *Push) ToMask() int64 {
	mask := int64(0)

	if a.GetBranch() {
		mask = mask | constants.AllowPushBranch
	}

	if a.GetTag() {
		mask = mask | constants.AllowPushTag
	}

	return mask
}

// GetBranch returns the Branch field from the provided Push. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *Push) GetBranch() bool {
	// return zero value if Push type or Branch field is nil
	if a == nil || a.Branch == nil {
		return false
	}

	return *a.Branch
}

// GetTag returns the Tag field from the provided Push. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *Push) GetTag() bool {
	// return zero value if Push type or Tag field is nil
	if a == nil || a.Tag == nil {
		return false
	}

	return *a.Tag
}

// SetBranch sets the Push Branch field.
//
// When the provided Push type is nil, it
// will set nothing and immediately return.
func (a *Push) SetBranch(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Branch = &v
}

// SetTag sets the Push Tag field.
//
// When the provided Push type is nil, it
// will set nothing and immediately return.
func (a *Push) SetTag(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Tag = &v
}
