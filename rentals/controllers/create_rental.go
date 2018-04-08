package controllers

import "github.com/poudels14/Neptro/brbn"

func CreateRental(ctx *brbn.Context) (*brbn.DataResponse, brbn.HTTPError) {
	ctx.Action = "rentals.create"
	return nil, brbn.Error202
}
