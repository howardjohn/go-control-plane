// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: envoy/config/metrics/v3/metrics_service.proto

package metricsv3

import (
	wrapperspb "github.com/planetscale/vtprotobuf/types/known/wrapperspb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	bits "math/bits"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *MetricsServiceConfig) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricsServiceConfig) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MetricsServiceConfig) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.HistogramEmitMode != 0 {
		i = encodeVarint(dAtA, i, uint64(m.HistogramEmitMode))
		i--
		dAtA[i] = 0x28
	}
	if m.EmitTagsAsLabels {
		i--
		if m.EmitTagsAsLabels {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.TransportApiVersion != 0 {
		i = encodeVarint(dAtA, i, uint64(m.TransportApiVersion))
		i--
		dAtA[i] = 0x18
	}
	if m.ReportCountersAsDeltas != nil {
		size, err := (*wrapperspb.BoolValue)(m.ReportCountersAsDeltas).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.GrpcService != nil {
		if vtmsg, ok := interface{}(m.GrpcService).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.GrpcService)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarint(dAtA []byte, offset int, v uint64) int {
	offset -= sov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MetricsServiceConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GrpcService != nil {
		if size, ok := interface{}(m.GrpcService).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.GrpcService)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.ReportCountersAsDeltas != nil {
		l = (*wrapperspb.BoolValue)(m.ReportCountersAsDeltas).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.TransportApiVersion != 0 {
		n += 1 + sov(uint64(m.TransportApiVersion))
	}
	if m.EmitTagsAsLabels {
		n += 2
	}
	if m.HistogramEmitMode != 0 {
		n += 1 + sov(uint64(m.HistogramEmitMode))
	}
	n += len(m.unknownFields)
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
