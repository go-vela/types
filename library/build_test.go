// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_Build_Getters(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	b := &Build{
		ID:           &num64,
		RepoID:       &num64,
		Number:       &num,
		Parent:       &num,
		Event:        &str,
		Status:       &str,
		Error:        &str,
		Enqueued:     &num64,
		Created:      &num64,
		Started:      &num64,
		Finished:     &num64,
		Deploy:       &str,
		Clone:        &str,
		Source:       &str,
		Title:        &str,
		Message:      &str,
		Commit:       &str,
		Sender:       &str,
		Author:       &str,
		Email:        &str,
		Link:         &str,
		Branch:       &str,
		Ref:          &str,
		BaseRef:      &str,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
	}
	wantID := num64
	wantRepoID := num64
	wantNumber := num
	wantParent := num
	wantEvent := str
	wantStatus := str
	wantError := str
	wantEnqueued := num64
	wantCreated := num64
	wantStarted := num64
	wantFinished := num64
	wantDeploy := str
	wantClone := str
	wantSource := str
	wantTitle := str
	wantMessage := str
	wantCommit := str
	wantSender := str
	wantAuthor := str
	wantEmail := str
	wantLink := str
	wantBranch := str
	wantRef := str
	wantBaseRef := str
	wantHost := str
	wantRuntime := str
	wantDistribution := str

	// run test
	gotID := b.GetID()
	gotRepoID := b.GetRepoID()
	gotNumber := b.GetNumber()
	gotParent := b.GetParent()
	gotEvent := b.GetEvent()
	gotStatus := b.GetStatus()
	gotError := b.GetError()
	gotEnqueued := b.GetEnqueued()
	gotCreated := b.GetCreated()
	gotStarted := b.GetStarted()
	gotFinished := b.GetFinished()
	gotDeploy := b.GetDeploy()
	gotClone := b.GetClone()
	gotSource := b.GetSource()
	gotTitle := b.GetTitle()
	gotMessage := b.GetMessage()
	gotCommit := b.GetCommit()
	gotSender := b.GetSender()
	gotAuthor := b.GetAuthor()
	gotEmail := b.GetEmail()
	gotLink := b.GetLink()
	gotBranch := b.GetBranch()
	gotRef := b.GetRef()
	gotBaseRef := b.GetBaseRef()
	gotHost := b.GetHost()
	gotRuntime := b.GetRuntime()
	gotDistribution := b.GetDistribution()

	if gotID != wantID {
		t.Errorf("GetID is %v, want %v", gotID, wantID)
	}
	if gotRepoID != wantRepoID {
		t.Errorf("GetRepoID is %v, want %v", gotRepoID, wantRepoID)
	}
	if gotNumber != wantNumber {
		t.Errorf("GetNumber is %v, want %v", gotNumber, wantNumber)
	}
	if gotParent != wantParent {
		t.Errorf("GetParent is %v, want %v", gotParent, wantParent)
	}
	if gotEvent != wantEvent {
		t.Errorf("GetEvent is %v, want %v", gotEvent, wantEvent)
	}
	if gotStatus != wantStatus {
		t.Errorf("GetStatus is %v, want %v", gotStatus, wantStatus)
	}
	if gotError != wantError {
		t.Errorf("GetError is %v, want %v", gotError, wantError)
	}
	if gotEnqueued != wantEnqueued {
		t.Errorf("GetEnqueued is %v, want %v", gotEnqueued, wantEnqueued)
	}
	if gotCreated != wantCreated {
		t.Errorf("GetCreated is %v, want %v", gotCreated, wantCreated)
	}
	if gotStarted != wantStarted {
		t.Errorf("GetStarted is %v, want %v", gotStarted, wantStarted)
	}
	if gotFinished != wantFinished {
		t.Errorf("GetFinished is %v, want %v", gotFinished, wantFinished)
	}
	if gotDeploy != wantDeploy {
		t.Errorf("GetDeploy is %v, want %v", gotDeploy, wantDeploy)
	}
	if gotClone != wantClone {
		t.Errorf("GetClone is %v, want %v", gotClone, wantClone)
	}
	if gotSource != wantSource {
		t.Errorf("GetSource is %v, want %v", gotSource, wantSource)
	}
	if gotTitle != wantTitle {
		t.Errorf("GetTitle is %v, want %v", gotTitle, wantTitle)
	}
	if gotMessage != wantMessage {
		t.Errorf("GetMessage is %v, want %v", gotMessage, wantMessage)
	}
	if gotCommit != wantCommit {
		t.Errorf("GetCommit is %v, want %v", gotCommit, wantCommit)
	}
	if gotSender != wantSender {
		t.Errorf("GetSender is %v, want %v", gotSender, wantSender)
	}
	if gotAuthor != wantAuthor {
		t.Errorf("GetAuthor is %v, want %v", gotAuthor, wantAuthor)
	}
	if gotEmail != wantEmail {
		t.Errorf("GetEmail is %v, want %v", gotEmail, wantEmail)
	}
	if gotLink != wantLink {
		t.Errorf("GetLink is %v, want %v", gotLink, wantLink)
	}
	if gotBranch != wantBranch {
		t.Errorf("GetBranch is %v, want %v", gotBranch, wantBranch)
	}
	if gotRef != wantRef {
		t.Errorf("GetRef is %v, want %v", gotRef, wantRef)
	}
	if gotBaseRef != wantBaseRef {
		t.Errorf("GetBaseRef is %v, want %v", gotBaseRef, wantBaseRef)
	}
	if gotHost != wantHost {
		t.Errorf("GetHost is %v, want %v", gotHost, wantHost)
	}
	if gotRuntime != wantRuntime {
		t.Errorf("GetRuntime is %v, want %v", gotRuntime, wantRuntime)
	}
	if gotDistribution != wantDistribution {
		t.Errorf("GetDistribution is %v, want %v", gotDistribution, wantDistribution)
	}
}

