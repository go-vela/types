// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"bytes"
	"fmt"

	"github.com/microcosm-cc/bluemonday"
)

// sanitize is a helper function to verify the provided input
// does not contain HTML content. If the input does contain
// HTML, then the function will return an error.
func sanitize(v interface{}) error {
	// create new HTML input microcosm-cc/bluemonday policy
	p := bluemonday.StrictPolicy()

	// create a new object string from the input
	object := fmt.Sprintf("%v", v)

	// create new bytes buffer from the object string
	buffer := bytes.NewBufferString(object)

	// check if the buffer bytes are different than the HTML sanitized bytes
	if !bytes.Equal(buffer.Bytes(), p.SanitizeBytes(buffer.Bytes())) {
		return fmt.Errorf("resource failed HTML input validation")
	}

	return nil
}
