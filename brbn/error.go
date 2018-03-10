package brbn

type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	return e.Message
}

// Static declaration of common errors
var (
	Error404 = HTTPError{404, "Request Not Found, 🤷🏽 "}
	Error500 = HTTPError{500, "Internal Server Error, 😰 "}
	Error403 = HTTPError{403, "Unauthorized Access, ✋ "}
)
