// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package types

import (
	"github.com/go-vela/types/library"
	"github.com/go-vela/types/pipeline"
)

// BuildPackage is the library representation of a build for a pipeline.
//
// swagger:model BuildPackage
type BuildPackage struct {
	Build    *library.Build    `json:"build,omitempty"`
	Secrets  []*library.Secret `json:"secrets,omitempty"`
	Pipeline *pipeline.Build   `json:"pipeline,omitempty"`
	Repo     *library.Repo     `json:"repo,omitempty"`
	User     *library.User     `json:"user,omitempty"`
	Token    string            `json:"token,omitempty"`
}

func (b *BuildPackage) WithBuild(build *library.Build) *BuildPackage {
	if build != nil {
		b.Build = build
	}

	return b
}

func (b *BuildPackage) WithPipeline(pipeline *pipeline.Build) *BuildPackage {
	if pipeline != nil {
		b.Pipeline = pipeline
	}

	return b
}

func (b *BuildPackage) WithRepo(repo *library.Repo) *BuildPackage {
	if repo != nil {
		b.Repo = repo
	}

	return b
}

func (b *BuildPackage) WithUser(user *library.User) *BuildPackage {
	if user != nil {
		b.User = user
	}

	return b
}

func (b *BuildPackage) WithToken(token string) *BuildPackage {
	if len(token) != 0 {
		b.Token = token
	}

	return b
}
