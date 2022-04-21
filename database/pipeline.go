// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"bytes"
	"compress/zlib"
	"database/sql"
	"errors"
	"io/ioutil"

	"github.com/go-vela/types/library"
	"github.com/sirupsen/logrus"
)

var (
	// ErrEmptyPipelineCommit defines the error type when a
	// Pipeline type has an empty Commit field provided.
	ErrEmptyPipelineCommit = errors.New("empty pipeline commit provided")

	// ErrEmptyPipelineRef defines the error type when a
	// Pipeline type has an empty Ref field provided.
	ErrEmptyPipelineRef = errors.New("empty pipeline ref provided")

	// ErrEmptyPipelineRepoID defines the error type when a
	// Pipeline type has an empty RepoID field provided.
	ErrEmptyPipelineRepoID = errors.New("empty pipeline repo_id provided")

	// ErrEmptyPipelineType defines the error type when a
	// Pipeline type has an empty Type field provided.
	ErrEmptyPipelineType = errors.New("empty pipeline type provided")

	// ErrEmptyPipelineVersion defines the error type when a
	// Pipeline type has an empty Version field provided.
	ErrEmptyPipelineVersion = errors.New("empty pipeline version provided")
)

// Pipeline is the database representation of a pipeline.
type Pipeline struct {
	ID              sql.NullInt64  `sql:"id"`
	RepoID          sql.NullInt64  `sql:"repo_id"`
	Commit          sql.NullString `sql:"commit"`
	Flavor          sql.NullString `sql:"flavor"`
	Platform        sql.NullString `sql:"platform"`
	Ref             sql.NullString `sql:"ref"`
	Type            sql.NullString `sql:"type"`
	Version         sql.NullString `sql:"version"`
	ExternalSecrets sql.NullBool   `sql:"external_secrets"`
	InternalSecrets sql.NullBool   `sql:"internal_secrets"`
	Services        sql.NullBool   `sql:"services"`
	Stages          sql.NullBool   `sql:"stages"`
	Steps           sql.NullBool   `sql:"steps"`
	Templates       sql.NullBool   `sql:"templates"`
	Data            []byte         `sql:"data"`
}

// Compress will manipulate the existing data for the
// pipeline by compressing that data. This produces
// a significantly smaller amount of data that is
// required to store in the system.
func (p *Pipeline) Compress(level int) error {
	// create new buffer for storing compressed pipeline data
	b := new(bytes.Buffer)

	// create new writer for writing compressed pipeline data
	w, err := zlib.NewWriterLevel(b, level)
	if err != nil {
		return err
	}

	// write compressed pipeline data to buffer
	_, err = w.Write(p.Data)
	if err != nil {
		return err
	}

	// close the writer
	//
	// compressed bytes are not flushed until the
	// writer is closed or explicitly flushed
	err = w.Close()
	if err != nil {
		logrus.Errorf("unable to close compression buffer: %v", err)
	}

	// overwrite database pipeline data with compressed pipeline data
	p.Data = b.Bytes()

	return nil
}

