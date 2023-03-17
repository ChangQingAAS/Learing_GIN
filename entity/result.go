package entity

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 定义错误码
const (
	CodeSuccess int = 1
	CodeError   int = -1
)

func (res *Result) SetCode(code int) *Result {
	res.Code = code
	return res
}

func (res *Result) SetMessage(msg string) *Result {
	res.Message = msg
	return res
}

func (res *Result) SetData(data interface{}) *Result {
	res.Data = data
	return res
}
