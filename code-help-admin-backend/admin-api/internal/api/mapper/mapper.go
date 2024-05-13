package api_mapper

import (
	codeHelpAdminGen "api-spec/generated/code-help-admin"
	codeHelpAdminCoreGen "api-spec/generated/code-help-admin-core"
	"strconv"
)

func MapClientProblemDetailToResponse(clientProblem codeHelpAdminCoreGen.ProblemDetail) codeHelpAdminGen.ProblemDetail {
	return codeHelpAdminGen.ProblemDetail{
		Category:    mapCategoryToResponse(clientProblem.Category),
		Difficulty:  codeHelpAdminGen.Difficulty(clientProblem.Difficulty),
		Id:          clientProblem.Id,
		Markdown:    clientProblem.Markdown,
		RunnerCode:  clientProblem.RunnerCode,
		StarterCode: clientProblem.StarterCode,
		TestCases:   clientProblem.TestCases,
		Title:       clientProblem.Title,
	}
}

func MapAllClientProblemToResponse(clientProblems *codeHelpAdminCoreGen.ProblemResponse) *codeHelpAdminGen.ProblemResponse {
	if clientProblems == nil || len(clientProblems.Problems) <= 0 {
		return &codeHelpAdminGen.ProblemResponse{
			Problems: make([]codeHelpAdminGen.Problem, 0),
		}
	}

	problems := make([]codeHelpAdminGen.Problem, len(clientProblems.Problems))
	for i, clientProblem := range clientProblems.Problems {
		problems[i] = mapClientProblemToResponse(clientProblem)
	}

	return &codeHelpAdminGen.ProblemResponse{
		Problems: problems,
	}
}

func mapClientProblemToResponse(clientProblem codeHelpAdminCoreGen.Problem) codeHelpAdminGen.Problem {
	return codeHelpAdminGen.Problem{
		Category:   mapCategoryToResponse(clientProblem.Category),
		Difficulty: codeHelpAdminGen.Difficulty(clientProblem.Difficulty),
		Id:         clientProblem.Id,
		Title:      clientProblem.Title,
	}
}

func MapClientResponseToResponse(response codeHelpAdminCoreGen.ProblemDetail) codeHelpAdminGen.ProblemDetail {
	return codeHelpAdminGen.ProblemDetail{
		Id:          response.Id,
		Title:       response.Title,
		Category:    mapCategoryToResponse(response.Category),
		Difficulty:  codeHelpAdminGen.Difficulty(response.Difficulty),
		Markdown:    response.Markdown,
		RunnerCode:  response.RunnerCode,
		StarterCode: response.StarterCode,
		TestCases:   response.TestCases,
	}
}

func MapToClientCreateReq(data codeHelpAdminGen.CreateProblemJSONRequestBody) codeHelpAdminCoreGen.CreateProblemMultipartRequestBody {
	return mapToClientReq(data)
}

func MapToClientUpdateReq(data codeHelpAdminGen.UpdateProblemJSONRequestBody) codeHelpAdminCoreGen.UpdateProblemMultipartRequestBody {
	return mapToClientReq(data)
}

func mapToClientReq(data codeHelpAdminGen.CreateProblemJSONRequestBody) codeHelpAdminCoreGen.ProblemRequest {
	return codeHelpAdminCoreGen.ProblemRequest{
		Category:    mapCategoryRequest(data.Category),
		Title:       data.Title,
		Difficulty:  mapDifficulty(data.Difficulty),
		Markdown:    data.Markdown,
		RunnerCode:  *createFileFromBytes(data.RunnerCode, "run.js"),
		StarterCode: *createFileFromBytes(data.StarterCode, "start.js"),
		TestCases:   toMultipart(data.TestCases),
	}
}

func mapCategoryToResponse(category *codeHelpAdminCoreGen.Category) *codeHelpAdminGen.Category {
	if category == nil {
		return nil
	}

	return &codeHelpAdminGen.Category{Id: category.Id, Name: category.Name}
}

func mapCategory(category *codeHelpAdminGen.Category) *codeHelpAdminCoreGen.Category {
	if category == nil {
		return nil
	}

	return &codeHelpAdminCoreGen.Category{Id: category.Id, Name: category.Name}
}

func mapCategoryRequest(category *codeHelpAdminGen.CategoryRequest) *codeHelpAdminCoreGen.CategoryRequest {
	if category == nil {
		return nil
	}

	return &codeHelpAdminCoreGen.CategoryRequest{Name: category.Name}
}

