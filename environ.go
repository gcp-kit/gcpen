package gcpen

import (
	"os"
)

var (
	// ProjectID - Google Cloud Platform Project ID
	ProjectID = getProjectID()

	// ServiceName - Google App Engine service name
	// Deprecated
	ServiceName = os.Getenv(EnvKeyServiceName)

	// ServiceVersion - Google App Engine service version
	// Deprecated
	ServiceVersion = os.Getenv(EnvKeyServiceVersion)

	// GAEServiceName - Google App Engine service name
	GAEServiceName = os.Getenv(EnvKeyGAEServiceName)
	// GAEServiceVersion - Google App Engine service version
	GAEServiceVersion = os.Getenv(EnvKeyGAEServiceVersion)
	// RUNServiceName - Google Cloud Run service name
	RUNServiceName = os.Getenv(EnvKeyRUNServiceName)
	// RUNServiceRevision - Google Cloud Run service revision
	RUNServiceRevision = os.Getenv(EnvKeyRUNServiceRevision)
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

	// Deprecated
	// ServiceName - Deprecated
	ServiceName = os.Getenv(EnvKeyServiceName)
	// Deprecated
	// ServiceVersion - Deprecated
	ServiceVersion = os.Getenv(EnvKeyServiceVersion)

	GAEServiceName = os.Getenv(EnvKeyGAEServiceName)
	GAEServiceVersion = os.Getenv(EnvKeyGAEServiceVersion)
	RUNServiceName = os.Getenv(EnvKeyRUNServiceName)
	RUNServiceRevision = os.Getenv(EnvKeyRUNServiceRevision)
}
