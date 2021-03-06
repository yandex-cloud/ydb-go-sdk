package table

import (
	"github.com/yandex-cloud/ydb-go-sdk/internal/cache/lru"
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/yandex-cloud/ydb-go-sdk"
	"github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb"
	"github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Scheme"
	"github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Table"
	"github.com/yandex-cloud/ydb-go-sdk/internal"
	"github.com/yandex-cloud/ydb-go-sdk/testutil"
)

func TestStatementRemoveNotFound(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	prepared := make(chan struct{}, 1)
	b := StubBuilder{
		T: t,
		Handler: methodHandlers{
			testutil.TablePrepareDataQuery: func(req, res interface{}) error {
				prepared <- struct{}{}
				return nil
			},
			testutil.TableExecuteDataQuery: func(req, res interface{}) error {
				return &ydb.OpError{
					Reason: ydb.StatusNotFound,
				}
			},
		},
	}
	s, err := b.CreateSession(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assertPrepared := func() {
		const timeout = 500 * time.Millisecond
		select {
		case <-prepared:
		case <-time.After(timeout):
			t.Fatalf("no prepare after %s", timeout)
		}
	}
	assertNotPrepared := func() {
		const timeout = 500 * time.Millisecond
		select {
		case <-prepared:
			t.Fatalf("unexpected prepare")
		case <-time.After(timeout):
		}
	}
	prepare := func() *Statement {
		const yql = "SOME YQL TEXT"
		stmt, err := s.Prepare(ctx, yql)
		if err != nil {
			t.Fatal(err)
		}
		return stmt
	}

	stmt1 := prepare()
	assertPrepared()

	prepare()
	assertNotPrepared() // Used from cache.

	_, _, _ = stmt1.Execute(ctx, TxControl(), nil)
	prepare()
	assertPrepared()
}

func TestSessionKeepAlive(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		status Ydb_Table.KeepAliveResult_SessionStatus
		e      error
	)
	b := StubBuilder{
		T: t,
		Handler: methodHandlers{
			testutil.TableKeepAlive: func(req, res interface{}) error {
				r, _ := res.(*Ydb_Table.KeepAliveResult)
				r.SessionStatus = status

				return e
			},
		},
	}
	s, err := b.CreateSession(ctx)
	if err != nil {
		t.Fatal(err)
	}

	{
		e = errors.New("any error")
		_, err := s.KeepAlive(ctx)
		if err == nil {
			t.Fatal(err)
		}
	}
	{
		status, e = Ydb_Table.KeepAliveResult_SESSION_STATUS_READY, nil
		info, err := s.KeepAlive(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if info.Status != SessionReady {
			t.Fatalf("Result %v differ from, expectd %v", info.Status, SessionReady)
		}
	}
	{
		status, e = Ydb_Table.KeepAliveResult_SESSION_STATUS_BUSY, nil
		info, err := s.KeepAlive(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if info.Status != SessionBusy {
			t.Fatalf("Result %v differ from, expectd %v", info.Status, SessionBusy)
		}
	}
}

func TestSessionDescribeTable(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		result *Ydb_Table.DescribeTableResult
		e      error
	)
	b := StubBuilder{
		T: t,
		Handler: methodHandlers{
			testutil.TableDescribeTable: func(req, res interface{}) error {
				r, _ := res.(*Ydb_Table.DescribeTableResult)
				r.Reset()
				proto.Merge(r, result)
				return e
			},
		},
	}
	s, err := b.CreateSession(ctx)
	if err != nil {
		t.Fatal(err)
	}

	{
		e = errors.New("any error")
		_, err = s.DescribeTable(ctx, "")
		if err == nil {
			t.Fatal(err)
		}
	}
	{
		e = nil
		expect := Description{
			Name:       "testName",
			PrimaryKey: []string{"testKey"},
			Columns: []Column{
				{
					Name:   "testColumn",
					Type:   ydb.Void(),
					Family: "testFamily",
				},
			},
			KeyRanges: []KeyRange{
				{
					From: nil,
					To:   ydb.Int64Value(100500),
				},
				{
					From: ydb.Int64Value(100500),
					To:   nil,
				},
			},
			ColumnFamilies: []ColumnFamily{
				{
					Name:         "testFamily",
					Data:         StoragePool{},
					Compression:  ColumnFamilyCompressionLZ4,
					KeepInMemory: ydb.FeatureEnabled,
				},
			},
			Attributes: map[string]string{},
			ReadReplicaSettings: ReadReplicasSettings{
				Type:  ReadReplicasAnyAzReadReplicas,
				Count: 42,
			},
			StorageSettings: StorageSettings{
				TableCommitLog0:    StoragePool{Media: "m1"},
				TableCommitLog1:    StoragePool{Media: "m2"},
				External:           StoragePool{Media: "m3"},
				StoreExternalBlobs: ydb.FeatureEnabled,
			},
		}
		result = &Ydb_Table.DescribeTableResult{
			Self: &Ydb_Scheme.Entry{
				Name:                 expect.Name,
				Owner:                "",
				Type:                 0,
				EffectivePermissions: nil,
				Permissions:          nil,
			},
			Columns: []*Ydb_Table.ColumnMeta{
				{
					Name:   expect.Columns[0].Name,
					Type:   internal.TypeToYDB(expect.Columns[0].Type),
					Family: "testFamily",
				},
			},
			PrimaryKey: expect.PrimaryKey,
			ShardKeyBounds: []*Ydb.TypedValue{
				internal.ValueToYDB(expect.KeyRanges[0].To),
			},
			Indexes:    nil,
			TableStats: nil,
			ColumnFamilies: []*Ydb_Table.ColumnFamily{
				{
					Name:         "testFamily",
					Data:         nil,
					Compression:  Ydb_Table.ColumnFamily_COMPRESSION_LZ4,
					KeepInMemory: Ydb.FeatureFlag_ENABLED,
				},
			},
			ReadReplicasSettings: expect.ReadReplicaSettings.toYDB(),
			StorageSettings:      expect.StorageSettings.toYDB(),
		}

		d, err := s.DescribeTable(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(d, expect) {
			t.Fatalf("Result %+v differ from, expectd %+v", d, expect)
		}
	}
}

func TestSessionOperationModeOnExecuteDataQuery(t *testing.T) {
	for _, c := range []struct {
		srcMode ydb.OperationMode
		dstMode ydb.OperationMode
	}{
		{
			srcMode: ydb.OperationModeUnknown,
			dstMode: ydb.OperationModeSync,
		},
		{
			srcMode: ydb.OperationModeSync,
			dstMode: ydb.OperationModeSync,
		},
		{
			srcMode: ydb.OperationModeAsync,
			dstMode: ydb.OperationModeAsync,
		},
	} {
		s := Session{
			c: Client{
				Driver: &testutil.Driver{
					OnCall: func(ctx context.Context, m testutil.MethodCode, req, res interface{}) error {
						require.Equal(t, testutil.TableExecuteDataQuery, m)
						mode, ok := ydb.ContextOperationMode(ctx)
						require.True(t, ok)
						require.Equal(t, c.dstMode, mode)
						r := testutil.TableBeginTransactionResult{R: res}
						r.SetTransactionID("")
						return nil
					},
				},
			},
			qcache: lru.New(0),
		}
		ctx, cancel := context.WithTimeout(
			context.Background(),
			time.Second,
		)
		if c.srcMode != 0 {
			ctx = ydb.WithOperationMode(
				ctx,
				c.srcMode,
			)
		}
		defer cancel()
		_, _, err := s.Execute(
			ctx,
			TxControl(),
			"",
			NewQueryParameters(),
		)
		require.NoError(t, err)
	}
}

func TestClientCache(t *testing.T) {
	for _, test := range [...]struct {
		name                 string
		cacheSize            int
		prepareCount         int
		prepareRequestsCount int
	}{
		{
			name:                 "fixed query cache size, one request proxed to server",
			cacheSize:            10,
			prepareCount:         10,
			prepareRequestsCount: 1,
		},
		{
			name:                 "default query cache size, one request proxed to server",
			cacheSize:            0,
			prepareCount:         10,
			prepareRequestsCount: 1,
		},
		{
			name:                 "disabled query cache, all requests proxed to server",
			cacheSize:            -1,
			prepareCount:         10,
			prepareRequestsCount: 10,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			prepareRequestsCount := 0
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			client := &Client{
				Driver: &testutil.Driver{
					OnCall: func(ctx context.Context, m testutil.MethodCode, req, res interface{}) error {
						switch m {
						case testutil.TableCreateSession:
						case testutil.TablePrepareDataQuery:
							prepareRequestsCount++
						default:
							t.Fatalf("Unexpected method %d", m)

						}
						return nil
					},
				},
				MaxQueryCacheSize: test.cacheSize,
			}
			s, err := client.CreateSession(ctx)
			require.NoError(t, err)
			for i := 0; i < test.prepareCount; i++ {
				stmt, err := s.Prepare(ctx, "SELECT 1")
				require.NoError(t, err)
				require.NotNil(t, stmt)
			}
			require.Equal(t, test.prepareRequestsCount, prepareRequestsCount)
		})
	}
}
