package category

import (
	codeHelpForumGen "api-spec/generated/code-help-forum"
	"context"
	"github.com/labstack/gommon/log"
	"net/http"
)

type ForumCategoryService interface {
	GetCategories(ctx context.Context) (*codeHelpForumGen.Categories, int, error)

	CreateCategory(ctx context.Context, body codeHelpForumGen.CreateCategoryJSONRequestBody) (int, error)

	UpdateCategory(ctx context.Context, uid string, body codeHelpForumGen.UpdateCategoryJSONRequestBody) (*codeHelpForumGen.Category, int, error)

	DeleteCategory(ctx context.Context, uid string) (int, error)
}

type forumCategoryServiceImpl struct {
	client  codeHelpForumGen.ClientInterface
	decoder Decoder
}

func NewCategoryServiceImpl(client codeHelpForumGen.ClientInterface) ForumCategoryService {
	return &forumCategoryServiceImpl{
		client:  client,
		decoder: NewDefaultDecoder(),
	}
}

func (it *forumCategoryServiceImpl) GetCategories(ctx context.Context) (*codeHelpForumGen.Categories, int, error) {
	res, err := it.client.GetAllCategories(ctx)

	if err != nil {
		log.Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, err
	}

	return it.decoder.DecodeAll(res), http.StatusOK, nil
}

func (it *forumCategoryServiceImpl) CreateCategory(ctx context.Context, body codeHelpForumGen.CreateCategoryJSONRequestBody) (int, error) {
	res, err := it.client.CreateCategory(ctx, body)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, err
	}

	return res.StatusCode, nil
}

func (it *forumCategoryServiceImpl) UpdateCategory(ctx context.Context, uid string, body codeHelpForumGen.UpdateCategoryJSONRequestBody) (*codeHelpForumGen.Category, int, error) {
	res, err := it.client.UpdateCategory(ctx, uid, body)
	if err != nil {
		log.Error(err)
		return nil, http.StatusInternalServerError, err
	}

	return it.decoder.Decode(res), res.StatusCode, nil
}

func (it *forumCategoryServiceImpl) DeleteCategory(ctx context.Context, uid string) (int, error) {
	res, err := it.client.DeleteCategory(ctx, uid)
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, err
	}

	return res.StatusCode, nil
}
