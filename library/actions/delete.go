// SPDX-License-Identifier: Apache-2.0
//
//nolint:dupl // similar code to push.go
package actions

import "github.com/go-vela/types/constants"

// Delete is the library representation of the various actions associated
// with the delete event webhook from the SCM.
type Delete struct {
	Branch *bool `json:"branch"`
	Tag    *bool `json:"tag"`
}

// FromMask returns the Delete type resulting from the provided integer mask.
func (a *Delete) FromMask(mask int64) *Delete {
	a.SetBranch(mask&constants.AllowDeleteBranch > 0)
	a.SetTag(mask&constants.AllowDeleteTag > 0)

	return a
}

// ToMask returns the integer mask of the values for the Delete set.
func (a *Delete) ToMask() int64 {
	mask := int64(0)

	if a.GetBranch() {
		mask = mask | constants.AllowDeleteBranch
	}

	if a.GetTag() {
		mask = mask | constants.AllowDeleteTag
	}

	return mask
}

// GetBranch returns the Branch field from the provided Delete. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *Delete) GetBranch() bool {
	// return zero value if Delete type or Branch field is nil
	if a == nil || a.Branch == nil {
		return false
	}

	return *a.Branch
}

// GetTag returns the Tag field from the provided Delete. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *Delete) GetTag() bool {
	// return zero value if Delete type or Tag field is nil
	if a == nil || a.Tag == nil {
		return false
	}

	return *a.Tag
}

// SetBranch sets the Delete Branch field.
//
// When the provided Delete type is nil, it
// will set nothing and immediately return.
func (a *Delete) SetBranch(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Branch = &v
}

// SetTag sets the Delete Tag field.
//
// When the provided Delete type is nil, it
// will set nothing and immediately return.
func (a *Delete) SetTag(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Tag = &v
}
