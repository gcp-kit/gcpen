package gcpen

import (
	"os"
)

var (
	// ProjectID - Google Cloud Platform Project ID
	ProjectID = getProjectID()
	// ServiceName - Google App Engine service name
	ServiceName = os.Getenv(EnvKeyServiceName)
	// ServiceVersion - Google App Engine service version
	ServiceVersion = os.Getenv(EnvKeyServiceVersion)
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
	ServiceName = os.Getenv(EnvKeyServiceName)
	ServiceVersion = os.Getenv(EnvKeyServiceVersion)
}
