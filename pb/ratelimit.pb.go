// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ratelimit.proto

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

type DescriptorStatus_Status int32

const (
	DescriptorStatus_UNKNOWN    DescriptorStatus_Status = 0
	DescriptorStatus_OK         DescriptorStatus_Status = 1
	DescriptorStatus_OVER_LIMIT DescriptorStatus_Status = 2
)

var DescriptorStatus_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "OVER_LIMIT",
}
var DescriptorStatus_Status_value = map[string]int32{
	"UNKNOWN":    0,
	"OK":         1,
	"OVER_LIMIT": 2,
}

func (x DescriptorStatus_Status) String() string {
	return proto.EnumName(DescriptorStatus_Status_name, int32(x))
}
func (DescriptorStatus_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{4, 0} }

type RateLimitRequest struct {
	// The domain scopes the request to avoid collisions with other services or applications.
	Domain string `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	// Rate limit requests must specify at least one descriptor
	Descriptors []*Descriptor `protobuf:"bytes,2,rep,name=descriptors" json:"descriptors,omitempty"`
}

func (m *RateLimitRequest) Reset()                    { *m = RateLimitRequest{} }
func (m *RateLimitRequest) String() string            { return proto.CompactTextString(m) }
func (*RateLimitRequest) ProtoMessage()               {}
func (*RateLimitRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *RateLimitRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitRequest) GetDescriptors() []*Descriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

type RateLimitResponse struct {
	// A list of DescriptorStatus messages which matches the length of the descriptor list passed
	// in the RateLimitRequest. This can be used by the caller to determine which individual
	// descriptors failed and/or what the currently configured limits are for all of them.
	Statuses []*DescriptorStatus `protobuf:"bytes,2,rep,name=statuses" json:"statuses,omitempty"`
}

func (m *RateLimitResponse) Reset()                    { *m = RateLimitResponse{} }
func (m *RateLimitResponse) String() string            { return proto.CompactTextString(m) }
func (*RateLimitResponse) ProtoMessage()               {}
func (*RateLimitResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *RateLimitResponse) GetStatuses() []*DescriptorStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

// This is used by Gubernator clients who already have the domain and descriptors and created a hash key
type RateLimitKeyRequest struct {
	// Requests must specify at least one KeyRequestEntry, but can specify multiple descriptors
	// to be individually evaluated
	Entries []*RateLimitKeyRequest_Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *RateLimitKeyRequest) Reset()                    { *m = RateLimitKeyRequest{} }
func (m *RateLimitKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*RateLimitKeyRequest) ProtoMessage()               {}
func (*RateLimitKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *RateLimitKeyRequest) GetEntries() []*RateLimitKeyRequest_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type RateLimitKeyRequest_Entry struct {
	// The key identifies the specific rate limit the request is for
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Rate limit requests can optionally specify the number of hits a request adds to the matched limit. If the
	// value is not set in the message, a request increases the matched limit by 1.
	Hits int64 `protobuf:"varint,2,opt,name=hits" json:"hits,omitempty"`
	// This is the rate limit the requested for this descriptor. If the rate limit doesn't currently exist for
	// this descriptor it will be created and enforced for subsequent requests
	RateLimitConfig *RateLimitConfig `protobuf:"bytes,3,opt,name=rate_limit_config,json=rateLimitConfig" json:"rate_limit_config,omitempty"`
}

func (m *RateLimitKeyRequest_Entry) Reset()                    { *m = RateLimitKeyRequest_Entry{} }
func (m *RateLimitKeyRequest_Entry) String() string            { return proto.CompactTextString(m) }
func (*RateLimitKeyRequest_Entry) ProtoMessage()               {}
func (*RateLimitKeyRequest_Entry) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2, 0} }

func (m *RateLimitKeyRequest_Entry) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RateLimitKeyRequest_Entry) GetHits() int64 {
	if m != nil {
		return m.Hits
	}
	return 0
}

func (m *RateLimitKeyRequest_Entry) GetRateLimitConfig() *RateLimitConfig {
	if m != nil {
		return m.RateLimitConfig
	}
	return nil
}

// Describes a rate limit
type Descriptor struct {
	// A map of key value pairs that make up the descriptor
	Values map[string]string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Rate limit requests optionally specify the number of hits a request adds to the matched limit. If the
	// value is not set, a request increases the matched limit by 1.
	Hits int64 `protobuf:"varint,2,opt,name=hits" json:"hits,omitempty"`
	// This is the rate limit the requested for this descriptor. If the rate limit doesn't currently exist for
	// this descriptor it will be created and enforced for subsequent requests
	RateLimitConfig *RateLimitConfig `protobuf:"bytes,3,opt,name=rate_limit_config,json=rateLimitConfig" json:"rate_limit_config,omitempty"`
}

func (m *Descriptor) Reset()                    { *m = Descriptor{} }
func (m *Descriptor) String() string            { return proto.CompactTextString(m) }
func (*Descriptor) ProtoMessage()               {}
func (*Descriptor) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *Descriptor) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Descriptor) GetHits() int64 {
	if m != nil {
		return m.Hits
	}
	return 0
}

func (m *Descriptor) GetRateLimitConfig() *RateLimitConfig {
	if m != nil {
		return m.RateLimitConfig
	}
	return nil
}

// The return status of a descriptor submitted via RateLimitRequest
type DescriptorStatus struct {
	// The status of this descriptor.
	Status DescriptorStatus_Status `protobuf:"varint,1,opt,name=status,enum=pb.gubernator.DescriptorStatus_Status" json:"status,omitempty"`
	// The current limit for the requested rate limit.
	CurrentLimit int64 `protobuf:"varint,2,opt,name=current_limit,json=currentLimit" json:"current_limit,omitempty"`
	// If Code is not OVER_LIMIT, This is the limit remaining
	LimitRemaining int64 `protobuf:"varint,3,opt,name=limit_remaining,json=limitRemaining" json:"limit_remaining,omitempty"`
	// This is the time when the rate limit span will be reset as a unix timestamp.
	ResetTime int64 `protobuf:"varint,4,opt,name=reset_time,json=resetTime" json:"reset_time,omitempty"`
	// This is additional metadata that a client might find useful. (IE: Additional headers, corrdinator ownership, etc..)
	Metadata map[string]string `protobuf:"bytes,5,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DescriptorStatus) Reset()                    { *m = DescriptorStatus{} }
func (m *DescriptorStatus) String() string            { return proto.CompactTextString(m) }
func (*DescriptorStatus) ProtoMessage()               {}
func (*DescriptorStatus) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *DescriptorStatus) GetStatus() DescriptorStatus_Status {
	if m != nil {
		return m.Status
	}
	return DescriptorStatus_UNKNOWN
}

