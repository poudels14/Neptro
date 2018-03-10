package brbn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestHandler() Handler {
	return func(c *Context) interface{} { return nil }
}

func TestNewRouter(t *testing.T) {
	router := NewRouter()
	assert.NotNil(t, router)
}

func TestAddingGetMethod(t *testing.T) {
	router := NewRouter()
	handler := getTestHandler()
	router.Add("GET", "/items", handler)
	assert.Equal(t, 1, len(router.routemap))

	handlers := router.routemap["/items"].handlers

	assert.NotNil(t, handlers.get)
	assert.Nil(t, handlers.post)
}

func TestAddingPostMethod(t *testing.T) {
	router := NewRouter()
	handler := getTestHandler()
	router.Add("POST", "/items", handler)
	assert.Equal(t, 1, len(router.routemap))

	handlers := router.routemap["/items"].handlers

	assert.Nil(t, handlers.get)
	assert.NotNil(t, handlers.post)
}

func TestAddingGetAndPostMethod(t *testing.T) {
	router := NewRouter()
	handler := getTestHandler()
	router.Add("GET", "/items", handler)
	router.Add("POST", "/items", handler)
	assert.Equal(t, 1, len(router.routemap))

	handlers := router.routemap["/items"].handlers

	assert.NotNil(t, handlers.get)
	assert.NotNil(t, handlers.post)
}

func TestGettingHandler(t *testing.T) {
	router := NewRouter()
	handler := getTestHandler()
	router.Add("GET", "/items", handler)

	retrievedHandler := router.GetHandler("GET", "/items")

	assert.NotNil(t, retrievedHandler)
}
