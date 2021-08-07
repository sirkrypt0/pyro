// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/agent/v1/agent.proto

package agent_v1

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

type ExecuteCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command     string            `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Environment map[string]string `protobuf:"bytes,2,rep,name=environment,proto3" json:"environment,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExecuteCommandRequest) Reset() {
	*x = ExecuteCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_agent_v1_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteCommandRequest) ProtoMessage() {}

func (x *ExecuteCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_agent_v1_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteCommandRequest.ProtoReflect.Descriptor instead.
func (*ExecuteCommandRequest) Descriptor() ([]byte, []int) {
	return file_api_agent_v1_agent_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteCommandRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *ExecuteCommandRequest) GetEnvironment() map[string]string {
	if x != nil {
		return x.Environment
	}
	return nil
}

type ExecuteCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stdout   *ExecuteIO `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr   *ExecuteIO `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	ExitCode int32      `protobuf:"varint,3,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *ExecuteCommandResponse) Reset() {
	*x = ExecuteCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_agent_v1_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteCommandResponse) ProtoMessage() {}

func (x *ExecuteCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_agent_v1_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteCommandResponse.ProtoReflect.Descriptor instead.
func (*ExecuteCommandResponse) Descriptor() ([]byte, []int) {
	return file_api_agent_v1_agent_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteCommandResponse) GetStdout() *ExecuteIO {
	if x != nil {
		return x.Stdout
	}
	return nil
}

func (x *ExecuteCommandResponse) GetStderr() *ExecuteIO {
	if x != nil {
		return x.Stderr
	}
	return nil
}

func (x *ExecuteCommandResponse) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

type ExecuteCommandStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prepare *ExecuteCommandStreamRequest_Prepare `protobuf:"bytes,1,opt,name=prepare,proto3" json:"prepare,omitempty"`
	Stdin   *ExecuteIO                           `protobuf:"bytes,2,opt,name=stdin,proto3" json:"stdin,omitempty"`
}

func (x *ExecuteCommandStreamRequest) Reset() {
	*x = ExecuteCommandStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_agent_v1_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteCommandStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteCommandStreamRequest) ProtoMessage() {}

func (x *ExecuteCommandStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_agent_v1_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteCommandStreamRequest.ProtoReflect.Descriptor instead.
func (*ExecuteCommandStreamRequest) Descriptor() ([]byte, []int) {
	return file_api_agent_v1_agent_proto_rawDescGZIP(), []int{2}
}

func (x *ExecuteCommandStreamRequest) GetPrepare() *ExecuteCommandStreamRequest_Prepare {
	if x != nil {
		return x.Prepare
	}
	return nil
}

func (x *ExecuteCommandStreamRequest) GetStdin() *ExecuteIO {
	if x != nil {
		return x.Stdin
	}
	return nil
}

type ExecuteCommandStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stdout *ExecuteIO     `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr *ExecuteIO     `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	Result *ExecuteResult `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ExecuteCommandStreamResponse) Reset() {
	*x = ExecuteCommandStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_agent_v1_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteCommandStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteCommandStreamResponse) ProtoMessage() {}

func (x *ExecuteCommandStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_agent_v1_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteCommandStreamResponse.ProtoReflect.Descriptor instead.
func (*ExecuteCommandStreamResponse) Descriptor() ([]byte, []int) {
	return file_api_agent_v1_agent_proto_rawDescGZIP(), []int{3}
}

func (x *ExecuteCommandStreamResponse) GetStdout() *ExecuteIO {
	if x != nil {
		return x.Stdout
	}
	return nil
}

func (x *ExecuteCommandStreamResponse) GetStderr() *ExecuteIO {
	if x != nil {
		return x.Stderr
	}
	return nil
}

func (x *ExecuteCommandStreamResponse) GetResult() *ExecuteResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type ExecuteResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exited   bool  `protobuf:"varint,1,opt,name=exited,proto3" json:"exited,omitempty"`
	ExitCode int32 `protobuf:"varint,2,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *ExecuteResult) Reset() {
	*x = ExecuteResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_agent_v1_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteResult) ProtoMessage() {}

func (x *ExecuteResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_agent_v1_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteResult.ProtoReflect.Descriptor instead.
func (*ExecuteResult) Descriptor() ([]byte, []int) {
	return file_api_agent_v1_agent_proto_rawDescGZIP(), []int{4}
}

func (x *ExecuteResult) GetExited() bool {
	if x != nil {
		return x.Exited
	}
	return false
}

func (x *ExecuteResult) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

type ExecuteIO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Close bool   `protobuf:"varint,1,opt,name=close,proto3" json:"close,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ExecuteIO) Reset() {
	*x = ExecuteIO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_agent_v1_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteIO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteIO) ProtoMessage() {}

