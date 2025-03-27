package api

import (
	"fmt"
	"io"
	"net/http"
	auth "proxy/internal/auth"
	"proxy/internal/environment"
	"regexp"
	"time"
)

func NewProxyServer(oidcService auth.OidcService) ProxyServer {

	regexString := [2]string{".*/admin/forum", ".*/admin"}
	forumAppRegex, _ := regexp.Compile(regexString[0])
	coreAppRegex, _ := regexp.Compile(regexString[1])

	regexMap := make(map[string]*regexp.Regexp)
	regexMap[regexString[0]] = forumAppRegex
	regexMap[regexString[1]] = coreAppRegex

	relocateMap := make(map[string]string)
	relocateMap[regexString[0]] = environment.Load().ForumServerUrl
	relocateMap[regexString[1]] = environment.Load().CoreServerUrl

	reverseProxyHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("[reverse proxy server] received request at: %s\n", time.Now())

		targetUrl := req.URL.String()
		fmt.Printf("[reverse proxy server] url: %s\n", targetUrl)
		for key, regex := range regexMap {
			if regex.MatchString(targetUrl) {
				targetUrl = regex.ReplaceAllString(targetUrl, relocateMap[key])
				fmt.Printf("[reverse proxy server] relocating to: %s\n", targetUrl)
			}
		}

		proxyReq, err := http.NewRequest(req.Method, targetUrl, req.Body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(rw, err)
		}

		// save the response from the origin server
		originServerResponse, err := http.DefaultClient.Do(proxyReq)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(rw, err)
			return
		}

		// return response to the client
		rw.WriteHeader(originServerResponse.StatusCode)
		io.Copy(rw, originServerResponse.Body)
	})

	middlewares := [1]func(http.Handler) http.Handler{auth.NewAuthMiddlewareWith(oidcService)}

	var httpHandler http.Handler = reverseProxyHandler
	for _, middleware := range middlewares {
		httpHandler = middleware(reverseProxyHandler)
	}

	server := http.Server{
		Addr:    ":4000",
		Handler: httpHandler,
	}

	return ProxyServer{&server}
}
