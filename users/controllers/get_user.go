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

	// TODO: get id from param
	var id int64 = 1
	user, err := store.Get(id)

	// TODO: well defined json structure + handling for unknown id

	userJson, err := json.Marshal(user)
	utils.PanicIfError(err)

	return &brbn.Response{
		Data: userJson,
	}, nil
}
