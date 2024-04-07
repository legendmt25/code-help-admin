package api

import (
	"net/http"

	codeHelpAdminGen "api-spec/generated/code-help-admin"
)

func NewAdminApiServer(service codeHelpAdminGen.ServerInterface) AdminApiServer {
	handler := codeHelpAdminGen.HandlerWithOptions(service, codeHelpAdminGen.StdHTTPServerOptions{
		BaseURL:     resolvePath(),
		Middlewares: []codeHelpAdminGen.MiddlewareFunc{NewAuthMiddleWare},
	})

	server := http.Server{
		Addr:    ":4000",
		Handler: handler,
	}

	return AdminApiServer{&server}
}

func resolvePath() string {
	swagger, err := codeHelpAdminGen.GetSwagger()
	if err != nil {
		return ""
	}

	s := swagger.Servers[0]

	path, err := s.BasePath()
	if err != nil {
		return ""
	}

	return path
}
