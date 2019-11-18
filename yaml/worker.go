// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import "github.com/go-vela/types/pipeline"

// Worker is the yaml representation of a worker
// from a worker block in a pipeline.
type Worker struct {
	Name    string `yaml:"name,omitempty"`
	Flavor  string `yaml:"flavor,omitempty"`
	Runtime string `yaml:"runtime,omitempty"`
}

// ToPipeline converts the Metadata type
// to a pipeline Metadata type.
func (w *Worker) ToPipeline() *pipeline.Worker {
	return &pipeline.Worker{
		Name:    w.Name,
		Flavor:  w.Flavor,
		Runtime: w.Runtime,
	}
}
