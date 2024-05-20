package main

import (
	"admin-api/internal/admin"
	"admin-api/internal/admin/forum"
	"admin-api/internal/api"
	"admin-api/internal/core"
	"admin-api/internal/environment"
)

func main() {
	env := environment.Load()

	oidcService := api.NewOidcService(env)

	codeHelpAdminClient, _ := core.CreateClient(env.CoreServerUrl, oidcService)
	coreService := admin.NewCoreService(codeHelpAdminClient)

	codeHelpForumClient, _ := forum.CreateClient(env.ForumServerUrl, oidcService)
	forumService := forum.NewForumService(codeHelpForumClient)

	adminService := api.NewServiceInterfaceImpl(coreService, forumService)

	api.NewAdminApiServer(adminService, oidcService).Serve()
}
