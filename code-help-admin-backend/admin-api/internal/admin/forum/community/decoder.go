package community

import (
	"admin-api/internal/util/decodeUtil"
	codeHelpForumGen "api-spec/generated/code-help-forum"
	"github.com/labstack/gommon/log"
	"net/http"
)

type Decoder interface {
	Decode(response *http.Response) *codeHelpForumGen.Community

	DecodeShort(response *http.Response) *codeHelpForumGen.ShortCommunity

	DecodeAll(response *http.Response) []codeHelpForumGen.Community

	DecodeAllShort(response *http.Response) []codeHelpForumGen.ShortCommunity
}

type defaultDecoder struct {
}

func NewDefaultDecoder() Decoder {
	return &defaultDecoder{}
}

func (d *defaultDecoder) Decode(response *http.Response) *codeHelpForumGen.Community {
	var community codeHelpForumGen.Community

	if err := decodeUtil.Decode(response, &community); err != nil {
		log.Error(err)
		return nil
	}

	return &community
}

func (d *defaultDecoder) DecodeShort(response *http.Response) *codeHelpForumGen.ShortCommunity {
	var community codeHelpForumGen.ShortCommunity

	if err := decodeUtil.Decode(response, &community); err != nil {
		log.Error(err)
		return nil
	}

	return &community
}

func (d *defaultDecoder) DecodeAll(response *http.Response) []codeHelpForumGen.Community {
	var communities []codeHelpForumGen.Community

	if err := decodeUtil.Decode(response, &communities); err != nil {
		log.Error(err)
		return make([]codeHelpForumGen.Community, 0)
	}

	return communities
}

func (d *defaultDecoder) DecodeAllShort(response *http.Response) []codeHelpForumGen.ShortCommunity {
	var communities []codeHelpForumGen.ShortCommunity

	if err := decodeUtil.Decode(response, &communities); err != nil {
		log.Error(err)
		return make([]codeHelpForumGen.ShortCommunity, 0)
	}

	return communities
}
