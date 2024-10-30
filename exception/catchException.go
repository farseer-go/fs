package exception

import (
	"fmt"

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
			switch catch.exp.(type) {
			case RefuseException, WebException:
			default:
				// 如果使用了链路追踪，则记录异常
				if traceContext := trace.CurTraceContext.Get(); traceContext != nil {
					traceContext.Error(fmt.Errorf("%s", catch.exp))
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
