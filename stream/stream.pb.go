// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

package stream

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type StreamRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RequesterName        string   `protobuf:"bytes,2,opt,name=requesterName,proto3" json:"requesterName,omitempty"`
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

func (m *StreamRequest) GetRequesterName() string {
	if m != nil {
		return m.RequesterName
	}
	return ""
}

type StreamEvent struct {
	Key                  []byte    `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value                []byte    `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Metadata             *Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
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

func (m *StreamEvent) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Metadata struct {
	EventTimestamp        int64             `protobuf:"varint,1,opt,name=EventTimestamp,proto3" json:"EventTimestamp,omitempty"`
	OriginStreamTimestamp int64             `protobuf:"varint,2,opt,name=OriginStreamTimestamp,proto3" json:"OriginStreamTimestamp,omitempty"`
	StreamTimestamp       int64             `protobuf:"varint,3,opt,name=StreamTimestamp,proto3" json:"StreamTimestamp,omitempty"`
	KeyValue              map[string]string `protobuf:"bytes,4,rep,name=keyValue,proto3" json:"keyValue,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral  struct{}          `json:"-"`
	XXX_unrecognized      []byte            `json:"-"`
	XXX_sizecache         int32             `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{2}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetEventTimestamp() int64 {
	if m != nil {
		return m.EventTimestamp
	}
	return 0
}

func (m *Metadata) GetOriginStreamTimestamp() int64 {
	if m != nil {
		return m.OriginStreamTimestamp
	}
	return 0
}

func (m *Metadata) GetStreamTimestamp() int64 {
	if m != nil {
		return m.StreamTimestamp
	}
	return 0
}

func (m *Metadata) GetKeyValue() map[string]string {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

type StreamDefinition struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DataType             string   `protobuf:"bytes,2,opt,name=dataType,proto3" json:"dataType,omitempty"`
	IsDelete             bool     `protobuf:"varint,3,opt,name=isDelete,proto3" json:"isDelete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamDefinition) Reset()         { *m = StreamDefinition{} }
func (m *StreamDefinition) String() string { return proto.CompactTextString(m) }
func (*StreamDefinition) ProtoMessage()    {}
func (*StreamDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{3}
}

func (m *StreamDefinition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamDefinition.Unmarshal(m, b)
}
func (m *StreamDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamDefinition.Marshal(b, m, deterministic)
}
func (m *StreamDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamDefinition.Merge(m, src)
}
func (m *StreamDefinition) XXX_Size() int {
	return xxx_messageInfo_StreamDefinition.Size(m)
}
func (m *StreamDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_StreamDefinition proto.InternalMessageInfo

func (m *StreamDefinition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamDefinition) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *StreamDefinition) GetIsDelete() bool {
	if m != nil {
		return m.IsDelete
	}
	return false
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "stream.StreamRequest")
	proto.RegisterType((*StreamEvent)(nil), "stream.StreamEvent")
	proto.RegisterType((*Metadata)(nil), "stream.Metadata")
	proto.RegisterMapType((map[string]string)(nil), "stream.Metadata.KeyValueEntry")
	proto.RegisterType((*StreamDefinition)(nil), "stream.StreamDefinition")
}

func init() { proto.RegisterFile("stream.proto", fileDescriptor_bb17ef3f514bfe54) }

var fileDescriptor_bb17ef3f514bfe54 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x41, 0x4b, 0x02, 0x41,
	0x14, 0x66, 0x77, 0x4d, 0xd6, 0xa7, 0x96, 0xbc, 0x12, 0xc4, 0x43, 0xc8, 0x12, 0xe1, 0x21, 0x24,
	0x2c, 0x22, 0xec, 0xd2, 0x41, 0x0f, 0x21, 0x15, 0x4c, 0xd2, 0x31, 0x98, 0xea, 0x15, 0x83, 0xee,
	0x6a, 0xe3, 0x28, 0xec, 0x6f, 0xe8, 0x4f, 0xe7, 0xcc, 0xec, 0x4e, 0xec, 0xe2, 0xed, 0xbd, 0xef,
	0xfb, 0xe6, 0xcd, 0xf7, 0xbd, 0x19, 0x68, 0xac, 0x95, 0x24, 0x1e, 0x0f, 0x56, 0x72, 0xa9, 0x96,
	0x58, 0xb5, 0x5d, 0xf4, 0x00, 0xcd, 0x17, 0x53, 0x31, 0xfa, 0xd9, 0xd0, 0x5a, 0x21, 0x42, 0x25,
	0xe1, 0x31, 0x75, 0xbc, 0x9e, 0xd7, 0xaf, 0x31, 0x53, 0xe3, 0x19, 0x34, 0xa5, 0xa5, 0x49, 0x3e,
	0x69, 0xd2, 0x37, 0x64, 0x11, 0x8c, 0x3e, 0xa0, 0x6e, 0x47, 0x4d, 0xb6, 0x94, 0x28, 0x6c, 0x41,
	0x30, 0xa5, 0xd4, 0xcc, 0x69, 0x30, 0x5d, 0xe2, 0x09, 0x1c, 0xbc, 0xf2, 0xc5, 0xc6, 0x1e, 0x6f,
	0x30, 0xdb, 0xe0, 0x05, 0x84, 0x31, 0x29, 0xfe, 0xc9, 0x15, 0xef, 0x04, 0x3b, 0xa2, 0x3e, 0x6c,
	0x0d, 0x32, 0xab, 0x8f, 0x19, 0xce, 0x9c, 0x22, 0xfa, 0xf5, 0x21, 0xcc, 0x61, 0x3c, 0x87, 0x43,
	0x73, 0xd7, 0x4c, 0xc4, 0x3b, 0x1b, 0x3c, 0x5e, 0x99, 0xdb, 0x02, 0x56, 0x42, 0xf1, 0x1a, 0xda,
	0xcf, 0x52, 0x7c, 0x8b, 0xc4, 0xfa, 0xfb, 0x97, 0xfb, 0x46, 0xbe, 0x9f, 0xc4, 0x3e, 0x1c, 0x95,
	0xf5, 0x81, 0xd1, 0x97, 0x61, 0x1c, 0x41, 0x38, 0xa7, 0xd4, 0x66, 0xab, 0xf4, 0x82, 0x5d, 0x84,
	0xd3, 0x72, 0x84, 0xc1, 0x34, 0x13, 0x4c, 0x12, 0x25, 0x53, 0xe6, 0xf4, 0xdd, 0x3b, 0x68, 0x16,
	0x28, 0xbd, 0xb7, 0x79, 0xb6, 0xb7, 0x1a, 0xd3, 0xa5, 0xde, 0xdb, 0xd6, 0xed, 0xad, 0xc6, 0x6c,
	0x33, 0xf2, 0x6f, 0xbd, 0xe8, 0x0d, 0x5a, 0xd6, 0xcb, 0x98, 0xbe, 0x44, 0x22, 0x94, 0x58, 0x26,
	0x7b, 0x1f, 0xb0, 0x0b, 0xa1, 0x36, 0x31, 0x4b, 0x57, 0xf9, 0x10, 0xd7, 0x6b, 0x4e, 0xac, 0xc7,
	0xb4, 0x20, 0x45, 0x26, 0x5f, 0xc8, 0x5c, 0x3f, 0xbc, 0x87, 0xaa, 0x9d, 0x8f, 0x37, 0xae, 0x6a,
	0xe7, 0xd1, 0x0a, 0xff, 0xa6, 0x7b, 0x5c, 0x84, 0xcd, 0x0b, 0x5c, 0x7a, 0xef, 0x55, 0xf3, 0xdd,
	0xae, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x13, 0xf0, 0x9a, 0xc1, 0x7e, 0x02, 0x00, 0x00,
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
	stream, err := c.cc.NewStream(ctx, &_Stream_serviceDesc.Streams[0], "/stream.Stream/Stream", opts...)
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

// UnimplementedStreamServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServer struct {
}

func (*UnimplementedStreamServer) Stream(req *StreamRequest, srv Stream_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
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
	ServiceName: "stream.Stream",
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
