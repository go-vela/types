// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"database/sql"
	"testing"
)

func TestDatabase_Sanitize(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		value   interface{}
	}{
		{ // invalid HTML in fields set for build
			failure: true,
			value: &Build{
				ID:      sql.NullInt64{Int64: 1, Valid: true},
				RepoID:  sql.NullInt64{Int64: 1, Valid: true},
				Number:  sql.NullInt32{Int32: 1, Valid: true},
				Message: sql.NullString{String: `%3cDIV%20STYLE%3d%22width%3a%20expression(alert('XSS'))%3b%22%3e`, Valid: true},
			},
		},
		{ // invalid HTML in fields set for repo
			failure: true,
			value: &Repo{
				ID:         sql.NullInt64{Int64: 1, Valid: true},
				UserID:     sql.NullInt64{Int64: 1, Valid: true},
				Hash:       sql.NullString{String: "superSecretHash", Valid: true},
				Org:        sql.NullString{String: "github", Valid: true},
				Name:       sql.NullString{String: "octocat", Valid: true},
				FullName:   sql.NullString{String: "github/octocat", Valid: true},
				Visibility: sql.NullString{String: `<SCRIPT/XSS SRC="http://ha.ckers.org/xss.js"></SCRIPT>`, Valid: true},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := sanitize(test.value)

		if test.failure {
			if err == nil {
				t.Errorf("Sanitize should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Sanitize returned err: %v", err)
		}
	}
}
