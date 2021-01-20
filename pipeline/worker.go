// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

// Worker is the yaml representation of the worker block for a pipeline.
//
// swagger:model PipelineWorker
type Worker struct {
	Flavor   string `json:"flavor,omitempty"   yaml:"flavor,omitempty"`
	Platform string `json:"platform,omitempty" yaml:"platform,omitempty"`
}

// Empty returns true if the provided worker is empty.
func (w *Worker) Empty() bool {
	// return true if every worker field is empty
	if len(w.Flavor) == 0 &&
		len(w.Platform) == 0 {
		return true
	}

	// return false if any of the worker fields are provided
	return false
}
