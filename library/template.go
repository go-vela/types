// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
)

// Template is the library representation of a template for a pipeline.
//
// swagger:model Template
type Template struct {
	HTMLURL *string `json:"html_url,omitempty"`
	Name    *string `json:"name,omitempty"`
	Source  *string `json:"source,omitempty"`
	Type    *string `json:"type,omitempty"`
}

// GetHTMLURL returns the HTMLURL field.
//
// When the provided Template type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (t *Template) GetHTMLURL() string {
	// return zero value if Template type or HTMLURL field is nil
	if t == nil || t.HTMLURL == nil {
		return ""
	}

	return *t.HTMLURL
}

// GetName returns the Name field.
//
// When the provided Template type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (t *Template) GetName() string {
	// return zero value if Template type or Name field is nil
	if t == nil || t.Name == nil {
		return ""
	}

	return *t.Name
}

// GetSource returns the Source field.
//
// When the provided Template type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (t *Template) GetSource() string {
	// return zero value if Template type or Source field is nil
	if t == nil || t.Source == nil {
		return ""
	}

	return *t.Source
}

// GetType returns the Type field.
//
// When the provided Template type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (t *Template) GetType() string {
	// return zero value if Template type or Type field is nil
	if t == nil || t.Type == nil {
		return ""
	}

	return *t.Type
}

// SetHTMLURL sets the HTMLURL field.
//
// When the provided Template type is nil, it
// will set nothing and immediately return.
func (t *Template) SetHTMLURL(v string) {
	// return if Template type is nil
	if t == nil {
		return
	}

	t.HTMLURL = &v
}

// SetName sets the Name field.
//
// When the provided Template type is nil, it
// will set nothing and immediately return.
func (t *Template) SetName(v string) {
	// return if Template type is nil
	if t == nil {
		return
	}

	t.Name = &v
}

// SetSource sets the Source field.
//
// When the provided Template type is nil, it
// will set nothing and immediately return.
func (t *Template) SetSource(v string) {
	// return if Template type is nil
	if t == nil {
		return
	}

	t.Source = &v
}

// SetType sets the Type field.
//
// When the provided Template type is nil, it
// will set nothing and immediately return.
func (t *Template) SetType(v string) {
	// return if Template type is nil
	if t == nil {
		return
	}

	t.Type = &v
}

// String implements the Stringer interface for the Template type.
func (t *Template) String() string {
	return fmt.Sprintf(`{
  Link: %s,
  Name: %s,
  Source: %s,
  Type: %s,
}`,
		t.GetHTMLURL(),
		t.GetName(),
		t.GetSource(),
		t.GetType(),
	)
}
