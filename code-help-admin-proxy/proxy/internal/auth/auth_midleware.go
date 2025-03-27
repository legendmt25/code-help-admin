package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
)

type AuthMiddleware struct {
	handler     http.Handler
	oidcService OidcService
}

func NewAuthMiddlewareWith(oidcService OidcService) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return AuthMiddleware{handler: handler, oidcService: oidcService}
	}
}

const (
	AuthenticationContextKey = "Authentication"
)

type AuthContext struct {
	IdToken *oidc.IDToken
	Token   string
}

func (a AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := context.WithValue(request.Context(), AuthenticationContextKey, AuthContext{})

	authorizationHeaderValue := request.Header.Get("Authorization")

	if len(authorizationHeaderValue) <= 0 {
		a.handler.ServeHTTP(writer, request)
		return
	}

	token := strings.TrimSpace(strings.Split(authorizationHeaderValue, " ")[1])

	idToken, err := a.oidcService.Verify(ctx, token)

	if err != nil {
		fmt.Printf("oauth2: error when fetching token: %s,%s\n", err.Error(), idToken)
		a.handler.ServeHTTP(writer, request.WithContext(ctx))
		return
	}

	newContext := context.WithValue(ctx, AuthenticationContextKey, AuthContext{Token: token, IdToken: idToken})
	a.handler.ServeHTTP(writer, request.WithContext(newContext))
}
