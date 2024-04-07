package api

import (
	"context"
	"net/http"
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
	//request.Header.Add("Authorization", "Bearer token")
	ctx := context.WithValue(request.Context(), "", AuthCtx{})
	a.handler.ServeHTTP(writer, request.WithContext(ctx))
}
