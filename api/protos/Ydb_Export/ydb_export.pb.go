// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_export.proto

package Ydb_Export

import (
	Ydb_Operations "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Operations"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type ExportProgress_Progress int32

const (
	ExportProgress_PROGRESS_UNSPECIFIED   ExportProgress_Progress = 0
	ExportProgress_PROGRESS_PREPARING     ExportProgress_Progress = 1
	ExportProgress_PROGRESS_TRANSFER_DATA ExportProgress_Progress = 2
	ExportProgress_PROGRESS_DONE          ExportProgress_Progress = 3
	ExportProgress_PROGRESS_CANCELLATION  ExportProgress_Progress = 4
	ExportProgress_PROGRESS_CANCELLED     ExportProgress_Progress = 5
)

var ExportProgress_Progress_name = map[int32]string{
	0: "PROGRESS_UNSPECIFIED",
	1: "PROGRESS_PREPARING",
	2: "PROGRESS_TRANSFER_DATA",
	3: "PROGRESS_DONE",
	4: "PROGRESS_CANCELLATION",
	5: "PROGRESS_CANCELLED",
}

var ExportProgress_Progress_value = map[string]int32{
	"PROGRESS_UNSPECIFIED":   0,
	"PROGRESS_PREPARING":     1,
	"PROGRESS_TRANSFER_DATA": 2,
	"PROGRESS_DONE":          3,
	"PROGRESS_CANCELLATION":  4,
	"PROGRESS_CANCELLED":     5,
}

func (x ExportProgress_Progress) String() string {
	return proto.EnumName(ExportProgress_Progress_name, int32(x))
}

func (ExportProgress_Progress) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{0, 0}
}

type ExportToS3Settings_Scheme int32

const (
	ExportToS3Settings_UNSPECIFIED ExportToS3Settings_Scheme = 0
	ExportToS3Settings_HTTP        ExportToS3Settings_Scheme = 1
	ExportToS3Settings_HTTPS       ExportToS3Settings_Scheme = 2
)

var ExportToS3Settings_Scheme_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "HTTP",
	2: "HTTPS",
}

var ExportToS3Settings_Scheme_value = map[string]int32{
	"UNSPECIFIED": 0,
	"HTTP":        1,
	"HTTPS":       2,
}

func (x ExportToS3Settings_Scheme) String() string {
	return proto.EnumName(ExportToS3Settings_Scheme_name, int32(x))
}

func (ExportToS3Settings_Scheme) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{7, 0}
}

type ExportToS3Settings_StorageClass int32

const (
	ExportToS3Settings_STORAGE_CLASS_UNSPECIFIED ExportToS3Settings_StorageClass = 0
	ExportToS3Settings_STANDARD                  ExportToS3Settings_StorageClass = 1
	ExportToS3Settings_REDUCED_REDUNDANCY        ExportToS3Settings_StorageClass = 2
	ExportToS3Settings_STANDARD_IA               ExportToS3Settings_StorageClass = 3
	ExportToS3Settings_ONEZONE_IA                ExportToS3Settings_StorageClass = 4
	ExportToS3Settings_INTELLIGENT_TIERING       ExportToS3Settings_StorageClass = 5
	ExportToS3Settings_GLACIER                   ExportToS3Settings_StorageClass = 6
	ExportToS3Settings_DEEP_ARCHIVE              ExportToS3Settings_StorageClass = 7
	ExportToS3Settings_OUTPOSTS                  ExportToS3Settings_StorageClass = 8
)

var ExportToS3Settings_StorageClass_name = map[int32]string{
	0: "STORAGE_CLASS_UNSPECIFIED",
	1: "STANDARD",
	2: "REDUCED_REDUNDANCY",
	3: "STANDARD_IA",
	4: "ONEZONE_IA",
	5: "INTELLIGENT_TIERING",
	6: "GLACIER",
	7: "DEEP_ARCHIVE",
	8: "OUTPOSTS",
}

var ExportToS3Settings_StorageClass_value = map[string]int32{
	"STORAGE_CLASS_UNSPECIFIED": 0,
	"STANDARD":                  1,
	"REDUCED_REDUNDANCY":        2,
	"STANDARD_IA":               3,
	"ONEZONE_IA":                4,
	"INTELLIGENT_TIERING":       5,
	"GLACIER":                   6,
	"DEEP_ARCHIVE":              7,
	"OUTPOSTS":                  8,
}

func (x ExportToS3Settings_StorageClass) String() string {
	return proto.EnumName(ExportToS3Settings_StorageClass_name, int32(x))
}

func (ExportToS3Settings_StorageClass) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{7, 1}
}

/// Common
type ExportProgress struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportProgress) Reset()         { *m = ExportProgress{} }
func (m *ExportProgress) String() string { return proto.CompactTextString(m) }
func (*ExportProgress) ProtoMessage()    {}
func (*ExportProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{0}
}

func (m *ExportProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportProgress.Unmarshal(m, b)
}
func (m *ExportProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportProgress.Marshal(b, m, deterministic)
}
func (m *ExportProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportProgress.Merge(m, src)
}
func (m *ExportProgress) XXX_Size() int {
	return xxx_messageInfo_ExportProgress.Size(m)
}
func (m *ExportProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportProgress.DiscardUnknown(m)
}

