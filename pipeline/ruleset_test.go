// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"testing"
)

func TestPipeline_Ruleset_Match(t *testing.T) {
	// setup types
	tests := []struct {
		ruleset *Ruleset
		data    *RuleData
		want    bool
	}{
		// Empty
		{ruleset: &Ruleset{}, data: &RuleData{Branch: "master"}, want: true},
		// If with and operator
		{
			ruleset: &Ruleset{If: Rules{Branch: []string{"master"}}},
			data:    &RuleData{Branch: "master", Comment: "rerun", Event: "push", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    true,
		},
		{
			ruleset: &Ruleset{If: Rules{Branch: []string{"master"}}},
			data:    &RuleData{Branch: "dev", Comment: "rerun", Event: "push", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    false,
		},
		{
			ruleset: &Ruleset{If: Rules{Branch: []string{"master"}, Event: []string{"push"}}},
			data:    &RuleData{Branch: "master", Comment: "rerun", Event: "push", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    true,
		},
		{
			ruleset: &Ruleset{If: Rules{Branch: []string{"master"}, Event: []string{"push"}}},
			data:    &RuleData{Branch: "master", Comment: "rerun", Event: "pull_request", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    false,
		},
		{
			ruleset: &Ruleset{If: Rules{Path: []string{"foo.txt", "/foo/bar.txt"}}},
			data:    &RuleData{Branch: "master", Comment: "rerun", Event: "pull_request", Path: []string{}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    false,
		},
		{
			ruleset: &Ruleset{If: Rules{Comment: []string{"rerun"}}},
			data:    &RuleData{Branch: "dev", Comment: "rerun", Event: "push", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    true,
		},
		{
			ruleset: &Ruleset{If: Rules{Comment: []string{"rerun"}}},
			data:    &RuleData{Branch: "dev", Comment: "ok to test", Event: "push", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    false,
		},
		{
			ruleset: &Ruleset{If: Rules{Event: []string{"deployment"}, Target: []string{"production"}}},
			data:    &RuleData{Branch: "dev", Comment: "", Event: "deployment", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: "production"},
			want:    true,
		},
		{
			ruleset: &Ruleset{If: Rules{Event: []string{"deployment"}, Target: []string{"production"}}},
			data:    &RuleData{Branch: "dev", Comment: "", Event: "deployment", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: "stage"},
			want:    false,
		},
		{
			ruleset: &Ruleset{If: Rules{Status: []string{"success", "failure"}}},
			data:    &RuleData{Branch: "dev", Comment: "ok to test", Event: "push", Repo: "octocat/hello-world", Status: "failure", Tag: "refs/heads/master", Target: ""},
			want:    true,
		},
		// If with or operator
		{
			ruleset: &Ruleset{If: Rules{Branch: []string{"master"}, Event: []string{"push"}}, Operator: "or"},
			data:    &RuleData{Branch: "master", Comment: "rerun", Event: "push", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    true,
		},
		{
			ruleset: &Ruleset{If: Rules{Branch: []string{"master"}, Event: []string{"push"}}, Operator: "or"},
			data:    &RuleData{Branch: "dev", Comment: "rerun", Event: "push", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    true,
		},
		{
			ruleset: &Ruleset{If: Rules{Branch: []string{"master"}, Event: []string{"push"}}, Operator: "or"},
			data:    &RuleData{Branch: "master", Comment: "rerun", Event: "pull_request", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    true,
		},
		{
			ruleset: &Ruleset{If: Rules{Branch: []string{"master"}, Event: []string{"push"}}, Operator: "or"},
			data:    &RuleData{Branch: "dev", Comment: "rerun", Event: "pull_request", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    false,
		},
		{
			ruleset: &Ruleset{If: Rules{Path: []string{"foo.txt", "/foo/bar.txt"}}, Operator: "or"},
			data:    &RuleData{Branch: "master", Comment: "rerun", Event: "pull_request", Path: []string{}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    false,
		},
		// Unless with and operator
		{
			ruleset: &Ruleset{Unless: Rules{Branch: []string{"master"}}},
			data:    &RuleData{Branch: "master", Comment: "rerun", Event: "push", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    false,
		},
		{
			ruleset: &Ruleset{Unless: Rules{Branch: []string{"master"}}},
			data:    &RuleData{Branch: "dev", Comment: "rerun", Event: "push", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    true,
		},
		{
			ruleset: &Ruleset{Unless: Rules{Branch: []string{"master"}, Event: []string{"push"}}},
			data:    &RuleData{Branch: "master", Comment: "rerun", Event: "push", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    false,
		},
		{
			ruleset: &Ruleset{Unless: Rules{Branch: []string{"master"}, Event: []string{"push"}}},
			data:    &RuleData{Branch: "master", Comment: "rerun", Event: "pull_request", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    true,
		},
		{
			ruleset: &Ruleset{Unless: Rules{Path: []string{"foo.txt", "/foo/bar.txt"}}},
			data:    &RuleData{Branch: "master", Comment: "rerun", Event: "pull_request", Path: []string{}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    true,
		},
		// Unless with or operator
		{
			ruleset: &Ruleset{Unless: Rules{Branch: []string{"master"}, Event: []string{"push"}}, Operator: "or"},
			data:    &RuleData{Branch: "master", Comment: "rerun", Event: "push", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    false,
		},
		{
			ruleset: &Ruleset{Unless: Rules{Branch: []string{"master"}, Event: []string{"push"}}, Operator: "or"},
			data:    &RuleData{Branch: "dev", Comment: "rerun", Event: "push", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    false,
		},
		{
			ruleset: &Ruleset{Unless: Rules{Branch: []string{"master"}, Event: []string{"push"}}, Operator: "or"},
			data:    &RuleData{Branch: "master", Comment: "rerun", Event: "pull_request", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    false,
		},
		{
			ruleset: &Ruleset{Unless: Rules{Branch: []string{"master"}, Event: []string{"push"}}, Operator: "or"},
			data:    &RuleData{Branch: "dev", Comment: "rerun", Event: "pull_request", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    true,
		},
		{
			ruleset: &Ruleset{Unless: Rules{Path: []string{"foo.txt", "/foo/bar.txt"}}, Operator: "or"},
			data:    &RuleData{Branch: "master", Comment: "rerun", Event: "pull_request", Path: []string{}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want:    true,
		},
		// Advanced Rulesets
		{
			ruleset: &Ruleset{
				If: Rules{
					Event: []string{"push", "pull_request"},
					Tag:   []string{"release/*"},
				},
				Operator: "or",
			},
			data: &RuleData{Branch: "master", Comment: "rerun", Event: "tag", Repo: "octocat/hello-world", Status: "pending", Tag: "release/*", Target: ""},
			want: true,
		},
		{
			ruleset: &Ruleset{
				If: Rules{
					Event: []string{"push", "pull_request"},
					Tag:   []string{"release/*"},
				},
				Operator: "or",
			},
			data: &RuleData{Branch: "master", Comment: "rerun", Event: "tag", Repo: "octocat/hello-world", Status: "pending", Tag: "release/*", Target: ""},
			want: true,
		},
		{
			ruleset: &Ruleset{
				If: Rules{
					Event: []string{"push", "pull_request"},
					Tag:   []string{"release/*"},
				},
				Operator: "or",
			},
			data: &RuleData{Branch: "master", Comment: "rerun", Event: "tag", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			want: false,
		},
	}

	// run test
	for _, test := range tests {
		got := test.ruleset.Match(test.data)

		if got != test.want {
			t.Errorf("Ruleset Match for %s operator is %v, want %v", test.ruleset.Operator, got, test.want)
		}
	}
}

func TestPipeline_Rules_Empty(t *testing.T) {
	// setup types
	r := Rules{}

	// run test
	got := r.Empty()

	if !got {
		t.Errorf("Rule IsEmpty is %v, want true", got)
	}
}

func TestPipeline_Rules_Empty_Invalid(t *testing.T) {
	// setup types
	r := Rules{Branch: []string{"master"}}

	// run test
	got := r.Empty()

	if got {
		t.Errorf("Rule IsEmpty is %v, want false", got)
	}
}

func TestPipeline_Rules_Version_Regex_Tag(t *testing.T) {
	// setup types
	tests := []struct {
		rules    *Rules
		data     *RuleData
		operator string
		want     bool
	}{
		{
			rules:    &Rules{Event: []string{"tag"}, Tag: []string{"refs/tags/20.*"}},
			data:     &RuleData{Branch: "master", Event: "tag", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/tags/20.4.42.167", Target: ""},
			operator: "and",
			want:     true,
		},
		{
			rules:    &Rules{Event: []string{"tag"}, Tag: []string{"[0-9][0-9].[0-9].[0-9][0-9].[0-9][0-9][0-9]"}},
			data:     &RuleData{Branch: "master", Event: "tag", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/tags/20.4.42.167", Target: ""},
			operator: "and",
			want:     true,
		},
		{
			rules:    &Rules{Event: []string{"tag"}, Tag: []string{"[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]+"}},
			data:     &RuleData{Branch: "master", Event: "tag", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/tags/20.4.42.167", Target: ""},
			operator: "and",
			want:     true,
		},
		{
			rules:    &Rules{Event: []string{"tag"}, Tag: []string{"^refs/tags/(\\d+\\.)+\\d+$"}},
			data:     &RuleData{Branch: "master", Event: "tag", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/tags/20.4.42.167", Target: ""},
			operator: "and",
			want:     true,
		},
		{
			rules:    &Rules{Event: []string{"tag"}, Tag: []string{"^refs/tags/(\\d+\\.)+\\d+"}},
			data:     &RuleData{Branch: "master", Event: "tag", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/tags/2.4.42.165-prod", Target: ""},
			operator: "and",
			want:     true,
		},
		{
			rules:    &Rules{Event: []string{"tag"}, Tag: []string{"^refs/tags/(\\d+\\.)+\\d+$"}},
			data:     &RuleData{Branch: "master", Event: "tag", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/tags/2.4.42.165-prod", Target: ""},
			operator: "and",
			want:     false,
		},
	}

	// run test
	for _, test := range tests {
		got := test.rules.Match(test.data, test.operator)

		if got != test.want {
			t.Errorf("Rules Match for %s operator is %v, want %v", test.operator, got, test.want)
		}
	}
}

func TestPipeline_Rules_Match(t *testing.T) {
	// setup types
	tests := []struct {
		rules    *Rules
		data     *RuleData
		operator string
		want     bool
	}{
		// Empty
		{
			rules:    &Rules{},
			data:     &RuleData{Branch: "master", Event: "push", Path: []string{"foo.txt", "/foo/bar.txt"}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			operator: "and",
			want:     true,
		},
		{
			rules:    &Rules{},
			data:     &RuleData{Branch: "master", Event: "push", Path: []string{"foo.txt", "/foo/bar.txt"}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			operator: "or",
			want:     false,
		},
		// and operator
		{
			rules:    &Rules{Branch: []string{"master"}},
			data:     &RuleData{Branch: "master", Event: "push", Path: []string{"foo.txt", "/foo/bar.txt"}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			operator: "and",
			want:     true,
		},
		{
			rules:    &Rules{Branch: []string{"master"}},
			data:     &RuleData{Branch: "dev", Event: "push", Path: []string{"foo.txt", "/foo/bar.txt"}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			operator: "and",
			want:     false,
		},
		{
			rules:    &Rules{Branch: []string{"master"}, Event: []string{"push"}},
			data:     &RuleData{Branch: "master", Event: "push", Path: []string{"foo.txt", "/foo/bar.txt"}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			operator: "and",
			want:     true,
		},
		{
			rules:    &Rules{Branch: []string{"master"}, Event: []string{"push"}},
			data:     &RuleData{Branch: "master", Event: "pull_request", Path: []string{"foo.txt", "/foo/bar.txt"}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			operator: "and",
			want:     false,
		},
		{
			rules:    &Rules{Path: []string{"foob.txt"}},
			data:     &RuleData{Branch: "master", Event: "pull_request", Path: []string{"foo.txt", "/foo/bar.txt"}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			operator: "and",
			want:     false,
		},
		{
			rules:    &Rules{Status: []string{"success", "failure"}},
			data:     &RuleData{Branch: "master", Event: "pull_request", Path: []string{"foo.txt", "/foo/bar.txt"}, Repo: "octocat/hello-world", Tag: "refs/heads/master", Target: ""},
			operator: "and",
			want:     true,
		},
		{
			rules:    &Rules{Event: []string{"tag"}, Tag: []string{"refs/tags/[0-9].*-prod"}},
			data:     &RuleData{Branch: "master", Event: "tag", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/tags/2.4.42.167-prod", Target: ""},
			operator: "and",
			want:     true,
		},
		{
			rules:    &Rules{Event: []string{"tag"}, Tag: []string{"path/to/thing/*/*"}},
			data:     &RuleData{Branch: "master", Event: "tag", Repo: "octocat/hello-world", Status: "pending", Tag: "path/to/thing/stage/1.0.2-rc", Target: ""},
			operator: "and",
			want:     true,
		},
		// or operator
		{
			rules:    &Rules{Branch: []string{"master"}, Event: []string{"push"}},
			data:     &RuleData{Branch: "master", Event: "push", Path: []string{"foo.txt", "/foo/bar.txt"}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			operator: "or",
			want:     true,
		},
		{
			rules:    &Rules{Branch: []string{"master"}, Event: []string{"push"}},
			data:     &RuleData{Branch: "dev", Event: "push", Path: []string{"foo.txt", "/foo/bar.txt"}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			operator: "or",
			want:     true,
		},
		{
			rules:    &Rules{Branch: []string{"master"}, Event: []string{"push"}},
			data:     &RuleData{Branch: "master", Event: "pull_request", Path: []string{"foo.txt", "/foo/bar.txt"}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			operator: "or",
			want:     true,
		},
		{
			rules:    &Rules{Branch: []string{"master"}, Event: []string{"push"}},
			data:     &RuleData{Branch: "dev", Event: "pull_request", Path: []string{"foo.txt", "/foo/bar.txt"}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			operator: "or",
			want:     false,
		},
		{
			rules:    &Rules{Path: []string{"foob.txt"}},
			data:     &RuleData{Branch: "dev", Event: "pull_request", Path: []string{"foo.txt", "/foo/bar.txt"}, Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			operator: "or",
			want:     false,
		},
		// Advanced Rulesets
		{
			rules:    &Rules{Event: []string{"push", "pull_request"}, Tag: []string{"release/*"}},
			data:     &RuleData{Branch: "master", Event: "push", Repo: "octocat/hello-world", Status: "pending", Tag: "release/*", Target: ""},
			operator: "or",
			want:     true,
		},
		{
			rules:    &Rules{Event: []string{"push", "pull_request"}, Tag: []string{"release/*"}},
			data:     &RuleData{Branch: "master", Event: "push", Repo: "octocat/hello-world", Status: "pending", Tag: "release/*", Target: ""},
			operator: "or",
			want:     true,
		},
		{
			rules:    &Rules{Event: []string{"push", "pull_request"}, Tag: []string{"release/*"}},
			data:     &RuleData{Branch: "master", Event: "tag", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: ""},
			operator: "or",
			want:     false,
		},
	}

	// run test
	for _, test := range tests {
		got := test.rules.Match(test.data, test.operator)

		if got != test.want {
			t.Errorf("Rules Match for %s operator is %v, want %v", test.operator, got, test.want)
		}
	}
}

func TestPipeline_Ruletype_MatchAnd(t *testing.T) {
	// setup types
	tests := []struct {
		rule    Ruletype
		pattern string
		want    bool
	}{
		// Empty
		{rule: []string{}, pattern: "master", want: true},
		{rule: []string{}, pattern: "push", want: true},
		{rule: []string{}, pattern: "foo/bar", want: true},
		{rule: []string{}, pattern: "success", want: true},
		{rule: []string{}, pattern: "release/*", want: true},
		// Branch
		{rule: []string{"master"}, pattern: "master", want: true},
		{rule: []string{"master"}, pattern: "dev", want: false},
		// Comment
		{rule: []string{"ok to test"}, pattern: "ok to test", want: true},
		{rule: []string{"ok to test"}, pattern: "rerun", want: false},
		// Event
		{rule: []string{"push"}, pattern: "push", want: true},
		{rule: []string{"push"}, pattern: "pull_request", want: false},
		// Repo
		{rule: []string{"foo/bar"}, pattern: "foo/bar", want: true},
		{rule: []string{"foo/bar"}, pattern: "test/foobar", want: false},
		// Status
		{rule: []string{"success"}, pattern: "success", want: true},
		{rule: []string{"success"}, pattern: "failure", want: false},
		// Tag
		{rule: []string{"release/*"}, pattern: "release/*", want: true},
		{rule: []string{"release/*"}, pattern: "stage/*", want: false},
		// Target
		{rule: []string{"production"}, pattern: "production", want: true},
		{rule: []string{"stage"}, pattern: "production", want: false},
	}

	// run test
	for _, test := range tests {
		got := test.rule.MatchAnd(test.pattern)

		if got != test.want {
			t.Errorf("Ruletype MatchAnd is %v, want %v", got, test.want)
		}
	}
}

func TestPipeline_Ruletype_MatchOr(t *testing.T) {
	// setup types
	tests := []struct {
		rule    Ruletype
		pattern string
		want    bool
	}{
		// Empty
		{rule: []string{}, pattern: "master", want: false},
		{rule: []string{}, pattern: "push", want: false},
		{rule: []string{}, pattern: "foo/bar", want: false},
		{rule: []string{}, pattern: "success", want: false},
		{rule: []string{}, pattern: "release/*", want: false},
		// Branch
		{rule: []string{"master"}, pattern: "master", want: true},
		{rule: []string{"master"}, pattern: "dev", want: false},
		// Comment
		{rule: []string{"ok to test"}, pattern: "ok to test", want: true},
		{rule: []string{"ok to test"}, pattern: "rerun", want: false},
		// Event
		{rule: []string{"push"}, pattern: "push", want: true},
		{rule: []string{"push"}, pattern: "pull_request", want: false},
		// Repo
		{rule: []string{"foo/bar"}, pattern: "foo/bar", want: true},
		{rule: []string{"foo/bar"}, pattern: "test/foobar", want: false},
		// Status
		{rule: []string{"success"}, pattern: "success", want: true},
		{rule: []string{"success"}, pattern: "failure", want: false},
		// Tag
		{rule: []string{"release/*"}, pattern: "release/*", want: true},
		{rule: []string{"release/*"}, pattern: "stage/*", want: false},
		// Target
		{rule: []string{"production"}, pattern: "production", want: true},
		{rule: []string{"stage"}, pattern: "production", want: false},
	}

	// run test
	for _, test := range tests {
		got := test.rule.MatchOr(test.pattern)

		if got != test.want {
			t.Errorf("Ruletype MatchOr is %v, want %v", got, test.want)
		}
	}
}
