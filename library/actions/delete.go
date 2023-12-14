// SPDX-License-Identifier: Apache-2.0

package actions

// Deploy is the library representation of the various actions associated
// with the deploy event webhook from the SCM.
type Delete struct {
	Created *bool `json:"created"`
}
