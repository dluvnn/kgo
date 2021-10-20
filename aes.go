package kgo

import (
	"crypto/aes"
	"crypto/cipher"
	"math/rand"
)

type AESCipher struct {
	gcm       cipher.AEAD
	nonceSize int
}

func (ac *AESCipher) Encrypt(value []byte) ([]byte, error) {
	nonce := make([]byte, ac.nonceSize)
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	return ac.gcm.Seal(nonce, nonce, value, nil), nil
}

func (ac *AESCipher) Decrypt(value []byte) ([]byte, error) {
	nonce, value := value[:ac.nonceSize], value[ac.nonceSize:]
	return ac.gcm.Open(nil, nonce, value, nil)
}

func CreateAESCipher(key []byte) (*AESCipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ac := new(AESCipher)
	ac.gcm = gcm
	ac.nonceSize = gcm.NonceSize()
	return ac, nil
}
