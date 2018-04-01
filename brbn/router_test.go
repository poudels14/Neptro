package brbn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestHandler() Handler {
	return func(c *Context) (*DataResponse, HTTPError) { return nil, nil }
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

func TestAddingOnlyGetMethod(t *testing.T) {
	router := NewRouter()
	handler := getTestHandler()
	router.Add("GET", "/items", handler)
	assert.Equal(t, 1, len(router.routemap))

	getHandler, _ := router.GetHandler("GET", "/items")
	postHandler, _ := router.GetHandler("POST", "/items")

	assert.NotNil(t, getHandler)
	assert.NotNil(t, postHandler)
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

	retrievedHandler, _ := router.GetHandler("GET", "/items")

	assert.NotNil(t, retrievedHandler)
}

func TestMatchingArgs(t *testing.T) {
	res, params := matches(parse("/items/:id"), parse("/items/10"))
	expectedParam := map[string]string{"id": "10"}
	assert.True(t, res)
	assert.Equal(t, expectedParam, params)

	res, params = matches(parse("/items/more"), parse("/items/10"))
	expectedParam = nil
	assert.False(t, res)
	assert.Equal(t, expectedParam, params)

	res, params = matches(parse("/items/more/items"), parse("/items/10/items"))
	expectedParam = nil
	assert.False(t, res)
	assert.Equal(t, expectedParam, params)

	res, params = matches(parse("/items/more/items/items"), parse("/items/10/items"))
	expectedParam = nil
	assert.False(t, res)
	assert.Equal(t, expectedParam, params)

	res, params = matches(parse("items/"), parse("/items"))
	expectedParam = map[string]string{}
	assert.True(t, res)
	assert.Equal(t, expectedParam, params)

	res, params = matches(parse("/items/:item/more/:adj"), parse("/items/ball/more/tall"))
	expectedParam = map[string]string{"item": "ball", "adj": "tall"}
	assert.True(t, res)
	assert.Equal(t, expectedParam, params)
}
