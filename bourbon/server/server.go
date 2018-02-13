package server

import (
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/poudels14/bourbon/router"

	"github.com/valyala/fasthttp"
)

type ServerStruct struct {
	address  string
	port     int
	routes   []router.RouteStruct
	routemap map[string]map[string]interface{}
}

func Server() *ServerStruct {
	return &ServerStruct{"0.0.0.0", 8080, []router.RouteStruct{}, make(map[string]map[string]interface{})}
}

func (s *ServerStruct) AddRoutes(routes []router.RouteStruct) {
	fmt.Println("Adding routes")
	for _, ele := range routes {
		if s.routemap[ele.Method] == nil {
			s.routemap[ele.Method] = make(map[string]interface{})
		}
		s.routemap[ele.Method][ele.Path] = ele.Controller
	}
}

func (s *ServerStruct) Start() {
	fmt.Println("Starting server")
	fmt.Println(s.routemap)

	requestHandler := func(ctxt *fasthttp.RequestCtx) {
		method := string(ctxt.Method())
		path := string(ctxt.Path())
		if s.routemap[method] != nil && s.routemap[method][path] != nil {
			fmt.Println(s.routemap[method][path])
			c := reflect.ValueOf(s.routemap[method][path])
			rargs := make([]reflect.Value, 0)
			// for i, a := range args {
			// 	rargs[i] = reflect.ValueOf(a)
			// }
			fmt.Println("Before calling")
			r := c.Call(rargs)
			fmt.Println("After calling")
			fmt.Fprintf(ctxt, string(r[0].Bytes()))
		} else {
			fmt.Fprintf(ctxt, "Hello, world!\n\n")
		}
	}

	if err := fasthttp.ListenAndServe(":"+strconv.Itoa(s.port), requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
