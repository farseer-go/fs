package exception

import "fmt"

type RefuseException struct {
	// 异常信息
	Message string
}

// ThrowRefuseException 抛出RefuseException异常
func ThrowRefuseException(err string) {
	panic(RefuseException{Message: err})
}

// ThrowRefuseExceptionf 抛出RefuseException异常
func ThrowRefuseExceptionf(format string, a ...any) {
	panic(RefuseException{Message: fmt.Sprintf(format, a...)})
}

// ThrowRefuseExceptionBool 抛出RefuseException异常
func ThrowRefuseExceptionBool(isTrue bool, err string) {
	if isTrue {
		panic(RefuseException{Message: err})
	}
}

// ThrowRefuseExceptionfBool 抛出RefuseException异常
func ThrowRefuseExceptionfBool(isTrue bool, format string, a ...any) {
	if isTrue {
		panic(RefuseException{Message: fmt.Sprintf(format, a...)})
	}
}

// ThrowRefuseExceptionError 抛出RefuseException异常
func ThrowRefuseExceptionError(err error) {
	if err != nil {
		panic(RefuseException{Message: err.Error()})
	}
}
