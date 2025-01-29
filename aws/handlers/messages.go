package handlers

// General use messages
const (
	handleBeginMsg string = "Handling request %s > %s"
	rqstEndMsg     string = "Request handled"
)

// Error messages
const (
	badRqstMsg     string = "An error was found, and your request could not be processed"
	noTokenMsg     string = "The mandatory authorization token was not found"
	ftlErrMsg      string = "Fatal error on method %s, path: %s, error: %v"
	intErrMsg      string = "An unexpected error occurred, please try again later"
	failAuthMsg    string = "Token failed to authorise"
	unexErrAuthMsg string = "Error occurred during authorizaton, error: %s"
)
