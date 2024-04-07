package admin

import (
	multipart_converter "admin-api/internal"
	codeHelpAdminCoreGen "api-spec/generated/code-help-admin-core"
	"context"
	"github.com/labstack/gommon/log"
	"net/http"
)

type ProblemCoreService interface {
	GetProblem(ctx context.Context, id codeHelpAdminCoreGen.ProblemId) (*codeHelpAdminCoreGen.ProblemDetail, error)

	GetProblems(ctx context.Context) ([]codeHelpAdminCoreGen.Problem, error)

	CreateProblem(ctx context.Context, req codeHelpAdminCoreGen.CreateProblemMultipartRequestBody) *codeHelpAdminCoreGen.ProblemDetail

	UpdateProblem(ctx context.Context, problemId codeHelpAdminCoreGen.ProblemId, req codeHelpAdminCoreGen.UpdateProblemMultipartRequestBody) *codeHelpAdminCoreGen.ProblemDetail

	DeleteProblem(ctx context.Context, problemId codeHelpAdminCoreGen.ProblemId) bool
}

type ProblemServiceContext struct {
}

type ProblemServiceWithContext struct {
	ctx ProblemServiceContext
}

type problemCoreServiceImpl struct {
	decoder ProblemDecoder
	client  codeHelpAdminCoreGen.ClientInterface
}

func NewCoreService(client codeHelpAdminCoreGen.ClientInterface, decoder ProblemDecoder) ProblemCoreService {
	return &problemCoreServiceImpl{
		client:  client,
		decoder: decoder,
	}
}

func (it *problemCoreServiceImpl) GetProblem(ctx context.Context, id codeHelpAdminCoreGen.ProblemId) (*codeHelpAdminCoreGen.ProblemDetail, error) {
	response, err := it.client.GetProblem(ctx, id)
	if err != nil {
		log.Errorf("Request Failed %s\n", err.Error())
		return nil, err
	}

	return it.decoder.DecodeDetail(response), nil
}

func (it *problemCoreServiceImpl) GetProblems(ctx context.Context) ([]codeHelpAdminCoreGen.Problem, error) {
	response, err := it.client.GetAllProblems(ctx)
	if err != nil {
		log.Errorf("Request Failed %s\n", err.Error())
		return []codeHelpAdminCoreGen.Problem{}, err
	}

	return it.decoder.DecodeAll(response), nil
}

func (it *problemCoreServiceImpl) CreateProblem(ctx context.Context, req codeHelpAdminCoreGen.CreateProblemMultipartRequestBody) *codeHelpAdminCoreGen.ProblemDetail {
	body, contentType := multipart_converter.EncodeMultipartFormData(req)
	response, err := it.client.CreateProblemWithBody(ctx, &codeHelpAdminCoreGen.CreateProblemParams{}, contentType, body)
	if err != nil {
		log.Error(err)
		return nil
	}

	return it.decoder.DecodeDetail(response)
}

func (it *problemCoreServiceImpl) UpdateProblem(ctx context.Context, problemId codeHelpAdminCoreGen.ProblemId, req codeHelpAdminCoreGen.UpdateProblemMultipartRequestBody) *codeHelpAdminCoreGen.ProblemDetail {
	body, contentType := multipart_converter.EncodeMultipartFormData(req)
	response, err := it.client.UpdateProblemWithBody(ctx, problemId, contentType, body)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return it.decoder.DecodeDetail(response)
}

func (it *problemCoreServiceImpl) DeleteProblem(ctx context.Context, problemId codeHelpAdminCoreGen.ProblemId) bool {
	problem, err := it.client.DeleteProblem(ctx, problemId)
	if err != nil {
		log.Error(err)
		return false
	}

	return problem.StatusCode == http.StatusNoContent
}
