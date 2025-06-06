package exception

import (
	"fmt"

	"github.com/farseer-go/fs/flog"
)

// WebException 此异常类型在链路上不会被认定为异常，仅为业务上作为拒绝的继续往下执行的一种中断行为
type WebException struct {
	// 异常信息
	Message    string
	StatusCode int
}

// ThrowWebException 抛出WebException异常
func ThrowWebException(statusCode int, err string) {
	flog.Debugf("WebException:%s", err)
	panic(WebException{StatusCode: statusCode, Message: err})
}

// ThrowWebExceptionf 抛出WebException异常
func ThrowWebExceptionf(statusCode int, format string, a ...any) {
	err := fmt.Sprintf(format, a...)
	flog.Debugf("WebException:%s", err)
	panic(WebException{StatusCode: statusCode, Message: err})
}

// ThrowWebExceptionBool 抛出WebException异常
func ThrowWebExceptionBool(isTrue bool, statusCode int, err string) {
	if isTrue {
		flog.Debugf("WebException:%s", err)
		panic(WebException{StatusCode: statusCode, Message: err})
	}
}

// ThrowWebExceptionfBool 抛出WebException异常
func ThrowWebExceptionfBool(isTrue bool, statusCode int, format string, a ...any) {
	if isTrue {
		err := fmt.Sprintf(format, a...)
		flog.Debugf("WebException:%s", err)
		panic(WebException{StatusCode: statusCode, Message: err})
	}
}

// ThrowWebExceptionError 抛出WebException异常
func ThrowWebExceptionError(statusCode int, err error) {
	if err != nil {
		flog.Debugf("WebException:%v", err)
		panic(WebException{StatusCode: statusCode, Message: err.Error()})
	}
}
