// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package raw

import (
	"encoding/json"
	"errors"
	"strings"
)

// StringSliceMap represents an array of strings or a map of strings.
type StringSliceMap map[string]string

// UnmarshalJSON implements the Unmarshaler interface for the StringSlice type.
func (s *StringSliceMap) UnmarshalJSON(b []byte) error {
	// return nil if no input is provided
	if len(b) == 0 {
		return nil
	}

	// target map we want to return
	targetMap := map[string]string{}

	// json slice we try unmarshaling to
	jsonSlice := StringSlice{}

	// attempt to unmarshal as a string slice type
	err := json.Unmarshal(b, &jsonSlice)
	if err == nil {
		// iterate through each element in the json slice
		for _, v := range jsonSlice {
			// split each slice element into key/value pairs
			kvPair := strings.SplitN(v, "=", 2)

			// append each key/value pair to our target map
			targetMap[kvPair[0]] = kvPair[1]
		}

		// overwrite existing StringSliceMap
		*s = targetMap

		return nil
	}

	// json map we try unmarshaling to
	jsonMap := map[string]string{}

	// attempt to unmarshal as map of strings
	err = json.Unmarshal(b, &jsonMap)
	if err == nil {
		// iterate through each item in the json map
		for k, v := range jsonMap {
			// append each key/value pair to our target map
			targetMap[k] = v
		}

		// overwrite existing StringSliceMap
		*s = targetMap

		return nil
	}

	return errors.New("unable to unmarshal into StringSliceMap")
}

// UnmarshalYAML implements the Unmarshaler interface for the StringSliceMap type.
func (s *StringSliceMap) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// target map we want to return
	targetMap := map[string]string{}

	// yaml slice we try unmarshaling to
	yamlSlice := StringSlice{}

	// attempt to unmarshal as a string slice type
	err := unmarshal(&yamlSlice)
	if err == nil {
		// iterate through each element in the yaml slice
		for _, v := range yamlSlice {
			// split each slice element into key/value pairs
			kvPair := strings.SplitN(v, "=", 2)

			// append each key/value pair to our target map
			targetMap[kvPair[0]] = kvPair[1]
		}

		// overwrite existing StringSliceMap
		*s = targetMap

		return nil
	}

	// yaml map we try unmarshaling to
	yamlMap := map[string]string{}

	// attempt to unmarshal as map of strings
	err = unmarshal(&yamlMap)
	if err == nil {
		// iterate through each item in the yaml map
		for k, v := range yamlMap {
			// append each key/value pair to our target map
			targetMap[k] = v
		}

		// overwrite existing StringSliceMap
		*s = targetMap

		return nil
	}

	return errors.New("unable to unmarshal into StringSliceMap")
}
