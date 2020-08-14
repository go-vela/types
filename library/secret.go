// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"strings"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/pipeline"
)

// Secret is the library representation of a secret.
//
// swagger:model Secret
type Secret struct {
	ID           *int64    `json:"id,omitempty"`
	Org          *string   `json:"org,omitempty"`
	Repo         *string   `json:"repo,omitempty"`
	Team         *string   `json:"team,omitempty"`
	Name         *string   `json:"name,omitempty"`
	Value        *string   `json:"value,omitempty"`
	Type         *string   `json:"type,omitempty"`
	Images       *[]string `json:"images,omitempty"`
	Events       *[]string `json:"events,omitempty"`
	AllowCommand *bool     `json:"allow_command,omitempty"`
}

// Sanitize creates a duplicate of the Secret without the value.
func (s *Secret) Sanitize() *Secret {
	// create a variable since constants can not be addressable
	//
	// https://golang.org/ref/spec#Address_operators
	value := constants.SecretMask

	return &Secret{
		ID:           s.ID,
		Org:          s.Org,
		Repo:         s.Repo,
		Team:         s.Team,
		Name:         s.Name,
		Value:        &value,
		Type:         s.Type,
		Images:       s.Images,
		Events:       s.Events,
		AllowCommand: s.AllowCommand,
	}
}

// Match returns true when the provided container matches
// the conditions to inject a secret into a pipeline container
// resource
func (s *Secret) Match(from *pipeline.Container) bool {
	eACL, iACL := false, false
	events, images, commands := s.GetEvents(), s.GetImages(), s.GetAllowCommand()

	// check if commands are utilized when not allowed
	if !commands && len(from.Commands) > 0 {
		return false
	}

	// check incoming events
	switch from.Environment["BUILD_EVENT"] {
	case constants.EventPush:
		eACL = checkEvent(events, constants.EventPush)
	case constants.EventPull:
		eACL = checkEvent(events, constants.EventPull)
	case constants.EventTag:
		eACL = checkEvent(events, constants.EventTag)
	case constants.EventDeploy:
		eACL = checkEvent(events, constants.EventDeploy)
	case constants.EventComment:
		eACL = checkEvent(events, constants.EventComment)
	}

	// check images whitelist
	for _, i := range images {
		if strings.HasPrefix(from.Image, i) && (len(i) != 0) {
			iACL = true
			break
		}
	}

	// inject secrets into environment
	switch {
	case iACL && (len(events) == 0):
		return true
	case eACL && (len(images) == 0):
		return true
	case eACL && iACL:
		return true
	}

	// return false if not match is found
	return false
}

// GetID returns the ID field.
//
// When the provided Secret type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Secret) GetID() int64 {
	// return zero value if Secret type or ID field is nil
	if s == nil || s.ID == nil {
		return 0
	}

	return *s.ID
}

// GetOrg returns the Org field.
//
// When the provided Secret type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Secret) GetOrg() string {
	// return zero value if Secret type or Org field is nil
	if s == nil || s.Org == nil {
		return ""
	}

	return *s.Org
}

// GetRepo returns the Repo field.
//
// When the provided Secret type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Secret) GetRepo() string {
	// return zero value if Secret type or Repo field is nil
	if s == nil || s.Repo == nil {
		return ""
	}

	return *s.Repo
}

// GetTeam returns the Team field.
//
// When the provided Secret type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Secret) GetTeam() string {
	// return zero value if Secret type or Team field is nil
	if s == nil || s.Team == nil {
		return ""
	}

	return *s.Team
}

// GetName returns the Name field.
//
// When the provided Secret type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Secret) GetName() string {
	// return zero value if Secret type or Name field is nil
	if s == nil || s.Name == nil {
		return ""
	}

	return *s.Name
}

// GetValue returns the Value field.
//
// When the provided Secret type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Secret) GetValue() string {
	// return zero value if Secret type or Value field is nil
	if s == nil || s.Value == nil {
		return ""
	}

	return *s.Value
}

// GetType returns the Type field.
//
// When the provided Secret type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Secret) GetType() string {
	// return zero value if Secret type or Type field is nil
	if s == nil || s.Type == nil {
		return ""
	}

	return *s.Type
}

