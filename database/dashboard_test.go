// SPDX-License-Identifier: Apache-2.0

package database

import (
	"database/sql"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/go-vela/types/library"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestDatabase_Dashboard_Nullify(t *testing.T) {
	// setup types
	var h *Dashboard

	want := &Dashboard{
		Name:      sql.NullString{String: "", Valid: false},
		CreatedAt: sql.NullInt64{Int64: 0, Valid: false},
		CreatedBy: sql.NullString{String: "", Valid: false},
		UpdatedAt: sql.NullInt64{Int64: 0, Valid: false},
		UpdatedBy: sql.NullString{String: "", Valid: false},
	}

	// setup tests
	tests := []struct {
		dashboard *Dashboard
		want      *Dashboard
	}{
		{
			dashboard: testDashboard(),
			want:      testDashboard(),
		},
		{
			dashboard: h,
			want:      nil,
		},
		{
			dashboard: new(Dashboard),
			want:      want,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.dashboard.Nullify()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Nullify is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_Dashboard_ToLibrary(t *testing.T) {
	// setup types
	want := new(library.Dashboard)
	want.SetID("c8da1302-07d6-11ea-882f-4893bca275b8")
	want.SetName("vela")
	want.SetCreatedAt(1)
	want.SetCreatedBy("octocat")
	want.SetUpdatedAt(2)
	want.SetUpdatedBy("octokitty")
	want.SetAdmins([]string{"octocat", "octokitty"})
	want.SetRepos(testDashReposJSON())

	uuid, _ := uuid.Parse("c8da1302-07d6-11ea-882f-4893bca275b8")
	h := &Dashboard{
		ID:        uuid,
		Name:      sql.NullString{String: "vela", Valid: true},
		CreatedAt: sql.NullInt64{Int64: 1, Valid: true},
		CreatedBy: sql.NullString{String: "octocat", Valid: true},
		UpdatedAt: sql.NullInt64{Int64: 2, Valid: true},
		UpdatedBy: sql.NullString{String: "octokitty", Valid: true},
		Admins:    []string{"octocat", "octokitty"},
		Repos:     testDashReposJSON(),
	}

	// run test
	got := h.ToLibrary()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToLibrary is %v, want %v", got, want)
	}
}

func TestDatabase_Dashboard_Validate(t *testing.T) {
	uuid, _ := uuid.Parse("c8da1302-07d6-11ea-882f-4893bca275b8")

	// setup tests
	tests := []struct {
		failure   bool
		dashboard *Dashboard
	}{
		{
			failure:   false,
			dashboard: testDashboard(),
		},
		{ // no name set for dashboard
			failure: true,
			dashboard: &Dashboard{
				ID: uuid,
			},
		},
		{ // invalid admin set for dashboard
			failure: true,
			dashboard: &Dashboard{
				ID:     uuid,
				Name:   sql.NullString{String: "vela", Valid: true},
				Admins: exceedAdmins(),
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.dashboard.Validate()

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

func TestDatabase_DashboardFromLibrary(t *testing.T) {
	uuid, err := uuid.Parse("c8da1302-07d6-11ea-882f-4893bca275b8")
	if err != nil {
		t.Errorf("error parsing uuid: %v", err)
	}

	// setup types
	want := &Dashboard{
		ID:        uuid,
		Name:      sql.NullString{String: "vela", Valid: true},
		CreatedAt: sql.NullInt64{Int64: 1, Valid: true},
		CreatedBy: sql.NullString{String: "octocat", Valid: true},
		UpdatedAt: sql.NullInt64{Int64: 2, Valid: true},
		UpdatedBy: sql.NullString{String: "octokitty", Valid: true},
		Admins:    []string{"octocat", "octokitty"},
		Repos:     testDashReposJSON(),
	}

	d := new(library.Dashboard)
	d.SetID("c8da1302-07d6-11ea-882f-4893bca275b8")
	d.SetName("vela")
	d.SetCreatedAt(1)
	d.SetCreatedBy("octocat")
	d.SetUpdatedAt(2)
	d.SetUpdatedBy("octokitty")
	d.SetAdmins([]string{"octocat", "octokitty"})
	d.SetRepos(testDashReposJSON())

	// run test
	got := DashboardFromLibrary(d)

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("DashboardFromLibrary() mismatch (-want +got):\n%s", diff)
	}

	d.SetID("")

	//nolint:staticcheck // linter is lying
	got = DashboardFromLibrary(d)

	if len(got.ID) != 16 {
		t.Errorf("Length is %d", len(got.ID))
	}

	d.SetID("123-abc")

	got = DashboardFromLibrary(d)

	if got != nil {
		t.Errorf("DashboardFromLibrary should have returned nil")
	}
}

// testDashboard is a test helper function to create a Dashboard
// type with all fields set to a fake value.
func testDashboard() *Dashboard {
	uuid, _ := uuid.Parse("c8da1302-07d6-11ea-882f-4893bca275b8")
	dRepos := testDashReposJSON()

	return &Dashboard{
		ID:        uuid,
		Name:      sql.NullString{String: "vela", Valid: true},
		CreatedAt: sql.NullInt64{Int64: time.Now().UTC().Unix(), Valid: true},
		CreatedBy: sql.NullString{String: "octocat", Valid: true},
		UpdatedAt: sql.NullInt64{Int64: time.Now().UTC().Unix(), Valid: true},
		UpdatedBy: sql.NullString{String: "octokitty", Valid: true},
		Admins:    []string{"octocat", "octokitty"},
		Repos:     dRepos,
	}
}

// testDashReposJSON is a test helper function to create a DashReposJSON
// type with all fields set to a fake value.
func testDashReposJSON() DashReposJSON {
	d := new(library.DashboardRepo)

	d.SetName("go-vela/server")
	d.SetID(1)
	d.SetBranches([]string{"main"})
	d.SetEvents([]string{"push", "tag"})

	return DashReposJSON{d}
}

// exceedAdmins returns a list of valid admins that exceed the maximum size.
func exceedAdmins() []string {
	// initialize empty favorites
	admins := []string{}

	// add enough favorites to exceed the character limit
	for i := 0; i < 500; i++ {
		// construct favorite
		// use i to adhere to unique favorites
		admin := "github/octocat-" + strconv.Itoa(i)

		admins = append(admins, admin)
	}

	return admins
}
