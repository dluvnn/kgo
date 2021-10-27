package kgo

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
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

// RandInt ...
func RandInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// ComputeMD5Checksum computes the MD5 checksum of the file 'filename'
func ComputeMD5Checksum(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// ComputeSHA1Checksum computes the SHA1 checksum of the file 'filename'
func ComputeSHA1Checksum(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
