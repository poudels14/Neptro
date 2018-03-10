package brbn

const (
	GET  = "GET"
	POST = "POST"
)

type Router struct {
	routemap map[string]*Route
}

type Route struct {
	Path     string
	handlers *methodhandlers
}

type methodhandlers struct {
	get  Handler
	post Handler
	// TODO: add more methods
}

// Handler function that performs some operation on the context
// TODO: determine if response should be explicitly returned and a more
// complex reponse writer object needs to be created similar to what "net/http" does
type Handler func(*Context) interface{}

type Response struct {
	Data  interface{}
	Error error
}

// Creates a new Router instance
func NewRouter() *Router {
	routemap := make(map[string]*Route)
	return &Router{routemap}
}

// Registers a new route (eventually we want to be able to add parameters to paths).
// Existence is not checked so methods will be overwritten.
func (r *Router) Add(method, path string, handler Handler) {
	var handlers *methodhandlers
	if route, ok := r.routemap[path]; ok {
		handlers = route.handlers
	} else {
		handlers = &methodhandlers{}
		r.routemap[path] = &Route{path, handlers}
	}

	switch method {
	case GET:
		handlers.get = handler
	case POST:
		handlers.post = handler
	}
}

// Retrieves the handler for the given method and path
func (r *Router) GetHandler(method, path string) Handler {
	if route, ok := r.routemap[path]; ok {
		handlers := route.handlers
		switch method {
		case GET:
			return handlers.get
		case POST:
			return handlers.post
		default:
			return nil
		}
	}

	return nil
}
