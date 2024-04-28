package admin

import (
	multipart_converter "admin-api/internal"
	codeHelpAdminCoreGen "api-spec/generated/code-help-admin-core"
	"bytes"
	"context"
	"github.com/labstack/gommon/log"
	"net/http"
)

type ProblemCoreService interface {
	GetProblem(ctx context.Context, id codeHelpAdminCoreGen.ProblemId) (*codeHelpAdminCoreGen.ProblemDetail, error)

	GetProblems(ctx context.Context) ([]codeHelpAdminCoreGen.Problem, error)

	CreateProblem(ctx context.Context, contestId *codeHelpAdminCoreGen.ContestId, req codeHelpAdminCoreGen.CreateProblemMultipartRequestBody) error

	UpdateProblem(ctx context.Context, problemId codeHelpAdminCoreGen.ProblemId, req codeHelpAdminCoreGen.UpdateProblemMultipartRequestBody) *codeHelpAdminCoreGen.ProblemDetail

	DeleteProblem(ctx context.Context, problemId codeHelpAdminCoreGen.ProblemId) bool

	GetContest(ctx context.Context, contestId codeHelpAdminCoreGen.ContestId) *codeHelpAdminCoreGen.ContestDetail

	GetContests(ctx context.Context) ([]codeHelpAdminCoreGen.Contest, error)

	CreateContest(ctx context.Context, body codeHelpAdminCoreGen.CreateContestJSONRequestBody) bool

	UpdateContest(ctx context.Context, contestId codeHelpAdminCoreGen.ContestId, body codeHelpAdminCoreGen.UpdateContestJSONRequestBody) bool

	DeleteContest(ctx context.Context, contestId codeHelpAdminCoreGen.ContestId) bool

	GetAllCategories(ctx context.Context) ([]codeHelpAdminCoreGen.Category, error)

	CreateCategory(ctx context.Context, category codeHelpAdminCoreGen.CreateCategoryJSONRequestBody) bool

	UpdateCategory(ctx context.Context, name codeHelpAdminCoreGen.CategoryName, body codeHelpAdminCoreGen.CreateCategoryJSONRequestBody) bool

	DeleteCategory(ctx context.Context, category codeHelpAdminCoreGen.CategoryName) bool
}

type ProblemServiceContext struct {
}

type ProblemServiceWithContext struct {
	ctx ProblemServiceContext
}

type problemCoreServiceImpl struct {
	client codeHelpAdminCoreGen.ClientInterface

	problemDecoder  ProblemDecoder
	contestDecoder  ContestDecoder
	categoryDecoder CategoryDecoder
}

func NewCoreService(client codeHelpAdminCoreGen.ClientInterface) ProblemCoreService {
	return &problemCoreServiceImpl{
		client:          client,
		problemDecoder:  NewProblemDecoder(),
		contestDecoder:  NewContestDecoder(),
		categoryDecoder: NewCategoryDecoder(),
	}
}

func (it *problemCoreServiceImpl) GetProblem(ctx context.Context, id codeHelpAdminCoreGen.ProblemId) (*codeHelpAdminCoreGen.ProblemDetail, error) {
	response, err := it.client.GetProblem(ctx, id)
	if err != nil {
		log.Errorf("Request Failed %s\n", err.Error())
		return nil, err
	}

	return it.problemDecoder.DecodeDetail(response), nil
}

func (it *problemCoreServiceImpl) GetProblems(ctx context.Context) ([]codeHelpAdminCoreGen.Problem, error) {
	response, err := it.client.GetAllProblems(ctx)
	if err != nil {
		log.Errorf("Request Failed %s\n", err.Error())
		return []codeHelpAdminCoreGen.Problem{}, err
	}

	return it.problemDecoder.DecodeAll(response), nil
}

func (it *problemCoreServiceImpl) CreateProblem(ctx context.Context, contestId *codeHelpAdminCoreGen.ContestId, req codeHelpAdminCoreGen.CreateProblemMultipartRequestBody) error {
	var body = &bytes.Buffer{}
	_, contentType := multipart_converter.EncodeMultipartFormData(req, body)
	_, err := it.client.CreateProblemWithBody(ctx, &codeHelpAdminCoreGen.CreateProblemParams{
		ContestId: contestId,
	}, contentType, body)
	if err != nil {
		log.Error("Error: ", err)
		return nil
	}

	return err
}

