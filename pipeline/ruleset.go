// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-vela/types/constants"
)

type (
	// Ruleset is the pipeline representation of
	// a ruleset block for a step in a pipeline.
	//
	// swagger:model PipelineRuleset
	Ruleset struct {
		If       Rules  `json:"if,omitempty"       yaml:"if,omitempty"`
		Unless   Rules  `json:"unless,omitempty"   yaml:"unless,omitempty"`
		Matcher  string `json:"matcher,omitempty"  yaml:"matcher,omitempty"`
		Operator string `json:"operator,omitempty" yaml:"operator,omitempty"`
		Continue bool   `json:"continue,omitempty" yaml:"continue,omitempty"`
	}

	// Rules is the pipeline representation of the ruletypes
	// from a ruleset block for a step in a pipeline.
	//
	// swagger:model PipelineRules
	Rules struct {
		Branch   Ruletype `json:"branch,omitempty"  yaml:"branch,omitempty"`
		Comment  Ruletype `json:"comment,omitempty" yaml:"comment,omitempty"`
		Event    Ruletype `json:"event,omitempty"   yaml:"event,omitempty"`
		Path     Ruletype `json:"path,omitempty"    yaml:"path,omitempty"`
		Repo     Ruletype `json:"repo,omitempty"    yaml:"repo,omitempty"`
		Status   Ruletype `json:"status,omitempty"  yaml:"status,omitempty"`
		Tag      Ruletype `json:"tag,omitempty"     yaml:"tag,omitempty"`
		Target   Ruletype `json:"target,omitempty"  yaml:"target,omitempty"`
		Parallel bool     `json:"parallel,omitempty" yaml:"parallel,omitempty"`
	}

	// Ruletype is the pipeline representation of an element
	// for a ruleset block for a step in a pipeline.
	//
	// swagger:model PipelineRuletype
	Ruletype []string

	// RuleData is the data to check our ruleset
	// against for a step in a pipeline.
	RuleData struct {
		Branch   string   `json:"branch,omitempty"  yaml:"branch,omitempty"`
		Comment  string   `json:"comment,omitempty" yaml:"comment,omitempty"`
		Event    string   `json:"event,omitempty"   yaml:"event,omitempty"`
		Path     []string `json:"path,omitempty"    yaml:"path,omitempty"`
		Repo     string   `json:"repo,omitempty"    yaml:"repo,omitempty"`
		Status   string   `json:"status,omitempty"  yaml:"status,omitempty"`
		Tag      string   `json:"tag,omitempty"     yaml:"tag,omitempty"`
		Target   string   `json:"target,omitempty"  yaml:"target,omitempty"`
		Parallel bool     `json:"parallel,omitempty" yaml:"parallel,omitempty"`
	}
)

// Match returns true when the provided ruledata matches
// the if rules and does not match any of the unless rules.
// When the provided if rules are empty, the function returns
// true. When both the provided if and unless rules are empty,
// the function also returns true.
func (r *Ruleset) Match(from *RuleData) bool {
	// return true when the if and unless rules are empty
	if r.If.Empty() && r.Unless.Empty() {
		return true
	}

	// return false when the unless rules are not empty and match
	if !r.Unless.Empty() {
		if r.Unless.Match(from, r.Matcher, r.Operator) {
			return false
		}
	}

	// return true when the if rules are empty
	if r.If.Empty() {
		return true
	}

	// return true when the if rules match
	if r.If.Match(from, r.Matcher, r.Operator) {
		return true
	}

	// return false if not match is found
	return false
}

// NoStatus returns true if the status field is empty.
func (r *Rules) NoStatus() bool {
	// return true if every ruletype is empty
	return len(r.Status) == 0
}

// Empty returns true if the provided ruletypes are empty.
func (r *Rules) Empty() bool {
	// return true if every ruletype is empty
	if len(r.Branch) == 0 &&
		len(r.Comment) == 0 &&
		len(r.Event) == 0 &&
		len(r.Path) == 0 &&
		len(r.Repo) == 0 &&
		len(r.Status) == 0 &&
		len(r.Tag) == 0 &&
		len(r.Target) == 0 {
		return true
	}

	// return false if any of the ruletype is provided
	return false
}

