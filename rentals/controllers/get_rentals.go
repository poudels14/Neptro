package controllers

import "github.com/poudels14/Neptro/brbn"

func GetRentals(ctx *brbn.Context) (*brbn.DataResponse, brbn.HTTPError) {
	ctx.Action = "rentals.getRentals"
	return nil, brbn.Error202
}
