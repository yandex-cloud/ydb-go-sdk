package tests

import "a.yandex-team.ru/kikimr/public/sdk/go/ydb/opt"

//ydb:generate value,scan,params,type
type Optional struct {
	Int64 opt.Int64  `ydb:"type:int16?,conv:assert"`
	Str   opt.String `ydb:"type:string"`
	Int32 int32
}
