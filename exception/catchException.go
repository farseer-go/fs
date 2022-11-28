package exception

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
		catch.exp = Try(func() {
			expFn(&exp)
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
func (catch *catchException) CatchWebException(expFn func(exp *WebException)) *catchException {
	if catch.exp == nil {
		return catch
	}
	if exp, ok := catch.exp.(WebException); ok {
		catch.exp = Try(func() {
			expFn(&exp)
		}).exp
	}
	return catch
}

// CatchException 捕获Any异常
func (catch *catchException) CatchException(expFn func(exp any)) *catchException {
	if catch.exp == nil {
		return nil
	}
	catch.exp = Try(func() {
		expFn(catch.exp)
	}).exp
	return catch
}

// ThrowUnCatch 抛出未捕获的异常
func (catch *catchException) ThrowUnCatch() {
	if catch.exp != nil {
		panic(catch.exp)
	}
}
