// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"context"
)

// contextKey defines the key type for
// storing library types in a context.
type contextKey int

const (
	// buildKey defines the key type for
	// storing a Build type in a context.
	buildKey contextKey = iota

	// logKey defines the key type for
	// storing a Log type in a context.
	logKey

	// repoKey defines the key type for
	// storing a Repo type in a context.
	repoKey

	// secretKey defines the key type for
	// storing a Secret type in a context.
	secretKey

	// stepKey defines the key type for
	// storing a Step type in a context.
	stepKey

	// userKey defines the key type for
	// storing a User type in a context.
	userKey
)

// BuildFromContext retrieves the Build type from the context.
func BuildFromContext(c context.Context) *Build {
	// get build value from context
	v := c.Value(buildKey)
	if v == nil {
		return nil
	}

	// cast build value to expected Build type
	b, ok := v.(*Build)
	if !ok {
		return nil
	}

	return b
}

// BuildWithContext inserts the Build type to the context.
func BuildWithContext(c context.Context, b *Build) context.Context {
	return context.WithValue(c, buildKey, b)
}

// LogFromContext retrieves the Log type from the context.
func LogFromContext(c context.Context) *Log {
	// get log value from context
	v := c.Value(logKey)
	if v == nil {
		return nil
	}

	// cast log value to expected Log type
	l, ok := v.(*Log)
	if !ok {
		return nil
	}

	return l
}

// LogWithContext inserts the Log type to the context.
func LogWithContext(c context.Context, l *Log) context.Context {
	return context.WithValue(c, logKey, l)
}

// RepoFromContext retrieves the Repo type from the context.
func RepoFromContext(c context.Context) *Repo {
	// get repo value from context
	v := c.Value(repoKey)
	if v == nil {
		return nil
	}

	// cast repo value to expected Repo type
	r, ok := v.(*Repo)
	if !ok {
		return nil
	}

	return r
}

// RepoWithContext inserts the library Repo type to the context.
func RepoWithContext(c context.Context, r *Repo) context.Context {
	return context.WithValue(c, repoKey, r)
}

// SecretFromContext retrieves the Secret type from the context.
func SecretFromContext(c context.Context) *Secret {
	// get secret value from context
	v := c.Value(secretKey)
	if v == nil {
		return nil
	}

	// cast secret value to expected Secret type
	s, ok := v.(*Secret)
	if !ok {
		return nil
	}

	return s
}

// SecretWithContext inserts the Secret type to the context.
func SecretWithContext(c context.Context, s *Secret) context.Context {
	return context.WithValue(c, secretKey, s)
}

// StepFromContext retrieves the Step type from the context.
func StepFromContext(c context.Context) *Step {
	// get step value from context
	v := c.Value(stepKey)
	if v == nil {
		return nil
	}

	// cast step value to expected Step type
	s, ok := v.(*Step)
	if !ok {
		return nil
	}

	return s
}

// StepWithContext inserts the Step type to the context.
func StepWithContext(c context.Context, s *Step) context.Context {
	return context.WithValue(c, stepKey, s)
}

// UserFromContext retrieves the User type from the context.
func UserFromContext(c context.Context) *User {
	// get user value from context
	v := c.Value(userKey)
	if v == nil {
		return nil
	}

	// cast user value to expected User type
	u, ok := v.(*User)
	if !ok {
		return nil
	}

	return u
}

// UserWithContext inserts the User type to the context.
func UserWithContext(c context.Context, u *User) context.Context {
	return context.WithValue(c, userKey, u)
}
