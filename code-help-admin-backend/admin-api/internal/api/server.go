package api

import (
	"admin-api/internal/admin"
	codeHelpAdminCoreGen "api-spec/generated/code-help-admin-core"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	codeHelpAdminGen "api-spec/generated/code-help-admin"
)

type AdminApiServer struct {
	server *http.Server
}

func (it AdminApiServer) Serve() {
	log.Println("Starting admin server...")
	log.Fatal(it.server.ListenAndServe())
}

type ServerInterfaceImpl struct {
	coreService admin.ProblemCoreService
}

func NewServiceInterfaceImpl(coreService admin.ProblemCoreService) codeHelpAdminGen.ServerInterface {
	return &ServerInterfaceImpl{coreService}
}

func (it *ServerInterfaceImpl) GetProblem(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ProblemId) {
	problem, _ := it.coreService.GetProblem(r.Context(), id)
	_ = json.NewEncoder(w).Encode(problem)
}

func (it *ServerInterfaceImpl) GetAllProblems(w http.ResponseWriter, r *http.Request) {
	problems, _ := it.coreService.GetProblems(r.Context())
	_ = json.NewEncoder(w).Encode(problems)
}

func (it *ServerInterfaceImpl) CreateProblem(w http.ResponseWriter, r *http.Request, params codeHelpAdminGen.CreateProblemParams) {
	var data codeHelpAdminGen.CreateProblemJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	response := it.coreService.CreateProblem(r.Context(), codeHelpAdminCoreGen.CreateProblemMultipartRequestBody{
		Category:    &codeHelpAdminCoreGen.Category{Name: data.Category.Name},
		Title:       data.Title,
		Difficulty:  codeHelpAdminCoreGen.Difficulty(data.Difficulty),
		Markdown:    data.Markdown,
		RunnerCode:  *createFileFromBytes(data.RunnerCode, "run.js"),
		StarterCode: *createFileFromBytes(data.StarterCode, "start.js"),
		TestCases:   toMultipart(data.TestCases),
	})

	_ = json.NewEncoder(w).Encode(codeHelpAdminGen.ProblemRequest{
		Category:    &codeHelpAdminGen.Category{Name: response.Category.Name},
		Difficulty:  codeHelpAdminGen.Difficulty(*response.Difficulty),
		Markdown:    response.Markdown,
		RunnerCode:  response.RunnerCode,
		StarterCode: response.StarterCode,
		TestCases:   response.TestCases,
		Title:       *response.Title,
	})
}

func (it *ServerInterfaceImpl) DeleteProblem(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ProblemId) {
	it.coreService.DeleteProblem(r.Context(), id)
}

func (it *ServerInterfaceImpl) UpdateProblem(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ProblemId) {
	var data codeHelpAdminGen.CreateProblemJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	response := it.coreService.UpdateProblem(r.Context(), id, codeHelpAdminCoreGen.UpdateProblemMultipartRequestBody{
		Category:    &codeHelpAdminCoreGen.Category{Name: data.Category.Name},
		Title:       data.Title,
		Difficulty:  codeHelpAdminCoreGen.Difficulty(data.Difficulty),
		Markdown:    data.Markdown,
		RunnerCode:  *createFileFromBytes(data.RunnerCode, "run.js"),
		StarterCode: *createFileFromBytes(data.StarterCode, "start.js"),
		TestCases:   toMultipart(data.TestCases),
	})

	_ = json.NewEncoder(w).Encode(codeHelpAdminGen.ProblemRequest{
		Category:    &codeHelpAdminGen.Category{Name: response.Category.Name},
		Difficulty:  codeHelpAdminGen.Difficulty(*response.Difficulty),
		Markdown:    response.Markdown,
		RunnerCode:  response.RunnerCode,
		StarterCode: response.StarterCode,
		TestCases:   response.TestCases,
		Title:       *response.Title,
	})
}

func (it *ServerInterfaceImpl) GetAllContests(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (it *ServerInterfaceImpl) CreateContest(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (it *ServerInterfaceImpl) DeleteContest(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ContestId) {
	panic("implement me")
}

func (it *ServerInterfaceImpl) GetContest(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ContestId) {
	panic("implement me")
}

func (it *ServerInterfaceImpl) UpdateContest(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ContestId) {
	panic("implement me")
}

func (it *ServerInterfaceImpl) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (it *ServerInterfaceImpl) CreateCategory(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (it *ServerInterfaceImpl) UpdateCategory(w http.ResponseWriter, r *http.Request, name codeHelpAdminGen.CategoryName) {
	//TODO implement me
	panic("implement me")
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
