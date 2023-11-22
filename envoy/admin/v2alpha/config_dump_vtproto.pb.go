// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: envoy/admin/v2alpha/config_dump.proto

package v2alpha

import (
	anypb "github.com/planetscale/vtprotobuf/types/known/anypb"
	timestamppb "github.com/planetscale/vtprotobuf/types/known/timestamppb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *ConfigDump) MarshalVT() (dAtA []byte, err error) {
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

func (m *ConfigDump) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConfigDump) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.Configs) > 0 {
		for iNdEx := len(m.Configs) - 1; iNdEx >= 0; iNdEx-- {
			size, err := (*anypb.Any)(m.Configs[iNdEx]).MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UpdateFailureState) MarshalVT() (dAtA []byte, err error) {
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

func (m *UpdateFailureState) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *UpdateFailureState) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = encodeVarint(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x1a
	}
	if m.LastUpdateAttempt != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdateAttempt).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.FailedConfiguration != nil {
		size, err := (*anypb.Any)(m.FailedConfiguration).MarshalToSizedBufferVT(dAtA[:i])
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

func (m *BootstrapConfigDump) MarshalVT() (dAtA []byte, err error) {
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

func (m *BootstrapConfigDump) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *BootstrapConfigDump) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.LastUpdated != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdated).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.Bootstrap != nil {
		if vtmsg, ok := interface{}(m.Bootstrap).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Bootstrap)
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

