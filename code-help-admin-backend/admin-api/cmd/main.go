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

	codeHelpAdminClient, _ := core.CreateClient(env.CoreServerUrl)
	coreService := admin.NewCoreService(codeHelpAdminClient)

	codeHelpForumClient, _ := forum.CreateClient(env.ForumServerUrl)
	forumService := forum.NewForumService(codeHelpForumClient)

	adminService := api.NewServiceInterfaceImpl(coreService, forumService)

	api.NewAdminApiServer(adminService).Serve()
}
