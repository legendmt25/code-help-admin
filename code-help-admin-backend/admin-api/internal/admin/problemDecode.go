package admin

import (
	"admin-api/internal/util/decodeUtil"
	codeHelpAdminCoreGen "api-spec/generated/code-help-admin-core"
	"net/http"
)

type ProblemDecoder interface {
	Decode(response *http.Response) *codeHelpAdminCoreGen.Problem

	DecodeDetail(response *http.Response) *codeHelpAdminCoreGen.ProblemDetail

	DecodeAll(response *http.Response) *codeHelpAdminCoreGen.ProblemResponse
}

type problemDecoderImpl struct{}

func NewProblemDecoder() ProblemDecoder {
	return &problemDecoderImpl{}
}

func (it *problemDecoderImpl) Decode(response *http.Response) *codeHelpAdminCoreGen.Problem {
	var problem codeHelpAdminCoreGen.Problem
	if err := decodeUtil.Decode(response, &problem); err != nil {
		return nil
	}

	return &problem
}

func (it *problemDecoderImpl) DecodeDetail(response *http.Response) *codeHelpAdminCoreGen.ProblemDetail {
	var problemDetail codeHelpAdminCoreGen.ProblemDetail
	if err := decodeUtil.Decode(response, &problemDetail); err != nil {
		return nil
	}

	return &problemDetail
}

func (it *problemDecoderImpl) DecodeAll(response *http.Response) *codeHelpAdminCoreGen.ProblemResponse {
	var problems codeHelpAdminCoreGen.ProblemResponse
	err := decodeUtil.Decode(response, &problems)
	if err != nil {
		return nil
	}

	return &problems
}
