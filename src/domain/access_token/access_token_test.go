package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	accessToken := GetNewAccessToken()
	assert.False(t, accessToken.isExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, "", accessToken.AccessToken, "new access token should not have defined token id")
	assert.True(t, accessToken.UserID == 0, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	accessToken := AccessToken{}
	assert.True(t, accessToken.isExpired(), "empty access token should be expired by default")

	accessToken.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, accessToken.isExpired(), "access token created 3 hours from now should NOT be expired")
}
