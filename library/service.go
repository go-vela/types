// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import "fmt"

// Service is the library representation of a service in a build.
//
// swagger:model Service
type Service struct {
	ID           *int64  `json:"id,omitempty"`
	BuildID      *int64  `json:"build_id,omitempty"`
	RepoID       *int64  `json:"repo_id,omitempty"`
	Number       *int    `json:"number,omitempty"`
	Name         *string `json:"name,omitempty"`
	Image        *string `json:"image,omitempty"`
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
// provided from the fields of the Service type.
func (s *Service) Environment() map[string]string {
	return map[string]string{
		"VELA_SERVICE_CREATED":      ToString(s.GetCreated()),
		"VELA_SERVICE_DISTRIBUTION": ToString(s.GetDistribution()),
		"VELA_SERVICE_EXIT_CODE":    ToString(s.GetExitCode()),
		"VELA_SERVICE_FINISHED":     ToString(s.GetFinished()),
		"VELA_SERVICE_HOST":         ToString(s.GetHost()),
		"VELA_SERVICE_IMAGE":        ToString(s.GetImage()),
		"VELA_SERVICE_NAME":         ToString(s.GetName()),
		"VELA_SERVICE_NUMBER":       ToString(s.GetNumber()),
		"VELA_SERVICE_RUNTIME":      ToString(s.GetRuntime()),
		"VELA_SERVICE_STARTED":      ToString(s.GetStarted()),
		"VELA_SERVICE_STATUS":       ToString(s.GetStatus()),
	}
}

// GetID returns the ID field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetID() int64 {
	// return zero value if Service type or ID field is nil
	if s == nil || s.ID == nil {
		return 0
	}

	return *s.ID
}

// GetBuildID returns the BuildID field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetBuildID() int64 {
	// return zero value if Service type or BuildID field is nil
	if s == nil || s.BuildID == nil {
		return 0
	}

	return *s.BuildID
}

// GetRepoID returns the RepoID field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetRepoID() int64 {
	// return zero value if Service type or RepoID field is nil
	if s == nil || s.RepoID == nil {
		return 0
	}

	return *s.RepoID
}

// GetNumber returns the Number field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetNumber() int {
	// return zero value if Service type or Number field is nil
	if s == nil || s.Number == nil {
		return 0
	}

	return *s.Number
}

// GetName returns the Name field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetName() string {
	// return zero value if Service type or Name field is nil
	if s == nil || s.Name == nil {
		return ""
	}

	return *s.Name
}

// GetImage returns the Image field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetImage() string {
	// return zero value if Service type or Image field is nil
	if s == nil || s.Image == nil {
		return ""
	}

	return *s.Image
}

// GetStatus returns the Status field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetStatus() string {
	// return zero value if Service type or Status field is nil
	if s == nil || s.Status == nil {
		return ""
	}

	return *s.Status
}

// GetError returns the Error field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetError() string {
	// return zero value if Service type or Error field is nil
	if s == nil || s.Error == nil {
		return ""
	}

	return *s.Error
}

// GetExitCode returns the ExitCode field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetExitCode() int {
	// return zero value if Service type or ExitCode field is nil
	if s == nil || s.ExitCode == nil {
		return 0
	}

	return *s.ExitCode
}

// GetCreated returns the Created field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetCreated() int64 {
	// return zero value if Service type or Created field is nil
	if s == nil || s.Created == nil {
		return 0
	}

	return *s.Created
}

// GetStarted returns the Started field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetStarted() int64 {
	// return zero value if Service type or Started field is nil
	if s == nil || s.Started == nil {
		return 0
	}

	return *s.Started
}

// GetFinished returns the Finished field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetFinished() int64 {
	// return zero value if Service type or Finished field is nil
	if s == nil || s.Finished == nil {
		return 0
	}

	return *s.Finished
}

// GetHost returns the Host field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetHost() string {
	// return zero value if Service type or Host field is nil
	if s == nil || s.Host == nil {
		return ""
	}

	return *s.Host
}

// GetRuntime returns the Runtime field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetRuntime() string {
	// return zero value if Service type or Runtime field is nil
	if s == nil || s.Runtime == nil {
		return ""
	}

	return *s.Runtime
}

// GetDistribution returns the Runtime field.
//
// When the provided Service type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Service) GetDistribution() string {
	// return zero value if Service type or Distribution field is nil
	if s == nil || s.Distribution == nil {
		return ""
	}

	return *s.Distribution
}

// SetID sets the ID field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetID(v int64) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.ID = &v
}

// SetBuildID sets the BuildID field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetBuildID(v int64) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.BuildID = &v
}

// SetRepoID sets the RepoID field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetRepoID(v int64) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.RepoID = &v
}

// SetNumber sets the Number field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetNumber(v int) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.Number = &v
}

// SetName sets the Name field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetName(v string) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.Name = &v
}

// SetImage sets the Image field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetImage(v string) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.Image = &v
}

// SetStatus sets the Status field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetStatus(v string) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.Status = &v
}

// SetError sets the Error field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetError(v string) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.Error = &v
}

// SetExitCode sets the ExitCode field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetExitCode(v int) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.ExitCode = &v
}

// SetCreated sets the Created field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetCreated(v int64) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.Created = &v
}

// SetStarted sets the Started field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetStarted(v int64) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.Started = &v
}

// SetFinished sets the Finished field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetFinished(v int64) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.Finished = &v
}

// SetHost sets the Host field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetHost(v string) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.Host = &v
}

// SetRuntime sets the Runtime field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetRuntime(v string) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.Runtime = &v
}

// SetDistribution sets the Runtime field.
//
// When the provided Service type is nil, it
// will set nothing and immediately return.
func (s *Service) SetDistribution(v string) {
	// return if Service type is nil
	if s == nil {
		return
	}

	s.Distribution = &v
}

// String implements the Stringer interface for the Service type.
func (s *Service) String() string {
	return fmt.Sprintf("%+v", *s)
}
