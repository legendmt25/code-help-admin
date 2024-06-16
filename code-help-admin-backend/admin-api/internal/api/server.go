package api

import (
	"admin-api/internal/admin"
	"admin-api/internal/admin/forum"
	apimapper "admin-api/internal/api/mapper"
	codeHelpAdminGen "api-spec/generated/code-help-admin"
	codeHelpAdminCoreGen "api-spec/generated/code-help-admin-core"
	codeHelpForumGen "api-spec/generated/code-help-forum"
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
	coreService  admin.ProblemCoreService
	forumService forum.Service
}

func NewServiceInterfaceImpl(coreService admin.ProblemCoreService, forumService forum.Service) codeHelpAdminGen.ServerInterface {
	return &ServerInterfaceImpl{
		coreService:  coreService,
		forumService: forumService,
	}
}

func (it *ServerInterfaceImpl) RunCode(w http.ResponseWriter, r *http.Request) {
	var data codeHelpAdminGen.RunCodeJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	response, statusCode, _ := it.coreService.RunCode(r.Context(), apimapper.MapToClientRunCodeRequest(data))

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(apimapper.MapClientRunCodeResponseToResponse(*response))
}

func (it *ServerInterfaceImpl) GetProblem(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ProblemId) {
	problem, _ := it.coreService.GetProblem(r.Context(), id)
	if problem == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(apimapper.MapClientProblemDetailToResponse(*problem))
}

func (it *ServerInterfaceImpl) GetAllProblems(w http.ResponseWriter, r *http.Request) {
	problems, statusCode, _ := it.coreService.GetProblems(r.Context())

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(apimapper.MapAllClientProblemToResponse(problems))
}

func (it *ServerInterfaceImpl) CreateProblem(w http.ResponseWriter, r *http.Request, params codeHelpAdminGen.CreateProblemParams) {
	var data codeHelpAdminGen.CreateProblemJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	statusCode, _ := it.coreService.CreateProblem(r.Context(), params.ContestId, apimapper.MapToClientCreateReq(data))

	w.WriteHeader(statusCode)
}

func (it *ServerInterfaceImpl) UpdateProblem(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ProblemId) {
	var data codeHelpAdminGen.CreateProblemJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	response := it.coreService.UpdateProblem(r.Context(), id, apimapper.MapToClientUpdateReq(data))
	if response == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(apimapper.MapClientResponseToResponse(*response))
}

func (it *ServerInterfaceImpl) DeleteProblem(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ProblemId) {
	it.coreService.DeleteProblem(r.Context(), id)
}

func (it *ServerInterfaceImpl) GetAllContests(w http.ResponseWriter, r *http.Request) {
	contests, _ := it.coreService.GetContests(r.Context())
	_ = json.NewEncoder(w).Encode(apimapper.MapAllClientContestToResponse(contests))
}

