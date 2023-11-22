// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: envoy/data/tap/v2alpha/transport.proto

package v2alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Connection properties.
type Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Local address.
	LocalAddress *core.Address `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	// Remote address.
	RemoteAddress *core.Address `protobuf:"bytes,3,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
}

func (x *Connection) Reset() {
	*x = Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_tap_v2alpha_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_tap_v2alpha_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_envoy_data_tap_v2alpha_transport_proto_rawDescGZIP(), []int{0}
}

func (x *Connection) GetLocalAddress() *core.Address {
	if x != nil {
		return x.LocalAddress
	}
	return nil
}

func (x *Connection) GetRemoteAddress() *core.Address {
	if x != nil {
		return x.RemoteAddress
	}
	return nil
}

// Event in a socket trace.
type SocketEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp for event.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Read or write with content as bytes string.
	//
	// Types that are assignable to EventSelector:
	//
	//	*SocketEvent_Read_
	//	*SocketEvent_Write_
	//	*SocketEvent_Closed_
	EventSelector isSocketEvent_EventSelector `protobuf_oneof:"event_selector"`
}

func (x *SocketEvent) Reset() {
	*x = SocketEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_tap_v2alpha_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocketEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocketEvent) ProtoMessage() {}

func (x *SocketEvent) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_tap_v2alpha_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocketEvent.ProtoReflect.Descriptor instead.
func (*SocketEvent) Descriptor() ([]byte, []int) {
	return file_envoy_data_tap_v2alpha_transport_proto_rawDescGZIP(), []int{1}
}

func (x *SocketEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (m *SocketEvent) GetEventSelector() isSocketEvent_EventSelector {
	if m != nil {
		return m.EventSelector
	}
	return nil
}

func (x *SocketEvent) GetRead() *SocketEvent_Read {
	if x, ok := x.GetEventSelector().(*SocketEvent_Read_); ok {
		return x.Read
	}
	return nil
}

func (x *SocketEvent) GetWrite() *SocketEvent_Write {
	if x, ok := x.GetEventSelector().(*SocketEvent_Write_); ok {
		return x.Write
	}
	return nil
}

func (x *SocketEvent) GetClosed() *SocketEvent_Closed {
	if x, ok := x.GetEventSelector().(*SocketEvent_Closed_); ok {
		return x.Closed
	}
	return nil
}

type isSocketEvent_EventSelector interface {
	isSocketEvent_EventSelector()
}

type SocketEvent_Read_ struct {
	Read *SocketEvent_Read `protobuf:"bytes,2,opt,name=read,proto3,oneof"`
}

type SocketEvent_Write_ struct {
	Write *SocketEvent_Write `protobuf:"bytes,3,opt,name=write,proto3,oneof"`
}

type SocketEvent_Closed_ struct {
	Closed *SocketEvent_Closed `protobuf:"bytes,4,opt,name=closed,proto3,oneof"`
}

func (*SocketEvent_Read_) isSocketEvent_EventSelector() {}

func (*SocketEvent_Write_) isSocketEvent_EventSelector() {}

func (*SocketEvent_Closed_) isSocketEvent_EventSelector() {}

// Sequence of read/write events that constitute a buffered trace on a socket.
// [#next-free-field: 6]
type SocketBufferedTrace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Trace ID unique to the originating Envoy only. Trace IDs can repeat and should not be used
	// for long term stable uniqueness. Matches connection IDs used in Envoy logs.
	TraceId uint64 `protobuf:"varint,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// Connection properties.
	Connection *Connection `protobuf:"bytes,2,opt,name=connection,proto3" json:"connection,omitempty"`
	// Sequence of observed events.
	Events []*SocketEvent `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	// Set to true if read events were truncated due to the :ref:`max_buffered_rx_bytes
	// <envoy_api_field_service.tap.v2alpha.OutputConfig.max_buffered_rx_bytes>` setting.
	ReadTruncated bool `protobuf:"varint,4,opt,name=read_truncated,json=readTruncated,proto3" json:"read_truncated,omitempty"`
	// Set to true if write events were truncated due to the :ref:`max_buffered_tx_bytes
	// <envoy_api_field_service.tap.v2alpha.OutputConfig.max_buffered_tx_bytes>` setting.
	WriteTruncated bool `protobuf:"varint,5,opt,name=write_truncated,json=writeTruncated,proto3" json:"write_truncated,omitempty"`
}

