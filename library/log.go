// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"regexp"

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
	InitID    *int64 `json:"init_id,omitempty"`
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
	data := l.GetData()

	for _, secret := range secrets {
		// escape regexp meta characters if they exist within value of secret
		//
		// https://pkg.go.dev/regexp#QuoteMeta
		escaped := regexp.QuoteMeta(secret)

		// create regexp to match secrets in the log data surrounded by regexp metacharacters
		//
		// https://pkg.go.dev/regexp#MustCompile
		re := regexp.MustCompile((`(\s|^|=|"|:|'|\.|,)` + escaped + `(\s|$|"|:|'|\.|,)`))

		// create a mask for the secret
		mask := fmt.Sprintf("$1%s$2", constants.SecretLogMask)

		// replace all regexp matches of secret with mask
		//
		// https://pkg.go.dev/regexp#Regexp.ReplaceAll
		data = re.ReplaceAll(data, []byte(mask))
	}

	// update data field to masked logs
	l.SetData(data)
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

// GetInitID returns the InitID field.
//
// When the provided Log type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Log) GetInitID() int64 {
	// return zero value if Log type or InitID field is nil
	if l == nil || l.InitID == nil {
		return 0
	}

	return *l.InitID
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

// SetInitID sets the InitID field.
//
// When the provided Log type is nil, it
// will set nothing and immediately return.
func (l *Log) SetInitID(v int64) {
	// return if Log type is nil
	if l == nil {
		return
	}

	l.InitID = &v
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
  InitID: %d,
}`,
		l.GetBuildID(),
		l.GetData(),
		l.GetID(),
		l.GetRepoID(),
		l.GetServiceID(),
		l.GetStepID(),
		l.GetInitID(),
	)
}
