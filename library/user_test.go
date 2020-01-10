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
		ID:        &num64,
		Name:      &str,
		Token:     &str,
		Hash:      &str,
		Favorites: &str,
		Active:    &booL,
		Admin:     &booL,
	}
	wantID := num64
	wantName := str
	wantToken := str
	wantHash := str
	wantFavorites := str
	wantActive := booL
	wantAdmin := booL

	// run test
	gotID := u.GetID()
	gotName := u.GetName()
	gotToken := u.GetToken()
	gotHash := u.GetHash()
	gotFavorites := u.GetFavorites()
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
	if gotFavorites != wantFavorites {
		t.Errorf("GetFavorites is %v, want %v", gotFavorites, wantFavorites)
	}
	if gotActive != wantActive {
		t.Errorf("GetActive is %v, want %v", gotActive, wantActive)
	}
	if gotAdmin != wantAdmin {
		t.Errorf("GetAdmin is %v, want %v", gotAdmin, wantAdmin)
	}
}

func TestLibrary_User_Getters_Empty(t *testing.T) {
	// setup types
	u := new(User)

	// run test
	gotID := u.GetID()
	gotName := u.GetName()
	gotToken := u.GetToken()
	gotHash := u.GetHash()
	gotFavorites := u.GetFavorites()
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
	if gotFavorites != "" {
		t.Errorf("GetFavorites is %v, want \"\"", gotFavorites)
	}
	if gotActive != false {
		t.Errorf("GetActive is %v, want false", gotActive)
	}
	if gotAdmin != false {
		t.Errorf("GetAdmin is %v, want false", gotAdmin)
	}
}

func TestLibrary_User_Setters(t *testing.T) {
	// setup types
	booL := false
	num64 := int64(1)
	str := "foo"
	u := new(User)

	wantID := num64
	wantName := str
	wantToken := str
	wantHash := str
	wantFavorites := str
	wantActive := booL
	wantAdmin := booL

	// run test
	u.SetID(wantID)
	u.SetName(wantName)
	u.SetToken(wantToken)
	u.SetHash(wantHash)
	u.SetFavorites(wantFavorites)
	u.SetActive(wantActive)
	u.SetAdmin(wantAdmin)

	if u.GetID() != wantID {
		t.Errorf("SetID is %v, want %v", u.GetID(), wantID)
	}
	if u.GetName() != wantName {
		t.Errorf("SetName is %v, want %v", u.GetName(), wantName)
	}
	if u.GetToken() != wantToken {
		t.Errorf("SetToken is %v, want %v", u.GetToken(), wantToken)
	}
	if u.GetHash() != wantHash {
		t.Errorf("SetHash is %v, want %v", u.GetHash(), wantHash)
	}
	if u.GetFavorites() != wantFavorites {
		t.Errorf("SetFavorites is %v, want %v", u.GetFavorites(), wantFavorites)
	}
	if u.GetActive() != wantActive {
		t.Errorf("SetActive is %v, want %v", u.GetActive(), wantActive)
	}
	if u.GetAdmin() != wantAdmin {
		t.Errorf("SetAdmin is %v, want %v", u.GetAdmin(), wantAdmin)
	}
}

func TestLibrary_User_Setters_Empty(t *testing.T) {
	// setup types
	var u *User

	// run test
	u.SetID(0)
	u.SetName("")
	u.SetToken("")
	u.SetHash("")
	u.SetFavorites("")
	u.SetActive(false)
	u.SetAdmin(false)

	if u.GetID() != 0 {
		t.Errorf("SetID is %v, want 0", u.GetID())
	}
	if u.GetName() != "" {
		t.Errorf("SetName is %v, want \"\"", u.GetName())
	}
	if u.GetToken() != "" {
		t.Errorf("SetToken is %v, want \"\"", u.GetToken())
	}
	if u.GetHash() != "" {
		t.Errorf("SetHash is %v, want \"\"", u.GetHash())
	}
	if u.GetFavorites() != "" {
		t.Errorf("GetFavorites is %v, want \"\"", u.GetFavorites())
	}
	if u.GetActive() != false {
		t.Errorf("SetActive is %v, want false", u.GetActive())
	}
	if u.GetAdmin() != false {
		t.Errorf("SetAdmin is %v, want false", u.GetAdmin())
	}
}

func TestLibrary_User_String(t *testing.T) {
	// setup types
	booL := false
	num64 := int64(1)
	str := "foo"
	u := &User{
		ID:        &num64,
		Name:      &str,
		Token:     &str,
		Hash:      &str,
		Favorites: &str,
		Active:    &booL,
		Admin:     &booL,
	}
	want := fmt.Sprintf("%+v", *u)

	// run test
	got := u.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}
