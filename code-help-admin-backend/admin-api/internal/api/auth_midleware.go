package api

import (
	"admin-api/internal/environment"
	codeHelpAdminGen "api-spec/generated/code-help-admin"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
)

type AuthMiddleware struct {
	handler     http.Handler
	environment environment.Environment
}

func NewAuthMiddlewareWith(environment environment.Environment) codeHelpAdminGen.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return AuthMiddleware{handler: handler, environment: environment}
	}
}

const (
	AuthenticationContextKey = "Authentication"
)

type AuthContext struct {
	Token *oidc.IDToken
}

func (a AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := context.WithValue(request.Context(), AuthenticationContextKey, AuthContext{})

	authorizationHeaderValues := request.Header["Authorization"]

	if len(authorizationHeaderValues) == 0 || authorizationHeaderValues[0] == "" {
		a.handler.ServeHTTP(writer, request)
		return
	}

	idToken := strings.Replace(authorizationHeaderValues[0], "Bearer ", "", 1)

	issuerUri := a.environment.Oauth2Jwt.IssuerUri
	jwkSetUri := a.environment.Oauth2Jwt.JwkSetUri

	remoteKeySet := oidc.NewRemoteKeySet(ctx, jwkSetUri)

	oidcConfig := oidc.Config{SkipClientIDCheck: true}

	verifier := oidc.NewVerifier(issuerUri, remoteKeySet, &oidcConfig)
	token, err := verifier.Verify(ctx, idToken)

	if err != nil {
		fmt.Printf("oauth2: error when fetching token: %s,%s\n", err.Error(), token)
		a.handler.ServeHTTP(writer, request.WithContext(ctx))
		return
	}

	fmt.Print("\n Token verification success! \n")
	context.WithValue(ctx, AuthenticationContextKey, AuthContext{Token: token})
	a.handler.ServeHTTP(writer, request.WithContext(ctx))
}
