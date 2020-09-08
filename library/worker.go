// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"time"
)

// Worker is the library representation of a worker.
//
// swagger:model Worker
type Worker struct {
	ID            *int64     `json:"id,omitempty"`
	Hostname      *string    `json:"hostname,omitempty"`
	Path          *string    `json:"path,omitempty"`
	Online        *bool      `json:"online,omitempty"`
	LastCheckedIn *time.Time `json:"last_checked_in,omitempty"`
}

// GetID returns the ID field.
//
// When the provided Worker type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (w *Worker) GetID() int64 {
	// return zero value if Worker type or ID field is nil
	if w == nil || w.ID == nil {
		return 0
	}

	return *w.ID
}

// GetHostname returns the Hostname field.
//
// When the provided Worker type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (w *Worker) GetHostname() string {
	// return zero value if Worker type or Hostname field is nil
	if w == nil || w.Hostname == nil {
		return ""
	}

	return *w.Hostname
}

// GetURL returns the URL field.
//
// When the provided Worker type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (w *Worker) GetURL() string {
	// return zero value if Worker type or URL field is nil
	if w == nil || w.Path == nil {
		return ""
	}

	return *w.Path
}

// GetOnline returns the Online field.
//
// When the provided Worker type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (w *Worker) GetOnline() bool {
	// return zero value if Worker type or Online field is nil
	if w == nil || w.Online == nil {
		return false
	}

	return *w.Online
}

// GetLastCheckedIn returns the Online field.
//
// When the provided Worker type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (w *Worker) GetLastCheckedIn() time.Time {
	// return zero value if Worker type or LastCheckedIn field is nil
	if w == nil || w.LastCheckedIn == nil {
		return time.Time{} // 0001-01-01 00:00:00 +0000 UTC
	}

	return *w.LastCheckedIn
}

// SetID sets the ID field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (w *Worker) SetID(v int64) {
	// return if Worker type is nil
	if w == nil {
		return
	}

	w.ID = &v
}

// SetHostname sets the Hostname field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (w *Worker) SetHostname(v string) {
	// return if Worker type is nil
	if w == nil {
		return
	}

	w.Hostname = &v
}

// SetURL sets the URL field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (w *Worker) SetURL(v string) {
	// return if Worker type is nil
	if w == nil {
		return
	}

	w.Path = &v
}

// SetOnline sets the Online field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (w *Worker) SetOnline(v bool) {
	// return if Worker type is nil
	if w == nil {
		return
	}

	w.Online = &v
}

// SetLastCheckedIn sets the LastCheckedIn field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (w *Worker) SetLastCheckedIn(v time.Time) {
	// return if Worker type is nil
	if w == nil {
		return
	}

	w.LastCheckedIn = &v
}

// String implements the Stringer interface for the Worker type.
func (w *Worker) String() string {
	return fmt.Sprintf(`{
  ID: %d,
  Hostname: %s,
  URL: %s,
  Online: %t,
  LastCheckedIn: %v,
}`,
		w.GetID(),
		w.GetHostname(),
		w.GetURL(),
		w.GetOnline(),
		w.GetLastCheckedIn(),
	)
}