// GetImages returns the Images field.
//
// When the provided Secret type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Secret) GetImages() []string {
	// return zero value if Secret type or Images field is nil
	if s == nil || s.Images == nil {
		return []string{}
	}

	return *s.Images
}

// GetEvents returns the Events field.
//
// When the provided Secret type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Secret) GetEvents() []string {
	// return zero value if Secret type or Events field is nil
	if s == nil || s.Events == nil {
		return []string{}
	}

	return *s.Events
}

// GetAllowCommand returns the AllowCommand field.
//
// When the provided Secret type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (s *Secret) GetAllowCommand() bool {
	// return zero value if Secret type or Images field is nil
	if s == nil || s.AllowCommand == nil {
		return false
	}

	return *s.AllowCommand
}

// SetID sets the ID field.
//
// When the provided Secret type is nil, it
// will set nothing and immediately return.
func (s *Secret) SetID(v int64) {
	// return if Secret type is nil
	if s == nil {
		return
	}

	s.ID = &v
}

// SetOrg sets the Org field.
//
// When the provided Secret type is nil, it
// will set nothing and immediately return.
func (s *Secret) SetOrg(v string) {
	// return if Secret type is nil
	if s == nil {
		return
	}

	s.Org = &v
}

// SetRepo sets the Repo field.
//
// When the provided Secret type is nil, it
// will set nothing and immediately return.
func (s *Secret) SetRepo(v string) {
	// return if Secret type is nil
	if s == nil {
		return
	}

	s.Repo = &v
}

// SetTeam sets the Team field.
//
// When the provided Secret type is nil, it
// will set nothing and immediately return.
func (s *Secret) SetTeam(v string) {
	// return if Secret type is nil
	if s == nil {
		return
	}

	s.Team = &v
}

// SetName sets the Name field.
//
// When the provided Secret type is nil, it
// will set nothing and immediately return.
func (s *Secret) SetName(v string) {
	// return if Secret type is nil
	if s == nil {
		return
	}

	s.Name = &v
}

// SetValue sets the Value field.
//
// When the provided Secret type is nil, it
// will set nothing and immediately return.
func (s *Secret) SetValue(v string) {
	// return if Secret type is nil
	if s == nil {
		return
	}

	s.Value = &v
}

// SetType sets the Type field.
//
// When the provided Secret type is nil, it
// will set nothing and immediately return.
func (s *Secret) SetType(v string) {
	// return if Secret type is nil
	if s == nil {
		return
	}

	s.Type = &v
}

// SetImages sets the Images field.
//
// When the provided Secret type is nil, it
// will set nothing and immediately return.
func (s *Secret) SetImages(v []string) {
	// return if Secret type is nil
	if s == nil {
		return
	}

	s.Images = &v
}

// SetEvents sets the Events field.
//
// When the provided Secret type is nil, it
// will set nothing and immediately return.
func (s *Secret) SetEvents(v []string) {
	// return if Secret type is nil
	if s == nil {
		return
	}

	s.Events = &v
}

// SetAllowCommand sets the AllowCommand field.
//
// When the provided Secret type is nil, it
// will set nothing and immediately return.
func (s *Secret) SetAllowCommand(v bool) {
	// return if Secret type is nil
	if s == nil {
		return
	}

	s.AllowCommand = &v
}

// String implements the Stringer interface for the Secret type.
func (s *Secret) String() string {
	return fmt.Sprintf(`{
	AllowCommand: %t,
	Events: %s,
	ID: %d,
	Images: %s,
	Name: %s,
	Org: %s,
	Repo: %s,
	Team: %s,
	Type: %s,
	Value: %s,
}`,
		s.GetAllowCommand(),
		s.GetEvents(),
		s.GetID(),
		s.GetImages(),
		s.GetName(),
		s.GetOrg(),
		s.GetRepo(),
		s.GetTeam(),
		s.GetType(),
		s.GetValue(),
	)
}

// checkEvent implements a function that iterates through
// a list to check the event is a member of the list
func checkEvent(events []string, event string) bool {
	for _, e := range events {
		if e == event {
			return true
		}
	}

	return false
}
