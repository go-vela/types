// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

//+build ignore

package main

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/jsonschema"

	"github.com/go-vela/types/yaml"
)

type EnumMap map[string][]string

func oneOfWithEnum(enum []string) []*jsonschema.Type {
	s := []*jsonschema.Type{}

	str := jsonschema.Type{}
	if len(enum) > 0 {
		str.Enum = sliceStrToSliceInterface(enum)
	} else {
		str.Type = "string"
	}

	strSlice := jsonschema.Type{}
	strSlice.Type = "array"
	strSlice.Items = &str

	s = append(s, &str)
	s = append(s, &strSlice)

	return s
}

func main() {
	r := jsonschema.Reflector{
		ExpandedStruct:             true,
		RequiredFromJSONSchemaTags: true,
	}
	s := r.Reflect(&yaml.Build{})

	// fix for when yaml anchors aren't expanded by yaml validator
	// TODO: do we need this?
	s.AdditionalProperties = []byte("true")

	// fix string or string slice
	rulesetMap := EnumMap{
		"branch":  []string{},
		"comment": []string{},
		"event":   []string{"push", "pull_request", "tag", "deployment", "comment"},
		"path":    []string{},
		"repo":    []string{},
		"status":  []string{"failure", "success"},
		"tag":     []string{},
		"target":  []string{},
	}

	for k, v := range rulesetMap {
		if p, ok := s.Definitions["Rules"].Properties.Get(k); ok {
			p.(*jsonschema.Type).Type = ""
			p.(*jsonschema.Type).Items = nil
			p.(*jsonschema.Type).OneOf = oneOfWithEnum(v)

			s.Definitions["Rules"].Properties.Set(k, p)
		}
	}

	// fix "needs"
	if needs, ok := s.Definitions["Stage"].Properties.Get("needs"); ok {
		needs.(*jsonschema.Type).Type = ""
		needs.(*jsonschema.Type).Items = nil
		needs.(*jsonschema.Type).OneOf = oneOfWithEnum([]string{})

		s.Definitions["Stage"].Properties.Set("needs", needs)
	}

	// stages fix
	if stages, ok := s.Properties.Get("stages"); ok {
		stages.(*jsonschema.Type).Type = "object"
		stages.(*jsonschema.Type).Items = nil
		stages.(*jsonschema.Type).PatternProperties = make(map[string]*jsonschema.Type)
		stages.(*jsonschema.Type).PatternProperties[".*"] = &jsonschema.Type{
			Version: "http://json-schema.org/draft-04/schema#",
			Ref:     "#/definitions/Stage",
		}
		stages.(*jsonschema.Type).AdditionalProperties = []byte("false")

		s.Properties.Set("stages", stages)
	}

	// stages, steps mutual exclusiveness
	// https://stackoverflow.com/questions/28162509/mutually-exclusive-property-groups
	s.OneOf = []*jsonschema.Type{
		{
			Required: []string{"steps"},
			Not: &jsonschema.Type{
				Required: []string{"stages"},
			},
		},
		{
			AllOf: []*jsonschema.Type{
				{
					Not: &jsonschema.Type{
						Required: []string{"steps"},
					},
				},
			},
		},
	}

	// fix ruleset
	currRulesetProperties := s.Definitions["Ruleset"].Properties
	s.Definitions["Ruleset"].Properties = nil
	s.Definitions["Ruleset"].Type = ""
	s.Definitions["Ruleset"].AdditionalProperties = []byte("")
	s.Definitions["Ruleset"].OneOf = []*jsonschema.Type{
		{
			Ref: "#/definitions/Rules",
		},
		{
			Properties:           currRulesetProperties,
			Type:                 "object",
			AdditionalProperties: []byte("false"),
		},
	}

	// transform to json and format
	j, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}

	// return json
	fmt.Printf("%s\n", j)
}

func sliceStrToSliceInterface(s []string) []interface{} {
	b := make([]interface{}, len(s))
	for i := range s {
		b[i] = s[i]
	}
	return b
}
