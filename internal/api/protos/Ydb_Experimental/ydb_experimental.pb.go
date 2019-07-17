// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_experimental.proto

package Ydb_Experimental

import (
	Ydb "github.com/yandex-cloud/ydb-go-sdk/internal/api/protos/Ydb"
	Ydb_Issue "github.com/yandex-cloud/ydb-go-sdk/internal/api/protos/Ydb_Issue"
	Ydb_Operations "github.com/yandex-cloud/ydb-go-sdk/internal/api/protos/Ydb_Operations"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UploadRowsRequest struct {
	Table                string                          `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Rows                 *Ydb.TypedValue                 `protobuf:"bytes,2,opt,name=rows,proto3" json:"rows,omitempty"`
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,3,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *UploadRowsRequest) Reset()         { *m = UploadRowsRequest{} }
func (m *UploadRowsRequest) String() string { return proto.CompactTextString(m) }
func (*UploadRowsRequest) ProtoMessage()    {}
func (*UploadRowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{0}
}

func (m *UploadRowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRowsRequest.Unmarshal(m, b)
}
func (m *UploadRowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRowsRequest.Marshal(b, m, deterministic)
}
func (m *UploadRowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRowsRequest.Merge(m, src)
}
func (m *UploadRowsRequest) XXX_Size() int {
	return xxx_messageInfo_UploadRowsRequest.Size(m)
}
func (m *UploadRowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRowsRequest proto.InternalMessageInfo

func (m *UploadRowsRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *UploadRowsRequest) GetRows() *Ydb.TypedValue {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *UploadRowsRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

type UploadRowsResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *UploadRowsResponse) Reset()         { *m = UploadRowsResponse{} }
func (m *UploadRowsResponse) String() string { return proto.CompactTextString(m) }
func (*UploadRowsResponse) ProtoMessage()    {}
func (*UploadRowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{1}
}

func (m *UploadRowsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRowsResponse.Unmarshal(m, b)
}
func (m *UploadRowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRowsResponse.Marshal(b, m, deterministic)
}
func (m *UploadRowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRowsResponse.Merge(m, src)
}
func (m *UploadRowsResponse) XXX_Size() int {
	return xxx_messageInfo_UploadRowsResponse.Size(m)
}
func (m *UploadRowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRowsResponse proto.InternalMessageInfo

func (m *UploadRowsResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type UploadRowsResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRowsResult) Reset()         { *m = UploadRowsResult{} }
func (m *UploadRowsResult) String() string { return proto.CompactTextString(m) }
func (*UploadRowsResult) ProtoMessage()    {}
func (*UploadRowsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{2}
}

func (m *UploadRowsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRowsResult.Unmarshal(m, b)
}
func (m *UploadRowsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRowsResult.Marshal(b, m, deterministic)
}
func (m *UploadRowsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRowsResult.Merge(m, src)
}
func (m *UploadRowsResult) XXX_Size() int {
	return xxx_messageInfo_UploadRowsResult.Size(m)
}
func (m *UploadRowsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRowsResult.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRowsResult proto.InternalMessageInfo

type ReadColumnsRequest struct {
	Table                string                          `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Columns              []string                        `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
	FromKey              []byte                          `protobuf:"bytes,3,opt,name=from_key,json=fromKey,proto3" json:"from_key,omitempty"`
	FromKeyInclusive     bool                            `protobuf:"varint,4,opt,name=from_key_inclusive,json=fromKeyInclusive,proto3" json:"from_key_inclusive,omitempty"`
	ToKey                []byte                          `protobuf:"bytes,5,opt,name=to_key,json=toKey,proto3" json:"to_key,omitempty"`
	ToKeyInclusive       bool                            `protobuf:"varint,6,opt,name=to_key_inclusive,json=toKeyInclusive,proto3" json:"to_key_inclusive,omitempty"`
	MaxRows              uint64                          `protobuf:"varint,7,opt,name=max_rows,json=maxRows,proto3" json:"max_rows,omitempty"`
	MaxBytes             uint64                          `protobuf:"varint,8,opt,name=max_bytes,json=maxBytes,proto3" json:"max_bytes,omitempty"`
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,9,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ReadColumnsRequest) Reset()         { *m = ReadColumnsRequest{} }
func (m *ReadColumnsRequest) String() string { return proto.CompactTextString(m) }
func (*ReadColumnsRequest) ProtoMessage()    {}
func (*ReadColumnsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{3}
}

func (m *ReadColumnsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadColumnsRequest.Unmarshal(m, b)
}
func (m *ReadColumnsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadColumnsRequest.Marshal(b, m, deterministic)
}
func (m *ReadColumnsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadColumnsRequest.Merge(m, src)
}
func (m *ReadColumnsRequest) XXX_Size() int {
	return xxx_messageInfo_ReadColumnsRequest.Size(m)
}
func (m *ReadColumnsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadColumnsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadColumnsRequest proto.InternalMessageInfo

func (m *ReadColumnsRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *ReadColumnsRequest) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *ReadColumnsRequest) GetFromKey() []byte {
	if m != nil {
		return m.FromKey
	}
	return nil
}

func (m *ReadColumnsRequest) GetFromKeyInclusive() bool {
	if m != nil {
		return m.FromKeyInclusive
	}
	return false
}

func (m *ReadColumnsRequest) GetToKey() []byte {
	if m != nil {
		return m.ToKey
	}
	return nil
}

func (m *ReadColumnsRequest) GetToKeyInclusive() bool {
	if m != nil {
		return m.ToKeyInclusive
	}
	return false
}

func (m *ReadColumnsRequest) GetMaxRows() uint64 {
	if m != nil {
		return m.MaxRows
	}
	return 0
}

func (m *ReadColumnsRequest) GetMaxBytes() uint64 {
	if m != nil {
		return m.MaxBytes
	}
	return 0
}

func (m *ReadColumnsRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

type ReadColumnsResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ReadColumnsResponse) Reset()         { *m = ReadColumnsResponse{} }
func (m *ReadColumnsResponse) String() string { return proto.CompactTextString(m) }
func (*ReadColumnsResponse) ProtoMessage()    {}
func (*ReadColumnsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{4}
}

func (m *ReadColumnsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadColumnsResponse.Unmarshal(m, b)
}
func (m *ReadColumnsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadColumnsResponse.Marshal(b, m, deterministic)
}
func (m *ReadColumnsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadColumnsResponse.Merge(m, src)
}
func (m *ReadColumnsResponse) XXX_Size() int {
	return xxx_messageInfo_ReadColumnsResponse.Size(m)
}
func (m *ReadColumnsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadColumnsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadColumnsResponse proto.InternalMessageInfo

func (m *ReadColumnsResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type ReadColumnsResult struct {
	Blocks               [][]byte `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Eof                  bool     `protobuf:"varint,2,opt,name=eof,proto3" json:"eof,omitempty"`
	LastKey              []byte   `protobuf:"bytes,3,opt,name=last_key,json=lastKey,proto3" json:"last_key,omitempty"`
	LastKeyInclusive     bool     `protobuf:"varint,4,opt,name=last_key_inclusive,json=lastKeyInclusive,proto3" json:"last_key_inclusive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadColumnsResult) Reset()         { *m = ReadColumnsResult{} }
func (m *ReadColumnsResult) String() string { return proto.CompactTextString(m) }
func (*ReadColumnsResult) ProtoMessage()    {}
func (*ReadColumnsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{5}
}

func (m *ReadColumnsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadColumnsResult.Unmarshal(m, b)
}
func (m *ReadColumnsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadColumnsResult.Marshal(b, m, deterministic)
}
func (m *ReadColumnsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadColumnsResult.Merge(m, src)
}
func (m *ReadColumnsResult) XXX_Size() int {
	return xxx_messageInfo_ReadColumnsResult.Size(m)
}
func (m *ReadColumnsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadColumnsResult.DiscardUnknown(m)
}

var xxx_messageInfo_ReadColumnsResult proto.InternalMessageInfo

func (m *ReadColumnsResult) GetBlocks() [][]byte {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *ReadColumnsResult) GetEof() bool {
	if m != nil {
		return m.Eof
	}
	return false
}

func (m *ReadColumnsResult) GetLastKey() []byte {
	if m != nil {
		return m.LastKey
	}
	return nil
}

func (m *ReadColumnsResult) GetLastKeyInclusive() bool {
	if m != nil {
		return m.LastKeyInclusive
	}
	return false
}

type GetShardLocationsRequest struct {
	TabletIds            []uint64                        `protobuf:"varint,1,rep,packed,name=tablet_ids,json=tabletIds,proto3" json:"tablet_ids,omitempty"`
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,9,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *GetShardLocationsRequest) Reset()         { *m = GetShardLocationsRequest{} }
func (m *GetShardLocationsRequest) String() string { return proto.CompactTextString(m) }
func (*GetShardLocationsRequest) ProtoMessage()    {}
func (*GetShardLocationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{6}
}

func (m *GetShardLocationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShardLocationsRequest.Unmarshal(m, b)
}
func (m *GetShardLocationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShardLocationsRequest.Marshal(b, m, deterministic)
}
func (m *GetShardLocationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShardLocationsRequest.Merge(m, src)
}
func (m *GetShardLocationsRequest) XXX_Size() int {
	return xxx_messageInfo_GetShardLocationsRequest.Size(m)
}
func (m *GetShardLocationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShardLocationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShardLocationsRequest proto.InternalMessageInfo

func (m *GetShardLocationsRequest) GetTabletIds() []uint64 {
	if m != nil {
		return m.TabletIds
	}
	return nil
}

func (m *GetShardLocationsRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

type GetShardLocationsResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetShardLocationsResponse) Reset()         { *m = GetShardLocationsResponse{} }
func (m *GetShardLocationsResponse) String() string { return proto.CompactTextString(m) }
func (*GetShardLocationsResponse) ProtoMessage()    {}
func (*GetShardLocationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{7}
}

func (m *GetShardLocationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShardLocationsResponse.Unmarshal(m, b)
}
func (m *GetShardLocationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShardLocationsResponse.Marshal(b, m, deterministic)
}
func (m *GetShardLocationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShardLocationsResponse.Merge(m, src)
}
func (m *GetShardLocationsResponse) XXX_Size() int {
	return xxx_messageInfo_GetShardLocationsResponse.Size(m)
}
func (m *GetShardLocationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShardLocationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetShardLocationsResponse proto.InternalMessageInfo

func (m *GetShardLocationsResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type TabletInfo struct {
	TabletId             uint64   `protobuf:"varint,1,opt,name=tablet_id,json=tabletId,proto3" json:"tablet_id,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TabletInfo) Reset()         { *m = TabletInfo{} }
func (m *TabletInfo) String() string { return proto.CompactTextString(m) }
func (*TabletInfo) ProtoMessage()    {}
func (*TabletInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{8}
}

func (m *TabletInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TabletInfo.Unmarshal(m, b)
}
func (m *TabletInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TabletInfo.Marshal(b, m, deterministic)
}
func (m *TabletInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TabletInfo.Merge(m, src)
}
func (m *TabletInfo) XXX_Size() int {
	return xxx_messageInfo_TabletInfo.Size(m)
}
func (m *TabletInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TabletInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TabletInfo proto.InternalMessageInfo

func (m *TabletInfo) GetTabletId() uint64 {
	if m != nil {
		return m.TabletId
	}
	return 0
}

func (m *TabletInfo) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type GetShardLocationsResult struct {
	Tablets              []*TabletInfo `protobuf:"bytes,1,rep,name=tablets,proto3" json:"tablets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetShardLocationsResult) Reset()         { *m = GetShardLocationsResult{} }
func (m *GetShardLocationsResult) String() string { return proto.CompactTextString(m) }
func (*GetShardLocationsResult) ProtoMessage()    {}
func (*GetShardLocationsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{9}
}

func (m *GetShardLocationsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShardLocationsResult.Unmarshal(m, b)
}
func (m *GetShardLocationsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShardLocationsResult.Marshal(b, m, deterministic)
}
func (m *GetShardLocationsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShardLocationsResult.Merge(m, src)
}
func (m *GetShardLocationsResult) XXX_Size() int {
	return xxx_messageInfo_GetShardLocationsResult.Size(m)
}
func (m *GetShardLocationsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShardLocationsResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetShardLocationsResult proto.InternalMessageInfo

func (m *GetShardLocationsResult) GetTablets() []*TabletInfo {
	if m != nil {
		return m.Tablets
	}
	return nil
}

type ExecuteStreamQueryRequest struct {
	YqlText              string                     `protobuf:"bytes,1,opt,name=yql_text,json=yqlText,proto3" json:"yql_text,omitempty"`
	Parameters           map[string]*Ydb.TypedValue `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ExecuteStreamQueryRequest) Reset()         { *m = ExecuteStreamQueryRequest{} }
func (m *ExecuteStreamQueryRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteStreamQueryRequest) ProtoMessage()    {}
func (*ExecuteStreamQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{10}
}

func (m *ExecuteStreamQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStreamQueryRequest.Unmarshal(m, b)
}
func (m *ExecuteStreamQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStreamQueryRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteStreamQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStreamQueryRequest.Merge(m, src)
}
func (m *ExecuteStreamQueryRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteStreamQueryRequest.Size(m)
}
func (m *ExecuteStreamQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStreamQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStreamQueryRequest proto.InternalMessageInfo

func (m *ExecuteStreamQueryRequest) GetYqlText() string {
	if m != nil {
		return m.YqlText
	}
	return ""
}

func (m *ExecuteStreamQueryRequest) GetParameters() map[string]*Ydb.TypedValue {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type ExecuteStreamQueryResponse struct {
	Status               Ydb.StatusIds_StatusCode  `protobuf:"varint,1,opt,name=status,proto3,enum=Ydb.StatusIds_StatusCode" json:"status,omitempty"`
	Issues               []*Ydb_Issue.IssueMessage `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	Result               *ExecuteStreamQueryResult `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ExecuteStreamQueryResponse) Reset()         { *m = ExecuteStreamQueryResponse{} }
func (m *ExecuteStreamQueryResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteStreamQueryResponse) ProtoMessage()    {}
func (*ExecuteStreamQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{11}
}

func (m *ExecuteStreamQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStreamQueryResponse.Unmarshal(m, b)
}
func (m *ExecuteStreamQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStreamQueryResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteStreamQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStreamQueryResponse.Merge(m, src)
}
func (m *ExecuteStreamQueryResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteStreamQueryResponse.Size(m)
}
func (m *ExecuteStreamQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStreamQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStreamQueryResponse proto.InternalMessageInfo

func (m *ExecuteStreamQueryResponse) GetStatus() Ydb.StatusIds_StatusCode {
	if m != nil {
		return m.Status
	}
	return Ydb.StatusIds_STATUS_CODE_UNSPECIFIED
}

func (m *ExecuteStreamQueryResponse) GetIssues() []*Ydb_Issue.IssueMessage {
	if m != nil {
		return m.Issues
	}
	return nil
}

func (m *ExecuteStreamQueryResponse) GetResult() *ExecuteStreamQueryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type ExecuteStreamQueryResult struct {
	ResultSet            *Ydb.ResultSet `protobuf:"bytes,1,opt,name=result_set,json=resultSet,proto3" json:"result_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExecuteStreamQueryResult) Reset()         { *m = ExecuteStreamQueryResult{} }
func (m *ExecuteStreamQueryResult) String() string { return proto.CompactTextString(m) }
func (*ExecuteStreamQueryResult) ProtoMessage()    {}
func (*ExecuteStreamQueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{12}
}

func (m *ExecuteStreamQueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStreamQueryResult.Unmarshal(m, b)
}
func (m *ExecuteStreamQueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStreamQueryResult.Marshal(b, m, deterministic)
}
func (m *ExecuteStreamQueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStreamQueryResult.Merge(m, src)
}
func (m *ExecuteStreamQueryResult) XXX_Size() int {
	return xxx_messageInfo_ExecuteStreamQueryResult.Size(m)
}
func (m *ExecuteStreamQueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStreamQueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStreamQueryResult proto.InternalMessageInfo

func (m *ExecuteStreamQueryResult) GetResultSet() *Ydb.ResultSet {
	if m != nil {
		return m.ResultSet
	}
	return nil
}

func init() {
	proto.RegisterType((*UploadRowsRequest)(nil), "Ydb.Experimental.UploadRowsRequest")
	proto.RegisterType((*UploadRowsResponse)(nil), "Ydb.Experimental.UploadRowsResponse")
	proto.RegisterType((*UploadRowsResult)(nil), "Ydb.Experimental.UploadRowsResult")
	proto.RegisterType((*ReadColumnsRequest)(nil), "Ydb.Experimental.ReadColumnsRequest")
	proto.RegisterType((*ReadColumnsResponse)(nil), "Ydb.Experimental.ReadColumnsResponse")
	proto.RegisterType((*ReadColumnsResult)(nil), "Ydb.Experimental.ReadColumnsResult")
	proto.RegisterType((*GetShardLocationsRequest)(nil), "Ydb.Experimental.GetShardLocationsRequest")
	proto.RegisterType((*GetShardLocationsResponse)(nil), "Ydb.Experimental.GetShardLocationsResponse")
	proto.RegisterType((*TabletInfo)(nil), "Ydb.Experimental.TabletInfo")
	proto.RegisterType((*GetShardLocationsResult)(nil), "Ydb.Experimental.GetShardLocationsResult")
	proto.RegisterType((*ExecuteStreamQueryRequest)(nil), "Ydb.Experimental.ExecuteStreamQueryRequest")
	proto.RegisterMapType((map[string]*Ydb.TypedValue)(nil), "Ydb.Experimental.ExecuteStreamQueryRequest.ParametersEntry")
	proto.RegisterType((*ExecuteStreamQueryResponse)(nil), "Ydb.Experimental.ExecuteStreamQueryResponse")
	proto.RegisterType((*ExecuteStreamQueryResult)(nil), "Ydb.Experimental.ExecuteStreamQueryResult")
}

func init() { proto.RegisterFile("ydb_experimental.proto", fileDescriptor_ac21a693e2c386a5) }

var fileDescriptor_ac21a693e2c386a5 = []byte{
	// 807 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0x96, 0x93, 0xcd, 0xee, 0xfa, 0xa4, 0x4a, 0xb6, 0x03, 0xb4, 0xde, 0x14, 0xc4, 0xca, 0x08,
	0xc9, 0x42, 0xc5, 0x0b, 0x01, 0x01, 0x02, 0x71, 0x93, 0x2a, 0x42, 0x0b, 0xb4, 0xa4, 0x93, 0x80,
	0x84, 0xb8, 0xb0, 0xc6, 0xf6, 0x09, 0xb5, 0x62, 0x7b, 0x1c, 0xcf, 0xb8, 0x5d, 0x3f, 0x00, 0x3c,
	0x04, 0xcf, 0xc3, 0x13, 0xf0, 0x0a, 0xbc, 0x04, 0x97, 0x68, 0x7e, 0xbc, 0x3f, 0xed, 0x6e, 0x40,
	0x84, 0x9b, 0xd5, 0xcc, 0x39, 0xdf, 0xf9, 0xfc, 0xcd, 0x77, 0xce, 0xce, 0xc0, 0xbd, 0x36, 0x8d,
	0x23, 0x9c, 0x57, 0x58, 0x67, 0x05, 0x96, 0x92, 0xe5, 0x61, 0x55, 0x73, 0xc9, 0xc9, 0xe8, 0xc7,
	0x34, 0x0e, 0x4f, 0x57, 0xe2, 0x47, 0x1f, 0x5c, 0x65, 0x57, 0x59, 0x51, 0x4f, 0xab, 0x26, 0xce,
	0xb3, 0x64, 0xca, 0xaa, 0x6c, 0xaa, 0xa1, 0x62, 0xaa, 0x28, 0x32, 0x21, 0x1a, 0x8c, 0x0a, 0x14,
	0x82, 0xfd, 0x8c, 0x86, 0xe3, 0xe8, 0xe1, 0x8d, 0x15, 0xbc, 0xc2, 0x9a, 0xc9, 0x8c, 0x97, 0x16,
	0x3d, 0xbd, 0x11, 0x2d, 0x24, 0x93, 0x8d, 0x88, 0x12, 0x9e, 0xa2, 0xb0, 0x05, 0xc1, 0x8d, 0x05,
	0xcf, 0x59, 0xde, 0x58, 0x21, 0xfe, 0x6f, 0x0e, 0xdc, 0xfd, 0xbe, 0xca, 0x39, 0x4b, 0x29, 0x7f,
	0x21, 0x28, 0x5e, 0x37, 0x28, 0x24, 0x79, 0x1d, 0xf6, 0x24, 0x8b, 0x73, 0xf4, 0x9c, 0x89, 0x13,
	0xb8, 0xd4, 0x6c, 0xc8, 0x3b, 0xd0, 0xab, 0xf9, 0x0b, 0xe1, 0xed, 0x4c, 0x9c, 0x60, 0xff, 0xf8,
	0x30, 0x54, 0x3e, 0x5c, 0xb4, 0x15, 0xa6, 0x3f, 0x28, 0x42, 0xaa, 0x93, 0xe4, 0x6b, 0x18, 0x2d,
	0xe4, 0x47, 0x15, 0xab, 0x59, 0x21, 0xbc, 0x5d, 0x5d, 0xf0, 0xb6, 0x2e, 0xf8, 0xae, 0x4b, 0x8a,
	0xe5, 0xf2, 0x4c, 0xc3, 0xe8, 0x21, 0x5f, 0x0f, 0xf8, 0x8f, 0x81, 0xac, 0x6a, 0x13, 0x15, 0x2f,
	0x05, 0x92, 0x4f, 0xc1, 0x5d, 0x00, 0xb5, 0xc0, 0xfd, 0xe3, 0xf1, 0x56, 0x6a, 0xba, 0xc4, 0xfa,
	0x04, 0x46, 0x6b, 0x74, 0x4d, 0x2e, 0xfd, 0x3f, 0x76, 0x80, 0x50, 0x64, 0xe9, 0x23, 0x9e, 0x37,
	0x45, 0xf9, 0x0f, 0x06, 0x78, 0x30, 0x48, 0x0c, 0xce, 0xdb, 0x99, 0xec, 0x06, 0x2e, 0xed, 0xb6,
	0x64, 0x0c, 0xc3, 0xcb, 0x9a, 0x17, 0xd1, 0x15, 0xb6, 0xfa, 0xb4, 0x77, 0xe8, 0x40, 0xed, 0xbf,
	0xc1, 0x96, 0x3c, 0x04, 0xd2, 0xa5, 0xa2, 0xac, 0x4c, 0xf2, 0x46, 0x64, 0xcf, 0xd1, 0xeb, 0x4d,
	0x9c, 0x60, 0x48, 0x47, 0x16, 0x34, 0xeb, 0xe2, 0xe4, 0x0d, 0xe8, 0x4b, 0xae, 0x69, 0xf6, 0x34,
	0xcd, 0x9e, 0xe4, 0x8a, 0x24, 0x80, 0x91, 0x09, 0xaf, 0x50, 0xf4, 0x35, 0xc5, 0x81, 0x06, 0x2c,
	0x09, 0xc6, 0x30, 0x2c, 0xd8, 0x3c, 0xd2, 0x8d, 0x1a, 0x4c, 0x9c, 0xa0, 0x47, 0x07, 0x05, 0x9b,
	0xab, 0x13, 0x93, 0x07, 0xe0, 0xaa, 0x54, 0xdc, 0x4a, 0x14, 0xde, 0x50, 0xe7, 0x14, 0xf6, 0x44,
	0xed, 0x37, 0xf6, 0xcd, 0xfd, 0x8f, 0x7d, 0x7b, 0x02, 0xaf, 0xad, 0x79, 0x7a, 0xdb, 0xc6, 0xfd,
	0xea, 0xc0, 0xdd, 0x75, 0xc2, 0x26, 0x97, 0xe4, 0x1e, 0xf4, 0xe3, 0x9c, 0x27, 0x57, 0xc2, 0x73,
	0x26, 0xbb, 0xc1, 0x1d, 0x6a, 0x77, 0x64, 0x04, 0xbb, 0xc8, 0x2f, 0xf5, 0x94, 0x0e, 0xa9, 0x5a,
	0x2a, 0x4f, 0x72, 0x26, 0xe4, 0x6a, 0x77, 0xd4, 0xde, 0x76, 0xa7, 0x4b, 0xbd, 0xda, 0x1d, 0x0b,
	0x5a, 0x98, 0xeb, 0xff, 0xe2, 0x80, 0xf7, 0x15, 0xca, 0xf3, 0x67, 0xac, 0x4e, 0xbf, 0xe5, 0x89,
	0xd1, 0xdc, 0xcd, 0xcc, 0x5b, 0x00, 0x7a, 0x4c, 0x64, 0x94, 0xa5, 0x46, 0x53, 0x8f, 0xba, 0x26,
	0x32, 0x4b, 0xff, 0x5f, 0x83, 0x2f, 0x60, 0xbc, 0x41, 0xc6, 0x6d, 0x6d, 0xfe, 0x12, 0xe0, 0xc2,
	0xc8, 0x2d, 0x2f, 0xb9, 0x9a, 0x96, 0xc5, 0x71, 0x34, 0x4d, 0x8f, 0x0e, 0xbb, 0xd3, 0x10, 0x02,
	0xbd, 0x67, 0x5c, 0x48, 0x6d, 0xb2, 0x4b, 0xf5, 0xda, 0x7f, 0x0a, 0xf7, 0x37, 0x89, 0x52, 0xad,
	0xfa, 0x04, 0x06, 0xa6, 0xd4, 0xf8, 0xb2, 0x7f, 0xfc, 0x66, 0xf8, 0xf2, 0x25, 0x1a, 0x2e, 0x3f,
	0x4d, 0x3b, 0xb0, 0xff, 0xa7, 0x03, 0xe3, 0xd3, 0x39, 0x26, 0x8d, 0xc4, 0x73, 0x59, 0x23, 0x2b,
	0x9e, 0x36, 0x58, 0xb7, 0x9d, 0xe1, 0x63, 0x18, 0xb6, 0xd7, 0x79, 0x24, 0x71, 0x2e, 0xed, 0xff,
	0x74, 0xd0, 0x5e, 0xe7, 0x17, 0x38, 0x97, 0xe4, 0x27, 0x00, 0x6d, 0x31, 0x4a, 0xac, 0xcd, 0x9f,
	0x75, 0xff, 0xf8, 0x8b, 0x57, 0xbf, 0xb9, 0x95, 0x3b, 0x3c, 0x5b, 0x54, 0x9f, 0x96, 0xb2, 0x6e,
	0xe9, 0x0a, 0xdd, 0xd1, 0x13, 0x38, 0x7c, 0x29, 0xad, 0x66, 0x4e, 0x0d, 0x97, 0x51, 0xa1, 0x96,
	0xe4, 0x5d, 0xd8, 0xd3, 0xf7, 0xec, 0xb6, 0xdb, 0xd2, 0x64, 0x3f, 0xdf, 0xf9, 0xcc, 0xf1, 0x7f,
	0x77, 0xe0, 0x68, 0x93, 0x12, 0xdb, 0xcf, 0x0f, 0xa1, 0x6f, 0xae, 0x78, 0x4d, 0x7f, 0x60, 0x9b,
	0x79, 0xae, 0x43, 0xb3, 0x54, 0xd8, 0xd5, 0x23, 0x9e, 0x22, 0xb5, 0x40, 0x32, 0x85, 0xbe, 0x7e,
	0x75, 0xba, 0xa3, 0xdf, 0xd7, 0x25, 0x33, 0x15, 0x32, 0xbf, 0x8f, 0xcd, 0x6b, 0x44, 0x2d, 0x8c,
	0x9c, 0x40, 0xbf, 0xd6, 0xad, 0xb2, 0x77, 0xf5, 0x7b, 0xff, 0xce, 0x2b, 0x55, 0x41, 0x6d, 0xa5,
	0x3f, 0x03, 0x6f, 0x1b, 0x86, 0xbc, 0x0f, 0x60, 0x50, 0x91, 0x40, 0x69, 0x87, 0xf2, 0x40, 0x7f,
	0xc3, 0x00, 0xce, 0x51, 0x52, 0xb7, 0xee, 0x96, 0x27, 0x1f, 0xc3, 0x83, 0x84, 0x17, 0x61, 0xcb,
	0xca, 0x14, 0xe7, 0x61, 0x9b, 0xc6, 0xe1, 0xea, 0x3b, 0x7c, 0x42, 0x56, 0x85, 0x9d, 0xe9, 0x87,
	0xed, 0x2f, 0xc7, 0x89, 0xfb, 0xfa, 0x49, 0xfb, 0xe8, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8,
	0x4c, 0x83, 0xb7, 0xb9, 0x07, 0x00, 0x00,
}

const ()