var xxx_messageInfo_ExportProgress proto.InternalMessageInfo

type ExportItemProgress struct {
	PartsTotal           uint32                 `protobuf:"varint,1,opt,name=parts_total,json=partsTotal,proto3" json:"parts_total,omitempty"`
	PartsCompleted       uint32                 `protobuf:"varint,2,opt,name=parts_completed,json=partsCompleted,proto3" json:"parts_completed,omitempty"`
	StartTime            *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ExportItemProgress) Reset()         { *m = ExportItemProgress{} }
func (m *ExportItemProgress) String() string { return proto.CompactTextString(m) }
func (*ExportItemProgress) ProtoMessage()    {}
func (*ExportItemProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{1}
}

func (m *ExportItemProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportItemProgress.Unmarshal(m, b)
}
func (m *ExportItemProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportItemProgress.Marshal(b, m, deterministic)
}
func (m *ExportItemProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportItemProgress.Merge(m, src)
}
func (m *ExportItemProgress) XXX_Size() int {
	return xxx_messageInfo_ExportItemProgress.Size(m)
}
func (m *ExportItemProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportItemProgress.DiscardUnknown(m)
}

var xxx_messageInfo_ExportItemProgress proto.InternalMessageInfo

func (m *ExportItemProgress) GetPartsTotal() uint32 {
	if m != nil {
		return m.PartsTotal
	}
	return 0
}

func (m *ExportItemProgress) GetPartsCompleted() uint32 {
	if m != nil {
		return m.PartsCompleted
	}
	return 0
}

func (m *ExportItemProgress) GetStartTime() *timestamppb.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ExportItemProgress) GetEndTime() *timestamppb.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

/// YT
type ExportToYtSettings struct {
	Host                 string                     `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 uint32                     `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Token                string                     `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Items                []*ExportToYtSettings_Item `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	Description          string                     `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	NumberOfRetries      uint32                     `protobuf:"varint,6,opt,name=number_of_retries,json=numberOfRetries,proto3" json:"number_of_retries,omitempty"`
	UseTypeV3            bool                       `protobuf:"varint,7,opt,name=use_type_v3,json=useTypeV3,proto3" json:"use_type_v3,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ExportToYtSettings) Reset()         { *m = ExportToYtSettings{} }
func (m *ExportToYtSettings) String() string { return proto.CompactTextString(m) }
func (*ExportToYtSettings) ProtoMessage()    {}
func (*ExportToYtSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{2}
}

func (m *ExportToYtSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToYtSettings.Unmarshal(m, b)
}
func (m *ExportToYtSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToYtSettings.Marshal(b, m, deterministic)
}
func (m *ExportToYtSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToYtSettings.Merge(m, src)
}
func (m *ExportToYtSettings) XXX_Size() int {
	return xxx_messageInfo_ExportToYtSettings.Size(m)
}
func (m *ExportToYtSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToYtSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToYtSettings proto.InternalMessageInfo

func (m *ExportToYtSettings) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ExportToYtSettings) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ExportToYtSettings) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ExportToYtSettings) GetItems() []*ExportToYtSettings_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ExportToYtSettings) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ExportToYtSettings) GetNumberOfRetries() uint32 {
	if m != nil {
		return m.NumberOfRetries
	}
	return 0
}

func (m *ExportToYtSettings) GetUseTypeV3() bool {
	if m != nil {
		return m.UseTypeV3
	}
	return false
}

