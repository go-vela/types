// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"html"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

// sanitize is a helper function to verify the provided input
// does not contain HTML content. If the input does contain
// HTML, then the function will return an error.
func sanitize(field string) string {
	// create new HTML input microcosm-cc/bluemonday policy
	p := bluemonday.StrictPolicy()

	// create an HTML escaped string from the field
	htmlEscaped := html.EscapeString(field) // &#34;hello&#34;

	// create a bluemonday escaped string from the field
	bluemondayEscaped := p.Sanitize(field) // &#34;hello&#34;

	// check if the field contains html
	if !strings.EqualFold(htmlEscaped, bluemondayEscaped) {
		// create new HTML input microcosm-cc/bluemonday policy
		return bluemondayEscaped
	}

	// return the unmodified field
	return field
}
