## What are the functions?
* exception
  * func
    * ThrowRefuseException （抛出RefuseException异常）
    * ThrowRefuseExceptionf （抛出RefuseException异常）
    * ThrowException （抛出Exception异常）
    * ThrowExceptionf （抛出Exception异常）
    * Try（执行有可能发生异常的代码块）
      * CatchRefuseException（捕获RefuseException异常）
      * CatchStringException（捕获String异常）
      * CatchException（捕获任意类型的异常）
      * ThrowUnCatch（异常没有捕获到时，向上层抛出异常）
      
## Getting Started

```go
try := exception.Try(func() {
    panic("panic throw")
})
try.CatchRefuseException(func(exp *exception.RefuseException) {
    flog.Warning(exp.Message)   // Type does not match, will not run
})
try.CatchStringException(func(exp string) {
    flog.Info(exp)  // this will run
})
try.CatchException(func(exp any) {
    flog.Error(exp) // StringException is match, will not run
})

// print: [Info] panic throw
```

```go
try := exception.Try(func() {
    exception.ThrowRefuseException("test is throw")
})
try.CatchStringException(func(exp string) {
    flog.Info(exp)  // Type does not match, will not run
})
try.CatchRefuseException(func(exp *exception.RefuseException) {
    flog.Warning(exp.Message)   // this will run
})
try.CatchException(func(exp any) {
    flog.Error(exp) // RefuseException is match, will not run
})

// print: [Warn] test is throw
```

```go
try := exception.Try(func() {
    exception.ThrowRefuseException("test is throw")
})
try.CatchStringException(func(exp string) {
    flog.Info(exp)  // Type does not match, will not run
})

try.ThrowUnCatch()  // will Throw Not Match RefuseException 
```