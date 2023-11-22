// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: envoy/data/dns/v3/dns_table.proto

package dnsv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This message contains the configuration for the DNS Filter if populated
// from the control plane
type DnsTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Control how many times Envoy makes an attempt to forward a query to an external DNS server
	ExternalRetryCount uint32 `protobuf:"varint,1,opt,name=external_retry_count,json=externalRetryCount,proto3" json:"external_retry_count,omitempty"`
	// Fully qualified domain names for which Envoy will respond to DNS queries. By leaving this
	// list empty, Envoy will forward all queries to external resolvers
	VirtualDomains []*DnsTable_DnsVirtualDomain `protobuf:"bytes,2,rep,name=virtual_domains,json=virtualDomains,proto3" json:"virtual_domains,omitempty"`
	// This field is deprecated and no longer used in Envoy. The filter's behavior has changed
	// internally to use a different data structure allowing the filter to determine whether a
	// query is for known domain without the use of this field.
	//
	// This field serves to help Envoy determine whether it can authoritatively answer a query
	// for a name matching a suffix in this list. If the query name does not match a suffix in
	// this list, Envoy will forward the query to an upstream DNS server
	//
	// Deprecated: Marked as deprecated in envoy/data/dns/v3/dns_table.proto.
	KnownSuffixes []*v3.StringMatcher `protobuf:"bytes,3,rep,name=known_suffixes,json=knownSuffixes,proto3" json:"known_suffixes,omitempty"`
}

func (x *DnsTable) Reset() {
	*x = DnsTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsTable) ProtoMessage() {}

func (x *DnsTable) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsTable.ProtoReflect.Descriptor instead.
func (*DnsTable) Descriptor() ([]byte, []int) {
	return file_envoy_data_dns_v3_dns_table_proto_rawDescGZIP(), []int{0}
}

func (x *DnsTable) GetExternalRetryCount() uint32 {
	if x != nil {
		return x.ExternalRetryCount
	}
	return 0
}

func (x *DnsTable) GetVirtualDomains() []*DnsTable_DnsVirtualDomain {
	if x != nil {
		return x.VirtualDomains
	}
	return nil
}

// Deprecated: Marked as deprecated in envoy/data/dns/v3/dns_table.proto.
func (x *DnsTable) GetKnownSuffixes() []*v3.StringMatcher {
	if x != nil {
		return x.KnownSuffixes
	}
	return nil
}

// This message contains a list of IP addresses returned for a query for a known name
type DnsTable_AddressList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field contains a well formed IP address that is returned in the answer for a
	// name query. The address field can be an IPv4 or IPv6 address. Address family
	// detection is done automatically when Envoy parses the string. Since this field is
	// repeated, Envoy will return as many entries from this list in the DNS response while
	// keeping the response under 512 bytes
	Address []string `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty"`
}

func (x *DnsTable_AddressList) Reset() {
	*x = DnsTable_AddressList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsTable_AddressList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsTable_AddressList) ProtoMessage() {}

func (x *DnsTable_AddressList) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsTable_AddressList.ProtoReflect.Descriptor instead.
func (*DnsTable_AddressList) Descriptor() ([]byte, []int) {
	return file_envoy_data_dns_v3_dns_table_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DnsTable_AddressList) GetAddress() []string {
	if x != nil {
		return x.Address
	}
	return nil
}

// Specify the service protocol using a numeric or string value
type DnsTable_DnsServiceProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ProtocolConfig:
	//
	//	*DnsTable_DnsServiceProtocol_Number
	//	*DnsTable_DnsServiceProtocol_Name
	ProtocolConfig isDnsTable_DnsServiceProtocol_ProtocolConfig `protobuf_oneof:"protocol_config"`
}

func (x *DnsTable_DnsServiceProtocol) Reset() {
	*x = DnsTable_DnsServiceProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsTable_DnsServiceProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsTable_DnsServiceProtocol) ProtoMessage() {}

