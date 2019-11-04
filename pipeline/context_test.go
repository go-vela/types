// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"context"
	"testing"
)

func TestPipeline_BuildFromContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	want := &Build{ID: "1"}

	// setup context
	ctx = context.WithValue(ctx, buildKey, want)

	// run test
	got := BuildFromContext(ctx)

	if got != want {
		t.Errorf("BuildFromContext is %v, want %v", got, want)
	}
}

func TestPipeline_BuildFromContext_Empty(t *testing.T) {
	// setup types
	ctx := context.Background()

	// run test
	got := BuildFromContext(ctx)

	if got != nil {
		t.Errorf("BuildFromContext is %v, want nil", got)
	}
}

func TestPipeline_BuildFromContext_WrongType(t *testing.T) {
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

func TestPipeline_BuildWithContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	want := &Build{ID: "1"}

	// setup context
	ctx = BuildWithContext(ctx, want)

	// run test
	got := ctx.Value(buildKey)

	if got != want {
		t.Errorf("BuildWithContext is %v, want %v", got, want)
	}
}

func TestPipeline_SecretFromContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	want := &Secret{Name: "1"}

	// setup context
	ctx = context.WithValue(ctx, secretKey, want)

	// run test
	got := SecretFromContext(ctx)

	if got != want {
		t.Errorf("SecretFromContext is %v, want %v", got, want)
	}
}

func TestPipeline_SecretFromContext_Empty(t *testing.T) {
	// setup types
	ctx := context.Background()

	// run test
	got := SecretFromContext(ctx)

	if got != nil {
		t.Errorf("SecretFromContext is %v, want nil", got)
	}
}

func TestPipeline_SecretFromContext_WrongType(t *testing.T) {
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

func TestPipeline_SecretWithContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	want := &Secret{Name: "1"}

	// setup context
	ctx = SecretWithContext(ctx, want)

	// run test
	got := ctx.Value(secretKey)

	if got != want {
		t.Errorf("SecretWithContext is %v, want %v", got, want)
	}
}

func TestPipeline_StageFromContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	want := &Stage{Name: "foo"}

	// setup context
	ctx = context.WithValue(ctx, stageKey, want)

	// run test
	got := StageFromContext(ctx)

	if got != want {
		t.Errorf("StageFromContext is %v, want %v", got, want)
	}
}

func TestPipeline_StageFromContext_Empty(t *testing.T) {
	// setup types
	ctx := context.Background()

	// run test
	got := StageFromContext(ctx)

	if got != nil {
		t.Errorf("StageFromContext is %v, want nil", got)
	}
}

func TestPipeline_StageFromContext_WrongType(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)

	// setup context
	ctx = context.WithValue(ctx, stageKey, id)

	// run test
	got := StageFromContext(ctx)

	if got != nil {
		t.Errorf("StageFromContext is %v, want nil", got)
	}
}

func TestPipeline_StageWithContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	want := &Stage{Name: "foo"}

	// setup context
	ctx = StageWithContext(ctx, want)

	// run test
	got := ctx.Value(stageKey)

	if got != want {
		t.Errorf("StageWithContext is %v, want %v", got, want)
	}
}

func TestPipeline_ContainerFromContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	want := &Container{ID: "1"}

	// setup context
	ctx = context.WithValue(ctx, containerKey, want)

	// run test
	got := ContainerFromContext(ctx)

	if got != want {
		t.Errorf("ContainerFromContext is %v, want %v", got, want)
	}
}

func TestPipeline_ContainerFromContext_Empty(t *testing.T) {
	// setup types
	ctx := context.Background()

	// run test
	got := ContainerFromContext(ctx)

	if got != nil {
		t.Errorf("ContainerFromContext is %v, want nil", got)
	}
}

func TestPipeline_ContainerFromContext_WrongType(t *testing.T) {
	// setup types
	ctx := context.Background()
	id := int64(1)

	// setup context
	ctx = context.WithValue(ctx, containerKey, id)

	// run test
	got := ContainerFromContext(ctx)

	if got != nil {
		t.Errorf("ContainerFromContext is %v, want nil", got)
	}
}

func TestPipeline_ContainerWithContext(t *testing.T) {
	// setup types
	ctx := context.Background()
	want := &Container{ID: "1"}

	// setup context
	ctx = ContainerWithContext(ctx, want)

	// run test
	got := ctx.Value(containerKey)

	if got != want {
		t.Errorf("ContainerWithContext is %v, want %v", got, want)
	}
}
