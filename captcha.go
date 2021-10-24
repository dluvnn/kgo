package kgo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	siteVerifyURL = "https://www.google.com/recaptcha/api/siteverify"
)

var (
	// ErrInvalidRecaptchaToken ...
	ErrInvalidRecaptchaToken = errors.New("invalid google recaptcha token")
)

// RecaptchaResponse ...
type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

// VerifyRecaptcha ...
func VerifyRecaptcha(ip, token, key string) (*RecaptchaResponse, error) {
	if len(token) == 0 {
		return nil, ErrInvalidRecaptchaToken
	}
	resp, err := http.PostForm(
		siteVerifyURL,
		url.Values{"secret": {key}, "remoteip": {ip}, "response": {token}},
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var x RecaptchaResponse
	err = json.Unmarshal(body, &x)
	if err != nil {
		return nil, err
	}
	return &x, nil
}

// VerifyRecaptchaV2 ...
func VerifyRecaptchaV2(ip, token, secretKey string) (*RecaptchaResponse, error) {
	return VerifyRecaptcha(ip, token, secretKey)
}

// VerifyRecaptchaV3 ...
func VerifyRecaptchaV3(ip, token, secretKey string) (*RecaptchaResponse, error) {
	return VerifyRecaptcha(ip, token, secretKey)
}
