// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: broker.proto

package broker

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BrokerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Request string `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *BrokerRequest) Reset() {
	*x = BrokerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerRequest) ProtoMessage() {}

func (x *BrokerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerRequest.ProtoReflect.Descriptor instead.
func (*BrokerRequest) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{0}
}

func (x *BrokerRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *BrokerRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type BrokerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *BrokerResponse) Reset() {
	*x = BrokerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerResponse) ProtoMessage() {}

func (x *BrokerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerResponse.ProtoReflect.Descriptor instead.
func (*BrokerResponse) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{1}
}

func (x *BrokerResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_broker_proto protoreflect.FileDescriptor

var file_broker_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x57, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_broker_proto_rawDescOnce sync.Once
	file_broker_proto_rawDescData = file_broker_proto_rawDesc
)

func file_broker_proto_rawDescGZIP() []byte {
	file_broker_proto_rawDescOnce.Do(func() {
		file_broker_proto_rawDescData = protoimpl.X.CompressGZIP(file_broker_proto_rawDescData)
	})
	return file_broker_proto_rawDescData
}

var file_broker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_broker_proto_goTypes = []interface{}{
	(*BrokerRequest)(nil),  // 0: broker.BrokerRequest
	(*BrokerResponse)(nil), // 1: broker.BrokerResponse
}
var file_broker_proto_depIdxs = []int32{
	0, // 0: broker.BrokerService.RequestConnection:input_type -> broker.BrokerRequest
	1, // 1: broker.BrokerService.RequestConnection:output_type -> broker.BrokerResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_broker_proto_init() }
func file_broker_proto_init() {
	if File_broker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_broker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrokerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_broker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrokerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_broker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_broker_proto_goTypes,
		DependencyIndexes: file_broker_proto_depIdxs,
		MessageInfos:      file_broker_proto_msgTypes,
	}.Build()
	File_broker_proto = out.File
	file_broker_proto_rawDesc = nil
	file_broker_proto_goTypes = nil
	file_broker_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BrokerServiceClient is the client API for BrokerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BrokerServiceClient interface {
	RequestConnection(ctx context.Context, opts ...grpc.CallOption) (BrokerService_RequestConnectionClient, error)
}

type brokerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerServiceClient(cc grpc.ClientConnInterface) BrokerServiceClient {
	return &brokerServiceClient{cc}
}

func (c *brokerServiceClient) RequestConnection(ctx context.Context, opts ...grpc.CallOption) (BrokerService_RequestConnectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BrokerService_serviceDesc.Streams[0], "/broker.BrokerService/RequestConnection", opts...)
	if err != nil {
		return nil, err
	}
	x := &brokerServiceRequestConnectionClient{stream}
	return x, nil
}

type BrokerService_RequestConnectionClient interface {
	Send(*BrokerRequest) error
	Recv() (*BrokerResponse, error)
	grpc.ClientStream
}

type brokerServiceRequestConnectionClient struct {
	grpc.ClientStream
}

func (x *brokerServiceRequestConnectionClient) Send(m *BrokerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *brokerServiceRequestConnectionClient) Recv() (*BrokerResponse, error) {
	m := new(BrokerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BrokerServiceServer is the server API for BrokerService service.
type BrokerServiceServer interface {
	RequestConnection(BrokerService_RequestConnectionServer) error
}

// UnimplementedBrokerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBrokerServiceServer struct {
}

func (*UnimplementedBrokerServiceServer) RequestConnection(BrokerService_RequestConnectionServer) error {
	return status.Errorf(codes.Unimplemented, "method RequestConnection not implemented")
}

func RegisterBrokerServiceServer(s *grpc.Server, srv BrokerServiceServer) {
	s.RegisterService(&_BrokerService_serviceDesc, srv)
}

func _BrokerService_RequestConnection_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BrokerServiceServer).RequestConnection(&brokerServiceRequestConnectionServer{stream})
}

type BrokerService_RequestConnectionServer interface {
	Send(*BrokerResponse) error
	Recv() (*BrokerRequest, error)
	grpc.ServerStream
}

type brokerServiceRequestConnectionServer struct {
	grpc.ServerStream
}

func (x *brokerServiceRequestConnectionServer) Send(m *BrokerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *brokerServiceRequestConnectionServer) Recv() (*BrokerRequest, error) {
	m := new(BrokerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BrokerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "broker.BrokerService",
	HandlerType: (*BrokerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestConnection",
			Handler:       _BrokerService_RequestConnection_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "broker.proto",
}
