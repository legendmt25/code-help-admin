package codehelpcore

import (
	"context"
	"errors"
	"net/http"

	codehelpcoreapi "mk.com.code-help/admin-api-docs/api/code-help-core/go"
)

type CategoryController struct{}

func (s CategoryController) GetCategories(ctx context.Context) (codehelpcoreapi.ImplResponse, error) {

	return codehelpcoreapi.Response(http.StatusNotImplemented, nil), errors.New("GetCategories method not implemented")
}