type ExportToYtSettings_Item struct {
	// Database path to a table to be exported
	SourcePath           string   `protobuf:"bytes,1,opt,name=source_path,json=sourcePath,proto3" json:"source_path,omitempty"`
	DestinationPath      string   `protobuf:"bytes,2,opt,name=destination_path,json=destinationPath,proto3" json:"destination_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportToYtSettings_Item) Reset()         { *m = ExportToYtSettings_Item{} }
func (m *ExportToYtSettings_Item) String() string { return proto.CompactTextString(m) }
func (*ExportToYtSettings_Item) ProtoMessage()    {}
func (*ExportToYtSettings_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{2, 0}
}

func (m *ExportToYtSettings_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToYtSettings_Item.Unmarshal(m, b)
}
func (m *ExportToYtSettings_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToYtSettings_Item.Marshal(b, m, deterministic)
}
func (m *ExportToYtSettings_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToYtSettings_Item.Merge(m, src)
}
func (m *ExportToYtSettings_Item) XXX_Size() int {
	return xxx_messageInfo_ExportToYtSettings_Item.Size(m)
}
func (m *ExportToYtSettings_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToYtSettings_Item.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToYtSettings_Item proto.InternalMessageInfo

func (m *ExportToYtSettings_Item) GetSourcePath() string {
	if m != nil {
		return m.SourcePath
	}
	return ""
}

func (m *ExportToYtSettings_Item) GetDestinationPath() string {
	if m != nil {
		return m.DestinationPath
	}
	return ""
}

type ExportToYtResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportToYtResult) Reset()         { *m = ExportToYtResult{} }
func (m *ExportToYtResult) String() string { return proto.CompactTextString(m) }
func (*ExportToYtResult) ProtoMessage()    {}
func (*ExportToYtResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{3}
}

func (m *ExportToYtResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToYtResult.Unmarshal(m, b)
}
func (m *ExportToYtResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToYtResult.Marshal(b, m, deterministic)
}
func (m *ExportToYtResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToYtResult.Merge(m, src)
}
func (m *ExportToYtResult) XXX_Size() int {
	return xxx_messageInfo_ExportToYtResult.Size(m)
}
func (m *ExportToYtResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToYtResult.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToYtResult proto.InternalMessageInfo

type ExportToYtMetadata struct {
	Settings             *ExportToYtSettings     `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	Progress             ExportProgress_Progress `protobuf:"varint,2,opt,name=progress,proto3,enum=Ydb.Export.ExportProgress_Progress" json:"progress,omitempty"`
	ItemsProgress        []*ExportItemProgress   `protobuf:"bytes,3,rep,name=items_progress,json=itemsProgress,proto3" json:"items_progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ExportToYtMetadata) Reset()         { *m = ExportToYtMetadata{} }
func (m *ExportToYtMetadata) String() string { return proto.CompactTextString(m) }
func (*ExportToYtMetadata) ProtoMessage()    {}
func (*ExportToYtMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{4}
}

func (m *ExportToYtMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToYtMetadata.Unmarshal(m, b)
}
func (m *ExportToYtMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToYtMetadata.Marshal(b, m, deterministic)
}
func (m *ExportToYtMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToYtMetadata.Merge(m, src)
}
func (m *ExportToYtMetadata) XXX_Size() int {
	return xxx_messageInfo_ExportToYtMetadata.Size(m)
}
func (m *ExportToYtMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToYtMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToYtMetadata proto.InternalMessageInfo

func (m *ExportToYtMetadata) GetSettings() *ExportToYtSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *ExportToYtMetadata) GetProgress() ExportProgress_Progress {
	if m != nil {
		return m.Progress
	}
	return ExportProgress_PROGRESS_UNSPECIFIED
}

func (m *ExportToYtMetadata) GetItemsProgress() []*ExportItemProgress {
	if m != nil {
		return m.ItemsProgress
	}
	return nil
}

type ExportToYtRequest struct {
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,1,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	Settings             *ExportToYtSettings             `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ExportToYtRequest) Reset()         { *m = ExportToYtRequest{} }
func (m *ExportToYtRequest) String() string { return proto.CompactTextString(m) }
func (*ExportToYtRequest) ProtoMessage()    {}
func (*ExportToYtRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{5}
}

func (m *ExportToYtRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToYtRequest.Unmarshal(m, b)
}
func (m *ExportToYtRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToYtRequest.Marshal(b, m, deterministic)
}
func (m *ExportToYtRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToYtRequest.Merge(m, src)
}
func (m *ExportToYtRequest) XXX_Size() int {
	return xxx_messageInfo_ExportToYtRequest.Size(m)
}
func (m *ExportToYtRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToYtRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToYtRequest proto.InternalMessageInfo

func (m *ExportToYtRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

func (m *ExportToYtRequest) GetSettings() *ExportToYtSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

type ExportToYtResponse struct {
	// operation.result = ExportToYtResult
	// operation.metadata = ExportToYtMetadata
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ExportToYtResponse) Reset()         { *m = ExportToYtResponse{} }
func (m *ExportToYtResponse) String() string { return proto.CompactTextString(m) }
func (*ExportToYtResponse) ProtoMessage()    {}
func (*ExportToYtResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{6}
}

func (m *ExportToYtResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToYtResponse.Unmarshal(m, b)
}
func (m *ExportToYtResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToYtResponse.Marshal(b, m, deterministic)
}
func (m *ExportToYtResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToYtResponse.Merge(m, src)
}
func (m *ExportToYtResponse) XXX_Size() int {
	return xxx_messageInfo_ExportToYtResponse.Size(m)
}
func (m *ExportToYtResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToYtResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToYtResponse proto.InternalMessageInfo

func (m *ExportToYtResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

/// S3
type ExportToS3Settings struct {
	Endpoint             string                          `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Scheme               ExportToS3Settings_Scheme       `protobuf:"varint,2,opt,name=scheme,proto3,enum=Ydb.Export.ExportToS3Settings_Scheme" json:"scheme,omitempty"`
	Bucket               string                          `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	AccessKey            string                          `protobuf:"bytes,4,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey            string                          `protobuf:"bytes,5,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Items                []*ExportToS3Settings_Item      `protobuf:"bytes,6,rep,name=items,proto3" json:"items,omitempty"`
	Description          string                          `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	NumberOfRetries      uint32                          `protobuf:"varint,8,opt,name=number_of_retries,json=numberOfRetries,proto3" json:"number_of_retries,omitempty"`
	StorageClass         ExportToS3Settings_StorageClass `protobuf:"varint,9,opt,name=storage_class,json=storageClass,proto3,enum=Ydb.Export.ExportToS3Settings_StorageClass" json:"storage_class,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ExportToS3Settings) Reset()         { *m = ExportToS3Settings{} }
func (m *ExportToS3Settings) String() string { return proto.CompactTextString(m) }
func (*ExportToS3Settings) ProtoMessage()    {}
func (*ExportToS3Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{7}
}

func (m *ExportToS3Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToS3Settings.Unmarshal(m, b)
}
func (m *ExportToS3Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToS3Settings.Marshal(b, m, deterministic)
}
func (m *ExportToS3Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToS3Settings.Merge(m, src)
}
func (m *ExportToS3Settings) XXX_Size() int {
	return xxx_messageInfo_ExportToS3Settings.Size(m)
}
func (m *ExportToS3Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToS3Settings.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToS3Settings proto.InternalMessageInfo

func (m *ExportToS3Settings) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *ExportToS3Settings) GetScheme() ExportToS3Settings_Scheme {
	if m != nil {
		return m.Scheme
	}
	return ExportToS3Settings_UNSPECIFIED
}

func (m *ExportToS3Settings) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *ExportToS3Settings) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *ExportToS3Settings) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *ExportToS3Settings) GetItems() []*ExportToS3Settings_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ExportToS3Settings) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ExportToS3Settings) GetNumberOfRetries() uint32 {
	if m != nil {
		return m.NumberOfRetries
	}
	return 0
}

