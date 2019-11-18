// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package yaml

import (
	"reflect"
	"testing"

	"github.com/go-vela/types/pipeline"
)

func TestYaml_Worker_ToPipeline(t *testing.T) {
	// setup types
	name := "worker_1"
	flavor := "16cpu8gb"
	runtime := "docker"
	want := &pipeline.Worker{
		Name:    name,
		Flavor:  flavor,
		Runtime: runtime,
	}

	v := &Worker{
		Name:    name,
		Flavor:  flavor,
		Runtime: runtime,
	}

	// run test
	got := v.ToPipeline()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToPipeline is %v, want %v", got, want)
	}
}
