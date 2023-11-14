package gocloaktokencache

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type TokenCacheOption func(*TokenCache)

func WithContext(ctx context.Context) TokenCacheOption {
	return func(t *TokenCache) {
		t.ctx = ctx
	}
}

func WithRealm(realm string) TokenCacheOption {
	return func(t *TokenCache) {
		t.realm = realm
	}
}

func WithExpiresSkew(expiresSkew int) TokenCacheOption {
	return func(t *TokenCache) {
		t.expiresSkew = expiresSkew
	}
}

func WithTokenOptions(tokenOptions gocloak.TokenOptions) TokenCacheOption {
	return func(t *TokenCache) {
		t.tokenOptions = tokenOptions
	}
}
