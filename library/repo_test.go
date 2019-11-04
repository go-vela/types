// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_Repo_Getters(t *testing.T) {
	// setup types
	booL := false
	num := 1
	num64 := int64(num)
	str := "foo"
	r := &Repo{
		ID:          &num64,
		UserID:      &num64,
		Org:         &str,
		Name:        &str,
		FullName:    &str,
		Link:        &str,
		Clone:       &str,
		Branch:      &str,
		Timeout:     &num64,
		Visibility:  &str,
		Private:     &booL,
		Trusted:     &booL,
		Active:      &booL,
		AllowPull:   &booL,
		AllowPush:   &booL,
		AllowDeploy: &booL,
		AllowTag:    &booL,
	}
	wantID := num64
	wantUserID := num64
	wantOrg := str
	wantName := str
	wantFullName := str
	wantLink := str
	wantClone := str
	wantBranch := str
	wantTimeout := num64
	wantVisibility := str
	wantPrivate := booL
	wantTrusted := booL
	wantActive := booL
	wantAllowPull := booL
	wantAllowPush := booL
	wantAllowDeploy := booL
	wantAllowTag := booL

	// run test
	gotID := r.GetID()
	gotUserID := r.GetUserID()
	gotOrg := r.GetOrg()
	gotName := r.GetName()
	gotFullName := r.GetFullName()
	gotLink := r.GetLink()
	gotClone := r.GetClone()
	gotBranch := r.GetBranch()
	gotTimeout := r.GetTimeout()
	gotVisibility := r.GetVisibility()
	gotPrivate := r.GetPrivate()
	gotTrusted := r.GetTrusted()
	gotActive := r.GetActive()
	gotAllowPull := r.GetAllowPull()
	gotAllowPush := r.GetAllowPush()
	gotAllowDeploy := r.GetAllowDeploy()
	gotAllowTag := r.GetAllowTag()

	if gotID != wantID {
		t.Errorf("GetID is %v, want %v", gotID, wantID)
	}
	if gotUserID != wantUserID {
		t.Errorf("GetUserID is %v, want %v", gotUserID, wantUserID)
	}
	if gotOrg != wantOrg {
		t.Errorf("GetOrg is %v, want %v", gotOrg, wantOrg)
	}
	if gotName != wantName {
		t.Errorf("GetName is %v, want %v", gotName, wantName)
	}
	if gotFullName != wantFullName {
		t.Errorf("GetFullName is %v, want %v", gotFullName, wantFullName)
	}
	if gotLink != wantLink {
		t.Errorf("GetLink is %v, want %v", gotLink, wantLink)
	}
	if gotClone != wantClone {
		t.Errorf("GetClone is %v, want %v", gotClone, wantClone)
	}
	if gotBranch != wantBranch {
		t.Errorf("GetBranch is %v, want %v", gotBranch, wantBranch)
	}
	if gotTimeout != wantTimeout {
		t.Errorf("GetTimeout is %v, want %v", gotTimeout, wantTimeout)
	}
	if gotVisibility != wantVisibility {
		t.Errorf("GetVisibility is %v, want %v", gotVisibility, wantVisibility)
	}
	if gotPrivate != wantPrivate {
		t.Errorf("GetPrivate is %v, want %v", gotPrivate, wantPrivate)
	}
	if gotTrusted != wantTrusted {
		t.Errorf("GetTrusted is %v, want %v", gotTrusted, wantTrusted)
	}
	if gotActive != wantActive {
		t.Errorf("GetActive is %v, want %v", gotActive, wantActive)
	}
	if gotAllowPull != wantAllowPull {
		t.Errorf("GetAllowPull is %v, want %v", gotAllowPull, wantAllowPull)
	}
	if gotAllowPush != wantAllowPush {
		t.Errorf("GetAllowPush is %v, want %v", gotAllowPush, wantAllowPush)
	}
	if gotAllowDeploy != wantAllowDeploy {
		t.Errorf("GetAllowDeploy is %v, want %v", gotAllowDeploy, wantAllowDeploy)
	}
	if gotAllowTag != wantAllowTag {
		t.Errorf("GetAllowTag is %v, want %v", gotAllowTag, wantAllowTag)
	}
}

