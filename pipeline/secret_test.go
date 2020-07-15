// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"testing"
)

func TestPipeline_Secret_ValidOrg(t *testing.T) {
	// setup types

	// setup tests
	tests := []struct {
		secret *Secret
		org    string
		want   bool
	}{
		{ // success with good data
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat/foo",
				Engine: "native",
				Type:   "org",
			},
			org:  "octocat",
			want: true,
		},
		{ // failure with bad org
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat/foo",
				Engine: "native",
				Type:   "org",
			},
			org:  "wrongorg",
			want: false,
		},
		{ // failure with bad key
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat",
				Engine: "native",
				Type:   "org",
			},
			org:  "octocat",
			want: false,
		},
	}

	// run tests
	for _, test := range tests {

		got, _ := test.secret.ValidOrg(test.org)

		if got != test.want {
			t.Errorf("ValidOrg is %v, want %v", got, test.want)
		}
	}
}

func TestPipeline_Secret_ValidRepo(t *testing.T) {
	// setup types

	// setup tests
	tests := []struct {
		secret *Secret
		org    string
		repo   string
		want   bool
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
			want: true,
		},
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
			want: false,
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
			want: false,
		},
		{ // failure with bad key
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat",
				Engine: "native",
				Type:   "repo",
			},
			org:  "octocat",
			want: false,
		},
	}

	// run tests
	for _, test := range tests {

		got, _ := test.secret.ValidRepo(test.org, test.repo)

		if got != test.want {
			t.Errorf("ValidRepo is %v, want %v", got, test.want)
		}
	}
}

func TestPipeline_Secret_ValidShared(t *testing.T) {
	// setup types

	// setup tests
	tests := []struct {
		secret *Secret
		org    string
		team   string
		want   bool
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
			team: "helloworld",
			want: true,
		},
		{ // failure with bad org
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat/helloworld/foo",
				Engine: "native",
				Type:   "repo",
			},
			org:  "wrongorg",
			team: "helloworld",
			want: false,
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
			team: "badrepo",
			want: false,
		},
		{ // failure with bad key
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat",
				Engine: "native",
				Type:   "repo",
			},
			org:  "octocat",
			team: "helloworld",
			want: false,
		},
	}

	// run tests
	for _, test := range tests {

		got, _ := test.secret.ValidShared(test.org, test.team)

		if got != test.want {
			t.Errorf("ValidShared is %v, want %v", got, test.want)
		}
	}
}
