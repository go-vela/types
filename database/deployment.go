// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-vela/types/library"
	"github.com/go-vela/types/raw"
	"github.com/lib/pq"
)

var (
	// ErrEmptyHookRepoID defines the error type when a
	// Hook type has an empty RepoID field provided.
	ErrEmptyDeploymentNumber = errors.New("empty deployment number provided")

	// ErrEmptyHookSourceID defines the error type when a
	// Hook type has an empty SourceID field provided.
	ErrEmptyDeploymentRepoID = errors.New("empty deployment repo_id provided")
)

// Deployment is the database representation of a deployment for a repo.
type Deployment struct {
	ID          sql.NullInt64      `sql:"id"`
	Number      sql.NullInt64      `sql:"number"`
	RepoID      sql.NullInt64      `sql:"repo_id"`
	URL         sql.NullString     `sql:"url"`
	User        sql.NullString     `sql:"user"`
	Commit      sql.NullString     `sql:"commit"`
	Ref         sql.NullString     `sql:"ref"`
	Task        sql.NullString     `sql:"task"`
	Target      sql.NullString     `sql:"target"`
	Description sql.NullString     `sql:"description"`
	Payload     raw.StringSliceMap `sql:"payload"`
	Builds      pq.StringArray     `sql:"builds" gorm:"type:varchar(50)"`
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the Hook type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (d *Deployment) Nullify() *Deployment {
	if d == nil {
		return nil
	}

	// check if the ID field should be false
	if d.ID.Int64 == 0 {
		d.ID.Valid = false
	}

	// check if the Number field should be false
	if d.Number.Int64 == 0 {
		d.Number.Valid = false
	}

	// check if the RepoID field should be false
	if d.RepoID.Int64 == 0 {
		d.RepoID.Valid = false
	}

	// check if the URL field should be false
	if len(d.URL.String) == 0 {
		d.URL.Valid = false
	}

	// check if the User field should be false
	if len(d.User.String) == 0 {
		d.User.Valid = false
	}

	// check if the Commit field should be false
	if len(d.Commit.String) == 0 {
		d.Commit.Valid = false
	}

	// check if the Ref field should be false
	if len(d.Ref.String) == 0 {
		d.Ref.Valid = false
	}

	// check if the Task field should be false
	if len(d.Task.String) == 0 {
		d.Task.Valid = false
	}

	// check if the Target field should be false
	if len(d.Target.String) == 0 {
		d.Target.Valid = false
	}

	// check if the Description field should be false
	if len(d.Description.String) == 0 {
		d.Description.Valid = false
	}

	return d
}

// ToLibrary converts the Deployment type
// to a library Deployment type.
func (d *Deployment) ToLibrary(builds *[]library.Build) *library.Deployment {
	deployment := new(library.Deployment)

	deployment.SetID(d.ID.Int64)
	deployment.SetNumber(d.Number.Int64)
	deployment.SetRepoID(d.RepoID.Int64)
	deployment.SetURL(d.URL.String)
	deployment.SetUser(d.User.String)
	deployment.SetCommit(d.Commit.String)
	deployment.SetRef(d.Ref.String)
	deployment.SetTask(d.Task.String)
	deployment.SetTarget(d.Target.String)
	deployment.SetDescription(d.Description.String)
	deployment.SetPayload(d.Payload)
	deployment.SetBuilds(builds)

	return deployment
}

// Validate verifies the necessary fields for
// the Deplotment type are populated correctly.
func (d *Deployment) Validate() error {
	// verify the RepoID field is populated
	if d.RepoID.Int64 <= 0 {
		return ErrEmptyDeploymentRepoID
	}

	// verify the Number field is populated
	if d.Number.Int64 <= 0 {
		return ErrEmptyDeploymentNumber
	}

	// ensure that all Hook string fields
	// that can be returned as JSON are sanitized
	// to avoid unsafe HTML content
	d.User = sql.NullString{String: sanitize(d.User.String), Valid: d.User.Valid}
	d.Commit = sql.NullString{String: sanitize(d.Commit.String), Valid: d.Commit.Valid}
	d.Ref = sql.NullString{String: sanitize(d.Ref.String), Valid: d.Ref.Valid}
	d.Task = sql.NullString{String: sanitize(d.Task.String), Valid: d.Task.Valid}
	d.Target = sql.NullString{String: sanitize(d.Target.String), Valid: d.Target.Valid}
	d.Description = sql.NullString{String: sanitize(d.Description.String), Valid: d.Description.Valid}

	return nil
}

// DeploymentFromLibrary converts the Deployment type
// to a library Deployment type.
func DeploymentFromLibrary(d *library.Deployment) *Deployment {
	buildIDs := []string{}
	for _, build := range d.GetBuilds() {
		buildIDs = append(buildIDs, fmt.Sprint(build.GetID()))
	}

	deployment := &Deployment{
		ID:          sql.NullInt64{Int64: d.GetID(), Valid: true},
		Number:      sql.NullInt64{Int64: d.GetNumber(), Valid: true},
		RepoID:      sql.NullInt64{Int64: d.GetRepoID(), Valid: true},
		URL:         sql.NullString{String: d.GetURL(), Valid: true},
		User:        sql.NullString{String: d.GetUser(), Valid: true},
		Commit:      sql.NullString{String: d.GetCommit(), Valid: true},
		Ref:         sql.NullString{String: d.GetRef(), Valid: true},
		Task:        sql.NullString{String: d.GetTask(), Valid: true},
		Target:      sql.NullString{String: d.GetTarget(), Valid: true},
		Description: sql.NullString{String: d.GetDescription(), Valid: true},
		Payload:     d.GetPayload(),
		Builds:      buildIDs,
	}

	return deployment.Nullify()
}
