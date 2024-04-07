package core

import (
	codeHelpAdminCoreGen "api-spec/generated/code-help-admin-core"
)

func CreateClient(server string) (codeHelpAdminCoreGen.ClientInterface, error) {
	return codeHelpAdminCoreGen.NewClient(server)
}
