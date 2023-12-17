package exception

import "fmt"

type WebException struct {
	// 异常信息
	Message    string
	StatusCode int
}

// ThrowWebException 抛出WebException异常
func ThrowWebException(statusCode int, err string) {
	panic(WebException{StatusCode: statusCode, Message: err})
}

// ThrowWebExceptionf 抛出WebException异常
func ThrowWebExceptionf(statusCode int, format string, a ...any) {
	panic(WebException{StatusCode: statusCode, Message: fmt.Sprintf(format, a...)})
}

// ThrowWebExceptionBool 抛出WebException异常
func ThrowWebExceptionBool(isTrue bool, statusCode int, err string) {
	if isTrue {
		panic(WebException{StatusCode: statusCode, Message: err})
	}
}

// ThrowWebExceptionfBool 抛出WebException异常
func ThrowWebExceptionfBool(isTrue bool, statusCode int, format string, a ...any) {
	if isTrue {
		panic(WebException{StatusCode: statusCode, Message: fmt.Sprintf(format, a...)})
	}
}

// ThrowWebExceptionError 抛出WebException异常
func ThrowWebExceptionError(statusCode int, err error) {
	if err != nil {
		panic(WebException{StatusCode: statusCode, Message: err.Error()})
	}
}
