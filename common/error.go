package common

import (
	"fmt"
	"net/http"
)

type ReqStat struct {
	HttpStatusCode int    `json:"httpCode"`
	Stat           string `json:"stat"`
	Message        string `json:"message"`
}

func (s ReqStat) Error() string {
	return fmt.Sprintf("Stat %s, Message %s", s.Stat, s.Message)
}

func NewReqStat(stat, message string, httpStatusCode ...int) ReqStat {
	statusCode := http.StatusOK
	if len(httpStatusCode) > 0 {
		statusCode = httpStatusCode[0]
	}
	return ReqStat{
		HttpStatusCode: statusCode,
		Stat:           stat,
		Message:        message,
	}
}

var (
	OK                   = NewReqStat("OK", "成功")
	ErrBadParams         = NewReqStat("ERR_BAD_PARAMS", "参数错误", http.StatusBadRequest)
	ErrServerException   = NewReqStat("ERR_SERVER_EXCEPTION", "系统服务异常", http.StatusInternalServerError)
	ErrUserAlreadyExists = NewReqStat("ERR_USER_ALREADY_EXISTS", "用户已存在")
	ErrUserNotFound      = NewReqStat("ERR_USER_NOT_FOUND", "用户不存在")
)

func Recover() {

}

func CheckError(err error) {
	if err == nil {
		return
	}
	unkown := NewReqStat("ERR_SERVER_EXCEPTION", err.Error(), http.StatusInternalServerError)
	panic(unkown)
}

func CheckErrorWithIgnore(err error, ignore ...error) {
	if err == nil {
		return
	}

	for _, ie := range ignore {
		if err == ie {
			return
		}
	}

	unkown := NewReqStat("ERR_SERVER_EXCEPTION", err.Error(), http.StatusInternalServerError)
	panic(unkown)
}

func Verify(exp bool, err error) {
	if !exp {
		return
	}

	if _, ok := err.(ReqStat); ok {
		panic(err)
	}

	unkown := NewReqStat("ERR_SERVER_EXCEPTION", err.Error(), http.StatusInternalServerError)
	panic(unkown)
}