func (m *DescriptorStatus) GetCurrentLimit() int64 {
	if m != nil {
		return m.CurrentLimit
	}
	return 0
}

func (m *DescriptorStatus) GetLimitRemaining() int64 {
	if m != nil {
		return m.LimitRemaining
	}
	return 0
}

func (m *DescriptorStatus) GetResetTime() int64 {
	if m != nil {
		return m.ResetTime
	}
	return 0
}

func (m *DescriptorStatus) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimitRequest)(nil), "pb.gubernator.RateLimitRequest")
	proto.RegisterType((*RateLimitResponse)(nil), "pb.gubernator.RateLimitResponse")
	proto.RegisterType((*RateLimitKeyRequest)(nil), "pb.gubernator.RateLimitKeyRequest")
	proto.RegisterType((*RateLimitKeyRequest_Entry)(nil), "pb.gubernator.RateLimitKeyRequest.Entry")
	proto.RegisterType((*Descriptor)(nil), "pb.gubernator.Descriptor")
	proto.RegisterType((*DescriptorStatus)(nil), "pb.gubernator.DescriptorStatus")
	proto.RegisterEnum("pb.gubernator.DescriptorStatus_Status", DescriptorStatus_Status_name, DescriptorStatus_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RateLimitService service

type RateLimitServiceClient interface {
	// Given a rate limit descriptor return a descriptor status
	GetRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error)
	// Client implementations can use this method if the decriptor key is already known
	GetRateLimitByKey(ctx context.Context, in *RateLimitKeyRequest, opts ...grpc.CallOption) (*RateLimitResponse, error)
}

type rateLimitServiceClient struct {
	cc *grpc.ClientConn
}

func NewRateLimitServiceClient(cc *grpc.ClientConn) RateLimitServiceClient {
	return &rateLimitServiceClient{cc}
}

