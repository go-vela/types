// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"io"

	"github.com/go-vela/types/library"
)

var (
	// ErrEmptyRepoFullName defines the error type when a
	// Repo type has an empty FullName field provided.
	ErrEmptyRepoFullName = errors.New("empty repo full_name provided")

	// ErrEmptyRepoHash defines the error type when a
	// Repo type has an empty Hash field provided.
	ErrEmptyRepoHash = errors.New("empty repo hash provided")

	// ErrEmptyRepoName defines the error type when a
	// Repo type has an empty Name field provided.
	ErrEmptyRepoName = errors.New("empty repo name provided")

	// ErrEmptyRepoOrg defines the error type when a
	// Repo type has an empty Org field provided.
	ErrEmptyRepoOrg = errors.New("empty repo org provided")

	// ErrEmptyRepoUserID defines the error type when a
	// Repo type has an empty UserID field provided.
	ErrEmptyRepoUserID = errors.New("empty repo user_id provided")

	// ErrEmptyRepoVisibility defines the error type when a
	// Repo type has an empty Visibility field provided.
	ErrEmptyRepoVisibility = errors.New("empty repo visibility provided")
)

// Repo is the database representation of a repo.
type Repo struct {
	ID           sql.NullInt64  `sql:"id"`
	UserID       sql.NullInt64  `sql:"user_id"`
	Hash         sql.NullString `sql:"hash"`
	Org          sql.NullString `sql:"org"`
	Name         sql.NullString `sql:"name"`
	FullName     sql.NullString `sql:"full_name"`
	Link         sql.NullString `sql:"link"`
	Clone        sql.NullString `sql:"clone"`
	Branch       sql.NullString `sql:"branch"`
	Timeout      sql.NullInt64  `sql:"timeout"`
	Visibility   sql.NullString `sql:"visibility"`
	Private      sql.NullBool   `sql:"private"`
	Trusted      sql.NullBool   `sql:"trusted"`
	Active       sql.NullBool   `sql:"active"`
	AllowPull    sql.NullBool   `sql:"allow_pull"`
	AllowPush    sql.NullBool   `sql:"allow_push"`
	AllowDeploy  sql.NullBool   `sql:"allow_deploy"`
	AllowTag     sql.NullBool   `sql:"allow_tag"`
	AllowComment sql.NullBool   `sql:"allow_comment"`
}

// Decrypt will manipulate the existing repo hash by
// base64 decoding that value. Then, a AES-256 cipher
// block is created from the encryption key in order to
// decrypt the base64 decoded secret value.
func (r *Repo) Decrypt(key string) error {
	// base64 decode the encrypted repo hash
	decoded, err := base64.StdEncoding.DecodeString(r.Hash.String)
	if err != nil {
		return err
	}

	// create a new cipher block from the encryption key
	//
	// the key should have a length of 64 bits to ensure
	// we are using the AES-256 standard
	//
	// https://en.wikipedia.org/wiki/Advanced_Encryption_Standard
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	// creates a new Galois Counter Mode cipher block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// nonce is an arbitrary number used to to ensure that
	// old communications cannot be reused in replay attacks.
	//
	// https://en.wikipedia.org/wiki/Cryptographic_nonce
	nonceSize := gcm.NonceSize()

	// verify the decoded repo hash is greater than nonce
	//
	// if the base64 decoded repo hash is less than the
	// nonce size, then we can reasonably assume the repo
	// hasn't been encrypted yet.
	if len(decoded) < nonceSize {
		return fmt.Errorf("invalid length for decoded repo hash")
	}

	// capture nonce and ciphertext from decoded repo hash
	nonce, ciphertext := decoded[:nonceSize], decoded[nonceSize:]

	// decrypt the decoded repo hash from the ciphertext
	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	// set the decrypted repo hash
	r.Hash = sql.NullString{String: string(decrypted), Valid: true}

	return nil
}

