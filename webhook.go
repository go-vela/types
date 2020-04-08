// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package types

import "github.com/go-vela/types/library"

// Webhook defines a struct that is used to return
// the required data when processing webhook event
// a for a source provider event.
type Webhook struct {
	Comment  string
	PRNumber int
	Hook     *library.Hook
	Repo     *library.Repo
	Build    *library.Build
}
