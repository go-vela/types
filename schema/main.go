// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

//+build ignore

package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/alecthomas/jsonschema"

	"github.com/go-vela/types/raw"
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
		TypeMapper: func(i reflect.Type) *jsonschema.Type {
			// array of strings or a map of strings
			if i == reflect.TypeOf(raw.StringSliceMap{}) {
				return &jsonschema.Type{
					OneOf: []*jsonschema.Type{
						{
							Type: "object",
							PatternProperties: map[string]*jsonschema.Type{
								".*": {Type: "string"},
							},
							AdditionalProperties: []byte("false"),
						},
						{
							Type: "array",
							Items: &jsonschema.Type{
								Type: "string",
							},
						},
					},
				}
			}
			// string or slice of strings
			if i == reflect.TypeOf(raw.StringSlice{}) {
				return &jsonschema.Type{
					OneOf: []*jsonschema.Type{
						{
							Type: "string",
						},
						{
							Type: "array",
							Items: &jsonschema.Type{
								Type: "string",
							},
						},
					},
				}
			}
			return nil
		},
	}
	s := r.Reflect(&yaml.Build{})

	// fix for when yaml anchors aren't expanded by yaml validator
	// TODO: do we need this?
	s.AdditionalProperties = []byte("true")

	// attach enums to Rules props (they're not raw.StringSlice)
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

	// fix ulimit on Step since it accepts an object
	// or specially formatted string
	ulimit := s.Definitions["Ulimit"].Properties
	ulimitReq := s.Definitions["Ulimit"].Required
	s.Definitions["Ulimit"].Type = ""
	s.Definitions["Ulimit"].Properties = nil
	s.Definitions["Ulimit"].AdditionalProperties = []byte("")
	s.Definitions["Ulimit"].Required = []string{}
	s.Definitions["Ulimit"].OneOf = []*jsonschema.Type{
		{
			Type:                 "string",
			Pattern:              "[a-z]+=[0-9]+:[0-9]+",
			AdditionalProperties: []byte("false"),
		},
		{
			Type:                 "object",
			Properties:           ulimit,
			Required:             ulimitReq,
			AdditionalProperties: []byte("false"),
		},
	}

	// fix volume on Step since it accepts an object
	// or a specially formatted string
	vol := s.Definitions["Volume"].Properties
	volReq := s.Definitions["Volume"].Required
	s.Definitions["Volume"].Type = ""
	s.Definitions["Volume"].Properties = nil
	s.Definitions["Volume"].AdditionalProperties = []byte("")
	s.Definitions["Volume"].Required = []string{}
	s.Definitions["Volume"].OneOf = []*jsonschema.Type{
		{
			Type:                 "string",
			Pattern:              "[a-z\\/]+:[a-z\\/]+:[row]+",
			AdditionalProperties: []byte("false"),
		},
		{
			Type:                 "object",
			Properties:           vol,
			Required:             volReq,
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
