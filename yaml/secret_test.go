// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/go-vela/types/pipeline"
	"github.com/goccy/go-yaml"
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
						Name:        "vault",
						Environment: map[string]string{"FOO": "bar"},
						Image:       "target/vela-vault:latest",
						Parameters: map[string]interface{}{
							"addr": "vault.company.com",
						},
						Pull: "always",
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
						Name:        "vault",
						Environment: map[string]string{"FOO": "bar"},
						Image:       "target/vela-vault:latest",
						Pull:        "always",
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
						Pull: "always",
						Ruleset: Ruleset{
							If: Rules{
								Event: []string{"push"},
							},
							Operator: "and",
							Matcher:  "filepath",
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
			file:    "testdata/step_secret_slice_invalid_no_source.yml",
			want:    nil,
		},
		{
			failure: true,
			file:    "testdata/step_secret_slice_invalid_no_target.yml",
			want:    nil,
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

func TestYaml_StepSecretSlice_Validate(t *testing.T) {
	//setup types
	tests := []struct {
		name    string
		file    string
		wantErr bool
	}{
		{
			name:    "failure: no secret name",
			file:    "testdata/secret/validate/no_name.yml",
			wantErr: true,
		},
		{
			name:    "success: repo secret block",
			file:    "testdata/secret/validate/repo.yml",
			wantErr: false,
		},
		{
			name:    "failure: repo secret bad engine yaml tag",
			file:    "testdata/secret/validate/repo_bad_engine.yml",
			wantErr: true,
		},
		{
			name:    "failure: repo secret bad key yaml tag",
			file:    "testdata/secret/validate/repo_bad_key.yml",
			wantErr: true,
		},
		{
			name:    "success: org secret block",
			file:    "testdata/secret/validate/org.yml",
			wantErr: false,
		},
		{
			name:    "failure: org secret bad engine yaml tag",
			file:    "testdata/secret/validate/org_bad_engine.yml",
			wantErr: true,
		},
		{
			name:    "failure: org secret bad engine yaml tag",
			file:    "testdata/secret/validate/org_bad_key.yml",
			wantErr: true,
		},
	}

	// run tests
	for _, test := range tests {
		b := new(Build)

		pipeline, err := ioutil.ReadFile(test.file)
		if err != nil {
			t.Errorf("Reading file for Validate returned err: %v", err)
		}

		err = yaml.Unmarshal(pipeline, b)

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}

		t.Run(test.name, func(t *testing.T) {
			if err := b.Secrets.Validate(pipeline); (err != nil) != test.wantErr {
				t.Errorf("Validate is %v, want %v", err, test.wantErr)
			}
		})
	}
}
