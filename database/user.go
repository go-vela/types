// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"errors"
	"regexp"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
	"github.com/lib/pq"
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

	// ErrEmptyRefreshUserToken defines the error type when a
	// User type has an empty RefreshToken field provided.
	ErrEmptyUserRefreshToken = errors.New("empty user refresh token provided")

	// ErrEmptyUserToken defines the error type when a
	// User type has an empty Token field provided.
	ErrEmptyUserToken = errors.New("empty user token provided")

	// ErrInvalidUserName defines the error type when a
	// User type has an invalid Name field provided.
	ErrInvalidUserName = errors.New("invalid user name provided")

	// ErrExceededFavoritesLimit defines the error type when a
	// User type has Favorites field provided that exceeds the database limit.
	ErrExceededFavoritesLimit = errors.New("exceeded favorites limit")
)

// User is the database representation of a user.
type User struct {
	ID           sql.NullInt64  `sql:"id"`
	Name         sql.NullString `sql:"name"`
	RefreshToken sql.NullString `sql:"refresh_token"`
	Token        sql.NullString `sql:"token"`
	Hash         sql.NullString `sql:"hash"`
	Favorites    pq.StringArray `sql:"favorites"`
	Active       sql.NullBool   `sql:"active"`
	Admin        sql.NullBool   `sql:"admin"`
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the User type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (u *User) Nullify() *User {
	if u == nil {
		return nil
	}

	// check if the ID field should be false
	if u.ID.Int64 == 0 {
		u.ID.Valid = false
	}

	// check if the Name field should be false
	if len(u.Name.String) == 0 {
		u.Name.Valid = false
	}

	// check if the RefreshToken field should be false
	if len(u.RefreshToken.String) == 0 {
		u.RefreshToken.Valid = false
	}

	// check if the Token field should be false
	if len(u.Token.String) == 0 {
		u.Token.Valid = false
	}

	// check if the Hash field should be false
	if len(u.Hash.String) == 0 {
		u.Hash.Valid = false
	}

	return u
}

// ToLibrary converts the User type
// to a library User type.
func (u *User) ToLibrary() *library.User {
	user := new(library.User)

	user.SetID(u.ID.Int64)
	user.SetName(u.Name.String)
	user.SetRefreshToken(u.RefreshToken.String)
	user.SetToken(u.Token.String)
	user.SetHash(u.Hash.String)
	user.SetActive(u.Active.Bool)
	user.SetAdmin(u.Admin.Bool)
	user.SetFavorites(u.Favorites)

	return user
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

	// calculate total size of favorites
	total := 0
	for _, f := range u.Favorites {
		total += len(f)
	}

	// verify the Favorites field is within the database constraints
	if total > constants.FavoritesMaxSize {
		return ErrExceededFavoritesLimit
	}

	return nil
}

// UserFromLibrary converts the library User type
// to a database User type.
func UserFromLibrary(u *library.User) *User {
	user := &User{
		ID:           sql.NullInt64{Int64: u.GetID(), Valid: true},
		Name:         sql.NullString{String: u.GetName(), Valid: true},
		RefreshToken: sql.NullString{String: u.GetRefreshToken(), Valid: true},
		Token:        sql.NullString{String: u.GetToken(), Valid: true},
		Hash:         sql.NullString{String: u.GetHash(), Valid: true},
		Active:       sql.NullBool{Bool: u.GetActive(), Valid: true},
		Admin:        sql.NullBool{Bool: u.GetAdmin(), Valid: true},
		Favorites:    u.GetFavorites(),
	}

	return user.Nullify()
}
