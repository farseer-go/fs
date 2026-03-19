package trace

type TraceDetailEtcd struct {
	EtcdKey     string `json:",omitempty"` // key
	EtcdLeaseID int64  `json:",omitempty"` // 租约ID
}
