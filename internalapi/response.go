// Deprecated: file was moved to ydb package
// To use it, rename the import to github.com/yandex-cloud/ydb-go-sdk/v2
package internalapi

import "github.com/yandex-cloud/ydb-go-sdk/v2/internal"

type Response = internal.Response

type OpResponse = internal.OpResponse

func WrapOpResponse(resp OpResponse) Response {
	return internal.WrapOpResponse(resp)
}

type NoOpResponse = internal.NoOpResponse

func WrapNoOpResponse(resp NoOpResponse) Response {
	return internal.WrapNoOpResponse(resp)
}
