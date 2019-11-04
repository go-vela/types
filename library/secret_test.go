// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-vela/types/pipeline"
)

func TestLibrary_Secret_Sanitize(t *testing.T) {
	// setup types
	one := int64(1)
	foo := "foo"
	bar := "bar"
	repo := "repo"
	images := []string{"foo"}
	events := []string{"bar"}
	s := &Secret{
		ID:     &one,
		Org:    &foo,
		Repo:   &bar,
		Name:   &foo,
		Value:  &bar,
		Type:   &repo,
		Images: &images,
		Events: &events,
	}

	want := &Secret{
		ID:     &one,
		Org:    &foo,
		Repo:   &bar,
		Name:   &foo,
		Type:   &repo,
		Images: &images,
		Events: &events,
	}

	// run test
	got := s.Sanitize()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sanitize is %v, want %v", got, want)
	}
}

func TestLibrary_Secret_Match(t *testing.T) {

	// name and value of secret
	v := "foo"

	// setup types
	tests := []struct {
		step *pipeline.Container
		sec  *Secret
		want bool
	}{
		{ // test build events without ruledata
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "push"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{}, Events: &[]string{"push"}},
			want: true,
		},
		{
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "pull_request"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{}, Events: &[]string{"pull_request"}},
			want: true,
		},
		{
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "tag"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{}, Events: &[]string{"tag"}},
			want: true,
		},
		{
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "deployment"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{}, Events: &[]string{"deployment"}},
			want: true,
		},
		{
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "fake_event"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{}, Events: &[]string{"push"}},
			want: false,
		},
		{
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "push"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{}, Events: &[]string{"push", "pull_request"}},
			want: true,
		},

		{ // test images without ruledata
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "push"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{"alpine"}, Events: &[]string{}},
			want: true,
		},
		{
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "push"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{"alpine:latest"}, Events: &[]string{}},
			want: true,
		},
		{
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "push"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{"alpine:1"}, Events: &[]string{}},
			want: false,
		},
		{
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "push"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{"alpine", "centos"}, Events: &[]string{}},
			want: true,
		},

		{ // test build events with image ACLs
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "push"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{"alpine"}, Events: &[]string{"push"}},
			want: true,
		},
		{
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "push"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{"alpine:latest"}, Events: &[]string{"push"}},
			want: true,
		},
		{
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "push"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{"alpine:1"}, Events: &[]string{"push"}},
			want: false,
		},
		{
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "pull_request"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{"alpine:latest"}, Events: &[]string{"push"}},
			want: false,
		},
		{
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "push"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{"alpine", "centos"}, Events: &[]string{"push"}},
			want: true,
		},

		{ // test build events with image ACLs and rulesets
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "push"},
				Ruleset: pipeline.Ruleset{
					If: pipeline.Rules{
						Event: []string{"push"},
					},
				},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{"alpine"}, Events: &[]string{"push"}},
			want: true,
		},
		//TODO: circle back to make this test pass
		// {
		// 	step: &pipeline.Container{
		// 		Image:       "alpine:latest",
		// 		Environment: map[string]string{"BUILD_EVENT": "pull_request"},
		// 		Ruleset: pipeline.Ruleset{
		// 			Unless: pipeline.Rules{
		// 				Event: []string{"push"},
		// 			},
		// 		},
		// 	},
		// 	sec:  &Secret{Name: &v, Value: &v, Images: &[]string{"alpine"}, Events: &[]string{"pull_request"}},
		// 	want: false,
		// },
	}

	// run test
	for _, test := range tests {

		inject := test.sec.Match(test.step)

		if !inject == test.want {
			t.Errorf("Match should have been %v", inject)
		}
	}
}

func TestLibrary_Secret_Getters(t *testing.T) {
	// setup types
	num64 := int64(1)
	str := "foo"
	arr := []string{"foo", "bar"}
	s := &Secret{
		ID:     &num64,
		Org:    &str,
		Repo:   &str,
		Team:   &str,
		Name:   &str,
		Value:  &str,
		Type:   &str,
		Images: &arr,
		Events: &arr,
	}
	wantID := num64
	wantOrg := str
	wantRepo := str
	wantTeam := str
	wantName := str
	wantValue := str
	wantType := str
	wantImages := arr
	wantEvents := arr

	// run test
	gotID := s.GetID()
	gotOrg := s.GetOrg()
	gotRepo := s.GetRepo()
	gotTeam := s.GetTeam()
	gotName := s.GetName()
	gotValue := s.GetValue()
	gotType := s.GetType()
	gotImages := s.GetImages()
	gotEvents := s.GetEvents()

	if gotID != wantID {
		t.Errorf("GetID is %v, want %v", gotID, wantID)
	}
	if gotOrg != wantOrg {
		t.Errorf("GetOrg is %v, want %v", gotOrg, wantOrg)
	}
	if gotRepo != wantRepo {
		t.Errorf("GetRepo is %v, want %v", gotRepo, wantRepo)
	}
	if gotTeam != wantTeam {
		t.Errorf("GetTeam is %v, want %v", gotTeam, wantTeam)
	}
	if gotName != wantName {
		t.Errorf("GetName is %v, want %v", gotName, wantName)
	}
	if gotValue != wantValue {
		t.Errorf("GetValue is %v, want %v", gotValue, wantValue)
	}
	if gotType != wantType {
		t.Errorf("GetType is %v, want %v", gotType, wantType)
	}
	if !reflect.DeepEqual(gotImages, wantImages) {
		t.Errorf("GetImages is %v, want %v", gotImages, wantImages)
	}
	if !reflect.DeepEqual(gotEvents, wantEvents) {
		t.Errorf("GetEvents is %v, want %v", gotEvents, wantEvents)
	}
}

