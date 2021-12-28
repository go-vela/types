// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import "fmt"

// Pipeline is the library representation of a Pipeline.
//
// swagger:model Pipeline
type Pipeline struct {
	ID        *int64  `json:"id,omitempty"`
	RepoID    *int64  `json:"repo_id,omitempty"`
	Flavor    *string `json:"flavor,omitempty"`
	Platform  *string `json:"platform,omitempty"`
	Ref       *string `json:"ref,omitempty"`
	Version   *string `json:"version,omitempty"`
	Services  *bool   `json:"services,omitempty"`
	Stages    *bool   `json:"stages,omitempty"`
	Steps     *bool   `json:"steps,omitempty"`
	Templates *bool   `json:"templates,omitempty"`
	// swagger:strfmt base64
	Data *[]byte `json:"data,omitempty"`
}

// GetID returns the ID field.
//
// When the provided Pipeline type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (p *Pipeline) GetID() int64 {
	// return zero value if Pipeline type or ID field is nil
	if p == nil || p.ID == nil {
		return 0
	}

	return *p.ID
}

// GetRepoID returns the RepoID field.
//
// When the provided Pipeline type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (p *Pipeline) GetRepoID() int64 {
	// return zero value if Pipeline type or RepoID field is nil
	if p == nil || p.RepoID == nil {
		return 0
	}

	return *p.RepoID
}

// GetFlavor returns the Flavor field.
//
// When the provided Pipeline type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (p *Pipeline) GetFlavor() string {
	// return zero value if Pipeline type or Flavor field is nil
	if p == nil || p.Flavor == nil {
		return ""
	}

	return *p.Flavor
}

// GetPlatform returns the Platform field.
//
// When the provided Pipeline type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (p *Pipeline) GetPlatform() string {
	// return zero value if Pipeline type or Platform field is nil
	if p == nil || p.Platform == nil {
		return ""
	}

	return *p.Platform
}

// GetRef returns the Ref field.
//
// When the provided Pipeline type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (p *Pipeline) GetRef() string {
	// return zero value if Pipeline type or Ref field is nil
	if p == nil || p.Ref == nil {
		return ""
	}

	return *p.Ref
}

// GetVersion returns the Version field.
//
// When the provided Pipeline type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (p *Pipeline) GetVersion() string {
	// return zero value if Pipeline type or Version field is nil
	if p == nil || p.Version == nil {
		return ""
	}

	return *p.Version
}

// GetServices returns the Services field.
//
// When the provided Pipeline type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (p *Pipeline) GetServices() bool {
	// return zero value if Pipeline type or Services field is nil
	if p == nil || p.Services == nil {
		return false
	}

	return *p.Services
}

// GetStages returns the Stages field.
//
// When the provided Pipeline type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (p *Pipeline) GetStages() bool {
	// return zero value if Pipeline type or Stages field is nil
	if p == nil || p.Stages == nil {
		return false
	}

	return *p.Stages
}

// GetSteps returns the Steps field.
//
// When the provided Pipeline type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (p *Pipeline) GetSteps() bool {
	// return zero value if Pipeline type or Steps field is nil
	if p == nil || p.Steps == nil {
		return false
	}

	return *p.Steps
}

// GetTemplates returns the Templates field.
//
// When the provided Pipeline type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (p *Pipeline) GetTemplates() bool {
	// return zero value if Pipeline type or Templates field is nil
	if p == nil || p.Templates == nil {
		return false
	}

	return *p.Templates
}

// GetData returns the Data field.
//
// When the provided Pipeline type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (p *Pipeline) GetData() []byte {
	// return zero value if Pipeline type or Data field is nil
	if p == nil || p.Data == nil {
		return []byte{}
	}

	return *p.Data
}

// SetID sets the ID field.
//
// When the provided Pipeline type is nil, it
// will set nothing and immediately return.
func (p *Pipeline) SetID(v int64) {
	// return if Pipeline type is nil
	if p == nil {
		return
	}

	p.ID = &v
}

// SetRepoID sets the RepoID field.
//
// When the provided Pipeline type is nil, it
// will set nothing and immediately return.
func (p *Pipeline) SetRepoID(v int64) {
	// return if Pipeline type is nil
	if p == nil {
		return
	}

	p.RepoID = &v
}

// SetFlavor sets the Flavor field.
//
// When the provided Pipeline type is nil, it
// will set nothing and immediately return.
func (p *Pipeline) SetFlavor(v string) {
	// return if Pipeline type is nil
	if p == nil {
		return
	}

	p.Flavor = &v
}

// SetPlatform sets the Platform field.
//
// When the provided Pipeline type is nil, it
// will set nothing and immediately return.
func (p *Pipeline) SetPlatform(v string) {
	// return if Pipeline type is nil
	if p == nil {
		return
	}

	p.Platform = &v
}

// SetRef sets the Ref field.
//
// When the provided Pipeline type is nil, it
// will set nothing and immediately return.
func (p *Pipeline) SetRef(v string) {
	// return if Pipeline type is nil
	if p == nil {
		return
	}

	p.Ref = &v
}

// SetVersion sets the Version field.
//
// When the provided Pipeline type is nil, it
// will set nothing and immediately return.
func (p *Pipeline) SetVersion(v string) {
	// return if Pipeline type is nil
	if p == nil {
		return
	}

	p.Version = &v
}

// SetServices sets the Services field.
//
// When the provided Pipeline type is nil, it
// will set nothing and immediately return.
func (p *Pipeline) SetServices(v bool) {
	// return if Pipeline type is nil
	if p == nil {
		return
	}

	p.Services = &v
}

// SetStages sets the Stages field.
//
// When the provided Pipeline type is nil, it
// will set nothing and immediately return.
func (p *Pipeline) SetStages(v bool) {
	// return if Pipeline type is nil
	if p == nil {
		return
	}

	p.Stages = &v
}

// SetSteps sets the Steps field.
//
// When the provided Pipeline type is nil, it
// will set nothing and immediately return.
func (p *Pipeline) SetSteps(v bool) {
	// return if Pipeline type is nil
	if p == nil {
		return
	}

	p.Steps = &v
}

// SetTemplates sets the Templates field.
//
// When the provided Pipeline type is nil, it
// will set nothing and immediately return.
func (p *Pipeline) SetTemplates(v bool) {
	// return if Pipeline type is nil
	if p == nil {
		return
	}

	p.Templates = &v
}

// SetData sets the Data field.
//
// When the provided Pipeline type is nil, it
// will set nothing and immediately return.
func (p *Pipeline) SetData(v []byte) {
	// return if Log type is nil
	if p == nil {
		return
	}

	p.Data = &v
}

// String implements the Stringer interface for the Pipeline type.
func (p *Pipeline) String() string {
	return fmt.Sprintf(`{
  Data: %s,
  Flavor: %s,
  ID: %d,
  Platform: %s,
  Ref: %s,
  RepoID: %d,
  Services: %t,
  Stages: %t,
  Steps: %t,
  Templates: %t,
  Version: %s,
}`,
		p.GetData(),
		p.GetFlavor(),
		p.GetID(),
		p.GetPlatform(),
		p.GetRef(),
		p.GetRepoID(),
		p.GetServices(),
		p.GetStages(),
		p.GetSteps(),
		p.GetTemplates(),
		p.GetVersion(),
	)
}
