package auth0

import (
	"context"
	"errors"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/withzeus/amimugi/config"
	"golang.org/x/oauth2"
)

type Auth0 struct {
	*oidc.Provider
	oauth2.Config
}

func NewAuth0() (*Auth0, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+config.Env.Auth0Domain+"/",
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		Endpoint:    provider.Endpoint(),
		RedirectURL: config.Env.Auth0CallbackUrl,
		Scopes:      []string{oidc.ScopeOpenID, "profile"},
	}

	return &Auth0{
		Provider: provider,
		Config:   conf,
	}, nil
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (a *Auth0) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
