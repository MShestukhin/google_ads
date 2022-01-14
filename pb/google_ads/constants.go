package google_ads

const (
	ServiceName         = "google_ads"
	CompleteServiceName = "ai.somin.service." + ServiceName

	GoogleErrorCode = "google_ads_error" // Micro error has instagram api error details

	TokenExpiredErrorCode  = 401 // Token expired or revoked
	UnknownGoogleErrorCode = 501
	UnknownErrorCode       = 500
)
