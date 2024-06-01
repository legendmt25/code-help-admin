package main

import (
	"admin-api/internal/admin"
	"admin-api/internal/admin/forum"
	"admin-api/internal/api"
	"admin-api/internal/client"
	"admin-api/internal/environment"
	"log"
	"sync"
	"time"
)

func main() {
	env := environment.Load()

	oidcService, err := api.NewOidcService(env)

	if err != nil {
		retry(func() {
			oidcService, err = api.NewOidcService(env)
		}, func() bool {
			return err == nil
		})
	}

	codeHelpAdminClient, _ := client.CreateAdminCoreClient(env.CoreServerUrl, oidcService)
	coreService := admin.NewCoreService(codeHelpAdminClient)

	codeHelpForumClient, _ := client.CreateForumClient(env.ForumServerUrl, oidcService)
	forumService := forum.NewForumService(codeHelpForumClient)

	adminService := api.NewServiceInterfaceImpl(coreService, forumService)

	api.NewAdminApiServer(adminService, oidcService).Serve()
}

func retry(task func(), condition func() bool) {
	var wg sync.WaitGroup
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Retrying...")
				if condition() {
					close(quit)
				} else {
					task()
				}
			case <-quit:
				ticker.Stop()
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()
}