func (x *SocketBufferedTrace) Reset() {
	*x = SocketBufferedTrace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_tap_v2alpha_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocketBufferedTrace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocketBufferedTrace) ProtoMessage() {}

func (x *SocketBufferedTrace) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_tap_v2alpha_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocketBufferedTrace.ProtoReflect.Descriptor instead.
func (*SocketBufferedTrace) Descriptor() ([]byte, []int) {
	return file_envoy_data_tap_v2alpha_transport_proto_rawDescGZIP(), []int{2}
}

func (x *SocketBufferedTrace) GetTraceId() uint64 {
	if x != nil {
		return x.TraceId
	}
	return 0
}

func (x *SocketBufferedTrace) GetConnection() *Connection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *SocketBufferedTrace) GetEvents() []*SocketEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *SocketBufferedTrace) GetReadTruncated() bool {
	if x != nil {
		return x.ReadTruncated
	}
	return false
}

func (x *SocketBufferedTrace) GetWriteTruncated() bool {
	if x != nil {
		return x.WriteTruncated
	}
	return false
}

// A streamed socket trace segment. Multiple segments make up a full trace.
type SocketStreamedTraceSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Trace ID unique to the originating Envoy only. Trace IDs can repeat and should not be used
	// for long term stable uniqueness. Matches connection IDs used in Envoy logs.
	TraceId uint64 `protobuf:"varint,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// Types that are assignable to MessagePiece:
	//
	//	*SocketStreamedTraceSegment_Connection
	//	*SocketStreamedTraceSegment_Event
	MessagePiece isSocketStreamedTraceSegment_MessagePiece `protobuf_oneof:"message_piece"`
}

func (x *SocketStreamedTraceSegment) Reset() {
	*x = SocketStreamedTraceSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_tap_v2alpha_transport_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocketStreamedTraceSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocketStreamedTraceSegment) ProtoMessage() {}

func (x *SocketStreamedTraceSegment) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_tap_v2alpha_transport_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocketStreamedTraceSegment.ProtoReflect.Descriptor instead.
func (*SocketStreamedTraceSegment) Descriptor() ([]byte, []int) {
	return file_envoy_data_tap_v2alpha_transport_proto_rawDescGZIP(), []int{3}
}

func (x *SocketStreamedTraceSegment) GetTraceId() uint64 {
	if x != nil {
		return x.TraceId
	}
	return 0
}

func (m *SocketStreamedTraceSegment) GetMessagePiece() isSocketStreamedTraceSegment_MessagePiece {
	if m != nil {
		return m.MessagePiece
	}
	return nil
}

func (x *SocketStreamedTraceSegment) GetConnection() *Connection {
	if x, ok := x.GetMessagePiece().(*SocketStreamedTraceSegment_Connection); ok {
		return x.Connection
	}
	return nil
}

func (x *SocketStreamedTraceSegment) GetEvent() *SocketEvent {
	if x, ok := x.GetMessagePiece().(*SocketStreamedTraceSegment_Event); ok {
		return x.Event
	}
	return nil
}

type isSocketStreamedTraceSegment_MessagePiece interface {
	isSocketStreamedTraceSegment_MessagePiece()
}

type SocketStreamedTraceSegment_Connection struct {
	// Connection properties.
	Connection *Connection `protobuf:"bytes,2,opt,name=connection,proto3,oneof"`
}

type SocketStreamedTraceSegment_Event struct {
	// Socket event.
	Event *SocketEvent `protobuf:"bytes,3,opt,name=event,proto3,oneof"`
}

func (*SocketStreamedTraceSegment_Connection) isSocketStreamedTraceSegment_MessagePiece() {}

func (*SocketStreamedTraceSegment_Event) isSocketStreamedTraceSegment_MessagePiece() {}

// Data read by Envoy from the transport socket.
type SocketEvent_Read struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Binary data read.
	Data *Body `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SocketEvent_Read) Reset() {
	*x = SocketEvent_Read{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_tap_v2alpha_transport_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocketEvent_Read) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocketEvent_Read) ProtoMessage() {}

