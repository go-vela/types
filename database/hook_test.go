// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/go-vela/types/library"
)

func TestDatabase_Hook_Nullify(t *testing.T) {
	// setup types
	var h *Hook

	want := &Hook{
		ID:          sql.NullInt64{Int64: 0, Valid: false},
		RepoID:      sql.NullInt64{Int64: 0, Valid: false},
		BuildID:     sql.NullInt64{Int64: 0, Valid: false},
		Number:      sql.NullInt32{Int32: 0, Valid: false},
		SourceID:    sql.NullString{String: "", Valid: false},
		Created:     sql.NullInt64{Int64: 0, Valid: false},
		Host:        sql.NullString{String: "", Valid: false},
		Event:       sql.NullString{String: "", Valid: false},
		EventAction: sql.NullString{String: "", Valid: false},
		Branch:      sql.NullString{String: "", Valid: false},
		Error:       sql.NullString{String: "", Valid: false},
		Status:      sql.NullString{String: "", Valid: false},
		Link:        sql.NullString{String: "", Valid: false},
		WebhookID:   sql.NullInt64{Int64: 0, Valid: false},
	}

	// setup tests
	tests := []struct {
		hook *Hook
		want *Hook
	}{
		{
			hook: testHook(),
			want: testHook(),
		},
		{
			hook: h,
			want: nil,
		},
		{
			hook: new(Hook),
			want: want,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.hook.Nullify()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Nullify is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_Hook_ToLibrary(t *testing.T) {
	// setup types
	want := new(library.Hook)
	want.SetID(1)
	want.SetRepoID(1)
	want.SetBuildID(1)
	want.SetNumber(1)
	want.SetSourceID("c8da1302-07d6-11ea-882f-4893bca275b8")
	want.SetCreated(time.Now().UTC().Unix())
	want.SetHost("github.com")
	want.SetEvent("push")
	want.SetEventAction("")
	want.SetBranch("master")
	want.SetError("")
	want.SetStatus("success")
	want.SetLink("https://github.com/github/octocat/settings/hooks/1")
	want.SetWebhookID(123456)

	h := &Hook{
		ID:          sql.NullInt64{Int64: 1, Valid: true},
		RepoID:      sql.NullInt64{Int64: 1, Valid: true},
		BuildID:     sql.NullInt64{Int64: 1, Valid: true},
		Number:      sql.NullInt32{Int32: 1, Valid: true},
		SourceID:    sql.NullString{String: "c8da1302-07d6-11ea-882f-4893bca275b8", Valid: true},
		Created:     sql.NullInt64{Int64: time.Now().UTC().Unix(), Valid: true},
		Host:        sql.NullString{String: "github.com", Valid: true},
		Event:       sql.NullString{String: "push", Valid: true},
		EventAction: sql.NullString{String: "", Valid: true},
		Branch:      sql.NullString{String: "master", Valid: true},
		Error:       sql.NullString{String: "", Valid: true},
		Status:      sql.NullString{String: "success", Valid: true},
		Link:        sql.NullString{String: "https://github.com/github/octocat/settings/hooks/1", Valid: true},
		WebhookID:   sql.NullInt64{Int64: 123456, Valid: true},
	}

	// run test
	got := h.ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Hook_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		hook    *Hook
	}{
		{
			failure: false,
			hook:    testHook(),
		},
		{ // no number set for hook
			failure: true,
			hook: &Hook{
				ID:        sql.NullInt64{Int64: 1, Valid: true},
				RepoID:    sql.NullInt64{Int64: 1, Valid: true},
				SourceID:  sql.NullString{String: "c8da1302-07d6-11ea-882f-4893bca275b8", Valid: true},
				WebhookID: sql.NullInt64{Int64: 1, Valid: true},
			},
		},
		{ // no repo_id set for hook
			failure: true,
			hook: &Hook{
				ID:        sql.NullInt64{Int64: 1, Valid: true},
				Number:    sql.NullInt32{Int32: 1, Valid: true},
				SourceID:  sql.NullString{String: "c8da1302-07d6-11ea-882f-4893bca275b8", Valid: true},
				WebhookID: sql.NullInt64{Int64: 1, Valid: true},
			},
		},
		{ // no source_id set for hook
			failure: true,
			hook: &Hook{
				ID:        sql.NullInt64{Int64: 1, Valid: true},
				Number:    sql.NullInt32{Int32: 1, Valid: true},
				RepoID:    sql.NullInt64{Int64: 1, Valid: true},
				WebhookID: sql.NullInt64{Int64: 1, Valid: true},
			},
		},
		{ // no webhook_id set for hook
			failure: true,
			hook: &Hook{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				Number:   sql.NullInt32{Int32: 1, Valid: true},
				RepoID:   sql.NullInt64{Int64: 1, Valid: true},
				SourceID: sql.NullString{String: "c8da1302-07d6-11ea-882f-4893bca275b8", Valid: true},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.hook.Validate()

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

func TestDatabase_HookFromLibrary(t *testing.T) {
	// setup types
	want := &Hook{
		ID:          sql.NullInt64{Int64: 1, Valid: true},
		RepoID:      sql.NullInt64{Int64: 1, Valid: true},
		BuildID:     sql.NullInt64{Int64: 1, Valid: true},
		Number:      sql.NullInt32{Int32: 1, Valid: true},
		SourceID:    sql.NullString{String: "c8da1302-07d6-11ea-882f-4893bca275b8", Valid: true},
		Created:     sql.NullInt64{Int64: time.Now().UTC().Unix(), Valid: true},
		Host:        sql.NullString{String: "github.com", Valid: true},
		Event:       sql.NullString{String: "pull_request", Valid: true},
		EventAction: sql.NullString{String: "opened", Valid: true},
		Branch:      sql.NullString{String: "master", Valid: true},
		Error:       sql.NullString{String: "", Valid: false},
		Status:      sql.NullString{String: "success", Valid: true},
		Link:        sql.NullString{String: "https://github.com/github/octocat/settings/hooks/1", Valid: true},
		WebhookID:   sql.NullInt64{Int64: 123456, Valid: true},
	}

	h := new(library.Hook)
	h.SetID(1)
	h.SetRepoID(1)
	h.SetBuildID(1)
	h.SetNumber(1)
	h.SetSourceID("c8da1302-07d6-11ea-882f-4893bca275b8")
	h.SetCreated(time.Now().UTC().Unix())
	h.SetHost("github.com")
	h.SetEvent("pull_request")
	h.SetEventAction("opened")
	h.SetBranch("master")
	h.SetError("")
	h.SetStatus("success")
	h.SetLink("https://github.com/github/octocat/settings/hooks/1")
	h.SetWebhookID(123456)

	// run test
	got := HookFromLibrary(h)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("HookFromLibrary is %v, want %v", got, want)
	}
}

// testHook is a test helper function to create a Hook
// type with all fields set to a fake value.
func testHook() *Hook {
	return &Hook{
		ID:          sql.NullInt64{Int64: 1, Valid: true},
		RepoID:      sql.NullInt64{Int64: 1, Valid: true},
		BuildID:     sql.NullInt64{Int64: 1, Valid: true},
		Number:      sql.NullInt32{Int32: 1, Valid: true},
		SourceID:    sql.NullString{String: "c8da1302-07d6-11ea-882f-4893bca275b8", Valid: true},
		Created:     sql.NullInt64{Int64: time.Now().UTC().Unix(), Valid: true},
		Host:        sql.NullString{String: "github.com", Valid: true},
		Event:       sql.NullString{String: "push", Valid: true},
		EventAction: sql.NullString{String: "", Valid: false},
		Branch:      sql.NullString{String: "master", Valid: true},
		Error:       sql.NullString{String: "", Valid: false},
		Status:      sql.NullString{String: "success", Valid: true},
		Link:        sql.NullString{String: "https://github.com/github/octocat/settings/hooks/1", Valid: true},
		WebhookID:   sql.NullInt64{Int64: 123456, Valid: true},
	}
}
