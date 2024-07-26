// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore false positive with template.go
package library

import "fmt"

// BuildQueue is the library representation of the builds in the queue.
//
// Deprecated: Use QueueBuild from github.com/go-vela/server/api/types instead.
//
// swagger:model BuildQueue
type BuildQueue struct {
	Status   *string `json:"status,omitempty"`
	Number   *int32  `json:"number,omitempty"`
	Created  *int64  `json:"created,omitempty"`
	FullName *string `json:"full_name,omitempty"`
}

// GetStatus returns the Status field.
//
// When the provided BuildQueue type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *BuildQueue) GetStatus() string {
	// return zero value if BuildQueue type or Status field is nil
	if b == nil || b.Status == nil {
		return ""
	}

	return *b.Status
}

// GetNumber returns the Number field.
//
// When the provided BuildQueue type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *BuildQueue) GetNumber() int32 {
	// return zero value if BuildQueue type or Number field is nil
	if b == nil || b.Number == nil {
		return 0
	}

	return *b.Number
}

// GetCreated returns the Created field.
//
// When the provided BuildQueue type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *BuildQueue) GetCreated() int64 {
	// return zero value if BuildQueue type or Created field is nil
	if b == nil || b.Created == nil {
		return 0
	}

	return *b.Created
}

// GetFullName returns the FullName field.
//
// When the provided BuildQueue type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *BuildQueue) GetFullName() string {
	// return zero value if BuildQueue type or FullName field is nil
	if b == nil || b.FullName == nil {
		return ""
	}

	return *b.FullName
}

// SetStatus sets the Status field.
//
// When the provided BuildQueue type is nil, it
// will set nothing and immediately return.
func (b *BuildQueue) SetStatus(v string) {
	// return if BuildQueue type is nil
	if b == nil {
		return
	}

	b.Status = &v
}

// SetNumber sets the Number field.
//
// When the provided BuildQueue type is nil, it
// will set nothing and immediately return.
func (b *BuildQueue) SetNumber(v int32) {
	// return if BuildQueue type is nil
	if b == nil {
		return
	}

	b.Number = &v
}

// SetCreated sets the Created field.
//
// When the provided BuildQueue type is nil, it
// will set nothing and immediately return.
func (b *BuildQueue) SetCreated(v int64) {
	// return if BuildQueue type is nil
	if b == nil {
		return
	}

	b.Created = &v
}

// SetFullName sets the FullName field.
//
// When the provided BuildQueue type is nil, it
// will set nothing and immediately return.
func (b *BuildQueue) SetFullName(v string) {
	// return if BuildQueue type is nil
	if b == nil {
		return
	}

	b.FullName = &v
}

// String implements the Stringer interface for the BuildQueue type.
func (b *BuildQueue) String() string {
	return fmt.Sprintf(`{
  Created: %d,
  FullName: %s,
  Number: %d,
  Status: %s,
}`,
		b.GetCreated(),
		b.GetFullName(),
		b.GetNumber(),
		b.GetStatus(),
	)
}
