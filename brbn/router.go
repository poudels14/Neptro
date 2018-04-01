package brbn

import (
	"strings"

	"github.com/poudels14/Neptro/utils/lists"
)

const (
	GET  = "GET"
	POST = "POST"
)

type Router struct {
	routemap map[string]*Route
}

type Route struct {
	Path     string
	args     []string
	handlers *methodhandlers
}

type methodhandlers struct {
	get  Handler
	post Handler
	// TODO: add more methods
}

// Handler function that performs some operation on the context
// complex reponse writer object needs to be created similar to what "net/http" does
type Handler func(*Context) (*DataResponse, HTTPError)

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
		args := parse(path)
		r.routemap[path] = &Route{path, args, handlers}
	}

	switch method {
	case GET:
		handlers.get = handler
	case POST:
		handlers.post = handler
	}
}

// Retrieves the handler for the given method and path
func (r *Router) GetHandler(method, path string) (Handler, map[string]string) {
	pathArgs := parse(path)
	for _, v := range r.routemap {
		routeArgs := v.args
		ok, params := matches(routeArgs, pathArgs)
		if ok {
			handlers := v.handlers
			var handler Handler

			switch method {
			case GET:
				handler = handlers.get
			case POST:
				handler = handlers.post
			default:
				handler = nil
			}

			if handler != nil {
				return handler, params
			} else {
				return get404Handler(), nil
			}
		}
	}

	return get404Handler(), nil
}

// A default handler that simply returns an error 404 response
func get404Handler() Handler {
	return func(c *Context) (*DataResponse, HTTPError) {
		return nil, Error404
	}
}

// Returns the arguments of a path as a slice
func parse(path string) []string {
	args := strings.Split(path, "/")
	return lists.RemoveEmptyStrings(args)
}

// Determines if the given route and path arguments are a match.
// The route is the general path originally added to the router.
func matches(routeArgs, pathArgs []string) (bool, map[string]string) {
	params := make(map[string]string)

	if len(pathArgs) == len(routeArgs) {
		for i, r := range routeArgs {
			p := pathArgs[i]
			if r[0] == ':' {
				key := r[1:]
				params[key] = p
			} else if r != p {
				return false, nil
			}
		}
		return true, params
	}

	return false, nil
}
