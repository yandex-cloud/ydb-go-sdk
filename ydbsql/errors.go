//go:build !go1.18
// +build !go1.18

package ydbsql

import (
	"database/sql/driver"
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
