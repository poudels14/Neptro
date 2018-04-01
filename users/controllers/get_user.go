package controllers

import (
	"github.com/poudels14/Neptro/brbn"
	"github.com/poudels14/Neptro/users/store"
	"github.com/poudels14/Neptro/utils"
)

func GetUser(ctxt *brbn.Context) (*brbn.DataResponse, brbn.HTTPError) {
	_, err := store.InitializeUserStore()
	utils.PanicIfError(err)

	// id := ctxt.Param("id")
	//
	// if id != nil {
	// 	user, err := store.Get(id.Int64())
	// 	userJson, err := json.Marshal(user)
	// 	utils.PanicIfError(err)
	//
	// 	return &brbn.Response{
	// 		Data: userJson,
	// 	}, nil
	// } else {
	//
	//
	// 	return &brbn.Response{
	// 		Data: errorJson,
	// 	}, nil
	// }

	return nil, brbn.Error202
}
