// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_Repo_Environment(t *testing.T) {
	// setup types
	want := map[string]string{
		"VELA_REPO_ACTIVE":         "true",
		"VELA_REPO_ALLOW_COMMENT":  "false",
		"VELA_REPO_ALLOW_DEPLOY":   "false",
		"VELA_REPO_ALLOW_PULL":     "false",
		"VELA_REPO_ALLOW_PUSH":     "true",
		"VELA_REPO_ALLOW_TAG":      "false",
		"VELA_REPO_ALLOW_RELEASE":  "false",
		"VELA_REPO_BRANCH":         "master",
		"VELA_REPO_BUILD_LIMIT":    "10",
		"VELA_REPO_CLONE":          "https://github.com/github/octocat.git",
		"VELA_REPO_FULL_NAME":      "github/octocat",
		"VELA_REPO_LINK":           "https://github.com/github/octocat",
		"VELA_REPO_NAME":           "octocat",
		"VELA_REPO_ORG":            "github",
		"VELA_REPO_PRIVATE":        "false",
		"VELA_REPO_TIMEOUT":        "30",
		"VELA_REPO_TRUSTED":        "false",
		"VELA_REPO_VISIBILITY":     "public",
		"VELA_REPO_PIPELINE_TYPE":  "",
		"REPOSITORY_ACTIVE":        "true",
		"REPOSITORY_ALLOW_COMMENT": "false",
		"REPOSITORY_ALLOW_DEPLOY":  "false",
		"REPOSITORY_ALLOW_PULL":    "false",
		"REPOSITORY_ALLOW_PUSH":    "true",
		"REPOSITORY_ALLOW_TAG":     "false",
		"REPOSITORY_ALLOW_RELEASE": "false",
		"REPOSITORY_BRANCH":        "master",
		"REPOSITORY_CLONE":         "https://github.com/github/octocat.git",
		"REPOSITORY_FULL_NAME":     "github/octocat",
		"REPOSITORY_LINK":          "https://github.com/github/octocat",
		"REPOSITORY_NAME":          "octocat",
		"REPOSITORY_ORG":           "github",
		"REPOSITORY_PRIVATE":       "false",
		"REPOSITORY_TIMEOUT":       "30",
		"REPOSITORY_TRUSTED":       "false",
		"REPOSITORY_VISIBILITY":    "public",
	}

	// run test
	got := testRepo().Environment()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Environment is %v, want %v", got, want)
	}
}

func TestLibrary_Repo_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		repo *Repo
		want *Repo
	}{
		{
			repo: testRepo(),
			want: testRepo(),
		},
		{
			repo: new(Repo),
			want: new(Repo),
		},
	}

	// run tests
	for _, test := range tests {
		if test.repo.GetID() != test.want.GetID() {
			t.Errorf("GetID is %v, want %v", test.repo.GetID(), test.want.GetID())
		}

		if test.repo.GetUserID() != test.want.GetUserID() {
			t.Errorf("GetUserID is %v, want %v", test.repo.GetUserID(), test.want.GetUserID())
		}

		if test.repo.GetHash() != test.want.GetHash() {
			t.Errorf("GetHash is %v, want %v", test.repo.GetHash(), test.want.GetHash())
		}

		if test.repo.GetOrg() != test.want.GetOrg() {
			t.Errorf("GetOrg is %v, want %v", test.repo.GetOrg(), test.want.GetOrg())
		}

		if test.repo.GetName() != test.want.GetName() {
			t.Errorf("GetName is %v, want %v", test.repo.GetName(), test.want.GetName())
		}

		if test.repo.GetFullName() != test.want.GetFullName() {
			t.Errorf("GetFullName is %v, want %v", test.repo.GetFullName(), test.want.GetFullName())
		}

		if test.repo.GetLink() != test.want.GetLink() {
			t.Errorf("GetLink is %v, want %v", test.repo.GetLink(), test.want.GetLink())
		}

		if test.repo.GetClone() != test.want.GetClone() {
			t.Errorf("GetClone is %v, want %v", test.repo.GetClone(), test.want.GetClone())
		}

		if test.repo.GetBranch() != test.want.GetBranch() {
			t.Errorf("GetBranch is %v, want %v", test.repo.GetBranch(), test.want.GetBranch())
		}

		if test.repo.GetBuildLimit() != test.want.GetBuildLimit() {
			t.Errorf("GetBuildLimit is %v, want %v", test.repo.GetBuildLimit(), test.want.GetBuildLimit())
		}

		if test.repo.GetTimeout() != test.want.GetTimeout() {
			t.Errorf("GetTimeout is %v, want %v", test.repo.GetTimeout(), test.want.GetTimeout())
		}

		if test.repo.GetVisibility() != test.want.GetVisibility() {
			t.Errorf("GetVisibility is %v, want %v", test.repo.GetVisibility(), test.want.GetVisibility())
		}

		if test.repo.GetPrivate() != test.want.GetPrivate() {
			t.Errorf("GetPrivate is %v, want %v", test.repo.GetPrivate(), test.want.GetPrivate())
		}

		if test.repo.GetTrusted() != test.want.GetTrusted() {
			t.Errorf("GetTrusted is %v, want %v", test.repo.GetTrusted(), test.want.GetTrusted())
		}

		if test.repo.GetActive() != test.want.GetActive() {
			t.Errorf("GetActive is %v, want %v", test.repo.GetActive(), test.want.GetActive())
		}

		if test.repo.GetAllowPull() != test.want.GetAllowPull() {
			t.Errorf("GetAllowPull is %v, want %v", test.repo.GetAllowPull(), test.want.GetAllowPull())
		}

		if test.repo.GetAllowPush() != test.want.GetAllowPush() {
			t.Errorf("GetAllowPush is %v, want %v", test.repo.GetAllowPush(), test.want.GetAllowPush())
		}

		if test.repo.GetAllowDeploy() != test.want.GetAllowDeploy() {
			t.Errorf("GetAllowDeploy is %v, want %v", test.repo.GetAllowDeploy(), test.want.GetAllowDeploy())
		}

		if test.repo.GetAllowTag() != test.want.GetAllowTag() {
			t.Errorf("GetAllowTag is %v, want %v", test.repo.GetAllowTag(), test.want.GetAllowTag())
		}

		if test.repo.GetAllowRelease() != test.want.GetAllowRelease() {
			t.Errorf("GetAllowRelease is %v, want %v", test.repo.GetAllowRelease(), test.want.GetAllowRelease())
		}

		if test.repo.GetAllowComment() != test.want.GetAllowComment() {
			t.Errorf("GetAllowComment is %v, want %v", test.repo.GetAllowComment(), test.want.GetAllowComment())
		}

		if test.repo.GetPipelineType() != test.want.GetPipelineType() {
			t.Errorf("GetPipelineType is %v, want %v", test.repo.GetPipelineType(), test.want.GetPipelineType())
		}

		if !reflect.DeepEqual(test.repo.GetPreviousName(), test.want.GetPreviousName()) {
			t.Errorf("GetPreviousName is %v, want %v", test.repo.GetPreviousName(), test.want.GetPreviousName())
		}
	}
}

