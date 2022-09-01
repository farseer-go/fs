package exception

import "fmt"

// ThrowException 抛出Exception异常
func ThrowException(err string) {
	panic(err)
}

// ThrowExceptionf 抛出Exception异常
func ThrowExceptionf(format string, a ...any) {
	panic(fmt.Sprintf(format, a...))
}
