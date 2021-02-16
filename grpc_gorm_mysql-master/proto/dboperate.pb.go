// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: dboperate.proto

package proto

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

type InsDelUpdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price      float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	TypeId     int32   `protobuf:"varint,4,opt,name=typeId,proto3" json:"typeId,omitempty"`
	CreateTime int64   `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
}

func (x *InsDelUpdRequest) Reset() {
	*x = InsDelUpdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dboperate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsDelUpdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsDelUpdRequest) ProtoMessage() {}

func (x *InsDelUpdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dboperate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsDelUpdRequest.ProtoReflect.Descriptor instead.
func (*InsDelUpdRequest) Descriptor() ([]byte, []int) {
	return file_dboperate_proto_rawDescGZIP(), []int{0}
}

func (x *InsDelUpdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InsDelUpdRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InsDelUpdRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *InsDelUpdRequest) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *InsDelUpdRequest) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type SelectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Columns   string `protobuf:"bytes,1,opt,name=columns,proto3" json:"columns,omitempty"`
	Table     string `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	Condition string `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *SelectRequest) Reset() {
	*x = SelectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dboperate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectRequest) ProtoMessage() {}

func (x *SelectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dboperate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectRequest.ProtoReflect.Descriptor instead.
func (*SelectRequest) Descriptor() ([]byte, []int) {
	return file_dboperate_proto_rawDescGZIP(), []int{1}
}

func (x *SelectRequest) GetColumns() string {
	if x != nil {
		return x.Columns
	}
	return ""
}

func (x *SelectRequest) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *SelectRequest) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

type SqlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sql string `protobuf:"bytes,1,opt,name=sql,proto3" json:"sql,omitempty"`
}

func (x *SqlRequest) Reset() {
	*x = SqlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dboperate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqlRequest) ProtoMessage() {}

func (x *SqlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dboperate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqlRequest.ProtoReflect.Descriptor instead.
func (*SqlRequest) Descriptor() ([]byte, []int) {
	return file_dboperate_proto_rawDescGZIP(), []int{2}
}

func (x *SqlRequest) GetSql() string {
	if x != nil {
		return x.Sql
	}
	return ""
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dboperate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_dboperate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_dboperate_proto_rawDescGZIP(), []int{3}
}

func (x *Reply) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_dboperate_proto protoreflect.FileDescriptor

var file_dboperate_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x62, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x10, 0x69, 0x6e, 0x73,
	0x44, 0x65, 0x6c, 0x55, 0x70, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x5d, 0x0a, 0x0d, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1e,
	0x0a, 0x0a, 0x73, 0x71, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x71, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x71, 0x6c, 0x22, 0x1f,
	0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32,
	0xd2, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a,
	0x06, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x69, 0x6e, 0x73, 0x44, 0x65, 0x6c, 0x55, 0x70, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x31, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x73, 0x44, 0x65, 0x6c, 0x55, 0x70, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x73, 0x44, 0x65, 0x6c, 0x55, 0x70, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x07, 0x65, 0x78, 0x65, 0x63, 0x53, 0x71,
	0x6c, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x71, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dboperate_proto_rawDescOnce sync.Once
	file_dboperate_proto_rawDescData = file_dboperate_proto_rawDesc
)

func file_dboperate_proto_rawDescGZIP() []byte {
	file_dboperate_proto_rawDescOnce.Do(func() {
		file_dboperate_proto_rawDescData = protoimpl.X.CompressGZIP(file_dboperate_proto_rawDescData)
	})
	return file_dboperate_proto_rawDescData
}

var file_dboperate_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dboperate_proto_goTypes = []interface{}{
	(*InsDelUpdRequest)(nil), // 0: proto.insDelUpdRequest
	(*SelectRequest)(nil),    // 1: proto.selectRequest
	(*SqlRequest)(nil),       // 2: proto.sqlRequest
	(*Reply)(nil),            // 3: proto.reply
}
var file_dboperate_proto_depIdxs = []int32{
	0, // 0: proto.Operation.insert:input_type -> proto.insDelUpdRequest
	0, // 1: proto.Operation.delete:input_type -> proto.insDelUpdRequest
	0, // 2: proto.Operation.update:input_type -> proto.insDelUpdRequest
	2, // 3: proto.Operation.execSql:input_type -> proto.sqlRequest
	3, // 4: proto.Operation.insert:output_type -> proto.reply
	3, // 5: proto.Operation.delete:output_type -> proto.reply
	3, // 6: proto.Operation.update:output_type -> proto.reply
	3, // 7: proto.Operation.execSql:output_type -> proto.reply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dboperate_proto_init() }
func file_dboperate_proto_init() {
	if File_dboperate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dboperate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsDelUpdRequest); i {
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
		file_dboperate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectRequest); i {
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
		file_dboperate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqlRequest); i {
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
		file_dboperate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_dboperate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dboperate_proto_goTypes,
		DependencyIndexes: file_dboperate_proto_depIdxs,
		MessageInfos:      file_dboperate_proto_msgTypes,
	}.Build()
	File_dboperate_proto = out.File
	file_dboperate_proto_rawDesc = nil
	file_dboperate_proto_goTypes = nil
	file_dboperate_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OperationClient is the client API for Operation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperationClient interface {
	Insert(ctx context.Context, in *InsDelUpdRequest, opts ...grpc.CallOption) (*Reply, error)
	Delete(ctx context.Context, in *InsDelUpdRequest, opts ...grpc.CallOption) (*Reply, error)
	Update(ctx context.Context, in *InsDelUpdRequest, opts ...grpc.CallOption) (*Reply, error)
	ExecSql(ctx context.Context, in *SqlRequest, opts ...grpc.CallOption) (*Reply, error)
}

type operationClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationClient(cc grpc.ClientConnInterface) OperationClient {
	return &operationClient{cc}
}

func (c *operationClient) Insert(ctx context.Context, in *InsDelUpdRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.Operation/insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationClient) Delete(ctx context.Context, in *InsDelUpdRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.Operation/delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationClient) Update(ctx context.Context, in *InsDelUpdRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.Operation/update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationClient) ExecSql(ctx context.Context, in *SqlRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.Operation/execSql", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationServer is the server API for Operation service.
type OperationServer interface {
	Insert(context.Context, *InsDelUpdRequest) (*Reply, error)
	Delete(context.Context, *InsDelUpdRequest) (*Reply, error)
	Update(context.Context, *InsDelUpdRequest) (*Reply, error)
	ExecSql(context.Context, *SqlRequest) (*Reply, error)
}

// UnimplementedOperationServer can be embedded to have forward compatible implementations.
type UnimplementedOperationServer struct {
}

func (*UnimplementedOperationServer) Insert(context.Context, *InsDelUpdRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (*UnimplementedOperationServer) Delete(context.Context, *InsDelUpdRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedOperationServer) Update(context.Context, *InsDelUpdRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedOperationServer) ExecSql(context.Context, *SqlRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecSql not implemented")
}

func RegisterOperationServer(s *grpc.Server, srv OperationServer) {
	s.RegisterService(&_Operation_serviceDesc, srv)
}

func _Operation_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsDelUpdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Operation/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).Insert(ctx, req.(*InsDelUpdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operation_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsDelUpdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Operation/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).Delete(ctx, req.(*InsDelUpdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operation_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsDelUpdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Operation/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).Update(ctx, req.(*InsDelUpdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operation_ExecSql_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).ExecSql(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Operation/ExecSql",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).ExecSql(ctx, req.(*SqlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Operation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Operation",
	HandlerType: (*OperationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "insert",
			Handler:    _Operation_Insert_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Operation_Delete_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Operation_Update_Handler,
		},
		{
			MethodName: "execSql",
			Handler:    _Operation_ExecSql_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dboperate.proto",
}
