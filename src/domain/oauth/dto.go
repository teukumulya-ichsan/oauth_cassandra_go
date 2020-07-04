package oauth

import (
	"oauth_cassandra_golang/src/utils/errors"
	"strings"
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

// Validate method to validate access Token
func (accessToken *AccessToken) Validate() *errors.RestErr {
	accessToken.AccessToken = strings.TrimSpace(accessToken.AccessToken)
	if accessToken.AccessToken == "" {
		return errors.NewBadRequestError(invalidAccessTokenID)
	}
	if accessToken.UserID <= 0 {
		return errors.NewBadRequestError(invalidUserID)
	}
	if accessToken.ClientID <= 0 {
		return errors.NewBadRequestError(invalidClientID)
	}
	if accessToken.Expires <= 0 {
		return errors.NewBadRequestError(invalidExpireTime)
	}
	return nil
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
