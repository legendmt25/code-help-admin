package api

import (
	"admin-api/internal/admin"
	"encoding/json"
	"log"
	"net/http"

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

func (s *ServerInterfaceImpl) GetProblem(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ProblemId) {
	problem, _ := s.coreService.GetProblem(r.Context(), id)
	_ = json.NewEncoder(w).Encode(problem)
}

func (s *ServerInterfaceImpl) GetAllProblems(w http.ResponseWriter, r *http.Request) {
	problems, _ := s.coreService.GetProblems(r.Context())
	_ = json.NewEncoder(w).Encode(problems)
}

func (s *ServerInterfaceImpl) CreateProblem(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *ServerInterfaceImpl) DeleteProblem(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ProblemId) {
	s.coreService.DeleteProblem(r.Context(), id)
}

func (s *ServerInterfaceImpl) UpdateProblem(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ProblemId) {
	panic("implement me")
}

func (s *ServerInterfaceImpl) GetAllContests(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *ServerInterfaceImpl) CreateContest(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *ServerInterfaceImpl) DeleteContest(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ContestId) {
	panic("implement me")
}

func (s *ServerInterfaceImpl) GetContest(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ContestId) {
	panic("implement me")
}

func (s *ServerInterfaceImpl) UpdateContest(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ContestId) {
	panic("implement me")
}
