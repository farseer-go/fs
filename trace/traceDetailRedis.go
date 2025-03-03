package trace

type TraceDetailRedis struct {
	RedisKey          string // redis key
	RedisField        string // hash field
	RedisRowsAffected int    // 影响行数
}

func (receiver *TraceDetailRedis) SetRows(rows int) {
	receiver.RedisRowsAffected = rows
}
