// SPDX-License-Identifier: Apache-2.0

package library

import (
	"fmt"
	"strings"
)

// Repo is the library representation of a repo.
//
// swagger:model Repo
type Repo struct {
	ID               *int64    `json:"id,omitempty"`
	UserID           *int64    `json:"user_id,omitempty"`
	Hash             *string   `json:"-"`
	Org              *string   `json:"org,omitempty"`
	Name             *string   `json:"name,omitempty"`
	FullName         *string   `json:"full_name,omitempty"`
	Link             *string   `json:"link,omitempty"`
	Clone            *string   `json:"clone,omitempty"`
	Branch           *string   `json:"branch,omitempty"`
	Topics           *[]string `json:"topics,omitempty"`
	BuildLimit       *int64    `json:"build_limit,omitempty"`
	Timeout          *int64    `json:"timeout,omitempty"`
	Counter          *int      `json:"counter,omitempty"`
	Visibility       *string   `json:"visibility,omitempty"`
	Private          *bool     `json:"private,omitempty"`
	Trusted          *bool     `json:"trusted,omitempty"`
	Active           *bool     `json:"active,omitempty"`
	AllowPull        *bool     `json:"allow_pull,omitempty"`
	AllowPush        *bool     `json:"allow_push,omitempty"`
	AllowDeploy      *bool     `json:"allow_deploy,omitempty"`
	AllowTag         *bool     `json:"allow_tag,omitempty"`
	AllowComment     *bool     `json:"allow_comment,omitempty"`
	PipelineType     *string   `json:"pipeline_type,omitempty"`
	PreviousName     *string   `json:"previous_name,omitempty"`
	ApproveForkBuild *string   `json:"approve_fork_build,omitempty"`
}

// Environment returns a list of environment variables
// provided from the fields of the Repo type.
func (r *Repo) Environment() map[string]string {
	return map[string]string{
		"VELA_REPO_ACTIVE":        ToString(r.GetActive()),
		"VELA_REPO_ALLOW_COMMENT": ToString(r.GetAllowComment()),
		"VELA_REPO_ALLOW_DEPLOY":  ToString(r.GetAllowDeploy()),
		"VELA_REPO_ALLOW_PULL":    ToString(r.GetAllowPull()),
		"VELA_REPO_ALLOW_PUSH":    ToString(r.GetAllowPush()),
		"VELA_REPO_ALLOW_TAG":     ToString(r.GetAllowTag()),
		"VELA_REPO_BRANCH":        ToString(r.GetBranch()),
		"VELA_REPO_TOPICS":        strings.Join(r.GetTopics()[:], ","),
		"VELA_REPO_BUILD_LIMIT":   ToString(r.GetBuildLimit()),
		"VELA_REPO_CLONE":         ToString(r.GetClone()),
		"VELA_REPO_FULL_NAME":     ToString(r.GetFullName()),
		"VELA_REPO_LINK":          ToString(r.GetLink()),
		"VELA_REPO_NAME":          ToString(r.GetName()),
		"VELA_REPO_ORG":           ToString(r.GetOrg()),
		"VELA_REPO_PRIVATE":       ToString(r.GetPrivate()),
		"VELA_REPO_TIMEOUT":       ToString(r.GetTimeout()),
		"VELA_REPO_TRUSTED":       ToString(r.GetTrusted()),
		"VELA_REPO_VISIBILITY":    ToString(r.GetVisibility()),
		"VELA_REPO_PIPELINE_TYPE": ToString(r.GetPipelineType()),

		// deprecated environment variables
		"REPOSITORY_ACTIVE":        ToString(r.GetActive()),
		"REPOSITORY_ALLOW_COMMENT": ToString(r.GetAllowComment()),
		"REPOSITORY_ALLOW_DEPLOY":  ToString(r.GetAllowDeploy()),
		"REPOSITORY_ALLOW_PULL":    ToString(r.GetAllowPull()),
		"REPOSITORY_ALLOW_PUSH":    ToString(r.GetAllowPush()),
		"REPOSITORY_ALLOW_TAG":     ToString(r.GetAllowTag()),
		"REPOSITORY_BRANCH":        ToString(r.GetBranch()),
		"REPOSITORY_CLONE":         ToString(r.GetClone()),
		"REPOSITORY_FULL_NAME":     ToString(r.GetFullName()),
		"REPOSITORY_LINK":          ToString(r.GetLink()),
		"REPOSITORY_NAME":          ToString(r.GetName()),
		"REPOSITORY_ORG":           ToString(r.GetOrg()),
		"REPOSITORY_PRIVATE":       ToString(r.GetPrivate()),
		"REPOSITORY_TIMEOUT":       ToString(r.GetTimeout()),
		"REPOSITORY_TRUSTED":       ToString(r.GetTrusted()),
		"REPOSITORY_VISIBILITY":    ToString(r.GetVisibility()),
	}
}

