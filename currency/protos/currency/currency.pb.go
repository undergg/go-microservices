// Code generated by protoc-gen-go. DO NOT EDIT.
// source: currency.proto

package currency

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

// Currencies is an enum which represents the allowed currencies for the API
type Currencies int32

const (
	Currencies_EUR Currencies = 0
	Currencies_USD Currencies = 1
	Currencies_JPY Currencies = 2
	Currencies_BGN Currencies = 3
	Currencies_CZK Currencies = 4
	Currencies_DKK Currencies = 5
	Currencies_GBP Currencies = 6
	Currencies_HUF Currencies = 7
	Currencies_PLN Currencies = 8
	Currencies_RON Currencies = 9
	Currencies_SEK Currencies = 10
	Currencies_CHF Currencies = 11
	Currencies_ISK Currencies = 12
	Currencies_NOK Currencies = 13
	Currencies_HRK Currencies = 14
	Currencies_RUB Currencies = 15
	Currencies_TRY Currencies = 16
	Currencies_AUD Currencies = 17
	Currencies_BRL Currencies = 18
	Currencies_CAD Currencies = 19
	Currencies_CNY Currencies = 20
	Currencies_HKD Currencies = 21
	Currencies_IDR Currencies = 22
	Currencies_ILS Currencies = 23
	Currencies_INR Currencies = 24
	Currencies_KRW Currencies = 25
	Currencies_MXN Currencies = 26
	Currencies_MYR Currencies = 27
	Currencies_NZD Currencies = 28
	Currencies_PHP Currencies = 29
	Currencies_SGD Currencies = 30
	Currencies_THB Currencies = 31
	Currencies_ZAR Currencies = 32
)

var Currencies_name = map[int32]string{
	0:  "EUR",
	1:  "USD",
	2:  "JPY",
	3:  "BGN",
	4:  "CZK",
	5:  "DKK",
	6:  "GBP",
	7:  "HUF",
	8:  "PLN",
	9:  "RON",
	10: "SEK",
	11: "CHF",
	12: "ISK",
	13: "NOK",
	14: "HRK",
	15: "RUB",
	16: "TRY",
	17: "AUD",
	18: "BRL",
	19: "CAD",
	20: "CNY",
	21: "HKD",
	22: "IDR",
	23: "ILS",
	24: "INR",
	25: "KRW",
	26: "MXN",
	27: "MYR",
	28: "NZD",
	29: "PHP",
	30: "SGD",
	31: "THB",
	32: "ZAR",
}

var Currencies_value = map[string]int32{
	"EUR": 0,
	"USD": 1,
	"JPY": 2,
	"BGN": 3,
	"CZK": 4,
	"DKK": 5,
	"GBP": 6,
	"HUF": 7,
	"PLN": 8,
	"RON": 9,
	"SEK": 10,
	"CHF": 11,
	"ISK": 12,
	"NOK": 13,
	"HRK": 14,
	"RUB": 15,
	"TRY": 16,
	"AUD": 17,
	"BRL": 18,
	"CAD": 19,
	"CNY": 20,
	"HKD": 21,
	"IDR": 22,
	"ILS": 23,
	"INR": 24,
	"KRW": 25,
	"MXN": 26,
	"MYR": 27,
	"NZD": 28,
	"PHP": 29,
	"SGD": 30,
	"THB": 31,
	"ZAR": 32,
}

func (x Currencies) String() string {
	return proto.EnumName(Currencies_name, int32(x))
}

func (Currencies) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{0}
}