func TestLibrary_Build_Getters_Empty(t *testing.T) {
	// setup types
	b := &Build{}

	// run test
	gotID := b.GetID()
	gotRepoID := b.GetRepoID()
	gotNumber := b.GetNumber()
	gotParent := b.GetParent()
	gotEvent := b.GetEvent()
	gotStatus := b.GetStatus()
	gotError := b.GetError()
	gotEnqueued := b.GetEnqueued()
	gotCreated := b.GetCreated()
	gotStarted := b.GetStarted()
	gotFinished := b.GetFinished()
	gotDeploy := b.GetDeploy()
	gotClone := b.GetClone()
	gotSource := b.GetSource()
	gotTitle := b.GetTitle()
	gotMessage := b.GetMessage()
	gotCommit := b.GetCommit()
	gotSender := b.GetSender()
	gotAuthor := b.GetAuthor()
	gotEmail := b.GetEmail()
	gotLink := b.GetLink()
	gotBranch := b.GetBranch()
	gotRef := b.GetRef()
	gotBaseRef := b.GetBaseRef()
	gotHost := b.GetHost()
	gotRuntime := b.GetRuntime()
	gotDistribution := b.GetDistribution()

	if gotID != 0 {
		t.Errorf("GetID is %v, want 0", gotID)
	}
	if gotRepoID != 0 {
		t.Errorf("GetRepoID is %v, want 0", gotRepoID)
	}
	if gotNumber != 0 {
		t.Errorf("GetNumber is %v, want 0", gotNumber)
	}
	if gotParent != 0 {
		t.Errorf("GetParent is %v, want 0", gotParent)
	}
	if gotEvent != "" {
		t.Errorf("GetEvent is %v, want \"\"", gotEvent)
	}
	if gotStatus != "" {
		t.Errorf("GetStatus is %v, want \"\"", gotStatus)
	}
	if gotError != "" {
		t.Errorf("GetError is %v, want \"\"", gotError)
	}
	if gotEnqueued != 0 {
		t.Errorf("GetEnqueued is %v, want 0", gotEnqueued)
	}
	if gotCreated != 0 {
		t.Errorf("GetCreated is %v, want 0", gotCreated)
	}
	if gotStarted != 0 {
		t.Errorf("GetStarted is %v, want 0", gotStarted)
	}
	if gotFinished != 0 {
		t.Errorf("GetFinished is %v, want 0", gotFinished)
	}
	if gotDeploy != "" {
		t.Errorf("GetDeploy is %v, want \"\"", gotDeploy)
	}
	if gotClone != "" {
		t.Errorf("GetClone is %v, want \"\"", gotClone)
	}
	if gotSource != "" {
		t.Errorf("GetSource is %v, want \"\"", gotSource)
	}
	if gotTitle != "" {
		t.Errorf("GetTitle is %v, want \"\"", gotTitle)
	}
	if gotMessage != "" {
		t.Errorf("GetMessage is %v, want \"\"", gotMessage)
	}
	if gotCommit != "" {
		t.Errorf("GetCommit is %v, want \"\"", gotCommit)
	}
	if gotSender != "" {
		t.Errorf("GetSender is %v, want \"\"", gotSender)
	}
	if gotAuthor != "" {
		t.Errorf("GetAuthor is %v, want \"\"", gotAuthor)
	}
	if gotEmail != "" {
		t.Errorf("GetEmail is %v, want \"\"", gotEmail)
	}
	if gotLink != "" {
		t.Errorf("GetLink is %v, want \"\"", gotLink)
	}
	if gotBranch != "" {
		t.Errorf("GetBranch is %v, want \"\"", gotBranch)
	}
	if gotRef != "" {
		t.Errorf("GetRef is %v, want \"\"", gotRef)
	}
	if gotBaseRef != "" {
		t.Errorf("GetBaseRef is %v, want \"\"", gotBaseRef)
	}
	if gotHost != "" {
		t.Errorf("GetBaseRef is %v, want \"\"", gotHost)
	}
	if gotRuntime != "" {
		t.Errorf("GetBaseRef is %v, want \"\"", gotRuntime)
	}
	if gotDistribution != "" {
		t.Errorf("GetBaseRef is %v, want \"\"", gotDistribution)
	}
}

