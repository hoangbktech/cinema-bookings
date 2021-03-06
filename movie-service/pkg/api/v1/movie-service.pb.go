// Code generated by protoc-gen-go. DO NOT EDIT.
// source: movie-service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

// Movie
type Movie struct {
	// Unique integer identifier of the movie
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Title of the movie
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Detail description of the movie
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Movie) Reset()         { *m = Movie{} }
func (m *Movie) String() string { return proto.CompactTextString(m) }
func (*Movie) ProtoMessage()    {}
func (*Movie) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eea73d65d8c40d8, []int{0}
}

func (m *Movie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Movie.Unmarshal(m, b)
}
func (m *Movie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Movie.Marshal(b, m, deterministic)
}
func (m *Movie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Movie.Merge(m, src)
}
func (m *Movie) XXX_Size() int {
	return xxx_messageInfo_Movie.Size(m)
}
func (m *Movie) XXX_DiscardUnknown() {
	xxx_messageInfo_Movie.DiscardUnknown(m)
}

var xxx_messageInfo_Movie proto.InternalMessageInfo

func (m *Movie) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Movie) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Movie) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ReadMovieRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Unique integer identifier of the movie
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadMovieRequest) Reset()         { *m = ReadMovieRequest{} }
func (m *ReadMovieRequest) String() string { return proto.CompactTextString(m) }
func (*ReadMovieRequest) ProtoMessage()    {}
func (*ReadMovieRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eea73d65d8c40d8, []int{1}
}

