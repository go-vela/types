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
	Metadata struct {
		// Architecture represents the architecture information for the application.
		Architecture string `json:"architecture,omitempty"`
		// BuildDate represents the build date information for the application.
		BuildDate string `json:"build_date,omitempty"`
		// Compiler represents the compiler information for the application.
		Compiler string `json:"compiler,omitempty"`
		// GitCommit represents the git commit information for the application.
		GitCommit string `json:"git_commit,omitempty"`
		// GoVersion represents the golang version information for the application.
		GoVersion string `json:"go_version,omitempty"`
		// OperatingSystem represents the operating system information for the application.
		OperatingSystem string `json:"operating_system,omitempty"`
	} `json:"metadata,omitempty"`
}

// Semantic implements a semantic string of the Stringer interface for the Version type.
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
