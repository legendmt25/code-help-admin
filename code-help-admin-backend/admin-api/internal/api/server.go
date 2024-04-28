package api

import (
	"admin-api/internal/admin"
	api_mapper "admin-api/internal/api/mapper"
	codeHelpAdminGen "api-spec/generated/code-help-admin"
	codeHelpAdminCoreGen "api-spec/generated/code-help-admin-core"
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

	if response := it.coreService.CreateProblem(r.Context(), params.ContestId, api_mapper.MapToClientCreateReq(data)); response == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
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

func (it *ServerInterfaceImpl) DeleteProblem(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ProblemId) {
	it.coreService.DeleteProblem(r.Context(), id)
}

func (it *ServerInterfaceImpl) GetAllContests(w http.ResponseWriter, r *http.Request) {
	contests, _ := it.coreService.GetContests(r.Context())
	_ = json.NewEncoder(w).Encode(api_mapper.MapAllClientContestToResponse(contests))
}

func (it *ServerInterfaceImpl) GetContest(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ContestId) {
	contest := it.coreService.GetContest(r.Context(), id)
	if contest == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(api_mapper.MapClientContestDetailToResponse(*contest))
}

func (it *ServerInterfaceImpl) CreateContest(w http.ResponseWriter, r *http.Request) {
	var data = codeHelpAdminGen.CreateContestJSONRequestBody{}
	_ = json.NewDecoder(r.Body).Decode(&data)

	res := it.coreService.CreateContest(r.Context(), codeHelpAdminCoreGen.CreateContestJSONRequestBody{
		Duration: data.Duration,
		Name:     data.Name,
		StartsOn: data.StartsOn,
		Status:   (*codeHelpAdminCoreGen.ContestStatus)(data.Status),
	})

	if res {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (it *ServerInterfaceImpl) UpdateContest(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ContestId) {
	var data = codeHelpAdminGen.UpdateContestJSONRequestBody{}
	_ = json.NewDecoder(r.Body).Decode(&data)

	res := it.coreService.UpdateContest(r.Context(), id, codeHelpAdminCoreGen.UpdateContestJSONRequestBody{
		Duration: data.Duration,
		Name:     data.Name,
		StartsOn: data.StartsOn,
		Status:   (*codeHelpAdminCoreGen.ContestStatus)(data.Status),
	})

	if res {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (it *ServerInterfaceImpl) DeleteContest(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ContestId) {
	it.coreService.DeleteContest(r.Context(), id)
	w.WriteHeader(http.StatusNoContent)
}

func (it *ServerInterfaceImpl) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := it.coreService.GetAllCategories(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(api_mapper.MapAllCategoryToResponse(categories))
}

func (it *ServerInterfaceImpl) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var data codeHelpAdminGen.CreateCategoryJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	res := it.coreService.CreateCategory(r.Context(), codeHelpAdminCoreGen.CreateCategoryJSONRequestBody{
		Name: data.Name,
	})

	if res {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (it *ServerInterfaceImpl) UpdateCategory(w http.ResponseWriter, r *http.Request, name codeHelpAdminGen.CategoryName) {

	res := it.coreService.UpdateCategory(r.Context(), name, codeHelpAdminCoreGen.CreateCategoryJSONRequestBody{
		Name: name,
	})

	if res {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
