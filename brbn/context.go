package brbn

import (
	"strconv"

	"github.com/valyala/fasthttp"
)

type Context struct {
	FContext *fasthttp.RequestCtx
	Params   map[string]string
}

// Retrieves the client context from the header if possible
func (c *Context) ClientContext() string {
	clientCtx := c.FContext.Request.Header.Peek("Context")
	if clientCtx != nil {
		return string(clientCtx)
	} else {
		return "Unknown"
	}
}

func (c *Context) Method() string {
	return string(c.FContext.Method())
}

func (c *Context) Path() string {
	return string(c.FContext.Path())
}

func (c *Context) SafeParams() map[string]string {
	return c.Params
}

// Passthrough methods for fast http context
func (c *Context) SetStatusCode(status int) {
	c.FContext.SetStatusCode(status)
}

func (c *Context) SetBody(body []byte) {
	c.FContext.SetBody(body)
}

//TODO(sagar): Can we use pool for this?
type Value struct {
	value string
}

func (v *Value) String() string {
	return v.value
}

func (v *Value) Int() int {
	result, err := strconv.Atoi(v.value)
	if err != nil {
		panic(err)
	}
	return result
}

func (v *Value) Int64() int64 {
	result, err := strconv.Atoi(v.value)
	if err != nil {
		panic(err)
	}
	return int64(result)
}

func (v *Value) Float() float64 {
	result, err := strconv.ParseFloat(v.value, 64)
	if err != nil {
		panic(err)
	}
	return result
}

func (v *Value) Bool() bool {
	result, err := strconv.ParseBool(v.value)
	if err != nil {
		panic(err)
	}
	return result
}

//TODO(sagar): use pool for this
func NewContext(r *fasthttp.RequestCtx) *Context {
	return &Context{
		FContext: r,
	}
}

func (c *Context) Param(key string) *Value {
	if val, ok := c.Params[key]; ok {
		return &Value{val}
	}

	query := c.FContext.QueryArgs()
	if query.Has(key) {
		return &Value{string(query.Peek(key))}
	}

	return nil
}
