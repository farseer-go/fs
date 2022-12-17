package eumLifecycle

type Enum int

const (
	Transient Enum = iota // 临时
	Single                // 单例
)
