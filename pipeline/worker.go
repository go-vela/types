// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

// Worker is the yaml representation of the worker block for a pipeline.
type Worker struct {
	Name    string `yaml:"name,omitempty"`
	Flavor  string `yaml:"flavor,omitempty"`
	Runtime string `yaml:"runtime,omitempty"`
}
