// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import "fmt"

// Step is the library representation of a step in a build.
type Step struct {
	ID           *int64  `json:"id,omitempty"`
	BuildID      *int64  `json:"build_id,omitempty"`
	RepoID       *int64  `json:"repo_id,omitempty"`
	Number       *int    `json:"number,omitempty"`
	Name         *string `json:"name,omitempty"`
	Stage        *string `json:"stage,omitempty"`
	Status       *string `json:"status,omitempty"`
	Error        *string `json:"error,omitempty"`
	ExitCode     *int    `json:"exit_code,omitempty"`
	Created      *int64  `json:"created,omitempty"`
	Started      *int64  `json:"started,omitempty"`
	Finished     *int64  `json:"finished,omitempty"`
	Host         *string `json:"host,omitempty"`
	Runtime      *string `json:"runtime,omitempty"`
	Distribution *string `json:"distribution,omitempty"`
}

// GetID returns the ID field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetID() int64 {
	// return zero value if Step type or ID field is nil
	if s == nil || s.ID == nil {
		return 0
	}
	return *s.ID
}

// GetBuildID returns the BuildID field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetBuildID() int64 {
	// return zero value if Step type or BuildID field is nil
	if s == nil || s.BuildID == nil {
		return 0
	}
	return *s.BuildID
}

// GetRepoID returns the RepoID field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetRepoID() int64 {
	// return zero value if Step type or RepoID field is nil
	if s == nil || s.RepoID == nil {
		return 0
	}
	return *s.RepoID
}

// GetNumber returns the Number field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetNumber() int {
	// return zero value if Step type or Number field is nil
	if s == nil || s.Number == nil {
		return 0
	}
	return *s.Number
}

// GetName returns the Name field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetName() string {
	// return zero value if Step type or Name field is nil
	if s == nil || s.Name == nil {
		return ""
	}
	return *s.Name
}

// GetStage returns the Stage field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetStage() string {
	// return zero value if Step type or Stage field is nil
	if s == nil || s.Stage == nil {
		return ""
	}
	return *s.Stage
}

// GetStatus returns the Status field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetStatus() string {
	// return zero value if Step type or Status field is nil
	if s == nil || s.Status == nil {
		return ""
	}
	return *s.Status
}

// GetError returns the Error field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetError() string {
	// return zero value if Step type or Error field is nil
	if s == nil || s.Error == nil {
		return ""
	}
	return *s.Error
}

// GetExitCode returns the ExitCode field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetExitCode() int {
	// return zero value if Step type or ExitCode field is nil
	if s == nil || s.ExitCode == nil {
		return 0
	}
	return *s.ExitCode
}

// GetCreated returns the Created field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetCreated() int64 {
	// return zero value if Step type or Created field is nil
	if s == nil || s.Created == nil {
		return 0
	}
	return *s.Created
}

// GetStarted returns the Started field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetStarted() int64 {
	// return zero value if Step type or Started field is nil
	if s == nil || s.Started == nil {
		return 0
	}
	return *s.Started
}

// GetFinished returns the Finished field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetFinished() int64 {
	// return zero value if Step type or Finished field is nil
	if s == nil || s.Finished == nil {
		return 0
	}
	return *s.Finished
}

// GetHost returns the Host field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetHost() string {
	// return zero value if Step type or Host field is nil
	if s == nil || s.Host == nil {
		return ""
	}
	return *s.Host
}

// GetRuntime returns the Runtime field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetRuntime() string {
	// return zero value if Step type or Runtime field is nil
	if s == nil || s.Runtime == nil {
		return ""
	}
	return *s.Runtime
}

// GetDistribution returns the Runtime field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetDistribution() string {
	// return zero value if Step type or Distribution field is nil
	if s == nil || s.Distribution == nil {
		return ""
	}
	return *s.Distribution
}

// SetID sets the ID field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetID(v int64) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.ID = &v
}

// SetBuildID sets the BuildID field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetBuildID(v int64) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.BuildID = &v
}

// SetRepoID sets the RepoID field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetRepoID(v int64) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.RepoID = &v
}

// SetNumber sets the Number field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetNumber(v int) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.Number = &v
}

// SetName sets the Name field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetName(v string) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.Name = &v
}

// SetStage sets the Stage field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetStage(v string) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.Stage = &v
}

// SetStatus sets the Status field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetStatus(v string) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.Status = &v
}

// SetError sets the Error field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetError(v string) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.Error = &v
}

// SetExitCode sets the ExitCode field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetExitCode(v int) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.ExitCode = &v
}

// SetCreated sets the Created field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetCreated(v int64) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.Created = &v
}

// SetStarted sets the Started field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetStarted(v int64) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.Started = &v
}

// SetFinished sets the Finished field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetFinished(v int64) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.Finished = &v
}

// SetHost sets the Host field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetHost(v string) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.Host = &v
}

// SetRuntime sets the Runtime field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetRuntime(v string) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.Runtime = &v
}

// SetDistribution sets the Runtime field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetDistribution(v string) {
	// return if Step type is nil
	if s == nil {
		return
	}
	s.Distribution = &v
}

// String implements the Stringer interface for the Step type.
func (s *Step) String() string {
	return fmt.Sprintf("%+v", *s)
}