type RateRequest struct {
	Base                 Currencies `protobuf:"varint,1,opt,name=Base,json=base,proto3,enum=Currencies" json:"Base,omitempty"`
	Destination          Currencies `protobuf:"varint,2,opt,name=Destination,json=destination,proto3,enum=Currencies" json:"Destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RateRequest) Reset()         { *m = RateRequest{} }
func (m *RateRequest) String() string { return proto.CompactTextString(m) }
func (*RateRequest) ProtoMessage()    {}
func (*RateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{0}
}

func (m *RateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateRequest.Unmarshal(m, b)
}
func (m *RateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateRequest.Marshal(b, m, deterministic)
}
func (m *RateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateRequest.Merge(m, src)
}
func (m *RateRequest) XXX_Size() int {
	return xxx_messageInfo_RateRequest.Size(m)
}
func (m *RateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateRequest proto.InternalMessageInfo

func (m *RateRequest) GetBase() Currencies {
	if m != nil {
		return m.Base
	}
	return Currencies_EUR
}

func (m *RateRequest) GetDestination() Currencies {
	if m != nil {
		return m.Destination
	}
	return Currencies_EUR
}

type RateResponse struct {
	Rate                 float64  `protobuf:"fixed64,1,opt,name=Rate,json=rate,proto3" json:"Rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateResponse) Reset()         { *m = RateResponse{} }
func (m *RateResponse) String() string { return proto.CompactTextString(m) }
func (*RateResponse) ProtoMessage()    {}
func (*RateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{1}
}

func (m *RateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateResponse.Unmarshal(m, b)
}
func (m *RateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateResponse.Marshal(b, m, deterministic)
}
func (m *RateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateResponse.Merge(m, src)
}
func (m *RateResponse) XXX_Size() int {
	return xxx_messageInfo_RateResponse.Size(m)
}
func (m *RateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateResponse proto.InternalMessageInfo

func (m *RateResponse) GetRate() float64 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func init() {
	proto.RegisterEnum("Currencies", Currencies_name, Currencies_value)
	proto.RegisterType((*RateRequest)(nil), "RateRequest")
	proto.RegisterType((*RateResponse)(nil), "RateResponse")
}

func init() {
	proto.RegisterFile("currency.proto", fileDescriptor_d3dc60ed002193ea)
}

var fileDescriptor_d3dc60ed002193ea = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd1, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0x06, 0x70, 0xcb, 0xad, 0x97, 0xeb, 0x94, 0x8b, 0x9f, 0xe3, 0x3f, 0xbc, 0xfe, 0x81, 0xb0,
	0x30, 0xc6, 0xc4, 0x2e, 0xf0, 0x09, 0x28, 0x03, 0xad, 0x0e, 0x96, 0x66, 0x6a, 0xa3, 0x25, 0x71,
	0x51, 0x70, 0x16, 0x6c, 0x5a, 0x6c, 0xcb, 0x82, 0x07, 0xf3, 0xfd, 0xcc, 0x39, 0x98, 0x40, 0x72,
	0x77, 0xbf, 0xa6, 0xf3, 0x9d, 0x7c, 0x39, 0x47, 0xf4, 0xb7, 0x87, 0xba, 0xb6, 0xe5, 0xf6, 0xe8,
	0xef, 0xeb, 0xaa, 0xad, 0xc6, 0xbf, 0x84, 0x67, 0x8a, 0xd6, 0x1a, 0xfb, 0xe7, 0x60, 0x9b, 0x56,
	0x0e, 0x85, 0x1b, 0x14, 0x8d, 0x1d, 0x38, 0x23, 0xe7, 0x43, 0x7f, 0xe2, 0xf9, 0xb3, 0xd3, 0xeb,
	0x9d, 0x6d, 0x8c, 0xbb, 0x29, 0x1a, 0x2b, 0x3f, 0x09, 0x4f, 0xd9, 0xa6, 0xdd, 0x95, 0x45, 0xbb,
	0xab, 0xca, 0x41, 0xe7, 0xfe, 0x3b, 0xef, 0xf7, 0xf9, 0xff, 0x78, 0x2c, 0x7a, 0xa7, 0xf1, 0xcd,
	0xbe, 0x2a, 0x1b, 0x2b, 0xa5, 0x70, 0xe9, 0x9b, 0xe7, 0x3b, 0xc6, 0xad, 0x8b, 0xd6, 0x7e, 0xfc,
	0xdb, 0x11, 0xe2, 0x9c, 0x97, 0x5d, 0x71, 0x35, 0xcf, 0x0c, 0x1e, 0x10, 0xb2, 0x54, 0xc1, 0x21,
	0x7c, 0x4d, 0x72, 0x74, 0x08, 0x41, 0x18, 0xe3, 0x8a, 0x30, 0x5b, 0x6b, 0xb8, 0x04, 0xa5, 0x35,
	0x1e, 0x12, 0xc2, 0x20, 0xc1, 0x35, 0x21, 0xca, 0x16, 0xe8, 0x12, 0x92, 0x65, 0x8c, 0x1b, 0x82,
	0x59, 0xc5, 0x78, 0x44, 0x48, 0xe7, 0x1a, 0x82, 0xe3, 0xd1, 0x02, 0x1e, 0xe1, 0x4b, 0xaa, 0xd1,
	0x23, 0xc4, 0x2b, 0x8d, 0x5b, 0x8e, 0x1b, 0x8d, 0x3e, 0xa7, 0xb2, 0x00, 0x8f, 0x09, 0xdf, 0x4d,
	0x0e, 0x10, 0xa6, 0x99, 0xc2, 0x13, 0xae, 0x61, 0x96, 0x90, 0x3c, 0x67, 0xaa, 0xf0, 0x94, 0x11,
	0xe7, 0x78, 0xc6, 0x71, 0xad, 0xf0, 0x9c, 0x27, 0x2b, 0x83, 0x17, 0x8c, 0x65, 0x8a, 0x97, 0x8c,
	0xd8, 0x60, 0x40, 0xd0, 0xe6, 0x07, 0x5e, 0x11, 0xbe, 0xfd, 0x8c, 0x71, 0xc7, 0xc8, 0x0d, 0x5e,
	0x73, 0x8d, 0xb5, 0xc2, 0x1b, 0x2e, 0x1f, 0x25, 0x78, 0xcb, 0x9d, 0x43, 0x85, 0x77, 0x5c, 0x23,
	0x0a, 0x30, 0x24, 0xac, 0xa7, 0x06, 0xa3, 0xc9, 0x44, 0xdc, 0xfc, 0x5f, 0xdb, 0x51, 0xbe, 0x17,
	0xdd, 0xd0, 0xb6, 0xb4, 0x5a, 0xd9, 0xf3, 0x2f, 0x0e, 0x7a, 0x77, 0xeb, 0x5f, 0xee, 0x7f, 0x73,
	0xcd, 0x57, 0xff, 0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x98, 0x51, 0x7e, 0x5a, 0x07, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrencyClient is the client API for Currency service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrencyClient interface {
	GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
}

type currencyClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyClient(cc grpc.ClientConnInterface) CurrencyClient {
	return &currencyClient{cc}
}

func (c *currencyClient) GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/Currency/GetRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyServer is the server API for Currency service.
type CurrencyServer interface {
	GetRate(context.Context, *RateRequest) (*RateResponse, error)
}

// UnimplementedCurrencyServer can be embedded to have forward compatible implementations.
type UnimplementedCurrencyServer struct {
}

func (*UnimplementedCurrencyServer) GetRate(ctx context.Context, req *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}

func RegisterCurrencyServer(s *grpc.Server, srv CurrencyServer) {
	s.RegisterService(&_Currency_serviceDesc, srv)
}

func _Currency_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Currency/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).GetRate(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Currency_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Currency",
	HandlerType: (*CurrencyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _Currency_GetRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "currency.proto",
}
