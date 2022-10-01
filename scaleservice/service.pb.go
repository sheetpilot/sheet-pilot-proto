// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ScaleRequest struct {
	Sheet                *ScaleRequest_Sheet `protobuf:"bytes,1,opt,name=sheet,proto3" json:"sheet,omitempty"`
	App                  *ScaleRequest_App   `protobuf:"bytes,2,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ScaleRequest) Reset()         { *m = ScaleRequest{} }
func (m *ScaleRequest) String() string { return proto.CompactTextString(m) }
func (*ScaleRequest) ProtoMessage()    {}
func (*ScaleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *ScaleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleRequest.Unmarshal(m, b)
}
func (m *ScaleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleRequest.Marshal(b, m, deterministic)
}
func (m *ScaleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleRequest.Merge(m, src)
}
func (m *ScaleRequest) XXX_Size() int {
	return xxx_messageInfo_ScaleRequest.Size(m)
}
func (m *ScaleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleRequest proto.InternalMessageInfo

func (m *ScaleRequest) GetSheet() *ScaleRequest_Sheet {
	if m != nil {
		return m.Sheet
	}
	return nil
}

func (m *ScaleRequest) GetApp() *ScaleRequest_App {
	if m != nil {
		return m.App
	}
	return nil
}

type ScaleRequest_Row struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScaleRequest_Row) Reset()         { *m = ScaleRequest_Row{} }
func (m *ScaleRequest_Row) String() string { return proto.CompactTextString(m) }
func (*ScaleRequest_Row) ProtoMessage()    {}
func (*ScaleRequest_Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 0}
}

func (m *ScaleRequest_Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleRequest_Row.Unmarshal(m, b)
}
func (m *ScaleRequest_Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleRequest_Row.Marshal(b, m, deterministic)
}
func (m *ScaleRequest_Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleRequest_Row.Merge(m, src)
}
func (m *ScaleRequest_Row) XXX_Size() int {
	return xxx_messageInfo_ScaleRequest_Row.Size(m)
}
func (m *ScaleRequest_Row) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleRequest_Row.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleRequest_Row proto.InternalMessageInfo

type ScaleRequest_Col struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScaleRequest_Col) Reset()         { *m = ScaleRequest_Col{} }
func (m *ScaleRequest_Col) String() string { return proto.CompactTextString(m) }
func (*ScaleRequest_Col) ProtoMessage()    {}
func (*ScaleRequest_Col) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 1}
}

func (m *ScaleRequest_Col) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleRequest_Col.Unmarshal(m, b)
}
func (m *ScaleRequest_Col) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleRequest_Col.Marshal(b, m, deterministic)
}
func (m *ScaleRequest_Col) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleRequest_Col.Merge(m, src)
}
func (m *ScaleRequest_Col) XXX_Size() int {
	return xxx_messageInfo_ScaleRequest_Col.Size(m)
}
func (m *ScaleRequest_Col) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleRequest_Col.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleRequest_Col proto.InternalMessageInfo

type ScaleRequest_Sheet struct {
	Row                  *ScaleRequest_Row `protobuf:"bytes,1,opt,name=row,proto3" json:"row,omitempty"`
	Col                  *ScaleRequest_Col `protobuf:"bytes,2,opt,name=col,proto3" json:"col,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ScaleRequest_Sheet) Reset()         { *m = ScaleRequest_Sheet{} }
func (m *ScaleRequest_Sheet) String() string { return proto.CompactTextString(m) }
func (*ScaleRequest_Sheet) ProtoMessage()    {}
func (*ScaleRequest_Sheet) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 2}
}

func (m *ScaleRequest_Sheet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleRequest_Sheet.Unmarshal(m, b)
}
func (m *ScaleRequest_Sheet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleRequest_Sheet.Marshal(b, m, deterministic)
}
func (m *ScaleRequest_Sheet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleRequest_Sheet.Merge(m, src)
}
func (m *ScaleRequest_Sheet) XXX_Size() int {
	return xxx_messageInfo_ScaleRequest_Sheet.Size(m)
}
func (m *ScaleRequest_Sheet) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleRequest_Sheet.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleRequest_Sheet proto.InternalMessageInfo

func (m *ScaleRequest_Sheet) GetRow() *ScaleRequest_Row {
	if m != nil {
		return m.Row
	}
	return nil
}

func (m *ScaleRequest_Sheet) GetCol() *ScaleRequest_Col {
	if m != nil {
		return m.Col
	}
	return nil
}