func (m *ListenersConfigDump_StaticListener) MarshalVT() (dAtA []byte, err error) {
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

func (m *ListenersConfigDump_StaticListener) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ListenersConfigDump_StaticListener) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.LastUpdated != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdated).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.Listener != nil {
		size, err := (*anypb.Any)(m.Listener).MarshalToSizedBufferVT(dAtA[:i])
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

func (m *ListenersConfigDump_DynamicListenerState) MarshalVT() (dAtA []byte, err error) {
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

func (m *ListenersConfigDump_DynamicListenerState) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ListenersConfigDump_DynamicListenerState) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.LastUpdated != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdated).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.Listener != nil {
		size, err := (*anypb.Any)(m.Listener).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if len(m.VersionInfo) > 0 {
		i -= len(m.VersionInfo)
		copy(dAtA[i:], m.VersionInfo)
		i = encodeVarint(dAtA, i, uint64(len(m.VersionInfo)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListenersConfigDump_DynamicListener) MarshalVT() (dAtA []byte, err error) {
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

func (m *ListenersConfigDump_DynamicListener) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ListenersConfigDump_DynamicListener) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.ErrorState != nil {
		size, err := m.ErrorState.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x2a
	}
	if m.DrainingState != nil {
		size, err := m.DrainingState.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.WarmingState != nil {
		size, err := m.WarmingState.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.ActiveState != nil {
		size, err := m.ActiveState.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
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

func (m *ListenersConfigDump) MarshalVT() (dAtA []byte, err error) {
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

func (m *ListenersConfigDump) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ListenersConfigDump) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.DynamicListeners) > 0 {
		for iNdEx := len(m.DynamicListeners) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.DynamicListeners[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.StaticListeners) > 0 {
		for iNdEx := len(m.StaticListeners) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.StaticListeners[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.VersionInfo) > 0 {
		i -= len(m.VersionInfo)
		copy(dAtA[i:], m.VersionInfo)
		i = encodeVarint(dAtA, i, uint64(len(m.VersionInfo)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClustersConfigDump_StaticCluster) MarshalVT() (dAtA []byte, err error) {
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

func (m *ClustersConfigDump_StaticCluster) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ClustersConfigDump_StaticCluster) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.LastUpdated != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdated).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.Cluster != nil {
		size, err := (*anypb.Any)(m.Cluster).MarshalToSizedBufferVT(dAtA[:i])
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

func (m *ClustersConfigDump_DynamicCluster) MarshalVT() (dAtA []byte, err error) {
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

func (m *ClustersConfigDump_DynamicCluster) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ClustersConfigDump_DynamicCluster) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.LastUpdated != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdated).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.Cluster != nil {
		size, err := (*anypb.Any)(m.Cluster).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if len(m.VersionInfo) > 0 {
		i -= len(m.VersionInfo)
		copy(dAtA[i:], m.VersionInfo)
		i = encodeVarint(dAtA, i, uint64(len(m.VersionInfo)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClustersConfigDump) MarshalVT() (dAtA []byte, err error) {
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

func (m *ClustersConfigDump) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ClustersConfigDump) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.DynamicWarmingClusters) > 0 {
		for iNdEx := len(m.DynamicWarmingClusters) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.DynamicWarmingClusters[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.DynamicActiveClusters) > 0 {
		for iNdEx := len(m.DynamicActiveClusters) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.DynamicActiveClusters[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.StaticClusters) > 0 {
		for iNdEx := len(m.StaticClusters) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.StaticClusters[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.VersionInfo) > 0 {
		i -= len(m.VersionInfo)
		copy(dAtA[i:], m.VersionInfo)
		i = encodeVarint(dAtA, i, uint64(len(m.VersionInfo)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RoutesConfigDump_StaticRouteConfig) MarshalVT() (dAtA []byte, err error) {
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

func (m *RoutesConfigDump_StaticRouteConfig) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *RoutesConfigDump_StaticRouteConfig) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.LastUpdated != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdated).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.RouteConfig != nil {
		size, err := (*anypb.Any)(m.RouteConfig).MarshalToSizedBufferVT(dAtA[:i])
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

func (m *RoutesConfigDump_DynamicRouteConfig) MarshalVT() (dAtA []byte, err error) {
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

func (m *RoutesConfigDump_DynamicRouteConfig) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *RoutesConfigDump_DynamicRouteConfig) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.LastUpdated != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdated).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.RouteConfig != nil {
		size, err := (*anypb.Any)(m.RouteConfig).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if len(m.VersionInfo) > 0 {
		i -= len(m.VersionInfo)
		copy(dAtA[i:], m.VersionInfo)
		i = encodeVarint(dAtA, i, uint64(len(m.VersionInfo)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RoutesConfigDump) MarshalVT() (dAtA []byte, err error) {
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

func (m *RoutesConfigDump) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *RoutesConfigDump) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.DynamicRouteConfigs) > 0 {
		for iNdEx := len(m.DynamicRouteConfigs) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.DynamicRouteConfigs[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.StaticRouteConfigs) > 0 {
		for iNdEx := len(m.StaticRouteConfigs) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.StaticRouteConfigs[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *ScopedRoutesConfigDump_InlineScopedRouteConfigs) MarshalVT() (dAtA []byte, err error) {
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

func (m *ScopedRoutesConfigDump_InlineScopedRouteConfigs) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ScopedRoutesConfigDump_InlineScopedRouteConfigs) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.LastUpdated != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdated).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ScopedRouteConfigs) > 0 {
		for iNdEx := len(m.ScopedRouteConfigs) - 1; iNdEx >= 0; iNdEx-- {
			size, err := (*anypb.Any)(m.ScopedRouteConfigs[iNdEx]).MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
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

func (m *ScopedRoutesConfigDump_DynamicScopedRouteConfigs) MarshalVT() (dAtA []byte, err error) {
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

func (m *ScopedRoutesConfigDump_DynamicScopedRouteConfigs) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ScopedRoutesConfigDump_DynamicScopedRouteConfigs) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.LastUpdated != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdated).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ScopedRouteConfigs) > 0 {
		for iNdEx := len(m.ScopedRouteConfigs) - 1; iNdEx >= 0; iNdEx-- {
			size, err := (*anypb.Any)(m.ScopedRouteConfigs[iNdEx]).MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.VersionInfo) > 0 {
		i -= len(m.VersionInfo)
		copy(dAtA[i:], m.VersionInfo)
		i = encodeVarint(dAtA, i, uint64(len(m.VersionInfo)))
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

func (m *ScopedRoutesConfigDump) MarshalVT() (dAtA []byte, err error) {
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

func (m *ScopedRoutesConfigDump) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ScopedRoutesConfigDump) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.DynamicScopedRouteConfigs) > 0 {
		for iNdEx := len(m.DynamicScopedRouteConfigs) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.DynamicScopedRouteConfigs[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.InlineScopedRouteConfigs) > 0 {
		for iNdEx := len(m.InlineScopedRouteConfigs) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.InlineScopedRouteConfigs[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SecretsConfigDump_DynamicSecret) MarshalVT() (dAtA []byte, err error) {
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

func (m *SecretsConfigDump_DynamicSecret) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *SecretsConfigDump_DynamicSecret) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.Secret != nil {
		size, err := (*anypb.Any)(m.Secret).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.LastUpdated != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdated).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.VersionInfo) > 0 {
		i -= len(m.VersionInfo)
		copy(dAtA[i:], m.VersionInfo)
		i = encodeVarint(dAtA, i, uint64(len(m.VersionInfo)))
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

func (m *SecretsConfigDump_StaticSecret) MarshalVT() (dAtA []byte, err error) {
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

func (m *SecretsConfigDump_StaticSecret) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *SecretsConfigDump_StaticSecret) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.Secret != nil {
		size, err := (*anypb.Any)(m.Secret).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.LastUpdated != nil {
		size, err := (*timestamppb.Timestamp)(m.LastUpdated).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
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

func (m *SecretsConfigDump) MarshalVT() (dAtA []byte, err error) {
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

func (m *SecretsConfigDump) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *SecretsConfigDump) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.DynamicWarmingSecrets) > 0 {
		for iNdEx := len(m.DynamicWarmingSecrets) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.DynamicWarmingSecrets[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DynamicActiveSecrets) > 0 {
		for iNdEx := len(m.DynamicActiveSecrets) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.DynamicActiveSecrets[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.StaticSecrets) > 0 {
		for iNdEx := len(m.StaticSecrets) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.StaticSecrets[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ConfigDump) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Configs) > 0 {
		for _, e := range m.Configs {
			l = (*anypb.Any)(e).SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *UpdateFailureState) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FailedConfiguration != nil {
		l = (*anypb.Any)(m.FailedConfiguration).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.LastUpdateAttempt != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdateAttempt).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *BootstrapConfigDump) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Bootstrap != nil {
		if size, ok := interface{}(m.Bootstrap).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Bootstrap)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.LastUpdated != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdated).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ListenersConfigDump_StaticListener) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Listener != nil {
		l = (*anypb.Any)(m.Listener).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.LastUpdated != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdated).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ListenersConfigDump_DynamicListenerState) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VersionInfo)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.Listener != nil {
		l = (*anypb.Any)(m.Listener).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.LastUpdated != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdated).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ListenersConfigDump_DynamicListener) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.ActiveState != nil {
		l = m.ActiveState.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.WarmingState != nil {
		l = m.WarmingState.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.DrainingState != nil {
		l = m.DrainingState.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.ErrorState != nil {
		l = m.ErrorState.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ListenersConfigDump) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VersionInfo)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if len(m.StaticListeners) > 0 {
		for _, e := range m.StaticListeners {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.DynamicListeners) > 0 {
		for _, e := range m.DynamicListeners {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *ClustersConfigDump_StaticCluster) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cluster != nil {
		l = (*anypb.Any)(m.Cluster).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.LastUpdated != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdated).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ClustersConfigDump_DynamicCluster) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VersionInfo)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.Cluster != nil {
		l = (*anypb.Any)(m.Cluster).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.LastUpdated != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdated).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ClustersConfigDump) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VersionInfo)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if len(m.StaticClusters) > 0 {
		for _, e := range m.StaticClusters {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.DynamicActiveClusters) > 0 {
		for _, e := range m.DynamicActiveClusters {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.DynamicWarmingClusters) > 0 {
		for _, e := range m.DynamicWarmingClusters {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *RoutesConfigDump_StaticRouteConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RouteConfig != nil {
		l = (*anypb.Any)(m.RouteConfig).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.LastUpdated != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdated).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RoutesConfigDump_DynamicRouteConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VersionInfo)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.RouteConfig != nil {
		l = (*anypb.Any)(m.RouteConfig).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.LastUpdated != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdated).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RoutesConfigDump) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.StaticRouteConfigs) > 0 {
		for _, e := range m.StaticRouteConfigs {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.DynamicRouteConfigs) > 0 {
		for _, e := range m.DynamicRouteConfigs {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *ScopedRoutesConfigDump_InlineScopedRouteConfigs) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if len(m.ScopedRouteConfigs) > 0 {
		for _, e := range m.ScopedRouteConfigs {
			l = (*anypb.Any)(e).SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.LastUpdated != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdated).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ScopedRoutesConfigDump_DynamicScopedRouteConfigs) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.VersionInfo)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if len(m.ScopedRouteConfigs) > 0 {
		for _, e := range m.ScopedRouteConfigs {
			l = (*anypb.Any)(e).SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.LastUpdated != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdated).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ScopedRoutesConfigDump) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.InlineScopedRouteConfigs) > 0 {
		for _, e := range m.InlineScopedRouteConfigs {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.DynamicScopedRouteConfigs) > 0 {
		for _, e := range m.DynamicScopedRouteConfigs {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *SecretsConfigDump_DynamicSecret) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.VersionInfo)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.LastUpdated != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdated).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.Secret != nil {
		l = (*anypb.Any)(m.Secret).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *SecretsConfigDump_StaticSecret) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.LastUpdated != nil {
		l = (*timestamppb.Timestamp)(m.LastUpdated).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.Secret != nil {
		l = (*anypb.Any)(m.Secret).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *SecretsConfigDump) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.StaticSecrets) > 0 {
		for _, e := range m.StaticSecrets {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.DynamicActiveSecrets) > 0 {
		for _, e := range m.DynamicActiveSecrets {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.DynamicWarmingSecrets) > 0 {
		for _, e := range m.DynamicWarmingSecrets {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}
