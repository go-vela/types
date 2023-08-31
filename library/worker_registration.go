// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

// WorkerRegistration is the library representation of a WorkerRegistration.
//
// swagger:model WorkerRegistration
type WorkerRegistration struct {
	RegistrationToken *string `json:"registration_token,omitempty"`
	QueuePublicKey    *string `json:"queue_public_key,omitempty"`
	QueueAddress      *string `json:"queue_address,omitempty"`
}

// GetRegistrationToken returns the RegistrationToken field.
//
// When the provided Worker type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (q *WorkerRegistration) GetRegistrationToken() string {
	// return zero value if Worker type or ID field is nil
	if q == nil || q.RegistrationToken == nil {
		return ""
	}

	return *q.RegistrationToken
}

// GetPublicKey returns the QueuePublicKey field.
//
// When the provided Worker type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (q *WorkerRegistration) GetPublicKey() string {
	// return zero value if Worker type or ID field is nil
	if q == nil || q.QueuePublicKey == nil {
		return ""
	}

	return *q.QueuePublicKey
}

// GetQueueAddress returns the QueueAddress field.
//
// When the provided Worker type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (q *WorkerRegistration) GetQueueAddress() string {
	// return zero value if Worker type or ID field is nil
	if q == nil || q.QueueAddress == nil {
		return ""
	}

	return *q.QueueAddress
}

// SetRegistrationToken sets the RegistrationToken field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (q *WorkerRegistration) SetRegistrationToken(v string) {
	// return if Worker type is nil
	if q == nil {
		return
	}

	q.RegistrationToken = &v
}

// SetPublicKey sets the QueuePublicKey field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (q *WorkerRegistration) SetPublicKey(v string) {
	// return if Worker type is nil
	if q == nil {
		return
	}

	q.QueuePublicKey = &v
}

// SetQueueAddress sets the QueueAddress field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (q *WorkerRegistration) SetQueueAddress(v string) {
	// return if Worker type is nil
	if q == nil {
		return
	}

	q.QueueAddress = &v
}
