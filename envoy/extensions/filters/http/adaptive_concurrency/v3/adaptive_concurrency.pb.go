// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: envoy/extensions/filters/http/adaptive_concurrency/v3/adaptive_concurrency.proto

package adaptive_concurrencyv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configuration parameters for the gradient controller.
type GradientControllerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The percentile to use when summarizing aggregated samples. Defaults to p50.
	SampleAggregatePercentile *v3.Percent                                                 `protobuf:"bytes,1,opt,name=sample_aggregate_percentile,json=sampleAggregatePercentile,proto3" json:"sample_aggregate_percentile,omitempty"`
	ConcurrencyLimitParams    *GradientControllerConfig_ConcurrencyLimitCalculationParams `protobuf:"bytes,2,opt,name=concurrency_limit_params,json=concurrencyLimitParams,proto3" json:"concurrency_limit_params,omitempty"`
	MinRttCalcParams          *GradientControllerConfig_MinimumRTTCalculationParams       `protobuf:"bytes,3,opt,name=min_rtt_calc_params,json=minRttCalcParams,proto3" json:"min_rtt_calc_params,omitempty"`
}

func (x *GradientControllerConfig) Reset() {
	*x = GradientControllerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GradientControllerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GradientControllerConfig) ProtoMessage() {}

func (x *GradientControllerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GradientControllerConfig.ProtoReflect.Descriptor instead.
func (*GradientControllerConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDescGZIP(), []int{0}
}

func (x *GradientControllerConfig) GetSampleAggregatePercentile() *v3.Percent {
	if x != nil {
		return x.SampleAggregatePercentile
	}
	return nil
}

func (x *GradientControllerConfig) GetConcurrencyLimitParams() *GradientControllerConfig_ConcurrencyLimitCalculationParams {
	if x != nil {
		return x.ConcurrencyLimitParams
	}
	return nil
}

func (x *GradientControllerConfig) GetMinRttCalcParams() *GradientControllerConfig_MinimumRTTCalculationParams {
	if x != nil {
		return x.MinRttCalcParams
	}
	return nil
}

type AdaptiveConcurrency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ConcurrencyControllerConfig:
	//
	//	*AdaptiveConcurrency_GradientControllerConfig
	ConcurrencyControllerConfig isAdaptiveConcurrency_ConcurrencyControllerConfig `protobuf_oneof:"concurrency_controller_config"`
	// If set to false, the adaptive concurrency filter will operate as a pass-through filter. If the
	// message is unspecified, the filter will be enabled.
	Enabled *v31.RuntimeFeatureFlag `protobuf:"bytes,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// This field allows for a custom HTTP response status code to the downstream client when
	// the concurrency limit has been exceeded.
	// Defaults to 503 (Service Unavailable).
	//
	// .. note::
	//
	//	If this is set to < 400, 503 will be used instead.
	ConcurrencyLimitExceededStatus *v3.HttpStatus `protobuf:"bytes,3,opt,name=concurrency_limit_exceeded_status,json=concurrencyLimitExceededStatus,proto3" json:"concurrency_limit_exceeded_status,omitempty"`
}

func (x *AdaptiveConcurrency) Reset() {
	*x = AdaptiveConcurrency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdaptiveConcurrency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdaptiveConcurrency) ProtoMessage() {}

func (x *AdaptiveConcurrency) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdaptiveConcurrency.ProtoReflect.Descriptor instead.
func (*AdaptiveConcurrency) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDescGZIP(), []int{1}
}

func (m *AdaptiveConcurrency) GetConcurrencyControllerConfig() isAdaptiveConcurrency_ConcurrencyControllerConfig {
	if m != nil {
		return m.ConcurrencyControllerConfig
	}
	return nil
}

func (x *AdaptiveConcurrency) GetGradientControllerConfig() *GradientControllerConfig {
	if x, ok := x.GetConcurrencyControllerConfig().(*AdaptiveConcurrency_GradientControllerConfig); ok {
		return x.GradientControllerConfig
	}
	return nil
}

func (x *AdaptiveConcurrency) GetEnabled() *v31.RuntimeFeatureFlag {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *AdaptiveConcurrency) GetConcurrencyLimitExceededStatus() *v3.HttpStatus {
	if x != nil {
		return x.ConcurrencyLimitExceededStatus
	}
	return nil
}

type isAdaptiveConcurrency_ConcurrencyControllerConfig interface {
	isAdaptiveConcurrency_ConcurrencyControllerConfig()
}

type AdaptiveConcurrency_GradientControllerConfig struct {
	// Gradient concurrency control will be used.
	GradientControllerConfig *GradientControllerConfig `protobuf:"bytes,1,opt,name=gradient_controller_config,json=gradientControllerConfig,proto3,oneof"`
}

func (*AdaptiveConcurrency_GradientControllerConfig) isAdaptiveConcurrency_ConcurrencyControllerConfig() {
}

// Parameters controlling the periodic recalculation of the concurrency limit from sampled request
// latencies.
type GradientControllerConfig_ConcurrencyLimitCalculationParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The allowed upper-bound on the calculated concurrency limit. Defaults to 1000.
	MaxConcurrencyLimit *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=max_concurrency_limit,json=maxConcurrencyLimit,proto3" json:"max_concurrency_limit,omitempty"`
	// The period of time samples are taken to recalculate the concurrency limit.
	ConcurrencyUpdateInterval *durationpb.Duration `protobuf:"bytes,3,opt,name=concurrency_update_interval,json=concurrencyUpdateInterval,proto3" json:"concurrency_update_interval,omitempty"`
}

