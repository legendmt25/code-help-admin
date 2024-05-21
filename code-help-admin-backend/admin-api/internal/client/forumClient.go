package client

import (
	"admin-api/internal/api"
	codeHelpForum "api-spec/generated/code-help-forum"
	"context"
	"net/http"
)

func CreateForumClient(server string, oidcService api.OidcService) (codeHelpForum.ClientInterface, error) {
	return codeHelpForum.NewClient(server, codeHelpForum.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		return oidcService.AttachAuthHeader(ctx, req)
	}))
}
