package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code int `json:"code"`
	msg string `json:"msg"`
	details[]string `json:"details"`
}
var codes=map[int]string{}//用来记录错误码1

func NewError(code int,msg string)*Error  {
	if _, ok:=codes[code]; ok{
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code]=msg
	return &Error{code:code,msg:msg}

}
func (e *Error)Error() string  {
	return fmt.Sprintf("错误码: %d,错误信息: %s",e.code,e.msg)
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}


func (e *Error) Msgf(args... interface{}) string {
	return fmt.Sprintf(e.msg,args...)
}
func (e *Error ) WithDetails(details ...string) *Error {
		newError :=*e//指针的指针是值
		newError.details=[]string{}
		for _, d:=range details{
			newError.details=append(newError.details,d)
		}
		return &newError//&值则是返回一个地址

}
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}

