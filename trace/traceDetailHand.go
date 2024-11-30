package trace

import (
	"fmt"
)

// TraceDetailHand 手动埋点
type TraceDetailHand struct {
	BaseTraceDetail
	Name string
}

func (receiver *TraceDetailHand) GetTraceDetail() *BaseTraceDetail {
	return &receiver.BaseTraceDetail
}

func (receiver *TraceDetailHand) ToString() string {
	return fmt.Sprintf("[%s]耗时：%s， %s", receiver.CallType.ToString(), receiver.UseTs.String(), receiver.Name)
}

func (receiver *TraceDetailHand) SetName(name string) {
	receiver.Name = name
}
