package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func (d *Decrypt) Pef1Decrypt(secret, salt, initialVector, tag, data string) (*string, error) {
	// Salt decode to base64
	salt64, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return nil, err
	}

	// InitialVector decode to base64
	iv64, err := base64.StdEncoding.DecodeString(initialVector)
	if err != nil {
		return nil, err
	}

	// Tag decode to base64
	tag64, err := base64.StdEncoding.DecodeString(tag)
	if err != nil {
		return nil, err
	}

	// Data decode to cipher base64
	ciphertext, _ := base64.StdEncoding.DecodeString(data)

	// Generate the HMAC-SHA256 key
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write(salt64)
	key := hash.Sum(nil)

	// Generate the AES-GCM cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	aesgcm, err := cipher.NewGCMWithNonceSize(block, len(iv64))
	if err != nil {
		panic(err)
	}

	// Set tag to ciphertext
	ciphertextWithTag := append(ciphertext, tag64...)

	// Decrypt the data
	plaintext, err := aesgcm.Open(nil, iv64, ciphertextWithTag, nil)
	if err != nil {
		panic(err)
	}

	// Convert base64 to string
	result := string(plaintext)
	return &result, nil
}