type ScaleRequest_Request struct {
	Cpu                  *field_mask.FieldMask `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Mem                  *field_mask.FieldMask `protobuf:"bytes,2,opt,name=mem,proto3" json:"mem,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ScaleRequest_Request) Reset()         { *m = ScaleRequest_Request{} }
func (m *ScaleRequest_Request) String() string { return proto.CompactTextString(m) }
func (*ScaleRequest_Request) ProtoMessage()    {}
func (*ScaleRequest_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 3}
}

func (m *ScaleRequest_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleRequest_Request.Unmarshal(m, b)
}
func (m *ScaleRequest_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleRequest_Request.Marshal(b, m, deterministic)
}
func (m *ScaleRequest_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleRequest_Request.Merge(m, src)
}
func (m *ScaleRequest_Request) XXX_Size() int {
	return xxx_messageInfo_ScaleRequest_Request.Size(m)
}
func (m *ScaleRequest_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleRequest_Request.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleRequest_Request proto.InternalMessageInfo

func (m *ScaleRequest_Request) GetCpu() *field_mask.FieldMask {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *ScaleRequest_Request) GetMem() *field_mask.FieldMask {
	if m != nil {
		return m.Mem
	}
	return nil
}

type ScaleRequest_Limit struct {
	Cpu                  *field_mask.FieldMask `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Mem                  *field_mask.FieldMask `protobuf:"bytes,2,opt,name=mem,proto3" json:"mem,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ScaleRequest_Limit) Reset()         { *m = ScaleRequest_Limit{} }
func (m *ScaleRequest_Limit) String() string { return proto.CompactTextString(m) }
func (*ScaleRequest_Limit) ProtoMessage()    {}
func (*ScaleRequest_Limit) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 4}
}

func (m *ScaleRequest_Limit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleRequest_Limit.Unmarshal(m, b)
}
func (m *ScaleRequest_Limit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleRequest_Limit.Marshal(b, m, deterministic)
}
func (m *ScaleRequest_Limit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleRequest_Limit.Merge(m, src)
}
func (m *ScaleRequest_Limit) XXX_Size() int {
	return xxx_messageInfo_ScaleRequest_Limit.Size(m)
}
func (m *ScaleRequest_Limit) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleRequest_Limit.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleRequest_Limit proto.InternalMessageInfo

func (m *ScaleRequest_Limit) GetCpu() *field_mask.FieldMask {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *ScaleRequest_Limit) GetMem() *field_mask.FieldMask {
	if m != nil {
		return m.Mem
	}
	return nil
}

