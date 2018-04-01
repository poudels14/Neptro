package controllers

import "github.com/poudels14/Neptro/brbn"

func GetUsers(ctx *brbn.Context) (*brbn.DataResponse, brbn.HTTPError) {
	ctx.Action = "users.getUsers"
	return nil, brbn.Error202
}
