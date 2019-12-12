// Code generated by protoc-gen-go. DO NOT EDIT.
// source: communication-testing-service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type TaskType int32

const (
	TaskType_PING    TaskType = 0
	TaskType_SLEEP   TaskType = 1
	TaskType_COMPUTE TaskType = 2
)

var TaskType_name = map[int32]string{
	0: "PING",
	1: "SLEEP",
	2: "COMPUTE",
}

var TaskType_value = map[string]int32{
	"PING":    0,
	"SLEEP":   1,
	"COMPUTE": 2,
}

func (x TaskType) String() string {
	return proto.EnumName(TaskType_name, int32(x))
}

func (TaskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_47327da0f66dd903, []int{0}
}

type TaskRequest struct {
	SenderCounter        int64                `protobuf:"varint,1,opt,name=senderCounter,proto3" json:"senderCounter,omitempty"`
	SenderName           string               `protobuf:"bytes,2,opt,name=senderName,proto3" json:"senderName,omitempty"`
	SentTimestamp        *timestamp.Timestamp `protobuf:"bytes,3,opt,name=sentTimestamp,proto3" json:"sentTimestamp,omitempty"`
	TaskType             TaskType             `protobuf:"varint,4,opt,name=taskType,proto3,enum=v1.TaskType" json:"taskType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskRequest) Reset()         { *m = TaskRequest{} }
func (m *TaskRequest) String() string { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()    {}
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47327da0f66dd903, []int{0}
}

func (m *TaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRequest.Unmarshal(m, b)
}
func (m *TaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRequest.Marshal(b, m, deterministic)
}
func (m *TaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRequest.Merge(m, src)
}
func (m *TaskRequest) XXX_Size() int {
	return xxx_messageInfo_TaskRequest.Size(m)
}
func (m *TaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRequest proto.InternalMessageInfo

func (m *TaskRequest) GetSenderCounter() int64 {
	if m != nil {
		return m.SenderCounter
	}
	return 0
}

func (m *TaskRequest) GetSenderName() string {
	if m != nil {
		return m.SenderName
	}
	return ""
}

func (m *TaskRequest) GetSentTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.SentTimestamp
	}
	return nil
}

func (m *TaskRequest) GetTaskType() TaskType {
	if m != nil {
		return m.TaskType
	}
	return TaskType_PING
}

type TaskResponse struct {
	SenderCounter        int64                `protobuf:"varint,1,opt,name=senderCounter,proto3" json:"senderCounter,omitempty"`
	SenderName           string               `protobuf:"bytes,2,opt,name=senderName,proto3" json:"senderName,omitempty"`
	SentTimestamp        *timestamp.Timestamp `protobuf:"bytes,3,opt,name=sentTimestamp,proto3" json:"sentTimestamp,omitempty"`
	ReceiverCounter      int64                `protobuf:"varint,4,opt,name=receiverCounter,proto3" json:"receiverCounter,omitempty"`
	ReceiverName         string               `protobuf:"bytes,5,opt,name=receiverName,proto3" json:"receiverName,omitempty"`
	ReceivedTimestamp    *timestamp.Timestamp `protobuf:"bytes,6,opt,name=receivedTimestamp,proto3" json:"receivedTimestamp,omitempty"`
	RespondedTimestamp   *timestamp.Timestamp `protobuf:"bytes,7,opt,name=respondedTimestamp,proto3" json:"respondedTimestamp,omitempty"`
	TaskType             TaskType             `protobuf:"varint,8,opt,name=taskType,proto3,enum=v1.TaskType" json:"taskType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskResponse) Reset()         { *m = TaskResponse{} }
func (m *TaskResponse) String() string { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()    {}
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47327da0f66dd903, []int{1}
}

func (m *TaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResponse.Unmarshal(m, b)
}
func (m *TaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResponse.Marshal(b, m, deterministic)
}
func (m *TaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResponse.Merge(m, src)
}
func (m *TaskResponse) XXX_Size() int {
	return xxx_messageInfo_TaskResponse.Size(m)
}
func (m *TaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResponse proto.InternalMessageInfo

func (m *TaskResponse) GetSenderCounter() int64 {
	if m != nil {
		return m.SenderCounter
	}
	return 0
}

func (m *TaskResponse) GetSenderName() string {
	if m != nil {
		return m.SenderName
	}
	return ""
}

func (m *TaskResponse) GetSentTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.SentTimestamp
	}
	return nil
}

func (m *TaskResponse) GetReceiverCounter() int64 {
	if m != nil {
		return m.ReceiverCounter
	}
	return 0
}

func (m *TaskResponse) GetReceiverName() string {
	if m != nil {
		return m.ReceiverName
	}
	return ""
}

func (m *TaskResponse) GetReceivedTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.ReceivedTimestamp
	}
	return nil
}

func (m *TaskResponse) GetRespondedTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.RespondedTimestamp
	}
	return nil
}

func (m *TaskResponse) GetTaskType() TaskType {
	if m != nil {
		return m.TaskType
	}
	return TaskType_PING
}

func init() {
	proto.RegisterEnum("v1.TaskType", TaskType_name, TaskType_value)
	proto.RegisterType((*TaskRequest)(nil), "v1.TaskRequest")
	proto.RegisterType((*TaskResponse)(nil), "v1.TaskResponse")
}

func init() {
	proto.RegisterFile("communication-testing-service.proto", fileDescriptor_47327da0f66dd903)
}

