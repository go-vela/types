// SPDX-License-Identifier: Apache-2.0

package database

import (
	"context"
	"testing"
)

func TestDatabase_BuildFromContext(t *testing.T) {
	// setup types
	b := testBuild()

	// setup tests
	tests := []struct {
		ctx  context.Context
		want *Build
	}{
		{
			ctx:  context.WithValue(context.Background(), buildKey, b),
			want: b,
		},
		{
			ctx:  context.Background(),
			want: nil,
		},
		{
			ctx:  context.WithValue(context.Background(), buildKey, "foo"),
			want: nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := BuildFromContext(test.ctx)

		if got != test.want {
			t.Errorf("BuildFromContext is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_BuildWithContext(t *testing.T) {
	// setup types
	want := testBuild()

	// setup context
	ctx := BuildWithContext(context.Background(), want)

	// run test
	got := ctx.Value(buildKey)

	if got != want {
		t.Errorf("BuildWithContext is %v, want %v", got, want)
	}
}

func TestDatabase_LogFromContext(t *testing.T) {
	// setup types
	l := testLog()

	// setup tests
	tests := []struct {
		ctx  context.Context
		want *Log
	}{
		{
			ctx:  context.WithValue(context.Background(), logKey, l),
			want: l,
		},
		{
			ctx:  context.Background(),
			want: nil,
		},
		{
			ctx:  context.WithValue(context.Background(), logKey, "foo"),
			want: nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := LogFromContext(test.ctx)

		if got != test.want {
			t.Errorf("LogFromContext is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_LogWithContext(t *testing.T) {
	// setup types
	want := testLog()

	// setup context
	ctx := LogWithContext(context.Background(), want)

	// run test
	got := ctx.Value(logKey)

	if got != want {
		t.Errorf("LogWithContext is %v, want %v", got, want)
	}
}

func TestDatabase_RepoFromContext(t *testing.T) {
	// setup types
	r := testRepo()

	// setup tests
	tests := []struct {
		ctx  context.Context
		want *Repo
	}{
		{
			ctx:  context.WithValue(context.Background(), repoKey, r),
			want: r,
		},
		{
			ctx:  context.Background(),
			want: nil,
		},
		{
			ctx:  context.WithValue(context.Background(), repoKey, "foo"),
			want: nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := RepoFromContext(test.ctx)

		if got != test.want {
			t.Errorf("RepoFromContext is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_RepoWithContext(t *testing.T) {
	// setup types
	want := testRepo()

	// setup context
	ctx := RepoWithContext(context.Background(), want)

	// run test
	got := ctx.Value(repoKey)

	if got != want {
		t.Errorf("RepoWithContext is %v, want %v", got, want)
	}
}

func TestDatabase_SecretFromContext(t *testing.T) {
	// setup types
	s := testSecret()

	// setup tests
	tests := []struct {
		ctx  context.Context
		want *Secret
	}{
		{
			ctx:  context.WithValue(context.Background(), secretKey, s),
			want: s,
		},
		{
			ctx:  context.Background(),
			want: nil,
		},
		{
			ctx:  context.WithValue(context.Background(), secretKey, "foo"),
			want: nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := SecretFromContext(test.ctx)

		if got != test.want {
			t.Errorf("SecretFromContext is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_SecretWithContext(t *testing.T) {
	// setup types
	want := testSecret()

	// setup context
	ctx := SecretWithContext(context.Background(), want)

	// run test
	got := ctx.Value(secretKey)

	if got != want {
		t.Errorf("SecretWithContext is %v, want %v", got, want)
	}
}

func TestDatabase_StepFromContext(t *testing.T) {
	// setup types
	s := testStep()

	// setup tests
	tests := []struct {
		ctx  context.Context
		want *Step
	}{
		{
			ctx:  context.WithValue(context.Background(), stepKey, s),
			want: s,
		},
		{
			ctx:  context.Background(),
			want: nil,
		},
		{
			ctx:  context.WithValue(context.Background(), stepKey, "foo"),
			want: nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := StepFromContext(test.ctx)

		if got != test.want {
			t.Errorf("StepFromContext is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_StepWithContext(t *testing.T) {
	// setup types
	want := testStep()

	// setup context
	ctx := StepWithContext(context.Background(), want)

	// run test
	got := ctx.Value(stepKey)

	if got != want {
		t.Errorf("StepWithContext is %v, want %v", got, want)
	}
}

func TestDatabase_UserFromContext(t *testing.T) {
	// setup types
	u := testUser()

	// setup tests
	tests := []struct {
		ctx  context.Context
		want *User
	}{
		{
			ctx:  context.WithValue(context.Background(), userKey, u),
			want: u,
		},
		{
			ctx:  context.Background(),
			want: nil,
		},
		{
			ctx:  context.WithValue(context.Background(), userKey, "foo"),
			want: nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := UserFromContext(test.ctx)

		if got != test.want {
			t.Errorf("UserFromContext is %v, want %v", got, test.want)
		}
	}
}

func TestDatabase_UserWithContext(t *testing.T) {
	// setup types
	want := testUser()

	// setup context
	ctx := UserWithContext(context.Background(), want)

	// run test
	got := ctx.Value(userKey)

	if got != want {
		t.Errorf("UserWithContext is %v, want %v", got, want)
	}
}
