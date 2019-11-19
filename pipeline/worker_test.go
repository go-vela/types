// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import "testing"

func TestPipeline_Worker_Empty(t *testing.T) {
	// setup types
	w := Worker{}

	// run test
	got := w.Empty()

	if !got {
		t.Errorf("Worker IsEmpty is %v, want true", got)
	}
}
