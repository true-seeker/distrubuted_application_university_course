package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type AES struct {
	key []byte
}

func (a *AES) Encrypt(data string) ([]byte, error) {
	plaintext := []byte(data)
	c, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func (a *AES) Decrypt(ciphertext []byte) ([]byte, error) {
	c, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, nil
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
