package brbn

import (
	"encoding/json"
)

const (
	apiVersion = "0.0"
)

type ResponseJson struct {
	ApiVersion string            `json:"apiVersion"`
	Context    string            `json:"context"`
	Method     string            `json:"method"`
	Params     map[string]string `json:"params,omitempty"`
	Error      *ErrorResponse    `json:"error,omitempty"`
	Data       *DataResponse     `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}

type DataResponse struct {
	TotalItems int           `json:"totalItems"`
	Items      []interface{} `json:"items"`
}

// Helper method to build a single item data response
func BuildSingleDataResponse(data interface{}) *DataResponse {
	items := []interface{}{data}
	return &DataResponse{
		TotalItems: 1,
		Items:      items,
	}
}

// Sets common response data
func buildResponse(b *Brbn, ctx *Context) {
	fCtx := ctx.FContext
	// fCtx.Response.Header.Set("Server", b.name)
	fCtx.SetContentType("application/json")
}

// Gets the json byte slice for the given HTTPError
func renderErrorResponse(ctx *Context, e HTTPError) []byte {
	err := ErrorResponse{
		Code:    e.Status(),
		Message: e.Error(),
	}

	response := ResponseJson{
		ApiVersion: apiVersion,
		Context:    ctx.ClientContext(),
		Method:     ctx.Method(),
		Params:     ctx.SafeParams(),
		Error:      &err,
	}

	responseBytes, _ := json.Marshal(response)
	return responseBytes
}

// Gets the json byte slice for the given DataResponse
func renderDataResponse(ctx *Context, data *DataResponse) []byte {
	response := ResponseJson{
		ApiVersion: apiVersion,
		Context:    ctx.ClientContext(),
		Method:     ctx.Method(),
		Params:     ctx.SafeParams(),
		Data:       data,
	}

	responseBytes, _ := json.Marshal(response)
	return responseBytes
}
