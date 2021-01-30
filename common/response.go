package common

// 统一的response结构体
type Response struct {
	Status int16       `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func SuccessResponse(data interface{}, msg string) *Response {
	return &Response{
		Status: 0,
		Msg:    msg,
		Data:   data,
	}
}
