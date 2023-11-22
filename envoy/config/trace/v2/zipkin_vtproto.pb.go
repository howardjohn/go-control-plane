// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: envoy/config/trace/v2/zipkin.proto

package tracev2

import (
	wrapperspb "github.com/planetscale/vtprotobuf/types/known/wrapperspb"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *ZipkinConfig) MarshalVT() (dAtA []byte, err error) {
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

func (m *ZipkinConfig) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ZipkinConfig) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.CollectorEndpointVersion != 0 {
		i = encodeVarint(dAtA, i, uint64(m.CollectorEndpointVersion))
		i--
		dAtA[i] = 0x28
	}
	if m.SharedSpanContext != nil {
		size, err := (*wrapperspb.BoolValue)(m.SharedSpanContext).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.TraceId_128Bit {
		i--
		if m.TraceId_128Bit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.CollectorEndpoint) > 0 {
		i -= len(m.CollectorEndpoint)
		copy(dAtA[i:], m.CollectorEndpoint)
		i = encodeVarint(dAtA, i, uint64(len(m.CollectorEndpoint)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CollectorCluster) > 0 {
		i -= len(m.CollectorCluster)
		copy(dAtA[i:], m.CollectorCluster)
		i = encodeVarint(dAtA, i, uint64(len(m.CollectorCluster)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ZipkinConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollectorCluster)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.CollectorEndpoint)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.TraceId_128Bit {
		n += 2
	}
	if m.SharedSpanContext != nil {
		l = (*wrapperspb.BoolValue)(m.SharedSpanContext).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.CollectorEndpointVersion != 0 {
		n += 1 + sov(uint64(m.CollectorEndpointVersion))
	}
	n += len(m.unknownFields)
	return n
}
