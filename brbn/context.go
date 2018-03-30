package brbn

import (
	"strconv"

	"github.com/valyala/fasthttp"
)

type Context struct {
	Request *fasthttp.RequestCtx
	Params  map[string]string
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
		Request: r,
	}
}

func (c *Context) Param(key string) *Value {
	if val, ok := c.Params[key]; ok {
		return &Value{val}
	}

	query := c.Request.QueryArgs()
	if query.Has(key) {
		return &Value{string(query.Peek(key))}
	}

	return nil
}
