// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/extensions/filters/network/dubbo_proxy/v3/route.proto

package dubbo_proxyv3

import (
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

func (m *RouteConfiguration) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RouteConfiguration) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RouteConfiguration) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.Routes) > 0 {
		for iNdEx := len(m.Routes) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Routes[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarint(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Group) > 0 {
		i -= len(m.Group)
		copy(dAtA[i:], m.Group)
		i = encodeVarint(dAtA, i, uint64(len(m.Group)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Interface) > 0 {
		i -= len(m.Interface)
		copy(dAtA[i:], m.Interface)
		i = encodeVarint(dAtA, i, uint64(len(m.Interface)))
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

func (m *Route) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *Route) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Route) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.Route != nil {
		size, err := m.Route.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.Match != nil {
		size, err := m.Match.MarshalToSizedBufferVTStrict(dAtA[:i])
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

func (m *RouteMatch) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RouteMatch) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RouteMatch) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.Headers) > 0 {
		for iNdEx := len(m.Headers) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.Headers[iNdEx]).(interface {
				MarshalToSizedBufferVTStrict([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.Headers[iNdEx])
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
	}
	if m.Method != nil {
		size, err := m.Method.MarshalToSizedBufferVTStrict(dAtA[:i])
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

func (m *RouteAction) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RouteAction) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RouteAction) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.MetadataMatch != nil {
		if vtmsg, ok := interface{}(m.MetadataMatch).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.MetadataMatch)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x1a
	}
	if msg, ok := m.ClusterSpecifier.(*RouteAction_WeightedClusters); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ClusterSpecifier.(*RouteAction_Cluster); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *RouteAction_Cluster) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RouteAction_Cluster) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Cluster)
	copy(dAtA[i:], m.Cluster)
	i = encodeVarint(dAtA, i, uint64(len(m.Cluster)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *RouteAction_WeightedClusters) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RouteAction_WeightedClusters) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.WeightedClusters != nil {
		if vtmsg, ok := interface{}(m.WeightedClusters).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.WeightedClusters)
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
func (m *MethodMatch_ParameterMatchSpecifier) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *MethodMatch_ParameterMatchSpecifier) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *MethodMatch_ParameterMatchSpecifier) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.ParameterMatchSpecifier.(*MethodMatch_ParameterMatchSpecifier_RangeMatch); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ParameterMatchSpecifier.(*MethodMatch_ParameterMatchSpecifier_ExactMatch); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *MethodMatch_ParameterMatchSpecifier_ExactMatch) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *MethodMatch_ParameterMatchSpecifier_ExactMatch) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ExactMatch)
	copy(dAtA[i:], m.ExactMatch)
	i = encodeVarint(dAtA, i, uint64(len(m.ExactMatch)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *MethodMatch_ParameterMatchSpecifier_RangeMatch) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *MethodMatch_ParameterMatchSpecifier_RangeMatch) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RangeMatch != nil {
		if vtmsg, ok := interface{}(m.RangeMatch).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.RangeMatch)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *MethodMatch) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *MethodMatch) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *MethodMatch) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.ParamsMatch) > 0 {
		for k := range m.ParamsMatch {
			v := m.ParamsMatch[k]
			baseI := i
			size, err := v.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
			i = encodeVarint(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarint(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Name != nil {
		if vtmsg, ok := interface{}(m.Name).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Name)
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

func (m *MultipleRouteConfiguration) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *MultipleRouteConfiguration) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *MultipleRouteConfiguration) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.RouteConfig) > 0 {
		for iNdEx := len(m.RouteConfig) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.RouteConfig[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x22
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
func (m *RouteConfiguration) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Interface)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Group)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if len(m.Routes) > 0 {
		for _, e := range m.Routes {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *Route) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Match != nil {
		l = m.Match.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.Route != nil {
		l = m.Route.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RouteMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Method != nil {
		l = m.Method.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
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

func (m *RouteAction) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.ClusterSpecifier.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if m.MetadataMatch != nil {
		if size, ok := interface{}(m.MetadataMatch).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.MetadataMatch)
		}
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RouteAction_Cluster) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cluster)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *RouteAction_WeightedClusters) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.WeightedClusters != nil {
		if size, ok := interface{}(m.WeightedClusters).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.WeightedClusters)
		}
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *MethodMatch_ParameterMatchSpecifier) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.ParameterMatchSpecifier.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *MethodMatch_ParameterMatchSpecifier_ExactMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ExactMatch)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *MethodMatch_ParameterMatchSpecifier_RangeMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RangeMatch != nil {
		if size, ok := interface{}(m.RangeMatch).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.RangeMatch)
		}
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *MethodMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Name != nil {
		if size, ok := interface{}(m.Name).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Name)
		}
		n += 1 + l + sov(uint64(l))
	}
	if len(m.ParamsMatch) > 0 {
		for k, v := range m.ParamsMatch {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.SizeVT()
			}
			l += 1 + sov(uint64(l))
			mapEntrySize := 1 + sov(uint64(k)) + l
			n += mapEntrySize + 1 + sov(uint64(mapEntrySize))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *MultipleRouteConfiguration) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if len(m.RouteConfig) > 0 {
		for _, e := range m.RouteConfig {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
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
