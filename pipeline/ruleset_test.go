// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
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
			ruleset: &Ruleset{If: Rules{Event: []string{"schedule"}, Target: []string{"weekly"}}},
			data:    &RuleData{Branch: "dev", Comment: "", Event: "schedule", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: "weekly"},
			want:    true,
		},
		{
			ruleset: &Ruleset{If: Rules{Event: []string{"schedule"}, Target: []string{"weekly"}}},
			data:    &RuleData{Branch: "dev", Comment: "", Event: "schedule", Repo: "octocat/hello-world", Status: "pending", Tag: "refs/heads/master", Target: "nightly"},
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

func TestPipeline_Rules_NoStatus(t *testing.T) {
	// setup types
	r := Rules{}

	// run test
	got := r.Empty()

	if !got {
		t.Errorf("Rule NoStatus is %v, want true", got)
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

func TestPipeline_Rules_Match_Regex_Tag(t *testing.T) {
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
		got := test.rules.Match(test.data, "regexp", test.operator)

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
		got := test.rules.Match(test.data, "filepath", test.operator)

		if got != test.want {
			t.Errorf("Rules Match for %s operator is %v, want %v", test.operator, got, test.want)
		}
	}
}

func TestPipeline_Ruletype_MatchAnd(t *testing.T) {
	// setup types
	tests := []struct {
		matcher string
		rule    Ruletype
		pattern string
		want    bool
	}{
		// Empty with filepath matcher
		{matcher: "filepath", rule: []string{}, pattern: "master", want: true},
		{matcher: "filepath", rule: []string{}, pattern: "push", want: true},
		{matcher: "filepath", rule: []string{}, pattern: "foo/bar", want: true},
		{matcher: "filepath", rule: []string{}, pattern: "success", want: true},
		{matcher: "filepath", rule: []string{}, pattern: "release/*", want: true},
		// Branch with filepath matcher
		{matcher: "filepath", rule: []string{"master"}, pattern: "master", want: true},
		{matcher: "filepath", rule: []string{"master"}, pattern: "dev", want: false},
		// Comment with filepath matcher
		{matcher: "filepath", rule: []string{"ok to test"}, pattern: "ok to test", want: true},
		{matcher: "filepath", rule: []string{"ok to test"}, pattern: "rerun", want: false},
		// Event with filepath matcher
		{matcher: "filepath", rule: []string{"push"}, pattern: "push", want: true},
		{matcher: "filepath", rule: []string{"push"}, pattern: "pull_request", want: false},
		// Repo with filepath matcher
		{matcher: "filepath", rule: []string{"foo/bar"}, pattern: "foo/bar", want: true},
		{matcher: "filepath", rule: []string{"foo/bar"}, pattern: "test/foobar", want: false},
		// Status with filepath matcher
		{matcher: "filepath", rule: []string{"success"}, pattern: "success", want: true},
		{matcher: "filepath", rule: []string{"success"}, pattern: "failure", want: false},
		// Tag with filepath matcher
		{matcher: "filepath", rule: []string{"release/*"}, pattern: "release/*", want: true},
		{matcher: "filepath", rule: []string{"release/*"}, pattern: "stage/*", want: false},
		{matcher: "filepath", rule: []string{"release/*"}, pattern: "release/111.2.3-rc", want: true},
		{matcher: "filepath", rule: []string{"release/**"}, pattern: "release/1.2.3-rc-hold", want: true},
		{matcher: "filepath", rule: []string{"release/*"}, pattern: "release/stage/1.2.3-rc", want: false},
		{matcher: "filepath", rule: []string{"release/*/*"}, pattern: "release/stage/1.2.3-rc", want: true},
		{matcher: "filepath", rule: []string{"release/stage/*"}, pattern: "release/stage/1.2.3-rc", want: true},
		{matcher: "filepath", rule: []string{"release/prod/*"}, pattern: "release/stage/1.2.3-rc", want: false},
		{matcher: "filepath", rule: []string{"release/*"}, pattern: "release/1.2.3-rc", want: true},
		{matcher: "filepath", rule: []string{"release/*"}, pattern: "release/1.2.3", want: true},
		// Target with filepath matcher
		{matcher: "filepath", rule: []string{"production"}, pattern: "production", want: true},
		{matcher: "filepath", rule: []string{"stage"}, pattern: "production", want: false},
		// Empty with regex matcher
		{matcher: "regexp", rule: []string{}, pattern: "master", want: true},
		{matcher: "regexp", rule: []string{}, pattern: "push", want: true},
		{matcher: "regexp", rule: []string{}, pattern: "foo/bar", want: true},
		{matcher: "regexp", rule: []string{}, pattern: "success", want: true},
		{matcher: "regexp", rule: []string{}, pattern: "release/*", want: true},
		// Branch with regex matcher
		{matcher: "regexp", rule: []string{"master"}, pattern: "master", want: true},
		{matcher: "regexp", rule: []string{"master"}, pattern: "dev", want: false},
		// Comment with regex matcher
		{matcher: "regexp", rule: []string{"ok to test"}, pattern: "ok to test", want: true},
		{matcher: "regexp", rule: []string{"ok to test"}, pattern: "rerun", want: false},
		// Event with regex matcher
		{matcher: "regexp", rule: []string{"push"}, pattern: "push", want: true},
		{matcher: "regexp", rule: []string{"push"}, pattern: "pull_request", want: false},
		// Repo with regex matcher
		{matcher: "regexp", rule: []string{"foo/bar"}, pattern: "foo/bar", want: true},
		{matcher: "regexp", rule: []string{"foo/bar"}, pattern: "test/foobar", want: false},
		// Status with regex matcher
		{matcher: "regexp", rule: []string{"success"}, pattern: "success", want: true},
		{matcher: "regexp", rule: []string{"success"}, pattern: "failure", want: false},
		// Tag with regex matcher
		{matcher: "regexp", rule: []string{"release/*"}, pattern: "release/*", want: true},
		{matcher: "regexp", rule: []string{"release/*"}, pattern: "stage/*", want: false},
		{matcher: "regex", rule: []string{"release/[0-9]+.*-rc$"}, pattern: "release/111.2.3-rc", want: true},
		{matcher: "regex", rule: []string{"release/[0-9]+.*-rc$"}, pattern: "release/1.2.3-rc-hold", want: false},
		{matcher: "regexp", rule: []string{"release/*"}, pattern: "release/stage/1.2.3-rc", want: true},
		{matcher: "regexp", rule: []string{"release/*/*"}, pattern: "release/stage/1.2.3-rc", want: true},
		{matcher: "regex", rule: []string{"release/stage/*"}, pattern: "release/stage/1.2.3-rc", want: true},
		{matcher: "regex", rule: []string{"release/prod/*"}, pattern: "release/stage/1.2.3-rc", want: false},
		{matcher: "regexp", rule: []string{"release/[0-9]+.[0-9]+.[0-9]+$"}, pattern: "release/1.2.3-rc", want: false},
		{matcher: "regexp", rule: []string{"release/[0-9]+.[0-9]+.[0-9]+$"}, pattern: "release/1.2.3", want: true},
		// Target with regex matcher
		{matcher: "regexp", rule: []string{"production"}, pattern: "production", want: true},
		{matcher: "regexp", rule: []string{"stage"}, pattern: "production", want: false},
	}

	// run test
	for _, test := range tests {
		got := test.rule.MatchAnd(test.pattern, test.matcher)

		if got != test.want {
			t.Errorf("MatchAnd for %s matcher is %v, want %v", test.matcher, got, test.want)
		}
	}
}

func TestPipeline_Ruletype_MatchOr(t *testing.T) {
	// setup types
	tests := []struct {
		matcher string
		rule    Ruletype
		pattern string
		want    bool
	}{
		// Empty with filepath matcher
		{matcher: "filepath", rule: []string{}, pattern: "master", want: false},
		{matcher: "filepath", rule: []string{}, pattern: "push", want: false},
		{matcher: "filepath", rule: []string{}, pattern: "foo/bar", want: false},
		{matcher: "filepath", rule: []string{}, pattern: "success", want: false},
		{matcher: "filepath", rule: []string{}, pattern: "release/*", want: false},
		// Branch with filepath matcher
		{matcher: "filepath", rule: []string{"master"}, pattern: "master", want: true},
		{matcher: "filepath", rule: []string{"master"}, pattern: "dev", want: false},
		// Comment with filepath matcher
		{matcher: "filepath", rule: []string{"ok to test"}, pattern: "ok to test", want: true},
		{matcher: "filepath", rule: []string{"ok to test"}, pattern: "rerun", want: false},
		// Event with filepath matcher
		{matcher: "filepath", rule: []string{"push"}, pattern: "push", want: true},
		{matcher: "filepath", rule: []string{"push"}, pattern: "pull_request", want: false},
		// Repo with filepath matcher
		{matcher: "filepath", rule: []string{"foo/bar"}, pattern: "foo/bar", want: true},
		{matcher: "filepath", rule: []string{"foo/bar"}, pattern: "test/foobar", want: false},
		// Status with filepath matcher
		{matcher: "filepath", rule: []string{"success"}, pattern: "success", want: true},
		{matcher: "filepath", rule: []string{"success"}, pattern: "failure", want: false},
		// Tag with filepath matcher
		{matcher: "filepath", rule: []string{"release/*"}, pattern: "release/*", want: true},
		{matcher: "filepath", rule: []string{"release/*"}, pattern: "stage/*", want: false},
		// Target with filepath matcher
		{matcher: "filepath", rule: []string{"production"}, pattern: "production", want: true},
		{matcher: "filepath", rule: []string{"stage"}, pattern: "production", want: false},
		// Empty with regexp matcher
		{matcher: "regexp", rule: []string{}, pattern: "master", want: false},
		{matcher: "regexp", rule: []string{}, pattern: "push", want: false},
		{matcher: "regexp", rule: []string{}, pattern: "foo/bar", want: false},
		{matcher: "regexp", rule: []string{}, pattern: "success", want: false},
		{matcher: "regexp", rule: []string{}, pattern: "release/*", want: false},
		// Branch with regexp matcher
		{matcher: "regexp", rule: []string{"master"}, pattern: "master", want: true},
		{matcher: "regexp", rule: []string{"master"}, pattern: "dev", want: false},
		// Comment with regexp matcher
		{matcher: "regexp", rule: []string{"ok to test"}, pattern: "ok to test", want: true},
		{matcher: "regexp", rule: []string{"ok to test"}, pattern: "rerun", want: false},
		// Event with regexp matcher
		{matcher: "regexp", rule: []string{"push"}, pattern: "push", want: true},
		{matcher: "regexp", rule: []string{"push"}, pattern: "pull_request", want: false},
		// Repo with regexp matcher
		{matcher: "regexp", rule: []string{"foo/bar"}, pattern: "foo/bar", want: true},
		{matcher: "regexp", rule: []string{"foo/bar"}, pattern: "test/foobar", want: false},
		// Status with regexp matcher
		{matcher: "regexp", rule: []string{"success"}, pattern: "success", want: true},
		{matcher: "regexp", rule: []string{"success"}, pattern: "failure", want: false},
		// Tag with regexp matcher
		{matcher: "regexp", rule: []string{"release/*"}, pattern: "release/*", want: true},
		{matcher: "regexp", rule: []string{"release/*"}, pattern: "stage/*", want: false},
		// Target with regexp matcher
		{matcher: "regexp", rule: []string{"production"}, pattern: "production", want: true},
		{matcher: "regexp", rule: []string{"stage"}, pattern: "production", want: false},
	}

	// run test
	for _, test := range tests {
		got := test.rule.MatchOr(test.pattern, test.matcher)

		if got != test.want {
			t.Errorf("MatchOr for %s matcher is %v, want %v", test.matcher, got, test.want)
		}
	}
}
