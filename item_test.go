// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package types

import (
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
)

func TestTypes_ToItem(t *testing.T) {
	// setup types
	booL := false
	num := 1
	num64 := int64(num)
	str := "foo"
	b := &library.Build{
		ID:       &num64,
		RepoID:   &num64,
		Number:   &num,
		Parent:   &num,
		Event:    &str,
		Status:   &str,
		Error:    &str,
		Enqueued: &num64,
		Created:  &num64,
		Started:  &num64,
		Finished: &num64,
		Deploy:   &str,
		Clone:    &str,
		Source:   &str,
		Title:    &str,
		Message:  &str,
		Commit:   &str,
		Sender:   &str,
		Author:   &str,
		Branch:   &str,
		Ref:      &str,
		BaseRef:  &str,
	}
	r := &library.Repo{
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
	u := &library.User{
		ID:     &num64,
		Name:   &str,
		Token:  &str,
		Active: &booL,
		Admin:  &booL,
	}
	want := &Item{
		Build: &library.Build{
			ID:       &num64,
			RepoID:   &num64,
			Number:   &num,
			Parent:   &num,
			Event:    &str,
			Status:   &str,
			Error:    &str,
			Enqueued: &num64,
			Created:  &num64,
			Started:  &num64,
			Finished: &num64,
			Deploy:   &str,
			Clone:    &str,
			Source:   &str,
			Title:    &str,
			Message:  &str,
			Commit:   &str,
			Sender:   &str,
			Author:   &str,
			Branch:   &str,
			Ref:      &str,
			BaseRef:  &str,
		},
		Repo: &library.Repo{
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
		},
		User: &library.User{
			ID:     &num64,
			Name:   &str,
			Token:  &str,
			Active: &booL,
			Admin:  &booL,
		},
		ItemVersion: ItemVersion,
	}

	// run test
	got := ToItem(b, r, u)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToItem is %v, want %v", got, want)
	}
}
