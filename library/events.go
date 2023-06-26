// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import "github.com/go-vela/types/constants"

// PRActions is the library representation of the various actions associated
// with the pull_request event webhook from the SCM.
type PRActions struct {
	Opened        *bool `json:"opened"`
	Edited        *bool `json:"edited"`
	Synchronize   *bool `json:"synchronize"`
	Labeled       *bool `json:"labeled"`
	ReviewRequest *bool `json:"review_request"`
}

// CommentActions is the library representation of the various actions associated
// with the comment event webhook from the SCM.
type CommentActions struct {
	Created *bool `json:"created"`
	Edited  *bool `json:"edited"`
}

// ReviewActions is the library representation of the various actions associated
// with the pull_request_review event webhook from the SCM.
type ReviewActions struct {
	Submitted *bool `json:"submitted"`
	Edited    *bool `json:"edited"`
}

// Events is the library representation of the various events that generate a
// webhook from the SCM.
type Events struct {
	Push        *bool           `json:"push"`
	PullRequest *PRActions      `json:"pull_request"`
	Tag         *bool           `json:"tag"`
	Deployment  *bool           `json:"deployment"`
	Comment     *CommentActions `json:"comment"`
	Schedule    *bool           `json:"schedule"`
	PullReview  *ReviewActions  `json:"pull_review"`
}

// NewEventsFromMask is an instatiation function for the Events type that
// takes in an event mask integer value and populates the nested Events struct.
func NewEventsFromMask(mask int64) *Events {
	prActions := new(PRActions)
	commentActions := new(CommentActions)
	reviewActions := new(ReviewActions)
	e := new(Events)

	prActions.SetOpened(mask&constants.AllowPROpen > 0)
	prActions.SetSynchronize(mask&constants.AllowPRSync > 0)
	prActions.SetEdited(mask&constants.AllowPREdit > 0)
	prActions.SetLabeled(mask&constants.AllowPRLabel > 0)
	prActions.SetReviewRequest(mask&constants.AllowPRReviewRequest > 0)

	commentActions.SetCreated(mask&constants.AllowCommentCreate > 0)
	commentActions.SetEdited(mask&constants.AllowCommentEdit > 0)

	reviewActions.SetSubmitted(mask&constants.AllowReviewSubmit > 0)
	reviewActions.SetEdited(mask&constants.AllowReviewEdit > 0)

	e.SetPush(mask&constants.AllowPush > 0)
	e.SetPullRequest(prActions)
	e.SetTag(mask&constants.AllowTag > 0)
	e.SetDeployment(mask&constants.AllowDeploy > 0)
	e.SetComment(commentActions)
	e.SetPullReview(reviewActions)
	e.SetSchedule(mask&constants.AllowSchedule > 0)

	return e
}