// GetID returns the ID field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetID() int64 {
	// return zero value if Repo type or ID field is nil
	if r == nil || r.ID == nil {
		return 0
	}

	return *r.ID
}

// GetUserID returns the UserID field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetUserID() int64 {
	// return zero value if Repo type or UserID field is nil
	if r == nil || r.UserID == nil {
		return 0
	}

	return *r.UserID
}

// GetHash returns the Hash field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetHash() string {
	// return zero value if Repo type or Hash field is nil
	if r == nil || r.Hash == nil {
		return ""
	}

	return *r.Hash
}

// GetOrg returns the Org field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetOrg() string {
	// return zero value if Repo type or Org field is nil
	if r == nil || r.Org == nil {
		return ""
	}

	return *r.Org
}

// GetName returns the Name field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetName() string {
	// return zero value if Repo type or Name field is nil
	if r == nil || r.Name == nil {
		return ""
	}

	return *r.Name
}

// GetFullName returns the FullName field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetFullName() string {
	// return zero value if Repo type or FullName field is nil
	if r == nil || r.FullName == nil {
		return ""
	}

	return *r.FullName
}

// GetLink returns the Link field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetLink() string {
	// return zero value if Repo type or Link field is nil
	if r == nil || r.Link == nil {
		return ""
	}

	return *r.Link
}

// GetClone returns the Clone field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetClone() string {
	// return zero value if Repo type or Clone field is nil
	if r == nil || r.Clone == nil {
		return ""
	}

	return *r.Clone
}

// GetBranch returns the Branch field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetBranch() string {
	// return zero value if Repo type or Branch field is nil
	if r == nil || r.Branch == nil {
		return ""
	}

	return *r.Branch
}

// GetTopics returns the Topics field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetTopics() []string {
	// return zero value if Repo type or Topics field is nil
	if r == nil || r.Topics == nil {
		return []string{}
	}

	return *r.Topics
}

// GetBuildLimit returns the BuildLimit field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetBuildLimit() int64 {
	// return zero value if Repo type or BuildLimit field is nil
	if r == nil || r.BuildLimit == nil {
		return 0
	}

	return *r.BuildLimit
}

// GetTimeout returns the Timeout field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetTimeout() int64 {
	// return zero value if Repo type or Timeout field is nil
	if r == nil || r.Timeout == nil {
		return 0
	}

	return *r.Timeout
}

// GetCounter returns the Counter field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetCounter() int {
	// return zero value if Repo type or Counter field is nil
	if r == nil || r.Counter == nil {
		return 0
	}

	return *r.Counter
}

// GetVisibility returns the Visibility field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetVisibility() string {
	// return zero value if Repo type or Visibility field is nil
	if r == nil || r.Visibility == nil {
		return ""
	}

	return *r.Visibility
}

// GetPrivate returns the Private field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetPrivate() bool {
	// return zero value if Repo type or Private field is nil
	if r == nil || r.Private == nil {
		return false
	}

	return *r.Private
}

// GetTrusted returns the Trusted field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetTrusted() bool {
	// return zero value if Repo type or Trusted field is nil
	if r == nil || r.Trusted == nil {
		return false
	}

	return *r.Trusted
}

// GetActive returns the Active field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetActive() bool {
	// return zero value if Repo type or Active field is nil
	if r == nil || r.Active == nil {
		return false
	}

	return *r.Active
}

// GetAllowPull returns the AllowPull field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetAllowPull() bool {
	// return zero value if Repo type or AllowPull field is nil
	if r == nil || r.AllowPull == nil {
		return false
	}

	return *r.AllowPull
}

// GetAllowPush returns the AllowPush field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetAllowPush() bool {
	// return zero value if Repo type or AllowPush field is nil
	if r == nil || r.AllowPush == nil {
		return false
	}

	return *r.AllowPush
}

// GetAllowDeploy returns the AllowDeploy field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetAllowDeploy() bool {
	// return zero value if Repo type or AllowDeploy field is nil
	if r == nil || r.AllowDeploy == nil {
		return false
	}

	return *r.AllowDeploy
}

// GetAllowTag returns the AllowTag field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetAllowTag() bool {
	// return zero value if Repo type or AllowTag field is nil
	if r == nil || r.AllowTag == nil {
		return false
	}

	return *r.AllowTag
}

// GetAllowComment returns the AllowComment field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetAllowComment() bool {
	// return zero value if Repo type or AllowTag field is nil
	if r == nil || r.AllowComment == nil {
		return false
	}

	return *r.AllowComment
}

