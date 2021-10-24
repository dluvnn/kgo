package kgo

import (
	"errors"

	"github.com/dluvnn/kgo/curl"
)

var ErrNotFoundCloudFlareID = errors.New("not found cloudflare id")

// GetCloudFlareIdentifier ...
func GetCloudFlareIdentifier(email, apiKey, domain string) (string, error) {
	var x struct {
		Result []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"result"`
	}
	err := curl.Get("https://api.cloudflare.com/client/v4/zones").
		SetHeaderList(
			"X-Auth-Email", email,
			"X-Auth-Key", apiKey,
			"Content-Type", "application/json",
		).Send().ReadJSON(&x)

	if err != nil {
		return "", err
	}
	ndomain := len(domain)
	for _, itm := range x.Result {
		n := len(itm.Name)
		if n > ndomain {
			continue
		}
		if itm.Name == domain[ndomain-n:] {
			return itm.ID, nil
		}
	}
	return "", ErrNotFoundCloudFlareID
}
