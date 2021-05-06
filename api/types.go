package api

type Request struct {
	CODE int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func NewRequest(code int, message string, data interface{}) *Request {
	return &Request{
		code, message, data,
	}
}