func TestLibrary_Build_Setters(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	b := &Build{}

	wantID := num64
	wantRepoID := num64
	wantNumber := num
	wantParent := num
	wantEvent := str
	wantStatus := str
	wantError := str
	wantEnqueued := num64
	wantCreated := num64
	wantStarted := num64
	wantFinished := num64
	wantDeploy := str
	wantClone := str
	wantSource := str
	wantTitle := str
	wantMessage := str
	wantCommit := str
	wantSender := str
	wantAuthor := str
	wantEmail := str
	wantLink := str
	wantBranch := str
	wantRef := str
	wantBaseRef := str
	wantHost := str
	wantRuntime := str
	wantDistribution := str

	// Run tests
	b.SetID(wantID)
	b.SetRepoID(wantRepoID)
	b.SetNumber(wantNumber)
	b.SetParent(wantParent)
	b.SetEvent(wantEvent)
	b.SetStatus(wantStatus)
	b.SetError(wantError)
	b.SetEnqueued(wantEnqueued)
	b.SetCreated(wantCreated)
	b.SetStarted(wantStarted)
	b.SetFinished(wantFinished)
	b.SetDeploy(wantDeploy)
	b.SetClone(wantClone)
	b.SetSource(wantSource)
	b.SetTitle(wantTitle)
	b.SetMessage(wantMessage)
	b.SetCommit(wantCommit)
	b.SetSender(wantSender)
	b.SetAuthor(wantAuthor)
	b.SetEmail(wantEmail)
	b.SetLink(wantLink)
	b.SetBranch(wantBranch)
	b.SetRef(wantRef)
	b.SetBaseRef(wantBaseRef)
	b.SetHost(wantHost)
	b.SetRuntime(wantRuntime)
	b.SetDistribution(wantDistribution)

	if b.GetID() != wantID {
		t.Errorf("SetID is %v, want %v", b.GetID(), wantID)
	}
	if b.GetRepoID() != wantRepoID {
		t.Errorf("SetRepoID is %v, want %v", b.GetRepoID(), wantRepoID)
	}
	if b.GetNumber() != wantNumber {
		t.Errorf("SetNumber is %v, want %v", b.GetNumber(), wantNumber)
	}
	if b.GetParent() != wantParent {
		t.Errorf("SetParent is %v, want %v", b.GetParent(), wantParent)
	}
	if b.GetEvent() != wantEvent {
		t.Errorf("SetEvent is %v, want %v", b.GetEvent(), wantEvent)
	}
	if b.GetStatus() != wantStatus {
		t.Errorf("SetStatus is %v, want %v", b.GetStatus(), wantStatus)
	}
	if b.GetError() != wantError {
		t.Errorf("SetError is %v, want %v", b.GetError(), wantError)
	}
	if b.GetEnqueued() != wantEnqueued {
		t.Errorf("SetEnqueued is %v, want %v", b.GetEnqueued(), wantEnqueued)
	}
	if b.GetCreated() != wantCreated {
		t.Errorf("SetCreated is %v, want %v", b.GetCreated(), wantCreated)
	}
	if b.GetStarted() != wantStarted {
		t.Errorf("SetStarted is %v, want %v", b.GetStarted(), wantStarted)
	}
	if b.GetFinished() != wantFinished {
		t.Errorf("SetFinished is %v, want %v", b.GetFinished(), wantFinished)
	}
	if b.GetDeploy() != wantDeploy {
		t.Errorf("SetDeploy is %v, want %v", b.GetDeploy(), wantDeploy)
	}
	if b.GetClone() != wantClone {
		t.Errorf("SetClone is %v, want %v", b.GetClone(), wantClone)
	}
	if b.GetSource() != wantSource {
		t.Errorf("SetSource is %v, want %v", b.GetSource(), wantSource)
	}
	if b.GetTitle() != wantTitle {
		t.Errorf("SetTitle is %v, want %v", b.GetTitle(), wantTitle)
	}
	if b.GetMessage() != wantMessage {
		t.Errorf("SetMessage is %v, want %v", b.GetMessage(), wantMessage)
	}
	if b.GetCommit() != wantCommit {
		t.Errorf("SetCommit is %v, want %v", b.GetCommit(), wantCommit)
	}
	if b.GetSender() != wantSender {
		t.Errorf("SetSender is %v, want %v", b.GetSender(), wantSender)
	}
	if b.GetAuthor() != wantAuthor {
		t.Errorf("SetAuthor is %v, want %v", b.GetAuthor(), wantAuthor)
	}
	if b.GetEmail() != wantEmail {
		t.Errorf("SetEmail is %v, want %v", b.GetEmail(), wantEmail)
	}
	if b.GetLink() != wantLink {
		t.Errorf("SetLink is %v, want %v", b.GetLink(), wantLink)
	}
	if b.GetBranch() != wantBranch {
		t.Errorf("SetBranch is %v, want %v", b.GetBranch(), wantBranch)
	}
	if b.GetRef() != wantRef {
		t.Errorf("SetRef is %v, want %v", b.GetRef(), wantRef)
	}
	if b.GetBaseRef() != wantBaseRef {
		t.Errorf("SetBaseRef is %v, want %v", b.GetBaseRef(), wantBaseRef)
	}
	if b.GetHost() != wantHost {
		t.Errorf("SetHost is %v, want %v", b.GetHost(), wantHost)
	}
	if b.GetRuntime() != wantRuntime {
		t.Errorf("SetRuntime is %v, want %v", b.GetRuntime(), wantRuntime)
	}
	if b.GetDistribution() != wantDistribution {
		t.Errorf("SetDistribution is %v, want %v", b.GetDistribution(), wantDistribution)
	}
}

