package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExpirationTimeConstant(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestAccessToken(t *testing.T) {
	at := GetAccessToken()
	assert.False(t, at.IsExpired(), "access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "access token should not be defined")
	assert.True(t, at.UserId == 0, "access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token created three hours from now should not be expired")
}