func (x *DnsTable_DnsServiceProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsTable_DnsServiceProtocol.ProtoReflect.Descriptor instead.
func (*DnsTable_DnsServiceProtocol) Descriptor() ([]byte, []int) {
	return file_envoy_data_dns_v3_dns_table_proto_rawDescGZIP(), []int{0, 1}
}

func (m *DnsTable_DnsServiceProtocol) GetProtocolConfig() isDnsTable_DnsServiceProtocol_ProtocolConfig {
	if m != nil {
		return m.ProtocolConfig
	}
	return nil
}

func (x *DnsTable_DnsServiceProtocol) GetNumber() uint32 {
	if x, ok := x.GetProtocolConfig().(*DnsTable_DnsServiceProtocol_Number); ok {
		return x.Number
	}
	return 0
}

func (x *DnsTable_DnsServiceProtocol) GetName() string {
	if x, ok := x.GetProtocolConfig().(*DnsTable_DnsServiceProtocol_Name); ok {
		return x.Name
	}
	return ""
}

type isDnsTable_DnsServiceProtocol_ProtocolConfig interface {
	isDnsTable_DnsServiceProtocol_ProtocolConfig()
}

type DnsTable_DnsServiceProtocol_Number struct {
	// Specify the protocol number for the service. Envoy will try to resolve the number to
	// the protocol name. For example, 6 will resolve to "tcp". Refer to:
	// https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	// for protocol names and numbers
	Number uint32 `protobuf:"varint,1,opt,name=number,proto3,oneof"`
}

type DnsTable_DnsServiceProtocol_Name struct {
	// Specify the protocol name for the service.
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*DnsTable_DnsServiceProtocol_Number) isDnsTable_DnsServiceProtocol_ProtocolConfig() {}

func (*DnsTable_DnsServiceProtocol_Name) isDnsTable_DnsServiceProtocol_ProtocolConfig() {}

// Specify the target for a given DNS service
// [#next-free-field: 6]
type DnsTable_DnsServiceTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify the name of the endpoint for the Service. The name is a hostname or a cluster
	//
	// Types that are assignable to EndpointType:
	//
	//	*DnsTable_DnsServiceTarget_HostName
	//	*DnsTable_DnsServiceTarget_ClusterName
	EndpointType isDnsTable_DnsServiceTarget_EndpointType `protobuf_oneof:"endpoint_type"`
	// The priority of the service record target
	Priority uint32 `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	// The weight of the service record target
	Weight uint32 `protobuf:"varint,4,opt,name=weight,proto3" json:"weight,omitempty"`
	// The port to which the service is bound. This value is optional if the target is a
	// cluster. Setting port to zero in this case makes the filter use the port value
	// from the cluster host
	Port uint32 `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *DnsTable_DnsServiceTarget) Reset() {
	*x = DnsTable_DnsServiceTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsTable_DnsServiceTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsTable_DnsServiceTarget) ProtoMessage() {}

func (x *DnsTable_DnsServiceTarget) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsTable_DnsServiceTarget.ProtoReflect.Descriptor instead.
func (*DnsTable_DnsServiceTarget) Descriptor() ([]byte, []int) {
	return file_envoy_data_dns_v3_dns_table_proto_rawDescGZIP(), []int{0, 2}
}

func (m *DnsTable_DnsServiceTarget) GetEndpointType() isDnsTable_DnsServiceTarget_EndpointType {
	if m != nil {
		return m.EndpointType
	}
	return nil
}

func (x *DnsTable_DnsServiceTarget) GetHostName() string {
	if x, ok := x.GetEndpointType().(*DnsTable_DnsServiceTarget_HostName); ok {
		return x.HostName
	}
	return ""
}

func (x *DnsTable_DnsServiceTarget) GetClusterName() string {
	if x, ok := x.GetEndpointType().(*DnsTable_DnsServiceTarget_ClusterName); ok {
		return x.ClusterName
	}
	return ""
}

func (x *DnsTable_DnsServiceTarget) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *DnsTable_DnsServiceTarget) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *DnsTable_DnsServiceTarget) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type isDnsTable_DnsServiceTarget_EndpointType interface {
	isDnsTable_DnsServiceTarget_EndpointType()
}

type DnsTable_DnsServiceTarget_HostName struct {
	// Use a resolvable hostname as the endpoint for a service.
	HostName string `protobuf:"bytes,1,opt,name=host_name,json=hostName,proto3,oneof"`
}

