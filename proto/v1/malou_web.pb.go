// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: v1/malou_web.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// CreateRunner
type CreateRunnerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Labels []string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *CreateRunnerReq) Reset() {
	*x = CreateRunnerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_malou_web_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRunnerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRunnerReq) ProtoMessage() {}

func (x *CreateRunnerReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_malou_web_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRunnerReq.ProtoReflect.Descriptor instead.
func (*CreateRunnerReq) Descriptor() ([]byte, []int) {
	return file_v1_malou_web_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRunnerReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRunnerReq) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type CreateRunnerResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CreateRunnerResp) Reset() {
	*x = CreateRunnerResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_malou_web_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRunnerResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRunnerResp) ProtoMessage() {}

func (x *CreateRunnerResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_malou_web_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRunnerResp.ProtoReflect.Descriptor instead.
func (*CreateRunnerResp) Descriptor() ([]byte, []int) {
	return file_v1_malou_web_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRunnerResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// RunnerInfo
type RunnerInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunnerId string `protobuf:"bytes,1,opt,name=runner_id,json=runnerId,proto3" json:"runner_id,omitempty"`
}

func (x *RunnerInfoReq) Reset() {
	*x = RunnerInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_malou_web_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunnerInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunnerInfoReq) ProtoMessage() {}

func (x *RunnerInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_malou_web_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunnerInfoReq.ProtoReflect.Descriptor instead.
func (*RunnerInfoReq) Descriptor() ([]byte, []int) {
	return file_v1_malou_web_proto_rawDescGZIP(), []int{2}
}

func (x *RunnerInfoReq) GetRunnerId() string {
	if x != nil {
		return x.RunnerId
	}
	return ""
}

// RunnerList
type RunnerListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size   int64    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Page   int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Name   string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Labels []string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *RunnerListReq) Reset() {
	*x = RunnerListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_malou_web_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunnerListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunnerListReq) ProtoMessage() {}

func (x *RunnerListReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_malou_web_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunnerListReq.ProtoReflect.Descriptor instead.
func (*RunnerListReq) Descriptor() ([]byte, []int) {
	return file_v1_malou_web_proto_rawDescGZIP(), []int{3}
}

func (x *RunnerListReq) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *RunnerListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *RunnerListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RunnerListReq) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type RunnerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuPercent   []float64         `protobuf:"fixed64,2,rep,packed,name=cpu_percent,json=cpuPercent,proto3" json:"cpu_percent,omitempty"`
	MemoryInfo   *MemoryInfo       `protobuf:"bytes,3,opt,name=memory_info,json=memoryInfo,proto3" json:"memory_info,omitempty"`
	DiskInfo     *DiskInfo         `protobuf:"bytes,4,opt,name=disk_info,json=diskInfo,proto3" json:"disk_info,omitempty"`
	WorkerStatus map[string]string `protobuf:"bytes,5,rep,name=worker_status,json=workerStatus,proto3" json:"worker_status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RunnerInfo) Reset() {
	*x = RunnerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_malou_web_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunnerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunnerInfo) ProtoMessage() {}

func (x *RunnerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v1_malou_web_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunnerInfo.ProtoReflect.Descriptor instead.
func (*RunnerInfo) Descriptor() ([]byte, []int) {
	return file_v1_malou_web_proto_rawDescGZIP(), []int{4}
}

func (x *RunnerInfo) GetCpuPercent() []float64 {
	if x != nil {
		return x.CpuPercent
	}
	return nil
}

func (x *RunnerInfo) GetMemoryInfo() *MemoryInfo {
	if x != nil {
		return x.MemoryInfo
	}
	return nil
}

func (x *RunnerInfo) GetDiskInfo() *DiskInfo {
	if x != nil {
		return x.DiskInfo
	}
	return nil
}

func (x *RunnerInfo) GetWorkerStatus() map[string]string {
	if x != nil {
		return x.WorkerStatus
	}
	return nil
}

type RunnerListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*RunnerInfo `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *RunnerListResp) Reset() {
	*x = RunnerListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_malou_web_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunnerListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunnerListResp) ProtoMessage() {}

func (x *RunnerListResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_malou_web_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunnerListResp.ProtoReflect.Descriptor instead.
func (*RunnerListResp) Descriptor() ([]byte, []int) {
	return file_v1_malou_web_proto_rawDescGZIP(), []int{5}
}

func (x *RunnerListResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RunnerListResp) GetData() []*RunnerInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

// PipelineLogList
type PipelineLogListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipelineId string `protobuf:"bytes,1,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id,omitempty"`
	Offset     int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *PipelineLogListReq) Reset() {
	*x = PipelineLogListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_malou_web_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineLogListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineLogListReq) ProtoMessage() {}

func (x *PipelineLogListReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_malou_web_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineLogListReq.ProtoReflect.Descriptor instead.
func (*PipelineLogListReq) Descriptor() ([]byte, []int) {
	return file_v1_malou_web_proto_rawDescGZIP(), []int{6}
}

func (x *PipelineLogListReq) GetPipelineId() string {
	if x != nil {
		return x.PipelineId
	}
	return ""
}

func (x *PipelineLogListReq) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type PipelineLogListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log []*PipelineLog `protobuf:"bytes,1,rep,name=log,proto3" json:"log,omitempty"`
}

func (x *PipelineLogListResp) Reset() {
	*x = PipelineLogListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_malou_web_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineLogListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineLogListResp) ProtoMessage() {}

