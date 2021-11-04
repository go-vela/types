// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/pipeline"
)

func TestLibrary_Secret_Sanitize(t *testing.T) {
	// setup types
	s := testSecret()

	want := testSecret()
	want.SetValue(constants.SecretMask)

	// run test
	got := s.Sanitize()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sanitize is %v, want %v", got, want)
	}
}

func TestLibrary_Secret_Match(t *testing.T) {
	// setup types
	v := "foo"
	booL := false

	// setup tests
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

	// run tests
	for _, test := range tests {
		got := test.sec.Match(test.step)

		if got != test.want {
			t.Errorf("Match is %v, want %v", got, test.want)
		}
	}
}

func TestLibrary_Secret_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		secret *Secret
		want   *Secret
	}{
		{
			secret: testSecret(),
			want:   testSecret(),
		},
		{
			secret: new(Secret),
			want:   new(Secret),
		},
	}

	// run tests
	for _, test := range tests {
		if test.secret.GetID() != test.want.GetID() {
			t.Errorf("GetID is %v, want %v", test.secret.GetID(), test.want.GetID())
		}
		if test.secret.GetOrg() != test.want.GetOrg() {
			t.Errorf("GetOrg is %v, want %v", test.secret.GetOrg(), test.want.GetOrg())
		}
		if test.secret.GetRepo() != test.want.GetRepo() {
			t.Errorf("GetRepo is %v, want %v", test.secret.GetRepo(), test.want.GetRepo())
		}
		if test.secret.GetTeam() != test.want.GetTeam() {
			t.Errorf("GetTeam is %v, want %v", test.secret.GetTeam(), test.want.GetTeam())
		}
		if test.secret.GetName() != test.want.GetName() {
			t.Errorf("GetName is %v, want %v", test.secret.GetName(), test.want.GetName())
		}
		if test.secret.GetValue() != test.want.GetValue() {
			t.Errorf("GetValue is %v, want %v", test.secret.GetValue(), test.want.GetValue())
		}
		if test.secret.GetType() != test.want.GetType() {
			t.Errorf("GetType is %v, want %v", test.secret.GetType(), test.want.GetType())
		}
		if !reflect.DeepEqual(test.secret.GetImages(), test.want.GetImages()) {
			t.Errorf("GetImages is %v, want %v", test.secret.GetImages(), test.want.GetImages())
		}
		if !reflect.DeepEqual(test.secret.GetEvents(), test.want.GetEvents()) {
			t.Errorf("GetEvents is %v, want %v", test.secret.GetEvents(), test.want.GetEvents())
		}
		if test.secret.GetAllowCommand() != test.want.GetAllowCommand() {
			t.Errorf("GetAllowCommand is %v, want %v", test.secret.GetAllowCommand(), test.want.GetAllowCommand())
		}
		if test.secret.GetCreatedAt() != test.want.GetCreatedAt() {
			t.Errorf("GetCreatedAt is %v, want %v", test.secret.GetCreatedAt(), test.want.GetCreatedAt())
		}
		if test.secret.GetCreatedBy() != test.want.GetCreatedBy() {
			t.Errorf("GetCreatedBy is %v, want %v", test.secret.GetCreatedBy(), test.want.GetCreatedBy())
		}
		if test.secret.GetUpdatedAt() != test.want.GetUpdatedAt() {
			t.Errorf("GetUpdatedAt is %v, want %v", test.secret.GetUpdatedAt(), test.want.GetUpdatedAt())
		}
		if test.secret.GetUpdatedBy() != test.want.GetUpdatedBy() {
			t.Errorf("GetUpdatedBy is %v, want %v", test.secret.GetUpdatedBy(), test.want.GetUpdatedBy())
		}
	}
}