type DnsTable_DnsServiceTarget_ClusterName struct {
	// Use a cluster name as the endpoint for a service.
	ClusterName string `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName,proto3,oneof"`
}

func (*DnsTable_DnsServiceTarget_HostName) isDnsTable_DnsServiceTarget_EndpointType() {}

func (*DnsTable_DnsServiceTarget_ClusterName) isDnsTable_DnsServiceTarget_EndpointType() {}

// This message defines a service selection record returned for a service query in a domain
type DnsTable_DnsService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the service without the protocol or domain name
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The service protocol. This can be specified as a string or the numeric value of the protocol
	Protocol *DnsTable_DnsServiceProtocol `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// The service entry time to live. This is independent from the DNS Answer record TTL
	Ttl *durationpb.Duration `protobuf:"bytes,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// The list of targets hosting the service
	Targets []*DnsTable_DnsServiceTarget `protobuf:"bytes,4,rep,name=targets,proto3" json:"targets,omitempty"`
}

func (x *DnsTable_DnsService) Reset() {
	*x = DnsTable_DnsService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsTable_DnsService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsTable_DnsService) ProtoMessage() {}

func (x *DnsTable_DnsService) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsTable_DnsService.ProtoReflect.Descriptor instead.
func (*DnsTable_DnsService) Descriptor() ([]byte, []int) {
	return file_envoy_data_dns_v3_dns_table_proto_rawDescGZIP(), []int{0, 3}
}

func (x *DnsTable_DnsService) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *DnsTable_DnsService) GetProtocol() *DnsTable_DnsServiceProtocol {
	if x != nil {
		return x.Protocol
	}
	return nil
}

func (x *DnsTable_DnsService) GetTtl() *durationpb.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

func (x *DnsTable_DnsService) GetTargets() []*DnsTable_DnsServiceTarget {
	if x != nil {
		return x.Targets
	}
	return nil
}

// Define a list of service records for a given service
type DnsTable_DnsServiceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*DnsTable_DnsService `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *DnsTable_DnsServiceList) Reset() {
	*x = DnsTable_DnsServiceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsTable_DnsServiceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsTable_DnsServiceList) ProtoMessage() {}

func (x *DnsTable_DnsServiceList) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsTable_DnsServiceList.ProtoReflect.Descriptor instead.
func (*DnsTable_DnsServiceList) Descriptor() ([]byte, []int) {
	return file_envoy_data_dns_v3_dns_table_proto_rawDescGZIP(), []int{0, 4}
}

func (x *DnsTable_DnsServiceList) GetServices() []*DnsTable_DnsService {
	if x != nil {
		return x.Services
	}
	return nil
}

type DnsTable_DnsEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to EndpointConfig:
	//
	//	*DnsTable_DnsEndpoint_AddressList
	//	*DnsTable_DnsEndpoint_ClusterName
	//	*DnsTable_DnsEndpoint_ServiceList
	EndpointConfig isDnsTable_DnsEndpoint_EndpointConfig `protobuf_oneof:"endpoint_config"`
}

func (x *DnsTable_DnsEndpoint) Reset() {
	*x = DnsTable_DnsEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsTable_DnsEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsTable_DnsEndpoint) ProtoMessage() {}

func (x *DnsTable_DnsEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsTable_DnsEndpoint.ProtoReflect.Descriptor instead.
func (*DnsTable_DnsEndpoint) Descriptor() ([]byte, []int) {
	return file_envoy_data_dns_v3_dns_table_proto_rawDescGZIP(), []int{0, 5}
}

func (m *DnsTable_DnsEndpoint) GetEndpointConfig() isDnsTable_DnsEndpoint_EndpointConfig {
	if m != nil {
		return m.EndpointConfig
	}
	return nil
}

func (x *DnsTable_DnsEndpoint) GetAddressList() *DnsTable_AddressList {
	if x, ok := x.GetEndpointConfig().(*DnsTable_DnsEndpoint_AddressList); ok {
		return x.AddressList
	}
	return nil
}

func (x *DnsTable_DnsEndpoint) GetClusterName() string {
	if x, ok := x.GetEndpointConfig().(*DnsTable_DnsEndpoint_ClusterName); ok {
		return x.ClusterName
	}
	return ""
}

func (x *DnsTable_DnsEndpoint) GetServiceList() *DnsTable_DnsServiceList {
	if x, ok := x.GetEndpointConfig().(*DnsTable_DnsEndpoint_ServiceList); ok {
		return x.ServiceList
	}
	return nil
}

type isDnsTable_DnsEndpoint_EndpointConfig interface {
	isDnsTable_DnsEndpoint_EndpointConfig()
}

type DnsTable_DnsEndpoint_AddressList struct {
	// Define a list of addresses to return for the specified endpoint
	AddressList *DnsTable_AddressList `protobuf:"bytes,1,opt,name=address_list,json=addressList,proto3,oneof"`
}

type DnsTable_DnsEndpoint_ClusterName struct {
	// Define a cluster whose addresses are returned for the specified endpoint
	ClusterName string `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName,proto3,oneof"`
}

