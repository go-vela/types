// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

//go:build ignore
// +build ignore

// This program will use the jsonschema struct tags
// defined on structs in /yaml/*.go files to generate
// a json schema via github.com/alecthomas/jsonschema.
//
// It can generate the majority of the schema
// but does require some 'manual' intervention to
// deal with instances where the struct definition
// is not telling the whole story.

// TODO: add test files

package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/alecthomas/jsonschema"
	"github.com/iancoleman/orderedmap"

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

	// -- Rules
	// struct says they are []string, but
	// they're basically raw.StringSlice,
	// with the addition of potential enums
	rulesEnumMap := map[string][]string{
		"branch":  {},
		"comment": {},
		"event": {
			"comment",
			"comment:created",
			"comment:edited",
			"deployment",
			"pull_request",
			"pull_request:edited",
			"pull_request:opened",
			"pull_request:synchronize",
			"push",
			"tag",
		},
		"path":   {},
		"repo":   {},
		"status": {"failure", "success"},
		"tag":    {},
		"target": {},
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

	// -- Ruleset
	// rules can currently live at ruleset level or
	// nested within 'if' (default) or 'unless'.
	// without changes the struct would only allow
	// the nested version.
	if rd, ok := s.Definitions["Ruleset"]; ok {
		// original
		rulesetProps := rd.Properties

		// create a flattened copy of ruleset
		rulesetPropsFlattened := orderedmap.New()
		for _, key := range rd.Properties.Keys() {
			if key != "if" && key != "unless" {
				if v, ok := rd.Properties.Get(key); ok {
					rulesetPropsFlattened.Set(key, v)
				}
			}
		}

		// merge in the props from rules
		// note: this causes an additional copy
		// of #/definitions/Rules in the final
		// output. using anyof/allof/oneof and
		// using a reference to Rules does not
		// seem to work as a way to "merge"
		// partially due to additionalProperties
		// being false
		if r, ok := s.Definitions["Rules"]; ok {
			for _, key := range r.Properties.Keys() {
				if v, ok := r.Properties.Get(key); ok {
					rulesetPropsFlattened.Set(key, v)
				}
			}
		}

		// clear props on the original definition
		s.Definitions["Ruleset"].Type = ""
		s.Definitions["Ruleset"].Properties = nil
		s.Definitions["Ruleset"].AdditionalProperties = []byte("")

		// at the top level for Ruleset
		s.Definitions["Ruleset"].AnyOf = []*jsonschema.Type{
			{
				Properties:           rulesetProps,
				Type:                 "object",
				AdditionalProperties: []byte("false"),
			},
			{
				Properties:           rulesetPropsFlattened,
				Type:                 "object",
				AdditionalProperties: []byte("false"),
			},
		}
	}

	// -- StepSecret
	// without changes it would only allow an array
	// of objects per the struct, but we also allow
	// an array of strings. this will allow one
	// or the other.

	// create a temp secret struct
	secret := &jsonschema.Type{
		OneOf: []*jsonschema.Type{
			{
				Type: "array",
				Items: &jsonschema.Type{
					Type:                 "string",
					AdditionalProperties: []byte("false"),
				},
			},
			{
				Type: "array",
				Items: &jsonschema.Type{
					Ref: "#/definitions/StepSecret",
				},
			},
		},
	}

	// replace "secrets" field for Step
	if step, ok := s.Definitions["Step"]; ok {
		if stepSecrets, ok := step.Properties.Get("secrets"); ok {
			stepSecret := stepSecrets.(*jsonschema.Type)

			// corry over current description
			secret.Description = stepSecret.Description
			step.Properties.Set("secrets", secret)
		}
	}

	// replace "secrets" field for Origin
	if origin, ok := s.Definitions["Origin"]; ok {
		if originSecrets, ok := origin.Properties.Get("secrets"); ok {
			// create a copy of the secret
			originSecretNew := &jsonschema.Type{}
			*originSecretNew = *secret

			originSecret := originSecrets.(*jsonschema.Type)

			// carry over current description
			originSecretNew.Description = originSecret.Description
			origin.Properties.Set("secrets", originSecretNew)
		}
	}

	// -- Stages
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

	// -- Ulimit
	// without changes it would only allow an object
	// per the struct, but we do some special handling
	// to allow specially formatted strings
	if ud, ok := s.Definitions["Ulimit"]; ok {
		ulimitProps := ud.Properties
		ulimitReq := ud.Required
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
				Properties:           ulimitProps,
				Required:             ulimitReq,
				AdditionalProperties: []byte("false"),
			},
		}
	}

	// -- Volume
	// without changes it would only allow an object
	// per the struct, but we do some special handling
	// to allow specially formatted strings
	if vd, ok := s.Definitions["Volume"]; ok {
		volProps := vd.Properties
		volReq := vd.Required
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
				Properties:           volProps,
				Required:             volReq,
				AdditionalProperties: []byte("false"),
			},
		}
	}

	// TODO: handle more granular step requirement, ie
	// https://github.com/go-vela/compiler/blob/ac17e426a4a62bac3ef8dd2dc587bafa374957c7/compiler/native/validate.go#L107-L111

	// TODO: better secret validation with regards to
	// expected 'key' formats depending on which 'type' is in use

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
// the both types.
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
