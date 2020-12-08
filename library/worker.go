// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
)

// Worker is the library representation of a worker.
//
// swagger:model Worker
type Worker struct {
	ID            *int64    `json:"id,omitempty"`
	Hostname      *string   `json:"hostname,omitempty"`
	Address       *string   `json:"address,omitempty"`
	Routes        *[]string `json:"routes,omitempty"`
	Active        *bool     `json:"active,omitempty"`
	LastCheckedIn *int64    `json:"last_checked_in,omitempty"`
	BuildLimit    *int      `json:"build_limit,omitempty"`
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

// GetAddress returns the Address field.
//
// When the provided Worker type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (w *Worker) GetAddress() string {
	// return zero value if Worker type or Address field is nil
	if w == nil || w.Address == nil {
		return ""
	}

	return *w.Address
}

// GetRoutes returns the Routes field.
//
// When the provided Worker type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (w *Worker) GetRoutes() []string {
	// return zero value if Worker type or Routes field is nil
	if w == nil || w.Routes == nil {
		return []string{}
	}

	return *w.Routes
}

// GetActive returns the Active field.
//
// When the provided Worker type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (w *Worker) GetActive() bool {
	// return zero value if Worker type or Active field is nil
	if w == nil || w.Active == nil {
		return false
	}

	return *w.Active
}

// GetLastCheckedIn returns the LastCheckedIn field.
//
// When the provided Worker type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (w *Worker) GetLastCheckedIn() int64 {
	// return zero value if Worker type or LastCheckedIn field is nil
	if w == nil || w.LastCheckedIn == nil {
		return 0
	}

	return *w.LastCheckedIn
}

// GetBuildLimit returns the BuildLimit field.
//
// When the provided Worker type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (w *Worker) GetBuildLimit() int {
	// return zero value if Worker type or BuildLimit field is nil
	if w == nil || w.BuildLimit == nil {
		return 0
	}

	return *w.BuildLimit
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

// SetAddress sets the Address field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (w *Worker) SetAddress(v string) {
	// return if Worker type is nil
	if w == nil {
		return
	}

	w.Address = &v
}

// SetRoutes sets the Routes field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (w *Worker) SetRoutes(v []string) {
	// return if Worker type is nil
	if w == nil {
		return
	}

	w.Routes = &v
}

// SetActive sets the Active field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (w *Worker) SetActive(v bool) {
	// return if Worker type is nil
	if w == nil {
		return
	}

	w.Active = &v
}

// SetLastCheckedIn sets the LastCheckedIn field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (w *Worker) SetLastCheckedIn(v int64) {
	// return if Worker type is nil
	if w == nil {
		return
	}

	w.LastCheckedIn = &v
}

// SetBuildLimit sets the LastBuildLimit field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (w *Worker) SetBuildLimit(v int) {
	// return if Worker type is nil
	if w == nil {
		return
	}

	w.BuildLimit = &v
}

// String implements the Stringer interface for the Worker type.
func (w *Worker) String() string {
	return fmt.Sprintf(`{
  ID: %d,
  Hostname: %s,
  Address: %s,
  Routes: %s,
  Active: %t,
  LastCheckedIn: %v,
  BuildLimit: %v,
}`,
		w.GetID(),
		w.GetHostname(),
		w.GetAddress(),
		w.GetRoutes(),
		w.GetActive(),
		w.GetLastCheckedIn(),
		w.GetBuildLimit(),
	)
}
