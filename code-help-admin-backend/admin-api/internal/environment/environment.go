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
	IssuerUri string
	JwkSetUri string
}

func Load() Environment {
	return Environment{
		CoreServerUrl:  getAdminApiServerUrl(),
		ForumServerUrl: getForumApiServerUrl(),
		Oauth2Jwt:      getOauth2JwtData(),
	}
}

func getAdminApiServerUrl() string {
	envValue := os.Getenv("CORE_ADMIN_API_SERVER_URL")

	if isNotBlank(envValue) {
		return envValue
	}

	return "http://localhost:3000/api"
}

func getForumApiServerUrl() string {
	envValue := os.Getenv("FORUM_API_SERVER_URL")

	if isNotBlank(envValue) {
		return envValue
	}

	return "http://localhost:3001/api/v1"
}

func getOauth2JwtData() Oauth2Jwt {
	return Oauth2Jwt{
		IssuerUri: getIssuerUri(),
		JwkSetUri: getJwkSetUri(),
	}
}

func getIssuerUri() string {
	issuerUri := os.Getenv("OAUTH2_RESOURCESERVER_JWT_ISSUERURI")

	if isNotBlank(issuerUri) {
		return issuerUri
	}

	return "http://localhost/iam/realms/code-help"
}

func getJwkSetUri() string {
	jwkSetUri := os.Getenv("OAUTH2_RESOURCESERVER_JWT_JWKSETURI")

	if isNotBlank(jwkSetUri) {
		return jwkSetUri
	}

	return "http://localhost/iam/realms/code-help/protocol/openid"
}

func isNotBlank(value string) bool {
	return len(strings.TrimSpace(value)) > 0
}