func (x *GradientControllerConfig_ConcurrencyLimitCalculationParams) Reset() {
	*x = GradientControllerConfig_ConcurrencyLimitCalculationParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GradientControllerConfig_ConcurrencyLimitCalculationParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GradientControllerConfig_ConcurrencyLimitCalculationParams) ProtoMessage() {}

func (x *GradientControllerConfig_ConcurrencyLimitCalculationParams) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GradientControllerConfig_ConcurrencyLimitCalculationParams.ProtoReflect.Descriptor instead.
func (*GradientControllerConfig_ConcurrencyLimitCalculationParams) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GradientControllerConfig_ConcurrencyLimitCalculationParams) GetMaxConcurrencyLimit() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxConcurrencyLimit
	}
	return nil
}

func (x *GradientControllerConfig_ConcurrencyLimitCalculationParams) GetConcurrencyUpdateInterval() *durationpb.Duration {
	if x != nil {
		return x.ConcurrencyUpdateInterval
	}
	return nil
}

// Parameters controlling the periodic minRTT recalculation.
// [#next-free-field: 6]
type GradientControllerConfig_MinimumRTTCalculationParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The time interval between recalculating the minimum request round-trip time. Has to be
	// positive.
	Interval *durationpb.Duration `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	// The number of requests to aggregate/sample during the minRTT recalculation window before
	// updating. Defaults to 50.
	RequestCount *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=request_count,json=requestCount,proto3" json:"request_count,omitempty"`
	// Randomized time delta that will be introduced to the start of the minRTT calculation window.
	// This is represented as a percentage of the interval duration. Defaults to 15%.
	//
	// Example: If the interval is 10s and the jitter is 15%, the next window will begin
	// somewhere in the range (10s - 11.5s).
	Jitter *v3.Percent `protobuf:"bytes,3,opt,name=jitter,proto3" json:"jitter,omitempty"`
	// The concurrency limit set while measuring the minRTT. Defaults to 3.
	MinConcurrency *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=min_concurrency,json=minConcurrency,proto3" json:"min_concurrency,omitempty"`
	// Amount added to the measured minRTT to add stability to the concurrency limit during natural
	// variability in latency. This is expressed as a percentage of the measured value and can be
	// adjusted to allow more or less tolerance to the sampled latency values.
	//
	// Defaults to 25%.
	Buffer *v3.Percent `protobuf:"bytes,5,opt,name=buffer,proto3" json:"buffer,omitempty"`
}

