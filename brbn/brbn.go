package brbn

import (
	"fmt"

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
	address    string
	port       string
	router     *Router
	server     *fasthttp.Server
	debug      bool
	middleware []MiddlewareFunc
}

type MiddlewareFunc func(Handler) Handler

func (b *Brbn) handleRequest(fCtxt *fasthttp.RequestCtx) {
	context := NewContext(fCtxt)

	method := string(fCtxt.Method())
	path := string(fCtxt.Path())
	handler := b.router.GetHandler(method, path)

	finalhandler := b.chainMiddleware(handler)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in handleRequest:", r)
			b.handleError(fCtxt, Error500)
		}
	}()

	if handler != nil {
		result := finalhandler(context)
		switch v := result.(type) {
		case HTTPError:
			b.handleError(fCtxt, v)
		case string:
			fmt.Fprintf(fCtxt, v)
		}
	} else {
		b.handleError(fCtxt, Error404)
	}
}

func (b *Brbn) handleError(ctx *fasthttp.RequestCtx, e HTTPError) {
	fmt.Fprintf(ctx, e.Message)
	ctx.SetStatusCode(e.Code)
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
	b.middleware = append(b.middleware, middleware...)
	return b
}

// Chains middleware together to create a final handler for a particular request
func (b *Brbn) chainMiddleware(handler Handler) Handler {
	if handler == nil {
		return nil
	}

	Log("Chaining middlewares...")
	middleware := b.middleware
	return func(c *Context) interface{} {
		h := handler
		for i := len(middleware) - 1; i >= 0; i -= 1 {
			h = middleware[i](h)
		}
		return h(c)
	}
}

// Starts a web server that is listening for requests.
func (b *Brbn) Start() {
	Log("Starting server: ")
	portStr := fmt.Sprintf(":%s", b.port)
	fasthttp.ListenAndServe(portStr, b.handleRequest)
}

// Stops a running web server
func (b *Brbn) Stop() {
	Log("Stopping server")
}

// Creates a new brbn instance with the given address and port
func New(address, port string) *Brbn {
	router := NewRouter()
	var middleware []MiddlewareFunc
	return &Brbn{
		address:    address,
		port:       port,
		router:     router,
		middleware: middleware,
	}
}