func (it *ServerInterfaceImpl) GetContest(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.ContestId) {
	contest := it.coreService.GetContest(r.Context(), id)
	if contest == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(apimapper.MapClientContestDetailToResponse(*contest))
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

	res := it.coreService.UpdateContest(r.Context(), id, apimapper.MapContestEditRequestToClient(data))

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
	categories, statusCode, err := it.coreService.GetAllCategories(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(apimapper.MapAllCategoryToResponse(categories))
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

func (it *ServerInterfaceImpl) UpdateCategory(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.CategoryId) {

	var data codeHelpAdminGen.UpdateCategoryJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	res := it.coreService.UpdateCategory(r.Context(), id, codeHelpAdminCoreGen.UpdateCategoryJSONRequestBody{
		Name: data.Name,
	})

	if res {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (it *ServerInterfaceImpl) DeleteCategory(w http.ResponseWriter, r *http.Request, id codeHelpAdminGen.CategoryId) {

	res := it.coreService.DeleteCategory(r.Context(), id)

	if res {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (it *ServerInterfaceImpl) CreateForumCategory(w http.ResponseWriter, r *http.Request) {
	var data codeHelpAdminGen.CreateForumCategoryJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	statusCode, _ := it.forumService.Category().CreateCategory(r.Context(), codeHelpForumGen.CreateCategoryJSONRequestBody{
		Name: data.Name,
	})

	w.WriteHeader(statusCode)
}

func (it *ServerInterfaceImpl) DeleteForumCategory(w http.ResponseWriter, r *http.Request, uid string) {
	statusCode, _ := it.forumService.Category().DeleteCategory(r.Context(), uid)

	w.WriteHeader(statusCode)
}

func (it *ServerInterfaceImpl) UpdateForumCategory(w http.ResponseWriter, r *http.Request, uid string) {
	var data codeHelpAdminGen.UpdateForumCategoryJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	res, statusCode, _ := it.forumService.Category().UpdateCategory(r.Context(), uid, codeHelpForumGen.UpdateCategoryJSONRequestBody{
		Name: data.Name,
	})

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func (it *ServerInterfaceImpl) GetCommentsForPost(w http.ResponseWriter, r *http.Request, params codeHelpAdminGen.GetCommentsForPostParams) {
	res, statusCode, _ := it.forumService.Comment().GetComments(r.Context(), codeHelpForumGen.GetCommentsForPostParams{
		Post: params.Post,
	})

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func (it *ServerInterfaceImpl) CommentOnPost(w http.ResponseWriter, r *http.Request, params codeHelpAdminGen.CommentOnPostParams) {
	var data codeHelpAdminGen.CommentOnPostJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	res, statusCode, _ := it.forumService.Comment().CreateComment(r.Context(), codeHelpForumGen.CommentOnPostParams{
		Post: params.Post,
	}, codeHelpForumGen.CommentOnPostJSONRequestBody{
		Content: data.Content,
	})

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func (it *ServerInterfaceImpl) DeleteComment(w http.ResponseWriter, r *http.Request, uid string) {
	statusCode, _ := it.forumService.Comment().DeleteComment(r.Context(), uid)

	w.WriteHeader(statusCode)
}

func (it *ServerInterfaceImpl) GetCommentReplies(w http.ResponseWriter, r *http.Request, uid string) {

	res, statusCode, _ := it.forumService.Comment().GetReplies(r.Context(), uid)

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func (it *ServerInterfaceImpl) ReplyToComment(w http.ResponseWriter, r *http.Request, uid string) {
	var data codeHelpAdminGen.ReplyToCommentJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	res, statusCode, _ := it.forumService.Comment().ReplyToComment(r.Context(), uid, codeHelpForumGen.ReplyToCommentJSONRequestBody{
		Content: data.Content,
	})

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func (it *ServerInterfaceImpl) UpdateComment(w http.ResponseWriter, r *http.Request, uid string) {
	var data codeHelpAdminGen.UpdateCommentJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	res, statusCode, _ := it.forumService.Comment().EditComment(r.Context(), uid, codeHelpForumGen.UpdateCommentJSONRequestBody{
		Content: data.Content,
	})

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func (it *ServerInterfaceImpl) GetAllCommunities(w http.ResponseWriter, r *http.Request) {
	res, statusCode, _ := it.forumService.Community().GetCommunities(r.Context())

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func (it *ServerInterfaceImpl) CreateCommunity(w http.ResponseWriter, r *http.Request) {
	var data codeHelpAdminGen.CreateCommunityJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	res, statusCode, _ := it.forumService.Community().CreateCommunity(r.Context(), codeHelpForumGen.CreateCommunityJSONRequestBody{
		Categories:  data.Categories,
		Description: data.Description,
		Image:       data.Image,
		Name:        data.Name,
	})

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func (it *ServerInterfaceImpl) LeaveCommunity(w http.ResponseWriter, r *http.Request, params codeHelpAdminGen.LeaveCommunityParams) {
	statusCode := it.forumService.Community().LeaveCommunity(r.Context(), codeHelpForumGen.LeaveCommunityParams{
		Community: params.Community,
	})

	w.WriteHeader(statusCode)
}

func (it *ServerInterfaceImpl) JoinCommunity(w http.ResponseWriter, r *http.Request, params codeHelpAdminGen.JoinCommunityParams) {
	statusCode := it.forumService.Community().JoinCommunity(r.Context(), codeHelpForumGen.JoinCommunityParams{
		Community: params.Community,
	})

	w.WriteHeader(statusCode)
}

func (it *ServerInterfaceImpl) RemoveModerator(w http.ResponseWriter, r *http.Request, params codeHelpAdminGen.RemoveModeratorParams) {
	statusCode := it.forumService.Community().RemoveModerator(r.Context(), codeHelpForumGen.RemoveModeratorParams{
		Community: params.Community,
		Username:  params.Username,
	})

	w.WriteHeader(statusCode)
}

func (it *ServerInterfaceImpl) GetCommunityModerators(w http.ResponseWriter, r *http.Request, params codeHelpAdminGen.GetCommunityModeratorsParams) {

	moderators, statusCode, _ := it.forumService.Community().GetModerators(
		r.Context(),
		codeHelpForumGen.GetCommunityModeratorsParams{Name: params.Name})

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(moderators)
}

func (it *ServerInterfaceImpl) AddModerator(w http.ResponseWriter, r *http.Request, params codeHelpAdminGen.AddModeratorParams) {
	var data codeHelpAdminGen.AddModeratorJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	statusCode := it.forumService.Community().AddModerator(r.Context(),
		codeHelpForumGen.AddModeratorParams{
			Community: params.Community,
		}, codeHelpForumGen.AddModeratorJSONRequestBody{
			Username: data.Username,
		})

	w.WriteHeader(statusCode)
}

func (it *ServerInterfaceImpl) DeleteCommunity(w http.ResponseWriter, r *http.Request, name string) {
	statusCode := it.forumService.Community().DeleteCommunity(r.Context(), name)

	w.WriteHeader(statusCode)
}

func (it *ServerInterfaceImpl) GetCommunityByUid(w http.ResponseWriter, r *http.Request, name string) {
	res, statusCode, _ := it.forumService.Community().GetCommunity(r.Context(), name)

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func (it *ServerInterfaceImpl) UpdateCommunity(w http.ResponseWriter, r *http.Request, name string) {
	var data codeHelpAdminGen.UpdateCommunityJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	res, statusCode, _ := it.forumService.Community().UpdateCommunity(r.Context(), name, codeHelpForumGen.UpdateCommunityJSONRequestBody{
		Categories:  data.Categories,
		Description: data.Description,
		Image:       data.Image,
		Name:        data.Name,
	})

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func (it *ServerInterfaceImpl) GetPosts(w http.ResponseWriter, r *http.Request, params codeHelpAdminGen.GetPostsParams) {
	res, statusCode, _ := it.forumService.Post().GetPosts(r.Context(), codeHelpForumGen.GetPostsParams{Community: params.Community})

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func (it *ServerInterfaceImpl) CreateCommunityPost(w http.ResponseWriter, r *http.Request, params codeHelpAdminGen.CreateCommunityPostParams) {
	var data codeHelpAdminGen.CreateCommunityPostJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	res, statusCode, _ := it.forumService.Post().CreatePost(r.Context(),
		codeHelpForumGen.CreateCommunityPostParams{Community: params.Community},
		codeHelpForumGen.CreateCommunityPostJSONRequestBody{
			Content: data.Content,
			Title:   data.Title,
		})

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func (it *ServerInterfaceImpl) DeletePost(w http.ResponseWriter, r *http.Request, uid string) {
	statusCode, _ := it.forumService.Post().DeletePost(r.Context(), uid)

	w.WriteHeader(statusCode)
}

func (it *ServerInterfaceImpl) GetPost(w http.ResponseWriter, r *http.Request, uid string) {
	res, statusCode, _ := it.forumService.Post().GetPost(r.Context(), uid)

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func (it *ServerInterfaceImpl) UpdatePost(w http.ResponseWriter, r *http.Request, uid string) {
	var data codeHelpAdminGen.UpdatePostJSONRequestBody
	_ = json.NewDecoder(r.Body).Decode(&data)

	res, statusCode, _ := it.forumService.Post().UpdatePost(r.Context(), uid, codeHelpForumGen.UpdatePostJSONRequestBody{
		Content: data.Content,
		Title:   data.Title,
	})

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}

func (it *ServerInterfaceImpl) GetAllForumCategories(w http.ResponseWriter, r *http.Request) {
	res, statusCode, _ := it.forumService.Category().GetCategories(r.Context())

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(res)
}
