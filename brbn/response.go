package brbn

import "time"

type ErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
}

type DataResponse struct {
	sentAt time.Time   `json:"sent_at"`
	data   interface{} `json:"data"`
}

type DataListResponse struct {
	sentAt time.Time     `json:"sent_at"`
	data   []interface{} `json:"data"`
}
