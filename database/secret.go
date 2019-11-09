// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"

	"github.com/lib/pq"
)

var (
	// ErrEmptySecretName defines the error type when a
	// Secret type has an empty Name field provided.
	ErrEmptySecretName = errors.New("empty secret name provided")

	// ErrEmptySecretOrg defines the error type when a
	// Secret type has an empty Org field provided.
	ErrEmptySecretOrg = errors.New("empty secret org provided")

	// ErrEmptySecretRepo defines the error type when a
	// Secret type has an empty Repo field provided.
	ErrEmptySecretRepo = errors.New("empty secret repo provided")

	// ErrEmptySecretTeam defines the error type when a
	// Secret type has an empty Team field provided.
	ErrEmptySecretTeam = errors.New("empty secret team provided")

	// ErrEmptySecretType defines the error type when a
	// Secret type has an empty Type field provided.
	ErrEmptySecretType = errors.New("empty secret type provided")

	// ErrEmptySecretValue defines the error type when a
	// Secret type has an empty Value field provided.
	ErrEmptySecretValue = errors.New("empty secret value provided")
)

// Secret is the database representation of a secret.
type Secret struct {
	ID     sql.NullInt64  `sql:"id"`
	Org    sql.NullString `sql:"org"`
	Repo   sql.NullString `sql:"repo"`
	Team   sql.NullString `sql:"team"`
	Name   sql.NullString `sql:"name"`
	Value  sql.NullString `sql:"value"`
	Type   sql.NullString `sql:"type"`
	Images pq.StringArray `sql:"images"`
	Events pq.StringArray `sql:"events"`
	Commands sql.NullBool `sql:"commands"`
}

// ToLibrary converts the Secret type
// to a library Secret type.
func (s *Secret) ToLibrary() *library.Secret {
	images := []string(s.Images)
	events := []string(s.Events)

	return &library.Secret{
		ID:     &s.ID.Int64,
		Org:    &s.Org.String,
		Repo:   &s.Repo.String,
		Team:   &s.Team.String,
		Name:   &s.Name.String,
		Value:  &s.Value.String,
		Type:   &s.Type.String,
		Images: &images,
		Events: &events,
		Commands: &s.Commands.Bool,
	}
}

// Nullify is a helper function to overwrite fields in the
// secret to ensure the valid flag is properly set for a sqlnull type.
func (s *Secret) nullify() {
	// check if the ID should be false
	if s.ID.Int64 == 0 {
		s.ID.Valid = false
	}

	// check if the Org should be false
	if strings.EqualFold(s.Org.String, "") {
		s.Org.Valid = false
	}

	// check if the Repo should be false
	if strings.EqualFold(s.Repo.String, "") {
		s.Repo.Valid = false
	}

	// check if the Team should be false
	if strings.EqualFold(s.Team.String, "") {
		s.Team.Valid = false
	}

	// check if the Name should be false
	if strings.EqualFold(s.Name.String, "") {
		s.Team.Valid = false
	}

	// check if the Value should be false
	if strings.EqualFold(s.Value.String, "") {
		s.Value.Valid = false
	}

	// check if the Value should be false
	if strings.EqualFold(s.Type.String, "") {
		s.Type.Valid = false
	}
}

// SecretFromLibrary converts the library Secret type
// to a database Secret type.
func SecretFromLibrary(s *library.Secret) *Secret {
	entry := &Secret{
		ID:     sql.NullInt64{Int64: s.GetID(), Valid: true},
		Org:    sql.NullString{String: s.GetOrg(), Valid: true},
		Repo:   sql.NullString{String: s.GetRepo(), Valid: true},
		Team:   sql.NullString{String: s.GetTeam(), Valid: true},
		Name:   sql.NullString{String: s.GetName(), Valid: true},
		Value:  sql.NullString{String: s.GetValue(), Valid: true},
		Type:   sql.NullString{String: s.GetType(), Valid: true},
		Images: s.GetImages(),
		Events: s.GetEvents(),
		Commands: sql.NullBool{Bool: s.GetCommands(), Valid: true},
	}

	entry.nullify()

	return entry
}

// Validate verifies the necessary fields for
// the Secret type are populated correctly.
func (s *Secret) Validate() error {
	// verify the Type field is populated
	if len(s.Type.String) == 0 {
		return ErrEmptySecretType
	}

	// verify the Org field is populated
	if len(s.Org.String) == 0 {
		return ErrEmptySecretOrg
	}

	// check if an org or repo secret
	if strings.EqualFold(s.Type.String, constants.SecretRepo) ||
		strings.EqualFold(s.Type.String, constants.SecretOrg) {
		// verify the Repo field is populated
		if len(s.Repo.String) == 0 {
			return ErrEmptySecretRepo
		}
	}

	// check if a shared secret
	if strings.EqualFold(s.Type.String, constants.SecretShared) {
		// verify the Team field is populated
		if len(s.Team.String) == 0 {
			return ErrEmptySecretTeam
		}
	}

	// verify the Name field is populated
	if len(s.Name.String) == 0 {
		return ErrEmptySecretName
	}

	// verify the Value field is populated
	if len(s.Value.String) == 0 {
		return ErrEmptySecretValue
	}

	return nil
}