func (m *ExportToS3Settings) GetStorageClass() ExportToS3Settings_StorageClass {
	if m != nil {
		return m.StorageClass
	}
	return ExportToS3Settings_STORAGE_CLASS_UNSPECIFIED
}

type ExportToS3Settings_Item struct {
	// Database path to a table to be exported
	SourcePath string `protobuf:"bytes,1,opt,name=source_path,json=sourcePath,proto3" json:"source_path,omitempty"`
	// Tables are exported to one or more S3 objects.
	//The object name begins with 'destination_prefix'.
	//This prefix will be followed by '/data_PartNumber', where 'PartNumber'
	//represents the index of the part, starting at zero.
	DestinationPrefix    string   `protobuf:"bytes,2,opt,name=destination_prefix,json=destinationPrefix,proto3" json:"destination_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportToS3Settings_Item) Reset()         { *m = ExportToS3Settings_Item{} }
func (m *ExportToS3Settings_Item) String() string { return proto.CompactTextString(m) }
func (*ExportToS3Settings_Item) ProtoMessage()    {}
func (*ExportToS3Settings_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{7, 0}
}

func (m *ExportToS3Settings_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToS3Settings_Item.Unmarshal(m, b)
}
func (m *ExportToS3Settings_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToS3Settings_Item.Marshal(b, m, deterministic)
}
func (m *ExportToS3Settings_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToS3Settings_Item.Merge(m, src)
}
func (m *ExportToS3Settings_Item) XXX_Size() int {
	return xxx_messageInfo_ExportToS3Settings_Item.Size(m)
}
func (m *ExportToS3Settings_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToS3Settings_Item.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToS3Settings_Item proto.InternalMessageInfo

func (m *ExportToS3Settings_Item) GetSourcePath() string {
	if m != nil {
		return m.SourcePath
	}
	return ""
}

func (m *ExportToS3Settings_Item) GetDestinationPrefix() string {
	if m != nil {
		return m.DestinationPrefix
	}
	return ""
}

type ExportToS3Result struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportToS3Result) Reset()         { *m = ExportToS3Result{} }
func (m *ExportToS3Result) String() string { return proto.CompactTextString(m) }
func (*ExportToS3Result) ProtoMessage()    {}
func (*ExportToS3Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{8}
}

func (m *ExportToS3Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToS3Result.Unmarshal(m, b)
}
func (m *ExportToS3Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToS3Result.Marshal(b, m, deterministic)
}
func (m *ExportToS3Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToS3Result.Merge(m, src)
}
func (m *ExportToS3Result) XXX_Size() int {
	return xxx_messageInfo_ExportToS3Result.Size(m)
}
func (m *ExportToS3Result) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToS3Result.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToS3Result proto.InternalMessageInfo

type ExportToS3Metadata struct {
	Settings             *ExportToS3Settings     `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	Progress             ExportProgress_Progress `protobuf:"varint,2,opt,name=progress,proto3,enum=Ydb.Export.ExportProgress_Progress" json:"progress,omitempty"`
	ItemsProgress        []*ExportItemProgress   `protobuf:"bytes,3,rep,name=items_progress,json=itemsProgress,proto3" json:"items_progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ExportToS3Metadata) Reset()         { *m = ExportToS3Metadata{} }
func (m *ExportToS3Metadata) String() string { return proto.CompactTextString(m) }
func (*ExportToS3Metadata) ProtoMessage()    {}
func (*ExportToS3Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{9}
}

func (m *ExportToS3Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToS3Metadata.Unmarshal(m, b)
}
func (m *ExportToS3Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToS3Metadata.Marshal(b, m, deterministic)
}
func (m *ExportToS3Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToS3Metadata.Merge(m, src)
}
func (m *ExportToS3Metadata) XXX_Size() int {
	return xxx_messageInfo_ExportToS3Metadata.Size(m)
}
func (m *ExportToS3Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToS3Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToS3Metadata proto.InternalMessageInfo

func (m *ExportToS3Metadata) GetSettings() *ExportToS3Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *ExportToS3Metadata) GetProgress() ExportProgress_Progress {
	if m != nil {
		return m.Progress
	}
	return ExportProgress_PROGRESS_UNSPECIFIED
}

func (m *ExportToS3Metadata) GetItemsProgress() []*ExportItemProgress {
	if m != nil {
		return m.ItemsProgress
	}
	return nil
}

type ExportToS3Request struct {
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,1,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	Settings             *ExportToS3Settings             `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ExportToS3Request) Reset()         { *m = ExportToS3Request{} }
func (m *ExportToS3Request) String() string { return proto.CompactTextString(m) }
func (*ExportToS3Request) ProtoMessage()    {}
func (*ExportToS3Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{10}
}

