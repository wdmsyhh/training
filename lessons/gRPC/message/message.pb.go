// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message/message.proto

/*
Package message is a generated protocol buffer package.

It is generated from these files:

	message/message.proto

It has these top-level messages:

	CalculateRequest
	CalculateResponse
	CheckResponse
*/
package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// 计算请求，三角形三条边
type CalculateRequest struct {
	FirstSide  float32 `protobuf:"fixed32,1,opt,name=firstSide" json:"firstSide,omitempty"`
	SecondSide float32 `protobuf:"fixed32,2,opt,name=secondSide" json:"secondSide,omitempty"`
	ThirdSide  float32 `protobuf:"fixed32,3,opt,name=thirdSide" json:"thirdSide,omitempty"`
}

func (m *CalculateRequest) Reset()                    { *m = CalculateRequest{} }
func (m *CalculateRequest) String() string            { return proto.CompactTextString(m) }
func (*CalculateRequest) ProtoMessage()               {}
func (*CalculateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CalculateRequest) GetFirstSide() float32 {
	if m != nil {
		return m.FirstSide
	}
	return 0
}

func (m *CalculateRequest) GetSecondSide() float32 {
	if m != nil {
		return m.SecondSide
	}
	return 0
}

func (m *CalculateRequest) GetThirdSide() float32 {
	if m != nil {
		return m.ThirdSide
	}
	return 0
}

type CalculateResponse struct {
	// 计算结果
	CalculateResult float32 `protobuf:"fixed32,1,opt,name=calculateResult" json:"calculateResult,omitempty"`
	// 描述计算的是周长或面积
	CircumferenceOrArea string `protobuf:"bytes,2,opt,name=circumferenceOrArea" json:"circumferenceOrArea,omitempty"`
}

func (m *CalculateResponse) Reset()                    { *m = CalculateResponse{} }
func (m *CalculateResponse) String() string            { return proto.CompactTextString(m) }
func (*CalculateResponse) ProtoMessage()               {}
func (*CalculateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CalculateResponse) GetCalculateResult() float32 {
	if m != nil {
		return m.CalculateResult
	}
	return 0
}

func (m *CalculateResponse) GetCircumferenceOrArea() string {
	if m != nil {
		return m.CircumferenceOrArea
	}
	return ""
}

type CheckResponse struct {
	// 结果为 false 则不是三角形，为 true 是三角形
	CheckResult bool `protobuf:"varint,1,opt,name=checkResult" json:"checkResult,omitempty"`
}