func TestLibrary_Secret_Setters(t *testing.T) {
	// setup types
	var s *Secret

	// setup tests
	tests := []struct {
		secret *Secret
		want   *Secret
	}{
		{
			secret: testSecret(),
			want:   testSecret(),
		},
		{
			secret: s,
			want:   new(Secret),
		},
	}

	// run tests
	for _, test := range tests {
		test.secret.SetID(test.want.GetID())
		test.secret.SetOrg(test.want.GetOrg())
		test.secret.SetRepo(test.want.GetRepo())
		test.secret.SetTeam(test.want.GetTeam())
		test.secret.SetName(test.want.GetName())
		test.secret.SetValue(test.want.GetValue())
		test.secret.SetType(test.want.GetType())
		test.secret.SetImages(test.want.GetImages())
		test.secret.SetEvents(test.want.GetEvents())
		test.secret.SetAllowCommand(test.want.GetAllowCommand())
		test.secret.SetCreatedAt(test.want.GetCreatedAt())
		test.secret.SetCreatedBy(test.want.GetCreatedBy())
		test.secret.SetUpdatedAt(test.want.GetUpdatedAt())
		test.secret.SetUpdatedBy(test.want.GetUpdatedBy())

		if test.secret.GetID() != test.want.GetID() {
			t.Errorf("SetID is %v, want %v", test.secret.GetID(), test.want.GetID())
		}
		if test.secret.GetOrg() != test.want.GetOrg() {
			t.Errorf("SetOrg is %v, want %v", test.secret.GetOrg(), test.want.GetOrg())
		}
		if test.secret.GetRepo() != test.want.GetRepo() {
			t.Errorf("SetRepo is %v, want %v", test.secret.GetRepo(), test.want.GetRepo())
		}
		if test.secret.GetTeam() != test.want.GetTeam() {
			t.Errorf("SetTeam is %v, want %v", test.secret.GetTeam(), test.want.GetTeam())
		}
		if test.secret.GetName() != test.want.GetName() {
			t.Errorf("SetName is %v, want %v", test.secret.GetName(), test.want.GetName())
		}
		if test.secret.GetValue() != test.want.GetValue() {
			t.Errorf("SetValue is %v, want %v", test.secret.GetValue(), test.want.GetValue())
		}
		if test.secret.GetType() != test.want.GetType() {
			t.Errorf("SetType is %v, want %v", test.secret.GetType(), test.want.GetType())
		}
		if !reflect.DeepEqual(test.secret.GetImages(), test.want.GetImages()) {
			t.Errorf("SetImages is %v, want %v", test.secret.GetImages(), test.want.GetImages())
		}
		if !reflect.DeepEqual(test.secret.GetEvents(), test.want.GetEvents()) {
			t.Errorf("SetEvents is %v, want %v", test.secret.GetEvents(), test.want.GetEvents())
		}
		if test.secret.GetAllowCommand() != test.want.GetAllowCommand() {
			t.Errorf("SetAllowCommand is %v, want %v", test.secret.GetAllowCommand(), test.want.GetAllowCommand())
		}
		if test.secret.GetCreatedAt() != test.want.GetCreatedAt() {
			t.Errorf("SetCreatedAt is %v, want %v", test.secret.GetCreatedAt(), test.want.GetCreatedAt())
		}
		if test.secret.GetCreatedBy() != test.want.GetCreatedBy() {
			t.Errorf("SetCreatedBy is %v, want %v", test.secret.GetCreatedBy(), test.want.GetCreatedBy())
		}
		if test.secret.GetUpdatedAt() != test.want.GetUpdatedAt() {
			t.Errorf("SetUpdatedAt is %v, want %v", test.secret.GetUpdatedAt(), test.want.GetUpdatedAt())
		}
		if test.secret.GetUpdatedBy() != test.want.GetUpdatedBy() {
			t.Errorf("SetUpdatedBy is %v, want %v", test.secret.GetUpdatedBy(), test.want.GetUpdatedBy())
		}
	}
}

func TestLibrary_Secret_String(t *testing.T) {
	// setup types
	s := testSecret()

	want := fmt.Sprintf(`{
	AllowCommand: %t,
	Events: %s,
	ID: %d,
	Images: %s,
	Name: %s,
	Org: %s,
	Repo: %s,
	Team: %s,
	Type: %s,
	Value: %s,
	CreatedAt: %d,
	CreatedBy: %d,
	UpdatedAt: %d,
	UpdatedBy: %d,
}`,
		s.GetAllowCommand(),
		s.GetEvents(),
		s.GetID(),
		s.GetImages(),
		s.GetName(),
		s.GetOrg(),
		s.GetRepo(),
		s.GetTeam(),
		s.GetType(),
		s.GetValue(),
		s.GetCreatedAt(),
		s.GetCreatedBy(),
		s.GetUpdatedAt(),
		s.GetUpdatedBy(),
	)

	// run test
	got := s.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

// testSecret is a test helper function to create a Secret
// type with all fields set to a fake value.
func testSecret() *Secret {
	currentTime := time.Now()
	tsCreate := currentTime.UTC().Unix()
	tsUpdate := currentTime.Add(time.Hour * 1).UTC().Unix()
	s := new(Secret)

	s.SetID(1)
	s.SetOrg("github")
	s.SetRepo("octocat")
	s.SetTeam("octokitties")
	s.SetName("foo")
	s.SetValue("bar")
	s.SetType("repo")
	s.SetImages([]string{"alpine"})
	s.SetEvents([]string{"push", "tag", "deployment"})
	s.SetAllowCommand(true)
	s.SetCreatedAt(tsCreate)
	s.SetCreatedBy(12345)
	s.SetUpdatedAt(tsUpdate)
	s.SetUpdatedBy(54321)
	return s
}
