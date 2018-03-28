package controllers

import "github.com/poudels14/Neptro/brbn"

func GetUsers(ctxt *brbn.Context) (*brbn.Response, brbn.HTTPError) {
	return &brbn.Response{
		Data: []byte("Getting some users"),
	}, nil
}
