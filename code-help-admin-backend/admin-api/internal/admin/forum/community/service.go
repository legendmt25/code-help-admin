package community

import (
	"admin-api/internal/admin/forum/user"
	codeHelpForumGen "api-spec/generated/code-help-forum"
	"context"
	"github.com/labstack/gommon/log"
	"net/http"
)

type ForumCommunityService interface {
	GetCommunity(ctx context.Context, name string) (*codeHelpForumGen.Community, error)

	GetCommunities(ctx context.Context) []codeHelpForumGen.ShortCommunity

	CreateCommunity(ctx context.Context, body codeHelpForumGen.CreateCommunityJSONRequestBody) (*codeHelpForumGen.Community, error)

	UpdateCommunity(ctx context.Context, name string, body codeHelpForumGen.UpdateCommunityJSONRequestBody) (*codeHelpForumGen.Community, error)

	DeleteCommunity(ctx context.Context, name string) int

	GetModerators(ctx context.Context, params codeHelpForumGen.GetCommunityModeratorsParams) (*codeHelpForumGen.Users, error)

	AddModerator(ctx context.Context, params codeHelpForumGen.AddModeratorParams, body codeHelpForumGen.AddModeratorJSONRequestBody) int

	RemoveModerator(ctx context.Context, params codeHelpForumGen.RemoveModeratorParams) int

	JoinCommunity(ctx context.Context, params codeHelpForumGen.JoinCommunityParams) int

	LeaveCommunity(ctx context.Context, params codeHelpForumGen.LeaveCommunityParams) int
}

type forumCommunityServiceImpl struct {
	client      codeHelpForumGen.ClientInterface
	decoder     Decoder
	userDecoder user.Decoder
}

func NewCommunityServiceImpl(client codeHelpForumGen.ClientInterface) ForumCommunityService {
	return &forumCommunityServiceImpl{
		client:      client,
		decoder:     NewDefaultDecoder(),
		userDecoder: user.NewDefaultDecoder(),
	}
}

func (it *forumCommunityServiceImpl) GetCommunities(ctx context.Context) []codeHelpForumGen.ShortCommunity {
	res, err := it.client.GetAllCommunities(ctx)

	if err != nil {
		return make([]codeHelpForumGen.ShortCommunity, 0)
	}

	return it.decoder.DecodeAllShort(res)
}

func (it *forumCommunityServiceImpl) CreateCommunity(ctx context.Context, body codeHelpForumGen.CreateCommunityJSONRequestBody) (*codeHelpForumGen.Community, error) {
	res, err := it.client.CreateCommunity(ctx, body)
	if err != nil {
		return nil, err
	}

	return it.decoder.Decode(res), nil
}

func (it *forumCommunityServiceImpl) GetCommunity(ctx context.Context, name string) (*codeHelpForumGen.Community, error) {
	res, err := it.client.GetCommunityByUid(ctx, name)
	if err != nil {
		return nil, err
	}

	return it.decoder.Decode(res), nil
}

func (it *forumCommunityServiceImpl) UpdateCommunity(ctx context.Context, name string, body codeHelpForumGen.UpdateCommunityJSONRequestBody) (*codeHelpForumGen.Community, error) {
	res, err := it.client.UpdateCommunity(ctx, name, body)
	if err != nil {
		return nil, err
	}

	return it.decoder.Decode(res), nil
}

func (it *forumCommunityServiceImpl) DeleteCommunity(ctx context.Context, name string) int {
	res, err := it.client.DeleteCommunity(ctx, name)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError
	}

	return res.StatusCode
}

func (it *forumCommunityServiceImpl) GetModerators(ctx context.Context, params codeHelpForumGen.GetCommunityModeratorsParams) (*codeHelpForumGen.Users, error) {
	res, err := it.client.GetCommunityModerators(ctx, &params)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return it.userDecoder.DecodeAll(res), err
}

func (it *forumCommunityServiceImpl) AddModerator(ctx context.Context, params codeHelpForumGen.AddModeratorParams, body codeHelpForumGen.AddModeratorJSONRequestBody) int {
	res, err := it.client.AddModerator(ctx, &params, body)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError
	}

	return res.StatusCode
}

func (it *forumCommunityServiceImpl) RemoveModerator(ctx context.Context, params codeHelpForumGen.RemoveModeratorParams) int {
	res, err := it.client.RemoveModerator(ctx, &params)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError
	}

	return res.StatusCode
}

func (it *forumCommunityServiceImpl) JoinCommunity(ctx context.Context, params codeHelpForumGen.JoinCommunityParams) int {
	res, err := it.client.JoinCommunity(ctx, &params)
	if err != nil {
		return http.StatusInternalServerError
	}

	return res.StatusCode
}

func (it *forumCommunityServiceImpl) LeaveCommunity(ctx context.Context, params codeHelpForumGen.LeaveCommunityParams) int {
	res, err := it.client.LeaveCommunity(ctx, &params)
	if err != nil {
		return http.StatusInternalServerError
	}

	return res.StatusCode
}
