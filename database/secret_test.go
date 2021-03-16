// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
)

func TestDatabase_Secret_Decrypt(t *testing.T) {
	// setup types
	key := "C639A572E14D5075C526FDDD43E4ECF6"

	encrypted := testSecret()
	err := encrypted.Encrypt(key)
	if err != nil {
		t.Errorf("unable to encrypt secret: %v", err)
	}

	// setup tests
	tests := []struct {
		failure bool
		key     string
		secret  Secret
	}{
		{
			failure: false,
			key:     key,
			secret:  *encrypted,
		},
		{
			failure: true,
			key:     "",
			secret:  *encrypted,
		},
		{
			failure: true,
			key:     key,
			secret:  *testSecret(),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.secret.Decrypt(test.key)

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

func TestDatabase_Secret_Encrypt(t *testing.T) {
	// setup types
	key := "C639A572E14D5075C526FDDD43E4ECF6"

	// setup tests
	tests := []struct {
		failure bool
		key     string
		secret  *Secret
	}{
		{
			failure: false,
			key:     key,
			secret:  testSecret(),
		},
		{
			failure: true,
			key:     "",
			secret:  testSecret(),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.secret.Encrypt(test.key)

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

func TestDatabase_Secret_Nullify(t *testing.T) {
	// setup types
	var s *Secret

	want := &Secret{
		ID:    sql.NullInt64{Int64: 0, Valid: false},
		Org:   sql.NullString{String: "", Valid: false},
		Repo:  sql.NullString{String: "", Valid: false},
		Team:  sql.NullString{String: "", Valid: false},
		Name:  sql.NullString{String: "", Valid: false},
		Value: sql.NullString{String: "", Valid: false},
		Type:  sql.NullString{String: "", Valid: false},
	}

	// setup tests
	tests := []struct {
		secret *Secret
		want   *Secret
	}{
		{
			secret: testSecret(),
			want:   testSecret(),
		},
		{
			secret: s,
			want:   nil,
		},
		{
			secret: new(Secret),
			want:   want,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.secret.Nullify()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Nullify is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_Secret_ToLibrary(t *testing.T) {
	// setup types
	want := new(library.Secret)

	want.SetID(1)
	want.SetOrg("github")
	want.SetRepo("octocat")
	want.SetTeam("octokitties")
	want.SetName("foo")
	want.SetValue("bar")
	want.SetType("repo")
	want.SetImages([]string{"alpine"})
	want.SetEvents([]string{"push", "tag", "deployment"})
	want.SetAllowCommand(true)

	// run test
	got := testSecret().ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Secret_Validate(t *testing.T) {
	// setup types
	tests := []struct {
		failure bool
		secret  *Secret
	}{
		{
			failure: false,
			secret:  testSecret(),
		},
		{ // no name set for secret
			failure: true,
			secret: &Secret{
				ID:    sql.NullInt64{Int64: 1, Valid: true},
				Org:   sql.NullString{String: "github", Valid: true},
				Repo:  sql.NullString{String: "octocat", Valid: true},
				Team:  sql.NullString{String: "octokitties", Valid: true},
				Value: sql.NullString{String: "bar", Valid: true},
				Type:  sql.NullString{String: "repo", Valid: true},
			},
		},
		{ // no org set for secret
			failure: true,
			secret: &Secret{
				ID:    sql.NullInt64{Int64: 1, Valid: true},
				Repo:  sql.NullString{String: "octocat", Valid: true},
				Team:  sql.NullString{String: "octokitties", Valid: true},
				Name:  sql.NullString{String: "foo", Valid: true},
				Value: sql.NullString{String: "bar", Valid: true},
				Type:  sql.NullString{String: "repo", Valid: true},
			},
		},
		{ // no repo set for secret
			failure: true,
			secret: &Secret{
				ID:    sql.NullInt64{Int64: 1, Valid: true},
				Org:   sql.NullString{String: "github", Valid: true},
				Team:  sql.NullString{String: "octokitties", Valid: true},
				Name:  sql.NullString{String: "foo", Valid: true},
				Value: sql.NullString{String: "bar", Valid: true},
				Type:  sql.NullString{String: "repo", Valid: true},
			},
		},
		{ // no team set for secret
			failure: true,
			secret: &Secret{
				ID:    sql.NullInt64{Int64: 1, Valid: true},
				Org:   sql.NullString{String: "github", Valid: true},
				Repo:  sql.NullString{String: "octocat", Valid: true},
				Name:  sql.NullString{String: "foo", Valid: true},
				Value: sql.NullString{String: "bar", Valid: true},
				Type:  sql.NullString{String: "shared", Valid: true},
			},
		},
		{ // no type set for secret
			failure: true,
			secret: &Secret{
				ID:    sql.NullInt64{Int64: 1, Valid: true},
				Org:   sql.NullString{String: "github", Valid: true},
				Repo:  sql.NullString{String: "octocat", Valid: true},
				Team:  sql.NullString{String: "octokitties", Valid: true},
				Name:  sql.NullString{String: "foo", Valid: true},
				Value: sql.NullString{String: "bar", Valid: true},
			},
		},
		{ // no value set for secret
			failure: true,
			secret: &Secret{
				ID:   sql.NullInt64{Int64: 1, Valid: true},
				Org:  sql.NullString{String: "github", Valid: true},
				Repo: sql.NullString{String: "octocat", Valid: true},
				Team: sql.NullString{String: "octokitties", Valid: true},
				Name: sql.NullString{String: "foo", Valid: true},
				Type: sql.NullString{String: "repo", Valid: true},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.secret.Validate()

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

func TestDatabase_SecretFromLibrary(t *testing.T) {
	// setup types
	s := new(library.Secret)

	s.SetID(1)
	s.SetOrg("github")
	s.SetRepo("octocat")
	s.SetTeam("octokitties")
	s.SetName("foo")
	s.SetValue("bar")
	s.SetType("repo")
	s.SetImages([]string{"alpine"})
	s.SetEvents([]string{"push", "tag", "deployment"})
	s.SetAllowCommand(true)

	want := testSecret()

	// run test
	got := SecretFromLibrary(s)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SecretFromLibrary is %v, want %v", got, want)
	}
}

// testSecret is a test helper function to create a Secret
// type with all fields set to a fake value.
func testSecret() *Secret {
	return &Secret{
		ID:           sql.NullInt64{Int64: 1, Valid: true},
		Org:          sql.NullString{String: "github", Valid: true},
		Repo:         sql.NullString{String: "octocat", Valid: true},
		Team:         sql.NullString{String: "octokitties", Valid: true},
		Name:         sql.NullString{String: "foo", Valid: true},
		Value:        sql.NullString{String: "bar", Valid: true},
		Type:         sql.NullString{String: "repo", Valid: true},
		Images:       []string{"alpine"},
		Events:       []string{"push", "tag", "deployment"},
		AllowCommand: sql.NullBool{Bool: true, Valid: true},
	}
}
