// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import "fmt"

// Build is the library representation of a build for a pipeline.
type Build struct {
	ID           *int64  `json:"id,omitempty"`
	RepoID       *int64  `json:"repo_id,omitempty"`
	Number       *int    `json:"number,omitempty"`
	Parent       *int    `json:"parent,omitempty"`
	Event        *string `json:"event,omitempty"`
	Status       *string `json:"status,omitempty"`
	Error        *string `json:"error,omitempty"`
	Enqueued     *int64  `json:"enqueued,omitempty"`
	Created      *int64  `json:"created,omitempty"`
	Started      *int64  `json:"started,omitempty"`
	Finished     *int64  `json:"finished,omitempty"`
	Deploy       *string `json:"deploy,omitempty"`
	Clone        *string `json:"clone,omitempty"`
	Source       *string `json:"source,omitempty"`
	Title        *string `json:"title,omitempty"`
	Message      *string `json:"message,omitempty"`
	Commit       *string `json:"commit,omitempty"`
	Sender       *string `json:"sender,omitempty"`
	Author       *string `json:"author,omitempty"`
	Email        *string `json:"email,omitempty"`
	Link         *string `json:"link,omitempty"`
	Branch       *string `json:"branch,omitempty"`
	Ref          *string `json:"ref,omitempty"`
	BaseRef      *string `json:"base_ref,omitempty"`
	Host         *string `json:"host,omitempty"`
	Runtime      *string `json:"runtime,omitempty"`
	Distribution *string `json:"distribution,omitempty"`
}

// GetID returns the ID field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetID() int64 {
	// return zero value if Build type or ID field is nil
	if b == nil || b.ID == nil {
		return 0
	}
	return *b.ID
}

// GetRepoID returns the RepoID field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetRepoID() int64 {
	// return zero value if Build type or RepoID field is nil
	if b == nil || b.RepoID == nil {
		return 0
	}
	return *b.RepoID
}

// GetNumber returns the Number field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetNumber() int {
	// return zero value if Build type or Number field is nil
	if b == nil || b.Number == nil {
		return 0
	}
	return *b.Number
}

// GetParent returns the Parent field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetParent() int {
	// return zero value if Build type or Parent field is nil
	if b == nil || b.Parent == nil {
		return 0
	}
	return *b.Parent
}

// GetEvent returns the Event field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetEvent() string {
	// return zero value if Build type or Event field is nil
	if b == nil || b.Event == nil {
		return ""
	}
	return *b.Event
}

// GetStatus returns the Status field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetStatus() string {
	// return zero value if Build type or Status field is nil
	if b == nil || b.Status == nil {
		return ""
	}
	return *b.Status
}

// GetError returns the Error field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetError() string {
	// return zero value if Build type or Error field is nil
	if b == nil || b.Error == nil {
		return ""
	}
	return *b.Error
}

// GetEnqueued returns the Enqueued field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetEnqueued() int64 {
	// return zero value if Build type or Enqueued field is nil
	if b == nil || b.Enqueued == nil {
		return 0
	}
	return *b.Enqueued
}

// GetCreated returns the Created field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetCreated() int64 {
	// return zero value if Build type or Created field is nil
	if b == nil || b.Created == nil {
		return 0
	}
	return *b.Created
}

// GetStarted returns the Started field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetStarted() int64 {
	// return zero value if Build type or Started field is nil
	if b == nil || b.Started == nil {
		return 0
	}
	return *b.Started
}

// GetFinished returns the Finished field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetFinished() int64 {
	// return zero value if Build type or Finished field is nil
	if b == nil || b.Finished == nil {
		return 0
	}
	return *b.Finished
}

// GetDeploy returns the Deploy field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetDeploy() string {
	// return zero value if Build type or Deploy field is nil
	if b == nil || b.Deploy == nil {
		return ""
	}
	return *b.Deploy
}

// GetClone returns the Clone field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetClone() string {
	// return zero value if Build type or Clone field is nil
	if b == nil || b.Clone == nil {
		return ""
	}
	return *b.Clone
}

// GetSource returns the Source field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetSource() string {
	// return zero value if Build type or Source field is nil
	if b == nil || b.Source == nil {
		return ""
	}
	return *b.Source
}

// GetTitle returns the Title field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetTitle() string {
	// return zero value if Build type or Title field is nil
	if b == nil || b.Title == nil {
		return ""
	}
	return *b.Title
}

// GetMessage returns the Message field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetMessage() string {
	// return zero value if Build type or Message field is nil
	if b == nil || b.Message == nil {
		return ""
	}
	return *b.Message
}

