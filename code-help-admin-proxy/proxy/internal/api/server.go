package api

import (
	"log"
	"net/http"
)

type ProxyServer struct {
	server *http.Server
}

func (it ProxyServer) Serve() {
	log.Println("Starting admin server...")
	log.Fatal(it.server.ListenAndServe())
}
