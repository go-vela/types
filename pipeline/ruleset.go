// SPDX-License-Identifier: Apache-2.0

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
//
//nolint:gocyclo // accepting complexity in this case
func (r *Rules) Match(from *RuleData, matcher, op string) (bool, error) {
	// set defaults
	status := true

	var err error

	// if the path ruletype is provided
	if len(from.Path) > 0 {
		// if the "or" operator is provided in the ruleset
		if strings.EqualFold(op, constants.OperatorOr) {
			// override the default to the "or"
			if len(from.Status) != 0 {
				status, err = r.Status.MatchOr(from.Status, matcher)
				if err != nil {
					return false, err
				}
			}

			// iterate through each path in the ruletype
			for _, p := range from.Path {
				matchBranch, err := r.Branch.MatchOr(from.Branch, matcher)
				if err != nil {
					return false, err
				}

				matchComment, err := r.Comment.MatchOr(from.Comment, matcher)
				if err != nil {
					return false, err
				}

				matchEvent, err := r.Event.MatchOr(from.Event, matcher)
				if err != nil {
					return false, err
				}

				matchPath, err := r.Path.MatchOr(p, matcher)
				if err != nil {
					return false, err
				}

				matchRepo, err := r.Repo.MatchOr(from.Repo, matcher)
				if err != nil {
					return false, err
				}

				matchTag, err := r.Tag.MatchOr(from.Tag, matcher)
				if err != nil {
					return false, err
				}

				matchTarget, err := r.Target.MatchOr(from.Target, matcher)
				if err != nil {
					return false, err
				}

				// return true if any ruletype matches the ruledata
				if matchBranch || matchComment || matchEvent || matchPath || matchRepo || matchTag || matchTarget || status {
					return true, nil
				}
			}

			// return false if no match is found
			return false, nil
		}

		// override the default to the "and"
		if len(from.Status) != 0 {
			status, err = r.Status.MatchAnd(from.Status, matcher)
			if err != nil {
				return false, err
			}
		}

		// iterate through each path in the ruletype
		for _, p := range from.Path {
			matchBranch, err := r.Branch.MatchAnd(from.Branch, matcher)
			if err != nil {
				return false, err
			}

			matchComment, err := r.Comment.MatchAnd(from.Comment, matcher)
			if err != nil {
				return false, err
			}

			matchEvent, err := r.Event.MatchAnd(from.Event, matcher)
			if err != nil {
				return false, err
			}

			matchPath, err := r.Path.MatchAnd(p, matcher)
			if err != nil {
				return false, err
			}

			matchRepo, err := r.Repo.MatchAnd(from.Repo, matcher)
			if err != nil {
				return false, err
			}

			matchTag, err := r.Tag.MatchAnd(from.Tag, matcher)
			if err != nil {
				return false, err
			}

			matchTarget, err := r.Target.MatchAnd(from.Target, matcher)
			if err != nil {
				return false, err
			}

			// return true if any ruletype matches the ruledata
			if matchBranch && matchComment && matchEvent && matchPath && matchRepo && matchTag && matchTarget && status {
				return true, nil
			}
		}

		// return false if no match is found
		return false, nil
	}

	// if the "or" operator is provided in the ruleset
	if strings.EqualFold(op, constants.OperatorOr) {
		// override the default to the "or"
		if len(from.Status) != 0 {
			status, err = r.Status.MatchOr(from.Status, matcher)
			if err != nil {
				return false, nil
			}
		}

		// return true if any ruletype matches the ruledata
		matchBranch, err := r.Branch.MatchOr(from.Branch, matcher)
		if err != nil {
			return false, err
		}

		matchComment, err := r.Comment.MatchOr(from.Comment, matcher)
		if err != nil {
			return false, err
		}

		matchEvent, err := r.Event.MatchOr(from.Event, matcher)
		if err != nil {
			return false, err
		}

		matchPath, err := r.Path.MatchOr("", matcher)
		if err != nil {
			return false, err
		}

		matchRepo, err := r.Repo.MatchOr(from.Repo, matcher)
		if err != nil {
			return false, err
		}

		matchTag, err := r.Tag.MatchOr(from.Tag, matcher)
		if err != nil {
			return false, err
		}

		matchTarget, err := r.Target.MatchOr(from.Target, matcher)
		if err != nil {
			return false, err
		}

		// return true if any ruletype matches the ruledata
		if matchBranch || matchComment || matchEvent || matchPath || matchRepo || matchTag || matchTarget || status {
			return true, nil
		}

		// return false if no match is found
		return false, nil
	}

	// override the default to the "and"
	if len(from.Status) != 0 {
		status, err = r.Status.MatchAnd(from.Status, matcher)
		if err != nil {
			return false, err
		}
	}

	matchBranch, err := r.Branch.MatchAnd(from.Branch, matcher)
	if err != nil {
		return false, err
	}

	matchComment, err := r.Comment.MatchAnd(from.Comment, matcher)
	if err != nil {
		return false, err
	}

	matchEvent, err := r.Event.MatchAnd(from.Event, matcher)
	if err != nil {
		return false, err
	}

	matchPath, err := r.Path.MatchAnd("", matcher)
	if err != nil {
		return false, err
	}

	matchRepo, err := r.Repo.MatchAnd(from.Repo, matcher)
	if err != nil {
		return false, err
	}

	matchTag, err := r.Tag.MatchAnd(from.Tag, matcher)
	if err != nil {
		return false, err
	}

	matchTarget, err := r.Target.MatchAnd(from.Target, matcher)
	if err != nil {
		return false, err
	}

	// return true if any ruletype matches the ruledata
	if matchBranch && matchComment && matchEvent && matchPath && matchRepo && matchTag && matchTarget && status {
		return true, nil
	}

	// return false if no match is found
	return false, nil
}

// MatchAnd returns true when the provided ruletype
// matches the provided ruledata. When the provided
// ruletype is empty, the function returns true.
func (r *Ruletype) MatchAnd(data, matcher string) (bool, error) {
	// return true if an empty ruletype is provided
	if len(*r) == 0 {
		return true, nil
	}

	// iterate through each pattern in the ruletype
	for _, pattern := range *r {
		// handle the pattern based off the matcher provided
		switch matcher {
		case constants.MatcherRegex, "regex":
			regExpPattern, err := regexp.Compile(pattern)
			if err != nil {
				return false, err
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

// MatchOr returns true when the provided ruletype
// matches the provided ruledata. When the provided
// ruletype is empty, the function returns false.
func (r *Ruletype) MatchOr(data, matcher string) (bool, error) {
	// return false if an empty ruletype is provided
	if len(*r) == 0 {
		return false, nil
	}

	// iterate through each pattern in the ruletype
	for _, pattern := range *r {
		// handle the pattern based off the matcher provided
		switch matcher {
		case constants.MatcherRegex, "regex":
			regExpPattern, err := regexp.Compile(pattern)
			if err != nil {
				return false, err
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