func (m *ReadMovieRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadMovieRequest.Unmarshal(m, b)
}
func (m *ReadMovieRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadMovieRequest.Marshal(b, m, deterministic)
}
func (m *ReadMovieRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadMovieRequest.Merge(m, src)
}
func (m *ReadMovieRequest) XXX_Size() int {
	return xxx_messageInfo_ReadMovieRequest.Size(m)
}
func (m *ReadMovieRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadMovieRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadMovieRequest proto.InternalMessageInfo

func (m *ReadMovieRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadMovieRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadMovieResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Movie entity read by ID
	Movie                *Movie   `protobuf:"bytes,2,opt,name=movie,proto3" json:"movie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadMovieResponse) Reset()         { *m = ReadMovieResponse{} }
func (m *ReadMovieResponse) String() string { return proto.CompactTextString(m) }
func (*ReadMovieResponse) ProtoMessage()    {}
func (*ReadMovieResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eea73d65d8c40d8, []int{2}
}

func (m *ReadMovieResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadMovieResponse.Unmarshal(m, b)
}
func (m *ReadMovieResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadMovieResponse.Marshal(b, m, deterministic)
}
func (m *ReadMovieResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadMovieResponse.Merge(m, src)
}
func (m *ReadMovieResponse) XXX_Size() int {
	return xxx_messageInfo_ReadMovieResponse.Size(m)
}
func (m *ReadMovieResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadMovieResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadMovieResponse proto.InternalMessageInfo

func (m *ReadMovieResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadMovieResponse) GetMovie() *Movie {
	if m != nil {
		return m.Movie
	}
	return nil
}

func init() {
	proto.RegisterType((*Movie)(nil), "v1.Movie")
	proto.RegisterType((*ReadMovieRequest)(nil), "v1.ReadMovieRequest")
	proto.RegisterType((*ReadMovieResponse)(nil), "v1.ReadMovieResponse")
}

func init() { proto.RegisterFile("movie-service.proto", fileDescriptor_6eea73d65d8c40d8) }

var fileDescriptor_6eea73d65d8c40d8 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0x69, 0x6a, 0x85, 0xde, 0x89, 0xcc, 0x38, 0xa1, 0x14, 0x61, 0xa5, 0x4f, 0x7b, 0xb1,
	0x65, 0x9b, 0xff, 0x40, 0xf0, 0x4d, 0x84, 0xf8, 0x0b, 0xb2, 0xf5, 0x3a, 0x2e, 0xac, 0x4d, 0x6c,
	0xb2, 0xfe, 0x7e, 0xc9, 0xcd, 0x84, 0x31, 0x7c, 0xcb, 0x3d, 0xf7, 0x9e, 0xef, 0x84, 0x03, 0x8f,
	0xbd, 0x99, 0x08, 0x5f, 0x1c, 0x8e, 0x13, 0xed, 0xb1, 0xb1, 0xa3, 0xf1, 0x46, 0x8a, 0x69, 0x5d,
	0x2e, 0x0f, 0xc6, 0x1c, 0x8e, 0xd8, 0xb2, 0xb2, 0x3b, 0x7d, 0xb7, 0x9e, 0x7a, 0x74, 0x5e, 0xf7,
	0x36, 0x1e, 0x95, 0xcf, 0xe7, 0x03, 0x6d, 0xa9, 0xd5, 0xc3, 0x60, 0xbc, 0xf6, 0x64, 0x06, 0x17,
	0xb7, 0xf5, 0x27, 0x64, 0x1f, 0x81, 0x2c, 0xef, 0x41, 0x50, 0x57, 0x24, 0x55, 0xb2, 0x4a, 0x95,
	0xa0, 0x4e, 0x2e, 0x20, 0xf3, 0xe4, 0x8f, 0x58, 0x88, 0x2a, 0x59, 0xe5, 0x2a, 0x0e, 0xb2, 0x82,
	0x59, 0x87, 0x6e, 0x3f, 0x92, 0x0d, 0x90, 0x22, 0xe5, 0xdd, 0xa5, 0x54, 0xbf, 0xc2, 0x5c, 0xa1,
	0xee, 0x18, 0xaa, 0xf0, 0xe7, 0x84, 0xce, 0xcb, 0x39, 0xa4, 0xda, 0x12, 0xc3, 0x73, 0x15, 0x9e,
	0xe7, 0x34, 0xf1, 0x97, 0x56, 0xbf, 0xc3, 0xc3, 0x85, 0xcb, 0x59, 0x33, 0x38, 0xfc, 0xc7, 0xb6,
	0x84, 0x8c, 0x7b, 0x60, 0xe7, 0x6c, 0x93, 0x37, 0xd3, 0xba, 0x89, 0x9e, 0xa8, 0x6f, 0xde, 0xe0,
	0x8e, 0xe7, 0xaf, 0xd8, 0x93, 0xdc, 0xc2, 0x4d, 0xe0, 0xca, 0x45, 0xb8, 0xbc, 0xfe, 0x57, 0xf9,
	0x74, 0xa5, 0xc6, 0xdc, 0xdd, 0x2d, 0x57, 0xb3, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x05, 0xa2,
	0x5f, 0xf1, 0x74, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MovieServiceClient is the client API for MovieService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MovieServiceClient interface {
	// Read movie
	Read(ctx context.Context, in *ReadMovieRequest, opts ...grpc.CallOption) (*ReadMovieResponse, error)
}

type movieServiceClient struct {
	cc *grpc.ClientConn
}

func NewMovieServiceClient(cc *grpc.ClientConn) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) Read(ctx context.Context, in *ReadMovieRequest, opts ...grpc.CallOption) (*ReadMovieResponse, error) {
	out := new(ReadMovieResponse)
	err := c.cc.Invoke(ctx, "/v1.MovieService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServiceServer is the server API for MovieService service.
type MovieServiceServer interface {
	// Read movie
	Read(context.Context, *ReadMovieRequest) (*ReadMovieResponse, error)
}

func RegisterMovieServiceServer(s *grpc.Server, srv MovieServiceServer) {
	s.RegisterService(&_MovieService_serviceDesc, srv)
}

func _MovieService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MovieService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).Read(ctx, req.(*ReadMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.MovieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _MovieService_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie-service.proto",
}
