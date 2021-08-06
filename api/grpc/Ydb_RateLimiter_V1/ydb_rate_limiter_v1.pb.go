// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_rate_limiter_v1.proto

package Ydb_RateLimiter_V1

import (
	_ "github.com/yandex-cloud/ydb-go-sdk/v2/api/protos/Ydb_RateLimiter"
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

func init() { proto.RegisterFile("ydb_rate_limiter_v1.proto", fileDescriptor_707f032cadb35942) }

var fileDescriptor_707f032cadb35942 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0x11, 0xc4, 0x43, 0x50, 0x27, 0xb9, 0xb9, 0x83, 0x27, 0x37, 0xe7, 0x25, 0x65, 0xee,
	0x17, 0x6c, 0x0e, 0xbc, 0xec, 0x20, 0x15, 0x04, 0x11, 0x2c, 0x49, 0xfa, 0x90, 0xb0, 0x76, 0xc9,
	0x5e, 0xd2, 0x62, 0x7f, 0xaf, 0x7f, 0x44, 0x58, 0xdb, 0xb5, 0x4d, 0x95, 0x7a, 0xcd, 0xfb, 0xf2,
	0x7d, 0x0f, 0x12, 0x72, 0x5d, 0xc4, 0x22, 0x42, 0xee, 0x20, 0x4a, 0x54, 0xaa, 0x1c, 0x60, 0x94,
	0xcf, 0x99, 0x41, 0xed, 0x34, 0xa5, 0x6f, 0xb1, 0x60, 0x21, 0x77, 0xb0, 0x29, 0x27, 0xec, 0x75,
	0x3e, 0x0e, 0xb6, 0x6a, 0xab, 0x52, 0x0c, 0x4c, 0x26, 0x12, 0x25, 0x03, 0x6e, 0x54, 0x70, 0x80,
	0x6d, 0xe0, 0x7b, 0x4a, 0xc9, 0xc3, 0xf7, 0x29, 0xa1, 0x2d, 0xc7, 0x0b, 0x60, 0xae, 0x24, 0x50,
	0x4e, 0x2e, 0x1f, 0x11, 0xb8, 0x83, 0x10, 0xac, 0xce, 0x50, 0x02, 0x9d, 0x32, 0x3f, 0xd7, 0x05,
	0x42, 0xd8, 0x67, 0x60, 0xdd, 0xf8, 0x6e, 0x90, 0xb3, 0x46, 0xef, 0x2c, 0xd0, 0x0f, 0x72, 0xb1,
	0x4c, 0x1c, 0xe0, 0xb1, 0x30, 0xe9, 0xdd, 0xec, 0xcc, 0xeb, 0xc0, 0x74, 0x08, 0xab, 0xfc, 0xef,
	0xe4, 0x7c, 0x8d, 0xda, 0x1c, 0xf5, 0xb7, 0xbd, 0x7b, 0xed, 0x71, 0x6d, 0x9f, 0x0c, 0x50, 0xcd,
	0xf2, 0x1b, 0x65, 0x5d, 0x7d, 0x6e, 0x7f, 0x59, 0xbe, 0x33, 0xff, 0x7b, 0x79, 0x0f, 0xab, 0xfc,
	0x9f, 0xe4, 0x6a, 0x0d, 0x56, 0xa2, 0x12, 0xcd, 0x0b, 0xcc, 0xfa, 0xab, 0x79, 0x48, 0x5d, 0xb9,
	0xff, 0x07, 0x59, 0x85, 0x62, 0x32, 0x5a, 0xca, 0x7d, 0xa6, 0xb0, 0xe9, 0xf4, 0x5f, 0xd0, 0x23,
	0xea, 0xcc, 0x6c, 0x18, 0x2c, 0x2b, 0xab, 0x05, 0xb9, 0x91, 0x3a, 0x65, 0x05, 0xdf, 0xc5, 0xf0,
	0xc5, 0x8a, 0x58, 0xb0, 0xce, 0x57, 0xcc, 0xe7, 0xab, 0x51, 0x4b, 0xf3, 0x84, 0x46, 0x3e, 0x9f,
	0x88, 0xb3, 0xc3, 0x0f, 0x5d, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x02, 0x6b, 0x3b, 0xcc, 0x03,
	0x03, 0x00, 0x00,
}

const (
	CreateResource   = "/Ydb.RateLimiter.V1.RateLimiterService/CreateResource"
	AlterResource    = "/Ydb.RateLimiter.V1.RateLimiterService/AlterResource"
	DropResource     = "/Ydb.RateLimiter.V1.RateLimiterService/DropResource"
	ListResources    = "/Ydb.RateLimiter.V1.RateLimiterService/ListResources"
	DescribeResource = "/Ydb.RateLimiter.V1.RateLimiterService/DescribeResource"
	AcquireResource  = "/Ydb.RateLimiter.V1.RateLimiterService/AcquireResource"
)