func (x *PipelineLogListResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_malou_web_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineLogListResp.ProtoReflect.Descriptor instead.
func (*PipelineLogListResp) Descriptor() ([]byte, []int) {
	return file_v1_malou_web_proto_rawDescGZIP(), []int{7}
}

func (x *PipelineLogListResp) GetLog() []*PipelineLog {
	if x != nil {
		return x.Log
	}
	return nil
}

var File_v1_malou_web_proto protoreflect.FileDescriptor

var file_v1_malou_web_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6c, 0x6f, 0x75, 0x5f, 0x77, 0x65, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x22, 0x28, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2c, 0x0a, 0x0d,
	0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x0d, 0x52, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22,
	0x88, 0x02, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x70, 0x75, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x70, 0x75, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12,
	0x2c, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a,
	0x09, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x64, 0x69, 0x73,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x52,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x47, 0x0a, 0x0e, 0x52, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x4d, 0x0a, 0x12, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c,
	0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x35, 0x0a, 0x13, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x6f,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x03, 0x6c, 0x6f, 0x67,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x32, 0x8e, 0x03, 0x0a, 0x08, 0x4d, 0x61,
	0x6c, 0x6f, 0x75, 0x57, 0x65, 0x62, 0x12, 0x51, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x52, 0x0a, 0x0a, 0x52, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d,
	0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x7b, 0x72, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x49, 0x0a,
	0x0a, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x52, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x52, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x8f, 0x01, 0x0a, 0x0f, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x14, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x51, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4b, 0x5a,
	0x20, 0x12, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x6c, 0x6f, 0x67, 0x2f, 0x7b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0x27, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x6c, 0x6f, 0x67, 0x2f, 0x7b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x7b, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x7d, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_malou_web_proto_rawDescOnce sync.Once
	file_v1_malou_web_proto_rawDescData = file_v1_malou_web_proto_rawDesc
)

func file_v1_malou_web_proto_rawDescGZIP() []byte {
	file_v1_malou_web_proto_rawDescOnce.Do(func() {
		file_v1_malou_web_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_malou_web_proto_rawDescData)
	})
	return file_v1_malou_web_proto_rawDescData
}

var file_v1_malou_web_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_v1_malou_web_proto_goTypes = []any{
	(*CreateRunnerReq)(nil),     // 0: CreateRunnerReq
	(*CreateRunnerResp)(nil),    // 1: CreateRunnerResp
	(*RunnerInfoReq)(nil),       // 2: RunnerInfoReq
	(*RunnerListReq)(nil),       // 3: RunnerListReq
	(*RunnerInfo)(nil),          // 4: RunnerInfo
	(*RunnerListResp)(nil),      // 5: RunnerListResp
	(*PipelineLogListReq)(nil),  // 6: PipelineLogListReq
	(*PipelineLogListResp)(nil), // 7: PipelineLogListResp
	nil,                         // 8: RunnerInfo.WorkerStatusEntry
	(*MemoryInfo)(nil),          // 9: MemoryInfo
	(*DiskInfo)(nil),            // 10: DiskInfo
	(*PipelineLog)(nil),         // 11: PipelineLog
}
var file_v1_malou_web_proto_depIdxs = []int32{
	9,  // 0: RunnerInfo.memory_info:type_name -> MemoryInfo
	10, // 1: RunnerInfo.disk_info:type_name -> DiskInfo
	8,  // 2: RunnerInfo.worker_status:type_name -> RunnerInfo.WorkerStatusEntry
	4,  // 3: RunnerListResp.data:type_name -> RunnerInfo
	11, // 4: PipelineLogListResp.log:type_name -> PipelineLog
	0,  // 5: MalouWeb.CreateRunner:input_type -> CreateRunnerReq
	2,  // 6: MalouWeb.RunnerInfo:input_type -> RunnerInfoReq
	3,  // 7: MalouWeb.RunnerList:input_type -> RunnerListReq
	6,  // 8: MalouWeb.PipelineLogList:input_type -> PipelineLogListReq
	1,  // 9: MalouWeb.CreateRunner:output_type -> CreateRunnerResp
	5,  // 10: MalouWeb.RunnerInfo:output_type -> RunnerListResp
	5,  // 11: MalouWeb.RunnerList:output_type -> RunnerListResp
	7,  // 12: MalouWeb.PipelineLogList:output_type -> PipelineLogListResp
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_v1_malou_web_proto_init() }
func file_v1_malou_web_proto_init() {
	if File_v1_malou_web_proto != nil {
		return
	}
	file_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_malou_web_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRunnerReq); i {
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
		file_v1_malou_web_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRunnerResp); i {
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
		file_v1_malou_web_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RunnerInfoReq); i {
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
		file_v1_malou_web_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RunnerListReq); i {
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
		file_v1_malou_web_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RunnerInfo); i {
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
		file_v1_malou_web_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RunnerListResp); i {
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
		file_v1_malou_web_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*PipelineLogListReq); i {
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
		file_v1_malou_web_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*PipelineLogListResp); i {
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
			RawDescriptor: file_v1_malou_web_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_malou_web_proto_goTypes,
		DependencyIndexes: file_v1_malou_web_proto_depIdxs,
		MessageInfos:      file_v1_malou_web_proto_msgTypes,
	}.Build()
	File_v1_malou_web_proto = out.File
	file_v1_malou_web_proto_rawDesc = nil
	file_v1_malou_web_proto_goTypes = nil
	file_v1_malou_web_proto_depIdxs = nil
}
