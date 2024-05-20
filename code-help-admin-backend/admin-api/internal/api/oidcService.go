package api

import (
	"admin-api/internal/environment"
	"context"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"log"
	"net/http"
)

type oidcServiceImpl struct {
	provider *oidc.Provider
	verifier *oidc.IDTokenVerifier
}

type OidcService interface {
	Verify(ctx context.Context, token string) (*oidc.IDToken, error)
	AttachAuthHeader(ctx context.Context, req *http.Request) error
}

func NewOidcService(environment environment.Environment) OidcService {
	ctx := context.Background()
	issuerUri := environment.Oauth2Jwt.IssuerUri

	provider, err := oidc.NewProvider(ctx, issuerUri)

	if err != nil {
		log.Fatal(err)
	}

	oidcConfig := &oidc.Config{
		SkipClientIDCheck: true,
	}

	verifier := provider.Verifier(oidcConfig)

	return &oidcServiceImpl{
		provider: provider,
		verifier: verifier,
	}
}

func (it *oidcServiceImpl) Verify(ctx context.Context, token string) (*oidc.IDToken, error) {
	return it.verifier.Verify(ctx, token)
}

func (it *oidcServiceImpl) AttachAuthHeader(ctx context.Context, req *http.Request) error {
	if authContext, ok := ctx.Value(AuthenticationContextKey).(AuthContext); ok {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authContext.Token))
	}

	return nil
}