func (it *problemCoreServiceImpl) UpdateProblem(ctx context.Context, problemId codeHelpAdminCoreGen.ProblemId, req codeHelpAdminCoreGen.UpdateProblemMultipartRequestBody) *codeHelpAdminCoreGen.ProblemDetail {
	var body = &bytes.Buffer{}
	_, contentType := multipart_converter.EncodeMultipartFormData(req, body)
	response, err := it.client.UpdateProblemWithBody(ctx, problemId, contentType, body)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return it.problemDecoder.DecodeDetail(response)
}

func (it *problemCoreServiceImpl) DeleteProblem(ctx context.Context, problemId codeHelpAdminCoreGen.ProblemId) bool {
	problem, err := it.client.DeleteProblem(ctx, problemId)
	if err != nil {
		log.Error(err)
		return false
	}

	return problem.StatusCode == http.StatusNoContent
}

func (it *problemCoreServiceImpl) GetContest(ctx context.Context, contestId codeHelpAdminCoreGen.ContestId) *codeHelpAdminCoreGen.ContestDetail {
	contest, err := it.client.GetContest(ctx, contestId)
	if err != nil {
		log.Error(err)
		return nil
	}

	return it.contestDecoder.DecodeDetail(contest)
}

func (it *problemCoreServiceImpl) GetContests(ctx context.Context) ([]codeHelpAdminCoreGen.Contest, error) {
	contests, err := it.client.GetAllContests(ctx)
	if err != nil {
		log.Error(err)
		return []codeHelpAdminCoreGen.Contest{}, err
	}

	return it.contestDecoder.DecodeAll(contests), nil
}

func (it *problemCoreServiceImpl) CreateContest(ctx context.Context, body codeHelpAdminCoreGen.CreateContestJSONRequestBody) bool {
	response, err := it.client.CreateContest(ctx, body)
	if err != nil {
		log.Error(err)
		return false
	}
	switch response.StatusCode {
	case http.StatusOK:
		return true
	case http.StatusUnauthorized, http.StatusForbidden:
		return false
	}

	return false
}

func (it *problemCoreServiceImpl) UpdateContest(ctx context.Context, id codeHelpAdminCoreGen.ContestId, body codeHelpAdminCoreGen.UpdateContestJSONRequestBody) bool {
	response, err := it.client.UpdateContest(ctx, id, body)
	if err != nil {
		log.Error(err)
		return false
	}

	switch response.StatusCode {
	case http.StatusNoContent:
		return true
	case http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound:
		return false
	}

	return false
}

func (it *problemCoreServiceImpl) DeleteContest(ctx context.Context, contestId codeHelpAdminCoreGen.ContestId) bool {
	response, err := it.client.DeleteContest(ctx, contestId)
	if err != nil {
		log.Error(err)
		return false
	}

	switch response.StatusCode {
	case http.StatusNoContent:
		return true
	case http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound:
		return false
	}

	return false
}

func (it *problemCoreServiceImpl) GetAllCategories(ctx context.Context) ([]codeHelpAdminCoreGen.Category, error) {
	response, err := it.client.GetAllCategories(ctx)
	if err != nil {
		log.Error(err)
		return []codeHelpAdminCoreGen.Category{}, err
	}

	return it.categoryDecoder.DecodeAll(response), nil
}

func (it *problemCoreServiceImpl) CreateCategory(ctx context.Context, body codeHelpAdminCoreGen.CreateCategoryJSONRequestBody) bool {
	res, err := it.client.CreateCategory(ctx, body)
	if err != nil {
		log.Error(err)
		return false
	}

	switch res.StatusCode {
	case http.StatusCreated:
		return true
	default:
		return false
	}
}

func (it *problemCoreServiceImpl) UpdateCategory(ctx context.Context, name codeHelpAdminCoreGen.CategoryName, body codeHelpAdminCoreGen.CreateCategoryJSONRequestBody) bool {
	res, err := it.client.UpdateCategory(ctx, name, body)
	if err != nil {
		log.Error(err)
		return false
	}

	switch res.StatusCode {
	case http.StatusNoContent:
		return true
	default:
		return false
	}
}

func (it *problemCoreServiceImpl) DeleteCategory(ctx context.Context, category codeHelpAdminCoreGen.CategoryName) bool {
	res, err := it.client.DeleteCategory(ctx, category)
	if err != nil {
		log.Error(err)
		return false
	}

	switch res.StatusCode {
	case http.StatusNoContent:
		return true
	case http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound:
		return false
	}

	return false
}
