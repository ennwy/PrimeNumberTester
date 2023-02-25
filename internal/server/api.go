package server

var ErrInvalidInput = "the given input is invalid"

type ErrResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
}

type Response []bool
