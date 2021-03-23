package gcpen

import (
	"os"
	"testing"
)

const (
	testProject = "test-project"
	testService = "test-service"
	testVersion = "test-version"
)

func init() {
	os.Setenv(EnvKeyGoogleCloudProject, testProject)
	os.Setenv(EnvKeyServiceName, testService)
	os.Setenv(EnvKeyServiceVersion, testVersion)
}

func TestEnviron(t *testing.T) {
	Reload()
	if ProjectID != testProject {
		t.Fatalf("unexpected, expected: %s, actual: %#v", testProject, ProjectID)
	}
	if ServiceName != testService {
		t.Fatalf("unexpected, expected: %s, actual: %#v", testService, ServiceName)
	}
	if ServiceVersion != testVersion {
		t.Fatalf("unexpected, expected: %s, actual: %#v", testVersion, ServiceVersion)
	}
}
