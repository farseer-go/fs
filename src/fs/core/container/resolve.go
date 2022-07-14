package container

// Resolve 从容器中获取实例
func Resolve[T any]() (t T, err error) {
	err = container.Resolve(&t)
	return
}
