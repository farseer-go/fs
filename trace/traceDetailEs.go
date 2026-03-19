package trace

type TraceDetailEs struct {
	EsIndexName   string `json:",omitempty"` // 索引名称
	EsAliasesName string `json:",omitempty"` // 别名
}