func (x *ExecuteIO) ProtoReflect() protoreflect.Message {
	mi := &file_api_agent_v1_agent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteIO.ProtoReflect.Descriptor instead.
func (*ExecuteIO) Descriptor() ([]byte, []int) {
	return file_api_agent_v1_agent_proto_rawDescGZIP(), []int{5}
}

func (x *ExecuteIO) GetClose() bool {
	if x != nil {
		return x.Close
	}
	return false
}

func (x *ExecuteIO) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ExecuteCommandStreamRequest_Prepare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command     string            `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Environment map[string]string `protobuf:"bytes,2,rep,name=environment,proto3" json:"environment,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExecuteCommandStreamRequest_Prepare) Reset() {
	*x = ExecuteCommandStreamRequest_Prepare{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_agent_v1_agent_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteCommandStreamRequest_Prepare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteCommandStreamRequest_Prepare) ProtoMessage() {}

func (x *ExecuteCommandStreamRequest_Prepare) ProtoReflect() protoreflect.Message {
	mi := &file_api_agent_v1_agent_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteCommandStreamRequest_Prepare.ProtoReflect.Descriptor instead.
func (*ExecuteCommandStreamRequest_Prepare) Descriptor() ([]byte, []int) {
	return file_api_agent_v1_agent_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ExecuteCommandStreamRequest_Prepare) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *ExecuteCommandStreamRequest_Prepare) GetEnvironment() map[string]string {
	if x != nil {
		return x.Environment
	}
	return nil
}

var File_api_agent_v1_agent_proto protoreflect.FileDescriptor

var file_api_agent_v1_agent_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0xc9, 0x01, 0x0a, 0x15, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x56, 0x0a, 0x0b,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x3e, 0x0a, 0x10, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x97, 0x01, 0x0a, 0x16, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x49, 0x4f, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74,
	0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x49, 0x4f, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xe5,
	0x02, 0x0a, 0x1b, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b,
	0x0a, 0x07, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x52, 0x07, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x73,
	0x74, 0x64, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x49, 0x4f, 0x52, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x1a, 0xc9, 0x01, 0x0a, 0x07, 0x50,
	0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x64, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x3e, 0x0a, 0x10, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb5, 0x01, 0x0a, 0x1c, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x49, 0x4f,
	0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65,
	0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x49,
	0x4f, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x12, 0x33, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x44,
	0x0a, 0x0d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x65, 0x78, 0x69, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x35, 0x0a, 0x09, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x49,
	0x4f, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xde, 0x01, 0x0a, 0x0c,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0e,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x23,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x14, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x72, 0x6b, 0x72,
	0x79, 0x70, 0x74, 0x30, 0x2f, 0x70, 0x79, 0x72, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_agent_v1_agent_proto_rawDescOnce sync.Once
	file_api_agent_v1_agent_proto_rawDescData = file_api_agent_v1_agent_proto_rawDesc
)

func file_api_agent_v1_agent_proto_rawDescGZIP() []byte {
	file_api_agent_v1_agent_proto_rawDescOnce.Do(func() {
		file_api_agent_v1_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_agent_v1_agent_proto_rawDescData)
	})
	return file_api_agent_v1_agent_proto_rawDescData
}

