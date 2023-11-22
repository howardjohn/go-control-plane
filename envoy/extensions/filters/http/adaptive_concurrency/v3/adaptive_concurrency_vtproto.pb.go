// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: envoy/extensions/filters/http/adaptive_concurrency/v3/adaptive_concurrency.proto

package adaptive_concurrencyv3

import (
	durationpb "github.com/planetscale/vtprotobuf/types/known/durationpb"
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

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) MarshalVT() (dAtA []byte, err error) {
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

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.ConcurrencyUpdateInterval != nil {
		size, err := (*durationpb.Duration)(m.ConcurrencyUpdateInterval).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.MaxConcurrencyLimit != nil {
		size, err := (*wrapperspb.UInt32Value)(m.MaxConcurrencyLimit).MarshalToSizedBufferVT(dAtA[:i])
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

func (m *GradientControllerConfig_MinimumRTTCalculationParams) MarshalVT() (dAtA []byte, err error) {
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

func (m *GradientControllerConfig_MinimumRTTCalculationParams) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.Buffer != nil {
		if vtmsg, ok := interface{}(m.Buffer).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Buffer)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.MinConcurrency != nil {
		size, err := (*wrapperspb.UInt32Value)(m.MinConcurrency).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.Jitter != nil {
		if vtmsg, ok := interface{}(m.Jitter).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Jitter)
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
	if m.RequestCount != nil {
		size, err := (*wrapperspb.UInt32Value)(m.RequestCount).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.Interval != nil {
		size, err := (*durationpb.Duration)(m.Interval).MarshalToSizedBufferVT(dAtA[:i])
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

func (m *GradientControllerConfig) MarshalVT() (dAtA []byte, err error) {
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

func (m *GradientControllerConfig) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *GradientControllerConfig) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.MinRttCalcParams != nil {
		size, err := m.MinRttCalcParams.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.ConcurrencyLimitParams != nil {
		size, err := m.ConcurrencyLimitParams.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.SampleAggregatePercentile != nil {
		if vtmsg, ok := interface{}(m.SampleAggregatePercentile).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.SampleAggregatePercentile)
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

func (m *AdaptiveConcurrency) MarshalVT() (dAtA []byte, err error) {
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

func (m *AdaptiveConcurrency) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *AdaptiveConcurrency) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if vtmsg, ok := m.ConcurrencyControllerConfig.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if m.ConcurrencyLimitExceededStatus != nil {
		if vtmsg, ok := interface{}(m.ConcurrencyLimitExceededStatus).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.ConcurrencyLimitExceededStatus)
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
	if m.Enabled != nil {
		if vtmsg, ok := interface{}(m.Enabled).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Enabled)
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

func (m *AdaptiveConcurrency_GradientControllerConfig) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *AdaptiveConcurrency_GradientControllerConfig) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.GradientControllerConfig != nil {
		size, err := m.GradientControllerConfig.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxConcurrencyLimit != nil {
		l = (*wrapperspb.UInt32Value)(m.MaxConcurrencyLimit).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.ConcurrencyUpdateInterval != nil {
		l = (*durationpb.Duration)(m.ConcurrencyUpdateInterval).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Interval != nil {
		l = (*durationpb.Duration)(m.Interval).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.RequestCount != nil {
		l = (*wrapperspb.UInt32Value)(m.RequestCount).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.Jitter != nil {
		if size, ok := interface{}(m.Jitter).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Jitter)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.MinConcurrency != nil {
		l = (*wrapperspb.UInt32Value)(m.MinConcurrency).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.Buffer != nil {
		if size, ok := interface{}(m.Buffer).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Buffer)
		}
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *GradientControllerConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SampleAggregatePercentile != nil {
		if size, ok := interface{}(m.SampleAggregatePercentile).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.SampleAggregatePercentile)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.ConcurrencyLimitParams != nil {
		l = m.ConcurrencyLimitParams.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.MinRttCalcParams != nil {
		l = m.MinRttCalcParams.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *AdaptiveConcurrency) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.ConcurrencyControllerConfig.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if m.Enabled != nil {
		if size, ok := interface{}(m.Enabled).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Enabled)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.ConcurrencyLimitExceededStatus != nil {
		if size, ok := interface{}(m.ConcurrencyLimitExceededStatus).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.ConcurrencyLimitExceededStatus)
		}
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *AdaptiveConcurrency_GradientControllerConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GradientControllerConfig != nil {
		l = m.GradientControllerConfig.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