func (c *rateLimitServiceClient) GetRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error) {
	out := new(RateLimitResponse)
	err := grpc.Invoke(ctx, "/pb.gubernator.RateLimitService/GetRateLimit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateLimitServiceClient) GetRateLimitByKey(ctx context.Context, in *RateLimitKeyRequest, opts ...grpc.CallOption) (*RateLimitResponse, error) {
	out := new(RateLimitResponse)
	err := grpc.Invoke(ctx, "/pb.gubernator.RateLimitService/GetRateLimitByKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RateLimitService service

type RateLimitServiceServer interface {
	// Given a rate limit descriptor return a descriptor status
	GetRateLimit(context.Context, *RateLimitRequest) (*RateLimitResponse, error)
	// Client implementations can use this method if the decriptor key is already known
	GetRateLimitByKey(context.Context, *RateLimitKeyRequest) (*RateLimitResponse, error)
}

func RegisterRateLimitServiceServer(s *grpc.Server, srv RateLimitServiceServer) {
	s.RegisterService(&_RateLimitService_serviceDesc, srv)
}

func _RateLimitService_GetRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).GetRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.gubernator.RateLimitService/GetRateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).GetRateLimit(ctx, req.(*RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateLimitService_GetRateLimitByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateLimitKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).GetRateLimitByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.gubernator.RateLimitService/GetRateLimitByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).GetRateLimitByKey(ctx, req.(*RateLimitKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.gubernator.RateLimitService",
	HandlerType: (*RateLimitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRateLimit",
			Handler:    _RateLimitService_GetRateLimit_Handler,
		},
		{
			MethodName: "GetRateLimitByKey",
			Handler:    _RateLimitService_GetRateLimitByKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ratelimit.proto",
}

func init() { proto.RegisterFile("ratelimit.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xee, 0x3a, 0x8d, 0xd3, 0x4c, 0x7e, 0xea, 0x2c, 0x08, 0x99, 0x48, 0xd0, 0xc8, 0x08, 0xc8,
	0xa5, 0x3e, 0x84, 0x0b, 0x10, 0xc1, 0x21, 0x50, 0xa1, 0x90, 0x36, 0x41, 0x9b, 0x52, 0x24, 0x38,
	0x44, 0x4e, 0x32, 0x84, 0x15, 0xb1, 0x1d, 0xd6, 0x9b, 0x4a, 0xbe, 0xf1, 0x70, 0x3c, 0x02, 0xcf,
	0xc0, 0x23, 0x70, 0x46, 0xd9, 0x75, 0x8c, 0x5b, 0x61, 0x05, 0x0e, 0xdc, 0xbc, 0xb3, 0xdf, 0x7c,
	0x3b, 0xdf, 0x37, 0x33, 0x86, 0x43, 0xe1, 0x49, 0x5c, 0x72, 0x9f, 0x4b, 0x77, 0x25, 0x42, 0x19,
	0xd2, 0xda, 0x6a, 0xea, 0x2e, 0xd6, 0x53, 0x14, 0x81, 0x27, 0x43, 0xd1, 0xac, 0xce, 0x42, 0xdf,
	0x0f, 0x03, 0x7d, 0xe9, 0x2c, 0xc0, 0x62, 0x9e, 0xc4, 0xd3, 0x0d, 0x9e, 0xe1, 0x97, 0x35, 0x46,
	0x92, 0xde, 0x02, 0x73, 0x1e, 0xfa, 0x1e, 0x0f, 0x6c, 0xd2, 0x22, 0xed, 0x32, 0x4b, 0x4e, 0xb4,
	0x0b, 0x95, 0x39, 0x46, 0x33, 0xc1, 0x57, 0x32, 0x14, 0x91, 0x6d, 0xb4, 0x0a, 0xed, 0x4a, 0xe7,
	0xb6, 0x7b, 0x85, 0xde, 0x7d, 0x99, 0x22, 0x58, 0x16, 0xed, 0xbc, 0x81, 0x46, 0xe6, 0xa1, 0x68,
	0x15, 0x06, 0x11, 0xd2, 0x2e, 0x1c, 0x44, 0xd2, 0x93, 0xeb, 0x08, 0xb7, 0x74, 0x47, 0xb9, 0x74,
	0x63, 0x05, 0x64, 0x69, 0x82, 0xf3, 0x9d, 0xc0, 0x8d, 0x94, 0x72, 0x80, 0xf1, 0xb6, 0xfc, 0x1e,
	0x94, 0x30, 0x90, 0x82, 0x63, 0x64, 0x13, 0xc5, 0xd9, 0xbe, 0xc6, 0xf9, 0x87, 0x24, 0xf7, 0x24,
	0x90, 0x22, 0x66, 0xdb, 0xc4, 0x66, 0x0c, 0x45, 0x15, 0xa1, 0x16, 0x14, 0x3e, 0x63, 0xac, 0x8c,
	0xa8, 0xb2, 0xcd, 0x27, 0xa5, 0xb0, 0xff, 0x89, 0xcb, 0x4d, 0xbd, 0xa4, 0x5d, 0x60, 0xea, 0x9b,
	0xbe, 0x86, 0xc6, 0xc6, 0xf5, 0x89, 0xb2, 0x7d, 0x32, 0x0b, 0x83, 0x8f, 0x7c, 0x61, 0x17, 0x5a,
	0xa4, 0x5d, 0xe9, 0xdc, 0xcd, 0x7b, 0xfc, 0x85, 0x42, 0x31, 0xd5, 0xae, 0x4c, 0xc0, 0xf9, 0x41,
	0x00, 0x7e, 0xab, 0xa6, 0xcf, 0xc0, 0xbc, 0xf4, 0x96, 0xeb, 0x54, 0xcc, 0xfd, 0x5c, 0x83, 0xdc,
	0x0b, 0x85, 0xd3, 0x4a, 0x92, 0xa4, 0xff, 0x5d, 0x6d, 0xf3, 0x09, 0x54, 0x32, 0xcf, 0x66, 0xed,
	0x2a, 0x6b, 0xbb, 0x6e, 0x42, 0x51, 0x95, 0xa2, 0x2a, 0x28, 0x33, 0x7d, 0x78, 0x6a, 0x3c, 0x26,
	0xce, 0x4f, 0x03, 0xac, 0xeb, 0xed, 0xa5, 0xcf, 0xc1, 0xd4, 0x0d, 0x56, 0x1c, 0xf5, 0xce, 0x83,
	0x1d, 0xf3, 0xe0, 0x26, 0x63, 0x91, 0x64, 0xd1, 0x7b, 0x50, 0x9b, 0xad, 0x85, 0xc0, 0x40, 0x6a,
	0x79, 0x89, 0xf0, 0x6a, 0x12, 0x54, 0xa5, 0xd3, 0x87, 0x70, 0xa8, 0xb5, 0x0b, 0xdc, 0x0c, 0x36,
	0x0f, 0xb4, 0xfc, 0x02, 0xab, 0x2f, 0xf5, 0x78, 0x26, 0x51, 0x7a, 0x07, 0x40, 0x60, 0x84, 0x72,
	0x22, 0xb9, 0x8f, 0xf6, 0xbe, 0xc2, 0x94, 0x55, 0xe4, 0x9c, 0xfb, 0x48, 0xfb, 0x70, 0xe0, 0xa3,
	0xf4, 0xe6, 0x9e, 0xf4, 0xec, 0xa2, 0xea, 0xce, 0xf1, 0xae, 0x72, 0xcf, 0x12, 0xbc, 0xee, 0x52,
	0x9a, 0xde, 0xec, 0x42, 0xed, 0xca, 0xd5, 0x3f, 0x39, 0x79, 0x0c, 0x66, 0x62, 0x5f, 0x05, 0x4a,
	0x6f, 0x87, 0x83, 0xe1, 0xe8, 0xdd, 0xd0, 0xda, 0xa3, 0x26, 0x18, 0xa3, 0x81, 0x45, 0x68, 0x1d,
	0x60, 0x74, 0x71, 0xc2, 0x26, 0xa7, 0xfd, 0xb3, 0xfe, 0xb9, 0x65, 0x74, 0xbe, 0x91, 0xcc, 0xd2,
	0x8f, 0x51, 0x5c, 0xf2, 0x19, 0xd2, 0x31, 0x54, 0x5f, 0xa1, 0x4c, 0xc3, 0xf4, 0x28, 0x6f, 0x12,
	0x92, 0x8d, 0x69, 0xb6, 0xf2, 0x01, 0x7a, 0xbb, 0x9d, 0x3d, 0xfa, 0x01, 0x1a, 0x59, 0xd2, 0x5e,
	0x3c, 0xc0, 0x98, 0x3a, 0xbb, 0xd7, 0xf1, 0x6f, 0xc8, 0x7b, 0xa5, 0xf7, 0xc6, 0x6a, 0xfa, 0x95,
	0x90, 0xa9, 0xa9, 0x7e, 0x65, 0x8f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x84, 0x27, 0x52, 0x77,
	0xfa, 0x04, 0x00, 0x00,
}
