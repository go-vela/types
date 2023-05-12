// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_Compiled_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		compiled *Compiled
		want     *Compiled
	}{
		{
			compiled: testCompiled(),
			want:     testCompiled(),
		},
		{
			compiled: new(Compiled),
			want:     new(Compiled),
		},
	}

	// run tests
	for _, test := range tests {
		if test.compiled.GetID() != test.want.GetID() {
			t.Errorf("GetID is %v, want %v", test.compiled.GetID(), test.want.GetID())
		}

		if test.compiled.GetBuildID() != test.want.GetBuildID() {
			t.Errorf("GetBuildID is %v, want %v", test.compiled.GetBuildID(), test.want.GetBuildID())
		}

		if !reflect.DeepEqual(test.compiled.GetData(), test.want.GetData()) {
			t.Errorf("GetData is %v, want %v", test.compiled.GetData(), test.want.GetData())
		}
	}
}

func TestLibrary_Compiled_Setters(t *testing.T) {
	// setup types
	var p *Compiled

	// setup tests
	tests := []struct {
		compiled *Compiled
		want     *Compiled
	}{
		{
			compiled: testCompiled(),
			want:     testCompiled(),
		},
		{
			compiled: p,
			want:     new(Compiled),
		},
	}

	// run tests
	for _, test := range tests {
		test.compiled.SetID(test.want.GetID())
		test.compiled.SetBuildID(test.want.GetBuildID())
		test.compiled.SetData(test.want.GetData())

		if test.compiled.GetID() != test.want.GetID() {
			t.Errorf("SetID is %v, want %v", test.compiled.GetID(), test.want.GetID())
		}

		if test.compiled.GetBuildID() != test.want.GetBuildID() {
			t.Errorf("SetRepoID is %v, want %v", test.compiled.GetBuildID(), test.want.GetBuildID())
		}

		if !reflect.DeepEqual(test.compiled.GetData(), test.want.GetData()) {
			t.Errorf("SetData is %v, want %v", test.compiled.GetData(), test.want.GetData())
		}
	}
}

func TestLibrary_Compiled_String(t *testing.T) {
	// setup types
	c := testCompiled()

	want := fmt.Sprintf(`{
  BuildID: %d,
  Data: %s,
  ID: %d,
}`,
		c.GetBuildID(),
		c.GetData(),
		c.GetID(),
	)

	// run test
	got := c.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

// testCompiled is a test helper function to create a Pipeline
// type with all fields set to a fake value.
func testCompiled() *Compiled {
	p := new(Compiled)

	p.SetID(1)
	p.SetBuildID(1)
	p.SetData(testCompiledData())

	return p
}

// testCompiledData is a test helper function to create the
// content for the Data field for the Pipeline type.
func testCompiledData() []byte {
	return []byte(`
{ 
    "id": "step_name",
    "version": "1",
    "metadata":{
        "clone":true,
        "environment":["steps","services","secrets"]},
    "worker":{},
    "steps":[
        {
            "id":"step_github_octocat_1_init",
            "directory":"/vela/src/github.com/github/octocat",
            "environment": {"BUILD_AUTHOR":"Octocat"}
        }
    ]
}
`)
}