// GetCommit returns the Commit field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetCommit() string {
	// return zero value if Build type or Commit field is nil
	if b == nil || b.Commit == nil {
		return ""
	}
	return *b.Commit
}

// GetSender returns the Sender field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetSender() string {
	// return zero value if Build type or Sender field is nil
	if b == nil || b.Sender == nil {
		return ""
	}
	return *b.Sender
}

// GetAuthor returns the Author field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetAuthor() string {
	// return zero value if Build type or Author field is nil
	if b == nil || b.Author == nil {
		return ""
	}
	return *b.Author
}

// GetEmail returns the Email field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetEmail() string {
	// return zero value if Build type or Email field is nil
	if b == nil || b.Email == nil {
		return ""
	}
	return *b.Email
}

// GetLink returns the Link field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetLink() string {
	// return zero value if Build type or Link field is nil
	if b == nil || b.Link == nil {
		return ""
	}
	return *b.Link
}

// GetBranch returns the Branch field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetBranch() string {
	// return zero value if Build type or Branch field is nil
	if b == nil || b.Branch == nil {
		return ""
	}
	return *b.Branch
}

// GetRef returns the Ref field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetRef() string {
	// return zero value if Build type or Ref field is nil
	if b == nil || b.Ref == nil {
		return ""
	}
	return *b.Ref
}

// GetBaseRef returns the BaseRef field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetBaseRef() string {
	// return zero value if Build type or BaseRef field is nil
	if b == nil || b.BaseRef == nil {
		return ""
	}
	return *b.BaseRef
}

// GetHost returns the Host field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetHost() string {
	// return zero value if Build type or Host field is nil
	if b == nil || b.Host == nil {
		return ""
	}
	return *b.Host
}

// GetRuntime returns the Runtime field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetRuntime() string {
	// return zero value if Build type or Runtime field is nil
	if b == nil || b.Runtime == nil {
		return ""
	}
	return *b.Runtime
}

// GetDistribution returns the Runtime field.
//
// When the provided Build type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *Build) GetDistribution() string {
	// return zero value if Build type or Distribution field is nil
	if b == nil || b.Distribution == nil {
		return ""
	}
	return *b.Distribution
}

// SetID sets the ID field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetID(v int64) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.ID = &v
}

// SetRepoID sets the RepoID field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetRepoID(v int64) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.RepoID = &v
}

// SetNumber sets the Number field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetNumber(v int) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Number = &v
}

// SetParent sets the Parent field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetParent(v int) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Parent = &v
}

// SetEvent sets the Event field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetEvent(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Event = &v
}

// SetStatus sets the Status field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetStatus(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Status = &v
}

// SetError sets the Error field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetError(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Error = &v
}

// SetEnqueued sets the Enqueued field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetEnqueued(v int64) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Enqueued = &v
}

// SetCreated sets the Created field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetCreated(v int64) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Created = &v
}

// SetStarted sets the Started field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetStarted(v int64) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Started = &v
}

// SetFinished sets the Finished field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetFinished(v int64) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Finished = &v
}

// SetDeploy sets the Deploy field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetDeploy(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Deploy = &v
}

// SetClone sets the Clone field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetClone(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Clone = &v
}

// SetSource sets the Source field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetSource(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Source = &v
}

// SetTitle sets the Title field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetTitle(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Title = &v
}

// SetMessage sets the Message field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetMessage(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Message = &v
}

// SetCommit sets the Commit field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetCommit(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Commit = &v
}

// SetSender sets the Sender field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetSender(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Sender = &v
}

// SetAuthor sets the Author field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetAuthor(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Author = &v
}

// SetEmail sets the Email field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetEmail(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Email = &v
}

// SetLink sets the Link field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetLink(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Link = &v
}

// SetBranch sets the Branch field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetBranch(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Branch = &v
}

// SetRef sets the Ref field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetRef(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Ref = &v
}

// SetBaseRef sets the BaseRef field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetBaseRef(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.BaseRef = &v
}

// SetHost sets the Host field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetHost(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Host = &v
}

// SetRuntime sets the Runtime field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetRuntime(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Runtime = &v
}

// SetDistribution sets the Runtime field.
//
// When the provided Build type is nil, it
// will set nothing and immediately return.
func (b *Build) SetDistribution(v string) {
	// return if Build type is nil
	if b == nil {
		return
	}
	b.Distribution = &v
}

// String implements the Stringer interface for the Build type.
func (b *Build) String() string {
	return fmt.Sprintf("%+v", *b)
}
