package snowflake

import (
	"sync"
	"time"
)

type snowflake struct {
	sync.Mutex         // 锁
	timestamp    int64 // 时间戳 ，毫秒
	workerId     int64 // 工作节点
	dataCenterId int64 // 数据中心机房id
	sequence     int64 // 序列号
}

const (
	epoch             = int64(1577808000000)                           // 设置起始时间(时间戳/毫秒)：2020-01-01 00:00:00，有效期69年
	timestampBits     = uint(41)                                       // 时间戳占用位数
	dataCenterIdBits  = uint(2)                                        // 数据中心id所占位数
	workerIdBits      = uint(7)                                        // 机器id所占位数
	sequenceBits      = uint(12)                                       // 序列所占的位数
	timestampMax      = int64(-1 ^ (-1 << timestampBits))              // 时间戳最大值
	dataCenterIdMax   = int64(-1 ^ (-1 << dataCenterIdBits))           // 支持的最大数据中心id数量
	workerIdMax       = int64(-1 ^ (-1 << workerIdBits))               // 支持的最大机器id数量
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))               // 支持的最大序列id数量
	workerIdShift     = sequenceBits                                   // 机器id左移位数
	dataCenterIdShift = sequenceBits + workerIdBits                    // 数据中心id左移位数
	timestampShift    = sequenceBits + workerIdBits + dataCenterIdBits // 时间戳左移位数
)

var snowflakeSetting snowflake

// Init 使用前，全局初始化一次
func Init(dataCenterId int64, workerId int64) {
	snowflakeSetting = snowflake{
		workerId:     workerId,
		dataCenterId: dataCenterId,
	}
}

// GenerateId 生成唯一ID
func GenerateId() int64 {
	return snowflakeSetting.nextVal()
}

func (s *snowflake) nextVal() int64 {
	s.Lock()
	now := time.Now().UnixMilli() // 毫秒
	if s.timestamp == now {
		// 当同一时间戳（精度：毫秒）下多次生成id会增加序列号
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			// 如果当前序列超出12bit长度，则需要等待下一毫秒
			// 下一毫秒将使用sequence:0
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		// 不同时间戳（精度：毫秒）下直接使用序列号：0
		s.sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		s.Unlock()
		return 0
	}
	s.timestamp = now
	r := (t)<<timestampShift | (s.dataCenterId << dataCenterIdShift) | (s.workerId << workerIdShift) | (s.sequence)
	s.Unlock()
	return r
}
