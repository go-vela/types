// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
)

func TestDatabase_User_Nullify(t *testing.T) {
	// setup types
	var u *User

	want := &User{
		ID:     sql.NullInt64{Int64: 0, Valid: false},
		Name:   sql.NullString{String: "", Valid: false},
		Token:  sql.NullString{String: "", Valid: false},
		Hash:   sql.NullString{String: "", Valid: false},
		Active: sql.NullBool{Bool: false, Valid: false},
		Admin:  sql.NullBool{Bool: false, Valid: false},
	}

	// setup tests
	tests := []struct {
		user *User
		want *User
	}{
		{
			user: testUser(),
			want: testUser(),
		},
		{
			user: u,
			want: nil,
		},
		{
			user: new(User),
			want: want,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.user.Nullify()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Nullify is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_User_ToLibrary(t *testing.T) {
	// setup types
	want := new(library.User)

	want.SetID(1)
	want.SetName("octocat")
	want.SetToken("superSecretToken")
	want.SetHash("superSecretHash")
	want.SetFavorites([]string{"github/octocat"})
	want.SetActive(true)
	want.SetAdmin(false)

	// run test
	got := testUser().ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_User_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		user    *User
	}{
		{
			failure: false,
			user:    testUser(),
		},
		{ // no name set for user
			failure: true,
			user: &User{
				ID:    sql.NullInt64{Int64: 1, Valid: true},
				Token: sql.NullString{String: "superSecretToken", Valid: true},
				Hash:  sql.NullString{String: "superSecretHash", Valid: true},
			},
		},
		{ // no token set for user
			failure: true,
			user: &User{
				ID:   sql.NullInt64{Int64: 1, Valid: true},
				Name: sql.NullString{String: "octocat", Valid: true},
				Hash: sql.NullString{String: "superSecretHash", Valid: true},
			},
		},
		{ // no hash set for user
			failure: true,
			user: &User{
				ID:    sql.NullInt64{Int64: 1, Valid: true},
				Name:  sql.NullString{String: "octocat", Valid: true},
				Token: sql.NullString{String: "superSecretToken", Valid: true},
			},
		},
		{ // invalid name set for user
			failure: true,
			user: &User{
				ID:    sql.NullInt64{Int64: 1, Valid: true},
				Name:  sql.NullString{String: "!@#$%^&*()", Valid: true},
				Token: sql.NullString{String: "superSecretToken", Valid: true},
				Hash:  sql.NullString{String: "superSecretHash", Valid: true},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.user.Validate()

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

func TestDatabase_UserFromLibrary(t *testing.T) {
	// setup types
	u := new(library.User)

	u.SetID(1)
	u.SetName("octocat")
	u.SetToken("superSecretToken")
	u.SetHash("superSecretHash")
	u.SetFavorites([]string{"github/octocat"})
	u.SetActive(true)
	u.SetAdmin(false)

	want := testUser()

	// run test
	got := UserFromLibrary(u)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UserFromLibrary is %v, want %v", got, want)
	}
}

// testUser is a test helper function to create a User
// type with all fields set to a fake value.
func testUser() *User {
	return &User{
		ID:        sql.NullInt64{Int64: 1, Valid: true},
		Name:      sql.NullString{String: "octocat", Valid: true},
		Token:     sql.NullString{String: "superSecretToken", Valid: true},
		Hash:      sql.NullString{String: "superSecretHash", Valid: true},
		Favorites: []string{"github/octocat"},
		Active:    sql.NullBool{Bool: true, Valid: true},
		Admin:     sql.NullBool{Bool: false, Valid: true},
	}
}
