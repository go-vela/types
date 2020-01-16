// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestLibrary_Hook_Getters(t *testing.T) {
	// setup types
	wantID := int64(1)
	wantRepoID := int64(1)
	wantBuildID := int64(1)
	wantNumber := 1
	wantSourceID := "c8da1302-07d6-11ea-882f-4893bca275b8"
	wantCreated := time.Now().UTC().Unix()
	wantHost := "github.com"
	wantEvent := "push"
	wantBranch := "master"
	wantError := ""
	wantStatus := "success"
	wantLink := "https://github.com/github/octocat/settings/hooks/1"

	h := new(Hook)

	h.SetID(wantID)
	h.SetRepoID(wantRepoID)
	h.SetBuildID(wantBuildID)
	h.SetNumber(wantNumber)
	h.SetSourceID(wantSourceID)
	h.SetCreated(wantCreated)
	h.SetHost(wantHost)
	h.SetEvent(wantEvent)
	h.SetBranch(wantBranch)
	h.SetError(wantError)
	h.SetStatus(wantStatus)
	h.SetLink(wantLink)

	// run test
	gotID := h.GetID()
	gotRepoID := h.GetRepoID()
	gotBuildID := h.GetBuildID()
	gotNumber := h.GetNumber()
	gotSourceID := h.GetSourceID()
	gotCreated := h.GetCreated()
	gotHost := h.GetHost()
	gotEvent := h.GetEvent()
	gotBranch := h.GetBranch()
	gotError := h.GetError()
	gotStatus := h.GetStatus()
	gotLink := h.GetLink()

	if gotID != wantID {
		t.Errorf("GetID is %v, want %v", gotID, wantID)
	}
	if gotRepoID != wantRepoID {
		t.Errorf("GetRepoID is %v, want %v", gotRepoID, wantRepoID)
	}
	if gotBuildID != wantBuildID {
		t.Errorf("GetBuildID is %v, want %v", gotBuildID, wantBuildID)
	}
	if gotNumber != wantNumber {
		t.Errorf("GetNumber is %v, want %v", gotNumber, wantNumber)
	}
	if gotSourceID != wantSourceID {
		t.Errorf("GetSourceID is %v, want %v", gotSourceID, wantSourceID)
	}
	if gotCreated != wantCreated {
		t.Errorf("GetCreated is %v, want %v", gotCreated, wantCreated)
	}
	if gotHost != wantHost {
		t.Errorf("GetHost is %v, want %v", gotHost, wantHost)
	}
	if gotEvent != wantEvent {
		t.Errorf("GetEvent is %v, want %v", gotEvent, wantEvent)
	}
	if gotBranch != wantBranch {
		t.Errorf("GetBranch is %v, want %v", gotBranch, wantBranch)
	}
	if gotError != wantError {
		t.Errorf("GetError is %v, want %v", gotError, wantError)
	}
	if gotStatus != wantStatus {
		t.Errorf("GetStatus is %v, want %v", gotStatus, wantStatus)
	}
	if gotLink != wantLink {
		t.Errorf("GetLink is %v, want %v", gotLink, wantLink)
	}
}

func TestLibrary_Hook_Getters_Empty(t *testing.T) {
	// setup types
	h := new(Hook)

	// run test
	gotID := h.GetID()
	gotRepoID := h.GetRepoID()
	gotBuildID := h.GetRepoID()
	gotNumber := h.GetNumber()
	gotSourceID := h.GetSourceID()
	gotCreated := h.GetCreated()
	gotHost := h.GetHost()
	gotEvent := h.GetEvent()
	gotBranch := h.GetBranch()
	gotError := h.GetError()
	gotStatus := h.GetStatus()
	gotLink := h.GetLink()

	if gotID != 0 {
		t.Errorf("GetID is %v, want 0", gotID)
	}
	if gotRepoID != 0 {
		t.Errorf("GetRepoID is %v, want 0", gotRepoID)
	}
	if gotBuildID != 0 {
		t.Errorf("GetBuildID is %v, want 0", gotBuildID)
	}
	if gotNumber != 0 {
		t.Errorf("GetNumber is %v, want 0", gotNumber)
	}
	if gotSourceID != "" {
		t.Errorf("GetSourceID is %v, want \"\"", gotSourceID)
	}
	if gotCreated != 0 {
		t.Errorf("GetCreated is %v, want 0", gotCreated)
	}
	if gotHost != "" {
		t.Errorf("GetHost is %v, want \"\"", gotHost)
	}
	if gotEvent != "" {
		t.Errorf("GetEvent is %v, want \"\"", gotEvent)
	}
	if gotBranch != "" {
		t.Errorf("GetBranch is %v, want \"\"", gotBranch)
	}
	if gotError != "" {
		t.Errorf("GetError is %v, want \"\"", gotError)
	}
	if gotStatus != "" {
		t.Errorf("GetStatus is %v, want \"\"", gotStatus)
	}
	if gotLink != "" {
		t.Errorf("GetStatus is %v, want \"\"", gotLink)
	}
}

