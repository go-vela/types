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

func TestLibrary_Login_Setters(t *testing.T) {
	// setup types
	str := "foo"
	l := &Login{}
	wantUsername := str
	wantPassword := str
	wantOTP := str
	wantToken := str

	// run test
	l.SetUsername(wantUsername)
	l.SetPassword(wantPassword)
	l.SetOTP(wantOTP)
	l.SetToken(wantToken)

	if *l.Username != wantUsername {
		t.Errorf("GetUsername is %v, want %v", *l.Username, wantUsername)
	}
	if *l.Password != wantPassword {
		t.Errorf("GetPassword is %v, want %v", *l.Password, wantPassword)
	}
	if *l.OTP != wantOTP {
		t.Errorf("GetOTP is %v, want %v", *l.OTP, wantOTP)
	}
	if *l.Token != wantToken {
		t.Errorf("GetToken is %v, want %v", *l.Token, wantToken)
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
