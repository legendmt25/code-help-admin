package comment

import (
	codeHelpForumGen "api-spec/generated/code-help-forum"
	"context"
	"github.com/labstack/gommon/log"
	"net/http"
)

type ForumCommentService interface {
	GetComments(ctx context.Context, params codeHelpForumGen.GetCommentsForPostParams) (*codeHelpForumGen.Comments, int, error)

	GetReplies(ctx context.Context, uid string) (*codeHelpForumGen.Comments, int, error)

	CreateComment(ctx context.Context, params codeHelpForumGen.CommentOnPostParams, body codeHelpForumGen.CommentOnPostJSONRequestBody) (*codeHelpForumGen.Comment, int, error)

	EditComment(ctx context.Context, uid string, body codeHelpForumGen.UpdateCommentJSONRequestBody) (*codeHelpForumGen.Comment, int, error)

	ReplyToComment(ctx context.Context, uid string, body codeHelpForumGen.ReplyToCommentJSONRequestBody) (*codeHelpForumGen.Comment, int, error)

	DeleteComment(ctx context.Context, uid string) (int, error)
}

type forumCommentServiceImpl struct {
	client  codeHelpForumGen.ClientInterface
	decoder Decoder
}

func NewCommentServiceImpl(client codeHelpForumGen.ClientInterface) ForumCommentService {
	return &forumCommentServiceImpl{
		client:  client,
		decoder: NewDefaultDecoder(),
	}
}

func (it *forumCommentServiceImpl) GetComments(ctx context.Context, params codeHelpForumGen.GetCommentsForPostParams) (*codeHelpForumGen.Comments, int, error) {
	res, err := it.client.GetCommentsForPost(ctx, &params)

	if err != nil {
		log.Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, nil
	}

	return it.decoder.DecodeAll(res), http.StatusOK, nil
}

func (it *forumCommentServiceImpl) GetReplies(ctx context.Context, uid string) (*codeHelpForumGen.Comments, int, error) {
	res, err := it.client.GetCommentReplies(ctx, uid)

	if err != nil {
		log.Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, nil
	}

	return it.decoder.DecodeAll(res), http.StatusOK, nil
}

func (it *forumCommentServiceImpl) CreateComment(ctx context.Context, params codeHelpForumGen.CommentOnPostParams, body codeHelpForumGen.CommentOnPostJSONRequestBody) (*codeHelpForumGen.Comment, int, error) {
	res, err := it.client.CommentOnPost(ctx, &params, body)

	if err != nil {
		log.Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, nil
	}

	return it.decoder.Decode(res), http.StatusOK, nil
}

func (it *forumCommentServiceImpl) ReplyToComment(ctx context.Context, uid string, body codeHelpForumGen.ReplyToCommentJSONRequestBody) (*codeHelpForumGen.Comment, int, error) {
	res, err := it.client.ReplyToComment(ctx, uid, body)

	if err != nil {
		log.Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, nil
	}

	return it.decoder.Decode(res), http.StatusOK, nil
}

func (it *forumCommentServiceImpl) EditComment(ctx context.Context, uid string, body codeHelpForumGen.UpdateCommentJSONRequestBody) (*codeHelpForumGen.Comment, int, error) {
	res, err := it.client.UpdateComment(ctx, uid, body)

	if err != nil {
		log.Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, nil
	}

	return it.decoder.Decode(res), http.StatusOK, nil
}

func (it *forumCommentServiceImpl) DeleteComment(ctx context.Context, uid string) (int, error) {
	res, err := it.client.DeleteComment(ctx, uid)

	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, err
	}

	return res.StatusCode, nil
}