var fileDescriptor_47327da0f66dd903 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0xdf, 0x4e, 0xe2, 0x50,
	0x10, 0x87, 0xb7, 0xfc, 0x2d, 0x03, 0xbb, 0x74, 0xe7, 0xaa, 0x61, 0x93, 0xb5, 0x41, 0x2f, 0x1a,
	0x95, 0x12, 0xf0, 0x05, 0x4c, 0x08, 0x11, 0x89, 0x22, 0x29, 0xf5, 0x01, 0x4a, 0x3b, 0x92, 0x06,
	0xda, 0x53, 0x7b, 0x4e, 0x9b, 0xf8, 0x3e, 0xbe, 0x87, 0xaf, 0x66, 0xd2, 0x52, 0xa0, 0x6a, 0xe4,
	0xd6, 0xcb, 0x7e, 0x3d, 0x33, 0xe7, 0x9b, 0x33, 0x3f, 0x38, 0x75, 0x98, 0xef, 0xc7, 0x81, 0xe7,
	0xd8, 0xc2, 0x63, 0x41, 0x4f, 0x10, 0x17, 0x5e, 0xb0, 0xea, 0x71, 0x8a, 0x12, 0xcf, 0x21, 0x23,
	0x8c, 0x98, 0x60, 0x58, 0x4a, 0x06, 0x9d, 0x93, 0x15, 0x63, 0xab, 0x0d, 0xf5, 0x53, 0xb2, 0x8c,
	0x9f, 0xfa, 0xc2, 0xf3, 0x89, 0x0b, 0xdb, 0x0f, 0xb3, 0x43, 0xdd, 0x37, 0x09, 0x9a, 0x96, 0xcd,
	0xd7, 0x26, 0x3d, 0xc7, 0xc4, 0x05, 0x9e, 0xc1, 0x6f, 0x4e, 0x81, 0x4b, 0xd1, 0x88, 0xc5, 0x81,
	0xa0, 0x48, 0x95, 0x34, 0x49, 0x2f, 0x9b, 0x45, 0x88, 0xff, 0x01, 0x32, 0x30, 0xb3, 0x7d, 0x52,
	0x4b, 0x9a, 0xa4, 0x37, 0xcc, 0x03, 0x82, 0xd7, 0x69, 0x17, 0x61, 0xe5, 0x97, 0xa9, 0x65, 0x4d,
	0xd2, 0x9b, 0xc3, 0x8e, 0x91, 0xe9, 0x18, 0xb9, 0x8e, 0xb1, 0x3b, 0x61, 0x16, 0x0b, 0x50, 0x07,
	0x59, 0xd8, 0x7c, 0x6d, 0xbd, 0x84, 0xa4, 0x56, 0x34, 0x49, 0xff, 0x33, 0x6c, 0x19, 0xc9, 0xc0,
	0xb0, 0xb6, 0xcc, 0xdc, 0xfd, 0xed, 0xbe, 0x96, 0xa1, 0x95, 0x4d, 0xc0, 0x43, 0x16, 0x70, 0xfa,
	0x41, 0x23, 0xb4, 0x23, 0x72, 0xc8, 0x4b, 0xf6, 0x26, 0x95, 0xd4, 0xe4, 0x23, 0xc6, 0x2e, 0xb4,
	0x72, 0x94, 0xda, 0x54, 0x53, 0x9b, 0x02, 0xc3, 0x09, 0xfc, 0xdd, 0x7e, 0xbb, 0x7b, 0xa7, 0xda,
	0x51, 0xa7, 0xcf, 0x45, 0x38, 0x05, 0x8c, 0xd2, 0xb7, 0x72, 0x0f, 0x5b, 0xd5, 0x8f, 0xb6, 0xfa,
	0xa2, 0xaa, 0xb0, 0x26, 0xf9, 0xbb, 0x35, 0x9d, 0x5f, 0x82, 0x9c, 0x53, 0x94, 0xa1, 0x32, 0xbf,
	0x9d, 0xdd, 0x28, 0xbf, 0xb0, 0x01, 0xd5, 0xc5, 0xdd, 0x78, 0x3c, 0x57, 0x24, 0x6c, 0x42, 0x7d,
	0xf4, 0x70, 0x3f, 0x7f, 0xb4, 0xc6, 0x4a, 0x69, 0x38, 0x85, 0x7f, 0xa3, 0xc3, 0x88, 0x5b, 0x59,
	0xc2, 0x17, 0x59, 0xc0, 0xf1, 0x02, 0x6a, 0x13, 0x3b, 0x70, 0x37, 0x84, 0xed, 0xfc, 0xba, 0x6d,
	0x80, 0x3b, 0xca, 0x1e, 0x64, 0x79, 0x58, 0xd6, 0xd2, 0x59, 0xae, 0xde, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xe3, 0x55, 0xe8, 0x25, 0x35, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommunicationTestingServiceClient is the client API for CommunicationTestingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommunicationTestingServiceClient interface {
	Handle(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
}

type communicationTestingServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommunicationTestingServiceClient(cc *grpc.ClientConn) CommunicationTestingServiceClient {
	return &communicationTestingServiceClient{cc}
}

func (c *communicationTestingServiceClient) Handle(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/v1.CommunicationTestingService/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationTestingServiceServer is the server API for CommunicationTestingService service.
type CommunicationTestingServiceServer interface {
	Handle(context.Context, *TaskRequest) (*TaskResponse, error)
}

// UnimplementedCommunicationTestingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommunicationTestingServiceServer struct {
}

func (*UnimplementedCommunicationTestingServiceServer) Handle(ctx context.Context, req *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterCommunicationTestingServiceServer(s *grpc.Server, srv CommunicationTestingServiceServer) {
	s.RegisterService(&_CommunicationTestingService_serviceDesc, srv)
}

func _CommunicationTestingService_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationTestingServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CommunicationTestingService/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationTestingServiceServer).Handle(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommunicationTestingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.CommunicationTestingService",
	HandlerType: (*CommunicationTestingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _CommunicationTestingService_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communication-testing-service.proto",
}