func (x *GradientControllerConfig_MinimumRTTCalculationParams) Reset() {
	*x = GradientControllerConfig_MinimumRTTCalculationParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GradientControllerConfig_MinimumRTTCalculationParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GradientControllerConfig_MinimumRTTCalculationParams) ProtoMessage() {}

func (x *GradientControllerConfig_MinimumRTTCalculationParams) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GradientControllerConfig_MinimumRTTCalculationParams.ProtoReflect.Descriptor instead.
func (*GradientControllerConfig_MinimumRTTCalculationParams) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDescGZIP(), []int{0, 1}
}

func (x *GradientControllerConfig_MinimumRTTCalculationParams) GetInterval() *durationpb.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *GradientControllerConfig_MinimumRTTCalculationParams) GetRequestCount() *wrapperspb.UInt32Value {
	if x != nil {
		return x.RequestCount
	}
	return nil
}

func (x *GradientControllerConfig_MinimumRTTCalculationParams) GetJitter() *v3.Percent {
	if x != nil {
		return x.Jitter
	}
	return nil
}

func (x *GradientControllerConfig_MinimumRTTCalculationParams) GetMinConcurrency() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MinConcurrency
	}
	return nil
}

func (x *GradientControllerConfig_MinimumRTTCalculationParams) GetBuffer() *v3.Percent {
	if x != nil {
		return x.Buffer
	}
	return nil
}

var File_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDesc = []byte{
	0x0a, 0x50, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x35, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x0a, 0x0a, 0x18, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x56, 0x0a, 0x1b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x19, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x12, 0xb5, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x71, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x64, 0x61, 0x70,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x76, 0x33, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x16, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0xa4, 0x01, 0x0a, 0x13, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x74, 0x74, 0x5f, 0x63, 0x61, 0x6c,
	0x63, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x6b,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61,
	0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x54, 0x54, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x52, 0x74, 0x74, 0x43, 0x61, 0x6c,
	0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0xde, 0x02, 0x0a, 0x21, 0x43, 0x6f, 0x6e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x59, 0x0a,
	0x15, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a,
	0x02, 0x20, 0x00, 0x52, 0x13, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x65, 0x0a, 0x1b, 0x63, 0x6f, 0x6e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0xaa, 0x01, 0x04,
	0x08, 0x01, 0x2a, 0x00, 0x52, 0x19, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x3a,
	0x77, 0x9a, 0xc5, 0x88, 0x1e, 0x72, 0x0a, 0x70, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x72,
	0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0xd3, 0x03, 0x0a, 0x1b, 0x4d, 0x69, 0x6e,
	0x69, 0x6d, 0x75, 0x6d, 0x52, 0x54, 0x54, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x45, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0xaa, 0x01, 0x08, 0x08, 0x01, 0x32,
	0x04, 0x10, 0xc0, 0x84, 0x3d, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x4a, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x0c, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x6a,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x52, 0x06, 0x6a, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x0f, 0x6d,
	0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x0e, 0x6d, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2e, 0x0a, 0x06, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x52, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x3a, 0x71, 0x9a, 0xc5, 0x88,
	0x1e, 0x6c, 0x0a, 0x6a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x54, 0x54, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x3a, 0x55,
	0x9a, 0xc5, 0x88, 0x1e, 0x50, 0x0a, 0x4e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x72, 0x61,
	0x64, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xd3, 0x03, 0x0a, 0x13, 0x41, 0x64, 0x61, 0x70, 0x74, 0x69,
	0x76, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x99, 0x01,
	0x0a, 0x1a, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52,
	0x18, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x42, 0x0a, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x46, 0x6c, 0x61, 0x67, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x64, 0x0a,
	0x21, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x5f, 0x65, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x1e, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x45, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x3a, 0x50, 0x9a, 0xc5, 0x88, 0x1e, 0x4b, 0x0a, 0x49, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x24, 0x0a, 0x1d, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0xde, 0x01, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x43, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x18, 0x41, 0x64, 0x61,
	0x70, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x73, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67,
	0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x61, 0x64,
	0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2f, 0x76, 0x33, 0x3b, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDescData = file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDesc
)

