// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package types

import "fmt"

// Error is the json error message from the server for a given http response.
type Error struct {
	Message *string `json:"error"`
}

func (e *Error) String() string {
	return fmt.Sprintf("%+v", *e)
}
