// Code generated by ydbgen; DO NOT EDIT.

package tests

import (
	"strconv"

	"github.com/yandex-cloud/ydb-go-sdk/v2"
	"github.com/yandex-cloud/ydb-go-sdk/v2/table"
)

var (
	_ = strconv.Itoa
	_ = ydb.StringValue
	_ = table.NewQueryParameters
)

func (o *Optional) Scan(res *table.Result) (err error) {
	res.SeekItem("int64")
	res.Unwrap()
	if !res.IsNull() {
		x0 := int64(res.Int16())
		o.Int64.Set(x0)
	}

	res.SeekItem("str")
	{
		x0 := string(res.String())
		o.Str.Set(x0)
	}

	res.SeekItem("int32")
	o.Int32 = res.OInt32()

	return res.Err()
}

func (o *Optional) QueryParameters() *table.QueryParameters {
	var v0 ydb.Value
	{
		var v1 ydb.Value
		x0, ok0 := o.Int64.Get()
		if ok0 {
			v1 = ydb.OptionalValue(ydb.Int16Value(ydbConvI64ToI16(x0)))
		} else {
			v1 = ydb.NullValue(ydb.TypeInt16)
		}
		v0 = v1
	}
	var v1 ydb.Value
	{
		var v2 ydb.Value
		x0, ok0 := o.Str.Get()
		if ok0 {
			v2 = ydb.StringValue([]uint8(x0))
		} else {
			panic("ydbgen: no value for non-optional type")
		}
		v1 = v2
	}
	var v2 ydb.Value
	{
		vp0 := ydb.OptionalValue(ydb.Int32Value(o.Int32))
		v2 = vp0
	}
	return table.NewQueryParameters(
		table.ValueParam("$int64", v0),
		table.ValueParam("$str", v1),
		table.ValueParam("$int32", v2),
	)
}

func (o *Optional) StructValue() ydb.Value {
	var v0 ydb.Value
	{
		var v1 ydb.Value
		{
			var v2 ydb.Value
			x0, ok0 := o.Int64.Get()
			if ok0 {
				v2 = ydb.OptionalValue(ydb.Int16Value(ydbConvI64ToI16(x0)))
			} else {
				v2 = ydb.NullValue(ydb.TypeInt16)
			}
			v1 = v2
		}
		var v2 ydb.Value
		{
			var v3 ydb.Value
			x0, ok0 := o.Str.Get()
			if ok0 {
				v3 = ydb.StringValue([]uint8(x0))
			} else {
				panic("ydbgen: no value for non-optional type")
			}
			v2 = v3
		}
		var v3 ydb.Value
		{
			vp0 := ydb.OptionalValue(ydb.Int32Value(o.Int32))
			v3 = vp0
		}
		v0 = ydb.StructValue(
			ydb.StructFieldValue("int64", v1),
			ydb.StructFieldValue("str", v2),
			ydb.StructFieldValue("int32", v3),
		)
	}
	return v0
}

func (o *Optional) StructType() ydb.Type {
	var t0 ydb.Type
	{
		fs0 := make([]ydb.StructOption, 3)
		var t1 ydb.Type
		{
			tp0 := ydb.TypeInt16
			t1 = ydb.Optional(tp0)
		}
		fs0[0] = ydb.StructField("int64", t1)
		var t2 ydb.Type
		{
			tp0 := ydb.TypeString
			t2 = tp0
		}
		fs0[1] = ydb.StructField("str", t2)
		var t3 ydb.Type
		{
			tp0 := ydb.TypeInt32
			t3 = ydb.Optional(tp0)
		}
		fs0[2] = ydb.StructField("int32", t3)
		t0 = ydb.Struct(fs0...)
	}
	return t0
}

func ydbConvI64ToI16(x int64) int16 {
	const (
		bits = 16
		mask = (1 << (bits - 1)) - 1
	)
	var abs uint64
	{
		v := int64(x)
		m := v >> 63
		abs = uint64(v ^ m - m)
	}
	if abs&mask != abs {
		panic(
			"ydbgen: convassert: " + strconv.FormatInt(int64(x), 10) +
				" (type int64) overflows int16",
		)
	}
	return int16(x)
}

