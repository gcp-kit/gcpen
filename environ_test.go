package gcpen

import (
	"os"
	"testing"
)

const (
	testProject     = "test-project"
	testGAEService  = "test-gae-service"
	testGAEVersion  = "test-gae-version"
	testRUNService  = "test-run-service"
	testRUNRevision = "test-run-revision"
)

func init() {
	os.Setenv(EnvKeyGoogleCloudProject, testProject)
	os.Setenv(EnvKeyGAEServiceName, testGAEService)
	os.Setenv(EnvKeyGAEServiceVersion, testGAEVersion)
	os.Setenv(EnvKeyRUNServiceName, testRUNService)
	os.Setenv(EnvKeyRUNServiceRevision, testRUNRevision)
}

func TestEnviron(t *testing.T) {
	Reload()
	if ProjectID != testProject {
		t.Fatalf("unexpected, expected: %s, actual: %#v", testProject, ProjectID)
	}
	if GAEServiceName != testGAEService {
		t.Fatalf("unexpected, expected: %s, actual: %#v", testGAEService, GAEServiceName)
	}
	if GAEServiceVersion != testGAEVersion {
		t.Fatalf("unexpected, expected: %s, actual: %#v", testGAEVersion, GAEServiceVersion)
	}
	if RUNServiceName != testRUNService {
		t.Fatalf("unexpected, expected: %s, actual: %#v", testRUNService, RUNServiceName)
	}
	if RUNServiceRevision != testRUNRevision {
		t.Fatalf("unexpected, expected: %s, actual: %#v", testRUNRevision, RUNServiceRevision)
	}
}
