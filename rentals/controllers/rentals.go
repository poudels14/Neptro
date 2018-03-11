package controllers

import (
	"time"

	"github.com/poudels14/Neptro/brbn"
	log "github.com/sirupsen/logrus"
)

func View(ctxt *brbn.Context) (*brbn.Response, brbn.HTTPError) {
	v := ctxt.Param("i")
	if v != nil {
		log.WithField("value", v.Int()).Info("Response value")
	}
	// panic("Something went wrong")
	time.Sleep(time.Duration(1) * time.Second)
	// return "Printing from View!"
	return nil, brbn.Error404
}

func Rental(ctxt *brbn.Context) (*brbn.Response, brbn.HTTPError) {
	return &brbn.Response{
		Data: []byte("Printing from Rental!"),
	}, nil
}
