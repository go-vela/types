// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import "github.com/go-vela/types/pipeline"

// Worker is the yaml representation of a worker
// from a worker block in a pipeline.
type Worker struct {
	Flavor   string `yaml:"flavor,omitempty"`
	Platform string `yaml:"platform,omitempty"`
}

// ToPipeline converts the Worker type
// to a pipeline Worker type.
func (w *Worker) ToPipeline() *pipeline.Worker {
	return &pipeline.Worker{
		Flavor:   w.Flavor,
		Platform: w.Platform,
	}
}