var file_api_agent_v1_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_agent_v1_agent_proto_goTypes = []interface{}{
	(*ExecuteCommandRequest)(nil),               // 0: api.agent.v1.ExecuteCommandRequest
	(*ExecuteCommandResponse)(nil),              // 1: api.agent.v1.ExecuteCommandResponse
	(*ExecuteCommandStreamRequest)(nil),         // 2: api.agent.v1.ExecuteCommandStreamRequest
	(*ExecuteCommandStreamResponse)(nil),        // 3: api.agent.v1.ExecuteCommandStreamResponse
	(*ExecuteResult)(nil),                       // 4: api.agent.v1.ExecuteResult
	(*ExecuteIO)(nil),                           // 5: api.agent.v1.ExecuteIO
	nil,                                         // 6: api.agent.v1.ExecuteCommandRequest.EnvironmentEntry
	(*ExecuteCommandStreamRequest_Prepare)(nil), // 7: api.agent.v1.ExecuteCommandStreamRequest.Prepare
	nil, // 8: api.agent.v1.ExecuteCommandStreamRequest.Prepare.EnvironmentEntry
}
var file_api_agent_v1_agent_proto_depIdxs = []int32{
	6,  // 0: api.agent.v1.ExecuteCommandRequest.environment:type_name -> api.agent.v1.ExecuteCommandRequest.EnvironmentEntry
	5,  // 1: api.agent.v1.ExecuteCommandResponse.stdout:type_name -> api.agent.v1.ExecuteIO
	5,  // 2: api.agent.v1.ExecuteCommandResponse.stderr:type_name -> api.agent.v1.ExecuteIO
	7,  // 3: api.agent.v1.ExecuteCommandStreamRequest.prepare:type_name -> api.agent.v1.ExecuteCommandStreamRequest.Prepare
	5,  // 4: api.agent.v1.ExecuteCommandStreamRequest.stdin:type_name -> api.agent.v1.ExecuteIO
	5,  // 5: api.agent.v1.ExecuteCommandStreamResponse.stdout:type_name -> api.agent.v1.ExecuteIO
	5,  // 6: api.agent.v1.ExecuteCommandStreamResponse.stderr:type_name -> api.agent.v1.ExecuteIO
	4,  // 7: api.agent.v1.ExecuteCommandStreamResponse.result:type_name -> api.agent.v1.ExecuteResult
	8,  // 8: api.agent.v1.ExecuteCommandStreamRequest.Prepare.environment:type_name -> api.agent.v1.ExecuteCommandStreamRequest.Prepare.EnvironmentEntry
	0,  // 9: api.agent.v1.AgentService.ExecuteCommand:input_type -> api.agent.v1.ExecuteCommandRequest
	2,  // 10: api.agent.v1.AgentService.ExecuteCommandStream:input_type -> api.agent.v1.ExecuteCommandStreamRequest
	1,  // 11: api.agent.v1.AgentService.ExecuteCommand:output_type -> api.agent.v1.ExecuteCommandResponse
	3,  // 12: api.agent.v1.AgentService.ExecuteCommandStream:output_type -> api.agent.v1.ExecuteCommandStreamResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_agent_v1_agent_proto_init() }
func file_api_agent_v1_agent_proto_init() {
	if File_api_agent_v1_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_agent_v1_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteCommandRequest); i {
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
		file_api_agent_v1_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteCommandResponse); i {
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
		file_api_agent_v1_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteCommandStreamRequest); i {
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
		file_api_agent_v1_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteCommandStreamResponse); i {
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
		file_api_agent_v1_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteResult); i {
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
		file_api_agent_v1_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteIO); i {
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
		file_api_agent_v1_agent_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteCommandStreamRequest_Prepare); i {
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
			RawDescriptor: file_api_agent_v1_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_agent_v1_agent_proto_goTypes,
		DependencyIndexes: file_api_agent_v1_agent_proto_depIdxs,
		MessageInfos:      file_api_agent_v1_agent_proto_msgTypes,
	}.Build()
	File_api_agent_v1_agent_proto = out.File
	file_api_agent_v1_agent_proto_rawDesc = nil
	file_api_agent_v1_agent_proto_goTypes = nil
	file_api_agent_v1_agent_proto_depIdxs = nil
}