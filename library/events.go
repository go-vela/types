// SPDX-License-Identifier: Apache-2.0

package library

import (
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library/actions"
)

// Events is the library representation of the various events that generate a
// webhook from the SCM.
type Events struct {
	Push        *actions.Push    `json:"push"`
	PullRequest *actions.Pull    `json:"pull_request"`
	Deployment  *actions.Deploy  `json:"deployment"`
	Comment     *actions.Comment `json:"comment"`
}

// NewEventsFromMask is an instatiation function for the Events type that
// takes in an event mask integer value and populates the nested Events struct.
func NewEventsFromMask(mask int64) *Events {
	pushActions := new(actions.Push).FromMask(mask)
	pullActions := new(actions.Pull).FromMask(mask)
	deployActions := new(actions.Deploy).FromMask(mask)
	commentActions := new(actions.Comment).FromMask(mask)

	e := new(Events)

	e.SetPush(pushActions)
	e.SetPullRequest(pullActions)
	e.SetDeployment(deployActions)
	e.SetComment(commentActions)

	return e
}

// List is an Events method that generates a comma-separated list of event:action
// combinations that are allowed for the repo.
func (e *Events) List() []string {
	eventSlice := []string{}

	if e.GetPush().GetBranch() {
		eventSlice = append(eventSlice, constants.EventPush)
	}

	if e.GetPullRequest().GetOpened() {
		eventSlice = append(eventSlice, constants.EventPull+":"+constants.ActionOpened)
	}

	if e.GetPullRequest().GetSynchronize() {
		eventSlice = append(eventSlice, constants.EventPull+":"+constants.ActionSynchronize)
	}

	if e.GetPullRequest().GetEdited() {
		eventSlice = append(eventSlice, constants.EventPull+":"+constants.ActionEdited)
	}

	if e.GetPush().GetTag() {
		eventSlice = append(eventSlice, constants.EventTag)
	}

	if e.GetDeployment().GetCreated() {
		eventSlice = append(eventSlice, constants.EventDeploy)
	}

	if e.GetComment().GetCreated() {
		eventSlice = append(eventSlice, constants.EventComment+":"+constants.ActionCreated)
	}

	if e.GetComment().GetEdited() {
		eventSlice = append(eventSlice, constants.EventComment+":"+constants.ActionEdited)
	}

	return eventSlice
}

// ToDatabase is an Events method that converts a nested Events struct into an integer event mask.
func (e *Events) ToDatabase() int64 {
	return 0 | e.GetPush().ToMask() | e.GetPullRequest().ToMask() | e.GetComment().ToMask() | e.GetDeployment().ToMask()
}

// GetPush returns the Push field from the provided Events. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetPush() *actions.Push {
	// return zero value if Events type or Push field is nil
	if e == nil || e.Push == nil {
		return new(actions.Push)
	}

	return e.Push
}

// GetPullRequest returns the PullRequest field from the provided Events. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetPullRequest() *actions.Pull {
	// return zero value if Events type or PullRequest field is nil
	if e == nil || e.PullRequest == nil {
		return new(actions.Pull)
	}

	return e.PullRequest
}

// GetDeployment returns the Deployment field from the provided Events. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetDeployment() *actions.Deploy {
	// return zero value if Events type or Deployment field is nil
	if e == nil || e.Deployment == nil {
		return new(actions.Deploy)
	}

	return e.Deployment
}

// GetComment returns the Comment field from the provided Events. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetComment() *actions.Comment {
	// return zero value if Events type or Comment field is nil
	if e == nil || e.Comment == nil {
		return new(actions.Comment)
	}

	return e.Comment
}

// SetPush sets the Events Push field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (e *Events) SetPush(v *actions.Push) {
	// return if Events type is nil
	if e == nil {
		return
	}

	e.Push = v
}

// SetPullRequest sets the Events PullRequest field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (e *Events) SetPullRequest(v *actions.Pull) {
	// return if Events type is nil
	if e == nil {
		return
	}

	e.PullRequest = v
}

// SetDeployment sets the Events Deployment field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (e *Events) SetDeployment(v *actions.Deploy) {
	// return if Events type is nil
	if e == nil {
		return
	}

	e.Deployment = v
}

// SetComment sets the Events Comment field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (e *Events) SetComment(v *actions.Comment) {
	// return if Events type is nil
	if e == nil {
		return
	}

	e.Comment = v
}
