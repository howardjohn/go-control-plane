// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: envoy/extensions/filters/http/grpc_stats/v3/config.proto

package grpc_statsv3

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

func (m *FilterConfig) MarshalVT() (dAtA []byte, err error) {
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

func (m *FilterConfig) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *FilterConfig) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if vtmsg, ok := m.PerMethodStatSpecifier.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if m.ReplaceDotsInGrpcServiceName {
		i--
		if m.ReplaceDotsInGrpcServiceName {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.EnableUpstreamStats {
		i--
		if m.EnableUpstreamStats {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.EmitFilterState {
		i--
		if m.EmitFilterState {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FilterConfig_IndividualMethodStatsAllowlist) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *FilterConfig_IndividualMethodStatsAllowlist) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IndividualMethodStatsAllowlist != nil {
		if vtmsg, ok := interface{}(m.IndividualMethodStatsAllowlist).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.IndividualMethodStatsAllowlist)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *FilterConfig_StatsForAllMethods) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *FilterConfig_StatsForAllMethods) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StatsForAllMethods != nil {
		size, err := (*wrapperspb.BoolValue)(m.StatsForAllMethods).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *FilterObject) MarshalVT() (dAtA []byte, err error) {
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

func (m *FilterObject) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *FilterObject) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.ResponseMessageCount != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ResponseMessageCount))
		i--
		dAtA[i] = 0x10
	}
	if m.RequestMessageCount != 0 {
		i = encodeVarint(dAtA, i, uint64(m.RequestMessageCount))
		i--
		dAtA[i] = 0x8
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
func (m *FilterConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EmitFilterState {
		n += 2
	}
	if vtmsg, ok := m.PerMethodStatSpecifier.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if m.EnableUpstreamStats {
		n += 2
	}
	if m.ReplaceDotsInGrpcServiceName {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *FilterConfig_IndividualMethodStatsAllowlist) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IndividualMethodStatsAllowlist != nil {
		if size, ok := interface{}(m.IndividualMethodStatsAllowlist).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.IndividualMethodStatsAllowlist)
		}
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *FilterConfig_StatsForAllMethods) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StatsForAllMethods != nil {
		l = (*wrapperspb.BoolValue)(m.StatsForAllMethods).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *FilterObject) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestMessageCount != 0 {
		n += 1 + sov(uint64(m.RequestMessageCount))
	}
	if m.ResponseMessageCount != 0 {
		n += 1 + sov(uint64(m.ResponseMessageCount))
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
