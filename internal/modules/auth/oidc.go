package auth

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type OIDCManager struct {
	Provider     *oidc.Provider
	Verifier     *oidc.IDTokenVerifier
	OAuth2Config *oauth2.Config
}

func NewOIDCManager(domain, clientID, clientSecret, callbackURL string) (*OIDCManager, error) {
	ctx := context.Background()

	// Auth0 issuer URL
	issuer := "https://" + domain + "/"

	// go-oidc docs অনুযায়ী provider init
	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize oidc provider: %w", err)
	}

	// oauth2 config
	oauth2Config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  callbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes: []string{
			oidc.ScopeOpenID, // mandatory for OIDC
			"profile",
			"email",
		},
	}

	// ID token verifier
	verifier := provider.Verifier(&oidc.Config{
		ClientID: clientID,
	})

	return &OIDCManager{
		Provider:     provider,
		Verifier:     verifier,
		OAuth2Config: oauth2Config,
	}, nil
}
