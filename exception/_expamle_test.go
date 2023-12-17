package exception

import "github.com/farseer-go/fs/flog"

func ExampleTry() {
	try := Try(func() {
		// 这里我们使用一个异常
		ThrowRefuseException("test is throw")
	})

	// 不会运行
	try.CatchStringException(func(exp string) {
		flog.Info(exp)
	})

	// 会运行
	try.CatchRefuseException(func(exp RefuseException) {
		flog.Warning(exp.Message)
	})

	// 不会运行，已经捕获到了
	try.CatchException(func(exp any) {
		_ = flog.Error(exp)
	})
}
