package trace

// TraceDetailHand 手动埋点
type TraceDetailHand struct {
	HandName string
}

func (receiver *TraceDetailHand) SetName(name string) {
	receiver.HandName = name
}
