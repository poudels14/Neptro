package middleware

import (
	"github.com/poudels14/Neptro/brbn"
	log "github.com/sirupsen/logrus"
)

// Logs a request
func Logger(handler brbn.Handler) brbn.Handler {
	return func(c *brbn.Context) (*brbn.Response, brbn.HTTPError) {
		request := c.Request
		log.WithFields(log.Fields{
			"uri":    request.URI(),
			"time":   request.Time(),
			"method": request.Method(),
			"query":  request.QueryArgs(),
			"args":   request.PostArgs(),
		}).Info("Incoming request")
		return handler(c)
	}
}
