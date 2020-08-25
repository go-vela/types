// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"errors"

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
	Email        sql.NullString `sql:"email"`
	Link         sql.NullString `sql:"link"`
	Branch       sql.NullString `sql:"branch"`
	Ref          sql.NullString `sql:"ref"`
	BaseRef      sql.NullString `sql:"base_ref"`
	HeadRef      sql.NullString `sql:"head_ref"`
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

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the Build type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (b *Build) Nullify() *Build {
	if b == nil {
		return nil
	}

	// check if the ID field should be false
	if b.ID.Int64 == 0 {
		b.ID.Valid = false
	}

	// check if the RepoID field should be false
	if b.RepoID.Int64 == 0 {
		b.RepoID.Valid = false
	}

	// check if the Number field should be false
	if b.Number.Int32 == 0 {
		b.Number.Valid = false
	}

	// check if the Parent field should be false
	if b.Parent.Int32 == 0 {
		b.Parent.Valid = false
	}

	// check if the Event field should be false
	if len(b.Event.String) == 0 {
		b.Event.Valid = false
	}

	// check if the Status field should be false
	if len(b.Status.String) == 0 {
		b.Status.Valid = false
	}

	// check if the Error field should be false
	if len(b.Error.String) == 0 {
		b.Error.Valid = false
	}

	// check if the Enqueued field should be false
	if b.Enqueued.Int64 == 0 {
		b.Enqueued.Valid = false
	}

	// check if the Created field should be false
	if b.Created.Int64 == 0 {
		b.Created.Valid = false
	}

	// check if the Started field should be false
	if b.Started.Int64 == 0 {
		b.Started.Valid = false
	}

	// check if the Finished field should be false
	if b.Finished.Int64 == 0 {
		b.Finished.Valid = false
	}

	// check if the Deploy field should be false
	if len(b.Deploy.String) == 0 {
		b.Deploy.Valid = false
	}

	// check if the Clone field should be false
	if len(b.Clone.String) == 0 {
		b.Clone.Valid = false
	}

	// check if the Source field should be false
	if len(b.Source.String) == 0 {
		b.Source.Valid = false
	}

	// check if the Title field should be false
	if len(b.Title.String) == 0 {
		b.Title.Valid = false
	}

	// check if the Message field should be false
	if len(b.Message.String) == 0 {
		b.Message.Valid = false
	}

	// check if the Commit field should be false
	if len(b.Commit.String) == 0 {
		b.Commit.Valid = false
	}

	// check if the Sender field should be false
	if len(b.Sender.String) == 0 {
		b.Sender.Valid = false
	}

	// check if the Author field should be false
	if len(b.Author.String) == 0 {
		b.Author.Valid = false
	}

	// check if the Email field should be false
	if len(b.Email.String) == 0 {
		b.Email.Valid = false
	}

	// check if the Link field should be false
	if len(b.Link.String) == 0 {
		b.Link.Valid = false
	}

	// check if the Branch field should be false
	if len(b.Branch.String) == 0 {
		b.Branch.Valid = false
	}

	// check if the Ref field should be false
	if len(b.Ref.String) == 0 {
		b.Ref.Valid = false
	}

	// check if the BaseRef field should be false
	if len(b.BaseRef.String) == 0 {
		b.BaseRef.Valid = false
	}

	// check if the HeadRef field should be false
	if len(b.HeadRef.String) == 0 {
		b.HeadRef.Valid = false
	}

	// check if the Host field should be false
	if len(b.Host.String) == 0 {
		b.Host.Valid = false
	}

	// check if the Runtime field should be false
	if len(b.Runtime.String) == 0 {
		b.Runtime.Valid = false
	}

	// check if the Distribution field should be false
	if len(b.Distribution.String) == 0 {
		b.Distribution.Valid = false
	}

	return b
}

// ToLibrary converts the Build type
// to a library Build type.
func (b *Build) ToLibrary() *library.Build {
	build := new(library.Build)

	build.SetID(b.ID.Int64)
	build.SetRepoID(b.RepoID.Int64)
	build.SetNumber(int(b.Number.Int32))
	build.SetParent(int(b.Parent.Int32))
	build.SetEvent(b.Event.String)
	build.SetStatus(b.Status.String)
	build.SetError(b.Error.String)
	build.SetEnqueued(b.Enqueued.Int64)
	build.SetCreated(b.Created.Int64)
	build.SetStarted(b.Started.Int64)
	build.SetFinished(b.Finished.Int64)
	build.SetDeploy(b.Deploy.String)
	build.SetClone(b.Clone.String)
	build.SetSource(b.Source.String)
	build.SetTitle(b.Title.String)
	build.SetMessage(b.Message.String)
	build.SetCommit(b.Commit.String)
	build.SetSender(b.Sender.String)
	build.SetAuthor(b.Author.String)
	build.SetEmail(b.Email.String)
	build.SetLink(b.Link.String)
	build.SetBranch(b.Branch.String)
	build.SetRef(b.Ref.String)
	build.SetBaseRef(b.BaseRef.String)
	build.SetHeadRef(b.HeadRef.String)
	build.SetHost(b.Host.String)
	build.SetRuntime(b.Runtime.String)
	build.SetDistribution(b.Distribution.String)

	return build
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

// BuildFromLibrary converts the library Build type
// to a database build type.
func BuildFromLibrary(b *library.Build) *Build {
	build := &Build{
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
		Email:        sql.NullString{String: b.GetEmail(), Valid: true},
		Link:         sql.NullString{String: b.GetLink(), Valid: true},
		Branch:       sql.NullString{String: b.GetBranch(), Valid: true},
		Ref:          sql.NullString{String: b.GetRef(), Valid: true},
		BaseRef:      sql.NullString{String: b.GetBaseRef(), Valid: true},
		HeadRef:      sql.NullString{String: b.GetHeadRef(), Valid: true},
		Host:         sql.NullString{String: b.GetHost(), Valid: true},
		Runtime:      sql.NullString{String: b.GetRuntime(), Valid: true},
		Distribution: sql.NullString{String: b.GetDistribution(), Valid: true},
	}

	return build.Nullify()
}
