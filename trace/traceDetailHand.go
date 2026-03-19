package trace

// TraceDetailHand 手动埋点
type TraceDetailHand struct {
	HandName string `json:",omitempty"` // 手动埋点名称
}

func (receiver *TraceDetailHand) SetName(name string) {
	receiver.HandName = name
}
