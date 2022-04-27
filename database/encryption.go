// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// decrypt is a helper function to decrypt values. First
// a AES-256 Galois Counter Mode cipher block is created
// from the encryption key to decrypt the value. Then, we
// verify the value isn't smaller than the nonce which
// would indicate the value isn't encrypted. Finally the
// cipher block and nonce is used to decrypt the value.
func decrypt(key string, value []byte) ([]byte, error) {
	// create a new cipher block from the encryption key
	//
	// the key should have a length of 64 bits to ensure
	// we are using the AES-256 standard
	//
	// https://en.wikipedia.org/wiki/Advanced_Encryption_Standard
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return value, err
	}

	// creates a new Galois Counter Mode cipher block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return value, err
	}

	// nonce is an arbitrary number used to to ensure that
	// old communications cannot be reused in replay attacks.
	//
	// https://en.wikipedia.org/wiki/Cryptographic_nonce
	nonceSize := gcm.NonceSize()

	// verify the value has a length greater than the nonce
	//
	// if the value is less than the nonce size, then we
	// can assume the value hasn't been encrypted yet.
	if len(value) < nonceSize {
		return value, fmt.Errorf("invalid value length for decrypt provided: %d", len(value))
	}

	// capture nonce and ciphertext from the value
	nonce, ciphertext := value[:nonceSize], value[nonceSize:]

	// decrypt the value from the ciphertext
	return gcm.Open(nil, nonce, ciphertext, nil)
}

// encrypt is a helper function to encrypt values. First
// a AES-256 Galois Counter Mode cipher block is created
// from the encryption key to encrypt the value. Then,
// we create the nonce from a cryptographically secure
// random number generator. Finally, the cipher block
// and nonce is used to encrypt the value.
func encrypt(key string, value []byte) ([]byte, error) {
	// create a new cipher block from the encryption key
	//
	// the key should have a length of 64 bits to ensure
	// we are using the AES-256 standard
	//
	// https://en.wikipedia.org/wiki/Advanced_Encryption_Standard
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return value, err
	}

	// creates a new Galois Counter Mode cipher block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return value, err
	}

	// nonce is an arbitrary number used to to ensure that
	// old communications cannot be reused in replay attacks.
	//
	// https://en.wikipedia.org/wiki/Cryptographic_nonce
	nonce := make([]byte, gcm.NonceSize())

	// set nonce from a cryptographically secure random number generator
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return value, err
	}

	// encrypt the value with the randomly generated nonce
	return gcm.Seal(nonce, nonce, value, nil), nil
}
