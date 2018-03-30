package controllers

import (
	"encoding/json"

	"github.com/poudels14/Neptro/brbn"
	"github.com/poudels14/Neptro/users/store"
	"github.com/poudels14/Neptro/utils"
)

func GetUser(ctxt *brbn.Context) (*brbn.Response, brbn.HTTPError) {
	store, err := store.InitializeUserStore()
	utils.PanicIfError(err)

	id := ctxt.Param("id")

	if id != nil {
		user, err := store.Get(id.Int64())
		userJson, err := json.Marshal(user)
		utils.PanicIfError(err)

		return &brbn.Response{
			Data: userJson,
		}, nil
	} else {
		errorResponse := brbn.ErrorResponse{
			ErrorCode: 404,
			Msg:       "invalid parameter supplied",
		}

		errorJson, err := json.Marshal(errorResponse)
		utils.PanicIfError(err)

		// the response creation is a bit verbose, we need to make this
		// very simple for a controller to the point where they can simply
		// return the data object and brbn would do the rest
		return &brbn.Response{
			Data: errorJson,
		}, nil
	}
}
