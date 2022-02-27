package entities

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseData struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ResponseError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
