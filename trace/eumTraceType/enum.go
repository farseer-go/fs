package eumTraceType

type Enum int

const (
	WebApi        Enum = iota // WebApi
	MqConsumer                // MQ消费
	QueueConsumer             // 本地消费
	FSchedule                 // 调度中心
	Task                      // 本地任务
	WatchKey                  // ETCD
	EventConsumer             // 事件消费
	WebSocket                 // WebSocket
)

func (e Enum) ToString() string {
	switch e {
	case WebApi:
		return "WebApi"
	case MqConsumer:
		return "MqConsumer"
	case QueueConsumer:
		return "QueueConsumer"
	case FSchedule:
		return "FSchedule"
	case Task:
		return "Task"
	case WatchKey:
		return "WatchKey"
	case EventConsumer:
		return "EventConsumer"
	}
	return ""
}
