// Code generated by protoc-gen-go.
// source: player.proto
// DO NOT EDIT!

/*
Package player is a generated protocol buffer package.

It is generated from these files:
	player.proto

It has these top-level messages:
	CardIn
	MeterUpdate
	CardOut
	PointUpdate
	Status
*/
package player

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

type CardIn struct {
	MachNumber string `protobuf:"bytes,1,opt,name=machNumber" json:"machNumber,omitempty"`
	CardNumber int32  `protobuf:"varint,2,opt,name=cardNumber" json:"cardNumber,omitempty"`
	StartTime  string `protobuf:"bytes,3,opt,name=startTime" json:"startTime,omitempty"`
}

func (m *CardIn) Reset()                    { *m = CardIn{} }
func (m *CardIn) String() string            { return proto.CompactTextString(m) }
func (*CardIn) ProtoMessage()               {}
func (*CardIn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CardIn) GetMachNumber() string {
	if m != nil {
		return m.MachNumber
	}
	return ""
}

func (m *CardIn) GetCardNumber() int32 {
	if m != nil {
		return m.CardNumber
	}
	return 0
}

func (m *CardIn) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

type MeterUpdate struct {
	MachNumber string `protobuf:"bytes,1,opt,name=machNumber" json:"machNumber,omitempty"`
	CardNumber int32  `protobuf:"varint,2,opt,name=cardNumber" json:"cardNumber,omitempty"`
	CoinIn     int32  `protobuf:"varint,3,opt,name=coinIn" json:"coinIn,omitempty"`
	CoinOut    int32  `protobuf:"varint,4,opt,name=coinOut" json:"coinOut,omitempty"`
	Jackpot    int32  `protobuf:"varint,5,opt,name=jackpot" json:"jackpot,omitempty"`
}

func (m *MeterUpdate) Reset()                    { *m = MeterUpdate{} }
func (m *MeterUpdate) String() string            { return proto.CompactTextString(m) }
func (*MeterUpdate) ProtoMessage()               {}
func (*MeterUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MeterUpdate) GetMachNumber() string {
	if m != nil {
		return m.MachNumber
	}
	return ""
}

func (m *MeterUpdate) GetCardNumber() int32 {
	if m != nil {
		return m.CardNumber
	}
	return 0
}

func (m *MeterUpdate) GetCoinIn() int32 {
	if m != nil {
		return m.CoinIn
	}
	return 0
}

func (m *MeterUpdate) GetCoinOut() int32 {
	if m != nil {
		return m.CoinOut
	}
	return 0
}

func (m *MeterUpdate) GetJackpot() int32 {
	if m != nil {
		return m.Jackpot
	}
	return 0
}

type CardOut struct {
	MachNumber string `protobuf:"bytes,1,opt,name=machNumber" json:"machNumber,omitempty"`
	CardNumber int32  `protobuf:"varint,2,opt,name=cardNumber" json:"cardNumber,omitempty"`
	CoinIn     int32  `protobuf:"varint,3,opt,name=coinIn" json:"coinIn,omitempty"`
	CoinOut    int32  `protobuf:"varint,4,opt,name=coinOut" json:"coinOut,omitempty"`
	Jackpot    int32  `protobuf:"varint,5,opt,name=jackpot" json:"jackpot,omitempty"`
	EndTime    int32  `protobuf:"varint,6,opt,name=endTime" json:"endTime,omitempty"`
}

func (m *CardOut) Reset()                    { *m = CardOut{} }
func (m *CardOut) String() string            { return proto.CompactTextString(m) }
func (*CardOut) ProtoMessage()               {}
func (*CardOut) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CardOut) GetMachNumber() string {
	if m != nil {
		return m.MachNumber
	}
	return ""
}

func (m *CardOut) GetCardNumber() int32 {
	if m != nil {
		return m.CardNumber
	}
	return 0
}

func (m *CardOut) GetCoinIn() int32 {
	if m != nil {
		return m.CoinIn
	}
	return 0
}

func (m *CardOut) GetCoinOut() int32 {
	if m != nil {
		return m.CoinOut
	}
	return 0
}

func (m *CardOut) GetJackpot() int32 {
	if m != nil {
		return m.Jackpot
	}
	return 0
}

func (m *CardOut) GetEndTime() int32 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

