// SPDX-License-Identifier: Apache-2.0

package database

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
	"github.com/lib/pq"
)

var (
	// ErrEmptyDashName defines the error type when a
	// User type has an empty Name field provided.
	ErrEmptyDashName = errors.New("empty dashboard name provided")

	// ErrExceededAdminLimit defines the error type when a
	// User type has Admins field provided that exceeds the database limit.
	ErrExceededAdminLimit = errors.New("exceeded admins limit")
)

// Dashboard is the database representation of a user.
type Dashboard struct {
	ID        sql.NullInt64  `sql:"id"`
	Name      sql.NullString `sql:"name"`
	CreatedAt sql.NullInt64  `sql:"created_at"`
	CreatedBy sql.NullString `sql:"created_by"`
	UpdatedAt sql.NullInt64  `sql:"updated_at"`
	UpdatedBy sql.NullString `sql:"updated_by"`
	Admins    pq.StringArray `sql:"admins" gorm:"type:varchar(5000)"`
	Repos     DashReposJSON
}

type DashReposJSON []*library.DashboardRepo

// Value - Implementation of valuer for database/sql for DashReposJSON.
func (r DashReposJSON) Value() (driver.Value, error) {
	valueString, err := json.Marshal(r)
	return string(valueString), err
}

// Scan - Implement the database/sql scanner interface for DashReposJSON.
func (r *DashReposJSON) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &r)
}

// Nullify ensures the valid flag for
// the sql.Null types are properly set.
//
// When a field within the Dashboard type is the zero
// value for the field, the valid flag is set to
// false causing it to be NULL in the database.
func (d *Dashboard) Nullify() *Dashboard {
	if d == nil {
		return nil
	}

	// check if the ID field should be false
	if d.ID.Int64 == 0 {
		d.ID.Valid = false
	}

	// check if the Name field should be false
	if len(d.Name.String) == 0 {
		d.Name.Valid = false
	}

	// check if the CreatedAt field should be false
	if d.CreatedAt.Int64 == 0 {
		d.CreatedAt.Valid = false
	}

	// check if the CreatedBy field should be false
	if len(d.CreatedBy.String) == 0 {
		d.CreatedBy.Valid = false
	}

	// check if the UpdatedAt field should be false
	if d.UpdatedAt.Int64 == 0 {
		d.UpdatedAt.Valid = false
	}

	// check if the UpdatedBy field should be false
	if len(d.UpdatedBy.String) == 0 {
		d.UpdatedBy.Valid = false
	}

	return d
}

// ToLibrary converts the Dashboard type
// to a library Dashboard type.
func (d *Dashboard) ToLibrary() *library.Dashboard {
	dashboard := new(library.Dashboard)

	dashboard.SetID(d.ID.Int64)
	dashboard.SetName(d.Name.String)
	dashboard.SetAdmins(d.Admins)
	dashboard.SetCreatedAt(d.CreatedAt.Int64)
	dashboard.SetCreatedBy(d.CreatedBy.String)
	dashboard.SetUpdatedAt(d.UpdatedAt.Int64)
	dashboard.SetUpdatedBy(d.UpdatedBy.String)
	dashboard.SetRepos(d.Repos)

	return dashboard
}

// Validate verifies the necessary fields for
// the Dashboard type are populated correctly.
func (d *Dashboard) Validate() error {
	// verify the Name field is populated
	if len(d.Name.String) == 0 {
		return ErrEmptyDashName
	}

	// calculate total size of favorites
	total := 0
	for _, f := range d.Admins {
		total += len(f)
	}

	// verify the Favorites field is within the database constraints
	// len is to factor in number of comma separators included in the database field,
	// removing 1 due to the last item not having an appended comma
	if (total + len(d.Admins) - 1) > constants.FavoritesMaxSize {
		return ErrExceededAdminLimit
	}

	// ensure that all Dashboard string fields
	// that can be returned as JSON are sanitized
	// to avoid unsafe HTML content
	d.Name = sql.NullString{String: sanitize(d.Name.String), Valid: d.Name.Valid}

	// ensure that all Favorites are sanitized
	// to avoid unsafe HTML content
	for i, v := range d.Admins {
		d.Admins[i] = sanitize(v)
	}

	return nil
}

// DashboardFromLibrary converts the library Dashboard type
// to a database Dashboard type.
func DashboardFromLibrary(d *library.Dashboard) *Dashboard {
	user := &Dashboard{
		ID:        sql.NullInt64{Int64: d.GetID(), Valid: true},
		Name:      sql.NullString{String: d.GetName(), Valid: true},
		CreatedAt: sql.NullInt64{Int64: d.GetCreatedAt(), Valid: true},
		CreatedBy: sql.NullString{String: d.GetCreatedBy(), Valid: true},
		UpdatedAt: sql.NullInt64{Int64: d.GetUpdatedAt(), Valid: true},
		UpdatedBy: sql.NullString{String: d.GetUpdatedBy(), Valid: true},
		Admins:    pq.StringArray(d.GetAdmins()),
		Repos:     d.GetRepos(),
	}

	return user.Nullify()
}
