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
// When the provided WorkerRegistration type is nil, or the field within
// the type is nil, it returns an empty string for the field.
func (w *WorkerRegistration) GetRegistrationToken() string {
	// return zero value if WorkerRegistration type or RegistrationToken field is nil
	if w == nil || w.RegistrationToken == nil {
		return ""
	}

	return *w.RegistrationToken
}

// GetPublicKey returns the QueuePublicKey field.
//
// When the provided WorkerRegistration type is nil, or the field within
// the type is nil, it returns an empty string for the field.
func (w *WorkerRegistration) GetPublicKey() string {
	// return zero value if WorkerRegistration type or QueuePublicKey field is nil
	if w == nil || w.QueuePublicKey == nil {
		return ""
	}

	return *w.QueuePublicKey
}

// GetQueueAddress returns the QueueAddress field.
//
// When the provided WorkerRegistration type is nil, or the field within
// the type is nil, it returns an empty string for the field.
func (w *WorkerRegistration) GetQueueAddress() string {
	// return zero value if WorkerRegistration type or QueueAddress field is nil
	if w == nil || w.QueueAddress == nil {
		return ""
	}

	return *w.QueueAddress
}

// SetRegistrationToken sets the RegistrationToken field.
//
// When the provided WorkerRegistration type is nil, it
// will set nothing and immediately return.
func (w *WorkerRegistration) SetRegistrationToken(v string) {
	// return if WorkerRegistration type is nil
	if w == nil {
		return
	}

	w.RegistrationToken = &v
}

// SetPublicKey sets the QueuePublicKey field.
//
// When the provided WorkerRegistration type is nil, it
// will set nothing and immediately return.
func (w *WorkerRegistration) SetPublicKey(v string) {
	// return if WorkerRegistration type is nil
	if w == nil {
		return
	}

	w.QueuePublicKey = &v
}

// SetQueueAddress sets the QueueAddress field.
//
// When the provided WorkerRegistration type is nil, it
// will set nothing and immediately return.
func (w *WorkerRegistration) SetQueueAddress(v string) {
	// return if WorkerRegistration type is nil
	if w == nil {
		return
	}

	w.QueueAddress = &v
}
