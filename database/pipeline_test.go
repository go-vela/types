// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
)

func TestDatabase_Pipeline_Compress(t *testing.T) {
	// setup tests
	tests := []struct {
		failure  bool
		level    int
		pipeline *Pipeline
	}{
		{
			failure:  false,
			level:    constants.CompressionNegOne,
			pipeline: testPipeline(),
		},
		{
			failure:  false,
			level:    constants.CompressionZero,
			pipeline: testPipeline(),
		},
		{
			failure:  false,
			level:    constants.CompressionOne,
			pipeline: testPipeline(),
		},
		{
			failure:  false,
			level:    constants.CompressionTwo,
			pipeline: testPipeline(),
		},
		{
			failure:  false,
			level:    constants.CompressionThree,
			pipeline: testPipeline(),
		},
		{
			failure:  false,
			level:    constants.CompressionFour,
			pipeline: testPipeline(),
		},
		{
			failure:  false,
			level:    constants.CompressionFive,
			pipeline: testPipeline(),
		},
		{
			failure:  false,
			level:    constants.CompressionSix,
			pipeline: testPipeline(),
		},
		{
			failure:  false,
			level:    constants.CompressionSeven,
			pipeline: testPipeline(),
		},
		{
			failure:  false,
			level:    constants.CompressionEight,
			pipeline: testPipeline(),
		},
		{
			failure:  false,
			level:    constants.CompressionNine,
			pipeline: testPipeline(),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.pipeline.Compress(test.level)

		if test.failure {
			if err == nil {
				t.Errorf("Compress should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Compress returned err: %v", err)
		}
	}
}

func TestDatabase_Pipeline_Decompress(t *testing.T) {
	// setup types
	p := testPipeline()
	err := p.Compress(constants.CompressionThree)
	if err != nil {
		t.Errorf("unable to compress log: %v", err)
	}

	// setup tests
	tests := []struct {
		failure  bool
		pipeline *Pipeline
	}{
		{
			failure:  false,
			pipeline: p,
		},
		{
			failure:  true,
			pipeline: testPipeline(),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.pipeline.Decompress()

		if test.failure {
			if err == nil {
				t.Errorf("Decompress should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Decompress returned err: %v", err)
		}
	}
}

func TestDatabase_Pipeline_Nullify(t *testing.T) {
	// setup types
	var p *Pipeline

	want := &Pipeline{
		ID:       sql.NullInt64{Int64: 0, Valid: false},
		RepoID:   sql.NullInt64{Int64: 0, Valid: false},
		Number:   sql.NullInt32{Int32: 0, Valid: false},
		Flavor:   sql.NullString{String: "", Valid: false},
		Platform: sql.NullString{String: "", Valid: false},
		Ref:      sql.NullString{String: "", Valid: false},
		Version:  sql.NullString{String: "", Valid: false},
	}

	// setup tests
	tests := []struct {
		pipeline *Pipeline
		want     *Pipeline
	}{
		{
			pipeline: testPipeline(),
			want:     testPipeline(),
		},
		{
			pipeline: p,
			want:     nil,
		},
		{
			pipeline: new(Pipeline),
			want:     want,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.pipeline.Nullify()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Nullify is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_Pipeline_ToLibrary(t *testing.T) {
	// setup types
	want := new(library.Pipeline)

	want.SetID(1)
	want.SetRepoID(1)
	want.SetNumber(1)
	want.SetFlavor("large")
	want.SetPlatform("docker")
	want.SetRef("refs/heads/master")
	want.SetVersion("1")
	want.SetServices(true)
	want.SetStages(false)
	want.SetSteps(true)
	want.SetTemplates(false)
	want.SetData(testPipelineData())

	// run test
	got := testPipeline().ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Pipeline_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure  bool
		pipeline *Pipeline
	}{
		{
			failure:  false,
			pipeline: testPipeline(),
		},
		{ // no number set for pipeline
			failure: true,
			pipeline: &Pipeline{
				ID:      sql.NullInt64{Int64: 1, Valid: true},
				RepoID:  sql.NullInt64{Int64: 1, Valid: true},
				Ref:     sql.NullString{String: "refs/heads/master", Valid: true},
				Version: sql.NullString{String: "1", Valid: true},
			},
		},
		{ // no ref set for pipeline
			failure: true,
			pipeline: &Pipeline{
				ID:      sql.NullInt64{Int64: 1, Valid: true},
				RepoID:  sql.NullInt64{Int64: 1, Valid: true},
				Number:  sql.NullInt32{Int32: 1, Valid: true},
				Version: sql.NullString{String: "1", Valid: true},
			},
		},
		{ // no repo_id set for pipeline
			failure: true,
			pipeline: &Pipeline{
				ID:      sql.NullInt64{Int64: 1, Valid: true},
				Number:  sql.NullInt32{Int32: 1, Valid: true},
				Ref:     sql.NullString{String: "refs/heads/master", Valid: true},
				Version: sql.NullString{String: "1", Valid: true},
			},
		},
		{ // no version set for pipeline
			failure: true,
			pipeline: &Pipeline{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
				Number: sql.NullInt32{Int32: 1, Valid: true},
				Ref:    sql.NullString{String: "refs/heads/master", Valid: true},
				RepoID: sql.NullInt64{Int64: 1, Valid: true},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.pipeline.Validate()

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

func TestDatabase_PipelineFromLibrary(t *testing.T) {
	// setup types
	p := new(library.Pipeline)

	p.SetID(1)
	p.SetRepoID(1)
	p.SetNumber(1)
	p.SetFlavor("large")
	p.SetPlatform("docker")
	p.SetRef("refs/heads/master")
	p.SetVersion("1")
	p.SetServices(true)
	p.SetStages(false)
	p.SetSteps(true)
	p.SetTemplates(false)
	p.SetData(testPipelineData())

	want := testPipeline()

	// run test
	got := PipelineFromLibrary(p)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("PipelineFromLibrary is %v, want %v", got, want)
	}
}

// testPipeline is a test helper function to create a Pipeline
// type with all fields set to a fake value.
func testPipeline() *Pipeline {
	return &Pipeline{
		ID:        sql.NullInt64{Int64: 1, Valid: true},
		RepoID:    sql.NullInt64{Int64: 1, Valid: true},
		Number:    sql.NullInt32{Int32: 1, Valid: true},
		Flavor:    sql.NullString{String: "large", Valid: true},
		Platform:  sql.NullString{String: "docker", Valid: true},
		Ref:       sql.NullString{String: "refs/heads/master", Valid: true},
		Version:   sql.NullString{String: "1", Valid: true},
		Services:  sql.NullBool{Bool: true, Valid: true},
		Stages:    sql.NullBool{Bool: false, Valid: true},
		Steps:     sql.NullBool{Bool: true, Valid: true},
		Templates: sql.NullBool{Bool: false, Valid: true},
		Data:      testPipelineData(),
	}
}

// testPipelineData is a test helper function to create the
// content for the Data field for the Pipeline type.
func testPipelineData() []byte {
	return []byte(`
version: 1

worker:
  flavor: large
  platform: docker

services:
  - name: redis
    image: redis

steps:
  - name: ping
    image: redis
    commands:
      - redis-cli -h redis ping
`)
}