// Decompress will manipulate the existing data for the
// pipeline by decompressing that data. This allows us
// to have a significantly smaller amount of data that is
// stored in the system.
func (p *Pipeline) Decompress() error {
	// create new buffer from the compressed pipeline data
	b := bytes.NewBuffer(p.Data)

	// create new reader for reading the compressed pipeline data
	r, err := zlib.NewReader(b)
	if err != nil {
		return err
	}

	// defer closing the reader
	defer r.Close()

	// capture decompressed pipeline data from the compressed pipeline data
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	// overwrite compressed pipeline data with decompressed pipeline data
	p.Data = data

	return nil
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the Pipeline type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (p *Pipeline) Nullify() *Pipeline {
	if p == nil {
		return nil
	}

	// check if the ID field should be false
	if p.ID.Int64 == 0 {
		p.ID.Valid = false
	}

	// check if the RepoID field should be false
	if p.RepoID.Int64 == 0 {
		p.RepoID.Valid = false
	}

	// check if the Commit field should be false
	if len(p.Commit.String) == 0 {
		p.Commit.Valid = false
	}

	// check if the Flavor field should be false
	if len(p.Flavor.String) == 0 {
		p.Flavor.Valid = false
	}

	// check if the Platform field should be false
	if len(p.Platform.String) == 0 {
		p.Platform.Valid = false
	}

	// check if the Ref field should be false
	if len(p.Ref.String) == 0 {
		p.Ref.Valid = false
	}

	// check if the Type field should be false
	if len(p.Type.String) == 0 {
		p.Type.Valid = false
	}

	// check if the Version field should be false
	if len(p.Version.String) == 0 {
		p.Version.Valid = false
	}

	return p
}

// ToLibrary converts the Pipeline type
// to a library Pipeline type.
func (p *Pipeline) ToLibrary() *library.Pipeline {
	pipeline := new(library.Pipeline)

	pipeline.SetID(p.ID.Int64)
	pipeline.SetRepoID(p.RepoID.Int64)
	pipeline.SetCommit(p.Commit.String)
	pipeline.SetFlavor(p.Flavor.String)
	pipeline.SetPlatform(p.Platform.String)
	pipeline.SetRef(p.Ref.String)
	pipeline.SetType(p.Type.String)
	pipeline.SetVersion(p.Version.String)
	pipeline.SetExternalSecrets(p.ExternalSecrets.Bool)
	pipeline.SetInternalSecrets(p.InternalSecrets.Bool)
	pipeline.SetServices(p.Services.Bool)
	pipeline.SetStages(p.Stages.Bool)
	pipeline.SetSteps(p.Steps.Bool)
	pipeline.SetTemplates(p.Templates.Bool)
	pipeline.SetData(p.Data)

	return pipeline
}

// Validate verifies the necessary fields for
// the Pipeline type are populated correctly.
func (p *Pipeline) Validate() error {
	// verify the Commit field is populated
	if len(p.Commit.String) == 0 {
		return ErrEmptyPipelineCommit
	}

	// verify the Ref field is populated
	if len(p.Ref.String) == 0 {
		return ErrEmptyPipelineRef
	}

	// verify the RepoID field is populated
	if p.RepoID.Int64 <= 0 {
		return ErrEmptyPipelineRepoID
	}

	// verify the Type field is populated
	if len(p.Type.String) == 0 {
		return ErrEmptyPipelineType
	}

	// verify the Version field is populated
	if len(p.Version.String) == 0 {
		return ErrEmptyPipelineVersion
	}

	// ensure that all Pipeline string fields
	// that can be returned as JSON are sanitized
	// to avoid unsafe HTML content
	p.Commit = sql.NullString{String: sanitize(p.Commit.String), Valid: p.Commit.Valid}
	p.Flavor = sql.NullString{String: sanitize(p.Flavor.String), Valid: p.Flavor.Valid}
	p.Platform = sql.NullString{String: sanitize(p.Platform.String), Valid: p.Platform.Valid}
	p.Ref = sql.NullString{String: sanitize(p.Ref.String), Valid: p.Ref.Valid}
	p.Type = sql.NullString{String: sanitize(p.Type.String), Valid: p.Type.Valid}
	p.Version = sql.NullString{String: sanitize(p.Version.String), Valid: p.Version.Valid}

	return nil
}

// PipelineFromLibrary converts the library Pipeline type
// to a database Pipeline type.
func PipelineFromLibrary(p *library.Pipeline) *Pipeline {
	pipeline := &Pipeline{
		ID:              sql.NullInt64{Int64: p.GetID(), Valid: true},
		RepoID:          sql.NullInt64{Int64: p.GetRepoID(), Valid: true},
		Commit:          sql.NullString{String: p.GetCommit(), Valid: true},
		Flavor:          sql.NullString{String: p.GetFlavor(), Valid: true},
		Platform:        sql.NullString{String: p.GetPlatform(), Valid: true},
		Ref:             sql.NullString{String: p.GetRef(), Valid: true},
		Type:            sql.NullString{String: p.GetType(), Valid: true},
		Version:         sql.NullString{String: p.GetVersion(), Valid: true},
		ExternalSecrets: sql.NullBool{Bool: p.GetExternalSecrets(), Valid: true},
		InternalSecrets: sql.NullBool{Bool: p.GetInternalSecrets(), Valid: true},
		Services:        sql.NullBool{Bool: p.GetServices(), Valid: true},
		Stages:          sql.NullBool{Bool: p.GetStages(), Valid: true},
		Steps:           sql.NullBool{Bool: p.GetSteps(), Valid: true},
		Templates:       sql.NullBool{Bool: p.GetTemplates(), Valid: true},
		Data:            p.GetData(),
	}

	return pipeline.Nullify()
}
