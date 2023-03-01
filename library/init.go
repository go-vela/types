// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"github.com/go-vela/types/pipeline"
)

// Init is the library representation of an init report in a build.
//
// swagger:model Init
type Init struct {
	ID       *int64  `json:"id,omitempty"`
	RepoID   *int64  `json:"repo_id,omitempty"`
	BuildID  *int64  `json:"build_id,omitempty"`
	Number   *int    `json:"number,omitempty"`
	Reporter *string `json:"reporter,omitempty"` // which layer created this: compile, runtime, ...
	Name     *string `json:"name,omitempty"`
	Mimetype *string `json:"mimetype,omitempty"`
}

// GetID returns the ID field.
//
// When the provided Init type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *Init) GetID() int64 {
	// return zero value if Init type or ID field is nil
	if i == nil || i.ID == nil {
		return 0
	}

	return *i.ID
}

// GetRepoID returns the RepoID field.
//
// When the provided Init type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *Init) GetRepoID() int64 {
	// return zero value if Init type or RepoID field is nil
	if i == nil || i.RepoID == nil {
		return 0
	}

	return *i.RepoID
}

// GetBuildID returns the BuildID field.
//
// When the provided Init type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *Init) GetBuildID() int64 {
	// return zero value if Init type or BuildID field is nil
	if i == nil || i.BuildID == nil {
		return 0
	}

	return *i.BuildID
}

// GetNumber returns the Number field.
//
// When the provided Init type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *Init) GetNumber() int {
	// return zero value if Init type or Number field is nil
	if i == nil || i.Number == nil {
		return 0
	}

	return *i.Number
}

// GetReporter returns the Reporter field.
//
// When the provided Init type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *Init) GetReporter() string {
	// return zero value if Init type or Stage field is nil
	if i == nil || i.Reporter == nil {
		return ""
	}

	return *i.Reporter
}

// GetName returns the Name field.
//
// When the provided Init type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *Init) GetName() string {
	// return zero value if Init type or Name field is nil
	if i == nil || i.Name == nil {
		return ""
	}

	return *i.Name
}

// GetMimetype returns the Mimetype field.
//
// When the provided Init type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *Init) GetMimetype() string {
	// return zero value if Init type of Image field is nil
	if i == nil || i.Mimetype == nil {
		return ""
	}

	return *i.Mimetype
}

// SetID sets the ID field.
//
// When the provided Init type is nil, it
// will set nothing and immediately return.
func (i *Init) SetID(v int64) {
	// return if Init type is nil
	if i == nil {
		return
	}

	i.ID = &v
}

// SetRepoID sets the RepoID field.
//
// When the provided Init type is nil, it
// will set nothing and immediately return.
func (i *Init) SetRepoID(v int64) {
	// return if Init type is nil
	if i == nil {
		return
	}

	i.RepoID = &v
}

// SetBuildID sets the BuildID field.
//
// When the provided Init type is nil, it
// will set nothing and immediately return.
func (i *Init) SetBuildID(v int64) {
	// return if Init type is nil
	if i == nil {
		return
	}

	i.BuildID = &v
}

// SetNumber sets the Number field.
//
// When the provided Init type is nil, it
// will set nothing and immediately return.
func (i *Init) SetNumber(v int) {
	// return if Init type is nil
	if i == nil {
		return
	}

	i.Number = &v
}

// SetReporter sets the Reporter field.
//
// When the provided Init type is nil, it
// will set nothing and immediately return.
func (i *Init) SetReporter(v string) {
	// return if Init type is nil
	if i == nil {
		return
	}

	i.Reporter = &v
}

// SetName sets the Name field.
//
// When the provided Init type is nil, it
// will set nothing and immediately return.
func (i *Init) SetName(v string) {
	// return if Init type is nil
	if i == nil {
		return
	}

	i.Name = &v
}

// SetMimetype sets the Mimetype field.
//
// When the provided Init type is nil, it
// will set nothing and immediately return.
func (i *Init) SetMimetype(v string) {
	// return if Init type is nil
	if i == nil {
		return
	}

	i.Mimetype = &v
}

// String implements the Stringer interface for the Init type.
func (i *Init) String() string {
	return fmt.Sprintf(`{
  BuildID: %d,
  ID: %d,
  Mimetype: %s,
  Name: %s,
  Number: %d,
  RepoID: %d,
  Reporter: %s,
}`,
		i.GetBuildID(),
		i.GetID(),
		i.GetMimetype(),
		i.GetName(),
		i.GetNumber(),
		i.GetRepoID(),
		i.GetReporter(),
	)
}

// InitFromBuildInit creates a new Init based on a Build and pipeline Init.
func InitFromBuildInit(init *pipeline.Init) *Init {
	// create new step type we want to return
	i := new(Init)

	// copy fields from init
	if init != nil && init.Name != "" {
		// set values from the init
		i.SetNumber(init.Number)
		i.SetReporter(init.Reporter)
		i.SetName(init.Name)
		i.SetMimetype(init.Mimetype)
	}

	return i
}
