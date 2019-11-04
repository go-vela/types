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
	b.SetBranch(wantBranch)
	b.SetRef(wantRef)
	b.SetBaseRef(wantBaseRef)
	b.SetHost(wantHost)
	b.SetRuntime(wantRuntime)
	b.SetDistribution(wantDistribution)

	if *b.ID != wantID {
		t.Errorf("GetID is %v, want %v", *b.ID, wantID)
	}
	if *b.RepoID != wantRepoID {
		t.Errorf("GetRepoID is %v, want %v", *b.RepoID, wantRepoID)
	}
	if *b.Number != wantNumber {
		t.Errorf("GetNumber is %v, want %v", *b.Number, wantNumber)
	}
	if *b.Parent != wantParent {
		t.Errorf("GetParent is %v, want %v", *b.Parent, wantParent)
	}
	if *b.Event != wantEvent {
		t.Errorf("GetEvent is %v, want %v", *b.Event, wantEvent)
	}
	if *b.Status != wantStatus {
		t.Errorf("GetStatus is %v, want %v", *b.Status, wantStatus)
	}
	if *b.Error != wantError {
		t.Errorf("GetError is %v, want %v", *b.Error, wantError)
	}
	if *b.Enqueued != wantEnqueued {
		t.Errorf("GetEnqueued is %v, want %v", *b.Enqueued, wantEnqueued)
	}
	if *b.Created != wantCreated {
		t.Errorf("GetCreated is %v, want %v", *b.Created, wantCreated)
	}
	if *b.Started != wantStarted {
		t.Errorf("GetStarted is %v, want %v", *b.Started, wantStarted)
	}
	if *b.Finished != wantFinished {
		t.Errorf("GetFinished is %v, want %v", *b.Finished, wantFinished)
	}
	if *b.Deploy != wantDeploy {
		t.Errorf("GetDeploy is %v, want %v", *b.Deploy, wantDeploy)
	}
	if *b.Clone != wantClone {
		t.Errorf("GetClone is %v, want %v", *b.Clone, wantClone)
	}
	if *b.Source != wantSource {
		t.Errorf("GetSource is %v, want %v", *b.Source, wantSource)
	}
	if *b.Title != wantTitle {
		t.Errorf("GetTitle is %v, want %v", *b.Title, wantTitle)
	}
	if *b.Message != wantMessage {
		t.Errorf("GetMessage is %v, want %v", *b.Message, wantMessage)
	}
	if *b.Commit != wantCommit {
		t.Errorf("GetCommit is %v, want %v", *b.Commit, wantCommit)
	}
	if *b.Sender != wantSender {
		t.Errorf("GetSender is %v, want %v", *b.Sender, wantSender)
	}
	if *b.Author != wantAuthor {
		t.Errorf("GetAuthor is %v, want %v", *b.Author, wantAuthor)
	}
	if *b.Branch != wantBranch {
		t.Errorf("GetBranch is %v, want %v", *b.Branch, wantBranch)
	}
	if *b.Ref != wantRef {
		t.Errorf("GetRef is %v, want %v", *b.Ref, wantRef)
	}
	if *b.BaseRef != wantBaseRef {
		t.Errorf("GetBaseRef is %v, want %v", *b.BaseRef, wantBaseRef)
	}
	if *b.Host != wantHost {
		t.Errorf("GetHost is %v, want %v", *b.Host, wantHost)
	}
	if *b.Runtime != wantRuntime {
		t.Errorf("GetRuntime is %v, want %v", *b.Runtime, wantRuntime)
	}
	if *b.Distribution != wantDistribution {
		t.Errorf("GetDistribution is %v, want %v", *b.Distribution, wantDistribution)
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
