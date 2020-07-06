// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/go-vela/types/pipeline"
	"github.com/google/go-cmp/cmp"
	yaml "gopkg.in/yaml.v2"
)

func TestYaml_SecretSlice_ToPipeline(t *testing.T) {
	// setup tests
	tests := []struct {
		secrets *SecretSlice
		want    *pipeline.SecretSlice
	}{
		{
			secrets: &SecretSlice{
				{
					Name:   "docker_username",
					Key:    "github/octocat/docker/username",
					Engine: "native",
					Type:   "repo",
					Origin: Origin{},
				},
				{
					Name:   "docker_username",
					Key:    "",
					Engine: "",
					Type:   "",
					Origin: Origin{
						Environment: map[string]string{"FOO": "bar"},
						Image:       "target/vela-vault:latest",
						Parameters: map[string]interface{}{
							"addr": "vault.company.com",
						},
						Pull: true,
						Ruleset: Ruleset{
							If: Rules{
								Event: []string{"push"},
							},
							Operator: "and",
						},
						Secrets: StepSecretSlice{
							{
								Source: "foo",
								Target: "foo",
							},
							{
								Source: "foobar",
								Target: "foobar",
							},
						},
					},
				},
			},
			want: &pipeline.SecretSlice{
				{
					Name:   "docker_username",
					Key:    "github/octocat/docker/username",
					Engine: "native",
					Type:   "repo",
					Origin: &pipeline.Container{},
				},
				{
					Name:   "docker_username",
					Key:    "",
					Engine: "",
					Type:   "",
					Origin: &pipeline.Container{
						Environment: map[string]string{"FOO": "bar"},
						Image:       "target/vela-vault:latest",
						Pull:        true,
						Ruleset: pipeline.Ruleset{
							If: pipeline.Rules{
								Event: []string{"push"},
							},
							Operator: "and",
						},
						Secrets: pipeline.StepSecretSlice{
							{
								Source: "foo",
								Target: "foo",
							},
							{
								Source: "foobar",
								Target: "foobar",
							},
						},
					},
				},
			},
		},
	}

	// run tests
	for _, test := range tests {
		got := test.secrets.ToPipeline()

		if !reflect.DeepEqual(got, test.want) {
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("MakeGatewayInfo() mismatch (-want +got):\n%s", diff)
			}
			t.Errorf("ToPipeline is %v, want %v", got, test.want)
		}
	}
}

func TestYaml_SecretSlice_UnmarshalYAML(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		file    string
		want    *SecretSlice
	}{
		{
			failure: false,
			file:    "testdata/secret.yml",
			want: &SecretSlice{
				{
					Name:   "foo",
					Key:    "bar",
					Engine: "native",
					Type:   "repo",
				},
				{
					Name:   "noKey",
					Key:    "noKey",
					Engine: "native",
					Type:   "repo",
				},
				{
					Name:   "noType",
					Key:    "bar",
					Engine: "native",
					Type:   "repo",
				},
				{
					Name:   "noEngine",
					Key:    "bar",
					Engine: "native",
					Type:   "repo",
				},
				{
					Name:   "noKeyEngineAndType",
					Key:    "noKeyEngineAndType",
					Engine: "native",
					Type:   "repo",
				},
				{
					Name:   "externalSecret",
					Key:    "",
					Engine: "",
					Type:   "",
					Origin: Origin{
						Environment: map[string]string{"FOO": "bar"},
						Image:       "target/vela-vault:latest",
						Parameters: map[string]interface{}{
							"addr": "vault.company.com",
						},
						Pull: true,
						Ruleset: Ruleset{
							If: Rules{
								Event: []string{"push"},
							},
							Operator: "and",
						},
						Secrets: StepSecretSlice{
							{
								Source: "foo",
								Target: "foo",
							},
							{
								Source: "foobar",
								Target: "foobar",
							},
						},
					},
				},
			},
		},
		{
			failure: true,
			file:    "testdata/invalid.yml",
			want:    nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := new(SecretSlice)

		// run test
		b, err := ioutil.ReadFile(test.file)
		if err != nil {
			t.Errorf("unable to read file: %v", err)
		}

		err = yaml.Unmarshal(b, got)

		if test.failure {
			if err == nil {
				t.Errorf("UnmarshalYAML should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("UnmarshalYAML returned err: %v", err)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("UnmarshalYAML is %v, want %v", got, test.want)
		}
	}
}

func TestYaml_StepSecretSlice_ToPipeline(t *testing.T) {
	// setup tests
	tests := []struct {
		secrets *StepSecretSlice
		want    *pipeline.StepSecretSlice
	}{
		{
			secrets: &StepSecretSlice{
				{
					Source: "docker_username",
					Target: "plugin_username",
				},
			},
			want: &pipeline.StepSecretSlice{
				{
					Source: "docker_username",
					Target: "plugin_username",
				},
			},
		},
	}

	// run tests
	for _, test := range tests {
		got := test.secrets.ToPipeline()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ToPipeline is %v, want %v", got, test.want)
		}
	}
}

func TestYaml_StepSecretSlice_UnmarshalYAML(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		file    string
		want    *StepSecretSlice
	}{
		{
			failure: false,
			file:    "testdata/step_secret_slice.yml",
			want: &StepSecretSlice{
				{
					Source: "foo",
					Target: "bar",
				},
				{
					Source: "hello",
					Target: "world",
				},
			},
		},
		{
			failure: false,
			file:    "testdata/step_secret_string.yml",
			want: &StepSecretSlice{
				{
					Source: "foo",
					Target: "foo",
				},
				{
					Source: "hello",
					Target: "hello",
				},
			},
		},
		{
			failure: true,
			file:    "testdata/invalid.yml",
			want:    nil,
		},
	}

	// run tests
	for _, test := range tests {
		got := new(StepSecretSlice)

		// run test
		b, err := ioutil.ReadFile(test.file)
		if err != nil {
			t.Errorf("unable to read file: %v", err)
		}

		err = yaml.Unmarshal(b, got)

		if test.failure {
			if err == nil {
				t.Errorf("UnmarshalYAML should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("UnmarshalYAML returned err: %v", err)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("UnmarshalYAML is %v, want %v", got, test.want)
		}
	}
}
