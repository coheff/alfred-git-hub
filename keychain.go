package main

import (
	"encoding/json"

	"github.com/deanishe/awgo/keychain"
	"golang.org/x/oauth2"
)

const key = "token"

var kc = keychain.New("alfred-git-hub")

// cacheToken adds a token to Keychain. If a token already exists, it is replaced.
func CacheToken(tok *oauth2.Token) error {
	jToken, err := json.Marshal(tok)
	if err != nil {
		return err
	}

	err = kc.Set(key, string(jToken))
	if err != nil {
		return err
	}
	return nil
}

// cachedToken retrieves a token from Keychain.
func CachedToken() (*oauth2.Token, error) {
	jToken, err := kc.Get(key)
	if err != nil {
		return nil, err
	}

	var tok oauth2.Token
	err = json.Unmarshal([]byte(jToken), &tok)
	if err != nil {
		return nil, err
	}
	return &tok, err
}
