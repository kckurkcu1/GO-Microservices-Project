// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package pb

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

type Account struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PostAccountRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostAccountRequest) Reset()         { *m = PostAccountRequest{} }
func (m *PostAccountRequest) String() string { return proto.CompactTextString(m) }
func (*PostAccountRequest) ProtoMessage()    {}
func (*PostAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *PostAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostAccountRequest.Unmarshal(m, b)
}
func (m *PostAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostAccountRequest.Marshal(b, m, deterministic)
}
func (m *PostAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostAccountRequest.Merge(m, src)
}
func (m *PostAccountRequest) XXX_Size() int {
	return xxx_messageInfo_PostAccountRequest.Size(m)
}
func (m *PostAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostAccountRequest proto.InternalMessageInfo

func (m *PostAccountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PostAccountResponse struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostAccountResponse) Reset()         { *m = PostAccountResponse{} }
func (m *PostAccountResponse) String() string { return proto.CompactTextString(m) }
func (*PostAccountResponse) ProtoMessage()    {}
func (*PostAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *PostAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostAccountResponse.Unmarshal(m, b)
}
func (m *PostAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostAccountResponse.Marshal(b, m, deterministic)
}
func (m *PostAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostAccountResponse.Merge(m, src)
}
func (m *PostAccountResponse) XXX_Size() int {
	return xxx_messageInfo_PostAccountResponse.Size(m)
}
func (m *PostAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostAccountResponse proto.InternalMessageInfo

func (m *PostAccountResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type GetAccountRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountRequest) Reset()         { *m = GetAccountRequest{} }
func (m *GetAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountRequest) ProtoMessage()    {}
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}

func (m *GetAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountRequest.Unmarshal(m, b)
}
func (m *GetAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountRequest.Marshal(b, m, deterministic)
}
func (m *GetAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountRequest.Merge(m, src)
}
func (m *GetAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountRequest.Size(m)
}
func (m *GetAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountRequest proto.InternalMessageInfo

func (m *GetAccountRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetAccountResponse struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountResponse) Reset()         { *m = GetAccountResponse{} }
func (m *GetAccountResponse) String() string { return proto.CompactTextString(m) }
func (*GetAccountResponse) ProtoMessage()    {}
func (*GetAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{4}
}

func (m *GetAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountResponse.Unmarshal(m, b)
}
func (m *GetAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountResponse.Marshal(b, m, deterministic)
}
func (m *GetAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountResponse.Merge(m, src)
}
func (m *GetAccountResponse) XXX_Size() int {
	return xxx_messageInfo_GetAccountResponse.Size(m)
}
func (m *GetAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountResponse proto.InternalMessageInfo

func (m *GetAccountResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type GetAccountsRequest struct {
	Skip                 uint64   `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Take                 uint64   `protobuf:"varint,2,opt,name=take,proto3" json:"take,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountsRequest) Reset()         { *m = GetAccountsRequest{} }
func (m *GetAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountsRequest) ProtoMessage()    {}
func (*GetAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{5}
}

func (m *GetAccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountsRequest.Unmarshal(m, b)
}
func (m *GetAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountsRequest.Marshal(b, m, deterministic)
}
func (m *GetAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountsRequest.Merge(m, src)
}
func (m *GetAccountsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountsRequest.Size(m)
}
func (m *GetAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountsRequest proto.InternalMessageInfo

func (m *GetAccountsRequest) GetSkip() uint64 {
	if m != nil {
		return m.Skip
	}
	return 0
}

func (m *GetAccountsRequest) GetTake() uint64 {
	if m != nil {
		return m.Take
	}
	return 0
}

type GetAccountsResponse struct {
	Accounts             []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetAccountsResponse) Reset()         { *m = GetAccountsResponse{} }
func (m *GetAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAccountsResponse) ProtoMessage()    {}
func (*GetAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{6}
}

func (m *GetAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountsResponse.Unmarshal(m, b)
}
func (m *GetAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountsResponse.Marshal(b, m, deterministic)
}
func (m *GetAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountsResponse.Merge(m, src)
}
func (m *GetAccountsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAccountsResponse.Size(m)
}
func (m *GetAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountsResponse proto.InternalMessageInfo

func (m *GetAccountsResponse) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "pb.Account")
	proto.RegisterType((*PostAccountRequest)(nil), "pb.PostAccountRequest")
	proto.RegisterType((*PostAccountResponse)(nil), "pb.PostAccountResponse")
	proto.RegisterType((*GetAccountRequest)(nil), "pb.GetAccountRequest")
	proto.RegisterType((*GetAccountResponse)(nil), "pb.GetAccountResponse")
	proto.RegisterType((*GetAccountsRequest)(nil), "pb.GetAccountsRequest")
	proto.RegisterType((*GetAccountsResponse)(nil), "pb.GetAccountsResponse")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xeb, 0x10, 0x51, 0xb8, 0x88, 0x4a, 0x5c, 0x45, 0x89, 0x3a, 0x55, 0x46, 0x88, 0x2e,
	0x64, 0x28, 0x23, 0x05, 0x95, 0x89, 0x15, 0x85, 0x5f, 0x90, 0xa4, 0x1e, 0xac, 0x8a, 0xd8, 0xf4,
	0x5c, 0xfe, 0x28, 0x7f, 0x08, 0xc5, 0x71, 0x83, 0xe3, 0xb0, 0xb0, 0x9d, 0x9e, 0xbf, 0x97, 0x77,
	0xf7, 0x14, 0xb8, 0x28, 0xaa, 0x4a, 0x1d, 0x6a, 0x93, 0xe9, 0xbd, 0x32, 0x0a, 0x23, 0x5d, 0xf2,
	0x7b, 0x18, 0xbf, 0xb4, 0x22, 0x4e, 0x20, 0x92, 0xdb, 0x94, 0x2d, 0xd8, 0xf2, 0x3c, 0x8f, 0xe4,
	0x16, 0x11, 0xe2, 0xba, 0xf8, 0x10, 0x69, 0x64, 0x15, 0x3b, 0xf3, 0x25, 0xe0, 0x9b, 0x22, 0xe3,
	0x2c, 0xb9, 0xf8, 0x3c, 0x08, 0x32, 0x1d, 0xc9, 0x3c, 0x72, 0x0d, 0xd3, 0x1e, 0x49, 0x5a, 0xd5,
	0x24, 0xf0, 0x16, 0xc6, 0x6e, 0x09, 0x4b, 0x27, 0xab, 0x24, 0xd3, 0x65, 0x76, 0xa4, 0x8e, 0x6f,
	0xfc, 0x06, 0x2e, 0x5f, 0x45, 0x18, 0x13, 0x2c, 0xc8, 0x1f, 0x01, 0x7d, 0xe8, 0x7f, 0x09, 0x6b,
	0xdf, 0x4c, 0xde, 0x25, 0xb4, 0x93, 0xda, 0x3a, 0xe3, 0xdc, 0xce, 0x8d, 0x66, 0x8a, 0x5d, 0xdb,
	0x43, 0x9c, 0xdb, 0x99, 0x3f, 0xc3, 0xb4, 0xe7, 0x76, 0xd9, 0x77, 0x70, 0xe6, 0xbe, 0x4f, 0x29,
	0x5b, 0x9c, 0x84, 0xe1, 0xdd, 0xe3, 0xea, 0x9b, 0xc1, 0xc4, 0xa9, 0xef, 0x62, 0xff, 0x25, 0x2b,
	0x81, 0x1b, 0x48, 0xbc, 0xc2, 0x70, 0xd6, 0x18, 0x87, 0x5d, 0xcf, 0xaf, 0x07, 0x7a, 0x9b, 0xcd,
	0x47, 0xf8, 0x04, 0xf0, 0xbb, 0x14, 0x5e, 0x35, 0xe0, 0xa0, 0xc4, 0xf9, 0x2c, 0x94, 0x3b, 0xfb,
	0x06, 0x12, 0xef, 0x26, 0x0c, 0x40, 0xea, 0x2d, 0xf0, 0xc7, 0xf1, 0x7c, 0x54, 0x9e, 0xda, 0xff,
	0xea, 0xe1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x06, 0x3f, 0x3f, 0xef, 0x68, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	PostAccount(ctx context.Context, in *PostAccountRequest, opts ...grpc.CallOption) (*PostAccountResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	GetAccounts(ctx context.Context, in *GetAccountsRequest, opts ...grpc.CallOption) (*GetAccountsResponse, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) PostAccount(ctx context.Context, in *PostAccountRequest, opts ...grpc.CallOption) (*PostAccountResponse, error) {
	out := new(PostAccountResponse)
	err := c.cc.Invoke(ctx, "/pb.AccountService/PostAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, "/pb.AccountService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccounts(ctx context.Context, in *GetAccountsRequest, opts ...grpc.CallOption) (*GetAccountsResponse, error) {
	out := new(GetAccountsResponse)
	err := c.cc.Invoke(ctx, "/pb.AccountService/GetAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	PostAccount(context.Context, *PostAccountRequest) (*PostAccountResponse, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	GetAccounts(context.Context, *GetAccountsRequest) (*GetAccountsResponse, error)
}

// UnimplementedAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (*UnimplementedAccountServiceServer) PostAccount(ctx context.Context, req *PostAccountRequest) (*PostAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostAccount not implemented")
}
func (*UnimplementedAccountServiceServer) GetAccount(ctx context.Context, req *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (*UnimplementedAccountServiceServer) GetAccounts(ctx context.Context, req *GetAccountsRequest) (*GetAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccounts not implemented")
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_PostAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).PostAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AccountService/PostAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).PostAccount(ctx, req.(*PostAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AccountService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AccountService/GetAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccounts(ctx, req.(*GetAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostAccount",
			Handler:    _AccountService_PostAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _AccountService_GetAccount_Handler,
		},
		{
			MethodName: "GetAccounts",
			Handler:    _AccountService_GetAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}