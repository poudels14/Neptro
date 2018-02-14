package controllers

import (
	"github.com/poudels14/Neptro/brbn"
)

func View(request *brbn.Request, params *brbn.Params) interface{} {
	panic("Something went wrong")
	return "Printing from View!"
	return brbn.Error_404
}

func Rental(request *brbn.Request, params *brbn.Params) interface{} {
	return "Printing from Rental!"
}