type PointUpdate struct {
	MachNumber      string `protobuf:"bytes,1,opt,name=machNumber" json:"machNumber,omitempty"`
	CardNumber      int32  `protobuf:"varint,2,opt,name=cardNumber" json:"cardNumber,omitempty"`
	NewPointBalance int32  `protobuf:"varint,3,opt,name=newPointBalance" json:"newPointBalance,omitempty"`
}

func (m *PointUpdate) Reset()                    { *m = PointUpdate{} }
func (m *PointUpdate) String() string            { return proto.CompactTextString(m) }
func (*PointUpdate) ProtoMessage()               {}
func (*PointUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PointUpdate) GetMachNumber() string {
	if m != nil {
		return m.MachNumber
	}
	return ""
}

func (m *PointUpdate) GetCardNumber() int32 {
	if m != nil {
		return m.CardNumber
	}
	return 0
}

func (m *PointUpdate) GetNewPointBalance() int32 {
	if m != nil {
		return m.NewPointBalance
	}
	return 0
}

type Status struct {
	Message     string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Points      int32  `protobuf:"varint,2,opt,name=points" json:"points,omitempty"`
	DisplayName string `protobuf:"bytes,3,opt,name=displayName" json:"displayName,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Status) GetPoints() int32 {
	if m != nil {
		return m.Points
	}
	return 0
}

func (m *Status) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func init() {
	proto.RegisterType((*CardIn)(nil), "player.CardIn")
	proto.RegisterType((*MeterUpdate)(nil), "player.MeterUpdate")
	proto.RegisterType((*CardOut)(nil), "player.CardOut")
	proto.RegisterType((*PointUpdate)(nil), "player.PointUpdate")
	proto.RegisterType((*Status)(nil), "player.Status")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PlayerTracking service

type PlayerTrackingClient interface {
	CardInEvent(ctx context.Context, in *CardIn, opts ...grpc.CallOption) (*Status, error)
	MeterUpdateEvent(ctx context.Context, in *MeterUpdate, opts ...grpc.CallOption) (*PointUpdate, error)
	CardOutEvent(ctx context.Context, in *CardOut, opts ...grpc.CallOption) (*Status, error)
}

type playerTrackingClient struct {
	cc *grpc.ClientConn
}

func NewPlayerTrackingClient(cc *grpc.ClientConn) PlayerTrackingClient {
	return &playerTrackingClient{cc}
}

func (c *playerTrackingClient) CardInEvent(ctx context.Context, in *CardIn, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/player.PlayerTracking/CardInEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackingClient) MeterUpdateEvent(ctx context.Context, in *MeterUpdate, opts ...grpc.CallOption) (*PointUpdate, error) {
	out := new(PointUpdate)
	err := grpc.Invoke(ctx, "/player.PlayerTracking/MeterUpdateEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackingClient) CardOutEvent(ctx context.Context, in *CardOut, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/player.PlayerTracking/CardOutEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PlayerTracking service

type PlayerTrackingServer interface {
	CardInEvent(context.Context, *CardIn) (*Status, error)
	MeterUpdateEvent(context.Context, *MeterUpdate) (*PointUpdate, error)
	CardOutEvent(context.Context, *CardOut) (*Status, error)
}

func RegisterPlayerTrackingServer(s *grpc.Server, srv PlayerTrackingServer) {
	s.RegisterService(&_PlayerTracking_serviceDesc, srv)
}

func _PlayerTracking_CardInEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackingServer).CardInEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player.PlayerTracking/CardInEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackingServer).CardInEvent(ctx, req.(*CardIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracking_MeterUpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeterUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackingServer).MeterUpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player.PlayerTracking/MeterUpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackingServer).MeterUpdateEvent(ctx, req.(*MeterUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracking_CardOutEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardOut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackingServer).CardOutEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player.PlayerTracking/CardOutEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackingServer).CardOutEvent(ctx, req.(*CardOut))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlayerTracking_serviceDesc = grpc.ServiceDesc{
	ServiceName: "player.PlayerTracking",
	HandlerType: (*PlayerTrackingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CardInEvent",
			Handler:    _PlayerTracking_CardInEvent_Handler,
		},
		{
			MethodName: "MeterUpdateEvent",
			Handler:    _PlayerTracking_MeterUpdateEvent_Handler,
		},
		{
			MethodName: "CardOutEvent",
			Handler:    _PlayerTracking_CardOutEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "player.proto",
}

func init() { proto.RegisterFile("player.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x53, 0x4f, 0x4f, 0xfa, 0x40,
	0x14, 0xa4, 0xbf, 0x9f, 0x94, 0xf0, 0x4a, 0xc0, 0xac, 0x89, 0x69, 0x88, 0x31, 0x64, 0x4f, 0x9c,
	0x30, 0xea, 0xdd, 0x83, 0xc6, 0x03, 0x07, 0x81, 0x20, 0xde, 0xbc, 0x3c, 0xda, 0x15, 0x2b, 0x74,
	0xdb, 0xec, 0x6e, 0x25, 0x7e, 0x14, 0x3f, 0x85, 0x17, 0x3f, 0xa0, 0xd9, 0x3f, 0xd4, 0x8d, 0x1e,
	0xf5, 0xe0, 0xad, 0x33, 0xf3, 0x5e, 0xe7, 0x65, 0x32, 0x0b, 0x9d, 0x72, 0x83, 0x2f, 0x4c, 0x8c,
	0x4a, 0x51, 0xa8, 0x82, 0x84, 0x16, 0xd1, 0x07, 0x08, 0xaf, 0x50, 0xa4, 0x63, 0x4e, 0x8e, 0x01,
	0x72, 0x4c, 0x1e, 0x27, 0x55, 0xbe, 0x64, 0x22, 0x0e, 0x06, 0xc1, 0xb0, 0x3d, 0xf7, 0x18, 0xad,
	0x27, 0x28, 0x52, 0xa7, 0xff, 0x1b, 0x04, 0xc3, 0xe6, 0xdc, 0x63, 0xc8, 0x11, 0xb4, 0xa5, 0x42,
	0xa1, 0x16, 0x59, 0xce, 0xe2, 0xff, 0x66, 0xfd, 0x93, 0xa0, 0xaf, 0x01, 0x44, 0x37, 0x4c, 0x31,
	0x71, 0x57, 0xa6, 0xa8, 0xd8, 0x8f, 0xdd, 0x0e, 0x21, 0x4c, 0x8a, 0x8c, 0x8f, 0xb9, 0xb1, 0x6a,
	0xce, 0x1d, 0x22, 0x31, 0xb4, 0xf4, 0xd7, 0xb4, 0x52, 0xf1, 0x9e, 0x11, 0x76, 0x50, 0x2b, 0x4f,
	0x98, 0xac, 0xcb, 0x42, 0xc5, 0x4d, 0xab, 0x38, 0x48, 0xdf, 0x02, 0x68, 0xe9, 0x10, 0xf4, 0xd4,
	0x1f, 0xba, 0x4b, 0x2b, 0x8c, 0xa7, 0x26, 0xcf, 0xd0, 0x2a, 0x0e, 0xd2, 0x2d, 0x44, 0xb3, 0x22,
	0xe3, 0xea, 0x97, 0xc2, 0x1c, 0x42, 0x8f, 0xb3, 0xad, 0xf9, 0xe3, 0x25, 0x6e, 0x90, 0x27, 0xcc,
	0x5d, 0xff, 0x95, 0xa6, 0xf7, 0x10, 0xde, 0x2a, 0x54, 0x95, 0xd4, 0xc7, 0xe5, 0x4c, 0x4a, 0x5c,
	0x31, 0x67, 0xb8, 0x83, 0x3a, 0x82, 0x52, 0xef, 0x48, 0xe7, 0xe4, 0x10, 0x19, 0x40, 0x94, 0x66,
	0x52, 0xf7, 0x6e, 0x82, 0x75, 0x45, 0x7c, 0xea, 0xec, 0x3d, 0x80, 0xee, 0xcc, 0xf4, 0x72, 0x21,
	0x30, 0x59, 0x67, 0x7c, 0x45, 0x4e, 0x20, 0xb2, 0xfd, 0xbc, 0x7e, 0x66, 0x5c, 0x91, 0xee, 0xc8,
	0xb5, 0xd8, 0x92, 0xfd, 0x1a, 0xdb, 0xab, 0x68, 0x83, 0x5c, 0xc0, 0xbe, 0xd7, 0x33, 0xbb, 0x75,
	0xb0, 0x9b, 0xf2, 0x94, 0x7e, 0x4d, 0x7a, 0x49, 0xd2, 0x06, 0x39, 0x85, 0x8e, 0xeb, 0x82, 0xdd,
	0xed, 0xf9, 0x8e, 0xd3, 0x4a, 0x7d, 0xb7, 0x5c, 0x86, 0xe6, 0x49, 0x9d, 0x7f, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x21, 0xc0, 0xda, 0x28, 0x62, 0x03, 0x00, 0x00,
}
