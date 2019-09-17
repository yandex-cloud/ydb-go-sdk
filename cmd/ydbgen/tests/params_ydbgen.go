// Code generated by ydbgen; DO NOT EDIT.

package tests

import (
	"strconv"

	"github.com/yandex-cloud/ydb-go-sdk"
	"github.com/yandex-cloud/ydb-go-sdk/table"
)

var (
	_ = strconv.Itoa
	_ = ydb.StringValue
	_ = table.NewQueryParameters
)

func (p *Params) QueryParameters() *table.QueryParameters {
	return table.NewQueryParameters(
		table.ValueParam("$name", ydb.OptionalValue(ydb.UTF8Value(p.Name))),
		table.ValueParam("$int16_to_uint32", ydb.Uint32Value(ydbConvI16ToU32(p.Int16ToUint32))),
		table.ValueParam("$int_to_int64", ydb.Int64Value(int64(p.IntToInt64))),
	)
}

func ydbConvI16ToU32(x int16) uint32 { 
	if x < 0 {
		panic("ydbgen: convassert: conversion of negative int16 to uint32")
	}
	return uint32(x)
}
