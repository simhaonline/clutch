// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/v1/annotations.proto

package apiv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

var E_Reference = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*Reference)(nil),
	Field:         58901,
	Name:          "clutch.api.v1.reference",
	Tag:           "bytes,58901,opt,name=reference",
	Filename:      "api/v1/annotations.proto",
}

var E_Id = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*Identifier)(nil),
	Field:         58902,
	Name:          "clutch.api.v1.id",
	Tag:           "bytes,58902,opt,name=id",
	Filename:      "api/v1/annotations.proto",
}

var E_Action = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*Action)(nil),
	Field:         58901,
	Name:          "clutch.api.v1.action",
	Tag:           "bytes,58901,opt,name=action",
	Filename:      "api/v1/annotations.proto",
}

func init() {
	proto.RegisterExtension(E_Reference)
	proto.RegisterExtension(E_Id)
	proto.RegisterExtension(E_Action)
}

func init() {
	proto.RegisterFile("api/v1/annotations.proto", fileDescriptor_2142d87d3bdfd889)
}

var fileDescriptor_2142d87d3bdfd889 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xce, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x71, 0x5c, 0x71, 0xc5, 0x8a, 0x97, 0x8a, 0x50, 0xf7, 0xa0, 0x7b, 0xf4, 0x34, 0xa5,
	0x7a, 0xeb, 0x4d, 0x6f, 0x82, 0xb2, 0x90, 0x9b, 0xde, 0x66, 0x93, 0x69, 0x3b, 0xb0, 0x66, 0x42,
	0x92, 0xed, 0x9b, 0xe8, 0x53, 0xf9, 0x50, 0x62, 0xda, 0x22, 0x16, 0xc1, 0x6b, 0xc2, 0xff, 0xf7,
	0x4d, 0x56, 0xa0, 0xe3, 0xb2, 0xaf, 0x4a, 0xb4, 0x56, 0x22, 0x46, 0x16, 0x1b, 0xc0, 0x79, 0x89,
	0x92, 0x9f, 0xe9, 0xdd, 0x3e, 0xea, 0x0e, 0xd0, 0x31, 0xf4, 0xd5, 0x6a, 0xdd, 0x8a, 0xb4, 0x3b,
	0x2a, 0xd3, 0xe7, 0x76, 0xdf, 0x94, 0x86, 0x82, 0xf6, 0xec, 0xa2, 0xf8, 0x21, 0x58, 0x9d, 0x8f,
	0x54, 0xd0, 0x1d, 0xbd, 0xe1, 0xf0, 0x58, 0xbf, 0x64, 0x27, 0x9e, 0x1a, 0xf2, 0x64, 0x35, 0xe5,
	0xd7, 0x30, 0x20, 0x30, 0x21, 0xf0, 0x4c, 0x21, 0x60, 0x4b, 0x1b, 0x97, 0x96, 0x8b, 0xf7, 0xcf,
	0xc3, 0xf5, 0xc1, 0xcd, 0xe9, 0x6d, 0x01, 0xbf, 0xc6, 0x41, 0x4d, 0x84, 0xfa, 0xd1, 0xea, 0xa7,
	0x6c, 0xc1, 0xe6, 0x7f, 0xf3, 0x63, 0x34, 0x2f, 0x67, 0xe6, 0xa3, 0x21, 0x1b, 0xb9, 0x61, 0xf2,
	0x6a, 0xc1, 0xa6, 0xde, 0x64, 0x4b, 0xd4, 0xdf, 0x45, 0x7e, 0xf5, 0x87, 0x18, 0x3b, 0x31, 0xf3,
	0x23, 0x2f, 0x66, 0xe0, 0x7d, 0xca, 0xd5, 0xc8, 0x3c, 0x1c, 0xbf, 0x1e, 0xa1, 0xe3, 0xbe, 0xda,
	0x2e, 0x93, 0x73, 0xf7, 0x15, 0x00, 0x00, 0xff, 0xff, 0x81, 0x89, 0xe9, 0xb1, 0x6b, 0x01, 0x00,
	0x00,
}
