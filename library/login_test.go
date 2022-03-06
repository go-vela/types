// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_Login_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		login *Login
		want  *Login
	}{
		{
			login: testLogin(),
			want:  testLogin(),
		},
		{
			login: new(Login),
			want:  new(Login),
		},
	}

	// run tests
	for _, test := range tests {
		if test.login.GetToken() != test.want.GetToken() {
			t.Errorf("GetToken is %v, want %v", test.login.GetToken(), test.want.GetToken())
		}
	}
}

func TestLibrary_Login_Setters(t *testing.T) {
	// setup types
	var l *Login

	// setup tests
	tests := []struct {
		login *Login
		want  *Login
	}{
		{
			login: testLogin(),
			want:  testLogin(),
		},
		{
			login: l,
			want:  new(Login),
		},
	}

	// run tests
	for _, test := range tests {
		test.login.SetToken(test.want.GetToken())

		if test.login.GetToken() != test.want.GetToken() {
			t.Errorf("SetToken is %v, want %v", test.login.GetToken(), test.want.GetToken())
		}
	}
}

func TestLogin_String(t *testing.T) {
	// setup types
	l := testLogin()

	want := fmt.Sprintf(`{
  Token: %s,
}`,
		l.GetToken(),
	)

	// run test
	got := l.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

// testLogin is a test helper function to create a Login
// type with all fields set to a fake value.
func testLogin() *Login {
	l := new(Login)

	l.SetToken("superSecretToken")

	return l
}
