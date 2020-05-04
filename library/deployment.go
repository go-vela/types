// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import "fmt"

// Deployment is the library representation of a deployment.
type Deployment struct {
	ID          *int64  `json:"id,omitempty"`
	RepoID      *int64  `json:"repo_id,omitempty"`
	URL         *string `json:"url,omitempty"`
	User        *string `json:"user,omitempty"`
	Commit      *string `json:"commit,omitempty"`
	Ref         *string `json:"ref,omitempty"`
	Task        *string `json:"task,omitempty"`
	Target      *string `json:"target,omitempty"`
	Description *string `json:"description,omitempty"`
}

// GetID returns the ID field.
//
// When the provided Deployment type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (d *Deployment) GetID() int64 {
	// return zero value if Deployment type or ID field is nil
	if d == nil || d.ID == nil {
		return 0
	}

	return *d.ID
}

// GetRepoID returns the RepoID field.
//
// When the provided Deployment type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (d *Deployment) GetRepoID() int64 {
	// return zero value if Deployment type or RepoID field is nil
	if d == nil || d.RepoID == nil {
		return 0
	}

	return *d.RepoID
}

// GetURL returns the URL field.
//
// When the provided Deployment type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (d *Deployment) GetURL() string {
	// return zero value if Deployment type or URL field is nil
	if d == nil || d.URL == nil {
		return ""
	}

	return *d.URL
}

// GetUser returns the User field.
//
// When the provided Deployment type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (d *Deployment) GetUser() string {
	// return zero value if Deployment type or User field is nil
	if d == nil || d.User == nil {
		return ""
	}

	return *d.User
}

// GetCommit returns the Commit field.
//
// When the provided Deployment type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (d *Deployment) GetCommit() string {
	// return zero value if Deployment type or Commit field is nil
	if d == nil || d.Commit == nil {
		return ""
	}

	return *d.Commit
}

// GetRef returns the Ref field.
//
// When the provided Deployment type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (d *Deployment) GetRef() string {
	// return zero value if Deployment type or Ref field is nil
	if d == nil || d.Ref == nil {
		return ""
	}

	return *d.Ref
}

// GetTask returns the Task field.
//
// When the provided Deployment type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (d *Deployment) GetTask() string {
	// return zero value if Deployment type or Task field is nil
	if d == nil || d.Task == nil {
		return ""
	}

	return *d.Task
}

// GetTarget returns the Target field.
//
// When the provided Deployment type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (d *Deployment) GetTarget() string {
	// return zero value if Deployment type or Target field is nil
	if d == nil || d.Target == nil {
		return ""
	}

	return *d.Target
}

// GetDescription returns the Description field.
//
// When the provided Deployment type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (d *Deployment) GetDescription() string {
	// return zero value if Deployment type or Description field is nil
	if d == nil || d.Description == nil {
		return ""
	}

	return *d.Description
}

// SetID sets the ID field.
//
// When the provided Deployment type is nil, it
// will set nothing and immediately return.
func (d *Deployment) SetID(v int64) {
	// return if Deployment type is nil
	if d == nil {
		return
	}

	d.ID = &v
}

// SetRepoID sets the RepoID field.
//
// When the provided Deployment type is nil, it
// will set nothing and immediately return.
func (d *Deployment) SetRepoID(v int64) {
	// return if Deployment type is nil
	if d == nil {
		return
	}

	d.RepoID = &v
}

// SetURL sets the URL field.
//
// When the provided Deployment type is nil, it
// will set nothing and immediately return.
func (d *Deployment) SetURL(v string) {
	// return if Deployment type is nil
	if d == nil {
		return
	}

	d.URL = &v
}

// SetUser sets the User field.
//
// When the provided Deployment type is nil, it
// will set nothing and immediately return.
func (d *Deployment) SetUser(v string) {
	// return if Deployment type is nil
	if d == nil {
		return
	}

	d.User = &v
}

// SetCommit sets the Commit field.
//
// When the provided Deployment type is nil, it
// will set nothing and immediately return.
func (d *Deployment) SetCommit(v string) {
	// return if Deployment type is nil
	if d == nil {
		return
	}

	d.Commit = &v
}

// SetRef sets the Ref field.
//
// When the provided Deployment type is nil, it
// will set nothing and immediately return.
func (d *Deployment) SetRef(v string) {
	// return if Deployment type is nil
	if d == nil {
		return
	}

	d.Ref = &v
}

// SetTask sets the Task field.
//
// When the provided Deployment type is nil, it
// will set nothing and immediately return.
func (d *Deployment) SetTask(v string) {
	// return if Deployment type is nil
	if d == nil {
		return
	}

	d.Task = &v
}

// SetTarget sets the Target field.
//
// When the provided Deployment type is nil, it
// will set nothing and immediately return.
func (d *Deployment) SetTarget(v string) {
	// return if Deployment type is nil
	if d == nil {
		return
	}

	d.Target = &v
}

// SetDescription sets the Description field.
//
// When the provided Deployment type is nil, it
// will set nothing and immediately return.
func (d *Deployment) SetDescription(v string) {
	// return if Deployment type is nil
	if d == nil {
		return
	}

	d.Description = &v
}

// String implements the Stringer interface for the Deployment type.
func (d *Deployment) String() string {
	return fmt.Sprintf("%+v", *d)
}
