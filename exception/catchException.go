package exception

import (
	"fmt"
	"strings"

	"github.com/farseer-go/fs/color"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/trace"
)

type catchException struct {
	exp any
}

// TryCatch 异常时返回error
func TryCatch(fn func()) error {
	var err error
	Try(fn).CatchException(func(exp any) {
		err = fmt.Errorf("%+v", exp)
	})
	return err
}

// Try 执行有可能发生异常的代码块
func Try(fn func()) (catch *catchException) {
	catch = &catchException{}
	defer func() {
		catch.exp = recover()
		if catch.exp != nil {
			switch e := catch.exp.(type) {
			case WebException:
			default:
				// 如果使用了链路追踪，则记录异常
				if traceContext := trace.CurTraceContext.Get(); traceContext != nil {
					traceContext.Error(fmt.Errorf("%+v", e))
				} else {
					lstLogs := []string{color.Red("【异常】") + color.Red(fmt.Sprintf("%+v", e))}
					for index, exceptionStackDetail := range trace.GetCallerInfo() {
						lstLogs = append(lstLogs, fmt.Sprintf("\t%d、%s:%s %s", index+1, exceptionStackDetail.ExceptionCallFile, color.Yellow(exceptionStackDetail.ExceptionCallLine), color.Red(exceptionStackDetail.ExceptionCallFuncName)))
					}
					flog.Printf(strings.Join(lstLogs, "\n") + "\n")
				}
			}
		}
	}()
	fn()
	return
}

// CatchRefuseException 捕获RefuseException异常
func (catch *catchException) CatchRefuseException(expFn func(exp RefuseException)) *catchException {
	if catch.exp == nil {
		return catch
	}
	if exp, ok := catch.exp.(RefuseException); ok {
		catch.exp = Try(func() {
			expFn(exp)
		}).exp
	}
	return catch
}

// CatchStringException 捕获String异常
func (catch *catchException) CatchStringException(expFn func(exp string)) *catchException {
	if catch.exp == nil {
		return catch
	}
	if exp, ok := catch.exp.(string); ok {
		catch.exp = Try(func() {
			expFn(exp)
		}).exp
	}
	return catch
}

// CatchWebException 捕获WebException异常
func (catch *catchException) CatchWebException(expFn func(exp WebException)) *catchException {
	if catch.exp == nil {
		return catch
	}
	if exp, ok := catch.exp.(WebException); ok {
		catch.exp = Try(func() {
			expFn(exp)
		}).exp
	}
	return catch
}

// CatchException 捕获Any异常
func (catch *catchException) CatchException(expFn func(exp any)) {
	if catch.exp == nil {
		return
	}
	expFn(catch.exp)
}

// ThrowUnCatch 抛出未捕获的异常
func (catch *catchException) ThrowUnCatch() {
	if catch.exp != nil {
		panic(catch.exp)
	}
}
