// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"reflect"
	"strings"
	"testing"
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
			t.Errorf("Purge is %v, want %v", got, test.want)
		}
	}
}

func TestPipeline_Secret_ParseOrg_success(t *testing.T) {
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

		org, key, err := test.secret.ParseOrg(test.org)
		if err != nil {
			t.Errorf("ParseOrg had an error occur: %+v", err)
		}

		p := strings.SplitN(test.secret.Key, "/", 2)

		if !strings.EqualFold(org, p[0]) {
			t.Errorf("org is %s want %s", org, p[0])
		}

		if !strings.EqualFold(key, p[1]) {
			t.Errorf("key is %s want %s", key, p[1])
		}
	}
}

func TestPipeline_Secret_ParseOrg_failure(t *testing.T) {
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

		_, _, err := test.secret.ParseOrg(test.org)
		if err == nil {
			t.Errorf("ParseOrg should have failed")
		}
	}
}

func TestPipeline_Secret_ParseRepo_success(t *testing.T) {
	// setup tests
	tests := []struct {
		secret *Secret
		org    string
		repo   string
	}{
		{ // success with explicit
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
		{ // success with implicit
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

		org, repo, key, err := test.secret.ParseRepo(test.org, test.repo)
		if err != nil {
			t.Errorf("ParseRepo had an error occur: %+v", err)
		}

		// checks for explicit only
		if strings.Contains(test.secret.Key, "/") {
			p := strings.SplitN(test.secret.Key, "/", 3)

			if !strings.EqualFold(org, p[0]) {
				t.Errorf("org is %s want %s", org, p[0])
			}

			if !strings.EqualFold(repo, p[1]) {
				t.Errorf("repo is %s want %s", key, p[1])
			}

			if !strings.EqualFold(key, p[2]) {
				t.Errorf("key is %s want %s", key, p[2])
			}
		}
	}
}

func TestPipeline_Secret_ParseRepo_failure(t *testing.T) {
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

		_, _, _, err := test.secret.ParseRepo(test.org, test.repo)
		if err == nil {
			t.Errorf("ParseRepo should have failed")
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

		org, team, key, err := test.secret.ParseShared(test.org)
		if err != nil {
			t.Errorf("ParseShared had an error occur: %+v", err)
		}

		p := strings.SplitN(test.secret.Key, "/", 3)

		if !strings.EqualFold(org, p[0]) {
			t.Errorf("org is %s want %s", org, p[0])
		}

		if !strings.EqualFold(team, p[1]) {
			t.Errorf("repo is %s want %s", key, p[1])
		}

		if !strings.EqualFold(key, p[2]) {
			t.Errorf("key is %s want %s", key, p[2])
		}
	}
}

func TestPipeline_Secret_ParseShared_failure(t *testing.T) {
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

		_, _, _, err := test.secret.ParseShared(test.org)
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
