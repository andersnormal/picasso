// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: pkg/proto/plugin.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Execute_Status int32

const (
	Execute_UNKNOWN Execute_Status = 0
	Execute_SUCCESS Execute_Status = 1
	Execute_FAILURE Execute_Status = 2
)

// Enum value maps for Execute_Status.
var (
	Execute_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "SUCCESS",
		2: "FAILURE",
	}
	Execute_Status_value = map[string]int32{
		"UNKNOWN": 0,
		"SUCCESS": 1,
		"FAILURE": 2,
	}
)

func (x Execute_Status) Enum() *Execute_Status {
	p := new(Execute_Status)
	*p = x
	return p
}

func (x Execute_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Execute_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_plugin_proto_enumTypes[0].Descriptor()
}

func (Execute_Status) Type() protoreflect.EnumType {
	return &file_pkg_proto_plugin_proto_enumTypes[0]
}

func (x Execute_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Execute_Status.Descriptor instead.
func (Execute_Status) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{0, 0}
}

type Diagnostic_Severity int32

const (
	Diagnostic_INVALID Diagnostic_Severity = 0
	Diagnostic_ERROR   Diagnostic_Severity = 1
	Diagnostic_WARNING Diagnostic_Severity = 2
)

// Enum value maps for Diagnostic_Severity.
var (
	Diagnostic_Severity_name = map[int32]string{
		0: "INVALID",
		1: "ERROR",
		2: "WARNING",
	}
	Diagnostic_Severity_value = map[string]int32{
		"INVALID": 0,
		"ERROR":   1,
		"WARNING": 2,
	}
)

func (x Diagnostic_Severity) Enum() *Diagnostic_Severity {
	p := new(Diagnostic_Severity)
	*p = x
	return p
}

func (x Diagnostic_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Diagnostic_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_plugin_proto_enumTypes[1].Descriptor()
}

func (Diagnostic_Severity) Type() protoreflect.EnumType {
	return &file_pkg_proto_plugin_proto_enumTypes[1]
}

func (x Diagnostic_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Diagnostic_Severity.Descriptor instead.
func (Diagnostic_Severity) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{2, 0}
}

// Start ...
type Execute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Execute) Reset() {
	*x = Execute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Execute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Execute) ProtoMessage() {}

func (x *Execute) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Execute.ProtoReflect.Descriptor instead.
func (*Execute) Descriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{0}
}

// Stop ...
type Stop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Stop) Reset() {
	*x = Stop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stop) ProtoMessage() {}

func (x *Stop) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stop.ProtoReflect.Descriptor instead.
func (*Stop) Descriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{1}
}

// Diagnostic ...
type Diagnostic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Severity  Diagnostic_Severity `protobuf:"varint,1,opt,name=severity,proto3,enum=proto.Diagnostic_Severity" json:"severity,omitempty"`
	Summary   string              `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	Detail    string              `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	FilePaths []string            `protobuf:"bytes,4,rep,name=FilePaths,proto3" json:"FilePaths,omitempty"`
}

func (x *Diagnostic) Reset() {
	*x = Diagnostic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Diagnostic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Diagnostic) ProtoMessage() {}

func (x *Diagnostic) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Diagnostic.ProtoReflect.Descriptor instead.
func (*Diagnostic) Descriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *Diagnostic) GetSeverity() Diagnostic_Severity {
	if x != nil {
		return x.Severity
	}
	return Diagnostic_INVALID
}

func (x *Diagnostic) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Diagnostic) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *Diagnostic) GetFilePaths() []string {
	if x != nil {
		return x.FilePaths
	}
	return nil
}

// PluginRequest ...
type PluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FilePaths ...
	FilePaths []string `protobuf:"bytes,1,rep,name=FilePaths,proto3" json:"FilePaths,omitempty"`
	// Parameters ...
	Parameters map[string]string `protobuf:"bytes,2,rep,name=Parameters,proto3" json:"Parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Spec ...
	Spec string `protobuf:"bytes,3,opt,name=Spec,proto3" json:"Spec,omitempty"`
	// Version ...
	Version string `protobuf:"bytes,4,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *PluginRequest) Reset() {
	*x = PluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginRequest) ProtoMessage() {}

