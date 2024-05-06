package xerr

import "fmt"

// CodeMsg is a struct that contains a code and a message.
// It implements the error interface.
type CodeMsg struct {
	Code int
	Msg  string
}

func (c *CodeMsg) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", c.Code, c.Msg)
}

// New creates a new CodeMsg.
func New(code int, msg string) error {
	return &CodeMsg{Code: code, Msg: msg}
}

/**
常用通用固定错误
*/

type CodeError struct {
	errCode uint32
	errMsg  string
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

// 自定义错误码和错误信息
func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

// 使用定义好的错误码，返回错误码和错误提示
func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

// 服务器开小差了错误，一般用在未知错误
func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}
