// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package api

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

type TaskMessage struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Time                 string   `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	EventId              int64    `protobuf:"varint,3,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	PartnerId            int64    `protobuf:"varint,4,opt,name=partner_id,json=partnerId,proto3" json:"partner_id,omitempty"`
	Revenue              int32    `protobuf:"varint,5,opt,name=revenue,proto3" json:"revenue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskMessage) Reset()         { *m = TaskMessage{} }
func (m *TaskMessage) String() string { return proto.CompactTextString(m) }
func (*TaskMessage) ProtoMessage()    {}
func (*TaskMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_657e7955b3ab0254, []int{0}
}
func (m *TaskMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskMessage.Unmarshal(m, b)
}
func (m *TaskMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskMessage.Marshal(b, m, deterministic)
}
func (dst *TaskMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskMessage.Merge(dst, src)
}
func (m *TaskMessage) XXX_Size() int {
	return xxx_messageInfo_TaskMessage.Size(m)
}
func (m *TaskMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TaskMessage proto.InternalMessageInfo

func (m *TaskMessage) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *TaskMessage) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *TaskMessage) GetEventId() int64 {
	if m != nil {
		return m.EventId
	}
	return 0
}

func (m *TaskMessage) GetPartnerId() int64 {
	if m != nil {
		return m.PartnerId
	}
	return 0
}

func (m *TaskMessage) GetRevenue() int32 {
	if m != nil {
		return m.Revenue
	}
	return 0
}

func init() {
	proto.RegisterType((*TaskMessage)(nil), "api.TaskMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatisticsClient is the client API for Statistics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatisticsClient interface {
	GetStatistics(ctx context.Context, in *TaskMessage, opts ...grpc.CallOption) (*TaskMessage, error)
}

type statisticsClient struct {
	cc *grpc.ClientConn
}

func NewStatisticsClient(cc *grpc.ClientConn) StatisticsClient {
	return &statisticsClient{cc}
}

func (c *statisticsClient) GetStatistics(ctx context.Context, in *TaskMessage, opts ...grpc.CallOption) (*TaskMessage, error) {
	out := new(TaskMessage)
	err := c.cc.Invoke(ctx, "/api.Statistics/GetStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatisticsServer is the server API for Statistics service.
type StatisticsServer interface {
	GetStatistics(context.Context, *TaskMessage) (*TaskMessage, error)
}

func RegisterStatisticsServer(s *grpc.Server, srv StatisticsServer) {
	s.RegisterService(&_Statistics_serviceDesc, srv)
}

func _Statistics_GetStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticsServer).GetStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Statistics/GetStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticsServer).GetStatistics(ctx, req.(*TaskMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Statistics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Statistics",
	HandlerType: (*StatisticsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatistics",
			Handler:    _Statistics_GetStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_data_657e7955b3ab0254) }

var fileDescriptor_data_657e7955b3ab0254 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0x2c, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0xea, 0x64, 0xe4, 0xe2,
	0x0e, 0x49, 0x2c, 0xce, 0xf6, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe2, 0x62, 0x49,
	0x49, 0x2c, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x41, 0x62, 0x25, 0x99,
	0xb9, 0xa9, 0x12, 0x4c, 0x10, 0x31, 0x10, 0x5b, 0x48, 0x92, 0x8b, 0x23, 0xb5, 0x2c, 0x35, 0xaf,
	0x24, 0x3e, 0x33, 0x45, 0x82, 0x59, 0x81, 0x51, 0x83, 0x39, 0x88, 0x1d, 0xcc, 0xf7, 0x4c, 0x11,
	0x92, 0xe5, 0xe2, 0x2a, 0x48, 0x2c, 0x2a, 0xc9, 0x4b, 0x2d, 0x02, 0x49, 0xb2, 0x80, 0x25, 0x39,
	0xa1, 0x22, 0x9e, 0x29, 0x42, 0x12, 0x5c, 0xec, 0x45, 0x20, 0xa5, 0xa5, 0xa9, 0x12, 0xac, 0x0a,
	0x8c, 0x1a, 0xac, 0x41, 0x30, 0xae, 0x91, 0x33, 0x17, 0x57, 0x70, 0x49, 0x62, 0x49, 0x66, 0x71,
	0x49, 0x66, 0x72, 0xb1, 0x90, 0x29, 0x17, 0xaf, 0x7b, 0x6a, 0x09, 0x92, 0x80, 0x80, 0x5e, 0x62,
	0x41, 0xa6, 0x1e, 0x92, 0x63, 0xa5, 0x30, 0x44, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x9e, 0x33, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x75, 0x59, 0x56, 0xf5, 0xea, 0x00, 0x00, 0x00,
}
