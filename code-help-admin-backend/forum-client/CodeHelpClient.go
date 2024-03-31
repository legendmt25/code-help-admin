package main

import (
	"context"
	"github.com/go-http-utils/headers"
	"net/http"

	codeHelpCoreGen "api-spec/generated/code-help-core"
)

func CreateClient() (*codeHelpCoreGen.ClientWithResponses, error) {
	return codeHelpCoreGen.NewClientWithResponses("http://localhost:3000",
		codeHelpCoreGen.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Add(headers.ContentType, "application/json")
			return nil
		}),
	)
}
