package brbn

import (
	"github.com/julienschmidt/httprouter"
	"github.com/valyala/fasthttp"
)

type RouteInfo struct {
	Method  string
	Path    string
	Handler interface{} // TODO: come up with better signature
}

type Router struct {
	router     *httprouter.Router
	middleware Middleware
}

// Chains middleware and creates a request context for
// fasthttp.
func (r *Router) Create() fasthttp.RequestHandler {
	// TODO: chaiin middleware
	return nil
}

func (r *Router) addRoute(method, path, description string, handler interface{}) {
	// TODO: add route to httprouter
}

func (r *Router) Get(path, description string, handler interface{}) {
	r.addRoute("GET", path, description, handler)
}

func (r *Router) Post(path, description string, handler interface{}) {
	r.addRoute("POST", path, description, handler)
}
