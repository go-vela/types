// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"bytes"
	"compress/zlib"
	"database/sql"
	"errors"
	"io/ioutil"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
	"github.com/sirupsen/logrus"
)

var (
	// ErrEmptyLogBuildID defines the error type when a
	// Log type has an empty BuildID field provided.
	ErrEmptyLogBuildID = errors.New("empty log build_id provided")

	// ErrEmptyLogRepoID defines the error type when a
	// Log type has an empty RepoID field provided.
	ErrEmptyLogRepoID = errors.New("empty log repo_id provided")

	// ErrEmptyLogStepOrServiceID defines the error type when a
	// Log type has an empty StepID or ServiceID field provided.
	ErrEmptyLogStepOrServiceID = errors.New("empty log step_id or service_id not provided")
)

// Log is the database representation of a log for a step in a build.
type Log struct {
	ID        sql.NullInt64 `sql:"id"`
	BuildID   sql.NullInt64 `sql:"build_id"`
	RepoID    sql.NullInt64 `sql:"repo_id"`
	ServiceID sql.NullInt64 `sql:"service_id"`
	StepID    sql.NullInt64 `sql:"step_id"`
	Data      []byte        `sql:"data"`
}

// Compress will manipulate the existing data for the
// log entry by compressing that data. This produces
// a significantly smaller amount of data that is
// required to store in the system.
func (l *Log) Compress() error {
	// create new buffer for storing compressed log data
	b := new(bytes.Buffer)

	// create new writer for writing compressed log data
	w, err := zlib.NewWriterLevel(b, constants.CompressionLevel)
	if err != nil {
		return err
	}

	// write compressed log data to buffer
	_, err = w.Write(l.Data)
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

	// overwrite database log data with compressed log data
	l.Data = b.Bytes()

	return nil
}

// Decompress will manipulate the existing data for the
// log entry by decompressing that data. This allows us
// to have a significantly smaller amount of data that is
// stored in the system.
func (l *Log) Decompress() error {
	// create new buffer from the compressed log data
	b := bytes.NewBuffer(l.Data)

	// create new reader for reading the compressed log data
	r, err := zlib.NewReader(b)
	if err != nil {
		return err
	}

	// defer closing the reader
	defer r.Close()

	// capture decompressed log data from the compressed log data
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	// overwrite compressed log data with decompressed log data
	l.Data = data

	return nil
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the Log type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (l *Log) Nullify() *Log {
	if l == nil {
		return nil
	}

	// check if the ID field should be false
	if l.ID.Int64 == 0 {
		l.ID.Valid = false
	}

	// check if the BuildID field should be false
	if l.BuildID.Int64 == 0 {
		l.BuildID.Valid = false
	}

	// check if the RepoID field should be false
	if l.RepoID.Int64 == 0 {
		l.RepoID.Valid = false
	}

	// check if the ServiceID field should be false
	if l.ServiceID.Int64 == 0 {
		l.ServiceID.Valid = false
	}

	// check if the StepID field should be false
	if l.StepID.Int64 == 0 {
		l.StepID.Valid = false
	}

	return l
}

// ToLibrary converts the Log type
// to a library Log type.
func (l *Log) ToLibrary() *library.Log {
	log := new(library.Log)

	log.SetID(l.ID.Int64)
	log.SetBuildID(l.BuildID.Int64)
	log.SetRepoID(l.RepoID.Int64)
	log.SetServiceID(l.ServiceID.Int64)
	log.SetStepID(l.StepID.Int64)
	log.SetData(l.Data)

	return log
}

// Validate verifies the necessary fields for
// the Log type are populated correctly.
func (l *Log) Validate() error {
	// verify the has StepID or ServiceID field populated
	if l.StepID.Int64 <= 0 && l.ServiceID.Int64 <= 0 {
		return ErrEmptyLogStepOrServiceID
	}

	// verify the BuildID field is populated
	if l.BuildID.Int64 <= 0 {
		return ErrEmptyLogBuildID
	}

	// verify the RepoID field is populated
	if l.RepoID.Int64 <= 0 {
		return ErrEmptyLogRepoID
	}

	return nil
}

// LogFromLibrary converts the Log type
// to a library Log type.
func LogFromLibrary(l *library.Log) *Log {
	log := &Log{
		ID:        sql.NullInt64{Int64: l.GetID(), Valid: true},
		BuildID:   sql.NullInt64{Int64: l.GetBuildID(), Valid: true},
		RepoID:    sql.NullInt64{Int64: l.GetRepoID(), Valid: true},
		ServiceID: sql.NullInt64{Int64: l.GetServiceID(), Valid: true},
		StepID:    sql.NullInt64{Int64: l.GetStepID(), Valid: true},
		Data:      l.GetData(),
	}

	return log.Nullify()
}
