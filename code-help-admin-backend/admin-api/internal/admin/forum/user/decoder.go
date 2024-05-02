package user

import (
	"admin-api/internal/util/decodeUtil"
	codeHelpForumGen "api-spec/generated/code-help-forum"
	"github.com/labstack/gommon/log"
	"net/http"
)

type Decoder interface {
	Decode(response *http.Response) *codeHelpForumGen.User

	DecodeAll(response *http.Response) *codeHelpForumGen.Users
}

type defaultDecoder struct {
}

func NewDefaultDecoder() Decoder {
	return &defaultDecoder{}
}

func (d defaultDecoder) Decode(response *http.Response) *codeHelpForumGen.User {
	var user codeHelpForumGen.User

	if err := decodeUtil.Decode(response, &user); err != nil {
		log.Error(err)
		return nil
	}

	return &user
}

func (d defaultDecoder) DecodeAll(response *http.Response) *codeHelpForumGen.Users {

	var users codeHelpForumGen.Users

	if err := decodeUtil.Decode(response, &users); err != nil {
		log.Error(err)
		return nil
	}

	return &users
}
