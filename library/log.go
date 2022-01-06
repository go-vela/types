// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-vela/types/constants"
)

// Log is the library representation of a log for a step in a build.
//
// swagger:model Log
type Log struct {
	ID        *int64 `json:"id,omitempty"`
	BuildID   *int64 `json:"build_id,omitempty"`
	RepoID    *int64 `json:"repo_id,omitempty"`
	ServiceID *int64 `json:"service_id,omitempty"`
	StepID    *int64 `json:"step_id,omitempty"`
	// swagger:strfmt base64
	Data *[]byte `json:"data,omitempty"`
}

// AppendData adds the provided data to the end of
// the Data field for the Log type. If the Data
// field is empty, then the function overwrites
// the entire Data field.
func (l *Log) AppendData(data []byte) {
	// check if Data field is empty
	if len(l.GetData()) == 0 {
		// overwrite the Data field
		l.SetData(data)

		return
	}

	// add the data to the Data field
	l.SetData(append(l.GetData(), data...))
}

// MaskData reads through the log data and masks
// all values provided in the string slice. If the
// log is empty, we do nothing.
func (l *Log) MaskData(secrets []string) {
	// convert data to string
	strData := string(l.GetData())
	for _, secret := range secrets {
		// escape regexp meta characters if they exist within value of secret
		sanitizedStr := regexp.QuoteMeta(secret)

		// find matches in logs for secrets surrounded by whitespace, start of string,
		// quotation, colon, period, comma, and end of string.
		// this prevents partial masking of secrets that contain a substring of another
		// secret.
		re := regexp.MustCompile((`(\s|^|"|:|'|\.|,)` + sanitizedStr + `(\s|$|"|:|'|\.|,)`))
		matches := re.FindAllString(strData, -1)

		// take all matches and mask them
		for _, match := range matches {
			// construct mask
			mask := ""

			// if secret was logged as the very first or last thing in the log, don't add
			// boundary mask values.
			if match[0] != secret[0] {
				mask += string(match[0])
			}
			mask += constants.SecretLogMask
			if match[len(match)-1] != secret[len(secret)-1] {
				mask += string(match[len(match)-1])
			}

			// replace log with new mask
			strData = strings.Replace(strData, match, mask, 1)
		}
	}

	// update data field to masked logs
	l.SetData([]byte(strData))
}

// GetID returns the ID field.
//
// When the provided Log type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Log) GetID() int64 {
	// return zero value if Log type or ID field is nil
	if l == nil || l.ID == nil {
		return 0
	}

	return *l.ID
}

// GetBuildID returns the BuildID field.
//
// When the provided Log type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Log) GetBuildID() int64 {
	// return zero value if Log type or BuildID field is nil
	if l == nil || l.BuildID == nil {
		return 0
	}

	return *l.BuildID
}

// GetRepoID returns the RepoID field.
//
// When the provided Log type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Log) GetRepoID() int64 {
	// return zero value if Log type or RepoID field is nil
	if l == nil || l.RepoID == nil {
		return 0
	}

	return *l.RepoID
}

// GetServiceID returns the ServiceID field.
//
// When the provided Log type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Log) GetServiceID() int64 {
	// return zero value if Log type or ServiceID field is nil
	if l == nil || l.ServiceID == nil {
		return 0
	}

	return *l.ServiceID
}

// GetStepID returns the StepID field.
//
// When the provided Log type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Log) GetStepID() int64 {
	// return zero value if Log type or StepID field is nil
	if l == nil || l.StepID == nil {
		return 0
	}

	return *l.StepID
}

// GetData returns the Data field.
//
// When the provided Log type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Log) GetData() []byte {
	// return zero value if Log type or Data field is nil
	if l == nil || l.Data == nil {
		return []byte{}
	}

	return *l.Data
}

// SetID sets the ID field.
//
// When the provided Log type is nil, it
// will set nothing and immediately return.
func (l *Log) SetID(v int64) {
	// return if Log type is nil
	if l == nil {
		return
	}

	l.ID = &v
}

// SetBuildID sets the BuildID field.
//
// When the provided Log type is nil, it
// will set nothing and immediately return.
func (l *Log) SetBuildID(v int64) {
	// return if Log type is nil
	if l == nil {
		return
	}

	l.BuildID = &v
}

// SetRepoID sets the RepoID field.
//
// When the provided Log type is nil, it
// will set nothing and immediately return.
func (l *Log) SetRepoID(v int64) {
	// return if Log type is nil
	if l == nil {
		return
	}

	l.RepoID = &v
}

// SetServiceID sets the ServiceID field.
//
// When the provided Log type is nil, it
// will set nothing and immediately return.
func (l *Log) SetServiceID(v int64) {
	// return if Log type is nil
	if l == nil {
		return
	}

	l.ServiceID = &v
}

// SetStepID sets the StepID field.
//
// When the provided Log type is nil, it
// will set nothing and immediately return.
func (l *Log) SetStepID(v int64) {
	// return if Log type is nil
	if l == nil {
		return
	}

	l.StepID = &v
}

// SetData sets the Data field.
//
// When the provided Log type is nil, it
// will set nothing and immediately return.
func (l *Log) SetData(v []byte) {
	// return if Log type is nil
	if l == nil {
		return
	}

	l.Data = &v
}

// String implements the Stringer interface for the Log type.
func (l *Log) String() string {
	return fmt.Sprintf(`{
  BuildID: %d,
  Data: %s,
  ID: %d,
  RepoID: %d,
  ServiceID: %d,
  StepID: %d,
}`,
		l.GetBuildID(),
		l.GetData(),
		l.GetID(),
		l.GetRepoID(),
		l.GetServiceID(),
		l.GetStepID(),
	)
}
