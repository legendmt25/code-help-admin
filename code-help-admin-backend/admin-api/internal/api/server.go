package api

import (
	"admin-api/internal/admin"
	api_mapper "admin-api/internal/api/mapper"
	codeHelpAdminGen "api-spec/generated/code-help-admin"
	"encoding/json"
	"log"
	"net/http"
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
	if problem == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(api_mapper.MapClientProblemDetailToResponse(*problem))
}

func (it *ServerInterfaceImpl) GetAllProblems(w http.ResponseWriter, r *http.Request) {
	problems, _ := it.coreService.GetProblems(r.Context())
	_ = json.NewEncoder(w).Encode(api_mapper.MapAllClientProblemToResponse(problems))
}

func (it *ServerInterfaceImpl) CreateProblem(w http.ResponseWriter, r *http.Request, params codeHelpAdminGen.CreateProblemParams) {
	var data codeHelpAdminGen.CreateProblemJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	if response := it.coreService.CreateProblem(r.Context(), api_mapper.MapToClientCreateReq(data)); response == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (it *ServerInterfaceImpl) DeleteProblem(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ProblemId) {
	it.coreService.DeleteProblem(r.Context(), id)
}

func (it *ServerInterfaceImpl) UpdateProblem(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ProblemId) {
	var data codeHelpAdminGen.CreateProblemJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	response := it.coreService.UpdateProblem(r.Context(), id, api_mapper.MapToClientUpdateReq(data))
	if response == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(api_mapper.MapClientResponseToResponse(*response))
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
