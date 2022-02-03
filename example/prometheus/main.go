package main

import "a.yandex-team.ru/kikimr/public/sdk/go/ydb/example/internal/cli"

func main() {
	cli.Run(new(Command))
}
