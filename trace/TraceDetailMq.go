package trace

type TraceDetailMq struct {
	MqServer     string `json:",omitempty"` // MQ服务器地址
	MqExchange   string `json:",omitempty"` // 交换器名称
	MqRoutingKey string `json:",omitempty"` // 路由key
}
