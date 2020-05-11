// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"path/filepath"
	"strings"
)

type (
	// Ruleset is the pipeline representation of
	// a ruleset block for a step in a pipeline.
	Ruleset struct {
		If       Rules  `json:"if,omitempty"       yaml:"if,omitempty"`
		Unless   Rules  `json:"unless,omitempty"   yaml:"unless,omitempty"`
		Operator string `json:"operator,omitempty" yaml:"operator,omitempty"`
		Continue bool   `json:"continue,omitempty" yaml:"continue,omitempty"`
	}

	// Rules is the pipeline representation of the ruletypes
	// from a ruleset block for a step in a pipeline.
	Rules struct {
		Branch  Ruletype `json:"branch,omitempty"  yaml:"branch,omitempty"`
		Comment Ruletype `json:"comment,omitempty" yaml:"comment,omitempty"`
		Event   Ruletype `json:"event,omitempty"   yaml:"event,omitempty"`
		Path    Ruletype `json:"path,omitempty"    yaml:"path,omitempty"`
		Repo    Ruletype `json:"repo,omitempty"    yaml:"repo,omitempty"`
		Status  Ruletype `json:"status,omitempty"  yaml:"status,omitempty"`
		Tag     Ruletype `json:"tag,omitempty"     yaml:"tag,omitempty"`
		Target  Ruletype `json:"target,omitempty"  yaml:"target,omitempty"`
	}

	// Ruletype is the pipeline representation of an element
	// for a ruleset block for a step in a pipeline.
	Ruletype []string

	// RuleData is the data to check our ruleset
	// against for a step in a pipeline.
	RuleData struct {
		Branch  string   `json:"branch,omitempty"  yaml:"branch,omitempty"`
		Comment string   `json:"comment,omitempty" yaml:"comment,omitempty"`
		Event   string   `json:"event,omitempty"   yaml:"event,omitempty"`
		Path    []string `json:"path,omitempty"    yaml:"path,omitempty"`
		Repo    string   `json:"repo,omitempty"    yaml:"repo,omitempty"`
		Status  string   `json:"status,omitempty"  yaml:"status,omitempty"`
		Tag     string   `json:"tag,omitempty"     yaml:"tag,omitempty"`
		Target  string   `json:"target,omitempty"  yaml:"target,omitempty"`
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
		if r.Unless.Match(from, r.Operator) {
			return false
		}
	}

	// return true when the if rules are empty
	if r.If.Empty() {
		return true
	}

	// return true when the if rules match
	if r.If.Match(from, r.Operator) {
		return true
	}

	// return false if not match is found
	return false
}

// Execute returns true when the provided ruledata matches
// the if rules and does not match any of the unless rules.
// This function does not check When the provided if and Unless
// rules are empty.
func (r *Ruleset) Execute(from *RuleData) bool {
	// return false when the unless rules are not empty and match
	if !r.Unless.Empty() {
		if r.Unless.Match(from, r.Operator) {
			return false
		}
	}

	// return true when the if rules match
	if r.If.Match(from, r.Operator) {
		return true
	}

	// return false if not match is found
	return false
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
func (r *Rules) Match(from *RuleData, op string) bool {
	// set defaults
	status := true

	// if the path ruletype is provided
	if len(from.Path) > 0 {
		// if the "or" operator is provided in the ruleset
		if strings.EqualFold(op, "or") {

			// override the default to the "or"
			if len(from.Status) != 0 {
				status = r.Status.MatchOr(from.Status)
			}

			// iterate through each path in the ruletype
			for _, p := range from.Path {

				// return true if any ruletype matches the ruledata
				if r.Branch.MatchOr(from.Branch) ||
					r.Comment.MatchOr(from.Comment) ||
					r.Event.MatchOr(from.Event) ||
					r.Path.MatchOr(p) ||
					r.Repo.MatchOr(from.Repo) ||
					status ||
					r.Tag.MatchOr(from.Tag) ||
					r.Target.MatchOr(from.Target) {
					return true
				}
			}

			// return false if no match is found
			return false
		}

		// override the default to the "and"
		if len(from.Status) != 0 {
			status = r.Status.MatchAnd(from.Status)
		}

		// iterate through each path in the ruletype
		for _, p := range from.Path {

			// return true if every ruletype matches the ruledata
			if r.Branch.MatchAnd(from.Branch) &&
				r.Comment.MatchAnd(from.Comment) &&
				r.Event.MatchAnd(from.Event) &&
				r.Path.MatchAnd(p) &&
				r.Repo.MatchAnd(from.Repo) &&
				status &&
				r.Tag.MatchAnd(from.Tag) &&
				r.Target.MatchAnd(from.Target) {
				return true
			}
		}

		// return false if no match is found
		return false
	}

	// if the "or" operator is provided in the ruleset
	if strings.EqualFold(op, "or") {

		// override the default to the "or"
		if len(from.Status) != 0 {
			status = r.Status.MatchOr(from.Status)
		}

		// return true if any ruletype matches the ruledata
		if r.Branch.MatchOr(from.Branch) ||
			r.Comment.MatchOr(from.Comment) ||
			r.Event.MatchOr(from.Event) ||
			r.Path.MatchOr("") ||
			r.Repo.MatchOr(from.Repo) ||
			status ||
			r.Tag.MatchOr(from.Tag) ||
			r.Target.MatchOr(from.Target) {
			return true
		}

		// return false if no match is found
		return false
	}

	// override the default to the "and"
	if len(from.Status) != 0 {
		status = r.Status.MatchAnd(from.Status)
	}

	// return true if every ruletype matches the ruledata
	if r.Branch.MatchAnd(from.Branch) &&
		r.Comment.MatchAnd(from.Comment) &&
		r.Event.MatchAnd(from.Event) &&
		r.Path.MatchAnd("") &&
		r.Repo.MatchAnd(from.Repo) &&
		status &&
		r.Tag.MatchAnd(from.Tag) &&
		r.Target.MatchAnd(from.Target) {
		return true
	}

	// return false if no match is found
	return false
}

// MatchAnd returns true when the provided ruletype
// matches the provided ruledata. When the provided
// ruletype is empty, the function returns true.
func (r *Ruletype) MatchAnd(data string) bool {
	// return true if an empty ruletype is provided
	if len(*r) == 0 {
		return true
	}

	// iterate through each pattern in the ruletype
	for _, pattern := range *r {

		// return true if the pattern matches the ruledata
		if ok, _ := filepath.Match(pattern, data); ok {
			return true
		}
	}

	// return false if no match is found
	return false
}

// MatchOr returns true when the provided ruletype
// matches the provided ruledata. When the provided
// ruletype is empty, the function returns false.
func (r *Ruletype) MatchOr(data string) bool {
	// return false if an empty ruletype is provided
	if len(*r) == 0 {
		return false
	}

	// iterate through each pattern in the ruletype
	for _, pattern := range *r {

		// return true if the pattern matches the ruledata
		if ok, _ := filepath.Match(pattern, data); ok {
			return true
		}
	}

	// return false if no match is found
	return false
}
