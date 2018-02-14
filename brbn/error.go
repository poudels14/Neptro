package brbn

type HTTPError struct {
	Code    int
	Message string
	Html    string
}

var Error_404 = HTTPError{404, "Not found", "<div>Page not found</div>"}

var Error_500 = HTTPError{500, "Internal Server Error", "<div>Internal Server Error</div>"}
