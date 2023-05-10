// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
)

// Compiled is the library representation of a Compiled.
//
// swagger:model Compiled
type Compiled struct {
	ID      *int64 `json:"id,omitempty"`
	BuildID *int64 `json:"build_id,omitempty"`
	// swagger:strfmt base64
	Data *[]byte `json:"data,omitempty"`
}

// GetID returns the ID field.
//
// When the provided Compiled type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (c *Compiled) GetID() int64 {
	// return zero value if Compiled type or ID field is nil
	if c == nil || c.ID == nil {
		return 0
	}

	return *c.ID
}

// GetBuildID returns the BuildID field.
//
// When the provided Compiled type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (c *Compiled) GetBuildID() int64 {
	// return zero value if Compiled type or BuildID field is nil
	if c == nil || c.BuildID == nil {
		return 0
	}

	return *c.BuildID
}

// GetData returns the Data field.
//
// When the provided Compiled type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (c *Compiled) GetData() []byte {
	// return zero value if Compiled type or Data field is nil
	if c == nil || c.Data == nil {
		return []byte{}
	}

	return *c.Data
}

// SetID sets the ID field.
//
// When the provided Compiled type is nil, it
// will set nothing and immediately return.
func (c *Compiled) SetID(v int64) {
	// return if Compiled type is nil
	if c == nil {
		return
	}

	c.ID = &v
}

// SetBuildID sets the BuildID field.
//
// When the provided Compiled type is nil, it
// will set nothing and immediately return.
func (c *Compiled) SetBuildID(v int64) {
	// return if Compiled type is nil
	if c == nil {
		return
	}

	c.BuildID = &v
}

// SetData sets the Data field.
//
// When the provided Compiled type is nil, it
// will set nothing and immediately return.
func (c *Compiled) SetData(v []byte) {
	// return if Log type is nil
	if c == nil {
		return
	}

	c.Data = &v
}

// String implements the Stringer interface for the Compiled type.
func (c *Compiled) String() string {
	return fmt.Sprintf(`{
  ID: %d,
  Data: %s,
  BuildID: %d,
}`,
		c.GetID(),
		c.GetData(),
		c.GetBuildID(),
	)
}
