// SPDX-License-Identifier: Apache-2.0

package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
	"github.com/go-vela/types/library/actions"
)

func TestDatabase_Repo_Decrypt(t *testing.T) {
	// setup types
	key := "C639A572E14D5075C526FDDD43E4ECF6"
	encrypted := testRepo()

	err := encrypted.Encrypt(key)
	if err != nil {
		t.Errorf("unable to encrypt repo: %v", err)
	}

	// setup tests
	tests := []struct {
		failure bool
		key     string
		repo    Repo
	}{
		{
			failure: false,
			key:     key,
			repo:    *encrypted,
		},
		{
			failure: true,
			key:     "",
			repo:    *encrypted,
		},
		{
			failure: true,
			key:     key,
			repo:    *testRepo(),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.repo.Decrypt(test.key)

		if test.failure {
			if err == nil {
				t.Errorf("Decrypt should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Decrypt returned err: %v", err)
		}
	}
}

func TestDatabase_Repo_Encrypt(t *testing.T) {
	// setup types
	key := "C639A572E14D5075C526FDDD43E4ECF6"

	// setup tests
	tests := []struct {
		failure bool
		key     string
		repo    *Repo
	}{
		{
			failure: false,
			key:     key,
			repo:    testRepo(),
		},
		{
			failure: true,
			key:     "",
			repo:    testRepo(),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.repo.Encrypt(test.key)

		if test.failure {
			if err == nil {
				t.Errorf("Encrypt should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Encrypt returned err: %v", err)
		}
	}
}

func TestDatabase_Repo_Nullify(t *testing.T) {
	// setup types
	var r *Repo

	want := &Repo{
		ID:           sql.NullInt64{Int64: 0, Valid: false},
		UserID:       sql.NullInt64{Int64: 0, Valid: false},
		Hash:         sql.NullString{String: "", Valid: false},
		Org:          sql.NullString{String: "", Valid: false},
		Name:         sql.NullString{String: "", Valid: false},
		FullName:     sql.NullString{String: "", Valid: false},
		Link:         sql.NullString{String: "", Valid: false},
		Clone:        sql.NullString{String: "", Valid: false},
		Branch:       sql.NullString{String: "", Valid: false},
		Timeout:      sql.NullInt64{Int64: 0, Valid: false},
		AllowEvents:  sql.NullInt64{Int64: 0, Valid: false},
		Visibility:   sql.NullString{String: "", Valid: false},
		PipelineType: sql.NullString{String: "", Valid: false},
	}

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
			want: nil,
		},
		{
			repo: new(Repo),
			want: want,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.repo.Nullify()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Nullify is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_Repo_ToLibrary(t *testing.T) {
	// setup types
	want := new(library.Repo)
	e := library.NewEventsFromMask(1)

	want.SetID(1)
	want.SetUserID(1)
	want.SetHash("superSecretHash")
	want.SetOrg("github")
	want.SetName("octocat")
	want.SetFullName("github/octocat")
	want.SetLink("https://github.com/github/octocat")
	want.SetClone("https://github.com/github/octocat.git")
	want.SetBranch("main")
	want.SetTopics([]string{"cloud", "security"})
	want.SetBuildLimit(10)
	want.SetTimeout(30)
	want.SetCounter(0)
	want.SetVisibility("public")
	want.SetPrivate(false)
	want.SetTrusted(false)
	want.SetActive(true)
	want.SetAllowEvents(e)
	want.SetPipelineType("yaml")
	want.SetPreviousName("oldName")

	// run test
	got := testRepo().ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Repo_Validate(t *testing.T) {
	// setup types
	topics := []string{}
	longTopic := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for len(topics) < 21 {
		topics = append(topics, longTopic)
	}

	// setup tests
	tests := []struct {
		failure bool
		repo    *Repo
	}{
		{
			failure: false,
			repo:    testRepo(),
		},
		{ // no user_id set for repo
			failure: true,
			repo: &Repo{
				ID:         sql.NullInt64{Int64: 1, Valid: true},
				Hash:       sql.NullString{String: "superSecretHash", Valid: true},
				Org:        sql.NullString{String: "github", Valid: true},
				Name:       sql.NullString{String: "octocat", Valid: true},
				FullName:   sql.NullString{String: "github/octocat", Valid: true},
				Visibility: sql.NullString{String: "public", Valid: true},
			},
		},
		{ // no hash set for repo
			failure: true,
			repo: &Repo{
				ID:         sql.NullInt64{Int64: 1, Valid: true},
				UserID:     sql.NullInt64{Int64: 1, Valid: true},
				Org:        sql.NullString{String: "github", Valid: true},
				Name:       sql.NullString{String: "octocat", Valid: true},
				FullName:   sql.NullString{String: "github/octocat", Valid: true},
				Visibility: sql.NullString{String: "public", Valid: true},
			},
		},
		{ // no org set for repo
			failure: true,
			repo: &Repo{
				ID:         sql.NullInt64{Int64: 1, Valid: true},
				UserID:     sql.NullInt64{Int64: 1, Valid: true},
				Hash:       sql.NullString{String: "superSecretHash", Valid: true},
				Name:       sql.NullString{String: "octocat", Valid: true},
				FullName:   sql.NullString{String: "github/octocat", Valid: true},
				Visibility: sql.NullString{String: "public", Valid: true},
			},
		},
		{ // no name set for repo
			failure: true,
			repo: &Repo{
				ID:         sql.NullInt64{Int64: 1, Valid: true},
				UserID:     sql.NullInt64{Int64: 1, Valid: true},
				Hash:       sql.NullString{String: "superSecretHash", Valid: true},
				Org:        sql.NullString{String: "github", Valid: true},
				FullName:   sql.NullString{String: "github/octocat", Valid: true},
				Visibility: sql.NullString{String: "public", Valid: true},
			},
		},
		{ // no full_name set for repo
			failure: true,
			repo: &Repo{
				ID:         sql.NullInt64{Int64: 1, Valid: true},
				UserID:     sql.NullInt64{Int64: 1, Valid: true},
				Hash:       sql.NullString{String: "superSecretHash", Valid: true},
				Org:        sql.NullString{String: "github", Valid: true},
				Name:       sql.NullString{String: "octocat", Valid: true},
				Visibility: sql.NullString{String: "public", Valid: true},
			},
		},
		{ // no visibility set for repo
			failure: true,
			repo: &Repo{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				UserID:   sql.NullInt64{Int64: 1, Valid: true},
				Hash:     sql.NullString{String: "superSecretHash", Valid: true},
				Org:      sql.NullString{String: "github", Valid: true},
				Name:     sql.NullString{String: "octocat", Valid: true},
				FullName: sql.NullString{String: "github/octocat", Valid: true},
			},
		},
		{ // topics exceed max size
			failure: true,
			repo: &Repo{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				UserID:   sql.NullInt64{Int64: 1, Valid: true},
				Hash:     sql.NullString{String: "superSecretHash", Valid: true},
				Org:      sql.NullString{String: "github", Valid: true},
				Name:     sql.NullString{String: "octocat", Valid: true},
				FullName: sql.NullString{String: "github/octocat", Valid: true},
				Topics:   topics,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.repo.Validate()

		if test.failure {
			if err == nil {
				t.Errorf("Validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}

func TestDatabase_RepoFromLibrary(t *testing.T) {
	// setup types
	r := new(library.Repo)
	e := new(library.Events)
	e.SetPush(new(actions.Push).FromMask(1))

	r.SetID(1)
	r.SetUserID(1)
	r.SetHash("superSecretHash")
	r.SetOrg("github")
	r.SetName("octocat")
	r.SetFullName("github/octocat")
	r.SetLink("https://github.com/github/octocat")
	r.SetClone("https://github.com/github/octocat.git")
	r.SetBranch("main")
	r.SetTopics([]string{"cloud", "security"})
	r.SetBuildLimit(10)
	r.SetTimeout(30)
	r.SetCounter(0)
	r.SetVisibility("public")
	r.SetPrivate(false)
	r.SetTrusted(false)
	r.SetActive(true)
	r.SetAllowEvents(e)
	r.SetPipelineType("yaml")
	r.SetPreviousName("oldName")

	want := testRepo()

	// run test
	got := RepoFromLibrary(r)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RepoFromLibrary is %v, want %v", got, want)
	}
}

// testRepo is a test helper function to create a Repo
// type with all fields set to a fake value.
func testRepo() *Repo {
	return &Repo{
		ID:           sql.NullInt64{Int64: 1, Valid: true},
		UserID:       sql.NullInt64{Int64: 1, Valid: true},
		Hash:         sql.NullString{String: "superSecretHash", Valid: true},
		Org:          sql.NullString{String: "github", Valid: true},
		Name:         sql.NullString{String: "octocat", Valid: true},
		FullName:     sql.NullString{String: "github/octocat", Valid: true},
		Link:         sql.NullString{String: "https://github.com/github/octocat", Valid: true},
		Clone:        sql.NullString{String: "https://github.com/github/octocat.git", Valid: true},
		Branch:       sql.NullString{String: "main", Valid: true},
		Topics:       []string{"cloud", "security"},
		BuildLimit:   sql.NullInt64{Int64: 10, Valid: true},
		Timeout:      sql.NullInt64{Int64: 30, Valid: true},
		Counter:      sql.NullInt32{Int32: 0, Valid: true},
		Visibility:   sql.NullString{String: "public", Valid: true},
		Private:      sql.NullBool{Bool: false, Valid: true},
		Trusted:      sql.NullBool{Bool: false, Valid: true},
		Active:       sql.NullBool{Bool: true, Valid: true},
		AllowEvents:  sql.NullInt64{Int64: 1, Valid: true},
		PipelineType: sql.NullString{String: "yaml", Valid: true},
		PreviousName: sql.NullString{String: "oldName", Valid: true},
	}
}
