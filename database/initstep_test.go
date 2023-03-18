// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
)

func TestDatabase_InitStep_Nullify(t *testing.T) {
	// setup types
	var s *InitStep

	want := &InitStep{
		ID:        sql.NullInt64{Int64: 0, Valid: false},
		RepoID:    sql.NullInt64{Int64: 0, Valid: false},
		BuildID:   sql.NullInt64{Int64: 0, Valid: false},
		StepID:    sql.NullInt64{Int64: 0, Valid: false},
		ServiceID: sql.NullInt64{Int64: 0, Valid: false},
		Number:    sql.NullInt32{Int32: 0, Valid: false},
		Reporter:  sql.NullString{String: "", Valid: false},
		Name:      sql.NullString{String: "", Valid: false},
	}

	// setup tests
	tests := []struct {
		initStep *InitStep
		want     *InitStep
	}{
		{
			initStep: testInitStep(),
			want:     testInitStep(),
		},
		{
			initStep: testStepInitStep(),
			want:     testStepInitStep(),
		},
		{
			initStep: testServiceInitStep(),
			want:     testServiceInitStep(),
		},
		{
			initStep: s,
			want:     nil,
		},
		{
			initStep: new(InitStep),
			want:     want,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.initStep.Nullify()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Nullify is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_InitStep_ToLibrary(t *testing.T) {
	// setup types
	forBuild := new(library.InitStep)
	forStep := new(library.InitStep)
	forService := new(library.InitStep)

	forBuild.SetID(1)
	forBuild.SetRepoID(1)
	forBuild.SetBuildID(1)
	forBuild.SetNumber(1)
	forBuild.SetReporter("Foobar Runtime")
	forBuild.SetName("foobar")

	forStep.SetID(1)
	forStep.SetRepoID(1)
	forStep.SetBuildID(1)
	forStep.SetStepID(1)
	forStep.SetNumber(1)
	forStep.SetReporter("Foobar Runtime")
	forStep.SetName("foobar")

	forService.SetID(1)
	forService.SetRepoID(1)
	forService.SetBuildID(1)
	forService.SetServiceID(1)
	forService.SetNumber(1)
	forService.SetReporter("Foobar Runtime")
	forService.SetName("foobar")

	tests := []struct {
		initStep *InitStep
		want     *library.InitStep
	}{
		{
			initStep: testInitStep(),
			want:     forBuild,
		},
		{
			initStep: testStepInitStep(),
			want:     forStep,
		},
		{
			initStep: testServiceInitStep(),
			want:     forService,
		},
	}
	// run tests
	for _, test := range tests {
		// run test
		got := test.initStep.ToLibrary()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ToLibrary is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_InitStep_Validate(t *testing.T) {
	// setup types
	tests := []struct {
		failure  bool
		initStep *InitStep
	}{
		{
			failure:  false,
			initStep: testInitStep(),
		},
		{ // no repo_id set for InitStep
			failure: true,
			initStep: &InitStep{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				BuildID:  sql.NullInt64{Int64: 1, Valid: true},
				Number:   sql.NullInt32{Int32: 1, Valid: true},
				Reporter: sql.NullString{String: "Foobar Runtime", Valid: true},
				Name:     sql.NullString{String: "foobar", Valid: true},
			},
		},
		{ // no build_id set for InitStep
			failure: true,
			initStep: &InitStep{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				RepoID:   sql.NullInt64{Int64: 1, Valid: true},
				Number:   sql.NullInt32{Int32: 1, Valid: true},
				Reporter: sql.NullString{String: "Foobar Runtime", Valid: true},
				Name:     sql.NullString{String: "foobar", Valid: true},
			},
		},
		{ // no number set for InitStep
			failure: true,
			initStep: &InitStep{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				RepoID:   sql.NullInt64{Int64: 1, Valid: true},
				BuildID:  sql.NullInt64{Int64: 1, Valid: true},
				Reporter: sql.NullString{String: "Foobar Runtime", Valid: true},
				Name:     sql.NullString{String: "foobar", Valid: true},
			},
		},
		{ // no reporter set for InitStep
			failure: true,
			initStep: &InitStep{
				ID:      sql.NullInt64{Int64: 1, Valid: true},
				RepoID:  sql.NullInt64{Int64: 1, Valid: true},
				BuildID: sql.NullInt64{Int64: 1, Valid: true},
				Number:  sql.NullInt32{Int32: 1, Valid: true},
				Name:    sql.NullString{String: "foobar", Valid: true},
			},
		},
		{ // no name set for InitStep
			failure: true,
			initStep: &InitStep{
				ID:       sql.NullInt64{Int64: 1, Valid: true},
				RepoID:   sql.NullInt64{Int64: 1, Valid: true},
				BuildID:  sql.NullInt64{Int64: 1, Valid: true},
				Number:   sql.NullInt32{Int32: 1, Valid: true},
				Reporter: sql.NullString{String: "Foobar Runtime", Valid: true},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.initStep.Validate()

		if test.failure {
			if err == nil {
				t.Errorf("Validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}

func TestDatabase_InitStepFromLibrary(t *testing.T) {
	// setup types
	forBuild := new(library.InitStep)
	forStep := new(library.InitStep)
	forService := new(library.InitStep)

	forBuild.SetID(1)
	forBuild.SetRepoID(1)
	forBuild.SetBuildID(1)
	forBuild.SetNumber(1)
	forBuild.SetReporter("Foobar Runtime")
	forBuild.SetName("foobar")

	forStep.SetID(1)
	forStep.SetRepoID(1)
	forStep.SetBuildID(1)
	forStep.SetStepID(1)
	forStep.SetNumber(1)
	forStep.SetReporter("Foobar Runtime")
	forStep.SetName("foobar")

	forService.SetID(1)
	forService.SetRepoID(1)
	forService.SetBuildID(1)
	forService.SetServiceID(1)
	forService.SetNumber(1)
	forService.SetReporter("Foobar Runtime")
	forService.SetName("foobar")

	tests := []struct {
		library *library.InitStep
		want    *InitStep
	}{
		{
			library: forBuild,
			want:    testInitStep(),
		},
		{
			library: forStep,
			want:    testStepInitStep(),
		},
		{
			library: forService,
			want:    testServiceInitStep(),
		},
	}
	// run tests
	for _, test := range tests {
		// run test
		got := InitStepFromLibrary(test.library)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("InitStepFromLibrary is %v, want %v", got, test.want)
		}
	}
}

// testInitStep is a test helper function to create a InitStep
// type for a Build with all fields set to a fake value.
func testInitStep() *InitStep {
	return &InitStep{
		ID:        sql.NullInt64{Int64: 1, Valid: true},
		RepoID:    sql.NullInt64{Int64: 1, Valid: true},
		BuildID:   sql.NullInt64{Int64: 1, Valid: false},
		StepID:    sql.NullInt64{Int64: 0, Valid: false},
		ServiceID: sql.NullInt64{Int64: 0, Valid: true},
		Number:    sql.NullInt32{Int32: 1, Valid: true},
		Reporter:  sql.NullString{String: "Foobar Runtime", Valid: true},
		Name:      sql.NullString{String: "foobar", Valid: true},
	}
}

// testStepInitStep is a test helper function to create a InitStep
// type for a Step with all fields set to a fake value.
func testStepInitStep() *InitStep {
	i := testInitStep()
	i.StepID = sql.NullInt64{Int64: 1, Valid: true}

	return i
}

// testServiceInitStep is a test helper function to create a InitStep
// type for a Service with all fields set to a fake value.
func testServiceInitStep() *InitStep {
	i := testInitStep()
	i.ServiceID = sql.NullInt64{Int64: 1, Valid: true}

	return i
}
