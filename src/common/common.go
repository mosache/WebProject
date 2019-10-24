package common

import "net/http"

//APIResponse ApiResponse
type APIResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"message"`
	Data   interface{} `json:"data"`
}

//ResponseOK ok
func ResponseOK(data interface{}) APIResponse {
	resp := APIResponse{}
	resp.Status = http.StatusOK
	resp.Msg = ""
	resp.Data = data
	return resp
}
