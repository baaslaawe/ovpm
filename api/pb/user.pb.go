// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	user.proto
	vpn.proto
	network.proto
	auth.proto

It has these top-level messages:
	UserListRequest
	UserCreateRequest
	UserUpdateRequest
	UserDeleteRequest
	UserRenewRequest
	UserGenConfigRequest
	UserResponse
	UserGenConfigResponse
	VPNStatusRequest
	VPNInitRequest
	VPNUpdateRequest
	VPNRestartRequest
	VPNStatusResponse
	VPNInitResponse
	VPNUpdateResponse
	VPNRestartResponse
	NetworkCreateRequest
	NetworkListRequest
	NetworkDeleteRequest
	NetworkGetAllTypesRequest
	NetworkAssociateRequest
	NetworkDissociateRequest
	NetworkGetAssociatedUsersRequest
	Network
	NetworkType
	NetworkCreateResponse
	NetworkListResponse
	NetworkDeleteResponse
	NetworkGetAllTypesResponse
	NetworkAssociateResponse
	NetworkDissociateResponse
	NetworkGetAssociatedUsersResponse
	AuthStatusRequest
	AuthAuthenticateRequest
	AuthStatusResponse
	AuthAuthenticateResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type UserUpdateRequest_GWPref int32

const (
	UserUpdateRequest_NOPREF UserUpdateRequest_GWPref = 0
	UserUpdateRequest_NOGW   UserUpdateRequest_GWPref = 1
	UserUpdateRequest_GW     UserUpdateRequest_GWPref = 2
)

var UserUpdateRequest_GWPref_name = map[int32]string{
	0: "NOPREF",
	1: "NOGW",
	2: "GW",
}
var UserUpdateRequest_GWPref_value = map[string]int32{
	"NOPREF": 0,
	"NOGW":   1,
	"GW":     2,
}

func (x UserUpdateRequest_GWPref) String() string {
	return proto.EnumName(UserUpdateRequest_GWPref_name, int32(x))
}
func (UserUpdateRequest_GWPref) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type UserUpdateRequest_StaticPref int32

const (
	UserUpdateRequest_NOPREFSTATIC UserUpdateRequest_StaticPref = 0
	UserUpdateRequest_NOSTATIC     UserUpdateRequest_StaticPref = 1
	UserUpdateRequest_STATIC       UserUpdateRequest_StaticPref = 2
)

var UserUpdateRequest_StaticPref_name = map[int32]string{
	0: "NOPREFSTATIC",
	1: "NOSTATIC",
	2: "STATIC",
}
var UserUpdateRequest_StaticPref_value = map[string]int32{
	"NOPREFSTATIC": 0,
	"NOSTATIC":     1,
	"STATIC":       2,
}

func (x UserUpdateRequest_StaticPref) String() string {
	return proto.EnumName(UserUpdateRequest_StaticPref_name, int32(x))
}
func (UserUpdateRequest_StaticPref) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 1}
}

type UserUpdateRequest_AdminPref int32

const (
	UserUpdateRequest_NOPREFADMIN UserUpdateRequest_AdminPref = 0
	UserUpdateRequest_NOADMIN     UserUpdateRequest_AdminPref = 1
	UserUpdateRequest_ADMIN       UserUpdateRequest_AdminPref = 2
)

var UserUpdateRequest_AdminPref_name = map[int32]string{
	0: "NOPREFADMIN",
	1: "NOADMIN",
	2: "ADMIN",
}
var UserUpdateRequest_AdminPref_value = map[string]int32{
	"NOPREFADMIN": 0,
	"NOADMIN":     1,
	"ADMIN":       2,
}

func (x UserUpdateRequest_AdminPref) String() string {
	return proto.EnumName(UserUpdateRequest_AdminPref_name, int32(x))
}
func (UserUpdateRequest_AdminPref) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 2}
}

type UserListRequest struct {
}

