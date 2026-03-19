package trace

type TraceDetailRedis struct {
	RedisKey          string `json:",omitempty"` // redis key
	RedisField        string `json:",omitempty"` // hash field
	RedisRowsAffected int    `json:",omitempty"` // 影响行数
}

func (receiver *TraceDetailRedis) SetRows(rows int) {
	receiver.RedisRowsAffected = rows
}
