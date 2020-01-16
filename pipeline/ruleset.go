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
		If       Rules  `json:"if,omitempty"`
		Unless   Rules  `json:"unless,omitempty"`
		Operator string `json:"operator,omitempty"`
		Continue bool   `json:"continue,omitempty"`
	}

	// Rules is the pipeline representation of the ruletypes
	// from a ruleset block for a step in a pipeline.
	Rules struct {
		Branch Ruletype `json:"branch,omitempty"`
		Event  Ruletype `json:"event,omitempty"`
		Path   Ruletype `json:"path,omitempty"`
		Repo   Ruletype `json:"repo,omitempty"`
		Status Ruletype `json:"status,omitempty"`
		Tag    Ruletype `json:"tag,omitempty"`
	}

	// Ruletype is the pipeline representation of an element
	// for a ruleset block for a step in a pipeline.
	Ruletype []string

	// RuleData is the data to check our ruleset
	// against for a step in a pipeline.
	RuleData struct {
		Branch string
		Event  string
		Path   []string
		Repo   string
		Status string
		Tag    string
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

// Empty returns true if the provided ruletypes are empty.
func (r *Rules) Empty() bool {
	// return true if every ruletype is empty
	if len(r.Branch) == 0 &&
		len(r.Event) == 0 &&
		len(r.Path) == 0 &&
		len(r.Repo) == 0 &&
		len(r.Status) == 0 &&
		len(r.Tag) == 0 {
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
	// if the path ruletype is provided
	if len(from.Path) > 0 {
		// if the "or" operator is provided in the ruleset
		if strings.EqualFold(op, "or") {
			// iterate through each path in the ruletype
			for _, p := range from.Path {

				// return true if any ruletype matches the ruledata
				if r.Branch.MatchOr(from.Branch) ||
					r.Event.MatchOr(from.Event) ||
					r.Path.MatchOr(p) ||
					r.Repo.MatchOr(from.Repo) ||
					r.Status.MatchOr(from.Status) ||
					r.Tag.MatchOr(from.Tag) {
					return true
				}
			}

			// return false if no match is found
			return false
		}

		// iterate through each path in the ruletype
		for _, p := range from.Path {

			// return true if every ruletype matches the ruledata
			if r.Branch.MatchAnd(from.Branch) &&
				r.Event.MatchAnd(from.Event) &&
				r.Path.MatchAnd(p) &&
				r.Repo.MatchAnd(from.Repo) &&
				r.Status.MatchAnd(from.Status) &&
				r.Tag.MatchAnd(from.Tag) {
				return true
			}
		}

		// return false if no match is found
		return false
	}

	// if the "or" operator is provided in the ruleset
	if strings.EqualFold(op, "or") {

		// return true if any ruletype matches the ruledata
		if r.Branch.MatchOr(from.Branch) ||
			r.Event.MatchOr(from.Event) ||
			r.Path.MatchOr("") ||
			r.Repo.MatchOr(from.Repo) ||
			r.Status.MatchOr(from.Status) ||
			r.Tag.MatchOr(from.Tag) {
			return true
		}

		// return false if no match is found
		return false
	}

	// return true if every ruletype matches the ruledata
	if r.Branch.MatchAnd(from.Branch) &&
		r.Event.MatchAnd(from.Event) &&
		r.Path.MatchAnd("") &&
		r.Repo.MatchAnd(from.Repo) &&
		r.Status.MatchAnd(from.Status) &&
		r.Tag.MatchAnd(from.Tag) {
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
