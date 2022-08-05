package elasticSearch

// elasticSearch配置结构
type elasticConfig struct {
	Server          string
	Username        string
	Password        string
	ReplicasCount   string
	ShardsCount     string
	RefreshInterval string
	IndexFormat     string
}
