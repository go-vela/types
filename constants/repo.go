// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package constants

// Repo get pipeline types.
const (
	// PipelineTypeYAML defines the pipeline type for allowing users
	// in Vela to control their pipeline being compiled as yaml.
	PipelineTypeYAML = "yaml"

	// PipelineTypeGo defines the pipeline type for allowing users
	// in Vela to control their pipeline being compiled as Go templates.
	PipelineTypeGo = "go"

	// PipelineTypeStarlark defines the pipeline type for allowing users
	// in Vela to control their pipeline being compiled as Starlark templates.
	PipelineTypeStarlark = "starlark"
)

// Repo ApproveForkBuild types.
const (
	// ApproveAlways defines the CI strategy of having a repo administrator approve
	// all builds triggered from a forked PR.
	ApproveAlways = "always"

	// ApproveNoWrite defines the CI strategy of having a repo administrator approve
	// all builds triggered from a forked PR where the author does not have write access.
	ApproveNoWrite = "no-write"

	// ApproveOnce defines the CI strategy of having a repo administrator approve
	// all builds triggered from an outside contributor if this is their first time contributing.
	ApproveOnce = "first-time"

	// ApproveNever defines the CI strategy of never having to approve CI builds from outside contributors.
	ApproveNever = "never"
)
