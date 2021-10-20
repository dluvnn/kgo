package kgo

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
)

const (
	allChars    = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	lenAllChars = len(allChars)
)

var (
	ErrTooSmallSize = errors.New("the size of input value is small than requirement")
)

// RandPassword generates password randomly
func RandPassword(length int) string {
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = allChars[rand.Intn(lenAllChars)]
	}
	return string(buf)
}

// GenKey ...
func GenKey(n int) string {
	b, err := GenRawKey(n)
	if err != nil {
		Error(err)
		return ""
	}
	return fmt.Sprintf("%02x", b)
}

// GenRawKey ...
func GenRawKey(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}

// SHA256 creates hash value of data
func SHA256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	tmp := h.Sum(nil)
	return base64.RawStdEncoding.EncodeToString(tmp)
}

// HashPassword ...
func HashPassword(pw, uname string) string {
	return SHA256(pw + "a7}6r!pjt@+(B]7" + uname)
}

// // Encrypt ...
// func Encrypt(key, value []byte) ([]byte, error) {
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		return nil, err
// 	}
// 	nonce := make([]byte, gcm.NonceSize())
// 	_, err = rand.Read(nonce)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return gcm.Seal(nonce, nonce, value, nil), nil
// }

// // Decrypt ...
// func Decrypt(key, value []byte) ([]byte, error) {
// 	c, err := aes.NewCipher(key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	gcm, err := cipher.NewGCM(c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	nonceSize := gcm.NonceSize()
// 	if len(value) < nonceSize {
// 		return nil, ErrTooSmallSize
// 	}
// 	nonce, value := value[:nonceSize], value[nonceSize:]
// 	return gcm.Open(nil, nonce, value, nil)
// }

// RandInt ...
func RandInt(min, max int) int {
	return rand.Intn(max-min) + min
}
