// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import "github.com/go-vela/types/pipeline"

// Worker is the yaml representation of a worker
// from a worker block in a pipeline.
type Worker struct {
	Flavor   string `yaml:"flavor,omitempty"   json:"flavor,omitempty" jsonschema:"minLength=1,description=Flavor identifier for worker.\nReference: https://go-vela.github.io/docs/reference/yaml/worker/#the-flavor-tag,example=large"`
	Platform string `yaml:"platform,omitempty" json:"platform,omitempty" jsonschema:"minLength=1,description=Platform identifier for the worker.\nReference: https://go-vela.github.io/docs/reference/yaml/worker/#the-platform-tag,example=kubernetes"`
}

// ToPipeline converts the Worker type
// to a pipeline Worker type.
func (w *Worker) ToPipeline() *pipeline.Worker {
	return &pipeline.Worker{
		Flavor:   w.Flavor,
		Platform: w.Platform,
	}
}
