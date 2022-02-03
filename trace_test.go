package ydb

import (
	"testing"

	"a.yandex-team.ru/kikimr/public/sdk/go/ydb/internal/tracetest"
)

func TestDriverTrace(t *testing.T) {
	tracetest.TestSingleTrace(t, DriverTrace{}, "DriverTrace")
}
