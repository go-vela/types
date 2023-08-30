// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

// QueueRegistration is the library representation of a QueueRegistration.
//
// swagger:model QueueRegistration
type QueueRegistration struct {
	QueuePublicKey *string `json:"queue-public-key,omitempty"`
	QueueAddress   *string `json:"queue-address,omitempty"`
}

// GetPublicKey returns the ID field.
//
// When the provided Worker type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (q *QueueRegistration) GetPublicKey() string {
	// return zero value if Worker type or ID field is nil
	if q == nil || q.QueuePublicKey == nil {
		return ""
	}

	return *q.QueuePublicKey
}

func (q *QueueRegistration) GetQueueAddress() string {
	// return zero value if Worker type or ID field is nil
	if q == nil || q.QueueAddress == nil {
		return ""
	}

	return *q.QueueAddress
}

// SetPublicKey sets the ID field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (q *QueueRegistration) SetPublicKey(v string) {
	// return if Worker type is nil
	if q == nil {
		return
	}

	q.QueuePublicKey = &v
}

// SetQueueAddress sets the ID field.
//
// When the provided Worker type is nil, it
// will set nothing and immediately return.
func (q *QueueRegistration) SetQueueAddress(v string) {
	// return if Worker type is nil
	if q == nil {
		return
	}

	q.QueueAddress = &v
}
