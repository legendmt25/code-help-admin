package environment

import (
	"os"
	"strings"
)

type Environment struct {
	CoreServerUrl  string
	ForumServerUrl string
	Oauth2Jwt      Oauth2Jwt
}

type Oauth2Jwt struct {
	IssuerUri          string
	DiscoveryIssuerUri string
	JwkSetUri          string
}

func Load() Environment {
	return Environment{
		CoreServerUrl:  getEnv("CORE_ADMIN_API_SERVER_URL"),
		ForumServerUrl: getEnv("FORUM_API_SERVER_URL"),
		Oauth2Jwt:      getOauth2JwtData(),
	}
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

	env["CORE_ADMIN_API_SERVER_URL"] = "http://localhost:3000"
	env["FORUM_API_SERVER_URL"] = "http://localhost:3001"
	env["OAUTH2_RESOURCESERVER_JWT_ISSUERURI"] = "http://localhost/iam/realms/code-help"
	env["OAUTH2_RESOURCESERVER_JWT__DISCOVERYISSUERURI"] = "http://localhost/iam/realms/code-help"
	env["OAUTH2_RESOURCESERVER_JWT_JWKSETURI"] = "http://localhost/iam/realms/code-help/protocol/openid-connect/certs"

	if isNotBlank(os.Getenv(key)) {
		return os.Getenv(key)
	}

	return env[key]
}

func isNotBlank(value string) bool {
	return len(strings.TrimSpace(value)) > 0
}
