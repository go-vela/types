// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import "fmt"

// Repo is the library representation of a repo.
type Repo struct {
	ID          *int64  `json:"id,omitempty"`
	UserID      *int64  `json:"user_id,omitempty"`
	Org         *string `json:"org,omitempty"`
	Name        *string `json:"name,omitempty"`
	FullName    *string `json:"full_name,omitempty"`
	Link        *string `json:"link,omitempty"`
	Clone       *string `json:"clone,omitempty"`
	Branch      *string `json:"branch,omitempty"`
	Timeout     *int64  `json:"timeout,omitempty"`
	Visibility  *string `json:"visibility,omitempty"`
	Private     *bool   `json:"private,omitempty"`
	Trusted     *bool   `json:"trusted,omitempty"`
	Active      *bool   `json:"active,omitempty"`
	AllowPull   *bool   `json:"allow_pr,omitempty"`
	AllowPush   *bool   `json:"allow_push,omitempty"`
	AllowDeploy *bool   `json:"allow_deploy,omitempty"`
	AllowTag    *bool   `json:"allow_tag,omitempty"`
}

// GetID returns the ID field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetID() int64 {
	// return zero value if Repo type or ID field is nil
	if r == nil || r.ID == nil {
		return 0
	}
	return *r.ID
}

// GetUserID returns the UserID field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetUserID() int64 {
	// return zero value if Repo type or UserID field is nil
	if r == nil || r.UserID == nil {
		return 0
	}
	return *r.UserID
}

// GetOrg returns the Org field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetOrg() string {
	// return zero value if Repo type or Org field is nil
	if r == nil || r.Org == nil {
		return ""
	}
	return *r.Org
}

// GetName returns the Name field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetName() string {
	// return zero value if Repo type or Name field is nil
	if r == nil || r.Name == nil {
		return ""
	}
	return *r.Name
}

// GetFullName returns the FullName field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetFullName() string {
	// return zero value if Repo type or FullName field is nil
	if r == nil || r.FullName == nil {
		return ""
	}
	return *r.FullName
}

// GetLink returns the Link field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetLink() string {
	// return zero value if Repo type or Link field is nil
	if r == nil || r.Link == nil {
		return ""
	}
	return *r.Link
}

// GetClone returns the Clone field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetClone() string {
	// return zero value if Repo type or Clone field is nil
	if r == nil || r.Clone == nil {
		return ""
	}
	return *r.Clone
}

// GetBranch returns the Branch field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetBranch() string {
	// return zero value if Repo type or Branch field is nil
	if r == nil || r.Branch == nil {
		return ""
	}
	return *r.Branch
}

// GetTimeout returns the Timeout field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetTimeout() int64 {
	// return zero value if Repo type or Timeout field is nil
	if r == nil || r.Timeout == nil {
		return 0
	}
	return *r.Timeout
}

// GetVisibility returns the Visibility field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetVisibility() string {
	// return zero value if Repo type or Visibility field is nil
	if r == nil || r.Visibility == nil {
		return ""
	}
	return *r.Visibility
}

// GetPrivate returns the Private field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetPrivate() bool {
	// return zero value if Repo type or Private field is nil
	if r == nil || r.Private == nil {
		return false
	}
	return *r.Private
}

// GetTrusted returns the Trusted field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetTrusted() bool {
	// return zero value if Repo type or Trusted field is nil
	if r == nil || r.Trusted == nil {
		return false
	}
	return *r.Trusted
}

// GetActive returns the Active field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetActive() bool {
	// return zero value if Repo type or Active field is nil
	if r == nil || r.Active == nil {
		return false
	}
	return *r.Active
}

// GetAllowPull returns the AllowPull field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetAllowPull() bool {
	// return zero value if Repo type or AllowPull field is nil
	if r == nil || r.AllowPull == nil {
		return false
	}
	return *r.AllowPull
}

// GetAllowPush returns the AllowPush field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetAllowPush() bool {
	// return zero value if Repo type or AllowPush field is nil
	if r == nil || r.AllowPush == nil {
		return false
	}
	return *r.AllowPush
}

// GetAllowDeploy returns the AllowDeploy field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetAllowDeploy() bool {
	// return zero value if Repo type or AllowDeploy field is nil
	if r == nil || r.AllowDeploy == nil {
		return false
	}
	return *r.AllowDeploy
}

// GetAllowTag returns the AllowTag field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetAllowTag() bool {
	// return zero value if Repo type or AllowTag field is nil
	if r == nil || r.AllowTag == nil {
		return false
	}
	return *r.AllowTag
}

// SetID sets the ID field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetID(v int64) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.ID = &v
}

// SetUserID sets the UserID field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetUserID(v int64) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.UserID = &v
}

// SetOrg sets the Org field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetOrg(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.Org = &v
}

// SetName sets the Name field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetName(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.Name = &v
}

// SetFullName sets the FullName field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetFullName(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.FullName = &v
}

// SetLink sets the Link field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetLink(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.Link = &v
}

// SetClone sets the Clone field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetClone(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.Clone = &v
}

// SetBranch sets the Branch field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetBranch(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.Branch = &v
}

// SetTimeout sets the Timeout field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetTimeout(v int64) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.Timeout = &v
}

// SetVisibility sets the Visibility field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetVisibility(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.Visibility = &v
}

// SetPrivate sets the Private field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetPrivate(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.Private = &v
}

// SetTrusted sets the Trusted field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetTrusted(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.Trusted = &v
}

// SetActive sets the Active field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetActive(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.Active = &v
}

// SetAllowPull sets the AllowPull field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetAllowPull(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.AllowPull = &v
}

// SetAllowPush sets the AllowPush field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetAllowPush(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.AllowPush = &v
}

// SetAllowDeploy sets the AllowDeploy field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetAllowDeploy(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.AllowDeploy = &v
}

// SetAllowTag sets the AllowTag field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetAllowTag(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}
	r.AllowTag = &v
}

// String implements the Stringer interface for the Repo type.
func (r *Repo) String() string {
	return fmt.Sprintf("%+v", *r)
}
