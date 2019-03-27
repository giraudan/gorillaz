// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

package stream

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type StreamRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{0}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StreamEvent struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=Key,json=key,proto3" json:"Key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=Value,json=value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamEvent) Reset()         { *m = StreamEvent{} }
func (m *StreamEvent) String() string { return proto.CompactTextString(m) }
func (*StreamEvent) ProtoMessage()    {}
func (*StreamEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{1}
}

func (m *StreamEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamEvent.Unmarshal(m, b)
}
func (m *StreamEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamEvent.Marshal(b, m, deterministic)
}
func (m *StreamEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEvent.Merge(m, src)
}
func (m *StreamEvent) XXX_Size() int {
	return xxx_messageInfo_StreamEvent.Size(m)
}
func (m *StreamEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEvent.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEvent proto.InternalMessageInfo

func (m *StreamEvent) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *StreamEvent) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "StreamRequest")
	proto.RegisterType((*StreamEvent)(nil), "StreamEvent")
}

func init() { proto.RegisterFile("stream.proto", fileDescriptor_bb17ef3f514bfe54) }

var fileDescriptor_bb17ef3f514bfe54 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe6, 0xe2, 0x0d, 0x06, 0xf3, 0x83,
	0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x53, 0x2e, 0x6e, 0x88, 0x22, 0xd7, 0xb2, 0xd4,
	0xbc, 0x12, 0x21, 0x01, 0x2e, 0x66, 0xef, 0xd4, 0x4a, 0xb0, 0x0a, 0x9e, 0x20, 0xe6, 0xec, 0xd4,
	0x4a, 0x21, 0x11, 0x2e, 0xd6, 0xb0, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0xb0, 0x18, 0x6b, 0x19,
	0x88, 0x63, 0x64, 0xc4, 0xc5, 0x06, 0xd1, 0x26, 0xa4, 0x01, 0x67, 0xf1, 0xe9, 0xa1, 0x58, 0x27,
	0xc5, 0xa3, 0x87, 0x64, 0xb2, 0x01, 0x63, 0x12, 0x1b, 0xd8, 0x59, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x07, 0xba, 0xb0, 0x86, 0xa6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamClient is the client API for Stream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamClient interface {
	Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Stream_StreamClient, error)
}

type streamClient struct {
	cc *grpc.ClientConn
}

func NewStreamClient(cc *grpc.ClientConn) StreamClient {
	return &streamClient{cc}
}

func (c *streamClient) Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Stream_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Stream_serviceDesc.Streams[0], "/Stream/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Stream_StreamClient interface {
	Recv() (*StreamEvent, error)
	grpc.ClientStream
}

type streamStreamClient struct {
	grpc.ClientStream
}

func (x *streamStreamClient) Recv() (*StreamEvent, error) {
	m := new(StreamEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServer is the server API for Stream service.
type StreamServer interface {
	Stream(*StreamRequest, Stream_StreamServer) error
}

func RegisterStreamServer(s *grpc.Server, srv StreamServer) {
	s.RegisterService(&_Stream_serviceDesc, srv)
}

func _Stream_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServer).Stream(m, &streamStreamServer{stream})
}

type Stream_StreamServer interface {
	Send(*StreamEvent) error
	grpc.ServerStream
}

type streamStreamServer struct {
	grpc.ServerStream
}

func (x *streamStreamServer) Send(m *StreamEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _Stream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Stream",
	HandlerType: (*StreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Stream_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stream.proto",
}
