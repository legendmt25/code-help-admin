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

func MapAllClientProblemToResponse(clientProblems []codeHelpAdminCoreGen.Problem) []codeHelpAdminGen.Problem {
	clientProblemSize := len(clientProblems)
	if clientProblemSize == 0 {
		return make([]codeHelpAdminGen.Problem, 0)
	}

	problems := make([]codeHelpAdminGen.Problem, clientProblemSize)
	for i, clientProblem := range clientProblems {
		problems[i] = mapClientProblemToResponse(clientProblem)
	}

	return problems
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
		Category:    mapCategory(data.Category),
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

	return &codeHelpAdminGen.Category{Name: category.Name}
}

func mapCategory(category *codeHelpAdminGen.Category) *codeHelpAdminCoreGen.Category {
	if category == nil {
		return nil
	}

	return &codeHelpAdminCoreGen.Category{Name: category.Name}
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
