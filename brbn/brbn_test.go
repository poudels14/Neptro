package brbn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultBrbn(t *testing.T) {
	b := New("TestServer", "0.0.0.0", "6666")
	assert.NotNil(t, b)
}

func TestChainingMiddleware(t *testing.T) {
	passthrough := func(handler Handler) Handler { return handler }
	b := New("TestServer", "0.0.0.0", "6666")
	oldLen := len(b.middlewares)
	b.Chain(passthrough).Chain(passthrough).Chain(passthrough)
	assert.Equal(t, 3, len(b.middlewares)-oldLen)
}

func TestChainingHandler(t *testing.T) {
	count := 0
	handler := func(c *Context) (*DataResponse, HTTPError) { return nil, nil }
	middleware := func(handler Handler) Handler {
		count += 1
		return handler
	}

	b := New("TestServer", "0.0.0.0", "6666")
	b.Chain(middleware, middleware, middleware)

	finalHandler := b.chainMiddleware(handler)
	finalHandler(&Context{})

	assert.Equal(t, 3, count)
}