func mapDifficulty(difficulty codeHelpAdminGen.Difficulty) codeHelpAdminCoreGen.Difficulty {
	return codeHelpAdminCoreGen.Difficulty(difficulty)
}

func createFileFromBytes[T string](data T, filename string) *codeHelpAdminCoreGen.File {
	file := &codeHelpAdminCoreGen.File{}
	file.InitFromBytes([]byte(data), filename)
	return file
}

func toMultipart(cases codeHelpAdminGen.TestCases) []codeHelpAdminCoreGen.File {
	var files []codeHelpAdminCoreGen.File

	for i, c := range cases {
		if inFileData := c.In; inFileData != nil {
			files = append(files, *createFileFromBytes(*inFileData, "IN"+strconv.Itoa(i)))
		}
		if outFileData := c.In; outFileData != nil {
			files = append(files, *createFileFromBytes(*outFileData, "OUT"+strconv.Itoa(i)))
		}
	}

	return files
}

func MapClientContestDetailToResponse(clientContest codeHelpAdminCoreGen.ContestDetail) codeHelpAdminGen.ContestDetail {
	return codeHelpAdminGen.ContestDetail{
		Duration: clientContest.Duration,
		Id:       clientContest.Id,
		Name:     clientContest.Name,
		Problems: mapContestProblemsToResponse(clientContest.Problems),
		StartsOn: clientContest.StartsOn,
		Status:   codeHelpAdminGen.ContestStatus(clientContest.Status),
	}
}

func mapContestProblemsToResponse(clientProblem []codeHelpAdminCoreGen.ContestProblem) []codeHelpAdminGen.ContestProblem {
	clientContestSize := len(clientProblem)
	if clientContestSize == 0 {
		return make([]codeHelpAdminGen.ContestProblem, 0)
	}

	problems := make([]codeHelpAdminGen.ContestProblem, clientContestSize)
	for i, clientContest := range clientProblem {
		problems[i] = mapContestProblemToResponse(clientContest)
	}

	return problems
}

func mapContestProblemToResponse(problem codeHelpAdminCoreGen.ContestProblem) codeHelpAdminGen.ContestProblem {
	return codeHelpAdminGen.ContestProblem{
		Id:               problem.Id,
		Title:            problem.Title,
		ContestProblemId: problem.ContestProblemId,
		Category:         mapCategoryToResponse(problem.Category),
		Difficulty:       codeHelpAdminGen.Difficulty(problem.Difficulty),
		Score:            problem.Score,
	}
}

func MapAllClientContestToResponse(clientContests *codeHelpAdminCoreGen.ContestResponse) *codeHelpAdminGen.ContestResponse {
	if clientContests == nil || len(clientContests.Contests) <= 0 {
		return &codeHelpAdminGen.ContestResponse{
			Contests: make([]codeHelpAdminGen.Contest, 0),
		}
	}

	contests := make([]codeHelpAdminGen.Contest, len(clientContests.Contests))
	for i, clientContest := range clientContests.Contests {
		contests[i] = mapClientContestToResponse(clientContest)
	}

	return &codeHelpAdminGen.ContestResponse{
		Contests: contests,
	}
}

func mapClientContestToResponse(clientContest codeHelpAdminCoreGen.Contest) codeHelpAdminGen.Contest {
	return codeHelpAdminGen.Contest{
		Id:       clientContest.Id,
		Name:     clientContest.Name,
		Duration: clientContest.Duration,
		StartsOn: clientContest.StartsOn,
		Status:   codeHelpAdminGen.ContestStatus(clientContest.Status),
	}
}

func MapAllCategoryToResponse(clientCategories *codeHelpAdminCoreGen.CategoryResponse) *codeHelpAdminGen.CategoryResponse {
	if clientCategories == nil || len(clientCategories.Categories) <= 0 {
		return &codeHelpAdminGen.CategoryResponse{
			Categories: make([]codeHelpAdminGen.Category, 0),
		}
	}

	categories := make([]codeHelpAdminGen.Category, len(clientCategories.Categories))
	for i, clientCategory := range clientCategories.Categories {
		categories[i] = *mapCategoryToResponse(&clientCategory)
	}

	return &codeHelpAdminGen.CategoryResponse{
		Categories: categories,
	}
}
