package oauth

import (
	"time"
)

const (
	expirationTime = 24
)

// AccessToken struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// GetNewAccessToken func to get new access token
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// isExpired func to check expires of access token
func (accessToken AccessToken) isExpired() bool {
	return time.Unix(accessToken.Expires, 0).Before(time.Now().UTC())
}
