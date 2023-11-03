// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"fmt"
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
		Parallel bool     `json:"-"                 yaml:"-"`
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
		Parallel bool     `json:"-"                 yaml:"-"`
	}
)

// Match returns true when the provided ruledata matches
// the if rules and does not match any of the unless rules.
// When the provided if rules are empty, the function returns
// true. When both the provided if and unless rules are empty,
// the function also returns true.
func (r *Ruleset) Match(from *RuleData) (bool, error) {
	// return true when the if and unless rules are empty
	if r.If.Empty() && r.Unless.Empty() {
		return true, nil
	}

	// return false when the unless rules are not empty and match
	if !r.Unless.Empty() {
		match, err := r.Unless.Match(from, r.Matcher, r.Operator)
		if err != nil {
			return false, err
		}

		if match {
			return false, nil
		}
	}

	// return true when the if rules are empty
	if r.If.Empty() {
		return true, nil
	}

	// return true when the if rules match
	match, err := r.If.Match(from, r.Matcher, r.Operator)

	return match, err
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
func (r *Rules) Match(from *RuleData, matcher, op string) (bool, error) {
	// if the path ruletype is provided
	if len(from.Path) > 0 {
		// if the "or" operator is provided in the ruleset
		if strings.EqualFold(op, constants.OperatorOr) {
			// iterate through each path in the ruletype
			for _, p := range from.Path {
				matches, err := matches(r, from, matcher, p, constants.OperatorOr)
				if err != nil {
					return false, err
				}

				// return true if any ruletype matches the ruledata
				if matches {
					return true, nil
				}
			}

			// return false if no match is found
			return false, nil
		}

		// iterate through each path in the ruletype
		for _, p := range from.Path {
			matches, err := matches(r, from, matcher, p, constants.OperatorAnd)
			if err != nil {
				return false, err
			}

			// return true if any ruletype matches the ruledata
			if matches {
				return true, nil
			}
		}

		// return false if no match is found
		return false, nil
	}

	// if the "or" operator is provided in the ruleset
	if strings.EqualFold(op, constants.OperatorOr) {
		// return true if any ruletype matches the ruledata
		return matches(r, from, matcher, "", constants.OperatorOr)
	}

	return matches(r, from, matcher, "", constants.OperatorAnd)
}

// Match returns true when the provided ruletype
// matches the provided ruledata. When the provided
// ruletype is empty, the function returns true for
// the `and` operator and false for the `or` operator.
func (r *Ruletype) Match(data, matcher, logic string) (bool, error) {
	// return true for `and`, false for `or` if an empty ruletype is provided
	if len(*r) == 0 {
		return strings.EqualFold(logic, constants.OperatorAnd), nil
	}

	// iterate through each pattern in the ruletype
	for _, pattern := range *r {
		// handle the pattern based off the matcher provided
		switch matcher {
		case constants.MatcherRegex, "regex":
			regExpPattern, err := regexp.Compile(pattern)
			if err != nil {
				return false, fmt.Errorf("error in regex pattern %s: %w", pattern, err)
			}

			// return true if the regexp pattern matches the ruledata
			if regExpPattern.MatchString(data) {
				return true, nil
			}
		case constants.MatcherFilepath:
			fallthrough
		default:
			// return true if the pattern matches the ruledata
			ok, _ := filepath.Match(pattern, data)
			if ok {
				return true, nil
			}
		}
	}

	// return false if no match is found
	return false, nil
}

// matches is a helper function which leverages the Match method for all rules
// and returns `true` if the ruleset is indeed a match.
func matches(r *Rules, from *RuleData, matcher, path, logic string) (bool, error) {
	status := true

	var err error

	if len(from.Status) != 0 {
		status, err = r.Status.Match(from.Status, matcher, logic)
		if err != nil {
			return false, err
		}
	}

	matchBranch, err := r.Branch.Match(from.Branch, matcher, logic)
	if err != nil {
		return false, err
	}

	matchComment, err := r.Comment.Match(from.Comment, matcher, logic)
	if err != nil {
		return false, err
	}

	matchEvent, err := r.Event.Match(from.Event, matcher, logic)
	if err != nil {
		return false, err
	}

	matchPath, err := r.Path.Match(path, matcher, logic)
	if err != nil {
		return false, err
	}

	matchRepo, err := r.Repo.Match(from.Repo, matcher, logic)
	if err != nil {
		return false, err
	}

	matchTag, err := r.Tag.Match(from.Tag, matcher, logic)
	if err != nil {
		return false, err
	}

	matchTarget, err := r.Target.Match(from.Target, matcher, logic)
	if err != nil {
		return false, err
	}

	switch logic {
	case constants.OperatorAnd:
		return (matchBranch && matchComment && matchEvent && matchPath && matchRepo && matchTag && matchTarget && status), nil
	default:
		return (matchBranch || matchComment || matchEvent || matchPath || matchRepo || matchTag || matchTarget || status), nil
	}
}
