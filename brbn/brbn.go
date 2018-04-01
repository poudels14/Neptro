package brbn

import (
	"fmt"
	"runtime/debug"

	"github.com/fatih/color"
	"github.com/valyala/fasthttp"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("brbn")

//	The main entry point to a brbn app. Brbn is an instance of a webserver.
//	A request will go through a server pipeline that will through a router
//	that directs it to the proper handler, a middleware sustem that processes
//	the request in some manner and finally to the handler that creates a final
//	response to send back to the client.
type Brbn struct {
	address     string
	debug       bool
	middlewares []MiddlewareFunc
	name        string
	port        string
	router      *Router
	server      *fasthttp.Server
}

type MiddlewareFunc func(Handler) Handler

func (b *Brbn) handleRequest(fctx *fasthttp.RequestCtx) {
	ctx := NewContext(fctx)
	method := ctx.Method()
	path := ctx.Path()

	errorRecovery := func() {
		if r := recover(); r != nil {
			log.Error(debug.Stack())
			b.handleError(ctx, Error500)
		}
	}
	defer errorRecovery()

	handler, params := b.router.GetHandler(method, path)
	ctx.Params = params
	finalhandler := b.chainMiddleware(handler)

	response, err := finalhandler(ctx)
	if response != nil {
		buildResponse(b, ctx)
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(renderDataResponse(ctx, response))
	} else if err != nil {
		b.handleError(ctx, err)
	} else {
		log.Error("Controller returned nil response/error")
		b.handleError(ctx, Error500)
	}
}

// Returns a json response for the given error
func (b *Brbn) handleError(ctx *Context, e HTTPError) {
	log.Error(e)
	buildResponse(b, ctx)
	ctx.SetStatusCode(e.Status())
	ctx.SetBody(renderErrorResponse(ctx, e))
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

	middleware := b.middlewares
	return func(c *Context) (*DataResponse, HTTPError) {
		h := handler
		for i := len(middleware) - 1; i >= 0; i -= 1 {
			h = middleware[i](h)
		}
		return h(c)
	}
}

// Starts a web server that is listening for requests.
func (b *Brbn) Start() {
	address := fmt.Sprintf("%s:%s", b.address, b.port)
	log.Infof("Starting %s at %s", color.GreenString(b.name), address)

	server := &fasthttp.Server{
		Handler: b.handleRequest,
		Name:    b.name,
	}

	b.server = server
	if err := server.ListenAndServe(address); err != nil {
		log.Error(err)
	}
}

// Stops a running web server
func (b *Brbn) Stop() {
	// TODO: find a graceful way to stop a server
	log.Warningf("Stopping %s", b.name)
}

// Creates a new brbn instance with the given address and port
func New(name, address, port string) *Brbn {
	router := NewRouter()
	var middlewares []MiddlewareFunc
	return &Brbn{
		address:     address,
		middlewares: middlewares,
		name:        name,
		port:        port,
		router:      router,
	}
}
