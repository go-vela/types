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
		name     string
		failure  bool
		level    int
		pipeline *Pipeline
		want     []byte
	}{
		{
			name:     "compression level -1",
			failure:  false,
			level:    constants.CompressionNegOne,
			pipeline: &Pipeline{Data: []byte("foo")},
			want:     []byte{120, 156, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69},
		},
		{
			name:     "compression level 0",
			failure:  false,
			level:    constants.CompressionZero,
			pipeline: &Pipeline{Data: []byte("foo")},
			want:     []byte{120, 1, 0, 3, 0, 252, 255, 102, 111, 111, 1, 0, 0, 255, 255, 2, 130, 1, 69},
		},
		{
			name:     "compression level 1",
			failure:  false,
			level:    constants.CompressionOne,
			pipeline: &Pipeline{Data: []byte("foo")},
			want:     []byte{120, 1, 0, 3, 0, 252, 255, 102, 111, 111, 1, 0, 0, 255, 255, 2, 130, 1, 69},
		},
		{
			name:     "compression level 2",
			failure:  false,
			level:    constants.CompressionTwo,
			pipeline: &Pipeline{Data: []byte("foo")},
			want:     []byte{120, 94, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69},
		},
		{
			name:     "compression level 3",
			failure:  false,
			level:    constants.CompressionThree,
			pipeline: &Pipeline{Data: []byte("foo")},
			want:     []byte{120, 94, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69},
		},
		{
			name:     "compression level 4",
			failure:  false,
			level:    constants.CompressionFour,
			pipeline: &Pipeline{Data: []byte("foo")},
			want:     []byte{120, 94, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69},
		},
		{
			name:     "compression level 5",
			failure:  false,
			level:    constants.CompressionFive,
			pipeline: &Pipeline{Data: []byte("foo")},
			want:     []byte{120, 94, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69},
		},
		{
			name:     "compression level 6",
			failure:  false,
			level:    constants.CompressionSix,
			pipeline: &Pipeline{Data: []byte("foo")},
			want:     []byte{120, 156, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69},
		},
		{
			name:     "compression level 7",
			failure:  false,
			level:    constants.CompressionSeven,
			pipeline: &Pipeline{Data: []byte("foo")},
			want:     []byte{120, 218, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69},
		},
		{
			name:     "compression level 8",
			failure:  false,
			level:    constants.CompressionEight,
			pipeline: &Pipeline{Data: []byte("foo")},
			want:     []byte{120, 218, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69},
		},
		{
			name:     "compression level 9",
			failure:  false,
			level:    constants.CompressionNine,
			pipeline: &Pipeline{Data: []byte("foo")},
			want:     []byte{120, 218, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.pipeline.Compress(test.level)

		if test.failure {
			if err == nil {
				t.Errorf("Compress for %s should have returned err", test.name)
			}

			continue
		}

		if err != nil {
			t.Errorf("Compress for %s returned err: %v", test.name, err)
		}

		if !reflect.DeepEqual(test.pipeline.Data, test.want) {
			t.Errorf("Compress for %s is %v, want %v", test.name, string(test.pipeline.Data), string(test.want))
		}
	}
}

func TestDatabase_Pipeline_Decompress(t *testing.T) {
	// setup tests
	tests := []struct {
		name     string
		failure  bool
		pipeline *Pipeline
		want     []byte
	}{
		{
			name:     "compression level -1",
			failure:  false,
			pipeline: &Pipeline{Data: []byte{120, 156, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69}},
			want:     []byte("foo"),
		},
		{
			name:     "compression level 0",
			failure:  false,
			pipeline: &Pipeline{Data: []byte{120, 1, 0, 3, 0, 252, 255, 102, 111, 111, 1, 0, 0, 255, 255, 2, 130, 1, 69}},
			want:     []byte("foo"),
		},
		{
			name:     "compression level 1",
			failure:  false,
			pipeline: &Pipeline{Data: []byte{120, 1, 0, 3, 0, 252, 255, 102, 111, 111, 1, 0, 0, 255, 255, 2, 130, 1, 69}},
			want:     []byte("foo"),
		},
		{
			name:     "compression level 2",
			failure:  false,
			pipeline: &Pipeline{Data: []byte{120, 94, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69}},
			want:     []byte("foo"),
		},
		{
			name:     "compression level 3",
			failure:  false,
			pipeline: &Pipeline{Data: []byte{120, 94, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69}},
			want:     []byte("foo"),
		},
		{
			name:     "compression level 4",
			failure:  false,
			pipeline: &Pipeline{Data: []byte{120, 94, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69}},
			want:     []byte("foo"),
		},
		{
			name:     "compression level 5",
			failure:  false,
			pipeline: &Pipeline{Data: []byte{120, 94, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69}},
			want:     []byte("foo"),
		},
		{
			name:     "compression level 6",
			failure:  false,
			pipeline: &Pipeline{Data: []byte{120, 156, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69}},
			want:     []byte("foo"),
		},
		{
			name:     "compression level 7",
			failure:  false,
			pipeline: &Pipeline{Data: []byte{120, 218, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69}},
			want:     []byte("foo"),
		},
		{
			name:     "compression level 8",
			failure:  false,
			pipeline: &Pipeline{Data: []byte{120, 218, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69}},
			want:     []byte("foo"),
		},
		{
			name:     "compression level 9",
			failure:  false,
			pipeline: &Pipeline{Data: []byte{120, 218, 74, 203, 207, 7, 4, 0, 0, 255, 255, 2, 130, 1, 69}},
			want:     []byte("foo"),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.pipeline.Decompress()

		if test.failure {
			if err == nil {
				t.Errorf("Decompress for %s should have returned err", test.name)
			}

			continue
		}

		if err != nil {
			t.Errorf("Decompress for %s returned err: %v", test.name, err)
		}

		if !reflect.DeepEqual(test.pipeline.Data, test.want) {
			t.Errorf("Decompress for %s is %v, want %v", test.name, string(test.pipeline.Data), string(test.want))
		}
	}
}

func TestDatabase_Pipeline_Nullify(t *testing.T) {
	// setup types
	var p *Pipeline

	want := &Pipeline{
		ID:       sql.NullInt64{Int64: 0, Valid: false},
		RepoID:   sql.NullInt64{Int64: 0, Valid: false},
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
		{ // no ref set for pipeline
			failure: true,
			pipeline: &Pipeline{
				ID:      sql.NullInt64{Int64: 1, Valid: true},
				RepoID:  sql.NullInt64{Int64: 1, Valid: true},
				Version: sql.NullString{String: "1", Valid: true},
			},
		},
		{ // no repo_id set for pipeline
			failure: true,
			pipeline: &Pipeline{
				ID:      sql.NullInt64{Int64: 1, Valid: true},
				Ref:     sql.NullString{String: "refs/heads/master", Valid: true},
				Version: sql.NullString{String: "1", Valid: true},
			},
		},
		{ // no version set for pipeline
			failure: true,
			pipeline: &Pipeline{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
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
