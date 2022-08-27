//go:build go1.18
// +build go1.18

package ydbsql

import (
	"context"
	"fmt"
	"testing"

	"github.com/yandex-cloud/ydb-go-sdk/v2"
)

func TestRetryModes(t *testing.T) {
	type CanRetry struct {
		idempotentOperation    bool // after an error we must retry idempotent operation or no
		nonIdempotentOperation bool // after an error we must retry non-idempotent operation or no
	}
	type Case struct {
		err           error // given error
		deleteSession bool  // close session and delete from pool
		canRetry      CanRetry
	}
	errs := []Case{
		{
			err:           fmt.Errorf("unknown error"), // retryer given unknown error - we will not operationCompleted and will close session
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err:           context.DeadlineExceeded, // golang context deadline exceeded
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err:           context.Canceled, // golang context cancelled
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorUnknownCode,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorCanceled,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    true,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorUnknown,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorInvalidArgument,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorDeadlineExceeded,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorNotFound,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorAlreadyExists,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorPermissionDenied,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorResourceExhausted,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    true,
				nonIdempotentOperation: true,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorFailedPrecondition,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorAborted,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    true,
				nonIdempotentOperation: true,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorOutOfRange,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorUnimplemented,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorInternal,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    true,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorUnavailable,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    true,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorDataLoss,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.TransportError{
				Reason: ydb.TransportErrorUnauthenticated,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusUnknownStatus,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusBadRequest,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusUnauthorized,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusInternalError,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusAborted,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    true,
				nonIdempotentOperation: true,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusUnavailable,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    true,
				nonIdempotentOperation: true,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusOverloaded,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    true,
				nonIdempotentOperation: true,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusSchemeError,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusGenericError,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusTimeout,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusBadSession,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    true,
				nonIdempotentOperation: true,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusPreconditionFailed,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusAlreadyExists,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusNotFound,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    true,
				nonIdempotentOperation: true,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusSessionExpired,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusCancelled,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusUndetermined,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    true,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusUnsupported,
			},
			deleteSession: false,
			canRetry: CanRetry{
				idempotentOperation:    false,
				nonIdempotentOperation: false,
			},
		},
		{
			err: &ydb.OpError{
				Reason: ydb.StatusSessionBusy,
			},
			deleteSession: true,
			canRetry: CanRetry{
				idempotentOperation:    true,
				nonIdempotentOperation: true,
			},
		},
	}
	r := ydb.DefaultRetryChecker
	t.Run("wrapped", func(t *testing.T) {
		for _, test := range errs {
			t.Run(test.err.Error(), func(t *testing.T) {
				m := r.Check(mapBadSessionError(test.err))
				if m.MustRetry(true) != test.canRetry.idempotentOperation {
					t.Errorf("unexpected must retry idempotent operation status: %v, want: %v", m.MustRetry(true), test.canRetry.idempotentOperation)
				}
				if m.MustRetry(false) != test.canRetry.nonIdempotentOperation {
					t.Errorf("unexpected must retry non-idempotent operation status: %v, want: %v", m.MustRetry(false), test.canRetry.nonIdempotentOperation)
				}
				if m.MustDeleteSession() != test.deleteSession {
					t.Errorf("unexpected delete session status: %v, want: %v", m.MustDeleteSession(), test.deleteSession)
				}
			})
		}
	})
	t.Run("wrapped->unwrapped", func(t *testing.T) {
		for _, test := range errs {
			t.Run(test.err.Error(), func(t *testing.T) {
				m := r.Check(unwrapErrBadConn(mapBadSessionError(test.err)))
				if m.MustRetry(true) != test.canRetry.idempotentOperation {
					t.Errorf("unexpected must retry idempotent operation status: %v, want: %v", m.MustRetry(true), test.canRetry.idempotentOperation)
				}
				if m.MustRetry(false) != test.canRetry.nonIdempotentOperation {
					t.Errorf("unexpected must retry non-idempotent operation status: %v, want: %v", m.MustRetry(false), test.canRetry.nonIdempotentOperation)
				}
				if m.MustDeleteSession() != test.deleteSession {
					t.Errorf("unexpected delete session status: %v, want: %v", m.MustDeleteSession(), test.deleteSession)
				}
			})
		}
	})
}
