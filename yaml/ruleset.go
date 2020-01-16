// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"
)

type (
	// Ruleset is the yaml representation of a
	// ruleset block for a step in a pipeline.
	Ruleset struct {
		If       Rules  `yaml:"if,omitempty"`
		Unless   Rules  `yaml:"unless,omitempty"`
		Operator string `yaml:"operator,omitempty"`
		Continue bool   `yaml:"continue,omitempty"`
	}

	// Rules is the yaml representation of the ruletypes
	// from a ruleset block for a step in a pipeline.
	Rules struct {
		Branch []string `yaml:"branch,omitempty"`
		Event  []string `yaml:"event,omitempty"`
		Path   []string `yaml:"path,omitempty"`
		Repo   []string `yaml:"repo,omitempty"`
		Status []string `yaml:"status,omitempty"`
		Tag    []string `yaml:"tag,omitempty"`
	}
)

// ToPipeline converts the Ruleset type
// to a pipeline Ruleset type.
func (r *Ruleset) ToPipeline() *pipeline.Ruleset {
	return &pipeline.Ruleset{
		If:       *r.If.ToPipeline(),
		Unless:   *r.Unless.ToPipeline(),
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
		Operator string
		Continue bool
	})

	// attempt to unmarshal simple ruleset
	unmarshal(simple)
	// attempt to unmarshal advanced ruleset
	unmarshal(advanced)

	// set ruleset `unless` to advanced `unless` rules
	r.Unless = advanced.Unless
	// set ruleset `operator` to advanced `operator`
	r.Operator = advanced.Operator
	// set ruleset `continue` to advanced `continue`
	r.Continue = advanced.Continue

	// implicitly add simple ruleset to the advanced ruleset for each rule type
	advanced.If.Branch = append(advanced.If.Branch, simple.Branch...)
	advanced.If.Event = append(advanced.If.Event, simple.Event...)
	advanced.If.Path = append(advanced.If.Path, simple.Path...)
	advanced.If.Repo = append(advanced.If.Repo, simple.Repo...)
	advanced.If.Status = append(advanced.If.Status, simple.Status...)
	advanced.If.Tag = append(advanced.If.Tag, simple.Tag...)

	// set ruleset `if` to advanced `if` rules
	r.If = advanced.If

	// implicitly set `operator` field if empty for ruleset
	if len(r.Operator) == 0 {
		r.Operator = "and"
	}

	return nil
}

// ToPipeline converts the Rules
// type to a pipeline Rules type.
func (r *Rules) ToPipeline() *pipeline.Rules {
	return &pipeline.Rules{
		Branch: r.Branch,
		Event:  r.Event,
		Path:   r.Path,
		Repo:   r.Repo,
		Status: r.Status,
		Tag:    r.Tag,
	}
}

// UnmarshalYAML implements the Unmarshaler interface for the Rules type.
func (r *Rules) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// rules struct we try unmarshaling to
	rules := new(struct {
		Branch raw.StringSlice
		Event  raw.StringSlice
		Path   raw.StringSlice
		Repo   raw.StringSlice
		Status raw.StringSlice
		Tag    raw.StringSlice
	})

	// attempt to unmarshal rules
	err := unmarshal(rules)
	if err == nil {
		r.Branch = []string(rules.Branch)
		r.Event = []string(rules.Event)
		r.Path = []string(rules.Path)
		r.Repo = []string(rules.Repo)
		r.Status = []string(rules.Status)
		r.Tag = []string(rules.Tag)
	}

	return err
}
