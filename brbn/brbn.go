package brbn

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// The main entry point to a brbn app.
// A request will go through a server pipeline that
// will go through a router that directs it to the proper handler,
// a middleware system that processes the request in some manner
// and finally the handler that creates a response to send back to
// the client.

// Server is an instance of a brbn app.
type Server struct {
	Address  string
	Port     string
	routemap map[string]map[string]Handler
}

func (s *Server) handleRequest(ctx *fasthttp.RequestCtx) {
	request := &Request{}
	params := &Params{}

	handler := s.routemap[string(ctx.Method())][string(ctx.Path())]
	fmt.Println("Method map")
	fmt.Println(s.routemap[string(ctx.Method())])

	fmt.Println("Path map: " + string(ctx.Path()))
	fmt.Println(s.routemap[string(ctx.Method())][string(ctx.Path())])

	fmt.Println(handler)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			s.handleError(ctx, Error_500)
		}
	}()
	if handler != nil {
		if error, ok := handler(request, params).(HTTPError); ok {
			s.handleError(ctx, error)
		} else if data, ok := handler(request, params).(string); ok {
			fmt.Fprintf(ctx, data)
		}
	} else {
		s.handleError(ctx, Error_404)
	}
}

func (s *Server) handleError(ctx *fasthttp.RequestCtx, e HTTPError) {
	fmt.Fprintf(ctx, e.Message)
	ctx.SetStatusCode(e.Code)
}

func (s *Server) AddRoutes(routes []Route) {
	if s.routemap == nil {
		s.routemap = make(map[string]map[string]Handler, len(routes))
	}
	for _, route := range routes {
		fmt.Println(route)
		if s.routemap[route.Method] == nil {
			s.routemap[route.Method] = make(map[string]Handler)
		}
		s.routemap[route.Method][route.Path] = route.Handler
		fmt.Println(route.Method + " -> " + route.Path)
	}
}

func (s *Server) Chain(args ...Middleware) {
	Log("Chaining middlewares...")
}

// Starts a web server that is listening for requests.
func (s *Server) Start() {
	fmt.Println(s.routemap)

	Log("Starting server: ")
	fasthttp.ListenAndServe(":"+s.Port, s.handleRequest)
}

// Stops a running web server
func (s *Server) Stop() {
	//stop
}
