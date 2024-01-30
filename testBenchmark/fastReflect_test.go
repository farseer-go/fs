package testBenchmark

import (
	"github.com/farseer-go/fs/dateTime"
	"reflect"
	"testing"
)

func init() {
}

func BenchmarkFastReflect(b *testing.B) {
	b.ReportAllocs()
	var a any = dateTime.DateTime{}
	aVal := reflect.ValueOf(a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		aVal.Interface()
		//fastReflect.ValueOf(a)
		//reflect.ValueOf(a)
		//reflect.TypeOf(a)
	}
}
