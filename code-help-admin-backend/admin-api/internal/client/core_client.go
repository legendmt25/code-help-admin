package client

import (
	"admin-api/internal/api"
	codeHelpAdminCoreGen "api-spec/generated/code-help-admin-core"
	"context"
	"net/http"
)

func CreateAdminCoreClient(server string, oidcService api.OidcService) (codeHelpAdminCoreGen.ClientInterface, error) {
	return codeHelpAdminCoreGen.NewClient(server, codeHelpAdminCoreGen.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		return oidcService.AttachAuthHeader(ctx, req)
	}))
}
