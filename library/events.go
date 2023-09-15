// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import "github.com/go-vela/types/constants"

// Events is the library representation of the various events that generate a
// webhook from the SCM.
type Events struct {
	Push        *PushActions    `json:"push"`
	PullRequest *PullActions    `json:"pull_request"`
	Deployment  *DeployActions  `json:"deployment"`
	Comment     *CommentActions `json:"comment"`
}

// NewEventsFromMask is an instatiation function for the Events type that
// takes in an event mask integer value and populates the nested Events struct.
func NewEventsFromMask(mask int64) *Events {
	pushActions := new(PushActions).FromMask(mask)
	pullActions := new(PullActions).FromMask(mask)
	deployActions := new(DeployActions).FromMask(mask)
	commentActions := new(CommentActions).FromMask(mask)

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
func (e *Events) GetPush() *PushActions {
	// return zero value if Events type or Push field is nil
	if e == nil || e.Push == nil {
		return new(PushActions)
	}

	return e.Push
}

// GetPullRequest returns the PullRequest field from the provided Events. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetPullRequest() *PullActions {
	// return zero value if Events type or PullRequest field is nil
	if e == nil || e.PullRequest == nil {
		return new(PullActions)
	}

	return e.PullRequest
}

// GetDeployment returns the Deployment field from the provided Events. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetDeployment() *DeployActions {
	// return zero value if Events type or Deployment field is nil
	if e == nil || e.Deployment == nil {
		return new(DeployActions)
	}

	return e.Deployment
}

// GetComment returns the Comment field from the provided Events. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetComment() *CommentActions {
	// return zero value if Events type or Comment field is nil
	if e == nil || e.Comment == nil {
		return new(CommentActions)
	}

	return e.Comment
}

// SetPush sets the Events Push field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (e *Events) SetPush(v *PushActions) {
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
func (e *Events) SetPullRequest(v *PullActions) {
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
func (e *Events) SetDeployment(v *DeployActions) {
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
func (e *Events) SetComment(v *CommentActions) {
	// return if Events type is nil
	if e == nil {
		return
	}

	e.Comment = v
}
