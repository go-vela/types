// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-vela/types/constants"
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
	cmd := true
	value := constants.SecretMask
	s := &Secret{
		ID:           &one,
		Org:          &foo,
		Repo:         &bar,
		Name:         &foo,
		Value:        &bar,
		Type:         &repo,
		Images:       &images,
		Events:       &events,
		AllowCommand: &cmd,
	}

	want := &Secret{
		ID:           &one,
		Org:          &foo,
		Repo:         &bar,
		Name:         &foo,
		Value:        &value,
		Type:         &repo,
		Images:       &images,
		Events:       &events,
		AllowCommand: &cmd,
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
	booL := false

	// setup types
	tests := []struct {
		step *pipeline.Container
		sec  *Secret
		want bool
	}{
		{ // test matching secret events
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
				Environment: map[string]string{"BUILD_EVENT": "comment"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{}, Events: &[]string{"comment"}},
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

		{ // test matching secret images
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

		{ // test matching secret events and images
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
		{
			step: &pipeline.Container{
				Image:       "alpine:latest",
				Environment: map[string]string{"BUILD_EVENT": "push"},
				Ruleset: pipeline.Ruleset{
					If: pipeline.Rules{
						Event: []string{"push"},
					},
				},
				Commands: []string{"echo hi"},
			},
			sec:  &Secret{Name: &v, Value: &v, Images: &[]string{"alpine"}, Events: &[]string{"push"}, AllowCommand: &booL},
			want: false,
		},
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
	booL := true
	s := &Secret{
		ID:           &num64,
		Org:          &str,
		Repo:         &str,
		Team:         &str,
		Name:         &str,
		Value:        &str,
		Type:         &str,
		Images:       &arr,
		Events:       &arr,
		AllowCommand: &booL,
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
	wantAllowCommand := booL

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
	gotAllowCommand := s.GetAllowCommand()

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
	if gotAllowCommand != wantAllowCommand {
		t.Errorf("GetAllowCommand is %v, want %v", gotAllowCommand, wantAllowCommand)
	}
}

func TestLibrary_Secret_Getters_Empty(t *testing.T) {
	// setup types
	s := new(Secret)

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
	gotAllowCommand := s.GetAllowCommand()

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
	if gotAllowCommand != false {
		t.Errorf("GetAllowCommand is %v, want false", gotAllowCommand)
	}
}

func TestLibrary_Secret_Setters(t *testing.T) {
	// setup types
	num64 := int64(1)
	str := "foo"
	arr := []string{"foo", "bar"}
	booL := true

	s := new(Secret)

	wantID := num64
	wantOrg := str
	wantRepo := str
	wantTeam := str
	wantName := str
	wantValue := str
	wantType := str
	wantImages := arr
	wantEvents := arr
	wantAllowCommand := booL

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
	s.SetAllowCommand(wantAllowCommand)

	if s.GetID() != wantID {
		t.Errorf("SetID is %v, want %v", s.GetID(), wantID)
	}
	if s.GetOrg() != wantOrg {
		t.Errorf("SetOrg is %v, want %v", s.GetOrg(), wantOrg)
	}
	if s.GetRepo() != wantRepo {
		t.Errorf("SetRepo is %v, want %v", s.GetRepo(), wantRepo)
	}
	if s.GetTeam() != wantTeam {
		t.Errorf("SetTeam is %v, want %v", s.GetTeam(), wantTeam)
	}
	if s.GetName() != wantName {
		t.Errorf("SetName is %v, want %v", s.GetName(), wantName)
	}
	if s.GetValue() != wantValue {
		t.Errorf("SetValue is %v, want %v", s.GetValue(), wantValue)
	}
	if s.GetType() != wantType {
		t.Errorf("SetType is %v, want %v", s.GetType(), wantType)
	}
	if !reflect.DeepEqual(s.GetImages(), wantImages) {
		t.Errorf("SetImages is %v, want %v", s.GetImages(), wantImages)
	}
	if !reflect.DeepEqual(s.GetEvents(), wantEvents) {
		t.Errorf("SetEvents is %v, want %v", s.GetEvents(), wantEvents)
	}
	if s.GetAllowCommand() != wantAllowCommand {
		t.Errorf("SetAllowCommand is %v, want %v", s.GetAllowCommand(), wantAllowCommand)
	}
}

func TestLibrary_Secret_Setters_Empty(t *testing.T) {
	// setup types
	var s *Secret

	// run test
	s.SetID(0)
	s.SetOrg("")
	s.SetRepo("")
	s.SetTeam("")
	s.SetName("")
	s.SetValue("")
	s.SetType("")
	s.SetImages([]string{})
	s.SetEvents([]string{})
	s.SetAllowCommand(false)

	if s.GetID() != 0 {
		t.Errorf("SetID is %v, want 0", s.GetID())
	}
	if s.GetOrg() != "" {
		t.Errorf("SetOrg is %v, want \"\"", s.GetOrg())
	}
	if s.GetRepo() != "" {
		t.Errorf("SetRepo is %v, want \"\"", s.GetRepo())
	}
	if s.GetTeam() != "" {
		t.Errorf("SetTeam is %v, want \"\"", s.GetTeam())
	}
	if s.GetName() != "" {
		t.Errorf("SetName is %v, want \"\"", s.GetName())
	}
	if s.GetValue() != "" {
		t.Errorf("SetValue is %v, want \"\"", s.GetValue())
	}
	if s.GetType() != "" {
		t.Errorf("SetType is %v, want \"\"", s.GetType())
	}
	if !reflect.DeepEqual(s.GetImages(), []string{}) {
		t.Errorf("SetImages is %v, want []string{}", s.GetImages())
	}
	if !reflect.DeepEqual(s.GetEvents(), []string{}) {
		t.Errorf("SetEvents is %v, want []string{}", s.GetEvents())
	}
	if s.GetAllowCommand() != false {
		t.Errorf("SetAllowCommand is %v, want false", s.GetAllowCommand())
	}
}

func TestLibrary_Secret_String(t *testing.T) {
	// setup types
	num64 := int64(1)
	str := "foo"
	arr := []string{"foo", "bar"}
	booL := true
	s := &Secret{
		ID:           &num64,
		Org:          &str,
		Repo:         &str,
		Team:         &str,
		Name:         &str,
		Value:        &str,
		Type:         &str,
		Images:       &arr,
		Events:       &arr,
		AllowCommand: &booL,
	}
	want := fmt.Sprintf("%+v", *s)

	// run test
	got := s.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}
