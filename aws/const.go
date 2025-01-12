package aws

// Environment constant names
const (
	urlPrefix  string = "URL_PREFIX"
	bucket     string = "BUCKET_NAME"
	secretName string = "SECRET_NAME"
)

// Error messages

const (
	envVaNotFoundErrorMsg string = "no environment variable \"%s\" found"
	fatalErrorMsg         string = "fatal error: %s"
	genFatalErrorMsg      string = "an unexpected error occured, please try again in a few minutes"
)
