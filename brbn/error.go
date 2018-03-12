package brbn

// A general interface for any error that occurs during
// HTTP processing.
type HTTPError interface {
	Status() int
	Error() string
}

// A simple implementation for HTTPError.
type SimpleHTTPError struct {
	StatusCode int
	Message    string
}

func (e SimpleHTTPError) Status() int {
	return e.StatusCode
}

func (e SimpleHTTPError) Error() string {
	return e.Message
}

// Static declaration of common errors
var (
	Error404 = SimpleHTTPError{404, "Request Not Found, 🤷🏽 "}
	Error500 = SimpleHTTPError{500, "Internal Server Error, 😰 "}
	Error403 = SimpleHTTPError{403, "Unauthorized Access, ✋ "}
)
