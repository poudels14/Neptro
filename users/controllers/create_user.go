package controllers

import "github.com/poudels14/Neptro/brbn"

func CreateUser(ctxt *brbn.Context) (*brbn.Response, brbn.HTTPError) {
	return &brbn.Response{
		Data: []byte("Creating a user"),
	}, nil
}
