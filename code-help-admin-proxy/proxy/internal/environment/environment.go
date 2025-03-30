package environment

import (
	"fmt"
	"os"
	"strings"
)

type Environment struct {
	Oauth2Jwt     Oauth2Jwt
	ProxyMappings []string
}

type Oauth2Jwt struct {
	IssuerUri          string
	DiscoveryIssuerUri string
	JwkSetUri          string
}

func Load() Environment {
	return Environment{
		Oauth2Jwt:     getOauth2JwtData(),
		ProxyMappings: getProxyMappings(),
	}
}

func getProxyMappings() []string {
	return strings.Split(getEnv("PROXY_MAPPINGS"), ",")
}

func getOauth2JwtData() Oauth2Jwt {
	return Oauth2Jwt{
		IssuerUri:          getEnv("OAUTH2_RESOURCESERVER_JWT_ISSUERURI"),
		JwkSetUri:          getEnv("OAUTH2_RESOURCESERVER_JWT_JWKSETURI"),
		DiscoveryIssuerUri: getEnv("OAUTH2_RESOURCESERVER_JWT__DISCOVERYISSUERURI"),
	}
}

func getEnv(key string) string {
	env := make(map[string]string)

	env["PROXY_MAPPINGS"] = ".*/api/v1/admin/forum->http://localhost:3001/api/v1,.*/api/v1/admin->http://localhost:3000/api/admin"
	env["OAUTH2_RESOURCESERVER_JWT_ISSUERURI"] = "http://localhost/iam/realms/code-help"
	env["OAUTH2_RESOURCESERVER_JWT__DISCOVERYISSUERURI"] = "http://localhost/iam/realms/code-help"
	env["OAUTH2_RESOURCESERVER_JWT_JWKSETURI"] = "http://localhost/iam/realms/code-help/protocol/openid-connect/certs"

	if isNotBlank(os.Getenv(key)) {
		fmt.Printf("Getting environment variable: %s - %s\n", key, os.Getenv(key))
		return os.Getenv(key)
	}

	fmt.Printf("Getting environment variable: %s - %s\n", key, env[key])
	return env[key]
}

func isNotBlank(value string) bool {
	return len(strings.TrimSpace(value)) > 0
}
