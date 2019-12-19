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

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_Login_Getters(t *testing.T) {
	// setup types
	str := "foo"
	l := &Login{
		Username: &str,
		Password: &str,
		OTP:      &str,
		Token:    &str,
	}
	wantUsername := str
	wantPassword := str
	wantOTP := str
	wantToken := str

	// run test
	gotUsername := l.GetUsername()
	gotPassword := l.GetPassword()
	gotOTP := l.GetOTP()
	gotToken := l.GetToken()

	if gotUsername != wantUsername {
		t.Errorf("GetUsername is %v, want %v", gotUsername, wantUsername)
	}
	if gotPassword != wantPassword {
		t.Errorf("GetPassword is %v, want %v", gotPassword, wantPassword)
	}
	if gotOTP != wantOTP {
		t.Errorf("GetOTP is %v, want %v", gotOTP, wantOTP)
	}
	if gotToken != wantToken {
		t.Errorf("GetToken is %v, want %v", gotToken, wantToken)
	}
}

func TestLibrary_Login_Getters_Empty(t *testing.T) {
	// setup types
	l := new(Login)

	// run test
	gotUsername := l.GetUsername()
	gotPassword := l.GetPassword()
	gotOTP := l.GetOTP()
	gotToken := l.GetToken()

	if gotUsername != "" {
		t.Errorf("GetUsername is %v, want \"\"", gotUsername)
	}
	if gotPassword != "" {
		t.Errorf("GetPassword is %v, want \"\"", gotPassword)
	}
	if gotOTP != "" {
		t.Errorf("GetOTP is %v, want \"\"", gotOTP)
	}
	if gotToken != "" {
		t.Errorf("GetToken is %v, want \"\"", gotToken)
	}
}

func TestLibrary_Login_Setters(t *testing.T) {
	// setup types
	str := "foo"
	l := new(Login)
	wantUsername := str
	wantPassword := str
	wantOTP := str
	wantToken := str

	// run test
	l.SetUsername(wantUsername)
	l.SetPassword(wantPassword)
	l.SetOTP(wantOTP)
	l.SetToken(wantToken)

	if l.GetUsername() != wantUsername {
		t.Errorf("SetUsername is %v, want %v", l.GetUsername(), wantUsername)
	}
	if l.GetPassword() != wantPassword {
		t.Errorf("SetPassword is %v, want %v", l.GetPassword(), wantPassword)
	}
	if l.GetOTP() != wantOTP {
		t.Errorf("SetOTP is %v, want %v", l.GetOTP(), wantOTP)
	}
	if l.GetToken() != wantToken {
		t.Errorf("SetToken is %v, want %v", l.GetToken(), wantToken)
	}
}

func TestLibrary_Login_Setters_Empty(t *testing.T) {
	// setup types
	var l *Login

	// run test
	l.SetUsername("")
	l.SetPassword("")
	l.SetOTP("")
	l.SetToken("")

	if l.GetUsername() != "" {
		t.Errorf("SetUsername is %v, want \"\"", l.GetUsername())
	}
	if l.GetPassword() != "" {
		t.Errorf("SetPassword is %v, want \"\"", l.GetPassword())
	}
	if l.GetOTP() != "" {
		t.Errorf("SetOTP is %v, want \"\"", l.GetOTP())
	}
	if l.GetToken() != "" {
		t.Errorf("SetToken is %v, want \"\"", l.GetToken())
	}
}

func TestLogin_String(t *testing.T) {
	// setup types
	str := "foo"
	l := &Login{
		Username: &str,
		Password: &str,
		OTP:      &str,
		Token:    &str,
	}
	want := fmt.Sprintf("%+v", *l)

	// run test
	got := l.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}
