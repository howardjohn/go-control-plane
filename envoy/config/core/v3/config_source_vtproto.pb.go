// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: envoy/config/core/v3/config_source.proto

package corev3

import (
	anypb "github.com/planetscale/vtprotobuf/types/known/anypb"
	durationpb "github.com/planetscale/vtprotobuf/types/known/durationpb"
	wrapperspb "github.com/planetscale/vtprotobuf/types/known/wrapperspb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *ApiConfigSource) MarshalVT() (dAtA []byte, err error) {
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

func (m *ApiConfigSource) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ApiConfigSource) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.ConfigValidators) > 0 {
		for iNdEx := len(m.ConfigValidators) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.ConfigValidators[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.TransportApiVersion != 0 {
		i = encodeVarint(dAtA, i, uint64(m.TransportApiVersion))
		i--
		dAtA[i] = 0x40
	}
	if m.SetNodeOnFirstMessageOnly {
		i--
		if m.SetNodeOnFirstMessageOnly {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.RateLimitSettings != nil {
		size, err := m.RateLimitSettings.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x32
	}
	if m.RequestTimeout != nil {
		size, err := (*durationpb.Duration)(m.RequestTimeout).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.GrpcServices) > 0 {
		for iNdEx := len(m.GrpcServices) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.GrpcServices[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.RefreshDelay != nil {
		size, err := (*durationpb.Duration)(m.RefreshDelay).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClusterNames) > 0 {
		for iNdEx := len(m.ClusterNames) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ClusterNames[iNdEx])
			copy(dAtA[i:], m.ClusterNames[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.ClusterNames[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.ApiType != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ApiType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AggregatedConfigSource) MarshalVT() (dAtA []byte, err error) {
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

func (m *AggregatedConfigSource) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *AggregatedConfigSource) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	return len(dAtA) - i, nil
}

func (m *SelfConfigSource) MarshalVT() (dAtA []byte, err error) {
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

func (m *SelfConfigSource) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *SelfConfigSource) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.TransportApiVersion != 0 {
		i = encodeVarint(dAtA, i, uint64(m.TransportApiVersion))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitSettings) MarshalVT() (dAtA []byte, err error) {
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

func (m *RateLimitSettings) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *RateLimitSettings) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.FillRate != nil {
		size, err := (*wrapperspb.DoubleValue)(m.FillRate).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.MaxTokens != nil {
		size, err := (*wrapperspb.UInt32Value)(m.MaxTokens).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PathConfigSource) MarshalVT() (dAtA []byte, err error) {
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

func (m *PathConfigSource) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *PathConfigSource) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.WatchedDirectory != nil {
		size, err := m.WatchedDirectory.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarint(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConfigSource) MarshalVT() (dAtA []byte, err error) {
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

func (m *ConfigSource) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConfigSource) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if vtmsg, ok := m.ConfigSourceSpecifier.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if len(m.Authorities) > 0 {
		for iNdEx := len(m.Authorities) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.Authorities[iNdEx]).(interface {
				MarshalToSizedBufferVT([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.Authorities[iNdEx])
				if err != nil {
					return 0, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = encodeVarint(dAtA, i, uint64(len(encoded)))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.ResourceApiVersion != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ResourceApiVersion))
		i--
		dAtA[i] = 0x30
	}
	if m.InitialFetchTimeout != nil {
		size, err := (*durationpb.Duration)(m.InitialFetchTimeout).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}

func (m *ConfigSource_Path) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConfigSource_Path) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Path)
	copy(dAtA[i:], m.Path)
	i = encodeVarint(dAtA, i, uint64(len(m.Path)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *ConfigSource_ApiConfigSource) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConfigSource_ApiConfigSource) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ApiConfigSource != nil {
		size, err := m.ApiConfigSource.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *ConfigSource_Ads) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConfigSource_Ads) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Ads != nil {
		size, err := m.Ads.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *ConfigSource_Self) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConfigSource_Self) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Self != nil {
		size, err := m.Self.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *ConfigSource_PathConfigSource) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConfigSource_PathConfigSource) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PathConfigSource != nil {
		size, err := m.PathConfigSource.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *ExtensionConfigSource) MarshalVT() (dAtA []byte, err error) {
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

func (m *ExtensionConfigSource) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ExtensionConfigSource) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.TypeUrls) > 0 {
		for iNdEx := len(m.TypeUrls) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TypeUrls[iNdEx])
			copy(dAtA[i:], m.TypeUrls[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.TypeUrls[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.ApplyDefaultConfigWithoutWarming {
		i--
		if m.ApplyDefaultConfigWithoutWarming {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.DefaultConfig != nil {
		size, err := (*anypb.Any)(m.DefaultConfig).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.ConfigSource != nil {
		size, err := m.ConfigSource.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ApiConfigSource) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ApiType != 0 {
		n += 1 + sov(uint64(m.ApiType))
	}
	if len(m.ClusterNames) > 0 {
		for _, s := range m.ClusterNames {
			l = len(s)
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.RefreshDelay != nil {
		l = (*durationpb.Duration)(m.RefreshDelay).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if len(m.GrpcServices) > 0 {
		for _, e := range m.GrpcServices {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.RequestTimeout != nil {
		l = (*durationpb.Duration)(m.RequestTimeout).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.RateLimitSettings != nil {
		l = m.RateLimitSettings.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.SetNodeOnFirstMessageOnly {
		n += 2
	}
	if m.TransportApiVersion != 0 {
		n += 1 + sov(uint64(m.TransportApiVersion))
	}
	if len(m.ConfigValidators) > 0 {
		for _, e := range m.ConfigValidators {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *AggregatedConfigSource) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *SelfConfigSource) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TransportApiVersion != 0 {
		n += 1 + sov(uint64(m.TransportApiVersion))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitSettings) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxTokens != nil {
		l = (*wrapperspb.UInt32Value)(m.MaxTokens).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.FillRate != nil {
		l = (*wrapperspb.DoubleValue)(m.FillRate).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *PathConfigSource) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.WatchedDirectory != nil {
		l = m.WatchedDirectory.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ConfigSource) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.ConfigSourceSpecifier.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if m.InitialFetchTimeout != nil {
		l = (*durationpb.Duration)(m.InitialFetchTimeout).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.ResourceApiVersion != 0 {
		n += 1 + sov(uint64(m.ResourceApiVersion))
	}
	if len(m.Authorities) > 0 {
		for _, e := range m.Authorities {
			if size, ok := interface{}(e).(interface {
				SizeVT() int
			}); ok {
				l = size.SizeVT()
			} else {
				l = proto.Size(e)
			}
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *ConfigSource_Path) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *ConfigSource_ApiConfigSource) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ApiConfigSource != nil {
		l = m.ApiConfigSource.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *ConfigSource_Ads) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ads != nil {
		l = m.Ads.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *ConfigSource_Self) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Self != nil {
		l = m.Self.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *ConfigSource_PathConfigSource) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PathConfigSource != nil {
		l = m.PathConfigSource.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *ExtensionConfigSource) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConfigSource != nil {
		l = m.ConfigSource.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.DefaultConfig != nil {
		l = (*anypb.Any)(m.DefaultConfig).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.ApplyDefaultConfigWithoutWarming {
		n += 2
	}
	if len(m.TypeUrls) > 0 {
		for _, s := range m.TypeUrls {
			l = len(s)
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}
