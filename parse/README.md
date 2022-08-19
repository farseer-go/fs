# Getting Started with parse
## Generic Type Conversion
```go
// Number 0, convert to bool, default to false if conversion fails
result := Convert(0, false)

// Number 1, convert to string, default to "" if conversion fails
result := Convert(1, "")

// Number 1, convert to int64, fail to convert then default to int64(0)
result := Convert(1, int64(0))

// string true, convert to bool, default to false if conversion fails
result := Convert("true", false)

// string 123, convert to int, default to 0 if conversion fails
result := Convert("123", 0)
```