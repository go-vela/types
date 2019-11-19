// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

// Worker is the yaml representation of the worker block for a pipeline.
type Worker struct {
	Flavor   string `yaml:"flavor,omitempty"`
	Platform string `yaml:"platform,omitempty"`
}

// Empty returns true if the provided ruletypes are empty.
func (w *Worker) Empty() bool {
	// return true if every ruletype is empty
	if len(w.Flavor) == 0 &&
		len(w.Platform) == 0 {
		return true
	}

	// return false if any of the ruletype is provided
	return false
}
