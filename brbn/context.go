package brbn

import "github.com/valyala/fasthttp"

type Context struct {
	Request *fasthttp.RequestCtx
	Params  map[string]string
}