type ScaleRequest_App struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Request              *ScaleRequest_Request `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Limit                *ScaleRequest_Limit   `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Replica              uint32                `protobuf:"varint,4,opt,name=replica,proto3" json:"replica,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ScaleRequest_App) Reset()         { *m = ScaleRequest_App{} }
func (m *ScaleRequest_App) String() string { return proto.CompactTextString(m) }
func (*ScaleRequest_App) ProtoMessage()    {}
func (*ScaleRequest_App) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 5}
}

func (m *ScaleRequest_App) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleRequest_App.Unmarshal(m, b)
}
func (m *ScaleRequest_App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleRequest_App.Marshal(b, m, deterministic)
}
func (m *ScaleRequest_App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleRequest_App.Merge(m, src)
}
func (m *ScaleRequest_App) XXX_Size() int {
	return xxx_messageInfo_ScaleRequest_App.Size(m)
}
func (m *ScaleRequest_App) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleRequest_App.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleRequest_App proto.InternalMessageInfo

func (m *ScaleRequest_App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScaleRequest_App) GetRequest() *ScaleRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ScaleRequest_App) GetLimit() *ScaleRequest_Limit {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *ScaleRequest_App) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

type ScaleResponse struct {
	Data                 string               `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Message              string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error                *ScaleResponse_ERROR `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ScaleResponse) Reset()         { *m = ScaleResponse{} }
func (m *ScaleResponse) String() string { return proto.CompactTextString(m) }
func (*ScaleResponse) ProtoMessage()    {}
func (*ScaleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *ScaleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleResponse.Unmarshal(m, b)
}
func (m *ScaleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleResponse.Marshal(b, m, deterministic)
}
func (m *ScaleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleResponse.Merge(m, src)
}
func (m *ScaleResponse) XXX_Size() int {
	return xxx_messageInfo_ScaleResponse.Size(m)
}
func (m *ScaleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleResponse proto.InternalMessageInfo

func (m *ScaleResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ScaleResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ScaleResponse) GetError() *ScaleResponse_ERROR {
	if m != nil {
		return m.Error
	}
	return nil
}

type ScaleResponse_ERROR struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScaleResponse_ERROR) Reset()         { *m = ScaleResponse_ERROR{} }
func (m *ScaleResponse_ERROR) String() string { return proto.CompactTextString(m) }
func (*ScaleResponse_ERROR) ProtoMessage()    {}
func (*ScaleResponse_ERROR) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1, 0}
}

func (m *ScaleResponse_ERROR) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleResponse_ERROR.Unmarshal(m, b)
}
func (m *ScaleResponse_ERROR) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleResponse_ERROR.Marshal(b, m, deterministic)
}
func (m *ScaleResponse_ERROR) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleResponse_ERROR.Merge(m, src)
}
func (m *ScaleResponse_ERROR) XXX_Size() int {
	return xxx_messageInfo_ScaleResponse_ERROR.Size(m)
}
func (m *ScaleResponse_ERROR) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleResponse_ERROR.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleResponse_ERROR proto.InternalMessageInfo

func (m *ScaleResponse_ERROR) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ScaleResponse_ERROR) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ScaleRequest)(nil), "pb.ScaleRequest")
	proto.RegisterType((*ScaleRequest_Row)(nil), "pb.ScaleRequest.Row")
	proto.RegisterType((*ScaleRequest_Col)(nil), "pb.ScaleRequest.Col")
	proto.RegisterType((*ScaleRequest_Sheet)(nil), "pb.ScaleRequest.Sheet")
	proto.RegisterType((*ScaleRequest_Request)(nil), "pb.ScaleRequest.Request")
	proto.RegisterType((*ScaleRequest_Limit)(nil), "pb.ScaleRequest.Limit")
	proto.RegisterType((*ScaleRequest_App)(nil), "pb.ScaleRequest.App")
	proto.RegisterType((*ScaleResponse)(nil), "pb.ScaleResponse")
	proto.RegisterType((*ScaleResponse_ERROR)(nil), "pb.ScaleResponse.ERROR")
}

func init() {
	proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626)
}

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xb1, 0xee, 0xd3, 0x30,
	0x10, 0xc6, 0x95, 0xa6, 0xa1, 0xaa, 0xa1, 0x12, 0x18, 0x04, 0x91, 0xa7, 0xaa, 0x03, 0xea, 0x00,
	0xae, 0x54, 0xc4, 0xc2, 0x56, 0x55, 0x30, 0x20, 0x10, 0xd2, 0x65, 0x60, 0x44, 0x4e, 0x7a, 0x2d,
	0x51, 0xed, 0xda, 0xc4, 0x29, 0x7d, 0x12, 0x9e, 0x80, 0x07, 0xe0, 0x15, 0xd1, 0xd9, 0x89, 0xa8,
	0x88, 0xfa, 0xdf, 0xfe, 0x93, 0x7d, 0xed, 0xef, 0xfc, 0x7d, 0xfe, 0xce, 0x61, 0x33, 0x8f, 0xcd,
	0xcf, 0xba, 0x42, 0xe9, 0x1a, 0xdb, 0x5a, 0x3e, 0x72, 0xa5, 0x98, 0x1f, 0xac, 0x3d, 0x68, 0x5c,
	0x85, 0x5f, 0xca, 0xf3, 0x7e, 0xb5, 0xaf, 0x51, 0xef, 0xbe, 0x19, 0xe5, 0x8f, 0x91, 0x5a, 0xfc,
	0x19, 0xb3, 0x47, 0x45, 0xa5, 0x34, 0x02, 0xfe, 0x38, 0xa3, 0x6f, 0xf9, 0x2b, 0x96, 0xf9, 0xef,
	0x88, 0x6d, 0x9e, 0xcc, 0x93, 0xe5, 0xc3, 0xf5, 0x73, 0xe9, 0x4a, 0x79, 0x0d, 0xc8, 0x82, 0xfe,
	0x85, 0x08, 0xf1, 0x97, 0x2c, 0x55, 0xce, 0xe5, 0xa3, 0xc0, 0x3e, 0x1b, 0xb0, 0x1b, 0xe7, 0x80,
	0x00, 0x91, 0xb1, 0x14, 0xec, 0x85, 0x96, 0xad, 0xd5, 0xe2, 0x2b, 0xcb, 0x8a, 0xbe, 0xbd, 0xb1,
	0x97, 0x4e, 0x6a, 0xd8, 0x0e, 0xf6, 0x02, 0x04, 0x10, 0x57, 0x59, 0x7d, 0x53, 0x66, 0x6b, 0x35,
	0x10, 0x20, 0x90, 0x4d, 0xfe, 0xdd, 0x23, 0xad, 0xdc, 0xb9, 0x3b, 0x5a, 0xc8, 0x18, 0x84, 0xec,
	0x83, 0x90, 0x1f, 0x28, 0x88, 0xcf, 0xca, 0x1f, 0x81, 0x30, 0xa2, 0x0d, 0x9a, 0x4e, 0xe0, 0x4e,
	0xda, 0xa0, 0x11, 0x15, 0xcb, 0x3e, 0xd5, 0xa6, 0xbe, 0x5f, 0x91, 0x5f, 0x09, 0x4b, 0x37, 0xce,
	0x71, 0xce, 0xc6, 0x27, 0x65, 0x30, 0x88, 0x4c, 0x21, 0xec, 0xf9, 0x9a, 0x4d, 0x9a, 0x78, 0xcf,
	0xee, 0xb4, 0x7c, 0x98, 0x5d, 0x5c, 0xa1, 0x07, 0x69, 0xb0, 0x9a, 0x4c, 0xe7, 0xe9, 0x8d, 0xc1,
	0x86, 0x2b, 0x41, 0x84, 0x78, 0x4e, 0x0a, 0x4e, 0xd7, 0x95, 0xca, 0xc7, 0xf3, 0x64, 0x39, 0x83,
	0xbe, 0x5c, 0xfc, 0x4e, 0xd8, 0xac, 0xeb, 0xf3, 0xce, 0x9e, 0x3c, 0x92, 0xc3, 0x9d, 0x6a, 0x55,
	0xef, 0x90, 0xf6, 0xd4, 0x6f, 0xd0, 0x7b, 0x75, 0xc0, 0xe0, 0x70, 0x0a, 0x7d, 0xc9, 0x5f, 0xb3,
	0x0c, 0x9b, 0xc6, 0x36, 0x9d, 0x8f, 0x17, 0x57, 0x3e, 0xe2, 0x79, 0xf2, 0x3d, 0xc0, 0x17, 0x80,
	0x48, 0x89, 0xb7, 0x2c, 0x0b, 0x35, 0xa9, 0x54, 0x76, 0x17, 0x73, 0xc8, 0x20, 0xec, 0x6f, 0xab,
	0xac, 0x3f, 0x76, 0xcf, 0xba, 0x88, 0xdf, 0x04, 0x7f, 0xc7, 0x9e, 0x5e, 0xd7, 0xfd, 0x2b, 0x79,
	0xfc, 0x7f, 0x0a, 0xe2, 0xc9, 0xc0, 0x4f, 0xf9, 0x20, 0x8c, 0xe8, 0xcd, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x3e, 0xf0, 0x45, 0x87, 0x61, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ScaleServiceClient is the client API for ScaleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScaleServiceClient interface {
	ScaleServiceRequest(ctx context.Context, in *ScaleRequest, opts ...grpc.CallOption) (*ScaleResponse, error)
}

type scaleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScaleServiceClient(cc grpc.ClientConnInterface) ScaleServiceClient {
	return &scaleServiceClient{cc}
}

func (c *scaleServiceClient) ScaleServiceRequest(ctx context.Context, in *ScaleRequest, opts ...grpc.CallOption) (*ScaleResponse, error) {
	out := new(ScaleResponse)
	err := c.cc.Invoke(ctx, "/pb.ScaleService/ScaleServiceRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScaleServiceServer is the server API for ScaleService service.
type ScaleServiceServer interface {
	ScaleServiceRequest(context.Context, *ScaleRequest) (*ScaleResponse, error)
}

// UnimplementedScaleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedScaleServiceServer struct {
}

func (*UnimplementedScaleServiceServer) ScaleServiceRequest(ctx context.Context, req *ScaleRequest) (*ScaleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScaleServiceRequest not implemented")
}

func RegisterScaleServiceServer(s *grpc.Server, srv ScaleServiceServer) {
	s.RegisterService(&_ScaleService_serviceDesc, srv)
}

func _ScaleService_ScaleServiceRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScaleServiceServer).ScaleServiceRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ScaleService/ScaleServiceRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScaleServiceServer).ScaleServiceRequest(ctx, req.(*ScaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScaleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ScaleService",
	HandlerType: (*ScaleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScaleServiceRequest",
			Handler:    _ScaleService_ScaleServiceRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
