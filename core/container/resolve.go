package container

// Resolve 从容器中获取实例
func Resolve[T any]() (t T) {
	_ = container.Resolve(&t)
	return
}
