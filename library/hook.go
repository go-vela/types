// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import "fmt"

// Hook is the library representation of a webhook for a repo.
//
// swagger:model Webhook
type Hook struct {
	ID       *int64  `json:"id,omitempty"`
	RepoID   *int64  `json:"repo_id,omitempty"`
	BuildID  *int64  `json:"build_id,omitempty"`
	Number   *int    `json:"number,omitempty"`
	SourceID *string `json:"source_id,omitempty"`
	Created  *int64  `json:"created,omitempty"`
	Host     *string `json:"host,omitempty"`
	Event    *string `json:"event,omitempty"`
	Branch   *string `json:"branch,omitempty"`
	Error    *string `json:"error,omitempty"`
	Status   *string `json:"status,omitempty"`
	Link     *string `json:"link,omitempty"`
}

// GetID returns the ID field.
//
// When the provided Hook type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (h *Hook) GetID() int64 {
	// return zero value if Hook type or ID field is nil
	if h == nil || h.ID == nil {
		return 0
	}

	return *h.ID
}

// GetRepoID returns the RepoID field.
//
// When the provided Hook type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (h *Hook) GetRepoID() int64 {
	// return zero value if Hook type or RepoID field is nil
	if h == nil || h.RepoID == nil {
		return 0
	}

	return *h.RepoID
}

// GetBuildID returns the BuildID field.
//
// When the provided Hook type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (h *Hook) GetBuildID() int64 {
	// return zero value if Hook type or BuildID field is nil
	if h == nil || h.BuildID == nil {
		return 0
	}

	return *h.BuildID
}

// GetNumber returns the Number field.
//
// When the provided Hook type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (h *Hook) GetNumber() int {
	// return zero value if Hook type or BuildID field is nil
	if h == nil || h.Number == nil {
		return 0
	}

	return *h.Number
}

// GetSourceID returns the SourceID field.
//
// When the provided Hook type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (h *Hook) GetSourceID() string {
	// return zero value if Hook type or SourceID field is nil
	if h == nil || h.SourceID == nil {
		return ""
	}

	return *h.SourceID
}

// GetCreated returns the Created field.
//
// When the provided Hook type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (h *Hook) GetCreated() int64 {
	// return zero value if Hook type or Created field is nil
	if h == nil || h.Created == nil {
		return 0
	}

	return *h.Created
}

// GetHost returns the Host field.
//
// When the provided Hook type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (h *Hook) GetHost() string {
	// return zero value if Hook type or Host field is nil
	if h == nil || h.Host == nil {
		return ""
	}

	return *h.Host
}

// GetEvent returns the Event field.
//
// When the provided Hook type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (h *Hook) GetEvent() string {
	// return zero value if Hook type or Event field is nil
	if h == nil || h.Event == nil {
		return ""
	}

	return *h.Event
}

// GetBranch returns the Branch field.
//
// When the provided Hook type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (h *Hook) GetBranch() string {
	// return zero value if Hook type or Branch field is nil
	if h == nil || h.Branch == nil {
		return ""
	}

	return *h.Branch
}

// GetError returns the Error field.
//
// When the provided Hook type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (h *Hook) GetError() string {
	// return zero value if Hook type or Error field is nil
	if h == nil || h.Error == nil {
		return ""
	}

	return *h.Error
}

// GetStatus returns the Status field.
//
// When the provided Hook type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (h *Hook) GetStatus() string {
	// return zero value if Hook type or Status field is nil
	if h == nil || h.Status == nil {
		return ""
	}

	return *h.Status
}

// GetLink returns the Link field.
//
// When the provided Hook type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (h *Hook) GetLink() string {
	// return zero value if Hook type or Link field is nil
	if h == nil || h.Link == nil {
		return ""
	}

	return *h.Link
}

// SetID sets the ID field.
//
// When the provided Hook type is nil, it
// will set nothing and immediately return.
func (h *Hook) SetID(v int64) {
	// return if Hook type is nil
	if h == nil {
		return
	}

	h.ID = &v
}

// SetRepoID sets the RepoID field.
//
// When the provided Hook type is nil, it
// will set nothing and immediately return.
func (h *Hook) SetRepoID(v int64) {
	// return if Hook type is nil
	if h == nil {
		return
	}

	h.RepoID = &v
}

// SetBuildID sets the BuildID field.
//
// When the provided Hook type is nil, it
// will set nothing and immediately return.
func (h *Hook) SetBuildID(v int64) {
	// return if Hook type is nil
	if h == nil {
		return
	}

	h.BuildID = &v
}

// SetNumber sets the Number field.
//
// When the provided Hook type is nil, it
// will set nothing and immediately return.
func (h *Hook) SetNumber(v int) {
	// return if Hook type is nil
	if h == nil {
		return
	}

	h.Number = &v
}

// SetSourceID sets the SourceID field.
//
// When the provided Hook type is nil, it
// will set nothing and immediately return.
func (h *Hook) SetSourceID(v string) {
	// return if Hook type is nil
	if h == nil {
		return
	}

	h.SourceID = &v
}

// SetCreated sets the Created field.
//
// When the provided Hook type is nil, it
// will set nothing and immediately return.
func (h *Hook) SetCreated(v int64) {
	// return if Hook type is nil
	if h == nil {
		return
	}

	h.Created = &v
}

// SetHost sets the Host field.
//
// When the provided Hook type is nil, it
// will set nothing and immediately return.
func (h *Hook) SetHost(v string) {
	// return if Hook type is nil
	if h == nil {
		return
	}

	h.Host = &v
}

// SetEvent sets the Event field.
//
// When the provided Hook type is nil, it
// will set nothing and immediately return.
func (h *Hook) SetEvent(v string) {
	// return if Hook type is nil
	if h == nil {
		return
	}
	h.Event = &v
}

// SetBranch sets the Branch field.
//
// When the provided Hook type is nil, it
// will set nothing and immediately return.
func (h *Hook) SetBranch(v string) {
	// return if Hook type is nil
	if h == nil {
		return
	}

	h.Branch = &v
}

// SetError sets the Error field.
//
// When the provided Hook type is nil, it
// will set nothing and immediately return.
func (h *Hook) SetError(v string) {
	// return if Hook type is nil
	if h == nil {
		return
	}

	h.Error = &v
}

// SetStatus sets the Status field.
//
// When the provided Hook type is nil, it
// will set nothing and immediately return.
func (h *Hook) SetStatus(v string) {
	// return if Hook type is nil
	if h == nil {
		return
	}

	h.Status = &v
}

// SetLink sets the Link field.
//
// When the provided Hook type is nil, it
// will set nothing and immediately return.
func (h *Hook) SetLink(v string) {
	// return if Hook type is nil
	if h == nil {
		return
	}

	h.Link = &v
}

// String implements the Stringer interface for the Hook type.
func (h *Hook) String() string {
	return fmt.Sprintf(`{
  Branch: %s,
  BuildID: %d,
  Created: %d,
  Error: %s,
  Event: %s,
  Host: %s,
  ID: %d,
  Link: %s,
  Number: %d,
  RepoID: %d,
  SourceID: %s,
  Status: %s,
}`,
		h.GetBranch(),
		h.GetBuildID(),
		h.GetCreated(),
		h.GetError(),
		h.GetEvent(),
		h.GetHost(),
		h.GetID(),
		h.GetLink(),
		h.GetNumber(),
		h.GetRepoID(),
		h.GetSourceID(),
		h.GetStatus(),
	)
}
