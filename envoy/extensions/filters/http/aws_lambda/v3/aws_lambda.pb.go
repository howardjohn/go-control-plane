// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: envoy/extensions/filters/http/aws_lambda/v3/aws_lambda.proto

package aws_lambdav3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

type Config_InvocationMode int32

const (
	// This is the more common mode of invocation, in which Lambda responds after it has completed the function. In
	// this mode the output of the Lambda function becomes the response of the HTTP request.
	Config_SYNCHRONOUS Config_InvocationMode = 0
	// In this mode Lambda responds immediately but continues to process the function asynchronously. This mode can be
	// used to signal events for example. In this mode, Lambda responds with an acknowledgment that it received the
	// call which is translated to an HTTP 200 OK by the filter.
	Config_ASYNCHRONOUS Config_InvocationMode = 1
)

// Enum value maps for Config_InvocationMode.
var (
	Config_InvocationMode_name = map[int32]string{
		0: "SYNCHRONOUS",
		1: "ASYNCHRONOUS",
	}
	Config_InvocationMode_value = map[string]int32{
		"SYNCHRONOUS":  0,
		"ASYNCHRONOUS": 1,
	}
)

func (x Config_InvocationMode) Enum() *Config_InvocationMode {
	p := new(Config_InvocationMode)
	*p = x
	return p
}

func (x Config_InvocationMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Config_InvocationMode) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_enumTypes[0].Descriptor()
}

func (Config_InvocationMode) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_enumTypes[0]
}

func (x Config_InvocationMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Config_InvocationMode.Descriptor instead.
func (Config_InvocationMode) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_rawDescGZIP(), []int{0, 0}
}

// AWS Lambda filter config
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ARN of the AWS Lambda to invoke when the filter is engaged
	// Must be in the following format:
	// arn:<partition>:lambda:<region>:<account-number>:function:<function-name>
	Arn string `protobuf:"bytes,1,opt,name=arn,proto3" json:"arn,omitempty"`
	// Whether to transform the request (headers and body) to a JSON payload or pass it as is.
	PayloadPassthrough bool `protobuf:"varint,2,opt,name=payload_passthrough,json=payloadPassthrough,proto3" json:"payload_passthrough,omitempty"`
	// Determines the way to invoke the Lambda function.
	InvocationMode Config_InvocationMode `protobuf:"varint,3,opt,name=invocation_mode,json=invocationMode,proto3,enum=envoy.extensions.filters.http.aws_lambda.v3.Config_InvocationMode" json:"invocation_mode,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetArn() string {
	if x != nil {
		return x.Arn
	}
	return ""
}

func (x *Config) GetPayloadPassthrough() bool {
	if x != nil {
		return x.PayloadPassthrough
	}
	return false
}

func (x *Config) GetInvocationMode() Config_InvocationMode {
	if x != nil {
		return x.InvocationMode
	}
	return Config_SYNCHRONOUS
}

// Per-route configuration for AWS Lambda. This can be useful when invoking a different Lambda function or a different
// version of the same Lambda depending on the route.
type PerRouteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvokeConfig *Config `protobuf:"bytes,1,opt,name=invoke_config,json=invokeConfig,proto3" json:"invoke_config,omitempty"`
}

func (x *PerRouteConfig) Reset() {
	*x = PerRouteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerRouteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerRouteConfig) ProtoMessage() {}

func (x *PerRouteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerRouteConfig.ProtoReflect.Descriptor instead.
func (*PerRouteConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_rawDescGZIP(), []int{1}
}

func (x *PerRouteConfig) GetInvokeConfig() *Config {
	if x != nil {
		return x.InvokeConfig
	}
	return nil
}

var File_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x61, 0x77, 0x73, 0x5f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x77,
	0x73, 0x5f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x77,
	0x73, 0x5f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x19, 0x0a, 0x03, 0x61, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x61, 0x72, 0x6e, 0x12, 0x2f, 0x0a, 0x13,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x74, 0x68, 0x72, 0x6f,
	0x75, 0x67, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x50, 0x61, 0x73, 0x73, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x12, 0x75, 0x0a,
	0x0f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x77, 0x73, 0x5f, 0x6c, 0x61, 0x6d, 0x62, 0x64,
	0x61, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x49, 0x6e, 0x76, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x6f, 0x64, 0x65, 0x22, 0x33, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x59, 0x4e, 0x43, 0x48, 0x52,
	0x4f, 0x4e, 0x4f, 0x55, 0x53, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x53, 0x59, 0x4e, 0x43,
	0x48, 0x52, 0x4f, 0x4e, 0x4f, 0x55, 0x53, 0x10, 0x01, 0x3a, 0x39, 0x9a, 0xc5, 0x88, 0x1e, 0x34,
	0x0a, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x77, 0x73, 0x5f, 0x6c,
	0x61, 0x6d, 0x62, 0x64, 0x61, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0xad, 0x01, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x58, 0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61,
	0x77, 0x73, 0x5f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x3a, 0x41, 0x9a, 0xc5, 0x88, 0x1e, 0x3c, 0x0a, 0x3a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x61, 0x77, 0x73, 0x5f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0xb6, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a,
	0x39, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x77, 0x73,
	0x5f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x2e, 0x76, 0x33, 0x42, 0x0e, 0x41, 0x77, 0x73, 0x4c,
	0x61, 0x6d, 0x62, 0x64, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x2f, 0x61, 0x77, 0x73, 0x5f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x2f, 0x76, 0x33,
	0x3b, 0x61, 0x77, 0x73, 0x5f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x76, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_rawDescData = file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_rawDesc
)

func file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_rawDescData
}

var file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_goTypes = []interface{}{
	(Config_InvocationMode)(0), // 0: envoy.extensions.filters.http.aws_lambda.v3.Config.InvocationMode
	(*Config)(nil),             // 1: envoy.extensions.filters.http.aws_lambda.v3.Config
	(*PerRouteConfig)(nil),     // 2: envoy.extensions.filters.http.aws_lambda.v3.PerRouteConfig
}
var file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_depIdxs = []int32{
	0, // 0: envoy.extensions.filters.http.aws_lambda.v3.Config.invocation_mode:type_name -> envoy.extensions.filters.http.aws_lambda.v3.Config.InvocationMode
	1, // 1: envoy.extensions.filters.http.aws_lambda.v3.PerRouteConfig.invoke_config:type_name -> envoy.extensions.filters.http.aws_lambda.v3.Config
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_init() }
func file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_init() {
	if File_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerRouteConfig); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto = out.File
	file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_rawDesc = nil
	file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_goTypes = nil
	file_envoy_extensions_filters_http_aws_lambda_v3_aws_lambda_proto_depIdxs = nil
}
