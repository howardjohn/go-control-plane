// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: envoy/service/tap/v2alpha/common.proto

package v2alpha

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

func (m *TapConfig) MarshalVT() (dAtA []byte, err error) {
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

func (m *TapConfig) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *TapConfig) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.TapEnabled != nil {
		if vtmsg, ok := interface{}(m.TapEnabled).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.TapEnabled)
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
	if m.OutputConfig != nil {
		size, err := m.OutputConfig.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.MatchConfig != nil {
		size, err := m.MatchConfig.MarshalToSizedBufferVT(dAtA[:i])
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

func (m *MatchPredicate_MatchSet) MarshalVT() (dAtA []byte, err error) {
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

func (m *MatchPredicate_MatchSet) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MatchPredicate_MatchSet) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.Rules) > 0 {
		for iNdEx := len(m.Rules) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Rules[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
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

func (m *MatchPredicate) MarshalVT() (dAtA []byte, err error) {
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

func (m *MatchPredicate) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MatchPredicate) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if vtmsg, ok := m.Rule.(interface {
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

func (m *MatchPredicate_OrMatch) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MatchPredicate_OrMatch) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.OrMatch != nil {
		size, err := m.OrMatch.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *MatchPredicate_AndMatch) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MatchPredicate_AndMatch) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AndMatch != nil {
		size, err := m.AndMatch.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *MatchPredicate_NotMatch) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MatchPredicate_NotMatch) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NotMatch != nil {
		size, err := m.NotMatch.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *MatchPredicate_AnyMatch) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MatchPredicate_AnyMatch) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i--
	if m.AnyMatch {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x20
	return len(dAtA) - i, nil
}
func (m *MatchPredicate_HttpRequestHeadersMatch) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MatchPredicate_HttpRequestHeadersMatch) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HttpRequestHeadersMatch != nil {
		size, err := m.HttpRequestHeadersMatch.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *MatchPredicate_HttpRequestTrailersMatch) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MatchPredicate_HttpRequestTrailersMatch) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HttpRequestTrailersMatch != nil {
		size, err := m.HttpRequestTrailersMatch.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *MatchPredicate_HttpResponseHeadersMatch) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MatchPredicate_HttpResponseHeadersMatch) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HttpResponseHeadersMatch != nil {
		size, err := m.HttpResponseHeadersMatch.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *MatchPredicate_HttpResponseTrailersMatch) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MatchPredicate_HttpResponseTrailersMatch) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HttpResponseTrailersMatch != nil {
		size, err := m.HttpResponseTrailersMatch.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *HttpHeadersMatch) MarshalVT() (dAtA []byte, err error) {
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

func (m *HttpHeadersMatch) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *HttpHeadersMatch) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
				MarshalToSizedBufferVT([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
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
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *OutputConfig) MarshalVT() (dAtA []byte, err error) {
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

func (m *OutputConfig) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OutputConfig) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.Streaming {
		i--
		if m.Streaming {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.MaxBufferedTxBytes != nil {
		size, err := (*wrapperspb.UInt32Value)(m.MaxBufferedTxBytes).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.MaxBufferedRxBytes != nil {
		size, err := (*wrapperspb.UInt32Value)(m.MaxBufferedRxBytes).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sinks) > 0 {
		for iNdEx := len(m.Sinks) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Sinks[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
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

func (m *OutputSink) MarshalVT() (dAtA []byte, err error) {
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

func (m *OutputSink) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OutputSink) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if vtmsg, ok := m.OutputSinkType.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if m.Format != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Format))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OutputSink_StreamingAdmin) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OutputSink_StreamingAdmin) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StreamingAdmin != nil {
		size, err := m.StreamingAdmin.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *OutputSink_FilePerTap) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OutputSink_FilePerTap) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.FilePerTap != nil {
		size, err := m.FilePerTap.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *OutputSink_StreamingGrpc) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OutputSink_StreamingGrpc) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StreamingGrpc != nil {
		size, err := m.StreamingGrpc.MarshalToSizedBufferVT(dAtA[:i])
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
func (m *StreamingAdminSink) MarshalVT() (dAtA []byte, err error) {
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

func (m *StreamingAdminSink) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *StreamingAdminSink) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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

func (m *FilePerTapSink) MarshalVT() (dAtA []byte, err error) {
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

func (m *FilePerTapSink) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *FilePerTapSink) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.PathPrefix) > 0 {
		i -= len(m.PathPrefix)
		copy(dAtA[i:], m.PathPrefix)
		i = encodeVarint(dAtA, i, uint64(len(m.PathPrefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StreamingGrpcSink) MarshalVT() (dAtA []byte, err error) {
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

func (m *StreamingGrpcSink) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *StreamingGrpcSink) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.GrpcService != nil {
		if vtmsg, ok := interface{}(m.GrpcService).(interface {
			MarshalToSizedBufferVT([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.GrpcService)
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
	if len(m.TapId) > 0 {
		i -= len(m.TapId)
		copy(dAtA[i:], m.TapId)
		i = encodeVarint(dAtA, i, uint64(len(m.TapId)))
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
func (m *TapConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MatchConfig != nil {
		l = m.MatchConfig.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.OutputConfig != nil {
		l = m.OutputConfig.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.TapEnabled != nil {
		if size, ok := interface{}(m.TapEnabled).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.TapEnabled)
		}
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *MatchPredicate_MatchSet) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for _, e := range m.Rules {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *MatchPredicate) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Rule.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *MatchPredicate_OrMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrMatch != nil {
		l = m.OrMatch.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *MatchPredicate_AndMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AndMatch != nil {
		l = m.AndMatch.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *MatchPredicate_NotMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NotMatch != nil {
		l = m.NotMatch.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *MatchPredicate_AnyMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *MatchPredicate_HttpRequestHeadersMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpRequestHeadersMatch != nil {
		l = m.HttpRequestHeadersMatch.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *MatchPredicate_HttpRequestTrailersMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpRequestTrailersMatch != nil {
		l = m.HttpRequestTrailersMatch.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *MatchPredicate_HttpResponseHeadersMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpResponseHeadersMatch != nil {
		l = m.HttpResponseHeadersMatch.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *MatchPredicate_HttpResponseTrailersMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpResponseTrailersMatch != nil {
		l = m.HttpResponseTrailersMatch.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *HttpHeadersMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
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

func (m *OutputConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Sinks) > 0 {
		for _, e := range m.Sinks {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.MaxBufferedRxBytes != nil {
		l = (*wrapperspb.UInt32Value)(m.MaxBufferedRxBytes).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.MaxBufferedTxBytes != nil {
		l = (*wrapperspb.UInt32Value)(m.MaxBufferedTxBytes).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.Streaming {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *OutputSink) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Format != 0 {
		n += 1 + sov(uint64(m.Format))
	}
	if vtmsg, ok := m.OutputSinkType.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *OutputSink_StreamingAdmin) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StreamingAdmin != nil {
		l = m.StreamingAdmin.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *OutputSink_FilePerTap) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FilePerTap != nil {
		l = m.FilePerTap.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *OutputSink_StreamingGrpc) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StreamingGrpc != nil {
		l = m.StreamingGrpc.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *StreamingAdminSink) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *FilePerTapSink) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PathPrefix)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *StreamingGrpcSink) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TapId)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.GrpcService != nil {
		if size, ok := interface{}(m.GrpcService).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.GrpcService)
		}
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