func (x *PluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginRequest.ProtoReflect.Descriptor instead.
func (*PluginRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *PluginRequest) GetFilePaths() []string {
	if x != nil {
		return x.FilePaths
	}
	return nil
}

func (x *PluginRequest) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *PluginRequest) GetSpec() string {
	if x != nil {
		return x.Spec
	}
	return ""
}

func (x *PluginRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// PluginResponse ...
type PluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FilePaths ...
	FilePaths []string `protobuf:"bytes,1,rep,name=FilePaths,proto3" json:"FilePaths,omitempty"`
	// Error ...
	Error string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *PluginResponse) Reset() {
	*x = PluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginResponse) ProtoMessage() {}

func (x *PluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginResponse.ProtoReflect.Descriptor instead.
func (*PluginResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *PluginResponse) GetFilePaths() []string {
	if x != nil {
		return x.FilePaths
	}
	return nil
}

func (x *PluginResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// PluginCommandRequest ...
type PluginCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Command ...
	Command string `protobuf:"bytes,1,opt,name=Command,proto3" json:"Command,omitempty"`
	// Parameters ...
	Parameters map[string]string `protobuf:"bytes,2,rep,name=Parameters,proto3" json:"Parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Spec ...
	Spec string `protobuf:"bytes,3,opt,name=Spec,proto3" json:"Spec,omitempty"`
	// Version ...
	Version string `protobuf:"bytes,4,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *PluginCommandRequest) Reset() {
	*x = PluginCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_plugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginCommandRequest) ProtoMessage() {}

func (x *PluginCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_plugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginCommandRequest.ProtoReflect.Descriptor instead.
func (*PluginCommandRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{5}
}

func (x *PluginCommandRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *PluginCommandRequest) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *PluginCommandRequest) GetSpec() string {
	if x != nil {
		return x.Spec
	}
	return ""
}

func (x *PluginCommandRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// PluginCommandResponse ...
type PluginCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Error ...
	Error string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *PluginCommandResponse) Reset() {
	*x = PluginCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_plugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginCommandResponse) ProtoMessage() {}

func (x *PluginCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_plugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginCommandResponse.ProtoReflect.Descriptor instead.
func (*PluginCommandResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{6}
}

func (x *PluginCommandResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// PluginUsageRequest ...
type PluginUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Spec ...
	Spec string `protobuf:"bytes,1,opt,name=Spec,proto3" json:"Spec,omitempty"`
	// Version ...
	Version string `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *PluginUsageRequest) Reset() {
	*x = PluginUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_plugin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginUsageRequest) ProtoMessage() {}

func (x *PluginUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_plugin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginUsageRequest.ProtoReflect.Descriptor instead.
func (*PluginUsageRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{7}
}

func (x *PluginUsageRequest) GetSpec() string {
	if x != nil {
		return x.Spec
	}
	return ""
}

func (x *PluginUsageRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// PluginUsageResponse ...
type PluginUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Error ...
	Error string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	// Parameters ...
	Parameters map[string]string `protobuf:"bytes,2,rep,name=Parameters,proto3" json:"Parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PluginUsageResponse) Reset() {
	*x = PluginUsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_plugin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginUsageResponse) ProtoMessage() {}

func (x *PluginUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_plugin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginUsageResponse.ProtoReflect.Descriptor instead.
func (*PluginUsageResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{8}
}

func (x *PluginUsageResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *PluginUsageResponse) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

// Request ...
type Execute_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string            `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Vars    map[string]string `protobuf:"bytes,2,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Params  []string          `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *Execute_Request) Reset() {
	*x = Execute_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_plugin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Execute_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Execute_Request) ProtoMessage() {}

func (x *Execute_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_plugin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Execute_Request.ProtoReflect.Descriptor instead.
func (*Execute_Request) Descriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Execute_Request) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Execute_Request) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}