func TestLibrary_Repo_Setters(t *testing.T) {
	// setup types
	booL := false
	num := 1
	num64 := int64(num)
	str := "foo"
	r := &Repo{}

	wantID := num64
	wantUserID := num64
	wantOrg := str
	wantName := str
	wantFullName := str
	wantLink := str
	wantClone := str
	wantBranch := str
	wantTimeout := num64
	wantVisibility := str
	wantPrivate := booL
	wantTrusted := booL
	wantActive := booL
	wantAllowPull := booL
	wantAllowPush := booL
	wantAllowDeploy := booL
	wantAllowTag := booL

	// run test
	r.SetID(wantID)
	r.SetUserID(wantUserID)
	r.SetOrg(wantOrg)
	r.SetName(wantName)
	r.SetFullName(wantFullName)
	r.SetLink(wantLink)
	r.SetClone(wantClone)
	r.SetBranch(wantBranch)
	r.SetTimeout(wantTimeout)
	r.SetVisibility(wantVisibility)
	r.SetPrivate(wantPrivate)
	r.SetTrusted(wantTrusted)
	r.SetActive(wantActive)
	r.SetAllowPull(wantAllowPull)
	r.SetAllowPush(wantAllowPush)
	r.SetAllowDeploy(wantAllowDeploy)
	r.SetAllowTag(wantAllowTag)

	if *r.ID != wantID {
		t.Errorf("GetID is %v, want %v", *r.ID, wantID)
	}
	if *r.UserID != wantUserID {
		t.Errorf("GetUserID is %v, want %v", *r.UserID, wantUserID)
	}
	if *r.Org != wantOrg {
		t.Errorf("GetOrg is %v, want %v", *r.Org, wantOrg)
	}
	if *r.Name != wantName {
		t.Errorf("GetName is %v, want %v", *r.Name, wantName)
	}
	if *r.FullName != wantFullName {
		t.Errorf("GetFullName is %v, want %v", *r.FullName, wantFullName)
	}
	if *r.Link != wantLink {
		t.Errorf("GetLink is %v, want %v", *r.Link, wantLink)
	}
	if *r.Clone != wantClone {
		t.Errorf("GetClone is %v, want %v", *r.Clone, wantClone)
	}
	if *r.Branch != wantBranch {
		t.Errorf("GetBranch is %v, want %v", *r.Branch, wantBranch)
	}
	if *r.Timeout != wantTimeout {
		t.Errorf("GetTimeout is %v, want %v", *r.Timeout, wantTimeout)
	}
	if *r.Visibility != wantVisibility {
		t.Errorf("GetVisibility is %v, want %v", *r.Visibility, wantVisibility)
	}
	if *r.Private != wantPrivate {
		t.Errorf("GetPrivate is %v, want %v", *r.Private, wantPrivate)
	}
	if *r.Trusted != wantTrusted {
		t.Errorf("GetTrusted is %v, want %v", *r.Trusted, wantTrusted)
	}
	if *r.Active != wantActive {
		t.Errorf("GetActive is %v, want %v", *r.Active, wantActive)
	}
	if *r.AllowPull != wantAllowPull {
		t.Errorf("GetAllowPull is %v, want %v", *r.AllowPull, wantAllowPull)
	}
	if *r.AllowPush != wantAllowPush {
		t.Errorf("GetAllowPush is %v, want %v", *r.AllowPush, wantAllowPush)
	}
	if *r.AllowDeploy != wantAllowDeploy {
		t.Errorf("GetAllowDeploy is %v, want %v", *r.AllowDeploy, wantAllowDeploy)
	}
	if *r.AllowTag != wantAllowTag {
		t.Errorf("GetAllowTag is %v, want %v", *r.AllowTag, wantAllowTag)
	}
}

func TestLibrary_Repo_Getters_Empty(t *testing.T) {
	// setup types
	r := &Repo{}

	// run test
	gotID := r.GetID()
	gotUserID := r.GetUserID()
	gotOrg := r.GetOrg()
	gotName := r.GetName()
	gotFullName := r.GetFullName()
	gotLink := r.GetLink()
	gotClone := r.GetClone()
	gotBranch := r.GetBranch()
	gotTimeout := r.GetTimeout()
	gotVisibility := r.GetVisibility()
	gotPrivate := r.GetPrivate()
	gotTrusted := r.GetTrusted()
	gotActive := r.GetActive()
	gotAllowPull := r.GetAllowPull()
	gotAllowPush := r.GetAllowPush()
	gotAllowDeploy := r.GetAllowDeploy()
	gotAllowTag := r.GetAllowTag()

	if gotID != 0 {
		t.Errorf("GetID is %v, want 0", gotID)
	}
	if gotUserID != 0 {
		t.Errorf("GetUserID is %v, want 0", gotUserID)
	}
	if gotOrg != "" {
		t.Errorf("GetOrg is %v, want \"\"", gotOrg)
	}
	if gotName != "" {
		t.Errorf("GetName is %v, want \"\"", gotName)
	}
	if gotFullName != "" {
		t.Errorf("GetFullName is %v, want \"\"", gotFullName)
	}
	if gotLink != "" {
		t.Errorf("GetLink is %v, want \"\"", gotLink)
	}
	if gotClone != "" {
		t.Errorf("GetClone is %v, want \"\"", gotClone)
	}
	if gotBranch != "" {
		t.Errorf("GetBranch is %v, want \"\"", gotBranch)
	}
	if gotTimeout != 0 {
		t.Errorf("GetTimeout is %v, want 0", gotTimeout)
	}
	if gotVisibility != "" {
		t.Errorf("GetVisibility is %v, want \"\"", gotVisibility)
	}
	if gotPrivate != false {
		t.Errorf("GetPrivate is %v, want false", gotPrivate)
	}
	if gotTrusted != false {
		t.Errorf("GetTrusted is %v, want false", gotTrusted)
	}
	if gotActive != false {
		t.Errorf("GetActive is %v, want false", gotActive)
	}
	if gotAllowPull != false {
		t.Errorf("GetAllowPull is %v, want false", gotAllowPull)
	}
	if gotAllowPush != false {
		t.Errorf("GetAllowPush is %v, want false", gotAllowPush)
	}
	if gotAllowDeploy != false {
		t.Errorf("GetAllowDeploy is %v, want false", gotAllowDeploy)
	}
	if gotAllowTag != false {
		t.Errorf("GetAllowTag is %v, want false", gotAllowTag)
	}
}

func TestLibrary_Repo_String(t *testing.T) {
	// setup types
	booL := false
	num := 1
	num64 := int64(num)
	str := "foo"
	r := &Repo{
		ID:          &num64,
		UserID:      &num64,
		Org:         &str,
		Name:        &str,
		FullName:    &str,
		Link:        &str,
		Clone:       &str,
		Branch:      &str,
		Timeout:     &num64,
		Visibility:  &str,
		Private:     &booL,
		Trusted:     &booL,
		Active:      &booL,
		AllowPull:   &booL,
		AllowPush:   &booL,
		AllowDeploy: &booL,
		AllowTag:    &booL,
	}
	want := fmt.Sprintf("%+v", *r)

	// run test
	got := r.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}
