// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package types

import (
	"github.com/go-vela/types/library"
	"github.com/go-vela/types/pipeline"
)

// ItemVersion allows the worker to detect items that were queued before an Vela server
// upgrade or downgrade, so it can handle such stale data gracefully.
// For example, the worker could fail a stale build or ask the server to recompile it.
// This is not a public API and is unrelated to the version key in pipeline yaml.
const ItemVersion uint64 = 1

// Item is the queue representation of an item to publish to the queue.
type Item struct {
	Build    *library.Build  `json:"build"`
	Pipeline *pipeline.Build `json:"pipeline"`
	Repo     *library.Repo   `json:"repo"`
	User     *library.User   `json:"user"`
	// The 0-value is the implicit ItemVersion for queued Items that pre-date adding the field.
	ItemVersion uint64 `json:"item_version"`
}

// ToItem creates a queue item from a pipeline, build, repo and user.
func ToItem(p *pipeline.Build, b *library.Build, r *library.Repo, u *library.User) *Item {
	return &Item{
		Pipeline:    p,
		Build:       b,
		Repo:        r,
		User:        u,
		ItemVersion: ItemVersion,
	}
}
