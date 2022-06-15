package gcpen

const (
	EnvKeyProjectID          = "GCP_PROJECT"          // for Google Cloud Functions
	EnvKeyGoogleCloudProject = "GOOGLE_CLOUD_PROJECT" // for Google App Engine

	// EnvKeyServiceName will be deleted in a major version upgrade
	// Deprecated
	EnvKeyServiceName = "GAE_SERVICE" // for Google App Engine
	// EnvKeyServiceVersion will be deleted in a major version upgrade
	// Deprecated
	EnvKeyServiceVersion = "GAE_VERSION" // for Google App Engine

	EnvKeyGAEServiceName     = "GAE_SERVICE" // for Google App Engine
	EnvKeyGAEServiceVersion  = "GAE_VERSION" // for Google App Engine
	EnvKeyRUNServiceName     = "K_SERVICE"   // for Google Cloud Run
	EnvKeyRUNServiceRevision = "K_REVISION"  // for Google Cloud Run
)
