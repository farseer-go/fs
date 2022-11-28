package exception

import "fmt"

type RefuseException struct {
	// 异常信息
	Message string
	r       any
}

// ThrowRefuseException 抛出RefuseException异常
func ThrowRefuseException(err string) {
	panic(RefuseException{Message: err})
}

// ThrowRefuseExceptionf 抛出RefuseException异常
func ThrowRefuseExceptionf(format string, a ...any) {
	panic(RefuseException{Message: fmt.Sprintf(format, a...)})
}