// List is an Events method that generates a comma-separated list of event:action
// combinations that are allowed for the repo.
func (e *Events) List() []string {
	eventSlice := []string{}

	if e.GetPush() {
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

	if e.GetPullRequest().GetLabeled() {
		eventSlice = append(eventSlice, constants.EventPull+":"+constants.ActionLabeled)
	}

	if e.GetPullRequest().GetReviewRequest() {
		eventSlice = append(eventSlice, constants.EventPull+":"+constants.ActionReviewRequested)
	}

	if e.GetTag() {
		eventSlice = append(eventSlice, constants.EventTag)
	}

	if e.GetDeployment() {
		eventSlice = append(eventSlice, constants.EventDeploy)
	}

	if e.GetComment().GetCreated() {
		eventSlice = append(eventSlice, constants.EventComment+":"+constants.ActionCreated)
	}

	if e.GetComment().GetEdited() {
		eventSlice = append(eventSlice, constants.EventComment+":"+constants.ActionEdited)
	}

	if e.GetPullReview().GetSubmitted() {
		eventSlice = append(eventSlice, constants.EventPullReview+":"+constants.ActionSubmitted)
	}

	if e.GetPullReview().GetEdited() {
		eventSlice = append(eventSlice, constants.EventPullReview+":"+constants.ActionEdited)
	}

	if e.GetSchedule() {
		eventSlice = append(eventSlice, constants.EventSchedule)
	}

	return eventSlice
}

// ToDatabase is an Events method that converts a nested Events struct into an integer event mask.
func (e *Events) ToDatabase() int64 {
	events := int64(0)

	// OR operator adds a "1" flag for each event (e.g. 0001 | 0101 == 0101)
	if e.GetPush() {
		events = events | constants.AllowPush
	}

	if e.GetPullRequest().GetOpened() {
		events = events | constants.AllowPROpen
	}

	if e.GetPullRequest().GetSynchronize() {
		events = events | constants.AllowPRSync
	}

	if e.GetPullRequest().GetEdited() {
		events = events | constants.AllowPREdit
	}

	if e.GetPullRequest().GetLabeled() {
		events = events | constants.AllowPRLabel
	}

	if e.GetPullRequest().GetReviewRequest() {
		events = events | constants.AllowPRReviewRequest
	}

	if e.GetTag() {
		events = events | constants.AllowTag
	}

	if e.GetDeployment() {
		events = events | constants.AllowDeploy
	}

	if e.GetComment().GetCreated() {
		events = events | constants.AllowCommentCreate
	}

	if e.GetComment().GetEdited() {
		events = events | constants.AllowCommentEdit
	}

	if e.GetPullReview().GetSubmitted() {
		events = events | constants.AllowReviewSubmit
	}

	if e.GetPullReview().GetEdited() {
		events = events | constants.AllowReviewEdit
	}

	if e.GetSchedule() {
		events = events | constants.AllowSchedule
	}

	return events
}

// GetPush returns the Push field from the provided Events. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetPush() bool {
	// return zero value if Events type or Push field is nil
	if e == nil || e.Push == nil {
		return false
	}

	return *e.Push
}

// GetPullRequest returns the PullRequest field from the provided Events. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetPullRequest() *PRActions {
	// return zero value if Events type or PullRequest field is nil
	if e == nil || e.PullRequest == nil {
		return new(PRActions)
	}

	return e.PullRequest
}

// GetTag returns the Tag field from the provided Events. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetTag() bool {
	// return zero value if Events type or Tag field is nil
	if e == nil || e.Tag == nil {
		return false
	}

	return *e.Tag
}

// GetDeployment returns the Deployment field from the provided Events. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetDeployment() bool {
	// return zero value if Events type or Deployment field is nil
	if e == nil || e.Deployment == nil {
		return false
	}

	return *e.Deployment
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

// GetSchedule returns the Schedule field from the provided Events. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetSchedule() bool {
	// return zero value if Events type or Schedule field is nil
	if e == nil || e.Schedule == nil {
		return false
	}

	return *e.Schedule
}

// GetPullReview returns the PullReview field from the provided Events. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetPullReview() *ReviewActions {
	// return zero value if Events type or PullReview field is nil
	if e == nil || e.PullReview == nil {
		return new(ReviewActions)
	}

	return e.PullReview
}

// GetOpened returns the Opened field from the provided PRActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *PRActions) GetOpened() bool {
	// return zero value if PRActions type or Opened field is nil
	if a == nil || a.Opened == nil {
		return false
	}

	return *a.Opened
}

// GetSynchronize returns the Synchronize field from the provided PRActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *PRActions) GetSynchronize() bool {
	// return zero value if PRActions type or Synchronize field is nil
	if a == nil || a.Synchronize == nil {
		return false
	}

	return *a.Synchronize
}

// GetEdited returns the Edited field from the provided PRActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *PRActions) GetEdited() bool {
	// return zero value if PRActions type or Edited field is nil
	if a == nil || a.Edited == nil {
		return false
	}

	return *a.Edited
}

