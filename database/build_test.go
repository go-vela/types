// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"math/rand"
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
	"github.com/go-vela/types/raw"
)

func TestDatabase_Build_Crop(t *testing.T) {
	// setup types
	title := randomString(1001)
	message := randomString(2001)

	b := testBuild()
	b.Title = sql.NullString{String: title, Valid: true}
	b.Message = sql.NullString{String: message, Valid: true}

	want := testBuild()
	want.Title = sql.NullString{String: title[:1000], Valid: true}
	want.Message = sql.NullString{String: message[:2000], Valid: true}

	// run test
	got := b.Crop()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Crop is %v, want %v", got, want)
	}
}

func TestDatabase_Build_Nullify(t *testing.T) {
	// setup types
	var b *Build

	want := &Build{
		ID:            sql.NullInt64{Int64: 0, Valid: false},
		RepoID:        sql.NullInt64{Int64: 0, Valid: false},
		Number:        sql.NullInt32{Int32: 0, Valid: false},
		Parent:        sql.NullInt32{Int32: 0, Valid: false},
		Event:         sql.NullString{String: "", Valid: false},
		Status:        sql.NullString{String: "", Valid: false},
		Error:         sql.NullString{String: "", Valid: false},
		Enqueued:      sql.NullInt64{Int64: 0, Valid: false},
		Created:       sql.NullInt64{Int64: 0, Valid: false},
		Started:       sql.NullInt64{Int64: 0, Valid: false},
		Finished:      sql.NullInt64{Int64: 0, Valid: false},
		Deploy:        sql.NullString{String: "", Valid: false},
		DeployPayload: nil,
		Clone:         sql.NullString{String: "", Valid: false},
		Source:        sql.NullString{String: "", Valid: false},
		Title:         sql.NullString{String: "", Valid: false},
		Message:       sql.NullString{String: "", Valid: false},
		Commit:        sql.NullString{String: "", Valid: false},
		Sender:        sql.NullString{String: "", Valid: false},
		Author:        sql.NullString{String: "", Valid: false},
		Email:         sql.NullString{String: "", Valid: false},
		Link:          sql.NullString{String: "", Valid: false},
		Branch:        sql.NullString{String: "", Valid: false},
		Ref:           sql.NullString{String: "", Valid: false},
		BaseRef:       sql.NullString{String: "", Valid: false},
		HeadRef:       sql.NullString{String: "", Valid: false},
		Host:          sql.NullString{String: "", Valid: false},
		Runtime:       sql.NullString{String: "", Valid: false},
		Distribution:  sql.NullString{String: "", Valid: false},
	}

	// setup tests
	tests := []struct {
		build *Build
		want  *Build
	}{
		{
			build: testBuild(),
			want:  testBuild(),
		},
		{
			build: b,
			want:  nil,
		},
		{
			build: new(Build),
			want:  want,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.build.Nullify()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Nullify is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_Build_ToLibrary(t *testing.T) {
	// setup types
	want := new(library.Build)

	want.SetID(1)
	want.SetRepoID(1)
	want.SetNumber(1)
	want.SetParent(1)
	want.SetEvent("push")
	want.SetStatus("running")
	want.SetError("")
	want.SetEnqueued(1563474077)
	want.SetCreated(1563474076)
	want.SetStarted(1563474078)
	want.SetFinished(1563474079)
	want.SetDeploy("")
	want.SetDeployPayload(nil)
	want.SetClone("https://github.com/github/octocat.git")
	want.SetSource("https://github.com/github/octocat/48afb5bdc41ad69bf22588491333f7cf71135163")
	want.SetTitle("push received from https://github.com/github/octocat")
	want.SetMessage("First commit...")
	want.SetCommit("48afb5bdc41ad69bf22588491333f7cf71135163")
	want.SetSender("OctoKitty")
	want.SetAuthor("OctoKitty")
	want.SetEmail("OctoKitty@github.com")
	want.SetLink("https://example.company.com/github/octocat/1")
	want.SetBranch("master")
	want.SetRef("refs/heads/master")
	want.SetBaseRef("")
	want.SetHeadRef("")
	want.SetHost("example.company.com")
	want.SetRuntime("docker")
	want.SetDistribution("linux")
	want.SetDeployPayload(raw.StringSliceMap{"foo": "test1", "bar": "test2"})

	// run test
	got := testBuild().ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Build_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		build   *Build
		failure bool
	}{
		{
			failure: false,
			build:   testBuild(),
		},
		{ // no repo_id set for build
			failure: true,
			build: &Build{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
				Number: sql.NullInt32{Int32: 1, Valid: true},
			},
		},
		{ // no number set for build
			failure: true,
			build: &Build{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
				RepoID: sql.NullInt64{Int64: 1, Valid: true},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.build.Validate()

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

func TestDatabase_BuildFromLibrary(t *testing.T) {
	// setup types
	b := new(library.Build)

	b.SetID(1)
	b.SetRepoID(1)
	b.SetNumber(1)
	b.SetParent(1)
	b.SetEvent("push")
	b.SetStatus("running")
	b.SetError("")
	b.SetEnqueued(1563474077)
	b.SetCreated(1563474076)
	b.SetStarted(1563474078)
	b.SetFinished(1563474079)
	b.SetDeploy("")
	b.SetDeployPayload(nil)
	b.SetClone("https://github.com/github/octocat.git")
	b.SetSource("https://github.com/github/octocat/48afb5bdc41ad69bf22588491333f7cf71135163")
	b.SetTitle("push received from https://github.com/github/octocat")
	b.SetMessage("First commit...")
	b.SetCommit("48afb5bdc41ad69bf22588491333f7cf71135163")
	b.SetSender("OctoKitty")
	b.SetAuthor("OctoKitty")
	b.SetEmail("OctoKitty@github.com")
	b.SetLink("https://example.company.com/github/octocat/1")
	b.SetBranch("master")
	b.SetRef("refs/heads/master")
	b.SetBaseRef("")
	b.SetHeadRef("")
	b.SetHost("example.company.com")
	b.SetRuntime("docker")
	b.SetDistribution("linux")
	b.SetDeployPayload(raw.StringSliceMap{"foo": "test1", "bar": "test2"})

	want := testBuild()

	// run test
	got := BuildFromLibrary(b)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BuildFromLibrary is %v, want %v", got, want)
	}
}

func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		// nolint:gosec // accepting weak RNG for test
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}

// testBuild is a test helper function to create a Build
// type with all fields set to a fake value.
func testBuild() *Build {
	return &Build{
		ID:            sql.NullInt64{Int64: 1, Valid: true},
		RepoID:        sql.NullInt64{Int64: 1, Valid: true},
		Number:        sql.NullInt32{Int32: 1, Valid: true},
		Parent:        sql.NullInt32{Int32: 1, Valid: true},
		Event:         sql.NullString{String: "push", Valid: true},
		Status:        sql.NullString{String: "running", Valid: true},
		Error:         sql.NullString{String: "", Valid: false},
		Enqueued:      sql.NullInt64{Int64: 1563474077, Valid: true},
		Created:       sql.NullInt64{Int64: 1563474076, Valid: true},
		Started:       sql.NullInt64{Int64: 1563474078, Valid: true},
		Finished:      sql.NullInt64{Int64: 1563474079, Valid: true},
		Deploy:        sql.NullString{String: "", Valid: false},
		DeployPayload: raw.StringSliceMap{"foo": "test1", "bar": "test2"},
		Clone:         sql.NullString{String: "https://github.com/github/octocat.git", Valid: true},
		Source:        sql.NullString{String: "https://github.com/github/octocat/48afb5bdc41ad69bf22588491333f7cf71135163", Valid: true},
		Title:         sql.NullString{String: "push received from https://github.com/github/octocat", Valid: true},
		Message:       sql.NullString{String: "First commit...", Valid: true},
		Commit:        sql.NullString{String: "48afb5bdc41ad69bf22588491333f7cf71135163", Valid: true},
		Sender:        sql.NullString{String: "OctoKitty", Valid: true},
		Author:        sql.NullString{String: "OctoKitty", Valid: true},
		Email:         sql.NullString{String: "OctoKitty@github.com", Valid: true},
		Link:          sql.NullString{String: "https://example.company.com/github/octocat/1", Valid: true},
		Branch:        sql.NullString{String: "master", Valid: true},
		Ref:           sql.NullString{String: "refs/heads/master", Valid: true},
		BaseRef:       sql.NullString{String: "", Valid: false},
		HeadRef:       sql.NullString{String: "", Valid: false},
		Host:          sql.NullString{String: "example.company.com", Valid: true},
		Runtime:       sql.NullString{String: "docker", Valid: true},
		Distribution:  sql.NullString{String: "linux", Valid: true},
	}
}
