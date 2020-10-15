// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package version

import (
	"fmt"
)

const versionFormat = `{
  Full: %s,
  Major: %d,
  Minor: %d,
  Patch: %d,
  PreRelease: %s,
  Metadata: {
    Architecture: %s,
    BuildDate: %s,
    Compiler: %s,
    GitCommit: %s,
    GoVersion: %s,
    OperatingSystem: %s,
  }
}`

// Version represents application information that
// follows semantic version guidelines from
// https://semver.org/.
type Version struct {
	// Major represents incompatible API changes.
	Major int64 `json:"major,omitempty"`
	// Minor represents added functionality in a backwards compatible manner.
	Minor int64 `json:"minor,omitempty"`
	// Patch represents backwards compatible bug fixes.
	Patch int64 `json:"patch,omitempty"`
	// PreRelease represents unstable changes that might not be compatible.
	PreRelease string `json:"pre_release,omitempty"`
	// Full represents a fully functional semantic version for the application.
	Full string `json:"full,omitempty"`
	// Metadata represents extra information surrounding the application version.
	Metadata Metadata
}

// Meta implements a formatted string containing only metadata for the Version type.
func (v *Version) Meta() string {
	return v.Metadata.String()
}

// Semantic implements a formatted string containing a formal semantic version for the Version type.
func (v *Version) Semantic() string {
	return v.Full
}

// String implements the Stringer interface for the Version type.
func (v *Version) String() string {
	return fmt.Sprintf(
		versionFormat,
		v.Full,
		v.Major,
		v.Minor,
		v.Patch,
		v.PreRelease,
		v.Metadata.Architecture,
		v.Metadata.BuildDate,
		v.Metadata.Compiler,
		v.Metadata.GitCommit,
		v.Metadata.GoVersion,
		v.Metadata.OperatingSystem,
	)
}
