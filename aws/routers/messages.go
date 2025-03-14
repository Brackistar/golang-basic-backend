package routers

// General messages
const (
	routeBeginMsg string = "Begin process for path %s\n"
	routeEndMsg   string = "process ended"
	routeSuccess  string = "Success"
)

// General Errors

const (
	unmarshalErrorMsg string = "Error while unmarshaling to type %T from value %s\n"
	fldEmptyMsg       string = "Field \"%s\" is mandatory and cannot be empty"
	invFieldValMsg    string = "The value of field \"%s\" is incorrect %s"
)

// Register User Messages
//Errors
const (
	usrFailBodyMsg   string = "Error found on user data"
	usrRegFailErrMsg string = "An error occurred on user registration. Error : %s\n"
	usrRegFailMsg    string = "User registration failed"
	usrRegExistsMsg  string = "An user with the same E-mail already exists"
)

// Messages
const (
	usrInserted string = "User creation success. User ID: %s"
)

// Login Messages
// Errors
const (
	usrPssInvalidMsg string = "Invalid password"
	usrNotFoundMsg   string = "User not found"
)

// Messages
const (
	usrLoginSuccess string = "User logged in succesfuly"
)
