// SPDX-License-Identifier: Apache-2.0

package database

import (
	"database/sql"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/lib/pq"

	"github.com/go-vela/types/library"
	"github.com/go-vela/types/raw"
)

func TestDatabase_Deployment_Nullify(t *testing.T) {
	// setup types
	var d *Deployment

	want := &Deployment{
		ID:          sql.NullInt64{Int64: 0, Valid: false},
		Number:      sql.NullInt64{Int64: 0, Valid: false},
		RepoID:      sql.NullInt64{Int64: 0, Valid: false},
		URL:         sql.NullString{String: "", Valid: false},
		Commit:      sql.NullString{String: "", Valid: false},
		Ref:         sql.NullString{String: "", Valid: false},
		Task:        sql.NullString{String: "", Valid: false},
		Target:      sql.NullString{String: "", Valid: false},
		Description: sql.NullString{String: "", Valid: false},
		Payload:     nil,
		CreatedAt:   sql.NullInt64{Int64: 0, Valid: false},
		CreatedBy:   sql.NullString{String: "", Valid: false},
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

		if diff := cmp.Diff(test.want, got); diff != "" {
			t.Errorf("(ToLibrary: -want +got):\n%s", diff)
		}
	}
}

func TestDatabase_Deployment_ToLibrary(t *testing.T) {
	builds := []*library.Build{}

	buildOne := new(library.Build)
	buildOne.SetID(1)
	buildOne.SetRepoID(1)
	buildOne.SetPipelineID(1)
	buildOne.SetNumber(1)
	buildOne.SetParent(1)
	buildOne.SetEvent("push")
	buildOne.SetEventAction("")
	buildOne.SetStatus("running")
	buildOne.SetError("")
	buildOne.SetEnqueued(1563474077)
	buildOne.SetCreated(1563474076)
	buildOne.SetStarted(1563474078)
	buildOne.SetFinished(1563474079)
	buildOne.SetDeploy("")
	buildOne.SetDeployNumber(0)
	buildOne.SetDeployPayload(nil)
	buildOne.SetClone("https://github.com/github/octocat.git")
	buildOne.SetSource("https://github.com/github/octocat/48afb5bdc41ad69bf22588491333f7cf71135163")
	buildOne.SetTitle("push received from https://github.com/github/octocat")
	buildOne.SetMessage("First commit...")
	buildOne.SetCommit("48afb5bdc41ad69bf22588491333f7cf71135163")
	buildOne.SetSender("OctoKitty")
	buildOne.SetAuthor("OctoKitty")
	buildOne.SetEmail("OctoKitty@github.com")
	buildOne.SetLink("https://example.company.com/github/octocat/1")
	buildOne.SetBranch("main")
	buildOne.SetRef("refs/heads/main")
	buildOne.SetBaseRef("")
	buildOne.SetHeadRef("")
	buildOne.SetHost("example.company.com")
	buildOne.SetRuntime("docker")
	buildOne.SetDistribution("linux")
	buildOne.SetDeployPayload(raw.StringSliceMap{"foo": "test1", "bar": "test2"})
	buildOne.SetApprovedAt(1563474076)
	buildOne.SetApprovedBy("OctoCat")

	builds = append(builds, buildOne)

	want := new(library.Deployment)
	want.SetID(1)
	want.SetNumber(1)
	want.SetRepoID(1)
	want.SetURL("https://github.com/github/octocat/deployments/1")
	want.SetCommit("1234")
	want.SetRef("refs/heads/main")
	want.SetTask("deploy:vela")
	want.SetTarget("production")
	want.SetDescription("Deployment request from Vela")
	want.SetPayload(raw.StringSliceMap{"foo": "test1"})
	want.SetCreatedAt(1)
	want.SetCreatedBy("octocat")
	want.SetBuilds(builds)

	got := testDeployment().ToLibrary(builds)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("(ToLibrary: -want +got):\n%s", diff)
	}
}

