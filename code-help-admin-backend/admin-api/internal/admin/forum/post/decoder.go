package post

import (
	"admin-api/internal/util/decodeUtil"
	codeHelpForumGen "api-spec/generated/code-help-forum"
	"github.com/labstack/gommon/log"
	"net/http"
)

type Decoder interface {
	Decode(response *http.Response) *codeHelpForumGen.Post

	DecodeShort(response *http.Response) *codeHelpForumGen.ShortPost

	DecodeAll(response *http.Response) []codeHelpForumGen.Post

	DecodeAllShort(response *http.Response) *codeHelpForumGen.ShortPosts
}

type defaultDecoder struct {
}

func NewDefaultDecoder() Decoder {
	return &defaultDecoder{}
}

func (d defaultDecoder) Decode(response *http.Response) *codeHelpForumGen.Post {
	var post codeHelpForumGen.Post

	if err := decodeUtil.Decode(response, &post); err != nil {
		log.Error(err)
		return nil
	}

	return &post
}

func (d defaultDecoder) DecodeShort(response *http.Response) *codeHelpForumGen.ShortPost {
	var post codeHelpForumGen.ShortPost

	if err := decodeUtil.Decode(response, &post); err != nil {
		log.Error(err)
		return nil
	}

	return &post
}

func (d defaultDecoder) DecodeAll(response *http.Response) []codeHelpForumGen.Post {
	var posts []codeHelpForumGen.Post

	if err := decodeUtil.Decode(response, &posts); err != nil {
		log.Error(err)
		return make([]codeHelpForumGen.Post, 0)
	}

	return posts
}

func (d defaultDecoder) DecodeAllShort(response *http.Response) *codeHelpForumGen.ShortPosts {
	var posts codeHelpForumGen.ShortPosts

	if err := decodeUtil.Decode(response, &posts); err != nil {
		log.Error(err)
		return nil
	}

	return &posts
}
