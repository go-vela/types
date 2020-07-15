// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"testing"
)

func TestPipeline_Secret_ValidOrg_success(t *testing.T) {
	// setup tests
	tests := []struct {
		secret *Secret
		org    string
	}{
		{ // success with good data
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat/foo",
				Engine: "native",
				Type:   "org",
			},
			org: "octocat",
		},
	}

	// run tests
	for _, test := range tests {

		err := test.secret.ValidOrg(test.org)
		if err != nil {
			t.Errorf("ValidOrg had an error occur: %+v", err)
		}
	}
}

func TestPipeline_Secret_ValidOrg_failure(t *testing.T) {
	// setup tests
	tests := []struct {
		secret *Secret
		org    string
	}{
		{ // failure with bad org
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat/foo",
				Engine: "native",
				Type:   "org",
			},
			org: "wrongorg",
		},
		{ // failure with bad key
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat",
				Engine: "native",
				Type:   "org",
			},
			org: "octocat",
		},
	}

	// run tests
	for _, test := range tests {

		err := test.secret.ValidOrg(test.org)
		if err == nil {
			t.Errorf("ValidOrg should have failed")
		}
	}
}

func TestPipeline_Secret_ValidRepo_success(t *testing.T) {
	// setup tests
	tests := []struct {
		secret *Secret
		org    string
		repo   string
	}{
		{ // success with good data
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat/helloworld/foo",
				Engine: "native",
				Type:   "repo",
			},
			org:  "octocat",
			repo: "helloworld",
		},
	}

	// run tests
	for _, test := range tests {

		err := test.secret.ValidRepo(test.org, test.repo)
		if err != nil {
			t.Errorf("ValidRepo had an error occur: %+v", err)
		}
	}
}

func TestPipeline_Secret_ValidRepo_failure(t *testing.T) {
	// setup tests
	tests := []struct {
		secret *Secret
		org    string
		repo   string
		want   error
	}{
		{ // failure with bad org
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat/helloworld/foo",
				Engine: "native",
				Type:   "repo",
			},
			org:  "wrongorg",
			repo: "helloworld",
		},
		{ // failure with bad repo
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat/helloworld/foo",
				Engine: "native",
				Type:   "repo",
			},
			org:  "octocat",
			repo: "badrepo",
		},
		{ // failure with bad key
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat",
				Engine: "native",
				Type:   "repo",
			},
			org: "octocat",
		},
	}

	// run tests
	for _, test := range tests {

		err := test.secret.ValidRepo(test.org, test.repo)
		if err == nil {
			t.Errorf("ValidOrg should have failed")
		}
	}
}

func TestPipeline_Secret_ValidShared_success(t *testing.T) {
	// setup tests
	tests := []struct {
		secret *Secret
		org    string
	}{
		{ // success with good data
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat/helloworld/foo",
				Engine: "native",
				Type:   "repo",
			},
			org: "octocat",
		},
	}

	// run tests
	for _, test := range tests {

		err := test.secret.ValidShared(test.org)
		if err != nil {
			t.Errorf("ValidShared had an error occur: %+v", err)
		}
	}
}

func TestPipeline_Secret_ValidShared_failure(t *testing.T) {
	// setup tests
	tests := []struct {
		secret *Secret
		org    string
	}{
		{ // failure with bad org
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat/helloworld/foo",
				Engine: "native",
				Type:   "repo",
			},
			org: "wrongorg",
		},
		{ // failure with bad key
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat",
				Engine: "native",
				Type:   "repo",
			},
			org: "octocat",
		},
	}

	// run tests
	for _, test := range tests {

		err := test.secret.ValidShared(test.org)
		if err == nil {
			t.Errorf("ValidOrg should have failed")
		}
	}
}
