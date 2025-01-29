package constants

// Environment constant names
const (
	UrlPrefix  string = "URL_PREFIX"
	Bucket     string = "BUCKET_NAME"
	SecretName string = "SECRET_NAME"
)

// Error messages

const (
	EnvVaNotFoundErrorMsg string = "no environment variable \"%s\" found"
	FatalErrorMsg         string = "fatal error: %s"
	GenFatalErrorMsg      string = "an unexpected error occured, please try again in a few minutes"
)
