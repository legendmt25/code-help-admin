package comment

import (
	"admin-api/internal/util/decodeUtil"
	codeHelpForumGen "api-spec/generated/code-help-forum"
	"github.com/labstack/gommon/log"
	"net/http"
)

type Decoder interface {
	Decode(response *http.Response) *codeHelpForumGen.Comment

	DecodeAll(response *http.Response) *codeHelpForumGen.Comments
}

type defaultDecoder struct {
}

func NewDefaultDecoder() Decoder {
	return &defaultDecoder{}
}

func (d defaultDecoder) Decode(response *http.Response) *codeHelpForumGen.Comment {
	var comment codeHelpForumGen.Comment

	if err := decodeUtil.Decode(response, &comment); err != nil {
		log.Error(err)
		return nil
	}

	return &comment
}

func (d defaultDecoder) DecodeAll(response *http.Response) *codeHelpForumGen.Comments {
	var comments codeHelpForumGen.Comments

	if err := decodeUtil.Decode(response, &comments); err != nil {
		log.Error(err)
		return nil
	}

	return &comments
}
