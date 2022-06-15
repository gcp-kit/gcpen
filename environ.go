package gcpen

import (
	"context"
	"os"

	"google.golang.org/api/cloudresourcemanager/v1"
)

var (
	// ProjectID - Google Cloud Platform Project ID
	ProjectID     = getProjectID()
	projectNumber int64

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

// Reload - reload environment variables.
func Reload() {
	ProjectID = getProjectID()

	// ServiceName - will be deleted in a major version upgrade
	// Deprecated
	ServiceName = os.Getenv(EnvKeyServiceName)
	// ServiceVersion - will be deleted in a major version upgrade
	// Deprecated
	ServiceVersion = os.Getenv(EnvKeyServiceVersion)

	GAEServiceName = os.Getenv(EnvKeyGAEServiceName)
	GAEServiceVersion = os.Getenv(EnvKeyGAEServiceVersion)
	RUNServiceName = os.Getenv(EnvKeyRUNServiceName)
	RUNServiceRevision = os.Getenv(EnvKeyRUNServiceRevision)
}

// GetProjectNumber - get project number
func GetProjectNumber(ctx context.Context, forceRefresh ...bool) (int64, error) {
	if projectNumber != 0 && !(len(forceRefresh) > 0 && forceRefresh[0]) {
		return projectNumber, nil
	}

	service, err := cloudresourcemanager.NewService(ctx)
	if err != nil {
		return 0, err
	}

	project, err := service.Projects.Get(getProjectID()).Context(ctx).Do()
	if err != nil {
		return 0, err
	}

	projectNumber = project.ProjectNumber

	return projectNumber, nil
}
