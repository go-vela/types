// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/go-vela/types/library"
)

// ItemVersion allows the worker to detect items that were queued before an Vela server
// upgrade or downgrade, so it can handle such stale data gracefully.
// For example, the worker could fail a stale build or ask the server to recompile it.
// This is not a public API and is unrelated to the version key in pipeline yaml.
const ItemVersion uint64 = 2

// Item is the queue representation of an item to publish to the queue.
//
// Deprecated: use Item from github.com/go-vela/server/queue/models instead.
type Item struct {
	Build *library.Build `json:"build"`
	Repo  *library.Repo  `json:"repo"`
	User  *library.User  `json:"user"`
	// The 0-value is the implicit ItemVersion for queued Items that pre-date adding the field.
	ItemVersion uint64 `json:"item_version"`
}

// ToItem creates a queue item from a build, repo and user.
func ToItem(b *library.Build, r *library.Repo, u *library.User) *Item {
	return &Item{
		Build:       b,
		Repo:        r,
		User:        u,
		ItemVersion: ItemVersion,
	}
}
