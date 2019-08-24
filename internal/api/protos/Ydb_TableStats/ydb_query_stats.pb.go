// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_query_stats.proto

package Ydb_TableStats

import (
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Describes select, update (insert, upsert, replace) and delete operations
type OperationStats struct {
	Rows                 uint64   `protobuf:"varint,1,opt,name=rows,proto3" json:"rows,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationStats) Reset()         { *m = OperationStats{} }
func (m *OperationStats) String() string { return proto.CompactTextString(m) }
func (*OperationStats) ProtoMessage()    {}
func (*OperationStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd6647573551cb14, []int{0}
}

func (m *OperationStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationStats.Unmarshal(m, b)
}
func (m *OperationStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationStats.Marshal(b, m, deterministic)
}
func (m *OperationStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationStats.Merge(m, src)
}
func (m *OperationStats) XXX_Size() int {
	return xxx_messageInfo_OperationStats.Size(m)
}
func (m *OperationStats) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationStats.DiscardUnknown(m)
}

var xxx_messageInfo_OperationStats proto.InternalMessageInfo

func (m *OperationStats) GetRows() uint64 {
	if m != nil {
		return m.Rows
	}
	return 0
}

func (m *OperationStats) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

// Describes all operations on a table
type TableAccessStats struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Reads                *OperationStats `protobuf:"bytes,3,opt,name=reads,proto3" json:"reads,omitempty"`
	Updates              *OperationStats `protobuf:"bytes,4,opt,name=updates,proto3" json:"updates,omitempty"`
	Deletes              *OperationStats `protobuf:"bytes,5,opt,name=deletes,proto3" json:"deletes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TableAccessStats) Reset()         { *m = TableAccessStats{} }
func (m *TableAccessStats) String() string { return proto.CompactTextString(m) }
func (*TableAccessStats) ProtoMessage()    {}
func (*TableAccessStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd6647573551cb14, []int{1}
}

func (m *TableAccessStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableAccessStats.Unmarshal(m, b)
}
func (m *TableAccessStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableAccessStats.Marshal(b, m, deterministic)
}
func (m *TableAccessStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableAccessStats.Merge(m, src)
}
func (m *TableAccessStats) XXX_Size() int {
	return xxx_messageInfo_TableAccessStats.Size(m)
}
func (m *TableAccessStats) XXX_DiscardUnknown() {
	xxx_messageInfo_TableAccessStats.DiscardUnknown(m)
}

var xxx_messageInfo_TableAccessStats proto.InternalMessageInfo

func (m *TableAccessStats) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TableAccessStats) GetReads() *OperationStats {
	if m != nil {
		return m.Reads
	}
	return nil
}

func (m *TableAccessStats) GetUpdates() *OperationStats {
	if m != nil {
		return m.Updates
	}
	return nil
}

func (m *TableAccessStats) GetDeletes() *OperationStats {
	if m != nil {
		return m.Deletes
	}
	return nil
}

type QueryPhaseStats struct {
	DurationUs           uint64              `protobuf:"varint,1,opt,name=duration_us,json=durationUs,proto3" json:"duration_us,omitempty"`
	TableAccess          []*TableAccessStats `protobuf:"bytes,2,rep,name=table_access,json=tableAccess,proto3" json:"table_access,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *QueryPhaseStats) Reset()         { *m = QueryPhaseStats{} }
func (m *QueryPhaseStats) String() string { return proto.CompactTextString(m) }
func (*QueryPhaseStats) ProtoMessage()    {}
func (*QueryPhaseStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd6647573551cb14, []int{2}
}

func (m *QueryPhaseStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryPhaseStats.Unmarshal(m, b)
}
func (m *QueryPhaseStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryPhaseStats.Marshal(b, m, deterministic)
}
func (m *QueryPhaseStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPhaseStats.Merge(m, src)
}
func (m *QueryPhaseStats) XXX_Size() int {
	return xxx_messageInfo_QueryPhaseStats.Size(m)
}
func (m *QueryPhaseStats) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPhaseStats.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPhaseStats proto.InternalMessageInfo

func (m *QueryPhaseStats) GetDurationUs() uint64 {
	if m != nil {
		return m.DurationUs
	}
	return 0
}

func (m *QueryPhaseStats) GetTableAccess() []*TableAccessStats {
	if m != nil {
		return m.TableAccess
	}
	return nil
}

type QueryStats struct {
	// A query might have one or more execution phases
	QueryPhases          []*QueryPhaseStats `protobuf:"bytes,1,rep,name=query_phases,json=queryPhases,proto3" json:"query_phases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *QueryStats) Reset()         { *m = QueryStats{} }
func (m *QueryStats) String() string { return proto.CompactTextString(m) }
func (*QueryStats) ProtoMessage()    {}
func (*QueryStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd6647573551cb14, []int{3}
}

func (m *QueryStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStats.Unmarshal(m, b)
}
func (m *QueryStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStats.Marshal(b, m, deterministic)
}
func (m *QueryStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStats.Merge(m, src)
}
func (m *QueryStats) XXX_Size() int {
	return xxx_messageInfo_QueryStats.Size(m)
}
func (m *QueryStats) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStats.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStats proto.InternalMessageInfo

func (m *QueryStats) GetQueryPhases() []*QueryPhaseStats {
	if m != nil {
		return m.QueryPhases
	}
	return nil
}

func init() {
	proto.RegisterType((*OperationStats)(nil), "Ydb.TableStats.OperationStats")
	proto.RegisterType((*TableAccessStats)(nil), "Ydb.TableStats.TableAccessStats")
	proto.RegisterType((*QueryPhaseStats)(nil), "Ydb.TableStats.QueryPhaseStats")
	proto.RegisterType((*QueryStats)(nil), "Ydb.TableStats.QueryStats")
}

func init() { proto.RegisterFile("ydb_query_stats.proto", fileDescriptor_bd6647573551cb14) }

var fileDescriptor_bd6647573551cb14 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0xe5, 0xfc, 0xf0, 0x73, 0x53, 0x85, 0xca, 0x80, 0x94, 0x89, 0x46, 0x99, 0x32, 0x79,
	0x28, 0x0c, 0x88, 0x8d, 0xb2, 0xb1, 0x50, 0x02, 0x0c, 0x4c, 0x91, 0x1d, 0x5b, 0x02, 0xa9, 0x4d,
	0xd2, 0xd8, 0x51, 0xf1, 0x8b, 0xf2, 0x2c, 0x8c, 0xc8, 0x36, 0x29, 0x6a, 0xa6, 0x6e, 0x37, 0x47,
	0xf7, 0xbb, 0xe7, 0xe8, 0xc4, 0x70, 0xa9, 0x39, 0x2b, 0x37, 0xbd, 0xe8, 0x74, 0x29, 0x15, 0x55,
	0x92, 0xb4, 0x5d, 0xa3, 0x1a, 0x1c, 0xbf, 0x73, 0x46, 0x5e, 0x29, 0x5b, 0x89, 0x17, 0xa3, 0x66,
	0x77, 0x10, 0x3f, 0xb5, 0xa2, 0xa3, 0xea, 0xb3, 0xa9, 0xad, 0x82, 0x31, 0x04, 0x5d, 0xb3, 0x95,
	0x09, 0x4a, 0x51, 0x1e, 0x14, 0x76, 0xc6, 0x17, 0x10, 0x32, 0xad, 0x84, 0x4c, 0x3c, 0x2b, 0xba,
	0x8f, 0xec, 0x1b, 0xc1, 0xd4, 0x9e, 0xba, 0xaf, 0x2a, 0x21, 0xe5, 0x0e, 0xaf, 0xe9, 0x5a, 0x58,
	0xfc, 0xb4, 0xb0, 0x33, 0xbe, 0x81, 0xb0, 0x13, 0x94, 0xcb, 0xc4, 0x4f, 0x51, 0x1e, 0xcd, 0xaf,
	0xc8, 0x7e, 0x08, 0xb2, 0x9f, 0xa0, 0x70, 0xcb, 0xf8, 0x16, 0x8e, 0xfb, 0x96, 0x53, 0x63, 0x1b,
	0x1c, 0xc4, 0x0d, 0xeb, 0x86, 0xe4, 0x62, 0x25, 0x0c, 0x19, 0x1e, 0x46, 0xfe, 0xad, 0x3f, 0x06,
	0x27, 0xde, 0xd4, 0xcf, 0xb6, 0x70, 0xf6, 0x6c, 0x9a, 0x5b, 0x7e, 0x50, 0xe9, 0x00, 0x3c, 0x83,
	0x88, 0xf7, 0x0e, 0x29, 0xfb, 0xa1, 0x1c, 0x18, 0xa4, 0x37, 0x89, 0x1f, 0x60, 0xa2, 0xcc, 0xfd,
	0x92, 0xda, 0x32, 0x12, 0x2f, 0xf5, 0xf3, 0x68, 0x9e, 0x8e, 0x8d, 0xc7, 0x7d, 0x15, 0x91, 0xfa,
	0x57, 0xb2, 0x25, 0x80, 0x35, 0x76, 0x9e, 0x0b, 0x98, 0xb8, 0x1f, 0xd8, 0x9a, 0x1c, 0xc6, 0xd4,
	0x9c, 0x9c, 0x8d, 0x4f, 0x8e, 0xa2, 0x16, 0xd1, 0x66, 0x27, 0xc8, 0xc5, 0x39, 0xc4, 0x55, 0xb3,
	0x26, 0x9a, 0xd6, 0x5c, 0x7c, 0x11, 0xcd, 0xd9, 0x0f, 0x42, 0xec, 0xc8, 0xbe, 0x85, 0xeb, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0x0c, 0xc9, 0x48, 0x24, 0x02, 0x00, 0x00,
}

const ()
