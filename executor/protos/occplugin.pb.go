// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/OccPlugin.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	protos/OccPlugin.proto

It has these top-level messages:
	StateStreamRequest
	StateStreamReply
	GetStateRequest
	GetStateReply
	ConfigEntry
	TransitionRequest
	TransitionReply
*/
package pb

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

type StateChangeTrigger int32

const (
	StateChangeTrigger_EXECUTOR           StateChangeTrigger = 0
	StateChangeTrigger_DEVICE_INTENTIONAL StateChangeTrigger = 1
	StateChangeTrigger_DEVICE_ERROR       StateChangeTrigger = 2
)

var StateChangeTrigger_name = map[int32]string{
	0: "EXECUTOR",
	1: "DEVICE_INTENTIONAL",
	2: "DEVICE_ERROR",
}
var StateChangeTrigger_value = map[string]int32{
	"EXECUTOR":           0,
	"DEVICE_INTENTIONAL": 1,
	"DEVICE_ERROR":       2,
}

func (x StateChangeTrigger) String() string {
	return proto.EnumName(StateChangeTrigger_name, int32(x))
}
func (StateChangeTrigger) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StateType int32

const (
	StateType_STATE_STABLE       StateType = 0
	StateType_STATE_INTERMEDIATE StateType = 1
)

var StateType_name = map[int32]string{
	0: "STATE_STABLE",
	1: "STATE_INTERMEDIATE",
}
var StateType_value = map[string]int32{
	"STATE_STABLE":       0,
	"STATE_INTERMEDIATE": 1,
}

