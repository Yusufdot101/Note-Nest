package oidc

import (
	"context"
	"errors"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

// ---------- Google OIDC ----------

func googleConfig(clientID, clientSecret, redirectURL string) (*oauth2.Config, *oidc.Provider, error) {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		return nil, nil, err
	}

	cfg := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	return cfg, provider, nil
}

func googleCallback(ctx context.Context, code string, config *oauth2.Config, provider *oidc.Provider) (*UserInfo, error) {
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: config.ClientID})
	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return nil, err
	}

	var claims struct {
		Sub   string `json:"sub"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := idToken.Claims(&claims); err != nil {
		return nil, err
	}

	return &UserInfo{
		ID:    claims.Sub,
		Name:  claims.Name,
		Email: claims.Email,
	}, nil
}
