// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"context"
	"testing"
)

func TestLibrary_BuildFromContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)
	want := &Build{ID: &id}

	// setup context
	ctx = context.WithValue(ctx, buildKey, want)

	// run test
	got := BuildFromContext(ctx)

	if got != want {
		t.Errorf("BuildFromContext is %v, want %v", got, want)
	}
}

func TestLibrary_BuildFromContext_Empty(t *testing.T) {
	// setup types
	ctx := context.Background()

	// run test
	got := BuildFromContext(ctx)

	if got != nil {
		t.Errorf("BuildFromContext is %v, want nil", got)
	}
}

func TestLibrary_BuildFromContext_WrongType(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)

	// setup context
	ctx = context.WithValue(ctx, buildKey, id)

	// run test
	got := BuildFromContext(ctx)

	if got != nil {
		t.Errorf("BuildFromContext is %v, want nil", got)
	}
}

func TestLibrary_BuildWithContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)
	want := &Build{ID: &id}

	// setup context
	ctx = BuildWithContext(ctx, want)

	// run test
	got := ctx.Value(buildKey)

	if got != want {
		t.Errorf("BuildWithContext is %v, want %v", got, want)
	}
}

func TestLibrary_LogFromContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)
	want := &Log{ID: &id}

	// setup context
	ctx = context.WithValue(ctx, logKey, want)

	// run test
	got := LogFromContext(ctx)

	if got != want {
		t.Errorf("LogFromContext is %v, want %v", got, want)
	}
}

func TestLibrary_LogFromContext_Empty(t *testing.T) {
	// setup types
	ctx := context.Background()

	// run test
	got := LogFromContext(ctx)

	if got != nil {
		t.Errorf("LogFromContext is %v, want nil", got)
	}
}

func TestLibrary_LogFromContext_WrongType(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)

	// setup context
	ctx = context.WithValue(ctx, logKey, id)

	// run test
	got := LogFromContext(ctx)

	if got != nil {
		t.Errorf("LogFromContext is %v, want nil", got)
	}
}

func TestLibrary_LogWithContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)
	want := &Log{ID: &id}

	// setup context
	ctx = LogWithContext(ctx, want)

	// run test
	got := ctx.Value(logKey)

	if got != want {
		t.Errorf("LogWithContext is %v, want %v", got, want)
	}
}

func TestLibrary_RepoFromContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)
	want := &Repo{ID: &id}

	// setup context
	ctx = context.WithValue(ctx, repoKey, want)

	// run test
	got := RepoFromContext(ctx)

	if got != want {
		t.Errorf("RepoFromContext is %v, want %v", got, want)
	}
}

func TestLibrary_RepoFromContext_Empty(t *testing.T) {
	// setup types
	ctx := context.Background()

	// run test
	got := RepoFromContext(ctx)

	if got != nil {
		t.Errorf("RepoFromContext is %v, want nil", got)
	}
}

func TestLibrary_RepoFromContext_WrongType(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)

	// setup context
	ctx = context.WithValue(ctx, repoKey, id)

	// run test
	got := RepoFromContext(ctx)

	if got != nil {
		t.Errorf("RepoFromContext is %v, want nil", got)
	}
}

func TestLibrary_RepoWithContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)
	want := &Repo{ID: &id}

	// setup context
	ctx = RepoWithContext(ctx, want)

	// run test
	got := ctx.Value(repoKey)

	if got != want {
		t.Errorf("RepoWithContext is %v, want %v", got, want)
	}
}

func TestLibrary_SecretFromContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)
	want := &Secret{ID: &id}

	// setup context
	ctx = context.WithValue(ctx, secretKey, want)

	// run test
	got := SecretFromContext(ctx)

	if got != want {
		t.Errorf("SecretFromContext is %v, want %v", got, want)
	}
}

func TestLibrary_SecretFromContext_Empty(t *testing.T) {
	// setup types
	ctx := context.Background()

	// run test
	got := SecretFromContext(ctx)

	if got != nil {
		t.Errorf("SecretFromContext is %v, want nil", got)
	}
}

func TestLibrary_SecretFromContext_WrongType(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)

	// setup context
	ctx = context.WithValue(ctx, secretKey, id)

	// run test
	got := SecretFromContext(ctx)

	if got != nil {
		t.Errorf("SecretFromContext is %v, want nil", got)
	}
}

func TestLibrary_SecretWithContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)
	want := &Secret{ID: &id}

	// setup context
	ctx = SecretWithContext(ctx, want)

	// run test
	got := ctx.Value(secretKey)

	if got != want {
		t.Errorf("SecretWithContext is %v, want %v", got, want)
	}
}

func TestLibrary_StepFromContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)
	want := &Step{ID: &id}

	// setup context
	ctx = context.WithValue(ctx, stepKey, want)

	// run test
	got := StepFromContext(ctx)

	if got != want {
		t.Errorf("StepFromContext is %v, want %v", got, want)
	}
}

func TestLibrary_StepFromContext_Empty(t *testing.T) {
	// setup types
	ctx := context.Background()

	// run test
	got := StepFromContext(ctx)

	if got != nil {
		t.Errorf("StepFromContext is %v, want nil", got)
	}
}

func TestLibrary_StepFromContext_WrongType(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)

	// setup context
	ctx = context.WithValue(ctx, stepKey, id)

	// run test
	got := StepFromContext(ctx)

	if got != nil {
		t.Errorf("StepFromContext is %v, want nil", got)
	}
}

func TestLibrary_StepWithContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)
	want := &Step{ID: &id}

	// setup context
	ctx = StepWithContext(ctx, want)

	// run test
	got := ctx.Value(stepKey)

	if got != want {
		t.Errorf("StepWithContext is %v, want %v", got, want)
	}
}

func TestLibrary_UserFromContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)
	want := &User{ID: &id}

	// setup context
	ctx = context.WithValue(ctx, userKey, want)

	// run test
	got := UserFromContext(ctx)

	if got != want {
		t.Errorf("UserFromContext is %v, want %v", got, want)
	}
}

func TestLibrary_UserFromContext_Empty(t *testing.T) {
	// setup types
	ctx := context.Background()

	// run test
	got := UserFromContext(ctx)

	if got != nil {
		t.Errorf("UserFromContext is %v, want nil", got)
	}
}

func TestLibrary_UserFromContext_WrongType(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)

	// setup context
	ctx = context.WithValue(ctx, userKey, id)

	// run test
	got := UserFromContext(ctx)

	if got != nil {
		t.Errorf("UserFromContext is %v, want nil", got)
	}
}

func TestLibrary_UserWithContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)
	want := &User{ID: &id}

	// setup context
	ctx = UserWithContext(ctx, want)

	// run test
	got := ctx.Value(userKey)

	if got != want {
		t.Errorf("UserWithContext is %v, want %v", got, want)
	}
}