func TestLibrary_Repo_Setters(t *testing.T) {
	// setup types
	var r *Repo

	// setup tests
	tests := []struct {
		repo *Repo
		want *Repo
	}{
		{
			repo: testRepo(),
			want: testRepo(),
		},
		{
			repo: r,
			want: new(Repo),
		},
	}

	// run tests
	for _, test := range tests {
		test.repo.SetID(test.want.GetID())
		test.repo.SetUserID(test.want.GetUserID())
		test.repo.SetHash(test.want.GetHash())
		test.repo.SetOrg(test.want.GetOrg())
		test.repo.SetName(test.want.GetName())
		test.repo.SetFullName(test.want.GetFullName())
		test.repo.SetLink(test.want.GetLink())
		test.repo.SetClone(test.want.GetClone())
		test.repo.SetBranch(test.want.GetBranch())
		test.repo.SetBuildLimit(test.want.GetBuildLimit())
		test.repo.SetTimeout(test.want.GetTimeout())
		test.repo.SetCounter(test.want.GetCounter())
		test.repo.SetVisibility(test.want.GetVisibility())
		test.repo.SetPrivate(test.want.GetPrivate())
		test.repo.SetTrusted(test.want.GetTrusted())
		test.repo.SetActive(test.want.GetActive())
		test.repo.SetAllowPull(test.want.GetAllowPull())
		test.repo.SetAllowPush(test.want.GetAllowPush())
		test.repo.SetAllowDeploy(test.want.GetAllowDeploy())
		test.repo.SetAllowTag(test.want.GetAllowTag())
		test.repo.SetAllowRelease(test.want.GetAllowRelease())
		test.repo.SetAllowComment(test.want.GetAllowComment())
		test.repo.SetPipelineType(test.want.GetPipelineType())
		test.repo.SetPreviousName(test.want.GetPreviousName())

		if test.repo.GetID() != test.want.GetID() {
			t.Errorf("SetID is %v, want %v", test.repo.GetID(), test.want.GetID())
		}

		if test.repo.GetUserID() != test.want.GetUserID() {
			t.Errorf("SetUserID is %v, want %v", test.repo.GetUserID(), test.want.GetUserID())
		}

		if test.repo.GetHash() != test.want.GetHash() {
			t.Errorf("SetHash is %v, want %v", test.repo.GetHash(), test.want.GetHash())
		}

		if test.repo.GetOrg() != test.want.GetOrg() {
			t.Errorf("SetOrg is %v, want %v", test.repo.GetOrg(), test.want.GetOrg())
		}

		if test.repo.GetName() != test.want.GetName() {
			t.Errorf("SetName is %v, want %v", test.repo.GetName(), test.want.GetName())
		}

		if test.repo.GetFullName() != test.want.GetFullName() {
			t.Errorf("SetFullName is %v, want %v", test.repo.GetFullName(), test.want.GetFullName())
		}

		if test.repo.GetLink() != test.want.GetLink() {
			t.Errorf("SetLink is %v, want %v", test.repo.GetLink(), test.want.GetLink())
		}

		if test.repo.GetClone() != test.want.GetClone() {
			t.Errorf("SetClone is %v, want %v", test.repo.GetClone(), test.want.GetClone())
		}

		if test.repo.GetBranch() != test.want.GetBranch() {
			t.Errorf("SetBranch is %v, want %v", test.repo.GetBranch(), test.want.GetBranch())
		}

		if test.repo.GetBuildLimit() != test.want.GetBuildLimit() {
			t.Errorf("SetBuildLimit is %v, want %v", test.repo.GetBuildLimit(), test.want.GetBuildLimit())
		}

		if test.repo.GetTimeout() != test.want.GetTimeout() {
			t.Errorf("SetTimeout is %v, want %v", test.repo.GetTimeout(), test.want.GetTimeout())
		}

		if test.repo.GetVisibility() != test.want.GetVisibility() {
			t.Errorf("SetVisibility is %v, want %v", test.repo.GetVisibility(), test.want.GetVisibility())
		}

		if test.repo.GetPrivate() != test.want.GetPrivate() {
			t.Errorf("SetPrivate is %v, want %v", test.repo.GetPrivate(), test.want.GetPrivate())
		}

		if test.repo.GetTrusted() != test.want.GetTrusted() {
			t.Errorf("SetTrusted is %v, want %v", test.repo.GetTrusted(), test.want.GetTrusted())
		}

		if test.repo.GetActive() != test.want.GetActive() {
			t.Errorf("SetActive is %v, want %v", test.repo.GetActive(), test.want.GetActive())
		}

		if test.repo.GetAllowPull() != test.want.GetAllowPull() {
			t.Errorf("SetAllowPull is %v, want %v", test.repo.GetAllowPull(), test.want.GetAllowPull())
		}

		if test.repo.GetAllowPush() != test.want.GetAllowPush() {
			t.Errorf("SetAllowPush is %v, want %v", test.repo.GetAllowPush(), test.want.GetAllowPush())
		}

		if test.repo.GetAllowDeploy() != test.want.GetAllowDeploy() {
			t.Errorf("SetAllowDeploy is %v, want %v", test.repo.GetAllowDeploy(), test.want.GetAllowDeploy())
		}

		if test.repo.GetAllowTag() != test.want.GetAllowTag() {
			t.Errorf("SetAllowTag is %v, want %v", test.repo.GetAllowTag(), test.want.GetAllowTag())
		}

		if test.repo.GetAllowRelease() != test.want.GetAllowRelease() {
			t.Errorf("SetAllowRelease is %v, want %v", test.repo.GetAllowRelease(), test.want.GetAllowRelease())
		}

		if test.repo.GetAllowComment() != test.want.GetAllowComment() {
			t.Errorf("SetAllowComment is %v, want %v", test.repo.GetAllowComment(), test.want.GetAllowComment())
		}

		if test.repo.GetPipelineType() != test.want.GetPipelineType() {
			t.Errorf("SetPipelineType is %v, want %v", test.repo.GetPipelineType(), test.want.GetPipelineType())
		}

		if !reflect.DeepEqual(test.repo.GetPreviousName(), test.want.GetPreviousName()) {
			t.Errorf("SetPreviousName is %v, want %v", test.repo.GetPreviousName(), test.want.GetPreviousName())
		}
	}
}

