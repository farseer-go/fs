package trace

type TraceDetailMq struct {
	MqServer     string // MQ服务器地址
	MqExchange   string // 交换器名称
	MqRoutingKey string // 路由key
}
