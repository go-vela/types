// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import "github.com/go-vela/types/constants"

type PRActions struct {
	Opened        *bool `json:"opened"`
	Edited        *bool `json:"edited"`
	Synchronize   *bool `json:"synchronize"`
	Labeled       *bool `json:"labeled"`
	ReviewRequest *bool `json:"review_request"`
}

type CommentActions struct {
	Created *bool `json:"created"`
	Edited  *bool `json:"edited"`
}

type ReviewActions struct {
	Submitted *bool `json:"submitted"`
	Edited    *bool `json:"edited"`
}

type Events struct {
	Push        *bool           `json:"push"`
	PullRequest *PRActions      `json:"pull_request"`
	Tag         *bool           `json:"tag"`
	Deployment  *bool           `json:"deployment"`
	Comment     *CommentActions `json:"comment"`
	Schedule    *bool           `json:"schedule"`
	PullReview  *ReviewActions  `json:"pull_review"`
}

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

func (e *Events) ToDatabase() int64 {
	events := int64(0)

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

// GetPush returns the ScheduledAt field from the provided Schedule. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetPush() bool {
	// return zero value if Schedule type or ScheduledAt field is nil
	if e == nil || e.Push == nil {
		return false
	}

	return *e.Push
}

// GetRepo returns the Repo field from the provided Schedule. If the object is nil,
// or the field within the object is nil, it returns the zero value instead.
func (e *Events) GetPullRequest() *PRActions {
	// return zero value if Schedule type or Repo field is nil
	if e == nil || e.PullRequest == nil {
		return new(PRActions)
	}

	return e.PullRequest
}

func (e *Events) GetTag() bool {
	// return zero value if Schedule type or ScheduledAt field is nil
	if e == nil || e.Tag == nil {
		return false
	}

	return *e.Tag
}

func (e *Events) GetDeployment() bool {
	// return zero value if Schedule type or ScheduledAt field is nil
	if e == nil || e.Deployment == nil {
		return false
	}

	return *e.Deployment
}

func (e *Events) GetComment() *CommentActions {
	// return zero value if Schedule type or Repo field is nil
	if e == nil || e.Comment == nil {
		return new(CommentActions)
	}

	return e.Comment
}

func (e *Events) GetSchedule() bool {
	// return zero value if Schedule type or ScheduledAt field is nil
	if e == nil || e.Schedule == nil {
		return false
	}

	return *e.Schedule
}

func (e *Events) GetPullReview() *ReviewActions {
	// return zero value if Schedule type or Repo field is nil
	if e == nil || e.PullReview == nil {
		return new(ReviewActions)
	}

	return e.PullReview
}

func (a *PRActions) GetOpened() bool {
	// return zero value if Schedule type or ScheduledAt field is nil
	if a == nil || a.Opened == nil {
		return false
	}

	return *a.Opened
}

func (a *PRActions) GetSynchronize() bool {
	// return zero value if Schedule type or ScheduledAt field is nil
	if a == nil || a.Synchronize == nil {
		return false
	}

	return *a.Synchronize
}

func (a *PRActions) GetEdited() bool {
	// return zero value if Schedule type or ScheduledAt field is nil
	if a == nil || a.Edited == nil {
		return false
	}

	return *a.Edited
}

func (a *PRActions) GetLabeled() bool {
	// return zero value if Schedule type or ScheduledAt field is nil
	if a == nil || a.Labeled == nil {
		return false
	}

	return *a.Labeled
}

func (a *PRActions) GetReviewRequest() bool {
	// return zero value if Schedule type or ScheduledAt field is nil
	if a == nil || a.ReviewRequest == nil {
		return false
	}

	return *a.ReviewRequest
}

func (a *CommentActions) GetCreated() bool {
	// return zero value if Schedule type or ScheduledAt field is nil
	if a == nil || a.Created == nil {
		return false
	}

	return *a.Created
}

func (a *CommentActions) GetEdited() bool {
	// return zero value if Schedule type or ScheduledAt field is nil
	if a == nil || a.Edited == nil {
		return false
	}

	return *a.Edited
}

func (a *ReviewActions) GetSubmitted() bool {
	// return zero value if Schedule type or ScheduledAt field is nil
	if a == nil || a.Submitted == nil {
		return false
	}

	return *a.Submitted
}

func (a *ReviewActions) GetEdited() bool {
	// return zero value if Schedule type or ScheduledAt field is nil
	if a == nil || a.Edited == nil {
		return false
	}

	return *a.Edited
}

func (e *Events) SetPush(v bool) {
	// return if Schedule type is nil
	if e == nil {
		return
	}

	e.Push = &v
}

func (e *Events) SetPullRequest(v *PRActions) {
	// return if Schedule type is nil
	if e == nil {
		return
	}

	e.PullRequest = v
}

func (e *Events) SetTag(v bool) {
	// return if Schedule type is nil
	if e == nil {
		return
	}

	e.Tag = &v
}

func (e *Events) SetDeployment(v bool) {
	// return if Schedule type is nil
	if e == nil {
		return
	}

	e.Deployment = &v
}

func (e *Events) SetComment(v *CommentActions) {
	// return if Schedule type is nil
	if e == nil {
		return
	}

	e.Comment = v
}

func (e *Events) SetSchedule(v bool) {
	// return if Schedule type is nil
	if e == nil {
		return
	}

	e.Schedule = &v
}

func (e *Events) SetPullReview(v *ReviewActions) {
	// return if Schedule type is nil
	if e == nil {
		return
	}

	e.PullReview = v
}

func (a *PRActions) SetOpened(v bool) {
	// return if Schedule type is nil
	if a == nil {
		return
	}

	a.Opened = &v
}

func (a *PRActions) SetSynchronize(v bool) {
	// return if Schedule type is nil
	if a == nil {
		return
	}

	a.Synchronize = &v
}

func (a *PRActions) SetEdited(v bool) {
	// return if Schedule type is nil
	if a == nil {
		return
	}

	a.Edited = &v
}

func (a *PRActions) SetLabeled(v bool) {
	// return if Schedule type is nil
	if a == nil {
		return
	}

	a.Labeled = &v
}

func (a *PRActions) SetReviewRequest(v bool) {
	// return if Schedule type is nil
	if a == nil {
		return
	}

	a.ReviewRequest = &v
}

func (a *CommentActions) SetCreated(v bool) {
	// return if Schedule type is nil
	if a == nil {
		return
	}

	a.Created = &v
}

func (a *CommentActions) SetEdited(v bool) {
	// return if Schedule type is nil
	if a == nil {
		return
	}

	a.Edited = &v
}

func (a *ReviewActions) SetSubmitted(v bool) {
	// return if Schedule type is nil
	if a == nil {
		return
	}

	a.Submitted = &v
}

func (a *ReviewActions) SetEdited(v bool) {
	// return if Schedule type is nil
	if a == nil {
		return
	}

	a.Edited = &v
}