func (m *CheckResponse) Reset()                    { *m = CheckResponse{} }
func (m *CheckResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()               {}
func (*CheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CheckResponse) GetCheckResult() bool {
	if m != nil {
		return m.CheckResult
	}
	return false
}

func init() {
	proto.RegisterType((*CalculateRequest)(nil), "message.CalculateRequest")
	proto.RegisterType((*CalculateResponse)(nil), "message.CalculateResponse")
	proto.RegisterType((*CheckResponse)(nil), "message.CheckResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CalculateService service

type CalculateServiceClient interface {
	// 计算周长或面积
	CalculateCircumferenceOrArea(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
	// 判断是否是三角形
	CheckTriangle(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type calculateServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculateServiceClient(cc *grpc.ClientConn) CalculateServiceClient {
	return &calculateServiceClient{cc}
}

func (c *calculateServiceClient) CalculateCircumferenceOrArea(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := grpc.Invoke(ctx, "/message.CalculateService/CalculateCircumferenceOrArea", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateServiceClient) CheckTriangle(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := grpc.Invoke(ctx, "/message.CalculateService/CheckTriangle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CalculateService service

type CalculateServiceServer interface {
	// 计算周长或面积
	CalculateCircumferenceOrArea(context.Context, *CalculateRequest) (*CalculateResponse, error)
	// 判断是否是三角形
	CheckTriangle(context.Context, *CalculateRequest) (*CheckResponse, error)
}

func RegisterCalculateServiceServer(s *grpc.Server, srv CalculateServiceServer) {
	s.RegisterService(&_CalculateService_serviceDesc, srv)
}

func _CalculateService_CalculateCircumferenceOrArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).CalculateCircumferenceOrArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.CalculateService/CalculateCircumferenceOrArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).CalculateCircumferenceOrArea(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculateService_CheckTriangle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).CheckTriangle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.CalculateService/CheckTriangle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).CheckTriangle(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.CalculateService",
	HandlerType: (*CalculateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateCircumferenceOrArea",
			Handler:    _CalculateService_CalculateCircumferenceOrArea_Handler,
		},
		{
			MethodName: "CheckTriangle",
			Handler:    _CalculateService_CheckTriangle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message/message.proto",
}

func init() { proto.RegisterFile("message/message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x5d, 0x4e, 0x02, 0x31,
	0x14, 0x85, 0x33, 0x98, 0xa8, 0x5c, 0x63, 0xd4, 0x1a, 0x8d, 0x12, 0x62, 0xc8, 0x3c, 0xf1, 0x84,
	0x7f, 0x2b, 0xd0, 0x59, 0x80, 0xc9, 0xa0, 0x0b, 0xa8, 0x97, 0x03, 0x34, 0x0e, 0x2d, 0xde, 0x76,
	0xdc, 0x95, 0x7b, 0x34, 0x0e, 0x65, 0x68, 0x90, 0xf8, 0x34, 0x99, 0xef, 0xcb, 0xed, 0xb9, 0x3d,
	0xa5, 0x8b, 0x05, 0xbc, 0xd7, 0x33, 0xdc, 0xc6, 0xef, 0x68, 0x29, 0x2e, 0x38, 0x75, 0x10, 0x7f,
	0x73, 0x4b, 0xa7, 0x85, 0xae, 0xb8, 0xae, 0x74, 0x40, 0x89, 0xcf, 0x1a, 0x3e, 0xa8, 0x3e, 0x75,
	0xa7, 0x46, 0x7c, 0x18, 0x9b, 0x09, 0xae, 0xb2, 0x41, 0x36, 0xec, 0x94, 0x1b, 0xa0, 0x6e, 0x88,
	0x3c, 0xd8, 0xd9, 0x49, 0xa3, 0x3b, 0x8d, 0x4e, 0xc8, 0xef, 0x74, 0x98, 0x1b, 0x59, 0xe9, 0xbd,
	0xd5, 0x74, 0x0b, 0x72, 0x47, 0x67, 0x49, 0x9e, 0x5f, 0x3a, 0xeb, 0xa1, 0x86, 0x74, 0xc2, 0x09,
	0xac, 0xab, 0x10, 0x63, 0xb7, 0xb1, 0xba, 0xa3, 0x73, 0x36, 0xc2, 0xf5, 0x62, 0x0a, 0x81, 0x65,
	0xbc, 0xc8, 0x93, 0x40, 0x37, 0x5b, 0x74, 0xcb, 0x5d, 0x2a, 0xbf, 0xa7, 0xe3, 0x62, 0x0e, 0xfe,
	0x68, 0xc3, 0x06, 0x74, 0xc4, 0x11, 0xac, 0x83, 0x0e, 0xcb, 0x14, 0x3d, 0x7c, 0x67, 0x49, 0x29,
	0x63, 0xc8, 0x97, 0x61, 0xa8, 0x37, 0xea, 0xb7, 0xac, 0xf8, 0x9b, 0xa3, 0xae, 0x47, 0xeb, 0x86,
	0xb7, 0xfb, 0xec, 0xf5, 0x76, 0xa9, 0xb8, 0xcd, 0x73, 0x5c, 0xef, 0x55, 0x8c, 0xb6, 0xb3, 0x0a,
	0xff, 0x9d, 0x73, 0xb9, 0x51, 0xe9, 0x8d, 0xde, 0xf7, 0x9b, 0x37, 0x7d, 0xfc, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xfa, 0xfd, 0xf2, 0x85, 0xec, 0x01, 0x00, 0x00,
}