func (x StateType) String() string {
	return proto.EnumName(StateType_name, int32(x))
}
func (StateType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type StateStreamRequest struct {
}

func (m *StateStreamRequest) Reset()                    { *m = StateStreamRequest{} }
func (m *StateStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*StateStreamRequest) ProtoMessage()               {}
func (*StateStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StateStreamReply struct {
	Type  StateType `protobuf:"varint,1,opt,name=type,enum=occplugin_pb.StateType" json:"type,omitempty"`
	State string    `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
}

func (m *StateStreamReply) Reset()                    { *m = StateStreamReply{} }
func (m *StateStreamReply) String() string            { return proto.CompactTextString(m) }
func (*StateStreamReply) ProtoMessage()               {}
func (*StateStreamReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StateStreamReply) GetType() StateType {
	if m != nil {
		return m.Type
	}
	return StateType_STATE_STABLE
}

func (m *StateStreamReply) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type GetStateRequest struct {
}

func (m *GetStateRequest) Reset()                    { *m = GetStateRequest{} }
func (m *GetStateRequest) String() string            { return proto.CompactTextString(m) }
func (*GetStateRequest) ProtoMessage()               {}
func (*GetStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GetStateReply struct {
	State string `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
}

func (m *GetStateReply) Reset()                    { *m = GetStateReply{} }
func (m *GetStateReply) String() string            { return proto.CompactTextString(m) }
func (*GetStateReply) ProtoMessage()               {}
func (*GetStateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetStateReply) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type ConfigEntry struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *ConfigEntry) Reset()                    { *m = ConfigEntry{} }
func (m *ConfigEntry) String() string            { return proto.CompactTextString(m) }
func (*ConfigEntry) ProtoMessage()               {}
func (*ConfigEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ConfigEntry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ConfigEntry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type TransitionRequest struct {
	SrcState  string         `protobuf:"bytes,1,opt,name=srcState" json:"srcState,omitempty"`
	Event     string         `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
	Arguments []*ConfigEntry `protobuf:"bytes,3,rep,name=arguments" json:"arguments,omitempty"`
}

func (m *TransitionRequest) Reset()                    { *m = TransitionRequest{} }
func (m *TransitionRequest) String() string            { return proto.CompactTextString(m) }
func (*TransitionRequest) ProtoMessage()               {}
func (*TransitionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TransitionRequest) GetSrcState() string {
	if m != nil {
		return m.SrcState
	}
	return ""
}

func (m *TransitionRequest) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *TransitionRequest) GetArguments() []*ConfigEntry {
	if m != nil {
		return m.Arguments
	}
	return nil
}

type TransitionReply struct {
	Trigger StateChangeTrigger `protobuf:"varint,1,opt,name=trigger,enum=occplugin_pb.StateChangeTrigger" json:"trigger,omitempty"`
	State   string             `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	Event   string             `protobuf:"bytes,3,opt,name=event" json:"event,omitempty"`
	Ok      bool               `protobuf:"varint,4,opt,name=ok" json:"ok,omitempty"`
}

func (m *TransitionReply) Reset()                    { *m = TransitionReply{} }
func (m *TransitionReply) String() string            { return proto.CompactTextString(m) }
func (*TransitionReply) ProtoMessage()               {}
func (*TransitionReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TransitionReply) GetTrigger() StateChangeTrigger {
	if m != nil {
		return m.Trigger
	}
	return StateChangeTrigger_EXECUTOR
}

func (m *TransitionReply) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *TransitionReply) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *TransitionReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*StateStreamRequest)(nil), "occplugin_pb.StateStreamRequest")
	proto.RegisterType((*StateStreamReply)(nil), "occplugin_pb.StateStreamReply")
	proto.RegisterType((*GetStateRequest)(nil), "occplugin_pb.GetStateRequest")
	proto.RegisterType((*GetStateReply)(nil), "occplugin_pb.GetStateReply")
	proto.RegisterType((*ConfigEntry)(nil), "occplugin_pb.ConfigEntry")
	proto.RegisterType((*TransitionRequest)(nil), "occplugin_pb.TransitionRequest")
	proto.RegisterType((*TransitionReply)(nil), "occplugin_pb.TransitionReply")
	proto.RegisterEnum("occplugin_pb.StateChangeTrigger", StateChangeTrigger_name, StateChangeTrigger_value)
	proto.RegisterEnum("occplugin_pb.StateType", StateType_name, StateType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for OccPlugin service

type OccPluginClient interface {
	// We have to have a notification stream because the FairMQDevice might transition
	// on its own for whatever reason.
	StateStream(ctx context.Context, in *StateStreamRequest, opts ...grpc.CallOption) (OccPlugin_StateStreamClient, error)
	GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateReply, error)
	Transition(ctx context.Context, in *TransitionRequest, opts ...grpc.CallOption) (*TransitionReply, error)
}

type occPluginClient struct {
	cc *grpc.ClientConn
}

func NewOccPluginClient(cc *grpc.ClientConn) OccPluginClient {
	return &occPluginClient{cc}
}

func (c *occPluginClient) StateStream(ctx context.Context, in *StateStreamRequest, opts ...grpc.CallOption) (OccPlugin_StateStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_OccPlugin_serviceDesc.Streams[0], c.cc, "/occplugin_pb.OccPlugin/StateStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &occPluginStateStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OccPlugin_StateStreamClient interface {
	Recv() (*StateStreamReply, error)
	grpc.ClientStream
}

type occPluginStateStreamClient struct {
	grpc.ClientStream
}

