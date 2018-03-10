package controllers

import (
	"time"

	"github.com/poudels14/Neptro/brbn"
	"github.com/valyala/fasthttp"
)

func View(ctxt *brbn.Context) interface{} {
	brbn.Log(fasthttp.AcquireRequest().URI())
	v := ctxt.Param("i")
	if v != nil {
		brbn.Log("Value of v: ")
		brbn.Log(v.Int())
	}
	// panic("Something went wrong")
	time.Sleep(time.Duration(1) * time.Second)
	return "Printing from View!"
	return brbn.Error404
}

func Rental(ctxt *brbn.Context) interface{} {
	return "Printing from Rental!"
}
