// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protoService/remiderService.proto

package protoService

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TaskRequest struct {
	Task                 *Task    `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskRequest) Reset()         { *m = TaskRequest{} }
func (m *TaskRequest) String() string { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()    {}
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d3f97059af3282a, []int{0}
}

func (m *TaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRequest.Unmarshal(m, b)
}
func (m *TaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRequest.Marshal(b, m, deterministic)
}
func (m *TaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRequest.Merge(m, src)
}
func (m *TaskRequest) XXX_Size() int {
	return xxx_messageInfo_TaskRequest.Size(m)
}
func (m *TaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRequest proto.InternalMessageInfo

func (m *TaskRequest) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type TaskResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskResponse) Reset()         { *m = TaskResponse{} }
func (m *TaskResponse) String() string { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()    {}
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d3f97059af3282a, []int{1}
}

func (m *TaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResponse.Unmarshal(m, b)
}
func (m *TaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResponse.Marshal(b, m, deterministic)
}
func (m *TaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResponse.Merge(m, src)
}
func (m *TaskResponse) XXX_Size() int {
	return xxx_messageInfo_TaskResponse.Size(m)
}
func (m *TaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResponse proto.InternalMessageInfo

func (m *TaskResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetTasksRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTasksRequest) Reset()         { *m = GetTasksRequest{} }
func (m *GetTasksRequest) String() string { return proto.CompactTextString(m) }
func (*GetTasksRequest) ProtoMessage()    {}
func (*GetTasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d3f97059af3282a, []int{2}
}

func (m *GetTasksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTasksRequest.Unmarshal(m, b)
}
func (m *GetTasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTasksRequest.Marshal(b, m, deterministic)
}
func (m *GetTasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTasksRequest.Merge(m, src)
}
func (m *GetTasksRequest) XXX_Size() int {
	return xxx_messageInfo_GetTasksRequest.Size(m)
}
func (m *GetTasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTasksRequest proto.InternalMessageInfo

func (m *GetTasksRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TasksResponse struct {
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TasksResponse) Reset()         { *m = TasksResponse{} }
func (m *TasksResponse) String() string { return proto.CompactTextString(m) }
func (*TasksResponse) ProtoMessage()    {}
func (*TasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d3f97059af3282a, []int{3}
}

func (m *TasksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TasksResponse.Unmarshal(m, b)
}
func (m *TasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TasksResponse.Marshal(b, m, deterministic)
}
func (m *TasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TasksResponse.Merge(m, src)
}
func (m *TasksResponse) XXX_Size() int {
	return xxx_messageInfo_TasksResponse.Size(m)
}
func (m *TasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TasksResponse proto.InternalMessageInfo

func (m *TasksResponse) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func init() {
	proto.RegisterType((*TaskRequest)(nil), "remiderService.TaskRequest")
	proto.RegisterType((*TaskResponse)(nil), "remiderService.TaskResponse")
	proto.RegisterType((*GetTasksRequest)(nil), "remiderService.GetTasksRequest")
	proto.RegisterType((*TasksResponse)(nil), "remiderService.TasksResponse")
}

func init() {
	proto.RegisterFile("protoService/remiderService.proto", fileDescriptor_2d3f97059af3282a)
}

var fileDescriptor_2d3f97059af3282a = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x0f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0x4a, 0xcd, 0xcd, 0x4c, 0x49, 0x2d,
	0x82, 0x72, 0xf5, 0xc0, 0x72, 0x42, 0x7c, 0xa8, 0xa2, 0x52, 0xe2, 0x28, 0x5a, 0x4a, 0x12, 0x8b,
	0xb3, 0x21, 0x0a, 0x95, 0xcc, 0xb9, 0xb8, 0x43, 0x12, 0x8b, 0xb3, 0x83, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0x84, 0x34, 0xb8, 0x58, 0x40, 0x92, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x22,
	0x7a, 0x68, 0x86, 0x83, 0x95, 0x82, 0x55, 0x28, 0x69, 0x70, 0xf1, 0x40, 0x34, 0x16, 0x17, 0xe4,
	0xe7, 0x15, 0xa7, 0x0a, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x82, 0x35,
	0x73, 0x06, 0xc1, 0xb8, 0x4a, 0xda, 0x5c, 0xfc, 0xee, 0xa9, 0x25, 0x20, 0xc5, 0xc5, 0x30, 0x6b,
	0x70, 0x2b, 0xb6, 0xe6, 0xe2, 0x85, 0xaa, 0x84, 0x9a, 0xab, 0xc5, 0xc5, 0x0a, 0xb2, 0xaf, 0x58,
	0x82, 0x51, 0x81, 0x19, 0xa7, 0x93, 0x20, 0x4a, 0x8c, 0x96, 0x32, 0x72, 0xf1, 0x05, 0xa1, 0x48,
	0x0b, 0xb9, 0x73, 0x71, 0x39, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0x82, 0xd4, 0x09, 0x49, 0x63, 0xd5,
	0x0d, 0x71, 0x94, 0x94, 0x0c, 0x76, 0x49, 0xa8, 0x3b, 0xbc, 0xb8, 0x38, 0x60, 0xbe, 0x10, 0x92,
	0x47, 0x57, 0x89, 0xe6, 0x3f, 0x29, 0x59, 0x6c, 0x46, 0xc1, 0xfd, 0xe4, 0xc4, 0x17, 0xc5, 0x83,
	0x1c, 0x1f, 0x49, 0x6c, 0x60, 0x9e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x66, 0x86, 0x97, 0x02,
	0xd9, 0x01, 0x00, 0x00,
}