func TestDatabase_Deployment_Validate(t *testing.T) {
	// setup types
	tests := []struct {
		failure    bool
		deployment *Deployment
		want       *Deployment
	}{
		{
			failure:    false,
			deployment: testDeployment(),
		},
		{ // no number set for deployment
			failure: true,
			deployment: &Deployment{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
				RepoID: sql.NullInt64{Int64: 1, Valid: true},
			},
			want: &Deployment{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
				RepoID: sql.NullInt64{Int64: 1, Valid: true},
			},
		},
		{ // no repoID set for deployment
			failure: true,
			deployment: &Deployment{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
				Number: sql.NullInt64{Int64: 1, Valid: true},
			},
			want: &Deployment{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
				RepoID: sql.NullInt64{Int64: 1, Valid: true},
			},
		},
		{ // too many builds
			failure: true,
			deployment: &Deployment{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
				Number: sql.NullInt64{Int64: 1, Valid: true},
				Builds: generateBuilds(100),
			},
			want: &Deployment{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
				RepoID: sql.NullInt64{Int64: 1, Valid: true},
				Builds: generateBuilds(50),
			},
		},
		{ // acceptable builds
			failure: true,
			deployment: &Deployment{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
				Number: sql.NullInt64{Int64: 1, Valid: true},
				Builds: generateBuilds(30),
			},
			want: &Deployment{
				ID:     sql.NullInt64{Int64: 1, Valid: true},
				RepoID: sql.NullInt64{Int64: 1, Valid: true},
				Builds: generateBuilds(30),
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
	builds := []*library.Build{}

	buildOne := new(library.Build)
	buildOne.SetID(1)
	buildOne.SetRepoID(1)
	buildOne.SetPipelineID(1)
	buildOne.SetNumber(1)
	buildOne.SetParent(1)
	buildOne.SetEvent("push")
	buildOne.SetEventAction("")
	buildOne.SetStatus("running")
	buildOne.SetError("")
	buildOne.SetEnqueued(1563474077)
	buildOne.SetCreated(1563474076)
	buildOne.SetStarted(1563474078)
	buildOne.SetFinished(1563474079)
	buildOne.SetDeploy("")
	buildOne.SetDeployNumber(0)
	buildOne.SetDeployPayload(nil)
	buildOne.SetClone("https://github.com/github/octocat.git")
	buildOne.SetSource("https://github.com/github/octocat/48afb5bdc41ad69bf22588491333f7cf71135163")
	buildOne.SetTitle("push received from https://github.com/github/octocat")
	buildOne.SetMessage("First commit...")
	buildOne.SetCommit("48afb5bdc41ad69bf22588491333f7cf71135163")
	buildOne.SetSender("OctoKitty")
	buildOne.SetAuthor("OctoKitty")
	buildOne.SetEmail("OctoKitty@github.com")
	buildOne.SetLink("https://example.company.com/github/octocat/1")
	buildOne.SetBranch("main")
	buildOne.SetRef("refs/heads/main")
	buildOne.SetBaseRef("")
	buildOne.SetHeadRef("")
	buildOne.SetHost("example.company.com")
	buildOne.SetRuntime("docker")
	buildOne.SetDistribution("linux")
	buildOne.SetDeployPayload(raw.StringSliceMap{"foo": "test1", "bar": "test2"})
	buildOne.SetApprovedAt(1563474076)
	buildOne.SetApprovedBy("OctoCat")

	builds = append(builds, buildOne)

	d := new(library.Deployment)
	d.SetID(1)
	d.SetNumber(1)
	d.SetRepoID(1)
	d.SetURL("https://github.com/github/octocat/deployments/1")
	d.SetCommit("1234")
	d.SetRef("refs/heads/main")
	d.SetTask("deploy:vela")
	d.SetTarget("production")
	d.SetDescription("Deployment request from Vela")
	d.SetPayload(raw.StringSliceMap{"foo": "test1"})
	d.SetCreatedAt(1)
	d.SetCreatedBy("octocat")
	d.SetBuilds(builds)

	want := testDeployment()

	// run test
	got := DeploymentFromLibrary(d)

	if diff := cmp.Diff(want, got); diff != "" {
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
		Commit:      sql.NullString{String: "1234", Valid: true},
		Ref:         sql.NullString{String: "refs/heads/main", Valid: true},
		Task:        sql.NullString{String: "deploy:vela", Valid: true},
		Target:      sql.NullString{String: "production", Valid: true},
		Description: sql.NullString{String: "Deployment request from Vela", Valid: true},
		Payload:     raw.StringSliceMap{"foo": "test1"},
		CreatedAt:   sql.NullInt64{Int64: 1, Valid: true},
		CreatedBy:   sql.NullString{String: "octocat", Valid: true},
		Builds:      pq.StringArray{"1"},
	}
}

// generateBuilds returns a list of valid builds that exceed the maximum size.
func generateBuilds(amount int) []string {
	// initialize empty builds
	builds := []string{}

	for i := 0; i < amount; i++ {
		builds = append(builds, "123456789")
	}

	return builds
}