// Encrypt will manipulate the existing repo hash by
// creating a AES-256 cipher block from the encryption
// key in order to encrypt the repo hash. Then, the
// repo hash is base64 encoded for transport across
// network boundaries.
func (r *Repo) Encrypt(key string) error {
	// create a new cipher block from the encryption key
	//
	// the key should have a length of 64 bits to ensure
	// we are using the AES-256 standard
	//
	// https://en.wikipedia.org/wiki/Advanced_Encryption_Standard
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	// creates a new Galois Counter Mode cipher block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// nonce is an arbitrary number used to to ensure that
	// old communications cannot be reused in replay attacks.
	//
	// https://en.wikipedia.org/wiki/Cryptographic_nonce
	nonce := make([]byte, gcm.NonceSize())

	// set nonce from a cryptographically secure random number generator
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return err
	}

	// encrypt the data with the randomly generated nonce
	encrypted := gcm.Seal(nonce, nonce, []byte(r.Hash.String), nil)

	// base64 encode the encrypted repo hash to make it network safe
	r.Hash = sql.NullString{String: base64.StdEncoding.EncodeToString(encrypted), Valid: true}

	return nil
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the Repo type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (r *Repo) Nullify() *Repo {
	if r == nil {
		return nil
	}

	// check if the ID field should be false
	if r.ID.Int64 == 0 {
		r.ID.Valid = false
	}

	// check if the UserID field should be false
	if r.UserID.Int64 == 0 {
		r.UserID.Valid = false
	}

	// check if the Hash field should be false
	if len(r.Hash.String) == 0 {
		r.Hash.Valid = false
	}

	// check if the Org field should be false
	if len(r.Org.String) == 0 {
		r.Org.Valid = false
	}

	// check if the Name field should be false
	if len(r.Name.String) == 0 {
		r.Name.Valid = false
	}

	// check if the FullName field should be false
	if len(r.FullName.String) == 0 {
		r.FullName.Valid = false
	}

	// check if the Link field should be false
	if len(r.Link.String) == 0 {
		r.Link.Valid = false
	}

	// check if the Clone field should be false
	if len(r.Clone.String) == 0 {
		r.Clone.Valid = false
	}

	// check if the Branch field should be false
	if len(r.Branch.String) == 0 {
		r.Branch.Valid = false
	}

	// check if the Timeout field should be false
	if r.Timeout.Int64 == 0 {
		r.Timeout.Valid = false
	}

	// check if the Visibility field should be false
	if len(r.Visibility.String) == 0 {
		r.Visibility.Valid = false
	}

	return r
}

// ToLibrary converts the Repo type
// to a library Repo type.
func (r *Repo) ToLibrary() *library.Repo {
	repo := new(library.Repo)

	repo.SetID(r.ID.Int64)
	repo.SetUserID(r.UserID.Int64)
	repo.SetHash(r.Hash.String)
	repo.SetOrg(r.Org.String)
	repo.SetName(r.Name.String)
	repo.SetFullName(r.FullName.String)
	repo.SetLink(r.Link.String)
	repo.SetClone(r.Clone.String)
	repo.SetBranch(r.Branch.String)
	repo.SetTimeout(r.Timeout.Int64)
	repo.SetVisibility(r.Visibility.String)
	repo.SetPrivate(r.Private.Bool)
	repo.SetTrusted(r.Trusted.Bool)
	repo.SetActive(r.Active.Bool)
	repo.SetAllowPull(r.AllowPull.Bool)
	repo.SetAllowPush(r.AllowPush.Bool)
	repo.SetAllowDeploy(r.AllowDeploy.Bool)
	repo.SetAllowTag(r.AllowTag.Bool)
	repo.SetAllowComment(r.AllowComment.Bool)

	return repo
}

// Validate verifies the necessary fields for
// the Repo type are populated correctly.
func (r *Repo) Validate() error {
	// verify the UserID field is populated
	if r.UserID.Int64 <= 0 {
		return ErrEmptyRepoUserID
	}

	// verify the Hash field is populated
	if len(r.Hash.String) == 0 {
		return ErrEmptyRepoHash
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

	// verify the Visibility field is populated
	if len(r.Visibility.String) == 0 {
		return ErrEmptyRepoVisibility
	}

	return nil
}

// RepoFromLibrary converts the library Repo type
// to a database repo type.
func RepoFromLibrary(r *library.Repo) *Repo {
	repo := &Repo{
		ID:           sql.NullInt64{Int64: r.GetID(), Valid: true},
		UserID:       sql.NullInt64{Int64: r.GetUserID(), Valid: true},
		Hash:         sql.NullString{String: r.GetHash(), Valid: true},
		Org:          sql.NullString{String: r.GetOrg(), Valid: true},
		Name:         sql.NullString{String: r.GetName(), Valid: true},
		FullName:     sql.NullString{String: r.GetFullName(), Valid: true},
		Link:         sql.NullString{String: r.GetLink(), Valid: true},
		Clone:        sql.NullString{String: r.GetClone(), Valid: true},
		Branch:       sql.NullString{String: r.GetBranch(), Valid: true},
		Timeout:      sql.NullInt64{Int64: r.GetTimeout(), Valid: true},
		Visibility:   sql.NullString{String: r.GetVisibility(), Valid: true},
		Private:      sql.NullBool{Bool: r.GetPrivate(), Valid: true},
		Trusted:      sql.NullBool{Bool: r.GetTrusted(), Valid: true},
		Active:       sql.NullBool{Bool: r.GetActive(), Valid: true},
		AllowPull:    sql.NullBool{Bool: r.GetAllowPull(), Valid: true},
		AllowPush:    sql.NullBool{Bool: r.GetAllowPush(), Valid: true},
		AllowDeploy:  sql.NullBool{Bool: r.GetAllowDeploy(), Valid: true},
		AllowTag:     sql.NullBool{Bool: r.GetAllowTag(), Valid: true},
		AllowComment: sql.NullBool{Bool: r.GetAllowComment(), Valid: true},
	}

	return repo.Nullify()
}
