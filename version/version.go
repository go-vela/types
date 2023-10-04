// SPDX-License-Identifier: Apache-2.0

package version

import (
	"fmt"
)

const versionFormat = `{
  Canonical: %s,
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
//
// swagger:model Version
type Version struct {
	// Canonical represents a canonical semantic version for the application.
	Canonical string `json:"canonical"`
	// Major represents incompatible API changes.
	Major uint64 `json:"major"`
	// Minor represents added functionality in a backwards compatible manner.
	Minor uint64 `json:"minor"`
	// Patch represents backwards compatible bug fixes.
	Patch uint64 `json:"patch"`
	// PreRelease represents unstable changes that might not be compatible.
	PreRelease string `json:"pre_release,omitempty"`
	// Metadata represents extra information surrounding the application version.
	Metadata Metadata `json:"metadata,omitempty"`
}

// Meta implements a formatted string containing only metadata for the Version type.
func (v *Version) Meta() string {
	return v.Metadata.String()
}

// Semantic implements a formatted string containing a formal semantic version for the Version type.
func (v *Version) Semantic() string {
	return v.Canonical
}

// String implements the Stringer interface for the Version type.
func (v *Version) String() string {
	return fmt.Sprintf(
		versionFormat,
		v.Canonical,
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