func (x *Execute_Request) GetParams() []string {
	if x != nil {
		return x.Params
	}
	return nil
}

// Response ...
type Execute_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     Execute_Status `protobuf:"varint,1,opt,name=status,proto3,enum=proto.Execute_Status" json:"status,omitempty"`
	Diagnostic []*Diagnostic  `protobuf:"bytes,10,rep,name=diagnostic,proto3" json:"diagnostic,omitempty"`
}

func (x *Execute_Response) Reset() {
	*x = Execute_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_plugin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Execute_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Execute_Response) ProtoMessage() {}

func (x *Execute_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_plugin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Execute_Response.ProtoReflect.Descriptor instead.
func (*Execute_Response) Descriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Execute_Response) GetStatus() Execute_Status {
	if x != nil {
		return x.Status
	}
	return Execute_UNKNOWN
}

func (x *Execute_Response) GetDiagnostic() []*Diagnostic {
	if x != nil {
		return x.Diagnostic
	}
	return nil
}

// Request ...
type Stop_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Stop_Request) Reset() {
	*x = Stop_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_plugin_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stop_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stop_Request) ProtoMessage() {}

func (x *Stop_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_plugin_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stop_Request.ProtoReflect.Descriptor instead.
func (*Stop_Request) Descriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{1, 0}
}

// Response ...
type Stop_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *Stop_Response) Reset() {
	*x = Stop_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_plugin_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stop_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stop_Response) ProtoMessage() {}

func (x *Stop_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_plugin_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stop_Response.ProtoReflect.Descriptor instead.
func (*Stop_Response) Descriptor() ([]byte, []int) {
	return file_pkg_proto_plugin_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Stop_Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_proto_plugin_proto protoreflect.FileDescriptor

var file_pkg_proto_plugin_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd5, 0x02, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x1a, 0xaa, 0x01, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x34, 0x0a, 0x04, 0x76, 0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x76, 0x61, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x37, 0x0a, 0x09, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x6c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69,
	0x63, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x52, 0x0a, 0x64, 0x69, 0x61, 0x67,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x22, 0x2f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41,
	0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x02, 0x22, 0x33, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x1a,
	0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xc5, 0x01, 0x0a,
	0x0a, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x12, 0x36, 0x0a, 0x08, 0x73,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63,
	0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x73, 0x22, 0x2f, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x22, 0xe0, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x73, 0x12, 0x44, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x70,
	0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x18,
	0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x44, 0x0a, 0x0e, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xea, 0x01,
	0x0a, 0x14, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x4b, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x53, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x3d, 0x0a, 0x0f, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2d, 0x0a, 0x15, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x42, 0x0a, 0x12, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xb6, 0x01,
	0x0a, 0x13, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x4a, 0x0a, 0x0a, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x7b, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x12, 0x3c, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x74, 0x6f, 0x70, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x2f, 0x70,
	0x69, 0x63, 0x61, 0x73, 0x73, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_plugin_proto_rawDescOnce sync.Once
	file_pkg_proto_plugin_proto_rawDescData = file_pkg_proto_plugin_proto_rawDesc
)

func file_pkg_proto_plugin_proto_rawDescGZIP() []byte {
	file_pkg_proto_plugin_proto_rawDescOnce.Do(func() {
		file_pkg_proto_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_plugin_proto_rawDescData)
	})
	return file_pkg_proto_plugin_proto_rawDescData
}

