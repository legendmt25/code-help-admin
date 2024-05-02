package category

import (
	"admin-api/internal/util/decodeUtil"
	codeHelpForumGen "api-spec/generated/code-help-forum"
	"github.com/labstack/gommon/log"
	"net/http"
)

type Decoder interface {
	Decode(response *http.Response) *codeHelpForumGen.Category

	DecodeAll(response *http.Response) *codeHelpForumGen.Categories
}

type defaultDecoder struct {
}

func NewDefaultDecoder() Decoder {
	return &defaultDecoder{}
}

func (d defaultDecoder) Decode(response *http.Response) *codeHelpForumGen.Category {
	var category codeHelpForumGen.Category

	if err := decodeUtil.Decode(response, &category); err != nil {
		log.Error(err)
		return nil
	}

	return &category
}

func (d defaultDecoder) DecodeAll(response *http.Response) *codeHelpForumGen.Categories {
	var categories codeHelpForumGen.Categories

	if err := decodeUtil.Decode(response, &categories); err != nil {
		log.Error(err)
		return nil
	}

	return &categories
}
