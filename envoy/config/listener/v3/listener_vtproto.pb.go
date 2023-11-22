// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: envoy/config/listener/v3/listener.proto

package listenerv3

import (
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

func (m *AdditionalAddress) MarshalVT() (dAtA []byte, err error) {
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

func (m *AdditionalAddress) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *AdditionalAddress) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.SocketOptions != nil {
		if vtmsg, ok := interface{}(m.SocketOptions).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.SocketOptions)
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
	if m.Address != nil {
		if vtmsg, ok := interface{}(m.Address).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Address)
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

func (m *ListenerCollection) MarshalVT() (dAtA []byte, err error) {
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

func (m *ListenerCollection) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ListenerCollection) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.Entries[iNdEx]).(interface {
				MarshalToSizedBufferVT([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.Entries[iNdEx])
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
	}
	return len(dAtA) - i, nil
}

func (m *Listener_DeprecatedV1) MarshalVT() (dAtA []byte, err error) {
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

func (m *Listener_DeprecatedV1) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Listener_DeprecatedV1) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.BindToPort != nil {
		size, err := (*wrapperspb.BoolValue)(m.BindToPort).MarshalToSizedBufferVT(dAtA[:i])
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

func (m *Listener_ConnectionBalanceConfig_ExactBalance) MarshalVT() (dAtA []byte, err error) {
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

func (m *Listener_ConnectionBalanceConfig_ExactBalance) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Listener_ConnectionBalanceConfig_ExactBalance) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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

func (m *Listener_ConnectionBalanceConfig) MarshalVT() (dAtA []byte, err error) {
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

func (m *Listener_ConnectionBalanceConfig) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Listener_ConnectionBalanceConfig) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if vtmsg, ok := m.BalanceType.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *Listener_ConnectionBalanceConfig_ExactBalance_) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Listener_ConnectionBalanceConfig_ExactBalance_) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ExactBalance != nil {
		size, err := m.ExactBalance.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *Listener_ConnectionBalanceConfig_ExtendBalance) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Listener_ConnectionBalanceConfig_ExtendBalance) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ExtendBalance != nil {
		if vtmsg, ok := interface{}(m.ExtendBalance).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.ExtendBalance)
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
func (m *Listener_InternalListenerConfig) MarshalVT() (dAtA []byte, err error) {
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

func (m *Listener_InternalListenerConfig) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Listener_InternalListenerConfig) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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

