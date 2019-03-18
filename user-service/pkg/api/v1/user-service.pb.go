// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user-service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,5,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

type ReadRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{1}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{2}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type FindUserByPhoneRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,2,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindUserByPhoneRequest) Reset()         { *m = FindUserByPhoneRequest{} }
func (m *FindUserByPhoneRequest) String() string { return proto.CompactTextString(m) }
func (*FindUserByPhoneRequest) ProtoMessage()    {}
func (*FindUserByPhoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{3}
}

func (m *FindUserByPhoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindUserByPhoneRequest.Unmarshal(m, b)
}
func (m *FindUserByPhoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindUserByPhoneRequest.Marshal(b, m, deterministic)
}
func (m *FindUserByPhoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindUserByPhoneRequest.Merge(m, src)
}
func (m *FindUserByPhoneRequest) XXX_Size() int {
	return xxx_messageInfo_FindUserByPhoneRequest.Size(m)
}
func (m *FindUserByPhoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindUserByPhoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindUserByPhoneRequest proto.InternalMessageInfo

func (m *FindUserByPhoneRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *FindUserByPhoneRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

type FindUserByPhoneResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindUserByPhoneResponse) Reset()         { *m = FindUserByPhoneResponse{} }
func (m *FindUserByPhoneResponse) String() string { return proto.CompactTextString(m) }
func (*FindUserByPhoneResponse) ProtoMessage()    {}
func (*FindUserByPhoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3086c73a75cdba, []int{4}
}

func (m *FindUserByPhoneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindUserByPhoneResponse.Unmarshal(m, b)
}
func (m *FindUserByPhoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindUserByPhoneResponse.Marshal(b, m, deterministic)
}
func (m *FindUserByPhoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindUserByPhoneResponse.Merge(m, src)
}
func (m *FindUserByPhoneResponse) XXX_Size() int {
	return xxx_messageInfo_FindUserByPhoneResponse.Size(m)
}
func (m *FindUserByPhoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindUserByPhoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindUserByPhoneResponse proto.InternalMessageInfo

func (m *FindUserByPhoneResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *FindUserByPhoneResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "v1.User")
	proto.RegisterType((*ReadRequest)(nil), "v1.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "v1.ReadResponse")
	proto.RegisterType((*FindUserByPhoneRequest)(nil), "v1.FindUserByPhoneRequest")
	proto.RegisterType((*FindUserByPhoneResponse)(nil), "v1.FindUserByPhoneResponse")
}

func init() { proto.RegisterFile("user-service.proto", fileDescriptor_2a3086c73a75cdba) }

var fileDescriptor_2a3086c73a75cdba = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0xc7, 0xc9, 0x36, 0x7d, 0x68, 0x27, 0x0f, 0xb6, 0x0c, 0xa2, 0x21, 0xf6, 0x50, 0x72, 0xaa,
	0x07, 0x1b, 0x5a, 0xef, 0x1e, 0x3c, 0x08, 0x8a, 0x14, 0x59, 0xf1, 0x03, 0x6c, 0xcd, 0x50, 0x17,
	0x92, 0xdd, 0x98, 0x4d, 0x02, 0x5e, 0x3c, 0xf9, 0xc1, 0x65, 0x37, 0x55, 0x4b, 0x1b, 0x2f, 0xde,
	0x66, 0xfe, 0xf3, 0xf6, 0x9b, 0xd9, 0x05, 0xac, 0x0d, 0x95, 0x17, 0x86, 0xca, 0x46, 0x3e, 0xd3,
	0xbc, 0x28, 0x75, 0xa5, 0x91, 0x35, 0x8b, 0x68, 0xb2, 0xd1, 0x7a, 0x93, 0x51, 0x22, 0x0a, 0x99,
	0x08, 0xa5, 0x74, 0x25, 0x2a, 0xa9, 0x95, 0x69, 0x33, 0xe2, 0x77, 0xf0, 0x9f, 0x0c, 0x95, 0x78,
	0x04, 0x4c, 0xa6, 0xa1, 0x37, 0xf5, 0x66, 0x3d, 0xce, 0x64, 0x8a, 0x08, 0xbe, 0x12, 0x39, 0x85,
	0x6c, 0xea, 0xcd, 0x86, 0xdc, 0xd9, 0x18, 0xc1, 0x20, 0x13, 0xa6, 0x5a, 0x59, 0xbd, 0xe7, 0xf4,
	0x6f, 0x1f, 0x8f, 0xa1, 0x4f, 0xb9, 0x90, 0x59, 0xe8, 0xbb, 0x40, 0xeb, 0xe0, 0x14, 0x82, 0xe2,
	0x45, 0x2b, 0x5a, 0xd5, 0xf9, 0x9a, 0xca, 0xb0, 0xef, 0x62, 0xbb, 0x52, 0x9c, 0x40, 0xc0, 0x49,
	0xa4, 0x9c, 0x5e, 0x6b, 0x32, 0x15, 0x8e, 0xa1, 0x27, 0x0a, 0xe9, 0x38, 0x86, 0xdc, 0x9a, 0x5b,
	0x30, 0xf6, 0x05, 0x16, 0x5f, 0xc1, 0xff, 0xb6, 0xc0, 0x14, 0x5a, 0x19, 0xea, 0xa8, 0x98, 0x80,
	0x6f, 0x4f, 0xe1, 0x6a, 0x82, 0xe5, 0x60, 0xde, 0x2c, 0xe6, 0x76, 0x45, 0xee, 0xd4, 0xf8, 0x1e,
	0x4e, 0x6e, 0xa4, 0x4a, 0xad, 0x72, 0xfd, 0xf6, 0x60, 0x49, 0x7e, 0x9f, 0xbd, 0x87, 0xcf, 0x0e,
	0xf1, 0x6f, 0xe1, 0xf4, 0xa0, 0xdb, 0xdf, 0xc0, 0x96, 0x1f, 0x1e, 0x04, 0xd6, 0x7d, 0x6c, 0x5f,
	0x10, 0xcf, 0xc1, 0xb7, 0x8b, 0xe2, 0xc8, 0xe6, 0xed, 0xdc, 0x28, 0x1a, 0xff, 0x08, 0xdb, 0x51,
	0x77, 0x30, 0xda, 0xa3, 0xc0, 0xc8, 0x26, 0x75, 0x2f, 0x1a, 0x9d, 0x75, 0xc6, 0xda, 0x5e, 0xeb,
	0x7f, 0xee, 0x5f, 0x5c, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x86, 0x35, 0x35, 0x4f, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	FindUserByPhone(ctx context.Context, in *FindUserByPhoneRequest, opts ...grpc.CallOption) (*FindUserByPhoneResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FindUserByPhone(ctx context.Context, in *FindUserByPhoneRequest, opts ...grpc.CallOption) (*FindUserByPhoneResponse, error) {
	out := new(FindUserByPhoneResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/FindUserByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	FindUserByPhone(context.Context, *FindUserByPhoneRequest) (*FindUserByPhoneResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FindUserByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserByPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindUserByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/FindUserByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindUserByPhone(ctx, req.(*FindUserByPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _UserService_Read_Handler,
		},
		{
			MethodName: "FindUserByPhone",
			Handler:    _UserService_FindUserByPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user-service.proto",
}