var file_pkg_proto_plugin_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_proto_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_pkg_proto_plugin_proto_goTypes = []interface{}{
	(Execute_Status)(0),           // 0: proto.Execute.Status
	(Diagnostic_Severity)(0),      // 1: proto.Diagnostic.Severity
	(*Execute)(nil),               // 2: proto.Execute
	(*Stop)(nil),                  // 3: proto.Stop
	(*Diagnostic)(nil),            // 4: proto.Diagnostic
	(*PluginRequest)(nil),         // 5: proto.PluginRequest
	(*PluginResponse)(nil),        // 6: proto.PluginResponse
	(*PluginCommandRequest)(nil),  // 7: proto.PluginCommandRequest
	(*PluginCommandResponse)(nil), // 8: proto.PluginCommandResponse
	(*PluginUsageRequest)(nil),    // 9: proto.PluginUsageRequest
	(*PluginUsageResponse)(nil),   // 10: proto.PluginUsageResponse
	(*Execute_Request)(nil),       // 11: proto.Execute.Request
	(*Execute_Response)(nil),      // 12: proto.Execute.Response
	nil,                           // 13: proto.Execute.Request.VarsEntry
	(*Stop_Request)(nil),          // 14: proto.Stop.Request
	(*Stop_Response)(nil),         // 15: proto.Stop.Response
	nil,                           // 16: proto.PluginRequest.ParametersEntry
	nil,                           // 17: proto.PluginCommandRequest.ParametersEntry
	nil,                           // 18: proto.PluginUsageResponse.ParametersEntry
}
var file_pkg_proto_plugin_proto_depIdxs = []int32{
	1,  // 0: proto.Diagnostic.severity:type_name -> proto.Diagnostic.Severity
	16, // 1: proto.PluginRequest.Parameters:type_name -> proto.PluginRequest.ParametersEntry
	17, // 2: proto.PluginCommandRequest.Parameters:type_name -> proto.PluginCommandRequest.ParametersEntry
	18, // 3: proto.PluginUsageResponse.Parameters:type_name -> proto.PluginUsageResponse.ParametersEntry
	13, // 4: proto.Execute.Request.vars:type_name -> proto.Execute.Request.VarsEntry
	0,  // 5: proto.Execute.Response.status:type_name -> proto.Execute.Status
	4,  // 6: proto.Execute.Response.diagnostic:type_name -> proto.Diagnostic
	11, // 7: proto.Plugin.Execute:input_type -> proto.Execute.Request
	14, // 8: proto.Plugin.Stop:input_type -> proto.Stop.Request
	12, // 9: proto.Plugin.Execute:output_type -> proto.Execute.Response
	15, // 10: proto.Plugin.Stop:output_type -> proto.Stop.Response
	9,  // [9:11] is the sub-list for method output_type
	7,  // [7:9] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_proto_plugin_proto_init() }
func file_pkg_proto_plugin_proto_init() {
	if File_pkg_proto_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Execute); i {
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
		file_pkg_proto_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stop); i {
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
		file_pkg_proto_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Diagnostic); i {
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
		file_pkg_proto_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginRequest); i {
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
		file_pkg_proto_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginResponse); i {
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
		file_pkg_proto_plugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginCommandRequest); i {
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
		file_pkg_proto_plugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginCommandResponse); i {
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
		file_pkg_proto_plugin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginUsageRequest); i {
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
		file_pkg_proto_plugin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginUsageResponse); i {
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
		file_pkg_proto_plugin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Execute_Request); i {
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
		file_pkg_proto_plugin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Execute_Response); i {
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
		file_pkg_proto_plugin_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stop_Request); i {
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
		file_pkg_proto_plugin_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stop_Response); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_plugin_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_plugin_proto_goTypes,
		DependencyIndexes: file_pkg_proto_plugin_proto_depIdxs,
		EnumInfos:         file_pkg_proto_plugin_proto_enumTypes,
		MessageInfos:      file_pkg_proto_plugin_proto_msgTypes,
	}.Build()
	File_pkg_proto_plugin_proto = out.File
	file_pkg_proto_plugin_proto_rawDesc = nil
	file_pkg_proto_plugin_proto_goTypes = nil
	file_pkg_proto_plugin_proto_depIdxs = nil
}
