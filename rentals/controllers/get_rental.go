package controllers

import "github.com/poudels14/Neptro/brbn"

func GetRental(ctxt *brbn.Context) (*brbn.Response, brbn.HTTPError) {
	return &brbn.Response{
		Data: []byte("Getting a rental"),
	}, nil
}
