package admin

import (
	codeHelpAdminCoreGen "api-spec/generated/code-help-admin-core"
	"net/http"
)

type ContestDecoder interface {
	Decode(response *http.Response) *codeHelpAdminCoreGen.Contest

	DecodeDetail(response *http.Response) *codeHelpAdminCoreGen.ContestDetail

	DecodeAll(response *http.Response) []codeHelpAdminCoreGen.Contest
}

type contestDecoderImpl struct{}

func NewContestDecoder() ContestDecoder {
	return &contestDecoderImpl{}
}

func (it *contestDecoderImpl) Decode(response *http.Response) *codeHelpAdminCoreGen.Contest {
	var problem codeHelpAdminCoreGen.Contest
	if err := decode(response, &problem); err != nil {
		return nil
	}

	return &problem
}

func (it *contestDecoderImpl) DecodeDetail(response *http.Response) *codeHelpAdminCoreGen.ContestDetail {
	var problemDetail codeHelpAdminCoreGen.ContestDetail
	if err := decode(response, &problemDetail); err != nil {
		return nil
	}

	return &problemDetail
}

func (it *contestDecoderImpl) DecodeAll(response *http.Response) []codeHelpAdminCoreGen.Contest {
	var problems []codeHelpAdminCoreGen.Contest
	err := decode(response, &problems)
	if err != nil {
		return []codeHelpAdminCoreGen.Contest{}
	}

	return problems
}
