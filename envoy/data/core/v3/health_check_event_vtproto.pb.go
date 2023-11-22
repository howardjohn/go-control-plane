// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: envoy/data/core/v3/health_check_event.proto

package corev3

import (
	timestamppb "github.com/planetscale/vtprotobuf/types/known/timestamppb"
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

func (m *HealthCheckEvent) MarshalVT() (dAtA []byte, err error) {
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

func (m *HealthCheckEvent) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *HealthCheckEvent) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if vtmsg, ok := m.Event.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if m.Locality != nil {
		if vtmsg, ok := interface{}(m.Locality).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Locality)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x5a
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
		dAtA[i] = 0x52
	}
	if m.Timestamp != nil {
		size, err := (*timestamppb.Timestamp)(m.Timestamp).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ClusterName) > 0 {
		i -= len(m.ClusterName)
		copy(dAtA[i:], m.ClusterName)
		i = encodeVarint(dAtA, i, uint64(len(m.ClusterName)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Host != nil {
		if vtmsg, ok := interface{}(m.Host).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Host)
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
	if m.HealthCheckerType != 0 {
		i = encodeVarint(dAtA, i, uint64(m.HealthCheckerType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *HealthCheckEvent_EjectUnhealthyEvent) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *HealthCheckEvent_EjectUnhealthyEvent) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EjectUnhealthyEvent != nil {
		size, err := m.EjectUnhealthyEvent.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *HealthCheckEvent_AddHealthyEvent) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *HealthCheckEvent_AddHealthyEvent) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AddHealthyEvent != nil {
		size, err := m.AddHealthyEvent.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *HealthCheckEvent_HealthCheckFailureEvent) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *HealthCheckEvent_HealthCheckFailureEvent) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HealthCheckFailureEvent != nil {
		size, err := m.HealthCheckFailureEvent.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *HealthCheckEvent_DegradedHealthyHost) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *HealthCheckEvent_DegradedHealthyHost) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DegradedHealthyHost != nil {
		size, err := m.DegradedHealthyHost.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *HealthCheckEvent_NoLongerDegradedHost) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *HealthCheckEvent_NoLongerDegradedHost) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoLongerDegradedHost != nil {
		size, err := m.NoLongerDegradedHost.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func (m *HealthCheckEjectUnhealthy) MarshalVT() (dAtA []byte, err error) {
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

func (m *HealthCheckEjectUnhealthy) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *HealthCheckEjectUnhealthy) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.FailureType != 0 {
		i = encodeVarint(dAtA, i, uint64(m.FailureType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *HealthCheckAddHealthy) MarshalVT() (dAtA []byte, err error) {
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

func (m *HealthCheckAddHealthy) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *HealthCheckAddHealthy) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.FirstCheck {
		i--
		if m.FirstCheck {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *HealthCheckFailure) MarshalVT() (dAtA []byte, err error) {
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

func (m *HealthCheckFailure) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *HealthCheckFailure) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.FirstCheck {
		i--
		if m.FirstCheck {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.FailureType != 0 {
		i = encodeVarint(dAtA, i, uint64(m.FailureType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DegradedHealthyHost) MarshalVT() (dAtA []byte, err error) {
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

func (m *DegradedHealthyHost) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *DegradedHealthyHost) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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

func (m *NoLongerDegradedHost) MarshalVT() (dAtA []byte, err error) {
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

func (m *NoLongerDegradedHost) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *NoLongerDegradedHost) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
func (m *HealthCheckEvent) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HealthCheckerType != 0 {
		n += 1 + sov(uint64(m.HealthCheckerType))
	}
	if m.Host != nil {
		if size, ok := interface{}(m.Host).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Host)
		}
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.ClusterName)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if vtmsg, ok := m.Event.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if m.Timestamp != nil {
		l = (*timestamppb.Timestamp)(m.Timestamp).SizeVT()
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
	if m.Locality != nil {
		if size, ok := interface{}(m.Locality).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Locality)
		}
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *HealthCheckEvent_EjectUnhealthyEvent) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EjectUnhealthyEvent != nil {
		l = m.EjectUnhealthyEvent.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *HealthCheckEvent_AddHealthyEvent) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AddHealthyEvent != nil {
		l = m.AddHealthyEvent.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *HealthCheckEvent_HealthCheckFailureEvent) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HealthCheckFailureEvent != nil {
		l = m.HealthCheckFailureEvent.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *HealthCheckEvent_DegradedHealthyHost) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DegradedHealthyHost != nil {
		l = m.DegradedHealthyHost.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *HealthCheckEvent_NoLongerDegradedHost) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoLongerDegradedHost != nil {
		l = m.NoLongerDegradedHost.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *HealthCheckEjectUnhealthy) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FailureType != 0 {
		n += 1 + sov(uint64(m.FailureType))
	}
	n += len(m.unknownFields)
	return n
}

func (m *HealthCheckAddHealthy) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FirstCheck {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *HealthCheckFailure) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FailureType != 0 {
		n += 1 + sov(uint64(m.FailureType))
	}
	if m.FirstCheck {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *DegradedHealthyHost) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *NoLongerDegradedHost) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
