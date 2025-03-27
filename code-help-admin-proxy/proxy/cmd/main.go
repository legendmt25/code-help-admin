package main

import (
	"log"
	"proxy/internal/api"
	"proxy/internal/auth"
	"proxy/internal/environment"
	"sync"
	"time"
)

func main() {
	env := environment.Load()

	oidcService, err := auth.NewOidcService(env)

	if err != nil {
		retry(func() {
			oidcService, err = auth.NewOidcService(env)
		}, func() bool {
			return err == nil
		})
	}

	api.NewProxyServer(oidcService).Serve()
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
