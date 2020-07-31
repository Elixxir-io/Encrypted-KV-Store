////////////////////////////////////////////////////////////////////////////////
// Copyright © 2019 Elixxir                                                    /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////
package ekv

import (
	"crypto/rand"
	"testing"
)

// TestCrypto smoke tests the crypto helper functions
func TestCrypto(t *testing.T) {
	plaintext := []byte("Hello, World!")
	password := "test_password"
	ciphertext := encrypt(plaintext, password, rand.Reader)
	decrypted, err := decrypt(ciphertext, password)
	if err != nil {
		t.Errorf("%+v", err)
	}

	for i := 0; i < len(plaintext); i++ {
		if plaintext[i] != decrypted[i] {
			t.Errorf("%b != %b", plaintext[i], decrypted[i])
		}
	}
}
