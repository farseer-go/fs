package core

type PageList[TData any] struct {
	// 总页数
	RecordCount int64
	// 数据列表
	List []TData
}

// NewPageList 数据分页列表及总数
func NewPageList[TData any](list []TData, recordCount int64) PageList[TData] {
	return PageList[TData]{
		List:        list,
		RecordCount: recordCount,
	}
}
