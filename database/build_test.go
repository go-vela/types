// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"math/rand"
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
)

func TestDatabase_Build_Crop(t *testing.T) {
	// setup types
	title := randomString(1001)
	message := randomString(2001)
	want := &Build{
		ID:      sql.NullInt64{Int64: 1, Valid: true},
		Title:   sql.NullString{String: title[:1000], Valid: true},
		Message: sql.NullString{String: message[:2000], Valid: true},
	}
	b := &Build{
		ID:      sql.NullInt64{Int64: 1, Valid: true},
		Title:   sql.NullString{String: title, Valid: true},
		Message: sql.NullString{String: message, Valid: true},
	}

	// run test
	got := b.Crop()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Crop is %v, want %v", got, want)
	}
}

func TestDatabase_Build_ToLibrary(t *testing.T) {
	// setup types
	sqlNum := sql.NullInt32{Int32: 1, Valid: true}
	num := 1
	num64 := int64(num)
	str := "foo"
	want := &library.Build{
		ID:           &num64,
		RepoID:       &num64,
		Number:       &num,
		Parent:       &num,
		Event:        &str,
		Status:       &str,
		Error:        &str,
		Enqueued:     &num64,
		Created:      &num64,
		Started:      &num64,
		Finished:     &num64,
		Deploy:       &str,
		Clone:        &str,
		Source:       &str,
		Title:        &str,
		Message:      &str,
		Commit:       &str,
		Sender:       &str,
		Author:       &str,
		Branch:       &str,
		Ref:          &str,
		BaseRef:      &str,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
	}
	b := &Build{
		ID:           sql.NullInt64{Int64: num64, Valid: true},
		RepoID:       sql.NullInt64{Int64: num64, Valid: true},
		Number:       sqlNum,
		Parent:       sqlNum,
		Event:        sql.NullString{String: str, Valid: true},
		Status:       sql.NullString{String: str, Valid: true},
		Error:        sql.NullString{String: str, Valid: true},
		Enqueued:     sql.NullInt64{Int64: num64, Valid: true},
		Created:      sql.NullInt64{Int64: num64, Valid: true},
		Started:      sql.NullInt64{Int64: num64, Valid: true},
		Finished:     sql.NullInt64{Int64: num64, Valid: true},
		Deploy:       sql.NullString{String: str, Valid: true},
		Clone:        sql.NullString{String: str, Valid: true},
		Source:       sql.NullString{String: str, Valid: true},
		Title:        sql.NullString{String: str, Valid: true},
		Message:      sql.NullString{String: str, Valid: true},
		Commit:       sql.NullString{String: str, Valid: true},
		Sender:       sql.NullString{String: str, Valid: true},
		Author:       sql.NullString{String: str, Valid: true},
		Branch:       sql.NullString{String: str, Valid: true},
		Ref:          sql.NullString{String: str, Valid: true},
		BaseRef:      sql.NullString{String: str, Valid: true},
		Host:         sql.NullString{String: str, Valid: true},
		Runtime:      sql.NullString{String: str, Valid: true},
		Distribution: sql.NullString{String: str, Valid: true},
	}

	// run test
	got := b.ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Build_BuildFromLibrary(t *testing.T) {
	// setup types
	sqlNum := sql.NullInt32{Int32: 1, Valid: true}
	num := 1
	num64 := int64(num)
	str := "foo"
	want := &Build{
		ID:           sql.NullInt64{Int64: num64, Valid: true},
		RepoID:       sql.NullInt64{Int64: num64, Valid: true},
		Number:       sqlNum,
		Parent:       sqlNum,
		Event:        sql.NullString{String: str, Valid: true},
		Status:       sql.NullString{String: str, Valid: true},
		Error:        sql.NullString{String: str, Valid: true},
		Enqueued:     sql.NullInt64{Int64: num64, Valid: true},
		Created:      sql.NullInt64{Int64: num64, Valid: true},
		Started:      sql.NullInt64{Int64: num64, Valid: true},
		Finished:     sql.NullInt64{Int64: num64, Valid: true},
		Deploy:       sql.NullString{String: str, Valid: true},
		Clone:        sql.NullString{String: str, Valid: true},
		Source:       sql.NullString{String: str, Valid: true},
		Title:        sql.NullString{String: str, Valid: true},
		Message:      sql.NullString{String: str, Valid: true},
		Commit:       sql.NullString{String: str, Valid: true},
		Sender:       sql.NullString{String: str, Valid: true},
		Author:       sql.NullString{String: str, Valid: true},
		Branch:       sql.NullString{String: str, Valid: true},
		Ref:          sql.NullString{String: str, Valid: true},
		BaseRef:      sql.NullString{String: str, Valid: true},
		Host:         sql.NullString{String: str, Valid: true},
		Runtime:      sql.NullString{String: str, Valid: true},
		Distribution: sql.NullString{String: str, Valid: true},
	}

	b := &library.Build{
		ID:           &num64,
		RepoID:       &num64,
		Number:       &num,
		Parent:       &num,
		Event:        &str,
		Status:       &str,
		Error:        &str,
		Enqueued:     &num64,
		Created:      &num64,
		Started:      &num64,
		Finished:     &num64,
		Deploy:       &str,
		Clone:        &str,
		Source:       &str,
		Title:        &str,
		Message:      &str,
		Commit:       &str,
		Sender:       &str,
		Author:       &str,
		Branch:       &str,
		Ref:          &str,
		BaseRef:      &str,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
	}

	// run test
	got := BuildFromLibrary(b)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BuildFromLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Build_Validate(t *testing.T) {
	// setup types
	b := &Build{
		ID:     sql.NullInt64{Int64: 1, Valid: true},
		RepoID: sql.NullInt64{Int64: 1, Valid: true},
		Number: sql.NullInt32{Int32: 1, Valid: true},
	}

	// run test
	err := b.Validate()

	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestDatabase_Build_Validate_NoRepoID(t *testing.T) {
	// setup types
	b := &Build{
		ID:     sql.NullInt64{Int64: 1, Valid: true},
		Number: sql.NullInt32{Int32: 1, Valid: true},
	}

	// run test
	err := b.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestDatabase_Build_Validate_NoNumber(t *testing.T) {
	// setup types
	b := &Build{
		ID:     sql.NullInt64{Int64: 1, Valid: true},
		RepoID: sql.NullInt64{Int64: 1, Valid: true},
	}

	// run test
	err := b.Validate()

	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}
