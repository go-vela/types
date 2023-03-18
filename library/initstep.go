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
	ID        *int64  `json:"id,omitempty"`
	RepoID    *int64  `json:"repo_id,omitempty"`
	BuildID   *int64  `json:"build_id,omitempty"`
	StepID    *int64  `json:"step_id,omitempty"`
	ServiceID *int64  `json:"service_id,omitempty"`
	Number    *int    `json:"number,omitempty"`
	Reporter  *string `json:"reporter,omitempty"` // which layer created this: compile, runtime, ...
	Name      *string `json:"name,omitempty"`
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

// GetStepID returns the StepID field.
//
// When the provided InitStep type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *InitStep) GetStepID() int64 {
	// return zero value if InitStep type or StepID field is nil
	if i == nil || i.StepID == nil {
		return 0
	}

	return *i.StepID
}

// GetServiceID returns the ServiceID field.
//
// When the provided InitStep type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *InitStep) GetServiceID() int64 {
	// return zero value if InitStep type or ServiceID field is nil
	if i == nil || i.ServiceID == nil {
		return 0
	}

	return *i.ServiceID
}

// GetNumber returns the Number field.
//
// When the provided InitStep type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (i *InitStep) GetNumber() int {
	// return zero value if InitStep type or Number field is nil
	if i == nil || i.Number == nil {
		return 0
	}

	return *i.Number
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

// SetStepID sets the StepID field.
//
// When the provided InitStep type is nil, it
// will set nothing and immediately return.
func (i *InitStep) SetStepID(v int64) {
	// return if InitStep type is nil
	if i == nil {
		return
	}

	i.StepID = &v
}

// SetServiceID sets the ServiceID field.
//
// When the provided InitStep type is nil, it
// will set nothing and immediately return.
func (i *InitStep) SetServiceID(v int64) {
	// return if InitStep type is nil
	if i == nil {
		return
	}

	i.ServiceID = &v
}

// SetNumber sets the Number field.
//
// When the provided InitStep type is nil, it
// will set nothing and immediately return.
func (i *InitStep) SetNumber(v int) {
	// return if InitStep type is nil
	if i == nil {
		return
	}

	i.Number = &v
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

// String implements the Stringer interface for the InitStep type.
func (i *InitStep) String() string {
	return fmt.Sprintf(`{
  ID: %d,
  RepoID: %d,
  BuildID: %d,
  StepID: %d,
  ServiceID: %d,
  Number: %d,
  Reporter: %s,
  Name: %s,
}`,
		i.GetID(),
		i.GetRepoID(),
		i.GetBuildID(),
		i.GetStepID(),
		i.GetServiceID(),
		i.GetNumber(),
		i.GetReporter(),
		i.GetName(),
	)
}

// InitStepFromPipelineInitStep creates a new InitStep based on a pipeline InitStep.
func InitStepFromPipelineInitStep(initStep *pipeline.InitStep) *InitStep {
	// create new InitStep type we want to return
	i := new(InitStep)

	// copy fields from initStep
	if initStep != nil && (initStep.Reporter != "" || initStep.Name != "") {
		// set values from the initStep
		i.SetNumber(initStep.Number)
		i.SetReporter(initStep.Reporter)
		i.SetName(initStep.Name)
	}

	return i
}

// InitStepLogFromBuild creates a new InitStep and Log based on a Build.
func InitStepLogFromBuild(build *Build) (*InitStep, *Log) {
	// create new InitStep type we want to return
	i := new(InitStep)
	l := new(Log)

	l.SetData([]byte{})

	// copy fields from build
	if build != nil {
		// set values from the initStep
		i.SetRepoID(build.GetRepoID())
		l.SetRepoID(build.GetRepoID())
		i.SetBuildID(build.GetID())
		l.SetBuildID(build.GetID())
	}

	return i, l
}

// InitStepLogFromStep creates a new InitStep and Log based on a Step.
func InitStepLogFromStep(step *Step) (*InitStep, *Log) {
	// create new InitStep type we want to return
	i := new(InitStep)
	l := new(Log)

	l.SetData([]byte{})

	// copy fields from step
	if step != nil {
		// set values from the step
		i.SetRepoID(step.GetRepoID())
		l.SetRepoID(step.GetRepoID())
		i.SetBuildID(step.GetBuildID())
		l.SetBuildID(step.GetBuildID())
		// do not set StepID on the log!
		i.SetStepID(step.GetID())
	}

	return i, l
}

// InitStepLogFromService creates a new InitStep and Log based on a Service.
func InitStepLogFromService(service *Service) (*InitStep, *Log) {
	// create new InitStep type we want to return
	i := new(InitStep)
	l := new(Log)

	l.SetData([]byte{})

	// copy fields from service
	if service != nil {
		// set values from the service
		i.SetRepoID(service.GetRepoID())
		l.SetRepoID(service.GetRepoID())
		i.SetBuildID(service.GetBuildID())
		l.SetBuildID(service.GetBuildID())
		// do not set ServiceID on the log!
		i.SetServiceID(service.GetID())
	}

	return i, l
}
