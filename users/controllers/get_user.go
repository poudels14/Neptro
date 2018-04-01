package controllers

import (
	logging "github.com/op/go-logging"
	"github.com/poudels14/Neptro/brbn"
	"github.com/poudels14/Neptro/users/store"
	"github.com/poudels14/Neptro/utils"
)

var log = logging.MustGetLogger("UserServer")

func GetUser(ctx *brbn.Context) (*brbn.DataResponse, brbn.HTTPError) {
	ctx.Action = "users.getUser"
	store, err := store.InitializeUserStore()
	utils.PanicIfError(err)

	id := ctx.Param("id")
	if id != nil {
		user, err := store.Get(id.Int64())
		if user != nil {
			return brbn.BuildSingleDataResponse(user), nil
		} else if err != nil {
			log.Error(err)
			return nil, brbn.Error500
		}
	}

	return nil, brbn.CustomError404("User not found")
}
