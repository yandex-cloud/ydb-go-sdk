//go:build !go1.18
// +build !go1.18

package ydbsql

import (
	"database/sql/driver"
	"errors"

	"github.com/yandex-cloud/ydb-go-sdk/v2"
)

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
		return driver.ErrBadConn
	default:
		return err
	}
}

func unwrapErrBadConn(err error) error {
	if errors.Is(err, driver.ErrBadConn) {
		return &ydb.OpError{
			Reason: ydb.StatusBadSession,
		}
	}
	return err
}
