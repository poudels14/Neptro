package brbn

import (
	"fmt"
	"runtime/debug"

	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

/*
	The main entry point to a brbn app. Brbn is an instance of a webserver.
	A request will go through a server pipeline that will through a router
	that directs it to the proper handler, a middleware sustem that processes
	the request in some manner and finally to the handler that creates a final
	response to send back to the client.
*/
type Brbn struct {
	address     string
	port        string
	router      *Router
	server      *fasthttp.Server
	debug       bool
	middlewares []MiddlewareFunc
}

type MiddlewareFunc func(Handler) Handler

func (b *Brbn) handleRequest(fCtxt *fasthttp.RequestCtx) {
	context := NewContext(fCtxt)

	method := string(fCtxt.Method())
	path := string(fCtxt.Path())
	handler, params := b.router.GetHandler(method, path)
	context.Params = params

	finalhandler := b.chainMiddleware(handler)

	defer func() {
		if r := recover(); r != nil {
			log.WithField("error", r).Error("Recovered an error in handleRequest")
			log.Error(debug.Stack())
			b.handleError(fCtxt, Error500)
		}
	}()

	if handler != nil {
		response, err := finalhandler(context)
		if response != nil {
			// TODO: expand this to make it more structured
			fmt.Fprintf(fCtxt, string(response.Data))
		} else {
			b.handleError(fCtxt, err)
		}
	} else {
		log.Warn("Could not find a handler")
		b.handleError(fCtxt, Error404)
	}
}

func (b *Brbn) handleError(ctx *fasthttp.RequestCtx, e HTTPError) {
	log.Error(e.Error())
	fmt.Fprintf(ctx, e.Error())
	ctx.SetStatusCode(e.Status())
}

// Adds a GET path to the router
func (b *Brbn) GET(path string, handler Handler) {
	b.add("GET", path, handler)
}

// Adds a POST path to the router
func (b *Brbn) POST(path string, handler Handler) {
	b.add("POST", path, handler)
}

func (b *Brbn) add(method, path string, handler Handler) {
	b.router.Add(method, path, handler)
}

// Appends a middleware to the chain
func (b *Brbn) Chain(middleware ...MiddlewareFunc) *Brbn {
	b.middlewares = append(b.middlewares, middleware...)
	return b
}

// Chains middleware together to create a final handler for a particular request
func (b *Brbn) chainMiddleware(handler Handler) Handler {
	if handler == nil {
		return nil
	}

	log.Info("Chaining middlewares")
	middleware := b.middlewares
	return func(c *Context) (*Response, HTTPError) {
		h := handler
		for i := len(middleware) - 1; i >= 0; i -= 1 {
			h = middleware[i](h)
		}
		return h(c)
	}
}

// Starts a web server that is listening for requests.
func (b *Brbn) Start() {
	log.Info("Starting brbn 🥃 ")
	portStr := fmt.Sprintf(":%s", b.port)
	fasthttp.ListenAndServe(portStr, b.handleRequest)
}

// Stops a running web server
func (b *Brbn) Stop() {
	log.Warn("🥃 Stopping brbn")
}

// Creates a new brbn instance with the given address and port
func New(address, port string) *Brbn {
	router := NewRouter()
	var middlewares []MiddlewareFunc
	return &Brbn{
		address:     address,
		port:        port,
		router:      router,
		middlewares: middlewares,
	}
}
