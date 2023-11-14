package gocloaktokencache

import (
	"context"
	"sync"
	"time"

	"github.com/Nerzal/gocloak/v13"
)

type TokenCache struct {
	tokenOptions gocloak.TokenOptions
	expireTime   time.Time
	ctx          context.Context
	client       *gocloak.GoCloak
	token        *gocloak.JWT
	realm        string
	expiresSkew  int
	m            sync.Mutex
}

func NewTokenCache(baseUrl string, options ...TokenCacheOption) *TokenCache {
	// sane defaults
	cache := &TokenCache{
		ctx:          context.Background(),
		client:       gocloak.NewClient(baseUrl),
		realm:        "master",
		expiresSkew:  30,
		tokenOptions: gocloak.TokenOptions{},
	}

	// apply options
	for _, option := range options {
		option(cache)
	}

	return cache
}

func (t *TokenCache) GetToken() (*gocloak.JWT, error) {
	t.m.Lock()
	defer t.m.Unlock()

	// if token is not nil and not expired, then no new token is needed
	if t.token != nil && !t.isExpired() {
		return t.token, nil
	}

	// get new token
	tok, err := t.client.GetToken(t.ctx, t.realm, t.tokenOptions)
	if err != nil {
		return nil, err
	}
	t.token = tok
	t.expireTime = time.Now().Add(time.Duration(tok.ExpiresIn) * time.Second).Add(time.Duration(-t.expiresSkew) * time.Second)

	return tok, nil
}

func (t *TokenCache) isExpired() bool {
	return time.Now().After(t.expireTime)
}
