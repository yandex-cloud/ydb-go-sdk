// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_coordination_v1.proto

package Ydb_Coordination_V1

import (
	_ "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Coordination"
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

func init() { proto.RegisterFile("ydb_coordination_v1.proto", fileDescriptor_20087c3a500ad53a) }

var fileDescriptor_20087c3a500ad53a = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd2, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0x06, 0x70, 0x8a, 0xe0, 0x9f, 0xe0, 0x41, 0xb2, 0x93, 0x3b, 0x88, 0x6e, 0x2a, 0x9e, 0x12,
	0xab, 0x7e, 0x01, 0xb7, 0x81, 0x37, 0xd1, 0x4d, 0x14, 0x0f, 0x32, 0x9a, 0xe4, 0x3d, 0x84, 0x6d,
	0x79, 0x63, 0x92, 0x15, 0xfb, 0x15, 0xfd, 0x54, 0x42, 0xdb, 0x68, 0xb5, 0xa3, 0x9e, 0x9f, 0xe7,
	0xf9, 0x95, 0x26, 0x21, 0x87, 0x85, 0x12, 0x73, 0x89, 0xe8, 0x94, 0x36, 0x59, 0xd0, 0x68, 0xe6,
	0x79, 0xca, 0xac, 0xc3, 0x80, 0xb4, 0xf7, 0xaa, 0x04, 0x1b, 0x37, 0x22, 0xf6, 0x9c, 0xf6, 0xf9,
	0x42, 0x2f, 0xf4, 0xca, 0x71, 0xbb, 0x16, 0x4b, 0x2d, 0x79, 0x66, 0x35, 0x2f, 0xdb, 0x9e, 0xff,
	0x85, 0x2a, 0xe5, 0xea, 0x73, 0x8b, 0xf4, 0x9a, 0xc8, 0x0c, 0x5c, 0xae, 0x25, 0xd0, 0x29, 0xd9,
	0x99, 0x81, 0xf7, 0x1a, 0x0d, 0x3d, 0x66, 0xad, 0x2f, 0xd5, 0xd1, 0x14, 0xde, 0xd7, 0xe0, 0x43,
	0xff, 0xa4, 0xa3, 0xe1, 0x2d, 0x1a, 0x0f, 0x17, 0xc9, 0x65, 0x42, 0x5f, 0x08, 0x19, 0x3b, 0xc8,
	0x02, 0xdc, 0xa3, 0x02, 0x3a, 0x6c, 0x8f, 0x7e, 0xd2, 0x28, 0x9f, 0x76, 0x97, 0x2a, 0x9c, 0x3e,
	0x91, 0xbd, 0xdb, 0x65, 0x00, 0x57, 0xba, 0x83, 0xf6, 0xe4, 0x3b, 0x8c, 0xec, 0xb0, 0xb3, 0x53,
	0xab, 0x8f, 0x64, 0x77, 0xe2, 0xd0, 0x96, 0xe8, 0x86, 0x3f, 0x8c, 0x59, 0x34, 0x07, 0x5d, 0x95,
	0x9a, 0x7c, 0x23, 0xfb, 0x13, 0xf0, 0xd2, 0x69, 0x51, 0x9d, 0xc1, 0xd9, 0x86, 0x4d, 0x23, 0x8f,
	0xf4, 0xf9, 0x7f, 0xb5, 0x8a, 0x1f, 0xdd, 0x90, 0x23, 0x89, 0x2b, 0x56, 0x64, 0x46, 0xc1, 0x07,
	0x2b, 0x94, 0x60, 0xbf, 0x6e, 0x3c, 0x4f, 0x47, 0x07, 0x4d, 0xe4, 0xce, 0x59, 0xf9, 0x90, 0x88,
	0xed, 0xf2, 0x25, 0x5c, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xf0, 0xc7, 0x86, 0x6c, 0x02,
	0x00, 0x00,
}

const (
	Session      = "/Ydb.Coordination.V1.CoordinationService/Session"
	CreateNode   = "/Ydb.Coordination.V1.CoordinationService/CreateNode"
	AlterNode    = "/Ydb.Coordination.V1.CoordinationService/AlterNode"
	DropNode     = "/Ydb.Coordination.V1.CoordinationService/DropNode"
	DescribeNode = "/Ydb.Coordination.V1.CoordinationService/DescribeNode"
)