func (m *UserListRequest) Reset()                    { *m = UserListRequest{} }
func (m *UserListRequest) String() string            { return proto.CompactTextString(m) }
func (*UserListRequest) ProtoMessage()               {}
func (*UserListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type UserCreateRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	NoGw     bool   `protobuf:"varint,3,opt,name=no_gw,json=noGw" json:"no_gw,omitempty"`
	HostId   uint32 `protobuf:"varint,4,opt,name=host_id,json=hostId" json:"host_id,omitempty"`
	IsAdmin  bool   `protobuf:"varint,5,opt,name=is_admin,json=isAdmin" json:"is_admin,omitempty"`
}

func (m *UserCreateRequest) Reset()                    { *m = UserCreateRequest{} }
func (m *UserCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*UserCreateRequest) ProtoMessage()               {}
func (*UserCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserCreateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserCreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserCreateRequest) GetNoGw() bool {
	if m != nil {
		return m.NoGw
	}
	return false
}

func (m *UserCreateRequest) GetHostId() uint32 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *UserCreateRequest) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

type UserUpdateRequest struct {
	Username   string                       `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password   string                       `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Gwpref     UserUpdateRequest_GWPref     `protobuf:"varint,3,opt,name=gwpref,enum=pb.UserUpdateRequest_GWPref" json:"gwpref,omitempty"`
	HostId     uint32                       `protobuf:"varint,4,opt,name=host_id,json=hostId" json:"host_id,omitempty"`
	StaticPref UserUpdateRequest_StaticPref `protobuf:"varint,5,opt,name=static_pref,json=staticPref,enum=pb.UserUpdateRequest_StaticPref" json:"static_pref,omitempty"`
	AdminPref  UserUpdateRequest_AdminPref  `protobuf:"varint,6,opt,name=admin_pref,json=adminPref,enum=pb.UserUpdateRequest_AdminPref" json:"admin_pref,omitempty"`
}

func (m *UserUpdateRequest) Reset()                    { *m = UserUpdateRequest{} }
func (m *UserUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UserUpdateRequest) ProtoMessage()               {}
func (*UserUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserUpdateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserUpdateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserUpdateRequest) GetGwpref() UserUpdateRequest_GWPref {
	if m != nil {
		return m.Gwpref
	}
	return UserUpdateRequest_NOPREF
}

func (m *UserUpdateRequest) GetHostId() uint32 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *UserUpdateRequest) GetStaticPref() UserUpdateRequest_StaticPref {
	if m != nil {
		return m.StaticPref
	}
	return UserUpdateRequest_NOPREFSTATIC
}

func (m *UserUpdateRequest) GetAdminPref() UserUpdateRequest_AdminPref {
	if m != nil {
		return m.AdminPref
	}
	return UserUpdateRequest_NOPREFADMIN
}

type UserDeleteRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *UserDeleteRequest) Reset()                    { *m = UserDeleteRequest{} }
func (m *UserDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*UserDeleteRequest) ProtoMessage()               {}
func (*UserDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserDeleteRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserRenewRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *UserRenewRequest) Reset()                    { *m = UserRenewRequest{} }
func (m *UserRenewRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRenewRequest) ProtoMessage()               {}
func (*UserRenewRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserRenewRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserGenConfigRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *UserGenConfigRequest) Reset()                    { *m = UserGenConfigRequest{} }
func (m *UserGenConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*UserGenConfigRequest) ProtoMessage()               {}
func (*UserGenConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserGenConfigRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserResponse struct {
	Users []*UserResponse_User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UserResponse) GetUsers() []*UserResponse_User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UserResponse_User struct {
	Username           string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	ServerSerialNumber string `protobuf:"bytes,2,opt,name=server_serial_number,json=serverSerialNumber" json:"server_serial_number,omitempty"`
	Cert               string `protobuf:"bytes,3,opt,name=cert" json:"cert,omitempty"`
	CreatedAt          string `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	IpNet              string `protobuf:"bytes,5,opt,name=ip_net,json=ipNet" json:"ip_net,omitempty"`
	NoGw               bool   `protobuf:"varint,6,opt,name=no_gw,json=noGw" json:"no_gw,omitempty"`
	HostId             uint32 `protobuf:"varint,7,opt,name=host_id,json=hostId" json:"host_id,omitempty"`
	IsAdmin            bool   `protobuf:"varint,8,opt,name=is_admin,json=isAdmin" json:"is_admin,omitempty"`
	IsConnected        bool   `protobuf:"varint,9,opt,name=is_connected,json=isConnected" json:"is_connected,omitempty"`
	ConnectedSince     string `protobuf:"bytes,10,opt,name=connected_since,json=connectedSince" json:"connected_since,omitempty"`
	BytesSent          uint64 `protobuf:"varint,11,opt,name=bytes_sent,json=bytesSent" json:"bytes_sent,omitempty"`
	BytesReceived      uint64 `protobuf:"varint,12,opt,name=bytes_received,json=bytesReceived" json:"bytes_received,omitempty"`
	ExpiresAt          string `protobuf:"bytes,13,opt,name=expires_at,json=expiresAt" json:"expires_at,omitempty"`
}

func (m *UserResponse_User) Reset()                    { *m = UserResponse_User{} }
func (m *UserResponse_User) String() string            { return proto.CompactTextString(m) }
func (*UserResponse_User) ProtoMessage()               {}
func (*UserResponse_User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *UserResponse_User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserResponse_User) GetServerSerialNumber() string {
	if m != nil {
		return m.ServerSerialNumber
	}
	return ""
}

func (m *UserResponse_User) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

func (m *UserResponse_User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UserResponse_User) GetIpNet() string {
	if m != nil {
		return m.IpNet
	}
	return ""
}

func (m *UserResponse_User) GetNoGw() bool {
	if m != nil {
		return m.NoGw
	}
	return false
}

func (m *UserResponse_User) GetHostId() uint32 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *UserResponse_User) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *UserResponse_User) GetIsConnected() bool {
	if m != nil {
		return m.IsConnected
	}
	return false
}

func (m *UserResponse_User) GetConnectedSince() string {
	if m != nil {
		return m.ConnectedSince
	}
	return ""
}

func (m *UserResponse_User) GetBytesSent() uint64 {
	if m != nil {
		return m.BytesSent
	}
	return 0
}

func (m *UserResponse_User) GetBytesReceived() uint64 {
	if m != nil {
		return m.BytesReceived
	}
	return 0
}

func (m *UserResponse_User) GetExpiresAt() string {
	if m != nil {
		return m.ExpiresAt
	}
	return ""
}

type UserGenConfigResponse struct {
	ClientConfig string `protobuf:"bytes,1,opt,name=client_config,json=clientConfig" json:"client_config,omitempty"`
}

func (m *UserGenConfigResponse) Reset()                    { *m = UserGenConfigResponse{} }
func (m *UserGenConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*UserGenConfigResponse) ProtoMessage()               {}
func (*UserGenConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UserGenConfigResponse) GetClientConfig() string {
	if m != nil {
		return m.ClientConfig
	}
	return ""
}

func init() {
	proto.RegisterType((*UserListRequest)(nil), "pb.UserListRequest")
	proto.RegisterType((*UserCreateRequest)(nil), "pb.UserCreateRequest")
	proto.RegisterType((*UserUpdateRequest)(nil), "pb.UserUpdateRequest")
	proto.RegisterType((*UserDeleteRequest)(nil), "pb.UserDeleteRequest")
	proto.RegisterType((*UserRenewRequest)(nil), "pb.UserRenewRequest")
	proto.RegisterType((*UserGenConfigRequest)(nil), "pb.UserGenConfigRequest")
	proto.RegisterType((*UserResponse)(nil), "pb.UserResponse")
	proto.RegisterType((*UserResponse_User)(nil), "pb.UserResponse.User")
	proto.RegisterType((*UserGenConfigResponse)(nil), "pb.UserGenConfigResponse")
	proto.RegisterEnum("pb.UserUpdateRequest_GWPref", UserUpdateRequest_GWPref_name, UserUpdateRequest_GWPref_value)
	proto.RegisterEnum("pb.UserUpdateRequest_StaticPref", UserUpdateRequest_StaticPref_name, UserUpdateRequest_StaticPref_value)
	proto.RegisterEnum("pb.UserUpdateRequest_AdminPref", UserUpdateRequest_AdminPref_name, UserUpdateRequest_AdminPref_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	List(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Renew(ctx context.Context, in *UserRenewRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GenConfig(ctx context.Context, in *UserGenConfigRequest, opts ...grpc.CallOption) (*UserGenConfigResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) List(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Renew(ctx context.Context, in *UserRenewRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/Renew", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GenConfig(ctx context.Context, in *UserGenConfigRequest, opts ...grpc.CallOption) (*UserGenConfigResponse, error) {
	out := new(UserGenConfigResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/GenConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	List(context.Context, *UserListRequest) (*UserResponse, error)
	Create(context.Context, *UserCreateRequest) (*UserResponse, error)
	Update(context.Context, *UserUpdateRequest) (*UserResponse, error)
	Delete(context.Context, *UserDeleteRequest) (*UserResponse, error)
	Renew(context.Context, *UserRenewRequest) (*UserResponse, error)
	GenConfig(context.Context, *UserGenConfigRequest) (*UserGenConfigResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).List(ctx, req.(*UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRenewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Renew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Renew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Renew(ctx, req.(*UserRenewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GenConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGenConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GenConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GenConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GenConfig(ctx, req.(*UserGenConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _UserService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
		{
			MethodName: "Renew",
			Handler:    _UserService_Renew_Handler,
		},
		{
			MethodName: "GenConfig",
			Handler:    _UserService_GenConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 825 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x80, 0xeb, 0xc4, 0x71, 0xe3, 0x93, 0xb4, 0x75, 0xa7, 0x2d, 0xeb, 0x86, 0x5d, 0xc8, 0x1a,
	0x01, 0xd1, 0x22, 0x25, 0x50, 0xf6, 0x02, 0xad, 0x10, 0x52, 0xd4, 0x85, 0x50, 0x09, 0x9c, 0xc5,
	0xd9, 0x55, 0x2f, 0x2d, 0xc7, 0x3e, 0x09, 0x23, 0xa5, 0x63, 0x33, 0x33, 0x49, 0xe0, 0x16, 0x89,
	0x27, 0x80, 0x0b, 0x5e, 0x84, 0x27, 0x41, 0xe2, 0x09, 0x78, 0x10, 0x34, 0x33, 0x4e, 0x42, 0x4a,
	0x16, 0xe5, 0x62, 0xef, 0xe6, 0xfc, 0x7d, 0x73, 0x3c, 0xe7, 0xc7, 0x00, 0x73, 0x81, 0xbc, 0x5b,
	0xf0, 0x5c, 0xe6, 0xa4, 0x52, 0x8c, 0x5b, 0x0f, 0xa7, 0x79, 0x3e, 0x9d, 0x61, 0x2f, 0x29, 0x68,
	0x2f, 0x61, 0x2c, 0x97, 0x89, 0xa4, 0x39, 0x13, 0xc6, 0x23, 0x38, 0x85, 0x93, 0x57, 0x02, 0xf9,
	0x37, 0x54, 0xc8, 0x08, 0x7f, 0x98, 0xa3, 0x90, 0xc1, 0x6f, 0x16, 0x9c, 0x2a, 0xdd, 0x35, 0xc7,
	0x44, 0x62, 0xa9, 0x25, 0x2d, 0xa8, 0x2b, 0x30, 0x4b, 0xee, 0xd0, 0xb7, 0xda, 0x56, 0xc7, 0x8d,
	0xd6, 0xb2, 0xb2, 0x15, 0x89, 0x10, 0xcb, 0x9c, 0x67, 0x7e, 0xc5, 0xd8, 0x56, 0x32, 0x39, 0x83,
	0x1a, 0xcb, 0xe3, 0xe9, 0xd2, 0xaf, 0xb6, 0xad, 0x4e, 0x3d, 0xb2, 0x59, 0x3e, 0x58, 0x92, 0x07,
	0x70, 0xf8, 0x7d, 0x2e, 0x64, 0x4c, 0x33, 0xdf, 0x6e, 0x5b, 0x9d, 0xa3, 0xc8, 0x51, 0xe2, 0x4d,
	0x46, 0x2e, 0xa1, 0x4e, 0x45, 0x9c, 0x64, 0x77, 0x94, 0xf9, 0x35, 0x1d, 0x70, 0x48, 0x45, 0x5f,
	0x89, 0xc1, 0x1f, 0x55, 0x93, 0xd6, 0xab, 0x22, 0x7b, 0x03, 0x69, 0x3d, 0x05, 0x67, 0xba, 0x2c,
	0x38, 0x4e, 0x74, 0x5e, 0xc7, 0x57, 0x0f, 0xbb, 0xc5, 0xb8, 0xfb, 0x1f, 0x7c, 0x77, 0x70, 0xfb,
	0x82, 0xe3, 0x24, 0x2a, 0x7d, 0x5f, 0x9f, 0x77, 0x1f, 0x1a, 0x42, 0x3d, 0x6c, 0x1a, 0x6b, 0x66,
	0x4d, 0x33, 0xdb, 0xbb, 0x99, 0x23, 0xed, 0xa8, 0xb9, 0x20, 0xd6, 0x67, 0xf2, 0x05, 0x80, 0xfe,
	0x6e, 0x43, 0x70, 0x34, 0xe1, 0xdd, 0xdd, 0x04, 0xfd, 0x20, 0x1a, 0xe0, 0x26, 0xab, 0x63, 0xf0,
	0x01, 0x38, 0x26, 0x5b, 0x02, 0xe0, 0x84, 0xc3, 0x17, 0xd1, 0x97, 0x5f, 0x79, 0x07, 0xa4, 0x0e,
	0x76, 0x38, 0x1c, 0xdc, 0x7a, 0x16, 0x71, 0xa0, 0x32, 0xb8, 0xf5, 0x2a, 0xc1, 0x67, 0x00, 0x9b,
	0x0c, 0x88, 0x07, 0x4d, 0xe3, 0x3b, 0x7a, 0xd9, 0x7f, 0x79, 0x73, 0xed, 0x1d, 0x90, 0x26, 0xd4,
	0xc3, 0x61, 0x29, 0x59, 0x8a, 0x55, 0x9e, 0x2b, 0xc1, 0x53, 0x70, 0xd7, 0x37, 0x93, 0x13, 0x68,
	0x98, 0xc0, 0xfe, 0xf3, 0x6f, 0x6f, 0x42, 0xef, 0x80, 0x34, 0xe0, 0x30, 0x1c, 0x1a, 0xc1, 0x22,
	0x2e, 0xd4, 0xcc, 0xb1, 0x12, 0xf4, 0x4c, 0xd9, 0x9e, 0xe3, 0x0c, 0xf7, 0x2a, 0x5b, 0xd0, 0x05,
	0x4f, 0x05, 0x44, 0xc8, 0x70, 0xb9, 0x8f, 0xff, 0x15, 0x9c, 0x2b, 0xff, 0x01, 0xb2, 0xeb, 0x9c,
	0x4d, 0xe8, 0x74, 0x9f, 0x98, 0xbf, 0xaa, 0xd0, 0x34, 0x97, 0x88, 0x22, 0x67, 0x02, 0xc9, 0x47,
	0x50, 0x53, 0x46, 0xe1, 0x5b, 0xed, 0x6a, 0xa7, 0x71, 0x75, 0xb1, 0x7a, 0xf8, 0x95, 0x83, 0x11,
	0x8c, 0x4f, 0xeb, 0xf7, 0x2a, 0xd8, 0x4a, 0xfe, 0xdf, 0xee, 0xfb, 0x18, 0xce, 0x05, 0xf2, 0x05,
	0xf2, 0x58, 0x20, 0xa7, 0xc9, 0x2c, 0x66, 0xf3, 0xbb, 0x31, 0xf2, 0xb2, 0x13, 0x89, 0xb1, 0x8d,
	0xb4, 0x29, 0xd4, 0x16, 0x42, 0xc0, 0x4e, 0x91, 0x4b, 0xdd, 0x91, 0x6e, 0xa4, 0xcf, 0xe4, 0x11,
	0x40, 0xaa, 0xe7, 0x30, 0x8b, 0x13, 0xa9, 0x9b, 0xce, 0x8d, 0xdc, 0x52, 0xd3, 0x97, 0xe4, 0x02,
	0x1c, 0x5a, 0xc4, 0x0c, 0xa5, 0x6e, 0x39, 0x37, 0xaa, 0xd1, 0x22, 0x44, 0xb9, 0x19, 0x3a, 0x67,
	0xf7, 0xd0, 0x1d, 0xbe, 0x76, 0xe8, 0xea, 0x5b, 0x43, 0x47, 0x1e, 0x43, 0x93, 0x8a, 0x38, 0xcd,
	0x19, 0xc3, 0x54, 0x62, 0xe6, 0xbb, 0xda, 0xdc, 0xa0, 0xe2, 0x7a, 0xa5, 0x22, 0x1f, 0xc2, 0xc9,
	0xda, 0x1e, 0x0b, 0xca, 0x52, 0xf4, 0x41, 0xe7, 0x72, 0xbc, 0x56, 0x8f, 0x94, 0x56, 0x7d, 0xca,
	0xf8, 0x27, 0x89, 0x22, 0x16, 0xc8, 0xa4, 0xdf, 0x68, 0x5b, 0x1d, 0x3b, 0x72, 0xb5, 0x66, 0x84,
	0x4c, 0x92, 0xf7, 0xe1, 0xd8, 0x98, 0x39, 0xa6, 0x48, 0x17, 0x98, 0xf9, 0x4d, 0xed, 0x72, 0xa4,
	0xb5, 0x51, 0xa9, 0x54, 0x14, 0xfc, 0xb1, 0xa0, 0x1c, 0x85, 0x7a, 0x90, 0x23, 0xf3, 0x20, 0xa5,
	0xa6, 0x2f, 0x83, 0xcf, 0xe1, 0xe2, 0x5e, 0x33, 0x94, 0x05, 0x7e, 0x0f, 0x8e, 0xd2, 0x19, 0x45,
	0x26, 0xd5, 0xd7, 0x4c, 0xe8, 0xb4, 0xac, 0x57, 0xd3, 0x28, 0x8d, 0xf3, 0xd5, 0x2f, 0x36, 0x34,
	0x54, 0xf8, 0x08, 0xf9, 0x82, 0xa6, 0x48, 0xbe, 0x06, 0x5b, 0x6d, 0x46, 0x72, 0xb6, 0x6a, 0x87,
	0x7f, 0xed, 0xc9, 0x96, 0x77, 0xbf, 0x47, 0x82, 0xcb, 0x9f, 0xff, 0xfc, 0xfb, 0xd7, 0xca, 0x19,
	0x39, 0xd5, 0xcb, 0x76, 0xf1, 0x49, 0x4f, 0x35, 0x43, 0x6f, 0xa6, 0x08, 0xdf, 0x81, 0x63, 0xf6,
	0x29, 0x59, 0xb7, 0xd6, 0xd6, 0x7e, 0xdd, 0x41, 0x7b, 0x47, 0xd3, 0xfc, 0xe0, 0x6c, 0x8b, 0x66,
	0x6a, 0xff, 0xcc, 0x7a, 0xa2, 0x90, 0x66, 0x2d, 0x6c, 0x90, 0x5b, 0x6b, 0x62, 0x6f, 0xe4, 0x5c,
	0x47, 0x95, 0x48, 0x33, 0xa7, 0x1b, 0xe4, 0xd6, 0xdc, 0xee, 0x8d, 0xcc, 0x74, 0x94, 0x42, 0x86,
	0x50, 0xd3, 0x93, 0x4c, 0xce, 0x37, 0xa1, 0x9b, 0xc1, 0xde, 0x01, 0x7c, 0xa4, 0x81, 0x0f, 0x02,
	0xb2, 0x05, 0xe4, 0x2a, 0x48, 0xf1, 0x52, 0x70, 0xd7, 0xc5, 0x25, 0xfe, 0x2a, 0xfa, 0xfe, 0xf0,
	0xb7, 0x2e, 0x77, 0x58, 0xca, 0x0b, 0x1e, 0xeb, 0x0b, 0xde, 0x0e, 0xde, 0xda, 0xba, 0x60, 0x8a,
	0xcc, 0x34, 0xc6, 0x33, 0xeb, 0xc9, 0xd8, 0xd1, 0x3f, 0xc7, 0x4f, 0xff, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x42, 0xdb, 0x7d, 0x5f, 0x4c, 0x07, 0x00, 0x00,
}
