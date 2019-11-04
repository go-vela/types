// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"errors"
	"regexp"
	"strings"

	"github.com/go-vela/types/library"
)

var (
	// userRegex defines the regex pattern for validating
	// the Name field for the User type.
	userRegex = regexp.MustCompile("^[a-zA-Z0-9_-]{0,38}$")

	// ErrEmptyUserHash defines the error type when a
	// User type has an empty Hash field provided.
	ErrEmptyUserHash = errors.New("empty user hash provided")

	// ErrEmptyUserName defines the error type when a
	// User type has an empty Name field provided.
	ErrEmptyUserName = errors.New("empty user name provided")

	// ErrEmptyUserToken defines the error type when a
	// User type has an empty Token field provided.
	ErrEmptyUserToken = errors.New("empty user token provided")

	// ErrInvalidUserName defines the error type when a
	// User type has an invalid Name field provided.
	ErrInvalidUserName = errors.New("invalid user name provided")
)

// User is the database representation of a user.
type User struct {
	ID     sql.NullInt64  `sql:"id"`
	Name   sql.NullString `sql:"name"`
	Token  sql.NullString `sql:"token"`
	Hash   sql.NullString `sql:"hash"`
	Active sql.NullBool   `sql:"active"`
	Admin  sql.NullBool   `sql:"admin"`
}

// ToLibrary converts the User type
// to a library User type.
func (u *User) ToLibrary() *library.User {
	return &library.User{
		ID:     &u.ID.Int64,
		Name:   &u.Name.String,
		Token:  &u.Token.String,
		Hash:   &u.Hash.String,
		Active: &u.Active.Bool,
		Admin:  &u.Admin.Bool,
	}
}

// Nullify is a helper function to overwrite fields in the
// user to ensure the valid flag is properly set for a sqlnull type.
func (u *User) nullify() {

	// check if the ID should be false
	if u.ID.Int64 == 0 {
		u.ID.Valid = false
	}

	// check if the Name should be false
	if strings.EqualFold(u.Name.String, "") {
		u.Name.Valid = false
	}

	// check if the Token should be false
	if strings.EqualFold(u.Token.String, "") {
		u.Token.Valid = false
	}

	// check if the Hash should be false
	if strings.EqualFold(u.Hash.String, "") {
		u.Hash.Valid = false
	}
}

// UserFromLibrary converts the library User type
// to a database User type.
func UserFromLibrary(u *library.User) *User {

	entry := &User{
		ID:     sql.NullInt64{Int64: u.GetID(), Valid: true},
		Name:   sql.NullString{String: u.GetName(), Valid: true},
		Token:  sql.NullString{String: u.GetToken(), Valid: true},
		Hash:   sql.NullString{String: u.GetHash(), Valid: true},
		Active: sql.NullBool{Bool: u.GetActive(), Valid: true},
		Admin:  sql.NullBool{Bool: u.GetAdmin(), Valid: true},
	}

	entry.nullify()

	return entry
}

// Validate verifies the necessary fields for
// the User type are populated correctly.
func (u *User) Validate() error {
	// verify the Name field is populated
	if len(u.Name.String) == 0 {
		return ErrEmptyUserName
	}

	// verify the Token field is populated
	if len(u.Token.String) == 0 {
		return ErrEmptyUserToken
	}

	// verify the Hash field is populated
	if len(u.Hash.String) == 0 {
		return ErrEmptyUserHash
	}

	// verify the Name field is valid
	if !userRegex.MatchString(u.Name.String) {
		return ErrInvalidUserName
	}

	return nil
}