func (m *Listener) MarshalVT() (dAtA []byte, err error) {
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

func (m *Listener) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Listener) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if vtmsg, ok := m.ListenerSpecifier.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if m.MaxConnectionsToAcceptPerSocketEvent != nil {
		size, err := (*wrapperspb.UInt32Value)(m.MaxConnectionsToAcceptPerSocketEvent).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x2
		i--
		dAtA[i] = 0x92
	}
	if len(m.AdditionalAddresses) > 0 {
		for iNdEx := len(m.AdditionalAddresses) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.AdditionalAddresses[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x2
			i--
			dAtA[i] = 0x8a
		}
	}
	if m.FilterChainMatcher != nil {
		if vtmsg, ok := interface{}(m.FilterChainMatcher).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.FilterChainMatcher)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x2
		i--
		dAtA[i] = 0x82
	}
	if m.IgnoreGlobalConnLimit {
		i--
		if m.IgnoreGlobalConnLimit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xf8
	}
	if m.EnableMptcp {
		i--
		if m.EnableMptcp {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xf0
	}
	if m.EnableReusePort != nil {
		size, err := (*wrapperspb.BoolValue)(m.EnableReusePort).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xea
	}
	if len(m.StatPrefix) > 0 {
		i -= len(m.StatPrefix)
		copy(dAtA[i:], m.StatPrefix)
		i = encodeVarint(dAtA, i, uint64(len(m.StatPrefix)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xe2
	}
	if m.BindToPort != nil {
		size, err := (*wrapperspb.BoolValue)(m.BindToPort).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xd2
	}
	if m.DefaultFilterChain != nil {
		size, err := m.DefaultFilterChain.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xca
	}
	if m.TcpBacklogSize != nil {
		size, err := (*wrapperspb.UInt32Value)(m.TcpBacklogSize).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xc2
	}
	if len(m.AccessLog) > 0 {
		for iNdEx := len(m.AccessLog) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.AccessLog[iNdEx]).(interface {
				MarshalToSizedBufferVT([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.AccessLog[iNdEx])
				if err != nil {
					return 0, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = encodeVarint(dAtA, i, uint64(len(encoded)))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0xb2
		}
	}
	if m.ReusePort {
		i--
		if m.ReusePort {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa8
	}
	if m.ConnectionBalanceConfig != nil {
		size, err := m.ConnectionBalanceConfig.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	if m.ApiListener != nil {
		size, err := m.ApiListener.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x9a
	}
	if m.UdpListenerConfig != nil {
		size, err := m.UdpListenerConfig.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if m.ContinueOnListenerFiltersTimeout {
		i--
		if m.ContinueOnListenerFiltersTimeout {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.TrafficDirection != 0 {
		i = encodeVarint(dAtA, i, uint64(m.TrafficDirection))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.ListenerFiltersTimeout != nil {
		size, err := (*durationpb.Duration)(m.ListenerFiltersTimeout).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.SocketOptions) > 0 {
		for iNdEx := len(m.SocketOptions) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.SocketOptions[iNdEx]).(interface {
				MarshalToSizedBufferVT([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.SocketOptions[iNdEx])
				if err != nil {
					return 0, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = encodeVarint(dAtA, i, uint64(len(encoded)))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if m.TcpFastOpenQueueLength != nil {
		size, err := (*wrapperspb.UInt32Value)(m.TcpFastOpenQueueLength).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x62
	}
	if m.Freebind != nil {
		size, err := (*wrapperspb.BoolValue)(m.Freebind).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x5a
	}
	if m.Transparent != nil {
		size, err := (*wrapperspb.BoolValue)(m.Transparent).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x52
	}
	if len(m.ListenerFilters) > 0 {
		for iNdEx := len(m.ListenerFilters) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.ListenerFilters[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.DrainType != 0 {
		i = encodeVarint(dAtA, i, uint64(m.DrainType))
		i--
		dAtA[i] = 0x40
	}
	if m.DeprecatedV1 != nil {
		size, err := m.DeprecatedV1.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x3a
	}
	if m.Metadata != nil {
		if vtmsg, ok := interface{}(m.Metadata).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Metadata)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.PerConnectionBufferLimitBytes != nil {
		size, err := (*wrapperspb.UInt32Value)(m.PerConnectionBufferLimitBytes).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x2a
	}
	if m.UseOriginalDst != nil {
		size, err := (*wrapperspb.BoolValue)(m.UseOriginalDst).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if len(m.FilterChains) > 0 {
		for iNdEx := len(m.FilterChains) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.FilterChains[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Address != nil {
		if vtmsg, ok := interface{}(m.Address).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Address)
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
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarint(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Listener_InternalListener) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Listener_InternalListener) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.InternalListener != nil {
		size, err := m.InternalListener.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xda
	}
	return len(dAtA) - i, nil
}
func (m *ListenerManager) MarshalVT() (dAtA []byte, err error) {
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

func (m *ListenerManager) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ListenerManager) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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

func (m *ValidationListenerManager) MarshalVT() (dAtA []byte, err error) {
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

func (m *ValidationListenerManager) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ValidationListenerManager) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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

func (m *ApiListenerManager) MarshalVT() (dAtA []byte, err error) {
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

func (m *ApiListenerManager) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ApiListenerManager) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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

func (m *AdditionalAddress) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Address != nil {
		if size, ok := interface{}(m.Address).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Address)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.SocketOptions != nil {
		if size, ok := interface{}(m.SocketOptions).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.SocketOptions)
		}
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ListenerCollection) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
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

func (m *Listener_DeprecatedV1) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BindToPort != nil {
		l = (*wrapperspb.BoolValue)(m.BindToPort).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *Listener_ConnectionBalanceConfig_ExactBalance) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *Listener_ConnectionBalanceConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.BalanceType.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *Listener_ConnectionBalanceConfig_ExactBalance_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExactBalance != nil {
		l = m.ExactBalance.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *Listener_ConnectionBalanceConfig_ExtendBalance) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExtendBalance != nil {
		if size, ok := interface{}(m.ExtendBalance).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.ExtendBalance)
		}
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *Listener_InternalListenerConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *Listener) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.Address != nil {
		if size, ok := interface{}(m.Address).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Address)
		}
		n += 1 + l + sov(uint64(l))
	}
	if len(m.FilterChains) > 0 {
		for _, e := range m.FilterChains {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.UseOriginalDst != nil {
		l = (*wrapperspb.BoolValue)(m.UseOriginalDst).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.PerConnectionBufferLimitBytes != nil {
		l = (*wrapperspb.UInt32Value)(m.PerConnectionBufferLimitBytes).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.Metadata != nil {
		if size, ok := interface{}(m.Metadata).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Metadata)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.DeprecatedV1 != nil {
		l = m.DeprecatedV1.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.DrainType != 0 {
		n += 1 + sov(uint64(m.DrainType))
	}
	if len(m.ListenerFilters) > 0 {
		for _, e := range m.ListenerFilters {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.Transparent != nil {
		l = (*wrapperspb.BoolValue)(m.Transparent).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.Freebind != nil {
		l = (*wrapperspb.BoolValue)(m.Freebind).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.TcpFastOpenQueueLength != nil {
		l = (*wrapperspb.UInt32Value)(m.TcpFastOpenQueueLength).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if len(m.SocketOptions) > 0 {
		for _, e := range m.SocketOptions {
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
	if m.ListenerFiltersTimeout != nil {
		l = (*durationpb.Duration)(m.ListenerFiltersTimeout).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.TrafficDirection != 0 {
		n += 2 + sov(uint64(m.TrafficDirection))
	}
	if m.ContinueOnListenerFiltersTimeout {
		n += 3
	}
	if m.UdpListenerConfig != nil {
		l = m.UdpListenerConfig.SizeVT()
		n += 2 + l + sov(uint64(l))
	}
	if m.ApiListener != nil {
		l = m.ApiListener.SizeVT()
		n += 2 + l + sov(uint64(l))
	}
	if m.ConnectionBalanceConfig != nil {
		l = m.ConnectionBalanceConfig.SizeVT()
		n += 2 + l + sov(uint64(l))
	}
	if m.ReusePort {
		n += 3
	}
	if len(m.AccessLog) > 0 {
		for _, e := range m.AccessLog {
			if size, ok := interface{}(e).(interface {
				SizeVT() int
			}); ok {
				l = size.SizeVT()
			} else {
				l = proto.Size(e)
			}
			n += 2 + l + sov(uint64(l))
		}
	}
	if m.TcpBacklogSize != nil {
		l = (*wrapperspb.UInt32Value)(m.TcpBacklogSize).SizeVT()
		n += 2 + l + sov(uint64(l))
	}
	if m.DefaultFilterChain != nil {
		l = m.DefaultFilterChain.SizeVT()
		n += 2 + l + sov(uint64(l))
	}
	if m.BindToPort != nil {
		l = (*wrapperspb.BoolValue)(m.BindToPort).SizeVT()
		n += 2 + l + sov(uint64(l))
	}
	if vtmsg, ok := m.ListenerSpecifier.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	l = len(m.StatPrefix)
	if l > 0 {
		n += 2 + l + sov(uint64(l))
	}
	if m.EnableReusePort != nil {
		l = (*wrapperspb.BoolValue)(m.EnableReusePort).SizeVT()
		n += 2 + l + sov(uint64(l))
	}
	if m.EnableMptcp {
		n += 3
	}
	if m.IgnoreGlobalConnLimit {
		n += 3
	}
	if m.FilterChainMatcher != nil {
		if size, ok := interface{}(m.FilterChainMatcher).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.FilterChainMatcher)
		}
		n += 2 + l + sov(uint64(l))
	}
	if len(m.AdditionalAddresses) > 0 {
		for _, e := range m.AdditionalAddresses {
			l = e.SizeVT()
			n += 2 + l + sov(uint64(l))
		}
	}
	if m.MaxConnectionsToAcceptPerSocketEvent != nil {
		l = (*wrapperspb.UInt32Value)(m.MaxConnectionsToAcceptPerSocketEvent).SizeVT()
		n += 2 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *Listener_InternalListener) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InternalListener != nil {
		l = m.InternalListener.SizeVT()
		n += 2 + l + sov(uint64(l))
	}
	return n
}
func (m *ListenerManager) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *ValidationListenerManager) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *ApiListenerManager) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}
