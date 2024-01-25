package fastReflect

import "reflect"

const (
	ListTypeString              = "collections.List["
	DictionaryTypeString        = "collections.Dictionary["
	PageListTypeString          = "collections.PageList["
	CollectionsTypeString       = "github.com/farseer-go/collections"
	CollectionsPrefixTypeString = "collections."
	DatetimeString              = "dateTime.DateTime"
	TimeString                  = "time.Time"
	DecimalString               = "decimal.Decimal"
	DomainSetString             = "data.DomainSet["
	TableSetString              = "data.TableSet["
	IndexSetString              = "elasticSearch.IndexSet["
)

// 数字类型
func isNumber(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}
