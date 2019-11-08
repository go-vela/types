// Copyright (c) 2018 Target Brands, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package library

import "fmt"

// Login is the library representation of a user login.
type Login struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	OTP      *string `json:"otp,omitempty"`
	Token    *string `json:"token,omitempty"`
}

// GetUsername returns the Username field.
//
// When the provided Login type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Login) GetUsername() string {
	// return zero value if Login type or Username field is nil
	if l == nil || l.Username == nil {
		return ""
	}
	return *l.Username
}

// GetPassword returns the Password field.
//
// When the provided Login type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Login) GetPassword() string {
	// return zero value if Login type or Password field is nil
	if l == nil || l.Password == nil {
		return ""
	}
	return *l.Password
}

// GetOTP returns the Username field.
//
// When the provided Login type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Login) GetOTP() string {
	// return zero value if Login type or OTP field is nil
	if l == nil || l.OTP == nil {
		return ""
	}
	return *l.OTP
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

// SetUsername Sets the Username field.
func (l *Login) SetUsername(v string) {
	if l == nil {
		return
	}
	l.Username = &v
}

// SetPassword Sets the Password field.
func (l *Login) SetPassword(v string) {
	if l == nil {
		return
	}
	l.Password = &v
}

// SetOTP Sets the OTP field.
func (l *Login) SetOTP(v string) {
	if l == nil {
		return
	}
	l.OTP = &v
}

// SetToken Sets the Token field.
func (l *Login) SetToken(v string) {
	if l == nil {
		return
	}
	l.Token = &v
}

// String implements the Stringer interface for the Login type.
func (l *Login) String() string {
	return fmt.Sprintf("%+v", *l)
}
