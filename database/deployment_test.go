// SPDX-License-Identifier: Apache-2.0

package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
	"github.com/go-vela/types/raw"
	"github.com/google/go-cmp/cmp"
	"github.com/lib/pq"
)

func TestDatabase_Deployment_Nullify(t *testing.T) {
	// setup types
	var d *Deployment

	want := &Deployment{
		ID:          sql.NullInt64{Int64: 0, Valid: false},
		Number:      sql.NullInt64{Int64: 0, Valid: false},
		RepoID:      sql.NullInt64{Int64: 0, Valid: false},
		URL:         sql.NullString{String: "", Valid: false},
		User:        sql.NullString{String: "", Valid: false},
		Commit:      sql.NullString{String: "", Valid: false},
		Ref:         sql.NullString{String: "", Valid: false},
		Task:        sql.NullString{String: "", Valid: false},
		Target:      sql.NullString{String: "", Valid: false},
		Description: sql.NullString{String: "", Valid: false},
		Payload:     nil,
		Builds:      nil,
	}

	// setup tests
	tests := []struct {
		deployment *Deployment
		want       *Deployment
	}{
		{
			deployment: testDeployment(),
			want:       testDeployment(),
		},
		{
			deployment: d,
			want:       nil,
		},
		{
			deployment: new(Deployment),
			want:       want,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.deployment.Nullify()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Nullify is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_Deployment_ToLibrary(t *testing.T) {
	want := new(library.Deployment)
	want.SetID(1)
	want.SetNumber(1)
	want.SetRepoID(1)
	want.SetURL("https://github.com/github/octocat/deployments/1")
	want.SetUser("octocat")
	want.SetCommit("1234")
	want.SetRef("refs/heads/main")
	want.SetTask("deploy:vela")
	want.SetTarget("production")
	want.SetDescription("Deployment request from Vela")
	want.SetPayload(raw.StringSliceMap{"foo": "test1"})
	want.SetBuilds(nil)

	got := testDeployment().ToLibrary(nil)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("(ToLibrary: -want +got):\n%s", diff)
	}
}

func TestDatabase_Deployment_Validate(t *testing.T) {
	// setup types
	tests := []struct {
		failure    bool
		deployment *Deployment
	}{
		{
			failure:    false,
			deployment: testDeployment(),
		},
		{ // no number set for secret
			failure: true,
			deployment: &Deployment{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
				RepoID: sql.NullInt64{Int64: 1, Valid: true},
			},
		},
		{ // no repoID set for secret
			failure: true,
			deployment: &Deployment{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
				Number: sql.NullInt64{Int64: 1, Valid: true},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.deployment.Validate()

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

func TestDatabase_DeploymentFromLibrary(t *testing.T) {
	d := new(library.Deployment)
	d.SetID(1)
	d.SetNumber(1)
	d.SetRepoID(1)
	d.SetURL("https://github.com/github/octocat/deployments/1")
	d.SetUser("octocat")
	d.SetCommit("1234")
	d.SetRef("refs/heads/main")
	d.SetTask("deploy:vela")
	d.SetTarget("production")
	d.SetDescription("Deployment request from Vela")
	d.SetPayload(raw.StringSliceMap{"foo": "test1"})
	d.SetBuilds(nil)

	want := testDeployment()

	// run test
	got := DeploymentFromLibrary(d)

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("(-want +got):\n%s", diff)
	}
}

// testDeployment is a test helper function to create a Deployment type with all fields set to a fake value.
func testDeployment() *Deployment {
	return &Deployment{
		ID:          sql.NullInt64{Int64: 1, Valid: true},
		Number:      sql.NullInt64{Int64: 1, Valid: true},
		RepoID:      sql.NullInt64{Int64: 1, Valid: true},
		URL:         sql.NullString{String: "https://github.com/github/octocat/deployments/1", Valid: true},
		User:        sql.NullString{String: "octocat", Valid: true},
		Commit:      sql.NullString{String: "1234", Valid: true},
		Ref:         sql.NullString{String: "refs/heads/main", Valid: true},
		Task:        sql.NullString{String: "deploy:vela", Valid: true},
		Target:      sql.NullString{String: "production", Valid: true},
		Description: sql.NullString{String: "Deployment request from Vela", Valid: true},
		Payload:     raw.StringSliceMap{"foo": "test1"},
		Builds:      pq.StringArray{},
	}
}
