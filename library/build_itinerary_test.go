// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_BuildItinerary_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		buildItinerary *BuildItinerary
		want           *BuildItinerary
	}{
		{
			buildItinerary: testBuildItinerary(),
			want:           testBuildItinerary(),
		},
		{
			buildItinerary: new(BuildItinerary),
			want:           new(BuildItinerary),
		},
	}

	// run tests
	for _, test := range tests {
		if test.buildItinerary.GetID() != test.want.GetID() {
			t.Errorf("GetID is %v, want %v", test.buildItinerary.GetID(), test.want.GetID())
		}

		if test.buildItinerary.GetBuildID() != test.want.GetBuildID() {
			t.Errorf("GetBuildID is %v, want %v", test.buildItinerary.GetBuildID(), test.want.GetBuildID())
		}

		if !reflect.DeepEqual(test.buildItinerary.GetData(), test.want.GetData()) {
			t.Errorf("GetData is %v, want %v", test.buildItinerary.GetData(), test.want.GetData())
		}
	}
}

func TestLibrary_BuildItinerary_Setters(t *testing.T) {
	// setup types
	var bItinerary *BuildItinerary

	// setup tests
	tests := []struct {
		buildItinerary *BuildItinerary
		want           *BuildItinerary
	}{
		{
			buildItinerary: testBuildItinerary(),
			want:           testBuildItinerary(),
		},
		{
			buildItinerary: bItinerary,
			want:           new(BuildItinerary),
		},
	}

	// run tests
	for _, test := range tests {
		test.buildItinerary.SetID(test.want.GetID())
		test.buildItinerary.SetBuildID(test.want.GetBuildID())
		test.buildItinerary.SetData(test.want.GetData())

		if test.buildItinerary.GetID() != test.want.GetID() {
			t.Errorf("SetID is %v, want %v", test.buildItinerary.GetID(), test.want.GetID())
		}

		if test.buildItinerary.GetBuildID() != test.want.GetBuildID() {
			t.Errorf("SetRepoID is %v, want %v", test.buildItinerary.GetBuildID(), test.want.GetBuildID())
		}

		if !reflect.DeepEqual(test.buildItinerary.GetData(), test.want.GetData()) {
			t.Errorf("SetData is %v, want %v", test.buildItinerary.GetData(), test.want.GetData())
		}
	}
}

func TestLibrary_BuildItinerary_String(t *testing.T) {
	// setup types
	bItinerary := testBuildItinerary()

	want := fmt.Sprintf(`{
  BuildID: %d,
  Data: %s,
  ID: %d,
}`,
		bItinerary.GetBuildID(),
		bItinerary.GetData(),
		bItinerary.GetID(),
	)

	// run test
	got := bItinerary.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

// testBuildItinerary is a test helper function to create a Pipeline
// type with all fields set to a fake value.
func testBuildItinerary() *BuildItinerary {
	p := new(BuildItinerary)

	p.SetID(1)
	p.SetBuildID(1)
	p.SetData(testBuildItineraryData())

	return p
}

// testBuildItineraryData is a test helper function to create the
// content for the Data field for the Pipeline type.
func testBuildItineraryData() []byte {
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