func (x *occPluginStateStreamClient) Recv() (*StateStreamReply, error) {
	m := new(StateStreamReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *occPluginClient) GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateReply, error) {
	out := new(GetStateReply)
	err := grpc.Invoke(ctx, "/occplugin_pb.OccPlugin/GetState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *occPluginClient) Transition(ctx context.Context, in *TransitionRequest, opts ...grpc.CallOption) (*TransitionReply, error) {
	out := new(TransitionReply)
	err := grpc.Invoke(ctx, "/occplugin_pb.OccPlugin/Transition", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OccPlugin service

type OccPluginServer interface {
	// We have to have a notification stream because the FairMQDevice might transition
	// on its own for whatever reason.
	StateStream(*StateStreamRequest, OccPlugin_StateStreamServer) error
	GetState(context.Context, *GetStateRequest) (*GetStateReply, error)
	Transition(context.Context, *TransitionRequest) (*TransitionReply, error)
}

func RegisterOccPluginServer(s *grpc.Server, srv OccPluginServer) {
	s.RegisterService(&_OccPlugin_serviceDesc, srv)
}

func _OccPlugin_StateStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StateStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OccPluginServer).StateStream(m, &occPluginStateStreamServer{stream})
}

type OccPlugin_StateStreamServer interface {
	Send(*StateStreamReply) error
	grpc.ServerStream
}

type occPluginStateStreamServer struct {
	grpc.ServerStream
}

func (x *occPluginStateStreamServer) Send(m *StateStreamReply) error {
	return x.ServerStream.SendMsg(m)
}

func _OccPlugin_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OccPluginServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/occplugin_pb.OccPlugin/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OccPluginServer).GetState(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OccPlugin_Transition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OccPluginServer).Transition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/occplugin_pb.OccPlugin/Transition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OccPluginServer).Transition(ctx, req.(*TransitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OccPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "occplugin_pb.OccPlugin",
	HandlerType: (*OccPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetState",
			Handler:    _OccPlugin_GetState_Handler,
		},
		{
			MethodName: "Transition",
			Handler:    _OccPlugin_Transition_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StateStream",
			Handler:       _OccPlugin_StateStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/OccPlugin.proto",
}

func init() { proto.RegisterFile("protos/OccPlugin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0xe7, 0x74, 0x40, 0x7b, 0x2d, 0x5b, 0x66, 0x4d, 0x5b, 0x29, 0x2a, 0x44, 0x91, 0x90,
	0xaa, 0x21, 0x05, 0x54, 0x34, 0x21, 0xf1, 0xd6, 0x75, 0x16, 0x54, 0x2a, 0x2d, 0x72, 0x3c, 0x84,
	0x78, 0xa9, 0xd2, 0xc8, 0x64, 0x51, 0xb3, 0x24, 0x38, 0xee, 0xa4, 0xbc, 0xf0, 0xce, 0xbf, 0xcc,
	0x13, 0x8a, 0x9b, 0xe6, 0x07, 0xb4, 0x7b, 0xcb, 0x7d, 0xfd, 0xb9, 0xef, 0x9d, 0xcf, 0x17, 0x38,
	0x8b, 0x45, 0x24, 0xa3, 0xe4, 0xcd, 0xdc, 0x75, 0xbf, 0x04, 0x6b, 0xcf, 0x0f, 0x2d, 0x25, 0xe0,
	0x4e, 0xe4, 0xba, 0xb1, 0x12, 0x16, 0xf1, 0xd2, 0x3c, 0x05, 0x6c, 0x4b, 0x47, 0x72, 0x5b, 0x0a,
	0xee, 0xdc, 0x51, 0xfe, 0x73, 0xcd, 0x13, 0x69, 0xde, 0x80, 0x5e, 0x53, 0xe3, 0x20, 0xc5, 0xaf,
	0xe1, 0x50, 0xa6, 0x31, 0xef, 0x22, 0x03, 0x0d, 0x8e, 0x86, 0xe7, 0x56, 0xd5, 0xc6, 0x52, 0x34,
	0x4b, 0x63, 0x4e, 0x15, 0x84, 0x4f, 0xe1, 0x51, 0x92, 0x49, 0x5d, 0xcd, 0x40, 0x83, 0x16, 0xdd,
	0x04, 0xe6, 0x09, 0x1c, 0x7f, 0xe4, 0x52, 0xb1, 0xdb, 0x4a, 0xaf, 0xe0, 0x69, 0x29, 0x65, 0x65,
	0x8a, 0x4c, 0x54, 0xcd, 0xbc, 0x84, 0xf6, 0x38, 0x0a, 0x7f, 0xf8, 0x1e, 0x09, 0xa5, 0x48, 0xb1,
	0x0e, 0x8d, 0x15, 0x4f, 0x73, 0x24, 0xfb, 0xcc, 0xd2, 0xee, 0x9d, 0x60, 0x5d, 0x14, 0x54, 0x81,
	0xf9, 0x0b, 0x4e, 0x98, 0x70, 0xc2, 0xc4, 0x97, 0x7e, 0x14, 0xe6, 0x25, 0x71, 0x0f, 0x9a, 0x89,
	0x70, 0xed, 0x4a, 0x91, 0x22, 0xce, 0x6c, 0xf8, 0x3d, 0x0f, 0xe5, 0xd6, 0x46, 0x05, 0xf8, 0x3d,
	0xb4, 0x1c, 0xe1, 0xad, 0xef, 0x78, 0x28, 0x93, 0x6e, 0xc3, 0x68, 0x0c, 0xda, 0xc3, 0x67, 0xf5,
	0xfb, 0x57, 0x9a, 0xa3, 0x25, 0x6b, 0xfe, 0x46, 0x70, 0x5c, 0x6d, 0x20, 0xbb, 0xe0, 0x07, 0x78,
	0x22, 0x85, 0xef, 0x79, 0x5c, 0xe4, 0xa3, 0x34, 0x76, 0x8c, 0x72, 0x7c, 0xeb, 0x84, 0x1e, 0x67,
	0x1b, 0x8e, 0x6e, 0x13, 0x76, 0x8f, 0xb5, 0x6c, 0xba, 0x51, 0x6d, 0xfa, 0x08, 0xb4, 0x68, 0xd5,
	0x3d, 0x34, 0xd0, 0xa0, 0x49, 0xb5, 0x68, 0x75, 0x31, 0xcd, 0x5f, 0xba, 0x66, 0x8d, 0x3b, 0xd0,
	0x24, 0xdf, 0xc8, 0xf8, 0x86, 0xcd, 0xa9, 0x7e, 0x80, 0xcf, 0x00, 0x5f, 0x93, 0xaf, 0x93, 0x31,
	0x59, 0x4c, 0x66, 0x8c, 0xcc, 0xd8, 0x64, 0x3e, 0x1b, 0x4d, 0x75, 0x84, 0x75, 0xe8, 0xe4, 0x3a,
	0xa1, 0x74, 0x4e, 0x75, 0xed, 0xe2, 0x12, 0x5a, 0xc5, 0x9b, 0x67, 0xc7, 0x36, 0x1b, 0x31, 0xb2,
	0xb0, 0xd9, 0xe8, 0x6a, 0x4a, 0x36, 0x46, 0x1b, 0x25, 0xf3, 0xa1, 0x9f, 0xc9, 0xf5, 0x64, 0xc4,
	0x88, 0x8e, 0x86, 0x7f, 0x10, 0xb4, 0x8a, 0x85, 0xc4, 0x36, 0xb4, 0x2b, 0x6b, 0x86, 0x77, 0x0d,
	0xa2, 0xb6, 0x97, 0xbd, 0x17, 0x0f, 0x10, 0x71, 0x90, 0x9a, 0x07, 0x6f, 0x11, 0xfe, 0x04, 0xcd,
	0xed, 0x46, 0xe1, 0x7e, 0x9d, 0xff, 0x67, 0xf9, 0x7a, 0xcf, 0xf7, 0x1d, 0x2b, 0x2f, 0x3c, 0x03,
	0x28, 0x1f, 0x0f, 0xbf, 0xac, 0xc3, 0xff, 0xed, 0x55, 0xaf, 0xbf, 0x1f, 0x50, 0x7e, 0x57, 0x7d,
	0x38, 0x77, 0x6f, 0x2d, 0x97, 0x8b, 0xd0, 0x72, 0x02, 0xdf, 0xe5, 0x25, 0xff, 0x5d, 0x8b, 0x97,
	0xcb, 0xc7, 0xea, 0xff, 0x7c, 0xf7, 0x37, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x3b, 0x0b, 0x6c, 0xb9,
	0x03, 0x00, 0x00,
}