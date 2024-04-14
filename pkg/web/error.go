package web

type HttpError struct {
	Code  int         `json:"code" example:"1" format:"int64"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty" example:"generic error!" format:"string"`
}
