// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_User_Getters(t *testing.T) {
	// setup types
	booL := false
	num64 := int64(1)
	str := "foo"
	u := &User{
		ID:     &num64,
		Name:   &str,
		Token:  &str,
		Hash:   &str,
		Active: &booL,
		Admin:  &booL,
	}
	wantID := num64
	wantName := str
	wantToken := str
	wantHash := str
	wantActive := booL
	wantAdmin := booL

	// run test
	gotID := u.GetID()
	gotName := u.GetName()
	gotToken := u.GetToken()
	gotHash := u.GetHash()
	gotActive := u.GetActive()
	gotAdmin := u.GetAdmin()

	if gotID != wantID {
		t.Errorf("GetID is %v, want %v", gotID, wantID)
	}
	if gotName != wantName {
		t.Errorf("GetName is %v, want %v", gotName, wantName)
	}
	if gotToken != wantToken {
		t.Errorf("GetToken is %v, want %v", gotToken, wantToken)
	}
	if gotHash != wantHash {
		t.Errorf("GetHash is %v, want %v", gotHash, wantHash)
	}
	if gotActive != wantActive {
		t.Errorf("GetActive is %v, want %v", gotActive, wantActive)
	}
	if gotAdmin != wantAdmin {
		t.Errorf("GetAdmin is %v, want %v", gotAdmin, wantAdmin)
	}
}

func TestLibrary_User_Setters(t *testing.T) {
	// setup types
	booL := false
	num64 := int64(1)
	str := "foo"
	u := &User{}

	wantID := num64
	wantName := str
	wantToken := str
	wantHash := str
	wantActive := booL
	wantAdmin := booL

	// run test
	u.SetID(wantID)
	u.SetName(wantName)
	u.SetToken(wantToken)
	u.SetHash(wantHash)
	u.SetActive(wantActive)
	u.SetAdmin(wantAdmin)

	if *u.ID != wantID {
		t.Errorf("GetID is %v, want %v", *u.ID, wantID)
	}
	if *u.Name != wantName {
		t.Errorf("GetName is %v, want %v", *u.Name, wantName)
	}
	if *u.Token != wantToken {
		t.Errorf("GetToken is %v, want %v", *u.Token, wantToken)
	}
	if *u.Hash != wantHash {
		t.Errorf("GetHash is %v, want %v", *u.Hash, wantHash)
	}
	if *u.Active != wantActive {
		t.Errorf("GetActive is %v, want %v", *u.Active, wantActive)
	}
	if *u.Admin != wantAdmin {
		t.Errorf("GetAdmin is %v, want %v", *u.Admin, wantAdmin)
	}
}

func TestLibrary_User_Getters_Empty(t *testing.T) {
	// setup types
	u := &User{}

	// run test
	gotID := u.GetID()
	gotName := u.GetName()
	gotToken := u.GetToken()
	gotHash := u.GetHash()
	gotActive := u.GetActive()
	gotAdmin := u.GetAdmin()

	if gotID != 0 {
		t.Errorf("GetID is %v, want 0", gotID)
	}
	if gotName != "" {
		t.Errorf("GetName is %v, want \"\"", gotName)
	}
	if gotToken != "" {
		t.Errorf("GetToken is %v, want \"\"", gotToken)
	}
	if gotHash != "" {
		t.Errorf("GetHash is %v, want \"\"", gotHash)
	}
	if gotActive != false {
		t.Errorf("GetActive is %v, want false", gotActive)
	}
	if gotAdmin != false {
		t.Errorf("GetAdmin is %v, want false", gotAdmin)
	}
}

func TestLibrary_User_String(t *testing.T) {
	// setup types
	booL := false
	num64 := int64(1)
	str := "foo"
	u := &User{
		ID:     &num64,
		Name:   &str,
		Token:  &str,
		Hash:   &str,
		Active: &booL,
		Admin:  &booL,
	}
	want := fmt.Sprintf("%+v", *u)

	// run test
	got := u.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}
