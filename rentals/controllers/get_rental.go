package controllers

import "github.com/poudels14/Neptro/brbn"

func GetRental(ctx *brbn.Context) (*brbn.DataResponse, brbn.HTTPError) {
	ctx.Action = "rentals.getRental"
	return nil, brbn.Error202
}
