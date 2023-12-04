// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: contrib/envoy/extensions/router/cluster_specifier/golang/v3alpha/golang.proto

package v3alpha

import (
	anypb "github.com/planetscale/vtprotobuf/types/known/anypb"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	bits "math/bits"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *Config) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Config) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.Config != nil {
		size, err := (*anypb.Any)(m.Config).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DefaultCluster) > 0 {
		i -= len(m.DefaultCluster)
		copy(dAtA[i:], m.DefaultCluster)
		i = encodeVarint(dAtA, i, uint64(len(m.DefaultCluster)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.LibraryPath) > 0 {
		i -= len(m.LibraryPath)
		copy(dAtA[i:], m.LibraryPath)
		i = encodeVarint(dAtA, i, uint64(len(m.LibraryPath)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.LibraryId) > 0 {
		i -= len(m.LibraryId)
		copy(dAtA[i:], m.LibraryId)
		i = encodeVarint(dAtA, i, uint64(len(m.LibraryId)))
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
func (m *Config) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LibraryId)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.LibraryPath)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.DefaultCluster)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.Config != nil {
		l = (*anypb.Any)(m.Config).SizeVT()
		n += 1 + l + sov(uint64(l))
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