// GetLabeled returns the Labeled field from the provided PRActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *PRActions) GetLabeled() bool {
	// return zero value if PRActions type or Labeled field is nil
	if a == nil || a.Labeled == nil {
		return false
	}

	return *a.Labeled
}

// GetReviewRequest returns the ReviewRequest field from the provided PRActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *PRActions) GetReviewRequest() bool {
	// return zero value if PRActions type or ReviewRequest field is nil
	if a == nil || a.ReviewRequest == nil {
		return false
	}

	return *a.ReviewRequest
}

// GetCreated returns the Created field from the provided CommentActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *CommentActions) GetCreated() bool {
	// return zero value if Events type or Created field is nil
	if a == nil || a.Created == nil {
		return false
	}

	return *a.Created
}

// GetEdited returns the Edited field from the provided CommentActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *CommentActions) GetEdited() bool {
	// return zero value if Events type or Edited field is nil
	if a == nil || a.Edited == nil {
		return false
	}

	return *a.Edited
}

// GetSubmitted returns the Submitted field from the provided ReviewActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *ReviewActions) GetSubmitted() bool {
	// return zero value if Events type or Submitted field is nil
	if a == nil || a.Submitted == nil {
		return false
	}

	return *a.Submitted
}

// GetEdited returns the Edited field from the provided ReviewActions. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (a *ReviewActions) GetEdited() bool {
	// return zero value if Events type or Edited field is nil
	if a == nil || a.Edited == nil {
		return false
	}

	return *a.Edited
}

// SetPush sets the Events Push field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (e *Events) SetPush(v bool) {
	// return if Events type is nil
	if e == nil {
		return
	}

	e.Push = &v
}

// SetPullRequest sets the Events PullRequest field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (e *Events) SetPullRequest(v *PRActions) {
	// return if Events type is nil
	if e == nil {
		return
	}

	e.PullRequest = v
}

// SetTag sets the Events Tag field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (e *Events) SetTag(v bool) {
	// return if Events type is nil
	if e == nil {
		return
	}

	e.Tag = &v
}

// SetDeployment sets the Events Deployment field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (e *Events) SetDeployment(v bool) {
	// return if Events type is nil
	if e == nil {
		return
	}

	e.Deployment = &v
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

// SetSchedule sets the Events Schedule field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (e *Events) SetSchedule(v bool) {
	// return if Events type is nil
	if e == nil {
		return
	}

	e.Schedule = &v
}

// SetPullReview sets the Events PullReview field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (e *Events) SetPullReview(v *ReviewActions) {
	// return if Events type is nil
	if e == nil {
		return
	}

	e.PullReview = v
}

// SetOpened sets the PRActions Opened field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (a *PRActions) SetOpened(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Opened = &v
}

// SetSynchronize sets the PRActions Synchronize field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (a *PRActions) SetSynchronize(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Synchronize = &v
}

// SetEdited sets the PRActions Edited field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (a *PRActions) SetEdited(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Edited = &v
}

// SetLabeled sets the PRActions Labeled field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (a *PRActions) SetLabeled(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Labeled = &v
}

// SetReviewRequest sets the PRActions ReviewRequest field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (a *PRActions) SetReviewRequest(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.ReviewRequest = &v
}

// SetCreated sets the CommentActions Created field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (a *CommentActions) SetCreated(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Created = &v
}

// SetEdited sets the CommentActions Edited field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (a *CommentActions) SetEdited(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Edited = &v
}

// SetSubmitted sets the ReviewActions Submitted field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (a *ReviewActions) SetSubmitted(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Submitted = &v
}

// SetEdited sets the ReviewActions Edited field.
//
// When the provided Events type is nil, it
// will set nothing and immediately return.
func (a *ReviewActions) SetEdited(v bool) {
	// return if Events type is nil
	if a == nil {
		return
	}

	a.Edited = &v
}
