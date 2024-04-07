package main

import (
	"admin-api/internal/admin"
	"admin-api/internal/api"
	"admin-api/internal/core"
)

func main() {
	codeHelpAdminClient, _ := core.CreateClient("http://localhost:3000/api")
	coreService := admin.NewCoreService(codeHelpAdminClient, admin.NewProblemDecoder())
	adminService := api.NewServiceInterfaceImpl(coreService)

	api.NewAdminApiServer(adminService).Serve()
}
