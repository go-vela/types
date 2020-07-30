// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// +build ignore

// This program will use the jsonschema struct tags
// defined on structs in /yaml/*.go files to generate
// a json schema via github.com/alecthomas/jsonschema.
//
// It can generate the majority of the schema
// but does require some 'manual' intervention to
// deal with instances where the struct definition
// is not telling the whole story.

package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/alecthomas/jsonschema"

	"github.com/go-vela/types/raw"
	"github.com/go-vela/types/yaml"
)

func main() {
	// setup the reflector for schema translation
	r := jsonschema.Reflector{
		ExpandedStruct:             true,
		RequiredFromJSONSchemaTags: true,
		// custom mapping of some Vela types to schema types
		TypeMapper: func(i reflect.Type) *jsonschema.Type {
			// handle array of strings or a map of strings
			if i == reflect.TypeOf(raw.StringSliceMap{}) {
				return &jsonschema.Type{
					OneOf: []*jsonschema.Type{
						{
							Type: "object",
							PatternProperties: map[string]*jsonschema.Type{
								".*": {
									Type: "string",
								},
							},
							AdditionalProperties: []byte("false"),
						},
						{
							Type: "array",
							Items: &jsonschema.Type{
								Type: "string",
							},
							AdditionalProperties: []byte("false"),
						},
					},
				}
			}
			// handle string or slice of strings
			if i == reflect.TypeOf(raw.StringSlice{}) {
				return &jsonschema.Type{
					OneOf: stringSliceOneOf([]string{}),
				}
			}
			return nil
		},
	}
	// build the schema with what is provided
	// by the struct tags and the typemapper
	s := r.Reflect(&yaml.Build{})

	// manual overrides

	// set the title of the schema
	s.Title = "Vela Pipeline Configuration"

	// defaults to false, less issues with
	// anchors when unset at top level
	s.AdditionalProperties = []byte("")

	// fix Rules
	// struct says they are []string, but
	// they're basically raw.StringSlice,
	// with the addition of potential enums
	rulesEnumMap := map[string][]string{
		"branch":  {},
		"comment": {},
		"event":   {"push", "pull_request", "tag", "deployment", "comment"},
		"path":    {},
		"repo":    {},
		"status":  {"failure", "success"},
		"tag":     {},
		"target":  {},
	}

	// iterate over the enum map to configure
	// each prop for the Rules definition
	if rules, ok := s.Definitions["Rules"]; ok {
		for k, v := range rulesEnumMap {
			if p, ok := rules.Properties.Get(k); ok {
				pj := p.(*jsonschema.Type)
				pj.Type = ""
				pj.Items = nil
				// should be either string or slice of string
				// with enums, if provided
				pj.OneOf = stringSliceOneOf(v)

				// set the above assigned schema Type to the current prop
				rules.Properties.Set(k, pj)
			}
		}
	}

	// fix stages
	// they claim to be a slice of stage, but
	// they're really a map/object
	if stages, ok := s.Properties.Get("stages"); ok {
		sj := stages.(*jsonschema.Type)
		sj.Type = "object"
		sj.Items = nil
		// make it an object. keys can be arbitrary names.
		sj.PatternProperties = map[string]*jsonschema.Type{
			".*": {
				Ref: "#/definitions/Stage",
			},
		}
		sj.AdditionalProperties = []byte("false")

		// set the above assigned schema Type to the current prop
		s.Properties.Set("stages", sj)
	}

	// fix ruleset
	// rules can currently live at ruleset level or
	// nested within 'if' (default) or 'unless'.
	// without changes the struct would only allow
	// the nested version.
	if _, ok := s.Definitions["Ruleset"]; ok {
		currRulesetProperties := s.Definitions["Ruleset"].Properties
		s.Definitions["Ruleset"].Type = ""
		s.Definitions["Ruleset"].Properties = nil
		s.Definitions["Ruleset"].AdditionalProperties = []byte("")

		// at the top level for Ruleset
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
	}

	// fix ulimit
	// without changes it would only allow an object
	// per the struct, but we do some special handling
	// to allow specially formatted strings
	if _, ok := s.Definitions["Ulimit"]; ok {
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
	}

	// fix volume
	// without changes it would only allow an object
	// per the struct, but we do some special handling
	// to allow specially formatted strings
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

	// transform to json and format
	j, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}

	// return json
	fmt.Printf("%s\n", j)
}

// stringSliceOneOf is a helper to create an slice of
// schema types that represent a string or a slice of
// string. optionally provide an enum list to set on
// the
func stringSliceOneOf(enum []string) []*jsonschema.Type {
	s := []*jsonschema.Type{}

	str := jsonschema.Type{}
	str.Type = "string"
	if len(enum) > 0 {
		str.Enum = sliceStringToSliceInterface(enum)
	}

	strSlice := jsonschema.Type{}
	strSlice.Type = "array"
	strSlice.Items = &str

	s = append(s, &str)
	s = append(s, &strSlice)

	return s
}

// sliceStrToSliceInterface is a helper to turn
// a slice of string to a slice of interface
func sliceStringToSliceInterface(s []string) []interface{} {
	b := make([]interface{}, len(s))
	for i := range s {
		b[i] = s[i]
	}
	return b
}
