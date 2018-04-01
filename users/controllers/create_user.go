package controllers

import "github.com/poudels14/Neptro/brbn"

func CreateUser(ctx *brbn.Context) (*brbn.DataResponse, brbn.HTTPError) {
	ctx.Action = "users.create"
	return nil, brbn.Error202
}