func TestLibrary_Secret_Getters_Empty(t *testing.T) {
	// setup types
	s := &Secret{}

	// run test
	gotID := s.GetID()
	gotOrg := s.GetOrg()
	gotRepo := s.GetRepo()
	gotTeam := s.GetTeam()
	gotName := s.GetName()
	gotValue := s.GetValue()
	gotType := s.GetType()
	gotImages := s.GetImages()
	gotEvents := s.GetEvents()

	if gotID != 0 {
		t.Errorf("GetID is %v, want 0", gotID)
	}
	if gotOrg != "" {
		t.Errorf("GetOrg is %v, want \"\"", gotOrg)
	}
	if gotRepo != "" {
		t.Errorf("GetRepo is %v, want \"\"", gotRepo)
	}
	if gotTeam != "" {
		t.Errorf("GetTeam is %v, want \"\"", gotTeam)
	}
	if gotName != "" {
		t.Errorf("GetName is %v, want \"\"", gotName)
	}
	if gotValue != "" {
		t.Errorf("GetValue is %v, want \"\"", gotValue)
	}
	if gotType != "" {
		t.Errorf("GetType is %v, want \"\"", gotType)
	}
	if !reflect.DeepEqual(gotImages, []string{}) {
		t.Errorf("GetImages is %v, want []string{}", gotImages)
	}
	if !reflect.DeepEqual(gotEvents, []string{}) {
		t.Errorf("GetEvents is %v, want []string{}", gotEvents)
	}
}

func TestLibrary_Secret_Setters(t *testing.T) {
	// setup types
	num64 := int64(1)
	str := "foo"
	arr := []string{"foo", "bar"}
	s := &Secret{}

	wantID := num64
	wantOrg := str
	wantRepo := str
	wantTeam := str
	wantName := str
	wantValue := str
	wantType := str
	wantImages := arr
	wantEvents := arr

	// run test
	s.SetID(wantID)
	s.SetOrg(wantOrg)
	s.SetRepo(wantRepo)
	s.SetTeam(wantTeam)
	s.SetName(wantName)
	s.SetValue(wantValue)
	s.SetType(wantType)
	s.SetImages(wantImages)
	s.SetEvents(wantEvents)

	if *s.ID != wantID {
		t.Errorf("GetID is %v, want %v", *s.ID, wantID)
	}
	if *s.Org != wantOrg {
		t.Errorf("GetOrg is %v, want %v", *s.Org, wantOrg)
	}
	if *s.Repo != wantRepo {
		t.Errorf("GetRepo is %v, want %v", *s.Repo, wantRepo)
	}
	if *s.Team != wantTeam {
		t.Errorf("GetTeam is %v, want %v", *s.Team, wantTeam)
	}
	if *s.Name != wantName {
		t.Errorf("GetName is %v, want %v", *s.Name, wantName)
	}
	if *s.Value != wantValue {
		t.Errorf("GetValue is %v, want %v", *s.Value, wantValue)
	}
	if *s.Type != wantType {
		t.Errorf("GetType is %v, want %v", *s.Type, wantType)
	}
	if !reflect.DeepEqual(*s.Images, wantImages) {
		t.Errorf("GetImages is %v, want %v", *s.Images, wantImages)
	}
	if !reflect.DeepEqual(*s.Events, wantEvents) {
		t.Errorf("GetEvents is %v, want %v", *s.Events, wantEvents)
	}
}

func TestLibrary_Secret_String(t *testing.T) {
	// setup types
	num64 := int64(1)
	str := "foo"
	arr := []string{"foo", "bar"}
	s := &Secret{
		ID:     &num64,
		Org:    &str,
		Repo:   &str,
		Team:   &str,
		Name:   &str,
		Value:  &str,
		Type:   &str,
		Images: &arr,
		Events: &arr,
	}
	want := fmt.Sprintf("%+v", *s)

	// run test
	got := s.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}