func (x *SocketEvent_Read) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_tap_v2alpha_transport_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocketEvent_Read.ProtoReflect.Descriptor instead.
func (*SocketEvent_Read) Descriptor() ([]byte, []int) {
	return file_envoy_data_tap_v2alpha_transport_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SocketEvent_Read) GetData() *Body {
	if x != nil {
		return x.Data
	}
	return nil
}

// Data written by Envoy to the transport socket.
type SocketEvent_Write struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Binary data written.
	Data *Body `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Stream was half closed after this write.
	EndStream bool `protobuf:"varint,2,opt,name=end_stream,json=endStream,proto3" json:"end_stream,omitempty"`
}

func (x *SocketEvent_Write) Reset() {
	*x = SocketEvent_Write{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_tap_v2alpha_transport_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocketEvent_Write) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocketEvent_Write) ProtoMessage() {}

func (x *SocketEvent_Write) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_tap_v2alpha_transport_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocketEvent_Write.ProtoReflect.Descriptor instead.
func (*SocketEvent_Write) Descriptor() ([]byte, []int) {
	return file_envoy_data_tap_v2alpha_transport_proto_rawDescGZIP(), []int{1, 1}
}

func (x *SocketEvent_Write) GetData() *Body {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SocketEvent_Write) GetEndStream() bool {
	if x != nil {
		return x.EndStream
	}
	return false
}

// The connection was closed.
type SocketEvent_Closed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SocketEvent_Closed) Reset() {
	*x = SocketEvent_Closed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_tap_v2alpha_transport_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocketEvent_Closed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocketEvent_Closed) ProtoMessage() {}

func (x *SocketEvent_Closed) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_tap_v2alpha_transport_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocketEvent_Closed.ProtoReflect.Descriptor instead.
func (*SocketEvent_Closed) Descriptor() ([]byte, []int) {
	return file_envoy_data_tap_v2alpha_transport_proto_rawDescGZIP(), []int{1, 2}
}

var File_envoy_data_tap_v2alpha_transport_proto protoreflect.FileDescriptor

var file_envoy_data_tap_v2alpha_transport_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x61, 0x70,
	0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x61,
	0x70, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xc0, 0x03, 0x0a, 0x0b, 0x53, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x3e, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74,
	0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x48, 0x00, 0x52, 0x04, 0x72,
	0x65, 0x61, 0x64, 0x12, 0x41, 0x0a, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x48, 0x00, 0x52,
	0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x64, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x1a, 0x38, 0x0a, 0x04,
	0x52, 0x65, 0x61, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x42, 0x6f, 0x64, 0x79,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x58, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12,
	0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x1a, 0x08, 0x0a, 0x06, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x81, 0x02, 0x0a,
	0x13, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x42, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x54, 0x72,
	0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x5f, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x22, 0xcb, 0x01, 0x0a, 0x1a, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3b, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0f, 0x0a,
	0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x69, 0x65, 0x63, 0x65, 0x42, 0x7f,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x0a, 0x24, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_data_tap_v2alpha_transport_proto_rawDescOnce sync.Once
	file_envoy_data_tap_v2alpha_transport_proto_rawDescData = file_envoy_data_tap_v2alpha_transport_proto_rawDesc
)

func file_envoy_data_tap_v2alpha_transport_proto_rawDescGZIP() []byte {
	file_envoy_data_tap_v2alpha_transport_proto_rawDescOnce.Do(func() {
		file_envoy_data_tap_v2alpha_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_data_tap_v2alpha_transport_proto_rawDescData)
	})
	return file_envoy_data_tap_v2alpha_transport_proto_rawDescData
}

var file_envoy_data_tap_v2alpha_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_envoy_data_tap_v2alpha_transport_proto_goTypes = []interface{}{
	(*Connection)(nil),                 // 0: envoy.data.tap.v2alpha.Connection
	(*SocketEvent)(nil),                // 1: envoy.data.tap.v2alpha.SocketEvent
	(*SocketBufferedTrace)(nil),        // 2: envoy.data.tap.v2alpha.SocketBufferedTrace
	(*SocketStreamedTraceSegment)(nil), // 3: envoy.data.tap.v2alpha.SocketStreamedTraceSegment
	(*SocketEvent_Read)(nil),           // 4: envoy.data.tap.v2alpha.SocketEvent.Read
	(*SocketEvent_Write)(nil),          // 5: envoy.data.tap.v2alpha.SocketEvent.Write
	(*SocketEvent_Closed)(nil),         // 6: envoy.data.tap.v2alpha.SocketEvent.Closed
	(*core.Address)(nil),               // 7: envoy.api.v2.core.Address
	(*timestamppb.Timestamp)(nil),      // 8: google.protobuf.Timestamp
	(*Body)(nil),                       // 9: envoy.data.tap.v2alpha.Body
}
var file_envoy_data_tap_v2alpha_transport_proto_depIdxs = []int32{
	7,  // 0: envoy.data.tap.v2alpha.Connection.local_address:type_name -> envoy.api.v2.core.Address
	7,  // 1: envoy.data.tap.v2alpha.Connection.remote_address:type_name -> envoy.api.v2.core.Address
	8,  // 2: envoy.data.tap.v2alpha.SocketEvent.timestamp:type_name -> google.protobuf.Timestamp
	4,  // 3: envoy.data.tap.v2alpha.SocketEvent.read:type_name -> envoy.data.tap.v2alpha.SocketEvent.Read
	5,  // 4: envoy.data.tap.v2alpha.SocketEvent.write:type_name -> envoy.data.tap.v2alpha.SocketEvent.Write
	6,  // 5: envoy.data.tap.v2alpha.SocketEvent.closed:type_name -> envoy.data.tap.v2alpha.SocketEvent.Closed
	0,  // 6: envoy.data.tap.v2alpha.SocketBufferedTrace.connection:type_name -> envoy.data.tap.v2alpha.Connection
	1,  // 7: envoy.data.tap.v2alpha.SocketBufferedTrace.events:type_name -> envoy.data.tap.v2alpha.SocketEvent
	0,  // 8: envoy.data.tap.v2alpha.SocketStreamedTraceSegment.connection:type_name -> envoy.data.tap.v2alpha.Connection
	1,  // 9: envoy.data.tap.v2alpha.SocketStreamedTraceSegment.event:type_name -> envoy.data.tap.v2alpha.SocketEvent
	9,  // 10: envoy.data.tap.v2alpha.SocketEvent.Read.data:type_name -> envoy.data.tap.v2alpha.Body
	9,  // 11: envoy.data.tap.v2alpha.SocketEvent.Write.data:type_name -> envoy.data.tap.v2alpha.Body
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_envoy_data_tap_v2alpha_transport_proto_init() }
func file_envoy_data_tap_v2alpha_transport_proto_init() {
	if File_envoy_data_tap_v2alpha_transport_proto != nil {
		return
	}
	file_envoy_data_tap_v2alpha_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_data_tap_v2alpha_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_data_tap_v2alpha_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocketEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_data_tap_v2alpha_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocketBufferedTrace); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_data_tap_v2alpha_transport_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocketStreamedTraceSegment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_data_tap_v2alpha_transport_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocketEvent_Read); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_data_tap_v2alpha_transport_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocketEvent_Write); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_data_tap_v2alpha_transport_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocketEvent_Closed); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_data_tap_v2alpha_transport_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SocketEvent_Read_)(nil),
		(*SocketEvent_Write_)(nil),
		(*SocketEvent_Closed_)(nil),
	}
	file_envoy_data_tap_v2alpha_transport_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SocketStreamedTraceSegment_Connection)(nil),
		(*SocketStreamedTraceSegment_Event)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_data_tap_v2alpha_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_data_tap_v2alpha_transport_proto_goTypes,
		DependencyIndexes: file_envoy_data_tap_v2alpha_transport_proto_depIdxs,
		MessageInfos:      file_envoy_data_tap_v2alpha_transport_proto_msgTypes,
	}.Build()
	File_envoy_data_tap_v2alpha_transport_proto = out.File
	file_envoy_data_tap_v2alpha_transport_proto_rawDesc = nil
	file_envoy_data_tap_v2alpha_transport_proto_goTypes = nil
	file_envoy_data_tap_v2alpha_transport_proto_depIdxs = nil
}
