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
	// ErrEmptyBuildNumber defines the error type when a
	// Build type has an empty Number field provided.
	ErrEmptyBuildNumber = errors.New("empty build number provided")

	// ErrEmptyBuildRepoID defines the error type when a
	// Build type has an empty `RepoID` field provided.
	ErrEmptyBuildRepoID = errors.New("empty build repo_id provided")
)

// Build is the database representation of a build for a pipeline.
type Build struct {
	ID           sql.NullInt64  `sql:"id"`
	RepoID       sql.NullInt64  `sql:"repo_id"`
	Number       sql.NullInt32  `sql:"number"`
	Parent       sql.NullInt32  `sql:"parent"`
	Event        sql.NullString `sql:"event"`
	Status       sql.NullString `sql:"status"`
	Error        sql.NullString `sql:"error"`
	Enqueued     sql.NullInt64  `sql:"enqueued"`
	Created      sql.NullInt64  `sql:"created"`
	Started      sql.NullInt64  `sql:"started"`
	Finished     sql.NullInt64  `sql:"finished"`
	Deploy       sql.NullString `sql:"deploy"`
	Clone        sql.NullString `sql:"clone"`
	Source       sql.NullString `sql:"source"`
	Title        sql.NullString `sql:"title"`
	Message      sql.NullString `sql:"message"`
	Commit       sql.NullString `sql:"commit"`
	Sender       sql.NullString `sql:"sender"`
	Author       sql.NullString `sql:"author"`
	Branch       sql.NullString `sql:"branch"`
	Ref          sql.NullString `sql:"ref"`
	BaseRef      sql.NullString `sql:"base_ref"`
	Host         sql.NullString `sql:"host"`
	Runtime      sql.NullString `sql:"runtime"`
	Distribution sql.NullString `sql:"distribution"`
}

// Crop prepares the Build type for inserting into the database by
// trimming values that may exceed the database column limit.
func (b *Build) Crop() *Build {
	// trim the Title field to 1000 characters
	if len(b.Title.String) > 1000 {
		b.Title = sql.NullString{String: b.Title.String[:1000], Valid: true}
	}

	// trim the Message field to 2000 characters
	if len(b.Message.String) > 2000 {
		b.Message = sql.NullString{String: b.Message.String[:2000], Valid: true}
	}

	return b
}

// ToLibrary converts the Build type
// to a library Build type.
func (b *Build) ToLibrary() *library.Build {
	n := int(b.Number.Int32)
	p := int(b.Parent.Int32)
	return &library.Build{
		ID:           &b.ID.Int64,
		RepoID:       &b.RepoID.Int64,
		Number:       &n,
		Parent:       &p,
		Event:        &b.Event.String,
		Status:       &b.Status.String,
		Error:        &b.Error.String,
		Enqueued:     &b.Enqueued.Int64,
		Created:      &b.Created.Int64,
		Started:      &b.Started.Int64,
		Finished:     &b.Finished.Int64,
		Deploy:       &b.Deploy.String,
		Clone:        &b.Clone.String,
		Source:       &b.Source.String,
		Title:        &b.Title.String,
		Message:      &b.Message.String,
		Commit:       &b.Commit.String,
		Sender:       &b.Sender.String,
		Author:       &b.Author.String,
		Branch:       &b.Branch.String,
		Ref:          &b.Ref.String,
		BaseRef:      &b.BaseRef.String,
		Host:         &b.Host.String,
		Runtime:      &b.Runtime.String,
		Distribution: &b.Distribution.String,
	}
}