// GetPipelineType returns the PipelineType field.
//
// When the provided Repo type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (r *Repo) GetPipelineType() string {
	// return zero value if Repo type or PipelineType field is nil
	if r == nil || r.PipelineType == nil {
		return ""
	}

	return *r.PipelineType
}

// GetPreviousName returns the PreviousName field.
//
// When the provided Repo type is nil, or the field within
//  the type is nil, it returns the zero value for the field.
func (r *Repo) GetPreviousName() string {
	// return zero value if Repo type or PreviousName field is nil
	if r == nil || r.PreviousName == nil {
		return ""
	}

	return *r.PreviousName
}

// GetApproveForkBuild returns the ApproveForkBuild field.
//
// When the provided Repo type is nil, or the field within
//  the type is nil, it returns the zero value for the field.
func (r *Repo) GetApproveForkBuild() string {
	// return zero value if Repo type or ApproveForkBuild field is nil
	if r == nil || r.ApproveForkBuild == nil {
		return ""
	}

	return *r.ApproveForkBuild
}

// SetID sets the ID field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetID(v int64) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.ID = &v
}

// SetUserID sets the UserID field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetUserID(v int64) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.UserID = &v
}

// SetHash sets the Hash field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetHash(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.Hash = &v
}

// SetOrg sets the Org field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetOrg(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.Org = &v
}

// SetName sets the Name field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetName(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.Name = &v
}

// SetFullName sets the FullName field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetFullName(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.FullName = &v
}

// SetLink sets the Link field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetLink(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.Link = &v
}

// SetClone sets the Clone field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetClone(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.Clone = &v
}

// SetBranch sets the Branch field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetBranch(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.Branch = &v
}

// SetTopics sets the Topics field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetTopics(v []string) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.Topics = &v
}

// SetBuildLimit sets the BuildLimit field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetBuildLimit(v int64) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.BuildLimit = &v
}

// SetTimeout sets the Timeout field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetTimeout(v int64) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.Timeout = &v
}

// SetCounter sets the Counter field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetCounter(v int) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.Counter = &v
}

// SetVisibility sets the Visibility field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetVisibility(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.Visibility = &v
}

// SetPrivate sets the Private field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetPrivate(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.Private = &v
}

// SetTrusted sets the Trusted field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetTrusted(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.Trusted = &v
}

// SetActive sets the Active field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetActive(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.Active = &v
}

// SetAllowPull sets the AllowPull field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetAllowPull(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.AllowPull = &v
}

// SetAllowPush sets the AllowPush field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetAllowPush(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.AllowPush = &v
}

// SetAllowDeploy sets the AllowDeploy field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetAllowDeploy(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.AllowDeploy = &v
}

// SetAllowTag sets the AllowTag field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetAllowTag(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.AllowTag = &v
}

// SetAllowComment sets the AllowComment field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetAllowComment(v bool) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.AllowComment = &v
}

// SetPipelineType sets the PipelineType field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetPipelineType(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.PipelineType = &v
}

// SetPreviousName sets the PreviousName field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetPreviousName(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.PreviousName = &v
}

// SetApproveForkBuild sets the ApproveForkBuild field.
//
// When the provided Repo type is nil, it
// will set nothing and immediately return.
func (r *Repo) SetApproveForkBuild(v string) {
	// return if Repo type is nil
	if r == nil {
		return
	}

	r.ApproveForkBuild = &v
}

// String implements the Stringer interface for the Repo type.
func (r *Repo) String() string {
	return fmt.Sprintf(`{
  Active: %t,
  AllowComment: %t,
  AllowDeploy: %t,
  AllowPull: %t,
  AllowPush: %t,
  AllowTag: %t,
  ApproveForkBuild: %s,
  Branch: %s,
  BuildLimit: %d,
  Clone: %s,
  Counter: %d,
  FullName: %s,
  ID: %d,
  Link: %s,
  Name: %s,
  Org: %s,
  PipelineType: %s,
  PreviousName: %s,
  Private: %t,
  Timeout: %d,
  Topics: %s,
  Trusted: %t,
  UserID: %d
  Visibility: %s,
}`,
		r.GetActive(),
		r.GetAllowComment(),
		r.GetAllowDeploy(),
		r.GetAllowPull(),
		r.GetAllowPush(),
		r.GetAllowTag(),
		r.GetApproveForkBuild(),
		r.GetBranch(),
		r.GetBuildLimit(),
		r.GetClone(),
		r.GetCounter(),
		r.GetFullName(),
		r.GetID(),
		r.GetLink(),
		r.GetName(),
		r.GetOrg(),
		r.GetPipelineType(),
		r.GetPreviousName(),
		r.GetPrivate(),
		r.GetTimeout(),
		r.GetTopics(),
		r.GetTrusted(),
		r.GetUserID(),
		r.GetVisibility(),
	)
}