type DnsTable_DnsEndpoint_ServiceList struct {
	// Define a DNS Service List for the specified endpoint
	ServiceList *DnsTable_DnsServiceList `protobuf:"bytes,3,opt,name=service_list,json=serviceList,proto3,oneof"`
}

func (*DnsTable_DnsEndpoint_AddressList) isDnsTable_DnsEndpoint_EndpointConfig() {}

func (*DnsTable_DnsEndpoint_ClusterName) isDnsTable_DnsEndpoint_EndpointConfig() {}

func (*DnsTable_DnsEndpoint_ServiceList) isDnsTable_DnsEndpoint_EndpointConfig() {}

type DnsTable_DnsVirtualDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A domain name for which Envoy will respond to query requests
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The configuration containing the method to determine the address of this endpoint
	Endpoint *DnsTable_DnsEndpoint `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Sets the TTL in DNS answers from Envoy returned to the client. The default TTL is 300s
	AnswerTtl *durationpb.Duration `protobuf:"bytes,3,opt,name=answer_ttl,json=answerTtl,proto3" json:"answer_ttl,omitempty"`
}

func (x *DnsTable_DnsVirtualDomain) Reset() {
	*x = DnsTable_DnsVirtualDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsTable_DnsVirtualDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsTable_DnsVirtualDomain) ProtoMessage() {}

func (x *DnsTable_DnsVirtualDomain) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_dns_v3_dns_table_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsTable_DnsVirtualDomain.ProtoReflect.Descriptor instead.
func (*DnsTable_DnsVirtualDomain) Descriptor() ([]byte, []int) {
	return file_envoy_data_dns_v3_dns_table_proto_rawDescGZIP(), []int{0, 6}
}

func (x *DnsTable_DnsVirtualDomain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DnsTable_DnsVirtualDomain) GetEndpoint() *DnsTable_DnsEndpoint {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *DnsTable_DnsVirtualDomain) GetAnswerTtl() *durationpb.Duration {
	if x != nil {
		return x.AnswerTtl
	}
	return nil
}

var File_envoy_data_dns_v3_dns_table_proto protoreflect.FileDescriptor

var file_envoy_data_dns_v3_dns_table_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x6e, 0x73,
	0x2f, 0x76, 0x33, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x64, 0x6e, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65,
	0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x0c, 0x0a, 0x08, 0x44,
	0x6e, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x03, 0x52, 0x12,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x55, 0x0a, 0x0f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x44, 0x6e, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x56, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x0e, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x58, 0x0a, 0x0e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x0b, 0x92, 0xc7, 0x86, 0xd8, 0x04, 0x03, 0x33,
	0x2e, 0x30, 0x18, 0x01, 0x52, 0x0d, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x53, 0x75, 0x66, 0x66, 0x69,
	0x78, 0x65, 0x73, 0x1a, 0x6b, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x92, 0x01, 0x08, 0x08, 0x01, 0x22, 0x04, 0x72,
	0x02, 0x10, 0x03, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x32, 0x9a, 0xc5,
	0x88, 0x1e, 0x2d, 0x0a, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x64, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x6e, 0x73, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x72, 0x0a, 0x12, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x22, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x2a, 0x03, 0x10, 0xff, 0x01,
	0x48, 0x00, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10,
	0x01, 0xc0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x16, 0x0a, 0x0f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x03, 0xf8, 0x42, 0x01, 0x1a, 0xed, 0x01, 0x0a, 0x10, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0x72, 0x05, 0x10, 0x01, 0xc0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72,
	0x05, 0x10, 0x01, 0xc0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x10, 0x80,
	0x80, 0x04, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x2a, 0x04, 0x10, 0x80, 0x80, 0x04, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x1d, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x2a, 0x04, 0x10, 0x80, 0x80, 0x04, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x14,
	0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x03, 0xf8, 0x42, 0x01, 0x1a, 0x92, 0x02, 0x0a, 0x0a, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05,
	0x10, 0x01, 0xc0, 0x01, 0x01, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6e, 0x73, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x37,
	0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0xaa, 0x01, 0x04, 0x32, 0x02,
	0x08, 0x01, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x50, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6e, 0x73,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01,
	0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x1a, 0x5e, 0x0a, 0x0e, 0x44, 0x6e, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76,
	0x33, 0x2e, 0x44, 0x6e, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x9d, 0x02, 0x0a, 0x0b, 0x44, 0x6e,
	0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x0c, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x64, 0x6e, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6e, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x64, 0x6e, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6e, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0x32, 0x9a,
	0xc5, 0x88, 0x1e, 0x2d, 0x0a, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x6e, 0x73,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x42, 0x16, 0x0a, 0x0f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x1a, 0xf6, 0x01, 0x0a, 0x10, 0x44, 0x6e,
	0x73, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1e,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0x72, 0x05, 0x10, 0x01, 0xc0, 0x01, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x43,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x64, 0x6e,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6e, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x44, 0x6e,
	0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x0a, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x74, 0x74,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0xaa, 0x01, 0x04, 0x32, 0x02, 0x08, 0x1e, 0x52, 0x09,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x54, 0x74, 0x6c, 0x3a, 0x37, 0x9a, 0xc5, 0x88, 0x1e, 0x32,
	0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x64, 0x6e, 0x73,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x6e, 0x73, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x44, 0x6e, 0x73, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x3a, 0x26, 0x9a, 0xc5, 0x88, 0x1e, 0x21, 0x0a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x44, 0x6e, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x7a, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x0a, 0x1f, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x64,
	0x6e, 0x73, 0x2e, 0x76, 0x33, 0x42, 0x0d, 0x44, 0x6e, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f,
	0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x76, 0x33,
	0x3b, 0x64, 0x6e, 0x73, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_data_dns_v3_dns_table_proto_rawDescOnce sync.Once
	file_envoy_data_dns_v3_dns_table_proto_rawDescData = file_envoy_data_dns_v3_dns_table_proto_rawDesc
)

func file_envoy_data_dns_v3_dns_table_proto_rawDescGZIP() []byte {
	file_envoy_data_dns_v3_dns_table_proto_rawDescOnce.Do(func() {
		file_envoy_data_dns_v3_dns_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_data_dns_v3_dns_table_proto_rawDescData)
	})
	return file_envoy_data_dns_v3_dns_table_proto_rawDescData
}

var file_envoy_data_dns_v3_dns_table_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_envoy_data_dns_v3_dns_table_proto_goTypes = []interface{}{
	(*DnsTable)(nil),                    // 0: envoy.data.dns.v3.DnsTable
	(*DnsTable_AddressList)(nil),        // 1: envoy.data.dns.v3.DnsTable.AddressList
	(*DnsTable_DnsServiceProtocol)(nil), // 2: envoy.data.dns.v3.DnsTable.DnsServiceProtocol
	(*DnsTable_DnsServiceTarget)(nil),   // 3: envoy.data.dns.v3.DnsTable.DnsServiceTarget
	(*DnsTable_DnsService)(nil),         // 4: envoy.data.dns.v3.DnsTable.DnsService
	(*DnsTable_DnsServiceList)(nil),     // 5: envoy.data.dns.v3.DnsTable.DnsServiceList
	(*DnsTable_DnsEndpoint)(nil),        // 6: envoy.data.dns.v3.DnsTable.DnsEndpoint
	(*DnsTable_DnsVirtualDomain)(nil),   // 7: envoy.data.dns.v3.DnsTable.DnsVirtualDomain
	(*v3.StringMatcher)(nil),            // 8: envoy.type.matcher.v3.StringMatcher
	(*durationpb.Duration)(nil),         // 9: google.protobuf.Duration
}
var file_envoy_data_dns_v3_dns_table_proto_depIdxs = []int32{
	7,  // 0: envoy.data.dns.v3.DnsTable.virtual_domains:type_name -> envoy.data.dns.v3.DnsTable.DnsVirtualDomain
	8,  // 1: envoy.data.dns.v3.DnsTable.known_suffixes:type_name -> envoy.type.matcher.v3.StringMatcher
	2,  // 2: envoy.data.dns.v3.DnsTable.DnsService.protocol:type_name -> envoy.data.dns.v3.DnsTable.DnsServiceProtocol
	9,  // 3: envoy.data.dns.v3.DnsTable.DnsService.ttl:type_name -> google.protobuf.Duration
	3,  // 4: envoy.data.dns.v3.DnsTable.DnsService.targets:type_name -> envoy.data.dns.v3.DnsTable.DnsServiceTarget
	4,  // 5: envoy.data.dns.v3.DnsTable.DnsServiceList.services:type_name -> envoy.data.dns.v3.DnsTable.DnsService
	1,  // 6: envoy.data.dns.v3.DnsTable.DnsEndpoint.address_list:type_name -> envoy.data.dns.v3.DnsTable.AddressList
	5,  // 7: envoy.data.dns.v3.DnsTable.DnsEndpoint.service_list:type_name -> envoy.data.dns.v3.DnsTable.DnsServiceList
	6,  // 8: envoy.data.dns.v3.DnsTable.DnsVirtualDomain.endpoint:type_name -> envoy.data.dns.v3.DnsTable.DnsEndpoint
	9,  // 9: envoy.data.dns.v3.DnsTable.DnsVirtualDomain.answer_ttl:type_name -> google.protobuf.Duration
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_envoy_data_dns_v3_dns_table_proto_init() }
func file_envoy_data_dns_v3_dns_table_proto_init() {
	if File_envoy_data_dns_v3_dns_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_data_dns_v3_dns_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsTable); i {
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
		file_envoy_data_dns_v3_dns_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsTable_AddressList); i {
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
		file_envoy_data_dns_v3_dns_table_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsTable_DnsServiceProtocol); i {
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
		file_envoy_data_dns_v3_dns_table_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsTable_DnsServiceTarget); i {
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
		file_envoy_data_dns_v3_dns_table_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsTable_DnsService); i {
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
		file_envoy_data_dns_v3_dns_table_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsTable_DnsServiceList); i {
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
		file_envoy_data_dns_v3_dns_table_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsTable_DnsEndpoint); i {
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
		file_envoy_data_dns_v3_dns_table_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsTable_DnsVirtualDomain); i {
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
	file_envoy_data_dns_v3_dns_table_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*DnsTable_DnsServiceProtocol_Number)(nil),
		(*DnsTable_DnsServiceProtocol_Name)(nil),
	}
	file_envoy_data_dns_v3_dns_table_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*DnsTable_DnsServiceTarget_HostName)(nil),
		(*DnsTable_DnsServiceTarget_ClusterName)(nil),
	}
	file_envoy_data_dns_v3_dns_table_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*DnsTable_DnsEndpoint_AddressList)(nil),
		(*DnsTable_DnsEndpoint_ClusterName)(nil),
		(*DnsTable_DnsEndpoint_ServiceList)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_data_dns_v3_dns_table_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_data_dns_v3_dns_table_proto_goTypes,
		DependencyIndexes: file_envoy_data_dns_v3_dns_table_proto_depIdxs,
		MessageInfos:      file_envoy_data_dns_v3_dns_table_proto_msgTypes,
	}.Build()
	File_envoy_data_dns_v3_dns_table_proto = out.File
	file_envoy_data_dns_v3_dns_table_proto_rawDesc = nil
	file_envoy_data_dns_v3_dns_table_proto_goTypes = nil
	file_envoy_data_dns_v3_dns_table_proto_depIdxs = nil
}
