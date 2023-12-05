package model

import "net/http"

type APIResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func ResponseOk(data interface{}, status string) APIResponse {
    return APIResponse {
        Code: http.StatusOK,
        Status: status,
        Data: data,
    }
}

func ResponseNotFound(reason string) APIResponse {
    return APIResponse {
        Code: http.StatusNotFound,
        Status: reason,
        Data: reason,
    }
}

func ResponseBadRequest(reason string) APIResponse {
    return APIResponse {
        Code: http.StatusBadRequest,
        Status: reason,
        Data: reason,
    }
}
