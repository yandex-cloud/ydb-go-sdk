package connect

import (
	"a.yandex-team.ru/kikimr/public/sdk/go/ydb/scheme"
	"context"
)

type schemeWrapper struct {
	ctx    context.Context
	client *scheme.Client
}

func newSchemeWrapper(ctx context.Context) *schemeWrapper {
	return &schemeWrapper{
		ctx:    ctx,
		client: &scheme.Client{},
	}
}
