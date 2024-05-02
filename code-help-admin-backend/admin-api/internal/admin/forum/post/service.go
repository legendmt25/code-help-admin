package post

import (
	codeHelpForumGen "api-spec/generated/code-help-forum"
	"context"
	"github.com/labstack/gommon/log"
	"net/http"
)

type ForumPostService interface {
	GetPost(ctx context.Context, uid string) (*codeHelpForumGen.Post, int, error)
	GetPosts(ctx context.Context, params codeHelpForumGen.GetPostsParams) (*codeHelpForumGen.ShortPosts, int, error)

	CreatePost(ctx context.Context, params codeHelpForumGen.CreateCommunityPostParams, body codeHelpForumGen.CreateCommunityPostJSONRequestBody) (*codeHelpForumGen.Post, int, error)

	UpdatePost(ctx context.Context, uid string, body codeHelpForumGen.UpdatePostJSONRequestBody) (*codeHelpForumGen.Post, int, error)

	DeletePost(ctx context.Context, uid string) (int, error)
}

type forumPostServiceImpl struct {
	client  codeHelpForumGen.ClientInterface
	decoder Decoder
}

func NewPostServiceImpl(client codeHelpForumGen.ClientInterface) ForumPostService {
	return &forumPostServiceImpl{
		client:  client,
		decoder: NewDefaultDecoder(),
	}
}

func (it *forumPostServiceImpl) GetPosts(ctx context.Context, params codeHelpForumGen.GetPostsParams) (*codeHelpForumGen.ShortPosts, int, error) {
	res, err := it.client.GetPosts(ctx, &params)
	if err != nil {
		log.Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, nil
	}

	return it.decoder.DecodeAllShort(res), http.StatusOK, nil
}

func (it *forumPostServiceImpl) CreatePost(ctx context.Context, params codeHelpForumGen.CreateCommunityPostParams, body codeHelpForumGen.CreateCommunityPostJSONRequestBody) (*codeHelpForumGen.Post, int, error) {
	res, err := it.client.CreateCommunityPost(ctx, &params, body)

	if err != nil {
		log.Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, nil
	}

	return it.decoder.Decode(res), http.StatusOK, nil
}

func (it *forumPostServiceImpl) UpdatePost(ctx context.Context, uid string, body codeHelpForumGen.UpdatePostJSONRequestBody) (*codeHelpForumGen.Post, int, error) {
	res, err := it.client.UpdatePost(ctx, uid, body)

	if err != nil {
		log.Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, nil
	}

	return it.decoder.Decode(res), http.StatusOK, nil
}

func (it *forumPostServiceImpl) DeletePost(ctx context.Context, uid string) (int, error) {
	res, err := it.client.DeletePost(ctx, uid)

	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, err
	}

	return res.StatusCode, err
}

func (it *forumPostServiceImpl) GetPost(ctx context.Context, uid string) (*codeHelpForumGen.Post, int, error) {
	res, err := it.client.GetPost(ctx, uid)

	if err != nil {
		log.Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, nil
	}

	return it.decoder.Decode(res), http.StatusOK, nil
}
