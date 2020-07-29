// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"
)

// nolint:lll // jsonschema will cause long lines
type (
	// Ruleset is the yaml representation of a
	// ruleset block for a step in a pipeline.
	Ruleset struct {
		If       Rules  `yaml:"if,omitempty"       jsonschema:"description=Limit execution to when all rules match.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/ruleset/#fields-1"`
		Unless   Rules  `yaml:"unless,omitempty"   jsonschema:"description=Limit execution to when all rules do not match.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/ruleset/#fields-1"`
		Matcher  string `yaml:"matcher,omitempty"  jsonschema:"enum=filepath,enum=regexp,default=filepath,description=Use the defined matching method.\nReference: coming soon"`
		Operator string `yaml:"operator,omitempty" jsonschema:"enum=or,enum=and,default=and,description=Whether all rule conditions must be met or just any one of them.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/ruleset/#fields-1"`
		Continue bool   `yaml:"continue,omitempty" jsonschema:"default=false,description=Limits the execution of a step to continuing on any failure.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/ruleset/#fields-1"`
	}

	// Rules is the yaml representation of the ruletypes
	// from a ruleset block for a step in a pipeline.
	Rules struct {
		Branch  []string `yaml:"branch,omitempty"  jsonschema:"description=Limits the execution of a step to matching build branches.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/ruleset/#fields"`
		Comment []string `yaml:"comment,omitempty" jsonschema:"description=Limits the execution of a step to matching a pull request comment.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/ruleset/#fields"`
		Event   []string `yaml:"event,omitempty"   jsonschema:"description=Limits the execution of a step to matching build events.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/ruleset/#fields"`
		Path    []string `yaml:"path,omitempty"    jsonschema:"description=Limits the execution of a step to matching files changed in a repository.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/ruleset/#fields"`
		Repo    []string `yaml:"repo,omitempty"    jsonschema:"description=Limits the execution of a step to matching repos.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/ruleset/#fields"`
		Status  []string `yaml:"status,omitempty"  jsonschema:"description=Limits the execution of a step to matching build statuses.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/ruleset/#fields"`
		Tag     []string `yaml:"tag,omitempty"     jsonschema:"description=Limits the execution of a step to matching build tag references.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/ruleset/#fields"`
		Target  []string `yaml:"target,omitempty"  jsonschema:"description=Limits the execution of a step to matching build deployment targets.\nReference: https://go-vela.github.io/docs/concepts/pipeline/steps/ruleset/#fields"`
	}
)

// ToPipeline converts the Ruleset type
// to a pipeline Ruleset type.
func (r *Ruleset) ToPipeline() *pipeline.Ruleset {
	return &pipeline.Ruleset{
		If:       *r.If.ToPipeline(),
		Unless:   *r.Unless.ToPipeline(),
		Matcher:  r.Matcher,
		Operator: r.Operator,
		Continue: r.Continue,
	}
}

// UnmarshalYAML implements the Unmarshaler interface for the Ruleset type.
func (r *Ruleset) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// simple struct we try unmarshaling to
	simple := new(Rules)

	// advanced struct we try unmarshaling to
	advanced := new(struct {
		If       Rules
		Unless   Rules
		Matcher  string
		Operator string
		Continue bool
	})

	// attempt to unmarshal simple ruleset
	unmarshal(simple)
	// attempt to unmarshal advanced ruleset
	unmarshal(advanced)

	// set ruleset `unless` to advanced `unless` rules
	r.Unless = advanced.Unless
	// set ruleset `matcher` to advanced `matcher`
	r.Matcher = advanced.Matcher
	// set ruleset `operator` to advanced `operator`
	r.Operator = advanced.Operator
	// set ruleset `continue` to advanced `continue`
	r.Continue = advanced.Continue

	// implicitly add simple ruleset to the advanced ruleset for each rule type
	advanced.If.Branch = append(advanced.If.Branch, simple.Branch...)
	advanced.If.Comment = append(advanced.If.Comment, simple.Comment...)
	advanced.If.Event = append(advanced.If.Event, simple.Event...)
	advanced.If.Path = append(advanced.If.Path, simple.Path...)
	advanced.If.Repo = append(advanced.If.Repo, simple.Repo...)
	advanced.If.Status = append(advanced.If.Status, simple.Status...)
	advanced.If.Tag = append(advanced.If.Tag, simple.Tag...)
	advanced.If.Target = append(advanced.If.Target, simple.Target...)

	// set ruleset `if` to advanced `if` rules
	r.If = advanced.If

	// implicitly set `matcher` field if empty for ruleset
	if len(r.Matcher) == 0 {
		r.Matcher = constants.MatcherFilepath
	}

	// implicitly set `operator` field if empty for ruleset
	if len(r.Operator) == 0 {
		r.Operator = constants.OperatorAnd
	}

	return nil
}

// ToPipeline converts the Rules
// type to a pipeline Rules type.
func (r *Rules) ToPipeline() *pipeline.Rules {
	return &pipeline.Rules{
		Branch:  r.Branch,
		Comment: r.Comment,
		Event:   r.Event,
		Path:    r.Path,
		Repo:    r.Repo,
		Status:  r.Status,
		Tag:     r.Tag,
		Target:  r.Target,
	}
}

// UnmarshalYAML implements the Unmarshaler interface for the Rules type.
func (r *Rules) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// rules struct we try unmarshaling to
	rules := new(struct {
		Branch  raw.StringSlice
		Comment raw.StringSlice
		Event   raw.StringSlice
		Path    raw.StringSlice
		Repo    raw.StringSlice
		Status  raw.StringSlice
		Tag     raw.StringSlice
		Target  raw.StringSlice
	})

	// attempt to unmarshal rules
	err := unmarshal(rules)
	if err == nil {
		r.Branch = []string(rules.Branch)
		r.Comment = []string(rules.Comment)
		r.Event = []string(rules.Event)
		r.Path = []string(rules.Path)
		r.Repo = []string(rules.Repo)
		r.Status = []string(rules.Status)
		r.Tag = []string(rules.Tag)
		r.Target = []string(rules.Target)
	}

	return err
}