func (m *ExportToS3Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToS3Request.Unmarshal(m, b)
}
func (m *ExportToS3Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToS3Request.Marshal(b, m, deterministic)
}
func (m *ExportToS3Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToS3Request.Merge(m, src)
}
func (m *ExportToS3Request) XXX_Size() int {
	return xxx_messageInfo_ExportToS3Request.Size(m)
}
func (m *ExportToS3Request) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToS3Request.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToS3Request proto.InternalMessageInfo

func (m *ExportToS3Request) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

func (m *ExportToS3Request) GetSettings() *ExportToS3Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

type ExportToS3Response struct {
	// operation.result = ExportToS3Result
	// operation.metadata = ExportToS3Metadata
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ExportToS3Response) Reset()         { *m = ExportToS3Response{} }
func (m *ExportToS3Response) String() string { return proto.CompactTextString(m) }
func (*ExportToS3Response) ProtoMessage()    {}
func (*ExportToS3Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_113d267532c68f6f, []int{11}
}

func (m *ExportToS3Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToS3Response.Unmarshal(m, b)
}
func (m *ExportToS3Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToS3Response.Marshal(b, m, deterministic)
}
func (m *ExportToS3Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToS3Response.Merge(m, src)
}
func (m *ExportToS3Response) XXX_Size() int {
	return xxx_messageInfo_ExportToS3Response.Size(m)
}
func (m *ExportToS3Response) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToS3Response.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToS3Response proto.InternalMessageInfo

func (m *ExportToS3Response) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func init() {
	proto.RegisterEnum("Ydb.Export.ExportProgress_Progress", ExportProgress_Progress_name, ExportProgress_Progress_value)
	proto.RegisterEnum("Ydb.Export.ExportToS3Settings_Scheme", ExportToS3Settings_Scheme_name, ExportToS3Settings_Scheme_value)
	proto.RegisterEnum("Ydb.Export.ExportToS3Settings_StorageClass", ExportToS3Settings_StorageClass_name, ExportToS3Settings_StorageClass_value)
	proto.RegisterType((*ExportProgress)(nil), "Ydb.Export.ExportProgress")
	proto.RegisterType((*ExportItemProgress)(nil), "Ydb.Export.ExportItemProgress")
	proto.RegisterType((*ExportToYtSettings)(nil), "Ydb.Export.ExportToYtSettings")
	proto.RegisterType((*ExportToYtSettings_Item)(nil), "Ydb.Export.ExportToYtSettings.Item")
	proto.RegisterType((*ExportToYtResult)(nil), "Ydb.Export.ExportToYtResult")
	proto.RegisterType((*ExportToYtMetadata)(nil), "Ydb.Export.ExportToYtMetadata")
	proto.RegisterType((*ExportToYtRequest)(nil), "Ydb.Export.ExportToYtRequest")
	proto.RegisterType((*ExportToYtResponse)(nil), "Ydb.Export.ExportToYtResponse")
	proto.RegisterType((*ExportToS3Settings)(nil), "Ydb.Export.ExportToS3Settings")
	proto.RegisterType((*ExportToS3Settings_Item)(nil), "Ydb.Export.ExportToS3Settings.Item")
	proto.RegisterType((*ExportToS3Result)(nil), "Ydb.Export.ExportToS3Result")
	proto.RegisterType((*ExportToS3Metadata)(nil), "Ydb.Export.ExportToS3Metadata")
	proto.RegisterType((*ExportToS3Request)(nil), "Ydb.Export.ExportToS3Request")
	proto.RegisterType((*ExportToS3Response)(nil), "Ydb.Export.ExportToS3Response")
}

func init() { proto.RegisterFile("ydb_export.proto", fileDescriptor_113d267532c68f6f) }

