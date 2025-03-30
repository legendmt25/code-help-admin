package api

import (
	"fmt"
	"io"
	"net/http"
	auth "proxy/internal/auth"
	"proxy/internal/environment"
	"regexp"
	"strings"
	"time"
)

func parseProxyMappings(environment environment.Environment) (map[string]string, []string) {
	mappingsArray := environment.ProxyMappings
	keys := make([]string, 0)
	mappings := make(map[string]string)

	for _, mapping := range mappingsArray {
		split := strings.SplitN(mapping, "->", 2)
		from := split[0]
		to := split[1]
		keys = append(keys, from)

		mappings[from] = to
	}

	return mappings, keys
}

func NewProxyServer(environment environment.Environment, oidcService auth.OidcService) ProxyServer {

	regexes := make(map[string]*regexp.Regexp)
	relocators := make(map[string]string)

	mappings, keys := parseProxyMappings(environment)

	for _, key := range keys {
		from := key
		to := mappings[key]
		regexMapFrom, _ := regexp.Compile(from)
		regexes[from] = regexMapFrom
		relocators[from] = to

		fmt.Printf("Proxy mapping: %s -> %s\n", from, to)
	}

	reverseProxyHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("[reverse proxy server] received request at: %s\n", time.Now())

		targetUrl := req.URL.String()
		fmt.Printf("[reverse proxy server] url: %s %s\n", req.Method, targetUrl)
		for _, key := range keys {
			regex := regexes[key]

			if regex.MatchString(targetUrl) {
				targetUrl = regex.ReplaceAllString(targetUrl, relocators[key])
				fmt.Printf("[reverse proxy server] relocating to: %s\n", targetUrl)
			}
		}

		proxyReq, err := http.NewRequest(req.Method, targetUrl, req.Body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(rw, err)
		}

		proxyReq.Header = req.Header

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