func file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDescData
}

var file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_goTypes = []interface{}{
	(*GradientControllerConfig)(nil),                                   // 0: envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig
	(*AdaptiveConcurrency)(nil),                                        // 1: envoy.extensions.filters.http.adaptive_concurrency.v3.AdaptiveConcurrency
	(*GradientControllerConfig_ConcurrencyLimitCalculationParams)(nil), // 2: envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig.ConcurrencyLimitCalculationParams
	(*GradientControllerConfig_MinimumRTTCalculationParams)(nil),       // 3: envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig.MinimumRTTCalculationParams
	(*v3.Percent)(nil),                                                 // 4: envoy.type.v3.Percent
	(*v31.RuntimeFeatureFlag)(nil),                                     // 5: envoy.config.core.v3.RuntimeFeatureFlag
	(*v3.HttpStatus)(nil),                                              // 6: envoy.type.v3.HttpStatus
	(*wrapperspb.UInt32Value)(nil),                                     // 7: google.protobuf.UInt32Value
	(*durationpb.Duration)(nil),                                        // 8: google.protobuf.Duration
}
var file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_depIdxs = []int32{
	4,  // 0: envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig.sample_aggregate_percentile:type_name -> envoy.type.v3.Percent
	2,  // 1: envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig.concurrency_limit_params:type_name -> envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig.ConcurrencyLimitCalculationParams
	3,  // 2: envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig.min_rtt_calc_params:type_name -> envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig.MinimumRTTCalculationParams
	0,  // 3: envoy.extensions.filters.http.adaptive_concurrency.v3.AdaptiveConcurrency.gradient_controller_config:type_name -> envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig
	5,  // 4: envoy.extensions.filters.http.adaptive_concurrency.v3.AdaptiveConcurrency.enabled:type_name -> envoy.config.core.v3.RuntimeFeatureFlag
	6,  // 5: envoy.extensions.filters.http.adaptive_concurrency.v3.AdaptiveConcurrency.concurrency_limit_exceeded_status:type_name -> envoy.type.v3.HttpStatus
	7,  // 6: envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig.ConcurrencyLimitCalculationParams.max_concurrency_limit:type_name -> google.protobuf.UInt32Value
	8,  // 7: envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig.ConcurrencyLimitCalculationParams.concurrency_update_interval:type_name -> google.protobuf.Duration
	8,  // 8: envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig.MinimumRTTCalculationParams.interval:type_name -> google.protobuf.Duration
	7,  // 9: envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig.MinimumRTTCalculationParams.request_count:type_name -> google.protobuf.UInt32Value
	4,  // 10: envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig.MinimumRTTCalculationParams.jitter:type_name -> envoy.type.v3.Percent
	7,  // 11: envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig.MinimumRTTCalculationParams.min_concurrency:type_name -> google.protobuf.UInt32Value
	4,  // 12: envoy.extensions.filters.http.adaptive_concurrency.v3.GradientControllerConfig.MinimumRTTCalculationParams.buffer:type_name -> envoy.type.v3.Percent
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_init()
}
func file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_init() {
	if File_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GradientControllerConfig); i {
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
		file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdaptiveConcurrency); i {
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
		file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GradientControllerConfig_ConcurrencyLimitCalculationParams); i {
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
		file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GradientControllerConfig_MinimumRTTCalculationParams); i {
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
	file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AdaptiveConcurrency_GradientControllerConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto = out.File
	file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_rawDesc = nil
	file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_goTypes = nil
	file_envoy_extensions_filters_http_adaptive_concurrency_v3_adaptive_concurrency_proto_depIdxs = nil
}
