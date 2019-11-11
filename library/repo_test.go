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

	if r.GetID() != wantID {
		t.Errorf("GetID is %v, want %v", r.GetID(), wantID)
	}
	if r.GetUserID() != wantUserID {
		t.Errorf("GetUserID is %v, want %v", r.GetUserID(), wantUserID)
	}
	if r.GetOrg() != wantOrg {
		t.Errorf("GetOrg is %v, want %v", r.GetOrg(), wantOrg)
	}
	if r.GetName() != wantName {
		t.Errorf("GetName is %v, want %v", r.GetName(), wantName)
	}
	if r.GetFullName() != wantFullName {
		t.Errorf("GetFullName is %v, want %v", r.GetFullName(), wantFullName)
	}
	if r.GetLink() != wantLink {
		t.Errorf("GetLink is %v, want %v", r.GetLink(), wantLink)
	}
	if r.GetClone() != wantClone {
		t.Errorf("GetClone is %v, want %v", r.GetClone(), wantClone)
	}
	if r.GetBranch() != wantBranch {
		t.Errorf("GetBranch is %v, want %v", r.GetBranch(), wantBranch)
	}
	if r.GetTimeout() != wantTimeout {
		t.Errorf("GetTimeout is %v, want %v", r.GetTimeout(), wantTimeout)
	}
	if r.GetVisibility() != wantVisibility {
		t.Errorf("GetVisibility is %v, want %v", r.GetVisibility(), wantVisibility)
	}
	if r.GetPrivate() != wantPrivate {
		t.Errorf("GetPrivate is %v, want %v", r.GetPrivate(), wantPrivate)
	}
	if r.GetTrusted() != wantTrusted {
		t.Errorf("GetTrusted is %v, want %v", r.GetTrusted(), wantTrusted)
	}
	if r.GetActive() != wantActive {
		t.Errorf("GetActive is %v, want %v", r.GetActive(), wantActive)
	}
	if r.GetAllowPull() != wantAllowPull {
		t.Errorf("GetAllowPull is %v, want %v", r.GetAllowPull(), wantAllowPull)
	}
	if r.GetAllowPush() != wantAllowPush {
		t.Errorf("GetAllowPush is %v, want %v", r.GetAllowPush(), wantAllowPush)
	}
	if r.GetAllowDeploy() != wantAllowDeploy {
		t.Errorf("GetAllowDeploy is %v, want %v", r.GetAllowDeploy(), wantAllowDeploy)
	}
	if r.GetAllowTag() != wantAllowTag {
		t.Errorf("GetAllowTag is %v, want %v", r.GetAllowTag(), wantAllowTag)
	}
}

func TestLibrary_Repo_Setters_Empty(t *testing.T) {
	// setup types
	r := &Repo{}
	r = nil

	// run test
	r.SetID(0)
	r.SetUserID(0)
	r.SetOrg("")
	r.SetName("")
	r.SetFullName("")
	r.SetLink("")
	r.SetClone("")
	r.SetBranch("")
	r.SetTimeout(0)
	r.SetVisibility("")
	r.SetPrivate(false)
	r.SetTrusted(false)
	r.SetActive(false)
	r.SetAllowPull(false)
	r.SetAllowPush(false)
	r.SetAllowDeploy(false)
	r.SetAllowTag(false)

	if r.GetID() != 0 {
		t.Errorf("GetID is %v, want 0", r.GetID())
	}
	if r.GetUserID() != 0 {
		t.Errorf("GetUserID is %v, want 0", r.GetUserID())
	}
	if r.GetOrg() != "" {
		t.Errorf("GetOrg is %v, want \"\"", r.GetOrg())
	}
	if r.GetName() != "" {
		t.Errorf("GetName is %v, want \"\"", r.GetName())
	}
	if r.GetFullName() != "" {
		t.Errorf("GetFullName is %v, want \"\"", r.GetFullName())
	}
	if r.GetLink() != "" {
		t.Errorf("GetLink is %v, want \"\"", r.GetLink())
	}
	if r.GetClone() != "" {
		t.Errorf("GetClone is %v, want \"\"", r.GetClone())
	}
	if r.GetBranch() != "" {
		t.Errorf("GetBranch is %v, want \"\"", r.GetBranch())
	}
	if r.GetTimeout() != 0 {
		t.Errorf("GetTimeout is %v, want 0", r.GetTimeout())
	}
	if r.GetVisibility() != "" {
		t.Errorf("GetVisibility is %v, want \"\"", r.GetVisibility())
	}
	if r.GetPrivate() != false {
		t.Errorf("GetPrivate is %v, want false", r.GetPrivate())
	}
	if r.GetTrusted() != false {
		t.Errorf("GetTrusted is %v, want false", r.GetTrusted())
	}
	if r.GetActive() != false {
		t.Errorf("GetActive is %v, want false", r.GetActive())
	}
	if r.GetAllowPull() != false {
		t.Errorf("GetAllowPull is %v, want false", r.GetAllowPull())
	}
	if r.GetAllowPush() != false {
		t.Errorf("GetAllowPush is %v, want false", r.GetAllowPush())
	}
	if r.GetAllowDeploy() != false {
		t.Errorf("GetAllowDeploy is %v, want false", r.GetAllowDeploy())
	}
	if r.GetAllowTag() != false {
		t.Errorf("GetAllowTag is %v, want false", r.GetAllowTag())
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
