package main

import (
	"admin-api/internal/admin"
	"admin-api/internal/api"
	"admin-api/internal/core"
	"admin-api/internal/environment"
)

func main() {
	env := environment.Load()

	codeHelpAdminClient, _ := core.CreateClient(env.CoreServerUrl)
	coreService := admin.NewCoreService(codeHelpAdminClient)
	adminService := api.NewServiceInterfaceImpl(coreService)

	api.NewAdminApiServer(adminService, env).Serve()
}
