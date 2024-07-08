package tools

type HttpCode struct {
	Msg  string      `json:"message"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
