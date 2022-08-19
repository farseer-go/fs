# Getting Started with ioc

```go
type ITest interface {
    TestCall() string
}

type testStruct struct {
}

func (r testStruct) TestCall() string {
	return "hello world"
}
```

## Registration
### Default Registration
```go
// Register to the ITest interface using a single instance
Register(func() ITest { return testStruct{} })
```

### Use func Registration
```go
// Name：Set ioc aliases to distinguish between different instances of the same ITest
// Transient：Each time it is used, reinitialize, temporary instance
Use[ITest](func() ITest { return testStruct{} }).Transient().Name("test1").Register()
```

### Use Instance Registration
```go
// Register with an existing instance
test := testStruct{}
Use[ITest](test).Name("test1").Transient().Register()
```

## Resolve
### When there is no alias
```go
instance:= Resolve[ITest]()
```

### When there is an alias
```go
instance := ResolveName[ITest]("test")
```