package exception

type RefuseException struct {
	// 异常信息
	Message string
	r       any
}

type catchException struct {
	exp any
}

// Try 执行有可能发生异常的代码块
func Try(fn func()) (catch *catchException) {
	catch = &catchException{}
	defer func() {
		catch.exp = recover()
	}()
	fn()
	return
}

// CatchRefuseException 捕获RefuseException异常
func (catch *catchException) CatchRefuseException(expFn func(exp *RefuseException)) *catchException {
	if catch.exp == nil {
		return catch
	}
	if exp, ok := catch.exp.(RefuseException); ok {
		expFn(&exp)
		catch.exp = nil
		if exp.r != nil {
			catch.exp = exp.r
		}
	}
	return catch
}

// ContinueRecover 是否继续让下一个捕获继续处理
func (exp *RefuseException) ContinueRecover(r any) {
	exp.r = &r
}

// CatchStringException 捕获String异常
func (catch *catchException) CatchStringException(expFn func(exp string)) *catchException {
	if catch.exp == nil {
		return catch
	}
	if exp, ok := catch.exp.(string); ok {
		expFn(exp)
		catch.exp = nil
	}
	return catch
}

// CatchException 捕获Any异常
func (catch *catchException) CatchException(expFn func(exp any)) {
	if catch.exp == nil {
		return
	}
	expFn(catch.exp)
	catch.exp = nil
}
