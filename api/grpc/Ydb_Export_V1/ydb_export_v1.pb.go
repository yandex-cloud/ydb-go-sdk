// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_export_v1.proto

package Ydb_Export_V1

import (
	_ "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Export"
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

func init() { proto.RegisterFile("ydb_export_v1.proto", fileDescriptor_c6aac3ac0f149119) }

var fileDescriptor_c6aac3ac0f149119 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xae, 0x4c, 0x49, 0x8a,
	0x4f, 0xad, 0x28, 0xc8, 0x2f, 0x2a, 0x89, 0x2f, 0x33, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x8d, 0x4c, 0x49, 0xd2, 0x73, 0x05, 0x0b, 0xea, 0x85, 0x19, 0x4a, 0x69, 0x66, 0x67, 0x66,
	0x67, 0xe6, 0x16, 0xe9, 0x17, 0x94, 0x26, 0xe5, 0x64, 0x26, 0xeb, 0x27, 0x16, 0x64, 0xea, 0x83,
	0xd5, 0x15, 0xeb, 0x23, 0x34, 0x43, 0x74, 0x1a, 0xad, 0x64, 0xe4, 0xe2, 0x85, 0x68, 0x0c, 0x4e,
	0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xf2, 0xe6, 0xe2, 0x82, 0x08, 0x84, 0xe4, 0x47, 0x96, 0x08,
	0xc9, 0xea, 0x21, 0x19, 0x8d, 0x10, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x91, 0x92, 0xc3,
	0x25, 0x5d, 0x5c, 0x90, 0x9f, 0x57, 0x8c, 0x62, 0x58, 0xb0, 0x31, 0x76, 0xc3, 0x82, 0x8d, 0xf1,
	0x1a, 0x06, 0x92, 0x86, 0x18, 0xe6, 0x24, 0xc5, 0x25, 0x91, 0x9c, 0x9f, 0xab, 0x57, 0x99, 0x98,
	0x97, 0x92, 0x5a, 0xa1, 0x57, 0x99, 0x92, 0xa4, 0x07, 0xf5, 0x4a, 0x99, 0x61, 0x12, 0x1b, 0xd8,
	0x3b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xc8, 0x77, 0x35, 0x1f, 0x01, 0x00, 0x00,
}

const (
	ExportToYt = "/Ydb.Export.V1.ExportService/ExportToYt"
	ExportToS3 = "/Ydb.Export.V1.ExportService/ExportToS3"
)
