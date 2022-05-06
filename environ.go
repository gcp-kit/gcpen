package gcpen

import (
	"os"
)

var (
	// ProjectID - Google Cloud Platform Project ID
	ProjectID = getProjectID()
	// GAEServiceName - Google App Engine service name
	GAEServiceName = os.Getenv(EnvKeyGAEServiceName)
	// GAEServiceVersion - Google App Engine service version
	GAEServiceVersion = os.Getenv(EnvKeyGAEServiceVersion)
)

func getProjectID() string {
	id, ok := os.LookupEnv(EnvKeyProjectID)
	if ok {
		return id
	}

	id, ok = os.LookupEnv(EnvKeyGoogleCloudProject)
	if ok {
		return id
	}

	return ""
}

// Reload - Reload environment variables.
func Reload() {
	ProjectID = getProjectID()
	GAEServiceName = os.Getenv(EnvKeyGAEServiceName)
	GAEServiceVersion = os.Getenv(EnvKeyGAEServiceVersion)
}
