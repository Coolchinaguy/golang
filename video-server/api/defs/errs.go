package defs

import "net/http"

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}


type ErroResponse struct {
	HttpSc int
	Error  Err
}

var (
	ErrorRequsetBodyParseFailed = ErroResponse{
		HttpSc:http.StatusBadRequest,
		Error:Err{
			Error:"Requset body is not correct",
			ErrorCode:"001",
		},
	}
	ErrorNotAuthUser = ErroResponse{
		HttpSc:http.StatusUnauthorized,
		Error:Err{
			Error:"User authentication failed",
			ErrorCode:"002",
		},
	}
)