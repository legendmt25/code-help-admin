package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
)

type AuthMiddleware struct {
	handler http.Handler
}

func NewAuthMiddleWare(handler http.Handler) http.Handler {
	return AuthMiddleware{handler: handler}
}

type AuthCtx struct {
	St string
}

func (a AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := context.WithValue(request.Context(), "", AuthCtx{})

	authorizationHeaderValues := request.Header["Authorization"]

	if len(authorizationHeaderValues) == 0 || authorizationHeaderValues[0] == "" {
		a.handler.ServeHTTP(writer, request.WithContext(ctx))
		return
	}

	idToken := strings.Replace(authorizationHeaderValues[0], "Bearer ", "", 1)

	issuerUri := os.Getenv("OAUTH2_RESOURCESERVER_JWT_ISSUERURI")
	jwtSetUri := os.Getenv("OAUTH2_RESOURCESERVER_JWT_JWKSETURI")

	remoteKeySet := oidc.NewRemoteKeySet(ctx, jwtSetUri)

	oidcConfig := oidc.Config{SkipClientIDCheck: true}

	verifier := oidc.NewVerifier(issuerUri, remoteKeySet, &oidcConfig)
	token, err := verifier.Verify(ctx, idToken)

	if err != nil {
		fmt.Printf("oauth2: error when fetching token: %s,%s\n", err.Error(), token)
		a.handler.ServeHTTP(writer, request.WithContext(ctx))
	} else {
		fmt.Print("\n Token verification success! \n")
	}

	a.handler.ServeHTTP(writer, request.WithContext(ctx))
}
