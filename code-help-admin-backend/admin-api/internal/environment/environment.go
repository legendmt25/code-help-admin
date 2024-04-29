package environment

import (
	"os"
	"strings"
)

type Environment struct {
	CoreServerUrl string
}

func Load() Environment {
	return Environment{
		CoreServerUrl: getAdminApiServerUtl(),
	}
}

func getAdminApiServerUtl() string {
	envValue := os.Getenv("CORE_ADMIN_API_SERVER_URL")

	if len(strings.TrimSpace(envValue)) > 0 {
		return envValue
	}

	return "http://localhost:3000/api"
}
