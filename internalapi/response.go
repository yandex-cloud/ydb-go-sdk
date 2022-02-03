// Deprecated: file was moved to ydb package
// To use it, rename the import to a.yandex-team.ru/kikimr/public/sdk/go/ydb
package internalapi

import "a.yandex-team.ru/kikimr/public/sdk/go/ydb/internal"

type Response = internal.Response

type OpResponse = internal.OpResponse

func WrapOpResponse(resp OpResponse) Response {
	return internal.WrapOpResponse(resp)
}

type NoOpResponse = internal.NoOpResponse

func WrapNoOpResponse(resp NoOpResponse) Response {
	return internal.WrapNoOpResponse(resp)
}
