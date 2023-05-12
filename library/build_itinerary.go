// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
)

// BuildItinerary is the library representation of a BuildItinerary.
//
// swagger:model BuildItinerary
type BuildItinerary struct {
	ID      *int64 `json:"id,omitempty"`
	BuildID *int64 `json:"build_id,omitempty"`
	// swagger:strfmt base64
	Data *[]byte `json:"data,omitempty"`
}

// GetID returns the ID field.
//
// When the provided BuildItinerary type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *BuildItinerary) GetID() int64 {
	// return zero value if BuildItinerary type or ID field is nil
	if b == nil || b.ID == nil {
		return 0
	}

	return *b.ID
}

// GetBuildID returns the BuildID field.
//
// When the provided BuildItinerary type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *BuildItinerary) GetBuildID() int64 {
	// return zero value if BuildItinerary type or BuildID field is nil
	if b == nil || b.BuildID == nil {
		return 0
	}

	return *b.BuildID
}

// GetData returns the Data field.
//
// When the provided BuildItinerary type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (b *BuildItinerary) GetData() []byte {
	// return zero value if BuildItinerary type or Data field is nil
	if b == nil || b.Data == nil {
		return []byte{}
	}

	return *b.Data
}

// SetID sets the ID field.
//
// When the provided BuildItinerary type is nil, it
// will set nothing and immediately return.
func (b *BuildItinerary) SetID(v int64) {
	// return if BuildItinerary type is nil
	if b == nil {
		return
	}

	b.ID = &v
}

// SetBuildID sets the BuildID field.
//
// When the provided BuildItinerary type is nil, it
// will set nothing and immediately return.
func (b *BuildItinerary) SetBuildID(v int64) {
	// return if BuildItinerary type is nil
	if b == nil {
		return
	}

	b.BuildID = &v
}

// SetData sets the Data field.
//
// When the provided BuildItinerary type is nil, it
// will set nothing and immediately return.
func (b *BuildItinerary) SetData(v []byte) {
	// return if Log type is nil
	if b == nil {
		return
	}

	b.Data = &v
}

// String implements the Stringer interface for the BuildItinerary type.
func (b *BuildItinerary) String() string {
	return fmt.Sprintf(`{
  BuildID: %d,
  Data: %s,
  ID: %d,
}`,
		b.GetBuildID(),
		b.GetData(),
		b.GetID(),
	)
}
