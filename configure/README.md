# Getting Started with configure
## Configuration file path
```
In the root of your executable
./farseer.yaml
```

## Get Configuration
```go
configString := configure.GetString("Database.default")
```

## Get child nodes
```go
config := configure.GetSubNodes("Database")  // return map[string]string
```

## Set default configuration
```go
// When the Database.test node, is not set, the default configuration is used
configure.SetDefault("Database.test", "DataType=MySql,PoolMaxSize=50,PoolMinSize=1,ConnectionString=root:xxxx@123456@tcp(mysql:3306)/test?charset=utf8&parseTime=True&loc=Local")
```