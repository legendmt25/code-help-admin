package admin

import (
	codeHelpAdminCoreGen "api-spec/generated/code-help-admin-core"
	"net/http"
)

type CategoryDecoder interface {
	Decode(response *http.Response) *codeHelpAdminCoreGen.Category

	DecodeAll(response *http.Response) []codeHelpAdminCoreGen.Category
}

type categoryDecoderImpl struct{}

func NewCategoryDecoder() CategoryDecoder {
	return &categoryDecoderImpl{}
}

func (it *categoryDecoderImpl) Decode(response *http.Response) *codeHelpAdminCoreGen.Category {
	var category codeHelpAdminCoreGen.Category
	if err := decode(response, &category); err != nil {
		return nil
	}

	return &category
}

func (it *categoryDecoderImpl) DecodeAll(response *http.Response) []codeHelpAdminCoreGen.Category {
	var problems []codeHelpAdminCoreGen.Category
	err := decode(response, &problems)
	if err != nil {
		return []codeHelpAdminCoreGen.Category{}
	}

	return problems
}
