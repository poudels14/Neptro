package middleware

import (
	"github.com/op/go-logging"
	"github.com/poudels14/Neptro/brbn"
)

var log = logging.MustGetLogger("brbn")

// Logs a request
func Logger(handler brbn.Handler) brbn.Handler {
	return func(c *brbn.Context) (*brbn.DataResponse, brbn.HTTPError) {
		ctx := c.FContext
		method := c.Method()
		log.Infof("Incoming request - %s %s", method, ctx.URI())

		if method == "POST" {
			log.Infof("Arguments: %v", ctx.PostArgs())
		}

		return handler(c)
	}
}
