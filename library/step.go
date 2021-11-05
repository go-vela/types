// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"strconv"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/pipeline"
)

// Step is the library representation of a step in a build.
//
// swagger:model Step
type Step struct {
	ID           *int64  `json:"id,omitempty"`
	BuildID      *int64  `json:"build_id,omitempty"`
	RepoID       *int64  `json:"repo_id,omitempty"`
	Number       *int    `json:"number,omitempty"`
	Name         *string `json:"name,omitempty"`
	Image        *string `json:"image,omitempty"`
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

// Environment returns a list of environment variables
// provided from the fields of the Step type.
func (s *Step) Environment() map[string]string {
	return map[string]string{
		"VELA_STEP_CREATED":      ToString(s.GetCreated()),
		"VELA_STEP_DISTRIBUTION": ToString(s.GetDistribution()),
		"VELA_STEP_EXIT_CODE":    ToString(s.GetExitCode()),
		"VELA_STEP_HOST":         ToString(s.GetHost()),
		"VELA_STEP_IMAGE":        ToString(s.GetImage()),
		"VELA_STEP_NAME":         ToString(s.GetName()),
		"VELA_STEP_NUMBER":       ToString(s.GetNumber()),
		"VELA_STEP_RUNTIME":      ToString(s.GetRuntime()),
		"VELA_STEP_STAGE":        ToString(s.GetStage()),
		"VELA_STEP_STARTED":      ToString(s.GetStarted()),
		"VELA_STEP_STATUS":       ToString(s.GetStatus()),
	}
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

// GetImage returns the Image field.
//
// When the provided Step type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Step) GetImage() string {
	// return zero value if Step type of Image field is nil
	if s == nil || s.Image == nil {
		return ""
	}

	return *s.Image
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

// SetImage sets the Image field.
//
// When the provided Step type is nil, it
// will set nothing and immediately return.
func (s *Step) SetImage(v string) {
	// return if Step type is nil
	if s == nil {
		return
	}

	s.Image = &v
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
	return fmt.Sprintf(`{
  BuildID: %d,
  Created: %d,
  Distribution: %s,
  Error: %s,
  ExitCode: %d,
  Finished: %d,
  Host: %s,
  ID: %d,
  Image: %s,
  Name: %s,
  Number: %d,
  RepoID: %d,
  Runtime: %s,
  Stage: %s,
  Started: %d,
  Status: %s,
}`,
		s.GetBuildID(),
		s.GetCreated(),
		s.GetDistribution(),
		s.GetError(),
		s.GetExitCode(),
		s.GetFinished(),
		s.GetHost(),
		s.GetID(),
		s.GetImage(),
		s.GetName(),
		s.GetNumber(),
		s.GetRepoID(),
		s.GetRuntime(),
		s.GetStage(),
		s.GetStarted(),
		s.GetStatus(),
	)
}

// InitFrom initializes some Step type fields based on a Container and a Build.
func (s *Step) InitFrom(ctn *pipeline.Container, build *Build) {
	if ctn != nil && ctn.Name != "" {
		// set values from the container
		s.SetName(ctn.Name)
		s.SetNumber(ctn.Number)
		s.SetImage(ctn.Image)

		// check if the VELA_STEP_STAGE environment variable exists
		value, ok := ctn.Environment["VELA_STEP_STAGE"]
		if ok {
			// set the Stage field to the value from environment variable
			s.SetStage(value)
		}
	}

	if build != nil {
		// set values from the build
		s.SetHost(build.GetHost())
		s.SetRuntime(build.GetRuntime())
		s.SetDistribution(build.GetDistribution())
	}

	// default status to Pending
	s.SetStatus(constants.StatusPending)
}

// StepFromContainerEnvironment converts the pipeline
// Container type to a library Step type using the container's Environment.
//
// nolint: funlen // ignore function length due to comments and conditionals
func StepFromContainerEnvironment(ctn *pipeline.Container) *Step {
	// check if container or container environment are nil
	if ctn == nil || ctn.Environment == nil {
		return nil
	}

	// create new step type we want to return
	s := new(Step)

	// check if the VELA_STEP_DISTRIBUTION environment variable exists
	value, ok := ctn.Environment["VELA_STEP_DISTRIBUTION"]
	if ok {
		// set the Distribution field to the value from environment variable
		s.SetDistribution(value)
	}

	// check if the VELA_STEP_HOST environment variable exists
	value, ok = ctn.Environment["VELA_STEP_HOST"]
	if ok {
		// set the Host field to the value from environment variable
		s.SetHost(value)
	}

	// check if the VELA_STEP_IMAGE environment variable exists
	value, ok = ctn.Environment["VELA_STEP_IMAGE"]
	if ok {
		// set the Image field to the value from environment variable
		s.SetImage(value)
	}

	// check if the VELA_STEP_NAME environment variable exists
	value, ok = ctn.Environment["VELA_STEP_NAME"]
	if ok {
		// set the Name field to the value from environment variable
		s.SetName(value)
	}

	// check if the VELA_STEP_RUNTIME environment variable exists
	value, ok = ctn.Environment["VELA_STEP_RUNTIME"]
	if ok {
		// set the Runtime field to the value from environment variable
		s.SetRuntime(value)
	}

	// check if the VELA_STEP_STAGE environment variable exists
	value, ok = ctn.Environment["VELA_STEP_STAGE"]
	if ok {
		// set the Stage field to the value from environment variable
		s.SetStage(value)
	}

	// check if the VELA_STEP_STATUS environment variable exists
	value, ok = ctn.Environment["VELA_STEP_STATUS"]
	if ok {
		// set the Status field to the value from environment variable
		s.SetStatus(value)
	}

	// check if the VELA_STEP_CREATED environment variable exists
	value, ok = ctn.Environment["VELA_STEP_CREATED"]
	if ok {
		// parse the environment variable value into an int64
		i, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			// set the Created field to the parsed int64
			s.SetCreated(i)
		}
	}

	// check if the VELA_STEP_EXIT_CODE environment variable exists
	value, ok = ctn.Environment["VELA_STEP_EXIT_CODE"]
	if ok {
		// parse the environment variable value into an int
		i, err := strconv.ParseInt(value, 10, 0)
		if err == nil {
			// set the ExitCode field to the parsed int
			s.SetExitCode(int(i))
		}
	}

	// check if the VELA_STEP_FINISHED environment variable exists
	value, ok = ctn.Environment["VELA_STEP_FINISHED"]
	if ok {
		// parse the environment variable value into an int64
		i, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			// set the Finished field to the parsed int64
			s.SetFinished(i)
		}
	}

	// check if the VELA_STEP_NUMBER environment variable exists
	value, ok = ctn.Environment["VELA_STEP_NUMBER"]
	if ok {
		// parse the environment variable value into an int
		i, err := strconv.ParseInt(value, 10, 0)
		if err == nil {
			// set the Number field to the parsed int
			s.SetNumber(int(i))
		}
	}

	// check if the VELA_STEP_STARTED environment variable exists
	value, ok = ctn.Environment["VELA_STEP_STARTED"]
	if ok {
		// parse the environment variable value into an int64
		i, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			// set the Started field to the parsed int64
			s.SetStarted(i)
		}
	}

	return s
}
