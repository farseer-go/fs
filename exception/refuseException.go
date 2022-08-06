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

type catchException struct {
	r any
}

// Catch 捕获异常
func Catch() *catchException {
	return &catchException{r: recover()}
}

// RefuseException 捕获RefuseException异常
func (catch *catchException) RefuseException(expFn func(exp *RefuseException)) *catchException {
	if catch.r == nil {
		return catch
	}
	if exp, ok := catch.r.(RefuseException); ok {
		expFn(&exp)
		if exp.r != nil {
			catch.r = exp.r
		}
	}
	return catch
}

// ContinueRecover 是否继续让下一个捕获继续处理
func (exp *RefuseException) ContinueRecover(r any) {
	exp.r = &r
}

// String 捕获String异常
func (catch *catchException) String(expFn func(exp string)) *catchException {
	if catch.r == nil {
		return catch
	}
	if exp, ok := catch.r.(string); ok {
		expFn(exp)
	}
	return catch
}

// Any 捕获Any异常
func (catch *catchException) Any(expFn func(exp any)) {
	if catch.r == nil {
		return
	}
	expFn(catch.r)
}
