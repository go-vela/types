// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_Template_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		template *Template
		want     *Template
	}{
		{
			template: testTemplate(),
			want:     testTemplate(),
		},
		{
			template: new(Template),
			want:     new(Template),
		},
	}

	// run tests
	for _, test := range tests {
		if test.template.GetHTMLURL() != test.want.GetHTMLURL() {
			t.Errorf("GetHTMLURL is %v, want %v", test.template.GetHTMLURL(), test.want.GetHTMLURL())
		}

		if test.template.GetName() != test.want.GetName() {
			t.Errorf("GetName is %v, want %v", test.template.GetName(), test.want.GetName())
		}

		if test.template.GetSource() != test.want.GetSource() {
			t.Errorf("GetSource is %v, want %v", test.template.GetSource(), test.want.GetSource())
		}

		if test.template.GetType() != test.want.GetType() {
			t.Errorf("GetType is %v, want %v", test.template.GetType(), test.want.GetType())
		}
	}
}

func TestLibrary_Template_Setters(t *testing.T) {
	// setup types
	var tmpl *Template

	// setup tests
	tests := []struct {
		template *Template
		want     *Template
	}{
		{
			template: testTemplate(),
			want:     testTemplate(),
		},
		{
			template: tmpl,
			want:     new(Template),
		},
	}

	// run tests
	for _, test := range tests {
		test.template.SetHTMLURL(test.want.GetHTMLURL())
		test.template.SetName(test.want.GetName())
		test.template.SetSource(test.want.GetSource())
		test.template.SetType(test.want.GetType())

		if test.template.GetHTMLURL() != test.want.GetHTMLURL() {
			t.Errorf("SetHTMLURL is %v, want %v", test.template.GetHTMLURL(), test.want.GetHTMLURL())
		}

		if test.template.GetName() != test.want.GetName() {
			t.Errorf("SetName is %v, want %v", test.template.GetName(), test.want.GetName())
		}

		if test.template.GetSource() != test.want.GetSource() {
			t.Errorf("SetSource is %v, want %v", test.template.GetSource(), test.want.GetSource())
		}

		if test.template.GetType() != test.want.GetType() {
			t.Errorf("SetType is %v, want %v", test.template.GetType(), test.want.GetType())
		}
	}
}

func TestLibrary_Template_String(t *testing.T) {
	// setup types
	tmpl := testTemplate()

	want := fmt.Sprintf(`{
  Link: %s,
  Name: %s,
  Source: %s,
  Type: %s,
}`,
		tmpl.GetHTMLURL(),
		tmpl.GetName(),
		tmpl.GetSource(),
		tmpl.GetType(),
	)

	// run test
	got := tmpl.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

// testTemplate is a test helper function to create a Template
// type with all fields set to a fake value.
func testTemplate() *Template {
	u := new(Template)

	u.SetHTMLURL("https://github.com/github/octocat/blob/branch/template.yml")
	u.SetName("template")
	u.SetSource("github.com/github/octocat/template.yml@branch")
	u.SetType("github")

	return u
}
