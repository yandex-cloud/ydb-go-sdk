//go:build go1.18
// +build go1.18

package ydbsql

import (
	"database/sql/driver"

	"github.com/yandex-cloud/ydb-go-sdk/v2"
)

type badConnError struct {
	err error
}

func (e badConnError) Error() string {
	return "ydbsql: bad connection: " + e.err.Error()
}

func (e badConnError) Unwrap() error {
	return driver.ErrBadConn
}

func mapBadSessionError(err error) error {
	if err == nil {
		return nil
	}
	m := (&ydb.RetryChecker{}).Check(err)
	switch {
	case
		m.MustDeleteSession(),
		ydb.IsOpError(err, ydb.StatusOverloaded),
		ydb.IsOpError(err, ydb.StatusUnavailable),
		ydb.IsTransportError(err, ydb.TransportErrorResourceExhausted):
		return badConnError{err: err}
	default:
		return err
	}
}
