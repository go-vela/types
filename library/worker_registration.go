// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

// QueueRegistration is the library representation of a QueueRegistration.
//
// swagger:model QueueRegistration
type QueueRegistration struct {
	QueuePublicKey *string `json:"queue_public_key,omitempty"`
	QueueAddress   *string `json:"queue_address,omitempty"`
}

// GetPublicKey returns the QueuePublicKey field.
//
// When the provided QueueRegistration type is nil, or the field within
// the type is nil, it returns an empty string for the field.
func (w *QueueRegistration) GetPublicKey() string {
	// return zero value if QueueRegistration type or QueuePublicKey field is nil
	if w == nil || w.QueuePublicKey == nil {
		return ""
	}

	return *w.QueuePublicKey
}

// GetQueueAddress returns the QueueAddress field.
//
// When the provided QueueRegistration type is nil, or the field within
// the type is nil, it returns an empty string for the field.
func (w *QueueRegistration) GetQueueAddress() string {
	// return zero value if QueueRegistration type or QueueAddress field is nil
	if w == nil || w.QueueAddress == nil {
		return ""
	}

	return *w.QueueAddress
}

// SetPublicKey sets the QueuePublicKey field.
//
// When the provided QueueRegistration type is nil, it
// will set nothing and immediately return.
func (w *QueueRegistration) SetPublicKey(v string) {
	// return if QueueRegistration type is nil
	if w == nil {
		return
	}

	w.QueuePublicKey = &v
}

// SetQueueAddress sets the QueueAddress field.
//
// When the provided QueueRegistration type is nil, it
// will set nothing and immediately return.
func (w *QueueRegistration) SetQueueAddress(v string) {
	// return if QueueRegistration type is nil
	if w == nil {
		return
	}

	w.QueueAddress = &v
}
