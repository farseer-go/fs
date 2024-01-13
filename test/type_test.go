package test

import (
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/types"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestType(t *testing.T) {
	_, b := types.IsSlice(reflect.ValueOf([]int{}))
	assert.True(t, b)
	_, b = types.IsSlice(reflect.ValueOf(""))
	assert.False(t, b)

	_, b = types.IsMap(reflect.ValueOf(map[string]string{}))
	assert.True(t, b)
	_, b = types.IsMap(reflect.ValueOf(""))
	assert.False(t, b)

	_, b = types.IsList(reflect.ValueOf(""))
	assert.False(t, b)
	_, b = types.IsDictionary(reflect.ValueOf(""))
	assert.False(t, b)

	_, b = types.IsPageList(reflect.ValueOf(""))
	assert.False(t, b)

	assert.False(t, types.IsCollections(reflect.TypeOf("")))

	assert.True(t, types.IsStruct(reflect.TypeOf(struct{}{})))
	assert.False(t, types.IsStruct(reflect.TypeOf(time.Now())))
	assert.False(t, types.IsStruct(reflect.TypeOf("")))

	assert.True(t, types.IsGoBasicType(reflect.TypeOf(int(1))))
	assert.True(t, types.IsGoBasicType(reflect.TypeOf(int8(1))))
	assert.True(t, types.IsGoBasicType(reflect.TypeOf(int16(1))))
	assert.True(t, types.IsGoBasicType(reflect.TypeOf(int32(1))))
	assert.True(t, types.IsGoBasicType(reflect.TypeOf(int64(1))))
	assert.True(t, types.IsGoBasicType(reflect.TypeOf(uint(1))))
	assert.True(t, types.IsGoBasicType(reflect.TypeOf(uint8(1))))
	assert.True(t, types.IsGoBasicType(reflect.TypeOf(uint16(1))))
	assert.True(t, types.IsGoBasicType(reflect.TypeOf(uint32(1))))
	assert.True(t, types.IsGoBasicType(reflect.TypeOf(uint64(1))))
	assert.True(t, types.IsGoBasicType(reflect.TypeOf(float32(1))))
	assert.True(t, types.IsGoBasicType(reflect.TypeOf(float64(1))))
	assert.True(t, types.IsGoBasicType(reflect.TypeOf("")))
	assert.True(t, types.IsGoBasicType(reflect.TypeOf(true)))
	assert.True(t, types.IsGoBasicType(reflect.TypeOf(time.Now())))
	assert.False(t, types.IsGoBasicType(reflect.TypeOf([]int{})))
	assert.False(t, types.IsGoBasicType(reflect.TypeOf(struct{}{})))
	assert.False(t, types.IsGoBasicType(reflect.TypeOf(map[string]string{})))

	_, b = types.IsEsIndexSet(reflect.ValueOf(""))
	assert.False(t, b)

	_, b = types.IsDataTableSet(reflect.ValueOf(""))
	assert.False(t, b)

	_, b = types.IsDataDomainSet(reflect.ValueOf(""))
	assert.False(t, b)

	assert.True(t, types.IsDtoModel([]reflect.Type{reflect.TypeOf(core.ApiResponseLongError403(""))}))
	assert.False(t, types.IsDtoModel([]reflect.Type{reflect.TypeOf(core.ApiResponseLongError403("")), reflect.TypeOf("")}))
	assert.False(t, types.IsDtoModel([]reflect.Type{reflect.TypeOf("")}))

	funcType := reflect.TypeOf(func(a int) string {
		return ""
	})

	assert.Equal(t, reflect.TypeOf(0).String(), types.GetInParam(funcType)[0].String())
	assert.Equal(t, reflect.TypeOf("").String(), types.GetOutParam(funcType)[0].String())

	var a any = 0
	assert.Equal(t, reflect.TypeOf(0).String(), types.GetRealType(reflect.ValueOf(&a)).String())
	assert.Panics(t, func() {
		types.GetRealType(reflect.ValueOf(nil))
	})

	v := 0
	assert.Equal(t, reflect.TypeOf(0).String(), types.GetRealType2(reflect.TypeOf(&v)).String())

	assert.False(t, types.IsDtoModelIgnoreInterface([]reflect.Type{}))
	assert.False(t, types.IsDtoModelIgnoreInterface([]reflect.Type{reflect.TypeOf(1)}))

	assert.False(t, types.IsDtoModelIgnoreInterface([]reflect.Type{reflect.TypeOf(sqlserver{}), reflect.TypeOf(sqlserver{})}))
	assert.True(t, types.IsDtoModelIgnoreInterface([]reflect.Type{reflect.TypeOf(sqlserver{})}))

	assert.Panics(t, func() {
		types.ListNew(nil)
	})
	assert.Panics(t, func() {
		types.ListAdd(reflect.ValueOf(""), nil)
	})
	assert.Panics(t, func() {
		types.GetListItemArrayType(nil)
	})
	assert.Panics(t, func() {
		types.GetListItemType(nil)
	})
	assert.Panics(t, func() {
		types.GetListToArray(reflect.ValueOf(""))
	})
	assert.Panics(t, func() {
		types.GetPageList(nil)
	})
}
