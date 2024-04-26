package main

import (
	"admin-api/internal/admin"
	"admin-api/internal/api"
	"admin-api/internal/core"
	"os"
)

func main() {
	coreAdminApiServerUrl := os.Getenv("CORE_ADMIN_API_SERVER_URL")

	codeHelpAdminClient, _ := core.CreateClient(coreAdminApiServerUrl)
	coreService := admin.NewCoreService(codeHelpAdminClient, admin.NewProblemDecoder())
	adminService := api.NewServiceInterfaceImpl(coreService)

	api.NewAdminApiServer(adminService).Serve()
}
