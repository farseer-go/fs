package testBenchmark

import (
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/fastReflect"
	"github.com/farseer-go/fs/parse"
	"testing"
	"time"
)

// BenchmarkConvertAll-12         446144	      2876 ns/op	     656 B/op	      30 allocs/op
// BenchmarkConvertAll-12    	  767450	      1527 ns/op	     424 B/op	      18 allocs/op
func BenchmarkConvertAll(b *testing.B) {
	b.ReportAllocs()
	timeNow := time.Now()
	dateTimeNow := dateTime.Now()
	unixMilli := time.UnixMilli(0)
	for i := 0; i < b.N; i++ {
		parse.Convert(1, 0)
		parse.Convert(1, "")
		parse.Convert(0, false)
		parse.Convert("8", int8(1))
		parse.Convert("8", uint(1))
		parse.Convert("8", float64(1))
		parse.Convert("123", "")
		parse.Convert("123", 0)
		parse.Convert("1,2,3", []int{})
		parse.Convert("4,5,6", []string{})
		parse.Convert("True", false)
		parse.Convert(true, "")
		parse.Convert(timeNow, unixMilli)
		parse.Convert(timeNow, dateTime.DateTime{})
		parse.Convert(dateTimeNow, unixMilli)
		parse.Convert("2023-09-15", dateTime.DateTime{})
	}
}

// BenchmarkConvert_1_0-12    	55259283	        20.54 ns/op	       0 B/op	       0 allocs/op
// BenchmarkConvert_1_0-12    	94150036	        12.57 ns/op	       0 B/op	       0 allocs/op
func BenchmarkConvert_1_0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parse.Convert(1, 0)
	}
}

// BenchmarkConvert_1_-12       20579169	        56.01 ns/op	      16 B/op	       1 allocs/op
// BenchmarkConvert_1_-12    	73151695	        16.07 ns/op	       0 B/op	       0 allocs/op
func BenchmarkConvert_1_(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parse.Convert(1, "")
	}
}

// BenchmarkConvert_0_false-12      52777993	        22.17 ns/op	       0 B/op	       0 allocs/op
// BenchmarkConvert_0_false-12    	86714840	        13.64 ns/op	       0 B/op	       0 allocs/op
func BenchmarkConvert_0_false(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parse.Convert(0, false)
	}
}

// BenchmarkConvert_8_int8_1-12     23033266	        51.93 ns/op	       0 B/op	       0 allocs/op
// BenchmarkConvert_8_int8_1-12    	47946500	        24.58 ns/op	       0 B/op	       0 allocs/op
func BenchmarkConvert_8_int8_1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parse.Convert("8", int8(1))
	}
}

// BenchmarkConvert_8_float64_1-12      13724061	        87.94 ns/op	      16 B/op	       2 allocs/op
// BenchmarkConvert_8_float64_1-12    	21466652	        55.50 ns/op	      16 B/op	       2 allocs/op
func BenchmarkConvert_8_float64_1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parse.Convert("8", float64(1))
	}
}

// BenchmarkConvert_123_-12     55135886	        22.83 ns/op	       0 B/op	       0 allocs/op
// BenchmarkConvert_123_-12    	92884807	        12.67 ns/op	       0 B/op	       0 allocs/op
func BenchmarkConvert_123_(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parse.Convert("123", "")
	}
}

// BenchmarkConvert_123_0-12        21935886	        55.41 ns/op	       0 B/op	       0 allocs/op
// BenchmarkConvert_123_0-12    	41578650	        28.64 ns/op	       0 B/op	       0 allocs/op
func BenchmarkConvert_123_0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parse.Convert("123", 0)
	}
}

// BenchmarkConvert_123_slice_int-12       	 1350770	       891.7 ns/op	     288 B/op	      16 allocs/op
// BenchmarkConvert_123_slice_int-12    	 2734504	       418.3 ns/op	     168 B/op	       7 allocs/op
func BenchmarkConvert_123_slice_int(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parse.Convert("1,2,3", []int{})
	}
}

// BenchmarkConvert_456_slice_string-12    	 8019180	       140.1 ns/op	      72 B/op	       2 allocs/op
// BenchmarkConvert_456_slice_string-12    	 8685973	       117.4 ns/op	      72 B/op	       2 allocs/op
func BenchmarkConvert_456_slice_string(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parse.Convert("4,5,6", []string{})
	}
}

// BenchmarkConvert_true_false-12       36805489	        33.05 ns/op	       0 B/op	       0 allocs/op
// BenchmarkConvert_true_false-12    	63328722	        18.37 ns/op	       0 B/op	       0 allocs/op
func BenchmarkConvert_true_false(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parse.Convert("True", false)
	}
}

// BenchmarkConvert_true_-12        53397604	        24.48 ns/op	       0 B/op	       0 allocs/op
// BenchmarkConvert_true_-12    	71195629	        14.24 ns/op	       0 B/op	       0 allocs/op
func BenchmarkConvert_true_(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parse.Convert(true, "")
	}
}

// BenchmarkConvert_time_time-12         6806041	       179.2 ns/op	      48 B/op	       2 allocs/op
// BenchmarkConvert_time_time-12    	15118432	        78.07 ns/op	      48 B/op	       2 allocs/op
func BenchmarkConvert_time_time(b *testing.B) {
	b.ReportAllocs()
	now := time.Now()
	unixMilli := time.UnixMilli(0)
	for i := 0; i < b.N; i++ {
		parse.Convert(now, unixMilli)
	}
}

// BenchmarkConvert_time_dateTime-12       	 6119569	       177.1 ns/op	      48 B/op	       2 allocs/op
// BenchmarkConvert_time_dateTime-12    	13504820	        79.19 ns/op	      48 B/op	       2 allocs/op
func BenchmarkConvert_time_dateTime(b *testing.B) {
	b.ReportAllocs()
	now := time.Now()
	for i := 0; i < b.N; i++ {
		parse.Convert(now, dateTime.DateTime{})
	}
}

// BenchmarkConvert_dateTime_time-12       	 6663381	       189.2 ns/op	      48 B/op	       2 allocs/op
// BenchmarkConvert_dateTime_time-12    	14563110	        80.72 ns/op	      48 B/op	       2 allocs/op
func BenchmarkConvert_dateTime_time(b *testing.B) {
	b.ReportAllocs()
	now := dateTime.Now()
	unixMilli := time.UnixMilli(0)
	for i := 0; i < b.N; i++ {
		parse.Convert(now, unixMilli)
	}
}

// BenchmarkConvert_string_dateTime-12     	 2564781	       477.2 ns/op	     120 B/op	       3 allocs/op
// BenchmarkConvert_string_dateTime-12    	 5185413	       207.0 ns/op	      24 B/op	       1 allocs/op
func BenchmarkConvert_string_dateTime(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parse.Convert("2023-09-15", dateTime.DateTime{})
	}
}

// BenchmarkConvert2-12                    	217430817	         5.386 ns/op	       0 B/op	       0 allocs/op
func BenchmarkConvert2(b *testing.B) {
	b.ReportAllocs()
	var a any = dateTime.DateTime{}
	fastReflect.PointerOf(a)
	//aVal := reflect.ValueOf(a)
	for i := 0; i < b.N; i++ {
		fastReflect.PointerOf(a)
	}
	return
	//println(aVal.Interface())
}
