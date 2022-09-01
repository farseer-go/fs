package exception

import "fmt"

// ThrowRefuseException 抛出RefuseException异常
func ThrowRefuseException(err string) {
	panic(RefuseException{Message: err})
}

// ThrowRefuseExceptionf 抛出RefuseException异常
func ThrowRefuseExceptionf(format string, a ...any) {
	panic(RefuseException{Message: fmt.Sprintf(format, a...)})
}