func TestLibrary_Build_Setters_Empty(t *testing.T) {
	// setup types
	b := &Build{}
	b = nil

	// Run tests
	b.SetID(0)
	b.SetRepoID(0)
	b.SetNumber(0)
	b.SetParent(0)
	b.SetEvent("")
	b.SetStatus("")
	b.SetError("")
	b.SetEnqueued(0)
	b.SetCreated(0)
	b.SetStarted(0)
	b.SetFinished(0)
	b.SetDeploy("")
	b.SetClone("")
	b.SetSource("")
	b.SetTitle("")
	b.SetMessage("")
	b.SetCommit("")
	b.SetSender("")
	b.SetAuthor("")
	b.SetEmail("")
	b.SetLink("")
	b.SetBranch("")
	b.SetRef("")
	b.SetBaseRef("")
	b.SetHost("")
	b.SetRuntime("")
	b.SetDistribution("")

	if b.GetID() != 0 {
		t.Errorf("SetID is %v, want 0", b.GetID())
	}
	if b.GetRepoID() != 0 {
		t.Errorf("SetRepoID is %v, want 0", b.GetRepoID())
	}
	if b.GetNumber() != 0 {
		t.Errorf("SetNumber is %v, want 0", b.GetNumber())
	}
	if b.GetParent() != 0 {
		t.Errorf("SetParent is %v, want 0", b.GetParent())
	}
	if b.GetEvent() != "" {
		t.Errorf("SetEvent is %v, want \"\"", b.GetEvent())
	}
	if b.GetStatus() != "" {
		t.Errorf("SetStatus is %v, want \"\"", b.GetStatus())
	}
	if b.GetError() != "" {
		t.Errorf("SetError is %v, want \"\"", b.GetError())
	}
	if b.GetEnqueued() != 0 {
		t.Errorf("SetEnqueued is %v, want 0", b.GetEnqueued())
	}
	if b.GetCreated() != 0 {
		t.Errorf("SetCreated is %v, want 0", b.GetCreated())
	}
	if b.GetStarted() != 0 {
		t.Errorf("SetStarted is %v, want 0", b.GetStarted())
	}
	if b.GetFinished() != 0 {
		t.Errorf("SetFinished is %v, want 0", b.GetFinished())
	}
	if b.GetDeploy() != "" {
		t.Errorf("SetDeploy is %v, want \"\"", b.GetDeploy())
	}
	if b.GetClone() != "" {
		t.Errorf("SetClone is %v, want \"\"", b.GetClone())
	}
	if b.GetSource() != "" {
		t.Errorf("SetSource is %v, want \"\"", b.GetSource())
	}
	if b.GetTitle() != "" {
		t.Errorf("SetTitle is %v, want \"\"", b.GetTitle())
	}
	if b.GetMessage() != "" {
		t.Errorf("SetMessage is %v, want \"\"", b.GetMessage())
	}
	if b.GetCommit() != "" {
		t.Errorf("SetCommit is %v, want \"\"", b.GetCommit())
	}
	if b.GetSender() != "" {
		t.Errorf("SetSender is %v, want \"\"", b.GetSender())
	}
	if b.GetAuthor() != "" {
		t.Errorf("SetAuthor is %v, want \"\"", b.GetAuthor())
	}
	if b.GetEmail() != "" {
		t.Errorf("SetEmail is %v, want \"\"", b.GetEmail())
	}
	if b.GetLink() != "" {
		t.Errorf("SetLink is %v, want \"\"", b.GetLink())
	}
	if b.GetBranch() != "" {
		t.Errorf("SetBranch is %v, want \"\"", b.GetBranch())
	}
	if b.GetRef() != "" {
		t.Errorf("SetRef is %v, want \"\"", b.GetRef())
	}
	if b.GetBaseRef() != "" {
		t.Errorf("SetBaseRef is %v, want \"\"", b.GetBaseRef())
	}
	if b.GetHost() != "" {
		t.Errorf("SetHost is %v, want \"\"", b.GetHost())
	}
	if b.GetRuntime() != "" {
		t.Errorf("SetRuntime is %v, want \"\"", b.GetRuntime())
	}
	if b.GetDistribution() != "" {
		t.Errorf("SetDistribution is %v, want \"\"", b.GetDistribution())
	}
}

func TestLibrary_Build_String(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	b := &Build{
		ID:           &num64,
		RepoID:       &num64,
		Number:       &num,
		Parent:       &num,
		Event:        &str,
		Status:       &str,
		Error:        &str,
		Enqueued:     &num64,
		Created:      &num64,
		Started:      &num64,
		Finished:     &num64,
		Deploy:       &str,
		Clone:        &str,
		Source:       &str,
		Title:        &str,
		Message:      &str,
		Commit:       &str,
		Sender:       &str,
		Author:       &str,
		Email:        &str,
		Link:         &str,
		Branch:       &str,
		Ref:          &str,
		BaseRef:      &str,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
	}
	want := fmt.Sprintf("%+v", *b)

	// run test
	got := b.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}
