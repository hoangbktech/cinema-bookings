// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification-service.proto

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

type EmailRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	RecipientIds         []string `protobuf:"bytes,2,rep,name=recipientIds,proto3" json:"recipientIds,omitempty"`
	Subject              string   `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailRequest) Reset()         { *m = EmailRequest{} }
func (m *EmailRequest) String() string { return proto.CompactTextString(m) }
func (*EmailRequest) ProtoMessage()    {}
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62e234be764a7e86, []int{0}
}

func (m *EmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailRequest.Unmarshal(m, b)
}
func (m *EmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailRequest.Marshal(b, m, deterministic)
}
func (m *EmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailRequest.Merge(m, src)
}
func (m *EmailRequest) XXX_Size() int {
	return xxx_messageInfo_EmailRequest.Size(m)
}
func (m *EmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailRequest proto.InternalMessageInfo

func (m *EmailRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *EmailRequest) GetRecipientIds() []string {
	if m != nil {
		return m.RecipientIds
	}
	return nil
}

func (m *EmailRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *EmailRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type SMSRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,2,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SMSRequest) Reset()         { *m = SMSRequest{} }
func (m *SMSRequest) String() string { return proto.CompactTextString(m) }
func (*SMSRequest) ProtoMessage()    {}
func (*SMSRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62e234be764a7e86, []int{1}
}

func (m *SMSRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SMSRequest.Unmarshal(m, b)
}
func (m *SMSRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SMSRequest.Marshal(b, m, deterministic)
}
func (m *SMSRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SMSRequest.Merge(m, src)
}
func (m *SMSRequest) XXX_Size() int {
	return xxx_messageInfo_SMSRequest.Size(m)
}
func (m *SMSRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SMSRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SMSRequest proto.InternalMessageInfo

func (m *SMSRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *SMSRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *SMSRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type NotificationResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationResponse) Reset()         { *m = NotificationResponse{} }
func (m *NotificationResponse) String() string { return proto.CompactTextString(m) }
func (*NotificationResponse) ProtoMessage()    {}
func (*NotificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62e234be764a7e86, []int{2}
}

func (m *NotificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationResponse.Unmarshal(m, b)
}
func (m *NotificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationResponse.Marshal(b, m, deterministic)
}
func (m *NotificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationResponse.Merge(m, src)
}
func (m *NotificationResponse) XXX_Size() int {
	return xxx_messageInfo_NotificationResponse.Size(m)
}
func (m *NotificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationResponse proto.InternalMessageInfo

func (m *NotificationResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *NotificationResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*EmailRequest)(nil), "v1.EmailRequest")
	proto.RegisterType((*SMSRequest)(nil), "v1.SMSRequest")
	proto.RegisterType((*NotificationResponse)(nil), "v1.NotificationResponse")
}

func init() { proto.RegisterFile("notification-service.proto", fileDescriptor_62e234be764a7e86) }

var fileDescriptor_62e234be764a7e86 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xd5, 0xa4, 0x2a, 0xea, 0xd1, 0x01, 0x99, 0x3f, 0xb2, 0x2a, 0x86, 0x28, 0x53, 0x17,
	0xa8, 0x80, 0x2f, 0xc0, 0xc2, 0xc0, 0x40, 0x87, 0x64, 0x63, 0x73, 0xdc, 0xa3, 0x18, 0xb5, 0x3e,
	0xe3, 0xbb, 0x54, 0xe2, 0xdb, 0xa3, 0x3a, 0x01, 0x05, 0x89, 0x6e, 0xef, 0x3d, 0x3f, 0xdf, 0x4f,
	0x7a, 0x30, 0xf7, 0x24, 0xee, 0xcd, 0x59, 0x23, 0x8e, 0xfc, 0x0d, 0x63, 0xdc, 0x3b, 0x8b, 0xb7,
	0x21, 0x92, 0x90, 0xca, 0xf6, 0x77, 0xf3, 0xeb, 0x0d, 0xd1, 0x66, 0x8b, 0x4b, 0x13, 0xdc, 0xd2,
	0x78, 0x4f, 0x92, 0x8a, 0xdc, 0x35, 0xca, 0x08, 0xb3, 0xa7, 0x9d, 0x71, 0xdb, 0x0a, 0x3f, 0x5b,
	0x64, 0x51, 0x67, 0x90, 0x9b, 0xe0, 0xf4, 0xa8, 0x18, 0x2d, 0xa6, 0xd5, 0x41, 0xaa, 0x12, 0x66,
	0x11, 0xad, 0x0b, 0x0e, 0xbd, 0x3c, 0xaf, 0x59, 0x67, 0x45, 0xbe, 0x98, 0x56, 0x7f, 0x32, 0xa5,
	0xe1, 0x84, 0xdb, 0xe6, 0x03, 0xad, 0xe8, 0x3c, 0xfd, 0xfc, 0xb1, 0x4a, 0xc1, 0xb8, 0xa1, 0xf5,
	0x97, 0x1e, 0xa7, 0x38, 0xe9, 0xf2, 0x15, 0xa0, 0x7e, 0xa9, 0x8f, 0x13, 0x0b, 0x38, 0x0d, 0xef,
	0xe4, 0x71, 0xd5, 0xee, 0x1a, 0x8c, 0x3a, 0x4b, 0x2f, 0xc3, 0xe8, 0xc0, 0xb3, 0xe4, 0x05, 0xfd,
	0x2f, 0xaf, 0xb7, 0xe5, 0x23, 0x5c, 0xac, 0x06, 0x7b, 0x54, 0xc8, 0x81, 0x3c, 0xe3, 0x3f, 0x94,
	0x2b, 0x98, 0xb0, 0x18, 0x69, 0xb9, 0x07, 0xf4, 0xee, 0xfe, 0x12, 0xce, 0x87, 0x17, 0xea, 0x6e,
	0xd0, 0x66, 0x92, 0xf6, 0x7a, 0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x60, 0x78, 0x43, 0x63, 0x6f,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "notification-service.proto",
}