func TestLibrary_Hook_Setters(t *testing.T) {
	// setup types
	wantID := int64(1)
	wantRepoID := int64(1)
	wantBuildID := int64(1)
	wantNumber := 1
	wantSourceID := "c8da1302-07d6-11ea-882f-4893bca275b8"
	wantCreated := time.Now().UTC().Unix()
	wantHost := "github.com"
	wantEvent := "push"
	wantBranch := "master"
	wantError := ""
	wantStatus := "success"
	wantLink := "https://github.com/github/octocat/settings/hooks/1"

	h := new(Hook)

	// Run tests
	h.SetID(wantID)
	h.SetRepoID(wantRepoID)
	h.SetBuildID(wantBuildID)
	h.SetNumber(wantNumber)
	h.SetSourceID(wantSourceID)
	h.SetCreated(wantCreated)
	h.SetHost(wantHost)
	h.SetEvent(wantEvent)
	h.SetBranch(wantBranch)
	h.SetError(wantError)
	h.SetStatus(wantStatus)
	h.SetLink(wantLink)

	if h.GetID() != wantID {
		t.Errorf("SetID is %v, want %v", h.GetID(), wantID)
	}
	if h.GetRepoID() != wantRepoID {
		t.Errorf("SetRepoID is %v, want %v", h.GetRepoID(), wantRepoID)
	}
	if h.GetBuildID() != wantBuildID {
		t.Errorf("SetBuildID is %v, want %v", h.GetBuildID(), wantBuildID)
	}
	if h.GetNumber() != wantNumber {
		t.Errorf("SetNumber is %v, want %v", h.GetNumber(), wantNumber)
	}
	if h.GetSourceID() != wantSourceID {
		t.Errorf("SetSourceID is %v, want %v", h.GetSourceID(), wantSourceID)
	}
	if h.GetCreated() != wantCreated {
		t.Errorf("SetCreated is %v, want %v", h.GetCreated(), wantCreated)
	}
	if h.GetHost() != wantHost {
		t.Errorf("SetHost is %v, want %v", h.GetHost(), wantHost)
	}
	if h.GetEvent() != wantEvent {
		t.Errorf("SetEvent is %v, want %v", h.GetEvent(), wantEvent)
	}
	if h.GetBranch() != wantBranch {
		t.Errorf("SetBranch is %v, want %v", h.GetBranch(), wantBranch)
	}
	if h.GetError() != wantError {
		t.Errorf("SetError is %v, want %v", h.GetError(), wantError)
	}
	if h.GetStatus() != wantStatus {
		t.Errorf("SetStatus is %v, want %v", h.GetStatus(), wantStatus)
	}
	if h.GetLink() != wantLink {
		t.Errorf("SetLink is %v, want %v", h.GetLink(), wantLink)
	}
}

func TestLibrary_Hook_Setters_Empty(t *testing.T) {
	// setup types
	var h *Hook

	// Run tests
	h.SetID(0)
	h.SetRepoID(0)
	h.SetBuildID(0)
	h.SetNumber(0)
	h.SetSourceID("")
	h.SetCreated(0)
	h.SetHost("")
	h.SetEvent("")
	h.SetBranch("")
	h.SetError("")
	h.SetStatus("")
	h.SetLink("")

	if h.GetID() != 0 {
		t.Errorf("SetID is %v, want 0", h.GetID())
	}
	if h.GetRepoID() != 0 {
		t.Errorf("SetRepoID is %v, want 0", h.GetRepoID())
	}
	if h.GetBuildID() != 0 {
		t.Errorf("SetBuildID is %v, want 0", h.GetBuildID())
	}
	if h.GetNumber() != 0 {
		t.Errorf("SetNumber is %v, want 0", h.GetNumber())
	}
	if h.GetSourceID() != "" {
		t.Errorf("SetSourceID is %v, want \"\"", h.GetSourceID())
	}
	if h.GetCreated() != 0 {
		t.Errorf("SetCreated is %v, want 0", h.GetCreated())
	}
	if h.GetHost() != "" {
		t.Errorf("SetHost is %v, want \"\"", h.GetHost())
	}
	if h.GetEvent() != "" {
		t.Errorf("SetEvent is %v, want \"\"", h.GetEvent())
	}
	if h.GetBranch() != "" {
		t.Errorf("SetBranch is %v, want \"\"", h.GetBranch())
	}
	if h.GetError() != "" {
		t.Errorf("SetError is %v, want \"\"", h.GetError())
	}
	if h.GetStatus() != "" {
		t.Errorf("SetStatus is %v, want \"\"", h.GetStatus())
	}
	if h.GetLink() != "" {
		t.Errorf("SetLink is %v, want \"\"", h.GetLink())
	}
}

func TestLibrary_Hook_String(t *testing.T) {
	// setup types
	h := new(Hook)
	h.SetID(1)
	h.SetRepoID(1)
	h.SetBuildID(1)
	h.SetNumber(1)
	h.SetSourceID("c8da1302-07d6-11ea-882f-4893bca275b8")
	h.SetCreated(time.Now().UTC().Unix())
	h.SetHost("github.com")
	h.SetEvent("push")
	h.SetBranch("master")
	h.SetError("")
	h.SetStatus("success")
	h.SetLink("https://github.com/github/octocat/settings/hooks/1")

	want := fmt.Sprintf("%+v", *h)

	// run test
	got := h.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}
