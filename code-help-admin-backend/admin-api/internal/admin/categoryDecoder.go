package admin

import (
	"admin-api/internal/util/decodeUtil"
	codeHelpAdminCoreGen "api-spec/generated/code-help-admin-core"
	"net/http"
)

type CategoryDecoder interface {
	Decode(response *http.Response) *codeHelpAdminCoreGen.Category

	DecodeAll(response *http.Response) *codeHelpAdminCoreGen.CategoryResponse
}

type categoryDecoderImpl struct{}

func NewCategoryDecoder() CategoryDecoder {
	return &categoryDecoderImpl{}
}

func (it *categoryDecoderImpl) Decode(response *http.Response) *codeHelpAdminCoreGen.Category {
	var category codeHelpAdminCoreGen.Category
	if err := decodeUtil.Decode(response, &category); err != nil {
		return nil
	}

	return &category
}

func (it *categoryDecoderImpl) DecodeAll(response *http.Response) *codeHelpAdminCoreGen.CategoryResponse {
	var problems codeHelpAdminCoreGen.CategoryResponse
	err := decodeUtil.Decode(response, &problems)
	if err != nil {
		return nil
	}

	return &problems
}
