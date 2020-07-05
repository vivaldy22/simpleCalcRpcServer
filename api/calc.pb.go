// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: api/calc.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ParameterMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int32 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int32 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *ParameterMessage) Reset() {
	*x = ParameterMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParameterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParameterMessage) ProtoMessage() {}

func (x *ParameterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParameterMessage.ProtoReflect.Descriptor instead.
func (*ParameterMessage) Descriptor() ([]byte, []int) {
	return file_api_calc_proto_rawDescGZIP(), []int{0}
}

func (x *ParameterMessage) GetNum1() int32 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *ParameterMessage) GetNum2() int32 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	string code = 1;
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_api_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_api_calc_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResultMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResNum int32  `protobuf:"varint,1,opt,name=resNum,proto3" json:"resNum,omitempty"`
	Error  *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ResultMessage) Reset() {
	*x = ResultMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultMessage) ProtoMessage() {}

func (x *ResultMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultMessage.ProtoReflect.Descriptor instead.
func (*ResultMessage) Descriptor() ([]byte, []int) {
	return file_api_calc_proto_rawDescGZIP(), []int{2}
}

func (x *ResultMessage) GetResNum() int32 {
	if x != nil {
		return x.ResNum
	}
	return 0
}

func (x *ResultMessage) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_api_calc_proto protoreflect.FileDescriptor

var file_api_calc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x3a, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d,
	0x32, 0x22, 0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x49, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x20, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32,
	0xee, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x63, 0x12, 0x38, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09,
	0x4d, 0x75, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09, 0x44, 0x69, 0x76, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_calc_proto_rawDescOnce sync.Once
	file_api_calc_proto_rawDescData = file_api_calc_proto_rawDesc
)

func file_api_calc_proto_rawDescGZIP() []byte {
	file_api_calc_proto_rawDescOnce.Do(func() {
		file_api_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_calc_proto_rawDescData)
	})
	return file_api_calc_proto_rawDescData
}

var file_api_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_calc_proto_goTypes = []interface{}{
	(*ParameterMessage)(nil), // 0: api.ParameterMessage
	(*Error)(nil),            // 1: api.Error
	(*ResultMessage)(nil),    // 2: api.ResultMessage
}
var file_api_calc_proto_depIdxs = []int32{
	1, // 0: api.ResultMessage.error:type_name -> api.Error
	0, // 1: api.Calc.AddNumber:input_type -> api.ParameterMessage
	0, // 2: api.Calc.SubNumber:input_type -> api.ParameterMessage
	0, // 3: api.Calc.MulNumber:input_type -> api.ParameterMessage
	0, // 4: api.Calc.DivNumber:input_type -> api.ParameterMessage
	2, // 5: api.Calc.AddNumber:output_type -> api.ResultMessage
	2, // 6: api.Calc.SubNumber:output_type -> api.ResultMessage
	2, // 7: api.Calc.MulNumber:output_type -> api.ResultMessage
	2, // 8: api.Calc.DivNumber:output_type -> api.ResultMessage
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_calc_proto_init() }
func file_api_calc_proto_init() {
	if File_api_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParameterMessage); i {
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
		file_api_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_api_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultMessage); i {
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
			RawDescriptor: file_api_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_calc_proto_goTypes,
		DependencyIndexes: file_api_calc_proto_depIdxs,
		MessageInfos:      file_api_calc_proto_msgTypes,
	}.Build()
	File_api_calc_proto = out.File
	file_api_calc_proto_rawDesc = nil
	file_api_calc_proto_goTypes = nil
	file_api_calc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalcClient is the client API for Calc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcClient interface {
	AddNumber(ctx context.Context, in *ParameterMessage, opts ...grpc.CallOption) (*ResultMessage, error)
	SubNumber(ctx context.Context, in *ParameterMessage, opts ...grpc.CallOption) (*ResultMessage, error)
	MulNumber(ctx context.Context, in *ParameterMessage, opts ...grpc.CallOption) (*ResultMessage, error)
	DivNumber(ctx context.Context, in *ParameterMessage, opts ...grpc.CallOption) (*ResultMessage, error)
}

type calcClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcClient(cc grpc.ClientConnInterface) CalcClient {
	return &calcClient{cc}
}

func (c *calcClient) AddNumber(ctx context.Context, in *ParameterMessage, opts ...grpc.CallOption) (*ResultMessage, error) {
	out := new(ResultMessage)
	err := c.cc.Invoke(ctx, "/api.Calc/AddNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcClient) SubNumber(ctx context.Context, in *ParameterMessage, opts ...grpc.CallOption) (*ResultMessage, error) {
	out := new(ResultMessage)
	err := c.cc.Invoke(ctx, "/api.Calc/SubNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcClient) MulNumber(ctx context.Context, in *ParameterMessage, opts ...grpc.CallOption) (*ResultMessage, error) {
	out := new(ResultMessage)
	err := c.cc.Invoke(ctx, "/api.Calc/MulNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcClient) DivNumber(ctx context.Context, in *ParameterMessage, opts ...grpc.CallOption) (*ResultMessage, error) {
	out := new(ResultMessage)
	err := c.cc.Invoke(ctx, "/api.Calc/DivNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServer is the server API for Calc service.
type CalcServer interface {
	AddNumber(context.Context, *ParameterMessage) (*ResultMessage, error)
	SubNumber(context.Context, *ParameterMessage) (*ResultMessage, error)
	MulNumber(context.Context, *ParameterMessage) (*ResultMessage, error)
	DivNumber(context.Context, *ParameterMessage) (*ResultMessage, error)
}

// UnimplementedCalcServer can be embedded to have forward compatible implementations.
type UnimplementedCalcServer struct {
}

func (*UnimplementedCalcServer) AddNumber(context.Context, *ParameterMessage) (*ResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNumber not implemented")
}
func (*UnimplementedCalcServer) SubNumber(context.Context, *ParameterMessage) (*ResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubNumber not implemented")
}
func (*UnimplementedCalcServer) MulNumber(context.Context, *ParameterMessage) (*ResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MulNumber not implemented")
}
func (*UnimplementedCalcServer) DivNumber(context.Context, *ParameterMessage) (*ResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DivNumber not implemented")
}

func RegisterCalcServer(s *grpc.Server, srv CalcServer) {
	s.RegisterService(&_Calc_serviceDesc, srv)
}

func _Calc_AddNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParameterMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).AddNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Calc/AddNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).AddNumber(ctx, req.(*ParameterMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calc_SubNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParameterMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).SubNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Calc/SubNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).SubNumber(ctx, req.(*ParameterMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calc_MulNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParameterMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).MulNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Calc/MulNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).MulNumber(ctx, req.(*ParameterMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calc_DivNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParameterMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).DivNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Calc/DivNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).DivNumber(ctx, req.(*ParameterMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Calc",
	HandlerType: (*CalcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNumber",
			Handler:    _Calc_AddNumber_Handler,
		},
		{
			MethodName: "SubNumber",
			Handler:    _Calc_SubNumber_Handler,
		},
		{
			MethodName: "MulNumber",
			Handler:    _Calc_MulNumber_Handler,
		},
		{
			MethodName: "DivNumber",
			Handler:    _Calc_DivNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/calc.proto",
}
