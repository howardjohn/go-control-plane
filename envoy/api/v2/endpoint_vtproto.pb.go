// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: envoy/api/v2/endpoint.proto

package apiv2

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

func (m *ClusterLoadAssignment_Policy_DropOverload) MarshalVT() (dAtA []byte, err error) {
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

func (m *ClusterLoadAssignment_Policy_DropOverload) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ClusterLoadAssignment_Policy_DropOverload) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.DropPercentage != nil {
		if vtmsg, ok := interface{}(m.DropPercentage).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.DropPercentage)
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
	if len(m.Category) > 0 {
		i -= len(m.Category)
		copy(dAtA[i:], m.Category)
		i = encodeVarint(dAtA, i, uint64(len(m.Category)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClusterLoadAssignment_Policy) MarshalVT() (dAtA []byte, err error) {
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

func (m *ClusterLoadAssignment_Policy) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ClusterLoadAssignment_Policy) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.DisableOverprovisioning {
		i--
		if m.DisableOverprovisioning {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.EndpointStaleAfter != nil {
		size, err := (*durationpb.Duration)(m.EndpointStaleAfter).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.OverprovisioningFactor != nil {
		size, err := (*wrapperspb.UInt32Value)(m.OverprovisioningFactor).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DropOverloads) > 0 {
		for iNdEx := len(m.DropOverloads) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.DropOverloads[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
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

func (m *ClusterLoadAssignment) MarshalVT() (dAtA []byte, err error) {
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

func (m *ClusterLoadAssignment) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ClusterLoadAssignment) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.NamedEndpoints) > 0 {
		for k := range m.NamedEndpoints {
			v := m.NamedEndpoints[k]
			baseI := i
			if vtmsg, ok := interface{}(v).(interface {
				MarshalToSizedBufferVT([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(v)
				if err != nil {
					return 0, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = encodeVarint(dAtA, i, uint64(len(encoded)))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarint(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarint(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Policy != nil {
		size, err := m.Policy.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Endpoints) > 0 {
		for iNdEx := len(m.Endpoints) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.Endpoints[iNdEx]).(interface {
				MarshalToSizedBufferVT([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.Endpoints[iNdEx])
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
	if len(m.ClusterName) > 0 {
		i -= len(m.ClusterName)
		copy(dAtA[i:], m.ClusterName)
		i = encodeVarint(dAtA, i, uint64(len(m.ClusterName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClusterLoadAssignment_Policy_DropOverload) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Category)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.DropPercentage != nil {
		if size, ok := interface{}(m.DropPercentage).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.DropPercentage)
		}
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ClusterLoadAssignment_Policy) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DropOverloads) > 0 {
		for _, e := range m.DropOverloads {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.OverprovisioningFactor != nil {
		l = (*wrapperspb.UInt32Value)(m.OverprovisioningFactor).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.EndpointStaleAfter != nil {
		l = (*durationpb.Duration)(m.EndpointStaleAfter).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.DisableOverprovisioning {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *ClusterLoadAssignment) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClusterName)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if len(m.Endpoints) > 0 {
		for _, e := range m.Endpoints {
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
	if m.Policy != nil {
		l = m.Policy.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if len(m.NamedEndpoints) > 0 {
		for k, v := range m.NamedEndpoints {
			_ = k
			_ = v
			l = 0
			if v != nil {
				if size, ok := interface{}(v).(interface {
					SizeVT() int
				}); ok {
					l = size.SizeVT()
				} else {
					l = proto.Size(v)
				}
			}
			l += 1 + sov(uint64(l))
			mapEntrySize := 1 + len(k) + sov(uint64(len(k))) + l
			n += mapEntrySize + 1 + sov(uint64(mapEntrySize))
		}
	}
	n += len(m.unknownFields)
	return n
}
