package eumCallType

type Enum int

const (
	Http          Enum = iota // Http
	Grpc                      // Grpc
	Database                  // Database
	Redis                     // Redis
	Mq                        // Mq
	Elasticsearch             // Elasticsearch
	Etcd                      // Etcd
	Hand                      // Hand
	EventPublish              // EventPublish
)

func (receiver Enum) ToString() string {
	switch receiver {
	case Http:
		return "Http"
	case Grpc:
		return "Grpc"
	case Database:
		return "Database"
	case Redis:
		return "Redis"
	case Mq:
		return "Mq"
	case Elasticsearch:
		return "Elasticsearch"
	case Hand:
		return "Hand"
	case Etcd:
		return "Etcd"
	case EventPublish:
		return "EventPublish"
	}
	return ""
}
