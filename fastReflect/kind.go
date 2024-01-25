package fastReflect

const (
	invalidKind uint8 = iota
	boolKind
	intKind
	int8Kind
	int16Kind
	int32Kind
	int64Kind
	uintKind
	uint8Kind
	uint16Kind
	uint32Kind
	uint64Kind
	uintptrKind
	float32Kind
	float64Kind
	complex64Kind
	complex128Kind
	arrayKind
	chanKind
	funcKind
	interfaceKind
	mapKind
	ptrKind
	sliceKind
	stringKind
	structKind
	unsafePointerKind
)
