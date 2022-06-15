package gcpen

import (
	"context"
	"os"

	"google.golang.org/api/cloudresourcemanager/v1"
)

var (
	// ProjectID - Google Cloud Platform Project ID
	ProjectID = getProjectID()
	// ServiceName - Google App Engine service name
	ServiceName = os.Getenv(EnvKeyServiceName)
	// ServiceVersion - Google App Engine service version
	ServiceVersion = os.Getenv(EnvKeyServiceVersion)
	projectNumber  int64
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
