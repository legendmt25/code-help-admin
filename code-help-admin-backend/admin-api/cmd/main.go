package main

import (
	"admin-api/internal/admin"
	"admin-api/internal/admin/forum"
	"admin-api/internal/api"
	"admin-api/internal/client"
	"admin-api/internal/environment"
)

func main() {
	env := environment.Load()

	oidcService := api.NewOidcService(env)

	codeHelpAdminClient, _ := client.CreateAdminCoreClient(env.CoreServerUrl, oidcService)
	coreService := admin.NewCoreService(codeHelpAdminClient)

	codeHelpForumClient, _ := client.CreateForumClient(env.ForumServerUrl, oidcService)
	forumService := forum.NewForumService(codeHelpForumClient)

	adminService := api.NewServiceInterfaceImpl(coreService, forumService)

	api.NewAdminApiServer(adminService, oidcService).Serve()
}
