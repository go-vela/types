// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"testing"
)

func TestDatabase_Sanitize(t *testing.T) {
	// setup tests
	tests := []struct {
		value string
		want  string
	}{
		{
			value: `"hello"`,
			want:  `"hello"`,
		},
		{
			value: `<SCRIPT/XSS SRC="http://ha.ckers.org/xss.js"></SCRIPT>`,
			want:  ``,
		},
		{
			value: `<script>alert('XSS')</script>`,
			want:  ``,
		},
	}

	// run tests
	for _, test := range tests {
		got := sanitize(test.value)

		if got != test.want {
			t.Errorf("error sanitizing. got %s, wanted %s", got, test.want)
		}
	}
}
