package brbn

import "github.com/valyala/fasthttp"

// The main entry point to a brbn app.
// A request will go through a server pipeline that
// will go through a router that directs it to the proper handler,
// a middleware system that processes the request in some manner
// and finally the handler that creates a response to send back to
// the client.

// Engine is an instance of a brbn app.
type Engine struct {
	Router
}

// Starts a web server that is listening for requests.
func (e *Engine) Start() {
	fasthttp.ListenAndServe(":8080", e.Create())
}

// TODO: return default engine
func Default() *Engine {
	return &Engine{}
}
