// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datamover.proto

/*
Package datamover is a generated protocol buffer package.

It is generated from these files:
	datamover.proto

It has these top-level messages:
	StoreFile
	RetrieveFile
	MoveStatus
*/
package datamover

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

// Source is full path of the customer source side of the file
// Destination is full path of the Campaign Storage file.
type StoreFile struct {
	SourcePath      string `protobuf:"bytes,1,opt,name=sourcePath" json:"sourcePath,omitempty"`
	DestinationPath string `protobuf:"bytes,2,opt,name=destinationPath" json:"destinationPath,omitempty"`
}

func (m *StoreFile) Reset()                    { *m = StoreFile{} }
func (m *StoreFile) String() string            { return proto.CompactTextString(m) }
func (*StoreFile) ProtoMessage()               {}
func (*StoreFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StoreFile) GetSourcePath() string {
	if m != nil {
		return m.SourcePath
	}
	return ""
}

func (m *StoreFile) GetDestinationPath() string {
	if m != nil {
		return m.DestinationPath
	}
	return ""
}

type RetrieveFile struct {
	SourcePath      string `protobuf:"bytes,1,opt,name=sourcePath" json:"sourcePath,omitempty"`
	DestinationPath string `protobuf:"bytes,2,opt,name=destinationPath" json:"destinationPath,omitempty"`
}

func (m *RetrieveFile) Reset()                    { *m = RetrieveFile{} }
func (m *RetrieveFile) String() string            { return proto.CompactTextString(m) }
func (*RetrieveFile) ProtoMessage()               {}
func (*RetrieveFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RetrieveFile) GetSourcePath() string {
	if m != nil {
		return m.SourcePath
	}
	return ""
}

func (m *RetrieveFile) GetDestinationPath() string {
	if m != nil {
		return m.DestinationPath
	}
	return ""
}

// The response message containing the greetings
type MoveStatus struct {
	Status int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *MoveStatus) Reset()                    { *m = MoveStatus{} }
func (m *MoveStatus) String() string            { return proto.CompactTextString(m) }
func (*MoveStatus) ProtoMessage()               {}
func (*MoveStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MoveStatus) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*StoreFile)(nil), "datamover.StoreFile")
	proto.RegisterType((*RetrieveFile)(nil), "datamover.RetrieveFile")
	proto.RegisterType((*MoveStatus)(nil), "datamover.MoveStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FileMover service

type FileMoverClient interface {
	// Store a file in Campaign Storage
	Store(ctx context.Context, in *StoreFile, opts ...grpc.CallOption) (*MoveStatus, error)
	// Retrieve a file from Campaign Storage
	Retrieve(ctx context.Context, in *RetrieveFile, opts ...grpc.CallOption) (*MoveStatus, error)
}

type fileMoverClient struct {
	cc *grpc.ClientConn
}

func NewFileMoverClient(cc *grpc.ClientConn) FileMoverClient {
	return &fileMoverClient{cc}
}

func (c *fileMoverClient) Store(ctx context.Context, in *StoreFile, opts ...grpc.CallOption) (*MoveStatus, error) {
	out := new(MoveStatus)
	err := grpc.Invoke(ctx, "/datamover.FileMover/Store", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileMoverClient) Retrieve(ctx context.Context, in *RetrieveFile, opts ...grpc.CallOption) (*MoveStatus, error) {
	out := new(MoveStatus)
	err := grpc.Invoke(ctx, "/datamover.FileMover/Retrieve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FileMover service

type FileMoverServer interface {
	// Store a file in Campaign Storage
	Store(context.Context, *StoreFile) (*MoveStatus, error)
	// Retrieve a file from Campaign Storage
	Retrieve(context.Context, *RetrieveFile) (*MoveStatus, error)
}

func RegisterFileMoverServer(s *grpc.Server, srv FileMoverServer) {
	s.RegisterService(&_FileMover_serviceDesc, srv)
}

func _FileMover_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileMoverServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datamover.FileMover/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileMoverServer).Store(ctx, req.(*StoreFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileMover_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileMoverServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datamover.FileMover/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileMoverServer).Retrieve(ctx, req.(*RetrieveFile))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileMover_serviceDesc = grpc.ServiceDesc{
	ServiceName: "datamover.FileMover",
	HandlerType: (*FileMoverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _FileMover_Store_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _FileMover_Retrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datamover.proto",
}

func init() { proto.RegisterFile("datamover.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x49, 0x2c, 0x49,
	0xcc, 0xcd, 0x2f, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x85, 0x72, 0x71, 0x06, 0x97, 0xe4, 0x17, 0xa5, 0xba, 0x65, 0xe6, 0xa4, 0x0a, 0xc9, 0x71, 0x71,
	0x15, 0xe7, 0x97, 0x16, 0x25, 0xa7, 0x06, 0x24, 0x96, 0x64, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0x21, 0x89, 0x08, 0x69, 0x00, 0x8d, 0x4a, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0x2c, 0xc9, 0xcc,
	0xcf, 0x03, 0x2b, 0x62, 0x02, 0x2b, 0x42, 0x17, 0x56, 0x8a, 0xe0, 0xe2, 0x09, 0x4a, 0x2d, 0x29,
	0xca, 0x4c, 0x2d, 0xa3, 0xb6, 0xc9, 0x2a, 0x5c, 0x5c, 0xbe, 0x40, 0x97, 0x07, 0x97, 0x24, 0x96,
	0x94, 0x16, 0x0b, 0x89, 0x71, 0xb1, 0x15, 0x83, 0x59, 0x60, 0x33, 0x59, 0x83, 0xa0, 0x3c, 0xa3,
	0x46, 0x46, 0x2e, 0x4e, 0x90, 0xc5, 0x20, 0xa5, 0x45, 0x42, 0x66, 0x5c, 0xac, 0x60, 0x4f, 0x0a,
	0x89, 0xe8, 0x21, 0x82, 0x02, 0xee, 0x6d, 0x29, 0x51, 0x24, 0x51, 0x84, 0xd9, 0x4a, 0x0c, 0x42,
	0x36, 0x5c, 0x1c, 0x30, 0x5f, 0x08, 0x89, 0x23, 0x29, 0x42, 0xf6, 0x1a, 0x4e, 0xdd, 0x49, 0x6c,
	0xe0, 0xc0, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x90, 0x9b, 0x3f, 0x17, 0x7f, 0x01, 0x00,
	0x00,
}