func TestLibrary_Repo_String(t *testing.T) {
	// setup types
	r := testRepo()

	want := fmt.Sprintf(`{
  Active: %t,
  AllowComment: %t,
  AllowDeploy: %t,
  AllowPull: %t,
  AllowPush: %t,
  AllowTag: %t,
  AllowRelease: %t,
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
		r.GetAllowRelease(),
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
		r.GetTrusted(),
		r.GetUserID(),
		r.GetVisibility(),
	)

	// run test
	got := r.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

// testRepo is a test helper function to create a Repo
// type with all fields set to a fake value.
func testRepo() *Repo {
	r := new(Repo)

	r.SetID(1)
	r.SetOrg("github")
	r.SetName("octocat")
	r.SetFullName("github/octocat")
	r.SetLink("https://github.com/github/octocat")
	r.SetClone("https://github.com/github/octocat.git")
	r.SetBranch("master")
	r.SetBuildLimit(10)
	r.SetTimeout(30)
	r.SetCounter(0)
	r.SetVisibility("public")
	r.SetPrivate(false)
	r.SetTrusted(false)
	r.SetActive(true)
	r.SetAllowPull(false)
	r.SetAllowPush(true)
	r.SetAllowDeploy(false)
	r.SetAllowTag(false)
	r.SetAllowRelease(false)
	r.SetAllowComment(false)
	r.SetPipelineType("")
	r.SetPreviousName("")

	return r
}
