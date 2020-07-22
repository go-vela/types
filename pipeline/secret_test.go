// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPipeline_SecretSlice_Purge(t *testing.T) {
	// setup types
	secrets := testSecrets()
	*secrets = (*secrets)[:len(*secrets)-1]

	// setup tests
	tests := []struct {
		secrets *SecretSlice
		want    *SecretSlice
	}{
		{
			secrets: testSecrets(),
			want:    secrets,
		},
		{
			secrets: new(SecretSlice),
			want:    new(SecretSlice),
		},
	}

	// run tests
	for _, test := range tests {
		r := &RuleData{
			Branch: "master",
			Event:  "push",
			Path:   []string{},
			Repo:   "foo/bar",
			Tag:    "refs/heads/master",
		}

		got := test.secrets.Purge(r)

		if !reflect.DeepEqual(got, test.want) {
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("MakeGatewayInfo() mismatch (-want +got):\n%s", diff)
			}
			t.Errorf("Purge is %v, want %v", got, test.want)
		}
	}
}

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
		{ // failure with bad engine
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat",
				Engine: "invalid",
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
		{ // success with implicit
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
		{ // success with explicit
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "foo",
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
		{ // failure with bad engine
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat",
				Engine: "invalid",
				Type:   "org",
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
		{ // failure with bad engine
			secret: &Secret{
				Name:   "foo",
				Value:  "bar",
				Key:    "octocat",
				Engine: "invalid",
				Type:   "org",
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

func testSecrets() *SecretSlice {
	return &SecretSlice{
		{
			Engine: "native",
			Key:    "github/octocat/foobar",
			Name:   "foobar",
			Type:   "repo",
			Origin: &Container{},
		},
		{
			Engine: "native",
			Key:    "github/foobar",
			Name:   "foobar",
			Type:   "org",
			Origin: &Container{},
		},
		{
			Engine: "native",
			Key:    "github/octokitties/foobar",
			Name:   "foobar",
			Type:   "shared",
			Origin: &Container{},
		},
		{
			Name: "",
			Origin: &Container{
				ID:          "secret_github octocat._1_vault",
				Directory:   "/vela/src/foo//",
				Environment: map[string]string{"FOO": "bar"},
				Image:       "vault:latest",
				Name:        "vault",
				Number:      1,
				Pull:        true,
				Ruleset: Ruleset{
					If:       Rules{Event: []string{"push"}},
					Operator: "and",
				},
			},
		},
		{
			Origin: &Container{
				ID:          "secret_github octocat._2_vault",
				Directory:   "/vela/src/foo//",
				Environment: map[string]string{"FOO": "bar"},
				Image:       "vault:latest",
				Name:        "vault",
				Number:      2,
				Pull:        true,
				Ruleset: Ruleset{
					If:       Rules{Event: []string{"pull_request"}},
					Operator: "and",
				},
			},
		},
	}
}
