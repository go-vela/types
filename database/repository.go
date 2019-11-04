// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/go-vela/types/library"
)

var (
	// ErrEmptyRepoFullName defines the error type when a
	// Repo type has an empty FullName field provided.
	ErrEmptyRepoFullName = errors.New("empty repo full_name provided")

	// ErrEmptyRepoName defines the error type when a
	// Repo type has an empty Name field provided.
	ErrEmptyRepoName = errors.New("empty repo name provided")

	// ErrEmptyRepoOrg defines the error type when a
	// Repo type has an empty Org field provided.
	ErrEmptyRepoOrg = errors.New("empty repo org provided")

	// ErrEmptyRepoUserID defines the error type when a
	// Repo type has an empty UserID field provided.
	ErrEmptyRepoUserID = errors.New("empty repo user_id provided")
)

// Repo is the database representation of a repo.
type Repo struct {
	ID          sql.NullInt64  `sql:"id"`
	UserID      sql.NullInt64  `sql:"user_id"`
	Org         sql.NullString `sql:"org"`
	Name        sql.NullString `sql:"name"`
	FullName    sql.NullString `sql:"full_name"`
	Link        sql.NullString `sql:"link"`
	Clone       sql.NullString `sql:"clone"`
	Branch      sql.NullString `sql:"branch"`
	Timeout     sql.NullInt64  `sql:"timeout"`
	Visibility  sql.NullString `sql:"visibility"`
	Private     sql.NullBool   `sql:"private"`
	Trusted     sql.NullBool   `sql:"trusted"`
	Active      sql.NullBool   `sql:"active"`
	AllowPull   sql.NullBool   `sql:"allow_pr"`
	AllowPush   sql.NullBool   `sql:"allow_push"`
	AllowDeploy sql.NullBool   `sql:"allow_deploy"`
	AllowTag    sql.NullBool   `sql:"allow_tag"`
}

// ToLibrary converts the Repo type
// to a library Repo type.
func (r *Repo) ToLibrary() *library.Repo {
	return &library.Repo{
		ID:          &r.ID.Int64,
		UserID:      &r.UserID.Int64,
		Org:         &r.Org.String,
		Name:        &r.Name.String,
		FullName:    &r.FullName.String,
		Link:        &r.Link.String,
		Clone:       &r.Clone.String,
		Branch:      &r.Branch.String,
		Timeout:     &r.Timeout.Int64,
		Visibility:  &r.Visibility.String,
		Private:     &r.Private.Bool,
		Trusted:     &r.Trusted.Bool,
		Active:      &r.Active.Bool,
		AllowPull:   &r.AllowPull.Bool,
		AllowPush:   &r.AllowPush.Bool,
		AllowDeploy: &r.AllowDeploy.Bool,
		AllowTag:    &r.AllowTag.Bool,
	}
}

// Nullify is a helper function to overwrite fields in the
// repo to ensure the valid flag is properly set for a sqlnull type.
func (r *Repo) nullify() {

	// check if the ID should be false
	if r.ID.Int64 == 0 {
		r.ID.Valid = false
	}

	// check if the UserID should be false
	if r.UserID.Int64 == 0 {
		r.UserID.Valid = false
	}

	// check if the Org should be false
	if strings.EqualFold(r.Org.String, "") {
		r.Org.Valid = false
	}

	// check if the Name should be false
	if strings.EqualFold(r.Name.String, "") {
		r.Name.Valid = false
	}

	// check if the FullName should be false
	if strings.EqualFold(r.FullName.String, "") {
		r.FullName.Valid = false
	}

	// check if the Link should be false
	if strings.EqualFold(r.Link.String, "") {
		r.Link.Valid = false
	}

	// check if the Clone should be false
	if strings.EqualFold(r.Clone.String, "") {
		r.Clone.Valid = false
	}

	// check if the Branch should be false
	if strings.EqualFold(r.Branch.String, "") {
		r.Branch.Valid = false
	}

	// check if the Timeout should be false
	if r.Timeout.Int64 == 0 {
		r.Timeout.Valid = false
	}

	// check if the Visibility should be false
	if strings.EqualFold(r.Visibility.String, "") {
		r.Visibility.Valid = false
	}
}

// RepoFromLibrary converts the libray Repo type
// to a database repo type.
func RepoFromLibrary(r *library.Repo) *Repo {
	entry := &Repo{
		ID:          sql.NullInt64{Int64: r.GetID(), Valid: true},
		UserID:      sql.NullInt64{Int64: r.GetUserID(), Valid: true},
		Org:         sql.NullString{String: r.GetOrg(), Valid: true},
		Name:        sql.NullString{String: r.GetName(), Valid: true},
		FullName:    sql.NullString{String: r.GetFullName(), Valid: true},
		Link:        sql.NullString{String: r.GetLink(), Valid: true},
		Clone:       sql.NullString{String: r.GetClone(), Valid: true},
		Branch:      sql.NullString{String: r.GetBranch(), Valid: true},
		Timeout:     sql.NullInt64{Int64: r.GetTimeout(), Valid: true},
		Visibility:  sql.NullString{String: r.GetVisibility(), Valid: true},
		Private:     sql.NullBool{Bool: r.GetPrivate(), Valid: true},
		Trusted:     sql.NullBool{Bool: r.GetTrusted(), Valid: true},
		Active:      sql.NullBool{Bool: r.GetActive(), Valid: true},
		AllowPull:   sql.NullBool{Bool: r.GetAllowPull(), Valid: true},
		AllowPush:   sql.NullBool{Bool: r.GetAllowPush(), Valid: true},
		AllowDeploy: sql.NullBool{Bool: r.GetAllowDeploy(), Valid: true},
		AllowTag:    sql.NullBool{Bool: r.GetAllowTag(), Valid: true},
	}

	entry.nullify()

	return entry
}

// Validate verifies the necessary fields for
// the Repo type are populated correctly.
func (r *Repo) Validate() error {
	// verify the UserID field is populated
	if r.UserID.Int64 <= 0 {
		return ErrEmptyRepoUserID
	}

	// verify the Org field is populated
	if len(r.Org.String) == 0 {
		return ErrEmptyRepoOrg
	}

	// verify the Name field is populated
	if len(r.Name.String) == 0 {
		return ErrEmptyRepoName
	}

	// verify the FullName field is populated
	if len(r.FullName.String) == 0 {
		return ErrEmptyRepoFullName
	}

	return nil
}
