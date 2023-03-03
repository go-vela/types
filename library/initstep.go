// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"

	"github.com/go-vela/types/pipeline"
)

// InitStep is the library representation of an initStep step in a build.
//
// swagger:model InitStep
type InitStep struct {
	ID       *int64  `json:"id,omitempty"`
	RepoID   *int64  `json:"repo_id,omitempty"`
	BuildID  *int64  `json:"build_id,omitempty"`
	Reporter *string `json:"reporter,omitempty"` // which layer created this: compile, runtime, ...
	Name     *string `json:"name,omitempty"`
	Mimetype *string `json:"mimetype,omitempty"`
}

// GetID returns the ID field.
//
// When the provided InitStep type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *InitStep) GetID() int64 {
	// return zero value if InitStep type or ID field is nil
	if i == nil || i.ID == nil {
		return 0
	}

	return *i.ID
}

// GetRepoID returns the RepoID field.
//
// When the provided InitStep type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *InitStep) GetRepoID() int64 {
	// return zero value if InitStep type or RepoID field is nil
	if i == nil || i.RepoID == nil {
		return 0
	}

	return *i.RepoID
}

// GetBuildID returns the BuildID field.
//
// When the provided InitStep type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *InitStep) GetBuildID() int64 {
	// return zero value if InitStep type or BuildID field is nil
	if i == nil || i.BuildID == nil {
		return 0
	}

	return *i.BuildID
}

// GetReporter returns the Reporter field.
//
// When the provided InitStep type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *InitStep) GetReporter() string {
	// return zero value if InitStep type or Stage field is nil
	if i == nil || i.Reporter == nil {
		return ""
	}

	return *i.Reporter
}

// GetName returns the Name field.
//
// When the provided InitStep type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *InitStep) GetName() string {
	// return zero value if InitStep type or Name field is nil
	if i == nil || i.Name == nil {
		return ""
	}

	return *i.Name
}

// GetMimetype returns the Mimetype field.
//
// When the provided InitStep type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *InitStep) GetMimetype() string {
	// return zero value if InitStep type of Image field is nil
	if i == nil || i.Mimetype == nil {
		return ""
	}

	return *i.Mimetype
}

// SetID sets the ID field.
//
// When the provided InitStep type is nil, it
// will set nothing and immediately return.
func (i *InitStep) SetID(v int64) {
	// return if InitStep type is nil
	if i == nil {
		return
	}

	i.ID = &v
}

// SetRepoID sets the RepoID field.
//
// When the provided InitStep type is nil, it
// will set nothing and immediately return.
func (i *InitStep) SetRepoID(v int64) {
	// return if InitStep type is nil
	if i == nil {
		return
	}

	i.RepoID = &v
}

// SetBuildID sets the BuildID field.
//
// When the provided InitStep type is nil, it
// will set nothing and immediately return.
func (i *InitStep) SetBuildID(v int64) {
	// return if InitStep type is nil
	if i == nil {
		return
	}

	i.BuildID = &v
}

// SetReporter sets the Reporter field.
//
// When the provided InitStep type is nil, it
// will set nothing and immediately return.
func (i *InitStep) SetReporter(v string) {
	// return if InitStep type is nil
	if i == nil {
		return
	}

	i.Reporter = &v
}

// SetName sets the Name field.
//
// When the provided InitStep type is nil, it
// will set nothing and immediately return.
func (i *InitStep) SetName(v string) {
	// return if InitStep type is nil
	if i == nil {
		return
	}

	i.Name = &v
}

// SetMimetype sets the Mimetype field.
//
// When the provided InitStep type is nil, it
// will set nothing and immediately return.
func (i *InitStep) SetMimetype(v string) {
	// return if InitStep type is nil
	if i == nil {
		return
	}

	i.Mimetype = &v
}

// String implements the Stringer interface for the InitStep type.
func (i *InitStep) String() string {
	return fmt.Sprintf(`{
  BuildID: %d,
  ID: %d,
  Mimetype: %s,
  Name: %s,
  RepoID: %d,
  Reporter: %s,
}`,
		i.GetBuildID(),
		i.GetID(),
		i.GetMimetype(),
		i.GetName(),
		i.GetRepoID(),
		i.GetReporter(),
	)
}

// InitStepFromBuildInitStep creates a new InitStep based on a Build and pipeline InitStep.
func InitStepFromBuildInitStep(initStep *pipeline.InitStep) *InitStep {
	// create new InitStep type we want to return
	i := new(InitStep)

	// copy fields from initStep
	if initStep != nil && (initStep.Reporter != "" || initStep.Name != "") {
		// set values from the initStep
		i.SetReporter(initStep.Reporter)
		i.SetName(initStep.Name)
		i.SetMimetype(initStep.Mimetype)
	}

	return i
}