// Match returns true for the or operator when one of the
// ruletypes from the rules match the provided ruledata.
// Match returns true for the and operator when all of the
// ruletypes from the rules match the provided ruledata. For
// both operators, when none of the ruletypes from the rules
// match the provided ruledata, the function returns false.
//
//nolint:gocyclo // accepting complexity in this case
func (r *Rules) Match(from *RuleData, matcher, op string) bool {
	// set defaults
	status := true

	// if the path ruletype is provided
	if len(from.Path) > 0 {
		// if the "or" operator is provided in the ruleset
		if strings.EqualFold(op, constants.OperatorOr) {
			// override the default to the "or"
			if len(from.Status) != 0 {
				status = r.Status.MatchOr(from.Status, matcher)
			}

			// iterate through each path in the ruletype
			for _, p := range from.Path {
				// return true if any ruletype matches the ruledata
				if r.Branch.MatchOr(from.Branch, matcher) ||
					r.Comment.MatchOr(from.Comment, matcher) ||
					r.Event.MatchOr(from.Event, matcher) ||
					r.Path.MatchOr(p, matcher) ||
					r.Repo.MatchOr(from.Repo, matcher) ||
					status ||
					r.Tag.MatchOr(from.Tag, matcher) ||
					r.Target.MatchOr(from.Target, matcher) {
					return true
				}
			}

			// return false if no match is found
			return false
		}

		// override the default to the "and"
		if len(from.Status) != 0 {
			status = r.Status.MatchAnd(from.Status, matcher)
		}

		// iterate through each path in the ruletype
		for _, p := range from.Path {
			// return true if every ruletype matches the ruledata
			if r.Branch.MatchAnd(from.Branch, matcher) &&
				r.Comment.MatchAnd(from.Comment, matcher) &&
				r.Event.MatchAnd(from.Event, matcher) &&
				r.Path.MatchAnd(p, matcher) &&
				r.Repo.MatchAnd(from.Repo, matcher) &&
				status &&
				r.Tag.MatchAnd(from.Tag, matcher) &&
				r.Target.MatchAnd(from.Target, matcher) {
				return true
			}
		}

		// return false if no match is found
		return false
	}

	// if the "or" operator is provided in the ruleset
	if strings.EqualFold(op, constants.OperatorOr) {
		// override the default to the "or"
		if len(from.Status) != 0 {
			status = r.Status.MatchOr(from.Status, matcher)
		}

		// return true if any ruletype matches the ruledata
		if r.Branch.MatchOr(from.Branch, matcher) ||
			r.Comment.MatchOr(from.Comment, matcher) ||
			r.Event.MatchOr(from.Event, matcher) ||
			r.Path.MatchOr("", matcher) ||
			r.Repo.MatchOr(from.Repo, matcher) ||
			status ||
			r.Tag.MatchOr(from.Tag, matcher) ||
			r.Target.MatchOr(from.Target, matcher) {
			return true
		}

		// return false if no match is found
		return false
	}

	// override the default to the "and"
	if len(from.Status) != 0 {
		status = r.Status.MatchAnd(from.Status, matcher)
	}

	// return true if every ruletype matches the ruledata
	if r.Branch.MatchAnd(from.Branch, matcher) &&
		r.Comment.MatchAnd(from.Comment, matcher) &&
		r.Event.MatchAnd(from.Event, matcher) &&
		r.Path.MatchAnd("", matcher) &&
		r.Repo.MatchAnd(from.Repo, matcher) &&
		status &&
		r.Tag.MatchAnd(from.Tag, matcher) &&
		r.Target.MatchAnd(from.Target, matcher) {
		return true
	}

	// return false if no match is found
	return false
}

// MatchAnd returns true when the provided ruletype
// matches the provided ruledata. When the provided
// ruletype is empty, the function returns true.
func (r *Ruletype) MatchAnd(data, matcher string) bool {
	// return true if an empty ruletype is provided
	if len(*r) == 0 {
		return true
	}

	// iterate through each pattern in the ruletype
	for _, pattern := range *r {
		// handle the pattern based off the matcher provided
		switch matcher {
		case constants.MatcherRegex, "regex":
			// return true if the regexp pattern matches the ruledata
			if regexp.MustCompile(pattern).MatchString(data) {
				return true
			}
		case constants.MatcherFilepath:
			fallthrough
		default:
			// return true if the pattern matches the ruledata
			ok, _ := filepath.Match(pattern, data)
			if ok {
				return true
			}
		}
	}

	// return false if no match is found
	return false
}

// MatchOr returns true when the provided ruletype
// matches the provided ruledata. When the provided
// ruletype is empty, the function returns false.
func (r *Ruletype) MatchOr(data, matcher string) bool {
	// return false if an empty ruletype is provided
	if len(*r) == 0 {
		return false
	}

	// iterate through each pattern in the ruletype
	for _, pattern := range *r {
		// handle the pattern based off the matcher provided
		switch matcher {
		case constants.MatcherRegex, "regex":
			// return true if the regexp pattern matches the ruledata
			if regexp.MustCompile(pattern).MatchString(data) {
				return true
			}
		case constants.MatcherFilepath:
			fallthrough
		default:
			// return true if the pattern matches the ruledata
			ok, _ := filepath.Match(pattern, data)
			if ok {
				return true
			}
		}
	}

	// return false if no match is found
	return false
}
