// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package types

import (
	"github.com/go-vela/types/library"
	"github.com/go-vela/types/pipeline"
)

// Item is the queue representation of an item to publish to the queue.
type Item struct {
	Build    *library.Build  `json:"build"`
	Pipeline *pipeline.Build `json:"pipeline"`
	Repo     *library.Repo   `json:"repo"`
	User     *library.User   `json:"user"`
}

// ToItem creates a queue item from a pipeline, build, repo and user.
func ToItem(p *pipeline.Build, b *library.Build, r *library.Repo, u *library.User) *Item {
	return &Item{
		Pipeline: p,
		Build:    b,
		Repo:     r,
		User:     u,
	}
}