// nullify is a to overwrite fields in the
// build to ensure the valid flag is properly set for a sqlnull type.
func (b *Build) nullify() {

	// check if the ID should be false
	if b.ID.Int64 == 0 {
		b.ID.Valid = false
	}

	// check if the RepoID should be false
	if b.RepoID.Int64 == 0 {
		b.RepoID.Valid = false
	}

	// check if the Number should be false
	if b.Number.Int32 == 0 {
		b.Number.Valid = false
	}

	// check if the Parent should be false
	if b.Parent.Int32 == 0 {
		b.Parent.Valid = false
	}

	// check if the Event should be false
	if strings.EqualFold(b.Event.String, "") {
		b.Event.Valid = false
	}

	// check if the Status should be false
	if strings.EqualFold(b.Status.String, "") {
		b.Status.Valid = false
	}

	// check if the Error should be false
	if strings.EqualFold(b.Error.String, "") {
		b.Error.Valid = false
	}

	// check if the Enqueued should be false
	if b.Enqueued.Int64 == 0 {
		b.Enqueued.Valid = false
	}

	// check if the Created should be false
	if b.Created.Int64 == 0 {
		b.Created.Valid = false
	}

	// check if the Started should be false
	if b.Started.Int64 == 0 {
		b.Started.Valid = false
	}

	// check if the Finished should be false
	if b.Finished.Int64 == 0 {
		b.Finished.Valid = false
	}

	// check if the Deploy should be false
	if strings.EqualFold(b.Deploy.String, "") {
		b.Deploy.Valid = false
	}

	// check if the Clone should be false
	if strings.EqualFold(b.Clone.String, "") {
		b.Clone.Valid = false
	}

	// check if the Source should be false
	if strings.EqualFold(b.Source.String, "") {
		b.Source.Valid = false
	}

	// check if the Title should be false
	if strings.EqualFold(b.Title.String, "") {
		b.Title.Valid = false
	}

	// check if the Message should be false
	if strings.EqualFold(b.Message.String, "") {
		b.Message.Valid = false
	}

	// check if the Author should be false
	if strings.EqualFold(b.Author.String, "") {
		b.Author.Valid = false
	}

	// check if the Branch should be false
	if strings.EqualFold(b.Branch.String, "") {
		b.Branch.Valid = false
	}

	// check if the Ref should be false
	if strings.EqualFold(b.Ref.String, "") {
		b.Ref.Valid = false
	}

	// check if the BaseRef should be false
	if strings.EqualFold(b.BaseRef.String, "") {
		b.BaseRef.Valid = false
	}

	// check if the Host should be false
	if strings.EqualFold(b.Host.String, "") {
		b.Host.Valid = false
	}

	// check if the Runtime should be false
	if strings.EqualFold(b.Runtime.String, "") {
		b.Runtime.Valid = false
	}

	// check if the Distribution should be false
	if strings.EqualFold(b.Distribution.String, "") {
		b.Distribution.Valid = false
	}
}

// BuildFromLibrary converts the libray Build type
// to a database build type.
func BuildFromLibrary(b *library.Build) *Build {

	entry := &Build{
		ID:           sql.NullInt64{Int64: b.GetID(), Valid: true},
		RepoID:       sql.NullInt64{Int64: b.GetRepoID(), Valid: true},
		Number:       sql.NullInt32{Int32: int32(b.GetNumber()), Valid: true},
		Parent:       sql.NullInt32{Int32: int32(b.GetParent()), Valid: true},
		Event:        sql.NullString{String: b.GetEvent(), Valid: true},
		Status:       sql.NullString{String: b.GetStatus(), Valid: true},
		Error:        sql.NullString{String: b.GetError(), Valid: true},
		Enqueued:     sql.NullInt64{Int64: b.GetEnqueued(), Valid: true},
		Created:      sql.NullInt64{Int64: b.GetCreated(), Valid: true},
		Started:      sql.NullInt64{Int64: b.GetStarted(), Valid: true},
		Finished:     sql.NullInt64{Int64: b.GetFinished(), Valid: true},
		Deploy:       sql.NullString{String: b.GetDeploy(), Valid: true},
		Clone:        sql.NullString{String: b.GetClone(), Valid: true},
		Source:       sql.NullString{String: b.GetSource(), Valid: true},
		Title:        sql.NullString{String: b.GetTitle(), Valid: true},
		Message:      sql.NullString{String: b.GetMessage(), Valid: true},
		Commit:       sql.NullString{String: b.GetCommit(), Valid: true},
		Sender:       sql.NullString{String: b.GetSender(), Valid: true},
		Author:       sql.NullString{String: b.GetAuthor(), Valid: true},
		Branch:       sql.NullString{String: b.GetBranch(), Valid: true},
		Ref:          sql.NullString{String: b.GetRef(), Valid: true},
		BaseRef:      sql.NullString{String: b.GetBaseRef(), Valid: true},
		Host:         sql.NullString{String: b.GetHost(), Valid: true},
		Runtime:      sql.NullString{String: b.GetRuntime(), Valid: true},
		Distribution: sql.NullString{String: b.GetDistribution(), Valid: true},
	}

	entry.nullify()

	return entry
}

// Validate verifies the necessary fields for
// the Build type are populated correctly.
func (b *Build) Validate() error {
	// verify the RepoID field is populated
	if b.RepoID.Int64 <= 0 {
		return ErrEmptyBuildRepoID
	}

	// verify the Number field is populated
	if b.Number.Int32 <= 0 {
		return ErrEmptyBuildNumber
	}

	return nil
}
