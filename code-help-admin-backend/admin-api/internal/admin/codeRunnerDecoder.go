package admin

import (
	"admin-api/internal/util/decodeUtil"
	codeHelpAdminCoreGen "api-spec/generated/code-help-admin-core"
	"net/http"
)

type CodeRunnerDecoder interface {
	Decode(response *http.Response) *codeHelpAdminCoreGen.CodeRunnerResponse
}

type codeRunnerDecoderImpl struct{}

func NewCodeRunnerDecoder() CodeRunnerDecoder {
	return &codeRunnerDecoderImpl{}
}

func (it *codeRunnerDecoderImpl) Decode(response *http.Response) *codeHelpAdminCoreGen.CodeRunnerResponse {
	var codeRunnerResponse codeHelpAdminCoreGen.CodeRunnerResponse
	if err := decodeUtil.Decode(response, &codeRunnerResponse); err != nil {
		return nil
	}

	return &codeRunnerResponse
}
