package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"proxy/internal/environment"

	"github.com/coreos/go-oidc/v3/oidc"
)

type oidcServiceImpl struct {
	provider *oidc.Provider
	verifier *oidc.IDTokenVerifier
}

type OidcService interface {
	Verify(ctx context.Context, token string) (*oidc.IDToken, error)
	AttachAuthHeader(ctx context.Context, req *http.Request) error
}

func NewOidcService(environment environment.Environment) (OidcService, error) {
	ctx := context.Background()
	issuerUri := environment.Oauth2Jwt.IssuerUri
	discoveryIssuerUrl := environment.Oauth2Jwt.DiscoveryIssuerUri

	ctx = oidc.InsecureIssuerURLContext(ctx, issuerUri)
	provider, err := oidc.NewProvider(ctx, discoveryIssuerUrl)

	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}

	oidcConfig := &oidc.Config{
		SkipClientIDCheck: true,
		SkipIssuerCheck:   true,
	}

	verifier := provider.Verifier(oidcConfig)

	return &oidcServiceImpl{
		provider: provider,
		verifier: verifier,
	}, nil
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
