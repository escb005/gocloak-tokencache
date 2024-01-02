package gocloaktokencache_test

import (
	"testing"
	"time"

	tokencache "github.com/escb005/gocloak-tokencache"

	"github.com/Nerzal/gocloak/v13"
	"github.com/stretchr/testify/require"
)

func ToPointer[T any](v T) *T {
	return &v
}

func TestGetToken(t *testing.T) {
	require := require.New(t)

	cache := tokencache.NewTokenCache("http://localhost:8080",
		tokencache.WithRealm("test"),
		tokencache.WithExpiresSkew(295),
		tokencache.WithTokenOptions(gocloak.TokenOptions{
			GrantType: ToPointer("password"),
			ClientID:  ToPointer("token-cache"),
			Username:  ToPointer("TestUser"),
			Password:  ToPointer("password"),
		}),
	)

	tok1, err := cache.GetToken()
	require.NoError(err)
	require.NotNil(tok1)

	tok2, err := cache.GetToken()
	require.NoError(err)
	require.NotNil(tok2)

	require.Equal(tok1.AccessToken, tok2.AccessToken)

	// wait for token to expire
	time.Sleep(5 * time.Second)

	tok3, err := cache.GetToken()
	require.NoError(err)
	require.NotNil(tok3)

	require.NotEqual(tok1.AccessToken, tok3.AccessToken)
}
