// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"encoding/base64"
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
	ID           sql.NullInt64  `sql:"id"`
	Org          sql.NullString `sql:"org"`
	Repo         sql.NullString `sql:"repo"`
	Team         sql.NullString `sql:"team"`
	Name         sql.NullString `sql:"name"`
	Value        sql.NullString `sql:"value"`
	Type         sql.NullString `sql:"type"`
	Images       pq.StringArray `sql:"images"`
	Events       pq.StringArray `sql:"events"`
	AllowCommand sql.NullBool   `sql:"allow_command"`
}

// Decrypt will manipulate the existing secret value by
// base64 decoding that value. Then, a AES-256 cipher
// block is created from the encryption key in order to
// decrypt the base64 decoded secret value.
func (s *Secret) Decrypt(key string) error {
	// base64 decode the encrypted secret value
	decoded, err := base64.StdEncoding.DecodeString(s.Value.String)
	if err != nil {
		return err
	}

	// decrypt the base64 decoded secret value
	decrypted, err := decrypt(key, decoded)
	if err != nil {
		return err
	}

	// set the decrypted secret value
	s.Value = sql.NullString{
		String: string(decrypted),
		Valid:  true,
	}

	return nil
}

// Encrypt will manipulate the existing secret value by
// creating a AES-256 cipher block from the encryption
// key in order to encrypt the secret value. Then, the
// secret value is base64 encoded for transport across
// network boundaries.
func (s *Secret) Encrypt(key string) error {
	// encrypt the secret value
	encrypted, err := encrypt(key, []byte(s.Value.String))
	if err != nil {
		return err
	}

	// base64 encode the encrypted secret data to make it network safe
	s.Value = sql.NullString{
		String: base64.StdEncoding.EncodeToString(encrypted),
		Valid:  true,
	}

	return nil
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the Secret type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (s *Secret) Nullify() *Secret {
	if s == nil {
		return nil
	}

	// check if the ID field should be false
	if s.ID.Int64 == 0 {
		s.ID.Valid = false
	}

	// check if the Org field should be false
	if len(s.Org.String) == 0 {
		s.Org.Valid = false
	}

	// check if the Repo field should be false
	if len(s.Repo.String) == 0 {
		s.Repo.Valid = false
	}

	// check if the Team field should be false
	if len(s.Team.String) == 0 {
		s.Team.Valid = false
	}

	// check if the Name field should be false
	if len(s.Name.String) == 0 {
		s.Name.Valid = false
	}

	// check if the Value field should be false
	if len(s.Value.String) == 0 {
		s.Value.Valid = false
	}

	// check if the Value should be false
	if len(s.Type.String) == 0 {
		s.Type.Valid = false
	}

	return s
}

// ToLibrary converts the Secret type
// to a library Secret type.
func (s *Secret) ToLibrary() *library.Secret {
	secret := new(library.Secret)

	secret.SetID(s.ID.Int64)
	secret.SetOrg(s.Org.String)
	secret.SetRepo(s.Repo.String)
	secret.SetTeam(s.Team.String)
	secret.SetName(s.Name.String)
	secret.SetValue(s.Value.String)
	secret.SetType(s.Type.String)
	secret.SetImages(s.Images)
	secret.SetEvents(s.Events)
	secret.SetAllowCommand(s.AllowCommand.Bool)

	return secret
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

	// ensure that all Secret string fields
	// that can be returned as JSON are sanitized
	// to avoid unsafe HTML content
	s.Org = sql.NullString{String: sanitize(s.Org.String), Valid: true}
	s.Repo = sql.NullString{String: sanitize(s.Repo.String), Valid: true}
	s.Team = sql.NullString{String: sanitize(s.Team.String), Valid: true}
	s.Name = sql.NullString{String: sanitize(s.Name.String), Valid: true}
	s.Type = sql.NullString{String: sanitize(s.Type.String), Valid: true}

	// ensure that all Images are sanitized
	// to avoid unsafe HTML content
	for i, v := range s.Images {
		s.Images[i] = sanitize(v)
	}

	// ensure that all Events are sanitized
	// to avoid unsafe HTML content
	for i, v := range s.Events {
		s.Events[i] = sanitize(v)
	}

	return nil
}

// SecretFromLibrary converts the library Secret type
// to a database Secret type.
func SecretFromLibrary(s *library.Secret) *Secret {
	secret := &Secret{
		ID:           sql.NullInt64{Int64: s.GetID(), Valid: true},
		Org:          sql.NullString{String: s.GetOrg(), Valid: true},
		Repo:         sql.NullString{String: s.GetRepo(), Valid: true},
		Team:         sql.NullString{String: s.GetTeam(), Valid: true},
		Name:         sql.NullString{String: s.GetName(), Valid: true},
		Value:        sql.NullString{String: s.GetValue(), Valid: true},
		Type:         sql.NullString{String: s.GetType(), Valid: true},
		Images:       s.GetImages(),
		Events:       s.GetEvents(),
		AllowCommand: sql.NullBool{Bool: s.GetAllowCommand(), Valid: true},
	}

	return secret.Nullify()
}
