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
			value: `%`,
			want:  `%`,
		},
		{
			value: `"hello"`,
			want:  `"hello"`,
		},
		{
			value: `OctoKitty@github.com`,
			want:  `OctoKitty@github.com`,
		},
		{
			value: `https://github.com/go-vela`,
			want:  `https://github.com/go-vela`,
		},
		{
			value: `Merge pull request #1 from me/patch-1\n\n<h1>hello</h1> is now <h2>hello</h2>`,
			want:  `Merge pull request #1 from me/patch-1\n\nhello is now hello`,
		},
		{
			value: `+ added foo %25 + updated bar %22 +`,
			want:  `+ added foo %25 + updated bar %22 +`,
		},
		{
			value: `Co-authored-by: OctoKitty <OctoKitty@github.com>`,
			want:  `Co-authored-by: OctoKitty `,
		},
		{
			value: `<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
			want:  `Google`,
		},
		{
			value: `<script>alert('XSS')</script>`,
			want:  ``,
		},
		{
			value: `<SCRIPT/XSS SRC="http://ha.ckers.org/xss.js"></SCRIPT>`,
			want:  ``,
		},
		{
			value: `%3cDIV%20STYLE%3d%22width%3a%20expression(alert('XSS'))%3b%22%3e`,
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
