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
