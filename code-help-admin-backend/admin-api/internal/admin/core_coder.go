package admin

import (
	codeHelpAdminCoreGen "api-spec/generated/code-help-admin-core"
	"encoding/json"
	"net/http"
)

type ProblemDecoder interface {
	Decode(response *http.Response) *codeHelpAdminCoreGen.Problem

	DecodeDetail(response *http.Response) *codeHelpAdminCoreGen.ProblemDetail

	DecodeAll(response *http.Response) []codeHelpAdminCoreGen.Problem
}

type problemDecoderImpl struct {
}

func (it *problemDecoderImpl) DecodeAll(response *http.Response) []codeHelpAdminCoreGen.Problem {
	var problems []codeHelpAdminCoreGen.Problem
	err := decode(response, &problems)
	if err != nil {
		return []codeHelpAdminCoreGen.Problem{}
	}

	return problems
}

func NewProblemDecoder() ProblemDecoder {
	return &problemDecoderImpl{}
}

func (it *problemDecoderImpl) Decode(response *http.Response) *codeHelpAdminCoreGen.Problem {
	var problem codeHelpAdminCoreGen.Problem
	if err := decode(response, &problem); err != nil {
		return nil
	}

	return &problem
}

func (it *problemDecoderImpl) DecodeDetail(response *http.Response) *codeHelpAdminCoreGen.ProblemDetail {
	var problemDetail codeHelpAdminCoreGen.ProblemDetail
	if err := decode(response, &problemDetail); err != nil {
		return nil
	}

	return &problemDetail
}

func decode[T any](response *http.Response, output *T) error {
	defer func() { _ = response.Body.Close() }()
	err := json.NewDecoder(response.Body).Decode(&output)
	if err != nil {
		return err
	}

	return nil
}
