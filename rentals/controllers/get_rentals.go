package controllers

import "github.com/poudels14/Neptro/brbn"

func GetRentals(ctxt *brbn.Context) (*brbn.Response, brbn.HTTPError) {
	return &brbn.Response{
		Data: []byte("Getting rentals"),
	}, nil
}
