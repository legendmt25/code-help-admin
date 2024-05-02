package forum

import codeHelpForum "api-spec/generated/code-help-forum"

func CreateClient(server string) (codeHelpForum.ClientInterface, error) {
	return codeHelpForum.NewClient(server)
}
