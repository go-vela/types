// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import "fmt"

// Login is the library representation of a user login.
//
// swagger:model Login
type Login struct {
	Token *string `json:"token,omitempty"`
}

// GetToken returns the Token field.
//
// When the provided Login type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Login) GetToken() string {
	// return zero value if Login type or Token field is nil
	if l == nil || l.Token == nil {
		return ""
	}

	return *l.Token
}

// SetToken sets the Token field.
//
// When the provided Login type is nil, it
// will set nothing and immediately return.
func (l *Login) SetToken(v string) {
	// return if Login type is nil
	if l == nil {
		return
	}

	l.Token = &v
}

// String implements the Stringer interface for the Login type.
func (l *Login) String() string {
	return fmt.Sprintf(`{
  Token: %s,
}`,
		l.GetToken(),
	)
}