var fileDescriptor_113d267532c68f6f = []byte{
	// 1095 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0xf9, 0xcf, 0x49, 0x7f, 0xdc, 0x81, 0x5d, 0xd2, 0x08, 0xb6, 0x55, 0x56, 0x2b, 0x4a,
	0x41, 0x8e, 0xd4, 0x80, 0x10, 0x48, 0x08, 0xdc, 0x64, 0xb6, 0x6b, 0xc8, 0x3a, 0xd6, 0xd8, 0x5d,
	0xa9, 0x5c, 0x60, 0x39, 0xf6, 0xb4, 0xb5, 0x9a, 0xc4, 0xc6, 0x33, 0x59, 0x35, 0x77, 0x3c, 0x02,
	0xd7, 0x7b, 0x81, 0x04, 0x6f, 0xc0, 0x2d, 0x4f, 0xc1, 0x25, 0x4f, 0xd0, 0x67, 0xd8, 0x4b, 0xe4,
	0x71, 0xec, 0x38, 0xda, 0x74, 0x77, 0x2b, 0x2e, 0xe0, 0x2a, 0x33, 0xdf, 0xf9, 0xce, 0x9c, 0x33,
	0xdf, 0xf1, 0x39, 0x13, 0x90, 0xe7, 0xde, 0xc8, 0xa6, 0xd7, 0x61, 0x10, 0x71, 0x25, 0x8c, 0x02,
	0x1e, 0x20, 0x38, 0xf3, 0x46, 0x0a, 0x16, 0x48, 0xeb, 0xb3, 0x2b, 0xff, 0xca, 0x9f, 0x44, 0x9d,
	0x70, 0x36, 0x1a, 0xfb, 0x6e, 0xc7, 0x09, 0xfd, 0x8e, 0x20, 0xb1, 0xce, 0x73, 0x67, 0xec, 0x7b,
	0x0e, 0xf7, 0x83, 0x69, 0x6e, 0x99, 0x9c, 0xd0, 0xfa, 0xf4, 0x56, 0xaf, 0x38, 0x58, 0x10, 0xd2,
	0x28, 0xcf, 0xde, 0xbb, 0x08, 0x82, 0x8b, 0x31, 0x4d, 0x28, 0xa3, 0xd9, 0x79, 0x87, 0xfb, 0x13,
	0xca, 0xb8, 0x33, 0x09, 0x13, 0x42, 0xfb, 0x0f, 0x09, 0xb6, 0x92, 0x7c, 0x8c, 0x28, 0xb8, 0x88,
	0x28, 0x63, 0xed, 0x5f, 0x25, 0xa8, 0xa5, 0x1b, 0xd4, 0x84, 0xf7, 0x0c, 0x32, 0x3c, 0x21, 0xd8,
	0x34, 0xed, 0x53, 0xdd, 0x34, 0x70, 0x4f, 0x7b, 0xac, 0xe1, 0xbe, 0xfc, 0x0e, 0xba, 0x0f, 0x28,
	0xb3, 0x18, 0x04, 0x1b, 0x2a, 0xd1, 0xf4, 0x13, 0x59, 0x42, 0x2d, 0xb8, 0x9f, 0xe1, 0x16, 0x51,
	0x75, 0xf3, 0x31, 0x26, 0x76, 0x5f, 0xb5, 0x54, 0xb9, 0x80, 0x76, 0x60, 0x33, 0xb3, 0xf5, 0x87,
	0x3a, 0x96, 0x8b, 0x68, 0x17, 0xee, 0x65, 0x50, 0x4f, 0xd5, 0x7b, 0x78, 0x30, 0x50, 0x2d, 0x6d,
	0xa8, 0xcb, 0xa5, 0x95, 0x08, 0x0b, 0x13, 0xee, 0xcb, 0xe5, 0xf6, 0x5f, 0x12, 0xa0, 0x24, 0x67,
	0x8d, 0xd3, 0x49, 0x96, 0xea, 0x1e, 0x34, 0x42, 0x27, 0xe2, 0xcc, 0xe6, 0x01, 0x77, 0xc6, 0x4d,
	0x69, 0x5f, 0x3a, 0xd8, 0x24, 0x20, 0x20, 0x2b, 0x46, 0xd0, 0x47, 0xb0, 0x9d, 0x10, 0xdc, 0x60,
	0x12, 0x8e, 0x29, 0xa7, 0x5e, 0xb3, 0x20, 0x48, 0x5b, 0x02, 0xee, 0xa5, 0x28, 0xfa, 0x12, 0x80,
	0x71, 0x27, 0xe2, 0x76, 0xac, 0x56, 0xb3, 0xb8, 0x2f, 0x1d, 0x34, 0x8e, 0x5a, 0x4a, 0x22, 0xa5,
	0x92, 0x4a, 0xa9, 0x58, 0xa9, 0x94, 0xa4, 0x2e, 0xd8, 0xf1, 0x1e, 0x7d, 0x0e, 0x35, 0x3a, 0xf5,
	0x12, 0xc7, 0xd2, 0x1b, 0x1d, 0xab, 0x74, 0xea, 0xc5, 0xbb, 0xf6, 0xcb, 0x42, 0x7a, 0x25, 0x2b,
	0x38, 0xe3, 0x26, 0xe5, 0xdc, 0x9f, 0x5e, 0xc4, 0xea, 0x97, 0x2e, 0x03, 0xc6, 0xc5, 0x5d, 0xea,
	0xc7, 0xa5, 0x5f, 0x6e, 0x0e, 0x25, 0x22, 0x10, 0x84, 0xa0, 0x14, 0xb3, 0x17, 0x17, 0x10, 0x6b,
	0xd4, 0x82, 0x32, 0x0f, 0xae, 0xe8, 0x54, 0x64, 0x9c, 0xd2, 0x13, 0x08, 0xa9, 0x50, 0xf6, 0x39,
	0x9d, 0xb0, 0x66, 0x69, 0xbf, 0x78, 0xd0, 0x38, 0x7a, 0xa8, 0x2c, 0x3f, 0x44, 0xe5, 0xd5, 0xc0,
	0x4a, 0x2c, 0xec, 0x71, 0xe5, 0xc5, 0xcd, 0x61, 0xe1, 0x40, 0x22, 0x89, 0x27, 0xfa, 0x18, 0x1a,
	0x1e, 0x65, 0x6e, 0xe4, 0x87, 0xf1, 0x07, 0xd6, 0x2c, 0x8b, 0x20, 0xd5, 0xdf, 0x6f, 0x0e, 0x8b,
	0xcd, 0x9f, 0x25, 0x92, 0xb7, 0xa1, 0x43, 0xd8, 0x99, 0xce, 0x26, 0x23, 0x1a, 0xd9, 0xc1, 0xb9,
	0x1d, 0x51, 0x1e, 0xf9, 0x94, 0x35, 0x2b, 0x22, 0xd5, 0xed, 0xc4, 0x30, 0x3c, 0x27, 0x09, 0x8c,
	0x1e, 0x40, 0x63, 0xc6, 0xa8, 0xcd, 0xe7, 0x21, 0xb5, 0x9f, 0x77, 0x9b, 0xd5, 0x7d, 0xe9, 0xa0,
	0x46, 0xea, 0x33, 0x46, 0xad, 0x79, 0x48, 0x9f, 0x75, 0x5b, 0x3f, 0x42, 0x29, 0xce, 0x06, 0x3d,
	0x82, 0x06, 0x0b, 0x66, 0x91, 0x4b, 0xed, 0xd0, 0xe1, 0x97, 0x2b, 0x92, 0x40, 0x62, 0x30, 0x1c,
	0x7e, 0x89, 0x3a, 0x20, 0x7b, 0x94, 0x71, 0x7f, 0x2a, 0xda, 0x20, 0xe1, 0x16, 0x72, 0xdc, 0xed,
	0x9c, 0x35, 0x76, 0x68, 0x23, 0x90, 0x97, 0x02, 0x10, 0xca, 0x66, 0x63, 0xde, 0xfe, 0x5b, 0xca,
	0x97, 0xe3, 0x29, 0xe5, 0x8e, 0xe7, 0x70, 0x07, 0x7d, 0x05, 0x35, 0xb6, 0x50, 0x48, 0xc4, 0x6f,
	0x1c, 0x3d, 0x78, 0xbd, 0x8e, 0x24, 0xe3, 0xa3, 0x6f, 0xa0, 0x16, 0x2e, 0xbe, 0x54, 0x91, 0xcf,
	0xd6, 0xba, 0x1a, 0xa4, 0xdf, 0xb2, 0x92, 0x2e, 0x48, 0xe6, 0x84, 0x30, 0x6c, 0x89, 0x3a, 0xd8,
	0xd9, 0x31, 0x45, 0x51, 0xca, 0x35, 0x29, 0xe4, 0xdb, 0x82, 0x6c, 0x0a, 0xaf, 0xac, 0xbb, 0x7f,
	0x93, 0x60, 0x27, 0x7f, 0xdf, 0x9f, 0x66, 0x94, 0x71, 0xf4, 0x1d, 0xc8, 0xd9, 0xe8, 0xb0, 0x43,
	0x27, 0x72, 0x26, 0xe9, 0x0d, 0xf7, 0xc4, 0xf1, 0xc3, 0xd4, 0xc8, 0x96, 0x4b, 0x43, 0xd0, 0xc8,
	0x76, 0xb0, 0x0a, 0xa0, 0x6f, 0x73, 0x2a, 0x15, 0xde, 0x46, 0xa5, 0x45, 0x65, 0x32, 0xaf, 0xf6,
	0xd3, 0xbc, 0xfa, 0x84, 0xb2, 0x30, 0x98, 0x32, 0x8a, 0xbe, 0x80, 0x7a, 0x16, 0x6a, 0x91, 0xdc,
	0xee, 0xad, 0xc9, 0x91, 0x25, 0xb7, 0xfd, 0xa2, 0xb2, 0x3c, 0xcf, 0xec, 0x66, 0xcd, 0xb5, 0x2f,
	0x5a, 0x35, 0x0c, 0xfc, 0xe9, 0x6a, 0x83, 0x65, 0x28, 0xfa, 0x1a, 0x2a, 0xcc, 0xbd, 0xa4, 0x13,
	0xba, 0xa8, 0xd8, 0xa3, 0x75, 0xf7, 0x58, 0x9e, 0xa8, 0x98, 0x82, 0x4c, 0x16, 0x4e, 0xe8, 0x03,
	0xa8, 0x8c, 0x66, 0xee, 0x15, 0xe5, 0x2b, 0x0d, 0xb9, 0xc0, 0xd0, 0x43, 0x00, 0xc7, 0x75, 0x29,
	0x63, 0xf6, 0x15, 0x9d, 0x8b, 0x59, 0x91, 0x32, 0xea, 0x09, 0xfe, 0x3d, 0x9d, 0xc7, 0x24, 0x46,
	0xdd, 0x88, 0x72, 0x41, 0x2a, 0xe7, 0x49, 0x09, 0x1e, 0x93, 0xb2, 0xde, 0xae, 0xdc, 0xde, 0xdb,
	0xb9, 0x2c, 0xdf, 0xa2, 0xb7, 0xab, 0x77, 0xed, 0xed, 0xda, 0xfa, 0xde, 0x36, 0x60, 0x93, 0xf1,
	0x20, 0x72, 0x2e, 0xa8, 0xed, 0x8e, 0x1d, 0xc6, 0x9a, 0x75, 0xa1, 0xe3, 0x27, 0x6f, 0xd2, 0x31,
	0xf1, 0xe9, 0xc5, 0x2e, 0x64, 0x83, 0xe5, 0x76, 0xad, 0xd1, 0xdd, 0xa6, 0x41, 0x17, 0xd0, 0xca,
	0x34, 0x88, 0xe8, 0xb9, 0x7f, 0xbd, 0x32, 0x0f, 0x76, 0xf2, 0xf3, 0x40, 0x98, 0xdb, 0x0a, 0x54,
	0x92, 0x4a, 0xa2, 0x6d, 0x68, 0xac, 0x3e, 0x7a, 0x35, 0x28, 0x3d, 0xb1, 0x2c, 0x43, 0x96, 0x50,
	0x1d, 0xca, 0xf1, 0xca, 0x94, 0x0b, 0xed, 0x3f, 0x25, 0xd8, 0xc8, 0xa7, 0x8c, 0x3e, 0x84, 0x5d,
	0xd3, 0x1a, 0x12, 0xf5, 0x04, 0xdb, 0xbd, 0x81, 0xfa, 0xca, 0xcb, 0xb9, 0x01, 0x35, 0xd3, 0x52,
	0xf5, 0xbe, 0x4a, 0xfa, 0xb2, 0x14, 0xbf, 0x72, 0x04, 0xf7, 0x4f, 0x7b, 0xb8, 0x6f, 0xc7, 0xbf,
	0x7a, 0x5f, 0xd5, 0x7b, 0x67, 0x72, 0x21, 0x8e, 0x9d, 0xb2, 0x6c, 0x4d, 0x95, 0x8b, 0x68, 0x0b,
	0x60, 0xa8, 0xe3, 0x1f, 0x86, 0x3a, 0x8e, 0xf7, 0x25, 0xf4, 0x3e, 0xbc, 0xab, 0xe9, 0x16, 0x1e,
	0x0c, 0xb4, 0x13, 0xac, 0x5b, 0xb6, 0xa5, 0x61, 0xf1, 0x02, 0x97, 0x51, 0x03, 0xaa, 0x27, 0x03,
	0xb5, 0xa7, 0x61, 0x22, 0x57, 0x90, 0x0c, 0x1b, 0x7d, 0x8c, 0x0d, 0x5b, 0x25, 0xbd, 0x27, 0xda,
	0x33, 0x2c, 0x57, 0xe3, 0xf0, 0xc3, 0x53, 0xcb, 0x18, 0x9a, 0x96, 0x29, 0xd7, 0xf2, 0xe3, 0xcf,
	0xec, 0xae, 0x19, 0x7f, 0x66, 0xf7, 0xae, 0xe3, 0x6f, 0x59, 0xc8, 0xff, 0xf9, 0xf8, 0x8b, 0xef,
	0xfb, 0x9f, 0x8d, 0xbf, 0xa5, 0x4a, 0xaf, 0x1b, 0x7f, 0xa2, 0x24, 0xff, 0x6e, 0xfc, 0x1d, 0xb7,
	0xe0, 0x9e, 0x1b, 0x4c, 0x94, 0xb9, 0x33, 0xf5, 0xe8, 0xb5, 0x32, 0xf7, 0x46, 0x4a, 0xf2, 0x97,
	0xf4, 0xa5, 0x24, 0x8d, 0x2a, 0xe2, 0x4f, 0x49, 0xf7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9,
	0x82, 0x74, 0x2d, 0xaa, 0x0a, 0x00, 0x00,
}

const ()

// SetOperationParams implements ydb generic interface for setting
// operation parameters inside driver implementation.
func (m *ExportToYtRequest) SetOperationParams(v *Ydb_Operations.OperationParams) {
	m.OperationParams = v
}

// SetOperationParams implements ydb generic interface for setting
// operation parameters inside driver implementation.
func (m *ExportToS3Request) SetOperationParams(v *Ydb_Operations.OperationParams) {
	m.OperationParams = v
}
