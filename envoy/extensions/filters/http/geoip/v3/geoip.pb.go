// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: envoy/extensions/filters/http/geoip/v3/geoip.proto

package geoipv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Geoip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If set, the :ref:`xff_num_trusted_hops <envoy_v3_api_field_extensions.filters.http.geoip.v3.Geoip.XffConfig.xff_num_trusted_hops>` field will be used to determine
	// trusted client address from “x-forwarded-for“ header.
	// Otherwise, the immediate downstream connection source address will be used.
	// [#next-free-field: 2]
	XffConfig *Geoip_XffConfig `protobuf:"bytes,1,opt,name=xff_config,json=xffConfig,proto3" json:"xff_config,omitempty"`
	// Geoip driver specific configuration which depends on the driver being instantiated.
	// See the geoip drivers for examples:
	//
	// - :ref:`MaxMindConfig <envoy_v3_api_msg_extensions.geoip_providers.maxmind.v3.MaxMindConfig>`
	// [#extension-category: envoy.geoip_providers]
	Provider *v3.TypedExtensionConfig `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *Geoip) Reset() {
	*x = Geoip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_geoip_v3_geoip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Geoip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Geoip) ProtoMessage() {}

func (x *Geoip) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_geoip_v3_geoip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Geoip.ProtoReflect.Descriptor instead.
func (*Geoip) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_geoip_v3_geoip_proto_rawDescGZIP(), []int{0}
}

func (x *Geoip) GetXffConfig() *Geoip_XffConfig {
	if x != nil {
		return x.XffConfig
	}
	return nil
}

func (x *Geoip) GetProvider() *v3.TypedExtensionConfig {
	if x != nil {
		return x.Provider
	}
	return nil
}

type Geoip_XffConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of additional ingress proxy hops from the right side of the
	// :ref:`config_http_conn_man_headers_x-forwarded-for` HTTP header to trust when
	// determining the origin client's IP address. The default is zero if this option
	// is not specified. See the documentation for
	// :ref:`config_http_conn_man_headers_x-forwarded-for` for more information.
	XffNumTrustedHops uint32 `protobuf:"varint,1,opt,name=xff_num_trusted_hops,json=xffNumTrustedHops,proto3" json:"xff_num_trusted_hops,omitempty"`
}

func (x *Geoip_XffConfig) Reset() {
	*x = Geoip_XffConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_geoip_v3_geoip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Geoip_XffConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Geoip_XffConfig) ProtoMessage() {}

func (x *Geoip_XffConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_geoip_v3_geoip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Geoip_XffConfig.ProtoReflect.Descriptor instead.
func (*Geoip_XffConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_geoip_v3_geoip_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Geoip_XffConfig) GetXffNumTrustedHops() uint32 {
	if x != nil {
		return x.XffNumTrustedHops
	}
	return 0
}

var File_envoy_extensions_filters_http_geoip_v3_geoip_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_geoip_v3_geoip_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x67, 0x65, 0x6f, 0x69, 0x70, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x2e, 0x76, 0x33, 0x1a, 0x24, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x05,
	0x47, 0x65, 0x6f, 0x69, 0x70, 0x12, 0x56, 0x0a, 0x0a, 0x78, 0x66, 0x66, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x2e,
	0x76, 0x33, 0x2e, 0x47, 0x65, 0x6f, 0x69, 0x70, 0x2e, 0x58, 0x66, 0x66, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x09, 0x78, 0x66, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x50, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a,
	0x3c, 0x0a, 0x09, 0x58, 0x66, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2f, 0x0a, 0x14,
	0x78, 0x66, 0x66, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x5f,
	0x68, 0x6f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x78, 0x66, 0x66, 0x4e,
	0x75, 0x6d, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x48, 0x6f, 0x70, 0x73, 0x42, 0xab, 0x01,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01,
	0x0a, 0x34, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x65,
	0x6f, 0x69, 0x70, 0x2e, 0x76, 0x33, 0x42, 0x0a, 0x47, 0x65, 0x6f, 0x69, 0x70, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x67, 0x65, 0x6f, 0x69, 0x70,
	0x2f, 0x76, 0x33, 0x3b, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_geoip_v3_geoip_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_geoip_v3_geoip_proto_rawDescData = file_envoy_extensions_filters_http_geoip_v3_geoip_proto_rawDesc
)

func file_envoy_extensions_filters_http_geoip_v3_geoip_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_geoip_v3_geoip_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_geoip_v3_geoip_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_geoip_v3_geoip_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_geoip_v3_geoip_proto_rawDescData
}

var file_envoy_extensions_filters_http_geoip_v3_geoip_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_http_geoip_v3_geoip_proto_goTypes = []interface{}{
	(*Geoip)(nil),                   // 0: envoy.extensions.filters.http.geoip.v3.Geoip
	(*Geoip_XffConfig)(nil),         // 1: envoy.extensions.filters.http.geoip.v3.Geoip.XffConfig
	(*v3.TypedExtensionConfig)(nil), // 2: envoy.config.core.v3.TypedExtensionConfig
}
var file_envoy_extensions_filters_http_geoip_v3_geoip_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.http.geoip.v3.Geoip.xff_config:type_name -> envoy.extensions.filters.http.geoip.v3.Geoip.XffConfig
	2, // 1: envoy.extensions.filters.http.geoip.v3.Geoip.provider:type_name -> envoy.config.core.v3.TypedExtensionConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_geoip_v3_geoip_proto_init() }
func file_envoy_extensions_filters_http_geoip_v3_geoip_proto_init() {
	if File_envoy_extensions_filters_http_geoip_v3_geoip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_geoip_v3_geoip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Geoip); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_filters_http_geoip_v3_geoip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Geoip_XffConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_geoip_v3_geoip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_geoip_v3_geoip_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_geoip_v3_geoip_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_geoip_v3_geoip_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_geoip_v3_geoip_proto = out.File
	file_envoy_extensions_filters_http_geoip_v3_geoip_proto_rawDesc = nil
	file_envoy_extensions_filters_http_geoip_v3_geoip_proto_goTypes = nil
	file_envoy_extensions_filters_http_geoip_v3_geoip_proto_depIdxs = nil
}
