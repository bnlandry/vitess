// Code generated by protoc-gen-go.
// source: automation.proto
// DO NOT EDIT!

/*
Package automation is a generated protocol buffer package.

It is generated from these files:
	automation.proto

It has these top-level messages:
	ClusterOperation
	TaskContainer
	Task
	EnqueueClusterOperationRequest
	EnqueueClusterOperationResponse
	GetClusterOperationStateRequest
	GetClusterOperationStateResponse
	GetClusterOperationDetailsRequest
	GetClusterOperationDetailsResponse
*/
package automation

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ClusterOperationState int32

const (
	ClusterOperationState_UNKNOWN_CLUSTER_OPERATION_STATE ClusterOperationState = 0
	ClusterOperationState_CLUSTER_OPERATION_NOT_STARTED   ClusterOperationState = 1
	ClusterOperationState_CLUSTER_OPERATION_RUNNING       ClusterOperationState = 2
	ClusterOperationState_CLUSTER_OPERATION_DONE          ClusterOperationState = 3
)

var ClusterOperationState_name = map[int32]string{
	0: "UNKNOWN_CLUSTER_OPERATION_STATE",
	1: "CLUSTER_OPERATION_NOT_STARTED",
	2: "CLUSTER_OPERATION_RUNNING",
	3: "CLUSTER_OPERATION_DONE",
}
var ClusterOperationState_value = map[string]int32{
	"UNKNOWN_CLUSTER_OPERATION_STATE": 0,
	"CLUSTER_OPERATION_NOT_STARTED":   1,
	"CLUSTER_OPERATION_RUNNING":       2,
	"CLUSTER_OPERATION_DONE":          3,
}

func (x ClusterOperationState) String() string {
	return proto.EnumName(ClusterOperationState_name, int32(x))
}
func (ClusterOperationState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TaskState int32

const (
	TaskState_UNKNOWN_TASK_STATE TaskState = 0
	TaskState_NOT_STARTED        TaskState = 1
	TaskState_RUNNING            TaskState = 2
	TaskState_DONE               TaskState = 3
)

var TaskState_name = map[int32]string{
	0: "UNKNOWN_TASK_STATE",
	1: "NOT_STARTED",
	2: "RUNNING",
	3: "DONE",
}
var TaskState_value = map[string]int32{
	"UNKNOWN_TASK_STATE": 0,
	"NOT_STARTED":        1,
	"RUNNING":            2,
	"DONE":               3,
}

func (x TaskState) String() string {
	return proto.EnumName(TaskState_name, int32(x))
}
func (TaskState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ClusterOperation struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// TaskContainer are processed sequentially, one at a time.
	SerialTasks []*TaskContainer `protobuf:"bytes,2,rep,name=serial_tasks" json:"serial_tasks,omitempty"`
	// Cached value. This has to be re-evaluated e.g. after a checkpoint load because running tasks may have already finished.
	State ClusterOperationState `protobuf:"varint,3,opt,name=state,enum=automation.ClusterOperationState" json:"state,omitempty"`
	// Error of the first task which failed. Set after state advanced to CLUSTER_OPERATION_DONE. If empty, all tasks succeeded. Cached value, see state above.
	Error string `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *ClusterOperation) Reset()                    { *m = ClusterOperation{} }
func (m *ClusterOperation) String() string            { return proto.CompactTextString(m) }
func (*ClusterOperation) ProtoMessage()               {}
func (*ClusterOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClusterOperation) GetSerialTasks() []*TaskContainer {
	if m != nil {
		return m.SerialTasks
	}
	return nil
}

// TaskContainer holds one or more task which may be executed in parallel.
// "concurrency", if > 0, limits the amount of concurrently executed tasks.
type TaskContainer struct {
	ParallelTasks []*Task `protobuf:"bytes,1,rep,name=parallel_tasks" json:"parallel_tasks,omitempty"`
	Concurrency   int32   `protobuf:"varint,2,opt,name=concurrency" json:"concurrency,omitempty"`
}

func (m *TaskContainer) Reset()                    { *m = TaskContainer{} }
func (m *TaskContainer) String() string            { return proto.CompactTextString(m) }
func (*TaskContainer) ProtoMessage()               {}
func (*TaskContainer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TaskContainer) GetParallelTasks() []*Task {
	if m != nil {
		return m.ParallelTasks
	}
	return nil
}

// Task represents a specific task which should be automatically executed.
type Task struct {
	// Task specification.
	Name       string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parameters map[string]string `protobuf:"bytes,2,rep,name=parameters" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Runtime data.
	Id    string    `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	State TaskState `protobuf:"varint,4,opt,name=state,enum=automation.TaskState" json:"state,omitempty"`
	// Set after state advanced to DONE.
	Output string `protobuf:"bytes,5,opt,name=output" json:"output,omitempty"`
	// Set after state advanced to DONE. If empty, the task did succeed.
	Error string `protobuf:"bytes,6,opt,name=error" json:"error,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Task) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type EnqueueClusterOperationRequest struct {
	Name       string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parameters map[string]string `protobuf:"bytes,2,rep,name=parameters" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EnqueueClusterOperationRequest) Reset()                    { *m = EnqueueClusterOperationRequest{} }
func (m *EnqueueClusterOperationRequest) String() string            { return proto.CompactTextString(m) }
func (*EnqueueClusterOperationRequest) ProtoMessage()               {}
func (*EnqueueClusterOperationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EnqueueClusterOperationRequest) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type EnqueueClusterOperationResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *EnqueueClusterOperationResponse) Reset()                    { *m = EnqueueClusterOperationResponse{} }
func (m *EnqueueClusterOperationResponse) String() string            { return proto.CompactTextString(m) }
func (*EnqueueClusterOperationResponse) ProtoMessage()               {}
func (*EnqueueClusterOperationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type GetClusterOperationStateRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetClusterOperationStateRequest) Reset()                    { *m = GetClusterOperationStateRequest{} }
func (m *GetClusterOperationStateRequest) String() string            { return proto.CompactTextString(m) }
func (*GetClusterOperationStateRequest) ProtoMessage()               {}
func (*GetClusterOperationStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type GetClusterOperationStateResponse struct {
	State ClusterOperationState `protobuf:"varint,1,opt,name=state,enum=automation.ClusterOperationState" json:"state,omitempty"`
}

func (m *GetClusterOperationStateResponse) Reset()         { *m = GetClusterOperationStateResponse{} }
func (m *GetClusterOperationStateResponse) String() string { return proto.CompactTextString(m) }
func (*GetClusterOperationStateResponse) ProtoMessage()    {}
func (*GetClusterOperationStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6}
}

type GetClusterOperationDetailsRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetClusterOperationDetailsRequest) Reset()         { *m = GetClusterOperationDetailsRequest{} }
func (m *GetClusterOperationDetailsRequest) String() string { return proto.CompactTextString(m) }
func (*GetClusterOperationDetailsRequest) ProtoMessage()    {}
func (*GetClusterOperationDetailsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7}
}

type GetClusterOperationDetailsResponse struct {
	// Full snapshot of the execution e.g. including output of each task.
	ClusterOp *ClusterOperation `protobuf:"bytes,2,opt,name=cluster_op" json:"cluster_op,omitempty"`
}

func (m *GetClusterOperationDetailsResponse) Reset()         { *m = GetClusterOperationDetailsResponse{} }
func (m *GetClusterOperationDetailsResponse) String() string { return proto.CompactTextString(m) }
func (*GetClusterOperationDetailsResponse) ProtoMessage()    {}
func (*GetClusterOperationDetailsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8}
}

func (m *GetClusterOperationDetailsResponse) GetClusterOp() *ClusterOperation {
	if m != nil {
		return m.ClusterOp
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterOperation)(nil), "automation.ClusterOperation")
	proto.RegisterType((*TaskContainer)(nil), "automation.TaskContainer")
	proto.RegisterType((*Task)(nil), "automation.Task")
	proto.RegisterType((*EnqueueClusterOperationRequest)(nil), "automation.EnqueueClusterOperationRequest")
	proto.RegisterType((*EnqueueClusterOperationResponse)(nil), "automation.EnqueueClusterOperationResponse")
	proto.RegisterType((*GetClusterOperationStateRequest)(nil), "automation.GetClusterOperationStateRequest")
	proto.RegisterType((*GetClusterOperationStateResponse)(nil), "automation.GetClusterOperationStateResponse")
	proto.RegisterType((*GetClusterOperationDetailsRequest)(nil), "automation.GetClusterOperationDetailsRequest")
	proto.RegisterType((*GetClusterOperationDetailsResponse)(nil), "automation.GetClusterOperationDetailsResponse")
	proto.RegisterEnum("automation.ClusterOperationState", ClusterOperationState_name, ClusterOperationState_value)
	proto.RegisterEnum("automation.TaskState", TaskState_name, TaskState_value)
}

var fileDescriptor0 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xd9, 0xc4, 0x09, 0x64, 0xdc, 0xa6, 0xd6, 0xa2, 0x56, 0x6e, 0x45, 0x69, 0x62, 0x38,
	0x44, 0x95, 0x48, 0x21, 0x70, 0x40, 0xbd, 0x45, 0x89, 0x55, 0x55, 0x41, 0xeb, 0xca, 0x71, 0xe0,
	0x68, 0x2d, 0x61, 0x0f, 0x51, 0x1d, 0xaf, 0x59, 0xaf, 0x91, 0xf2, 0x16, 0x9c, 0x79, 0x0b, 0x5e,
	0x84, 0x67, 0xc2, 0x76, 0x62, 0xc7, 0x75, 0xdc, 0x08, 0x71, 0xcc, 0xce, 0x3f, 0xff, 0x7c, 0xff,
	0x68, 0x62, 0xd0, 0x68, 0x24, 0xf9, 0x92, 0xca, 0x05, 0xf7, 0xfb, 0x81, 0xe0, 0x92, 0x63, 0xd8,
	0xbe, 0x18, 0x3f, 0x11, 0x68, 0x23, 0x2f, 0x0a, 0x25, 0x13, 0x56, 0xc0, 0x44, 0xfa, 0x88, 0x01,
	0x6a, 0x8b, 0x6f, 0x3a, 0xea, 0xa0, 0x5e, 0x0b, 0x5f, 0xc1, 0x41, 0xc8, 0xc4, 0x82, 0x7a, 0xae,
	0xa4, 0xe1, 0x7d, 0xa8, 0xd7, 0x3a, 0xf5, 0x9e, 0x3a, 0x38, 0xed, 0x17, 0x5c, 0x9d, 0xb8, 0x30,
	0xe2, 0xbe, 0xa4, 0x0b, 0x9f, 0x09, 0xfc, 0x16, 0x1a, 0xa1, 0xa4, 0x92, 0xe9, 0xf5, 0xb8, 0xbf,
	0x3d, 0xe8, 0x16, 0x95, 0xe5, 0x49, 0xd3, 0x44, 0x88, 0x0f, 0xa1, 0xc1, 0x84, 0xe0, 0x42, 0x57,
	0x92, 0x89, 0x06, 0x81, 0xc3, 0x87, 0x8e, 0x3d, 0x68, 0x07, 0x54, 0x50, 0xcf, 0x63, 0x19, 0x04,
	0x4a, 0x21, 0xb4, 0x32, 0x04, 0x7e, 0x0e, 0xea, 0x9c, 0xfb, 0xf3, 0x48, 0x08, 0xe6, 0xcf, 0x57,
	0x31, 0x2b, 0xea, 0x35, 0x8c, 0x3f, 0x08, 0x94, 0xb4, 0x7a, 0x00, 0x8a, 0x4f, 0x97, 0x6c, 0x13,
	0xec, 0x03, 0x40, 0xe2, 0xba, 0x64, 0x31, 0x50, 0x16, 0xab, 0x53, 0x76, 0xec, 0xdf, 0xe5, 0x12,
	0xd3, 0x97, 0x62, 0xb5, 0x59, 0x4d, 0x3d, 0x75, 0x78, 0x9d, 0x25, 0x55, 0xd2, 0xa4, 0xc7, 0xe5,
	0xe6, 0x75, 0xba, 0x36, 0x34, 0x79, 0x24, 0x83, 0x48, 0xea, 0x8d, 0xb4, 0x2b, 0x4f, 0xdb, 0x4c,
	0x7e, 0x9e, 0xbd, 0x83, 0xa3, 0xf2, 0x0c, 0x15, 0xea, 0xf7, 0x6c, 0xb5, 0xc1, 0x8c, 0xe5, 0x3f,
	0xa8, 0x17, 0xb1, 0x34, 0x4c, 0xeb, 0xba, 0xf6, 0x11, 0x19, 0xbf, 0x11, 0xbc, 0x34, 0xfd, 0xef,
	0x11, 0x8b, 0x58, 0x79, 0xa1, 0x36, 0x8b, 0x9f, 0x43, 0x59, 0x8a, 0x4a, 0x2a, 0xa2, 0x5e, 0x17,
	0x69, 0xf7, 0xbb, 0x95, 0x97, 0xf0, 0x3f, 0xcc, 0x6f, 0xe0, 0xe2, 0xd1, 0x21, 0x61, 0xc0, 0xfd,
	0x90, 0x15, 0xaf, 0x2e, 0x91, 0xdf, 0x30, 0x59, 0x79, 0x2e, 0x59, 0xc4, 0xa2, 0xdc, 0x81, 0xce,
	0xe3, 0xf2, 0x8d, 0x7d, 0x7e, 0x97, 0xe8, 0x1f, 0xef, 0xd2, 0xb8, 0x82, 0x6e, 0x85, 0xeb, 0x98,
	0xc5, 0x67, 0xe9, 0x85, 0x55, 0x18, 0x9f, 0xc1, 0xd8, 0xd7, 0x90, 0x83, 0xc0, 0x7c, 0x2d, 0x71,
	0x79, 0x90, 0xae, 0x48, 0x1d, 0xbc, 0xd8, 0x47, 0x73, 0xf9, 0x0b, 0xc1, 0x71, 0xf5, 0x5f, 0xe7,
	0x15, 0x5c, 0xcc, 0xc8, 0x84, 0x58, 0x5f, 0x88, 0x3b, 0xfa, 0x34, 0x9b, 0x3a, 0xa6, 0xed, 0x5a,
	0x77, 0xa6, 0x3d, 0x74, 0x6e, 0x2d, 0xe2, 0x4e, 0x9d, 0xa1, 0x63, 0x6a, 0x4f, 0x70, 0x17, 0xce,
	0x77, 0x8b, 0xc4, 0x72, 0x12, 0x81, 0xed, 0x98, 0x63, 0x0d, 0xe1, 0x73, 0x38, 0xdd, 0x95, 0xd8,
	0x33, 0x42, 0x6e, 0xc9, 0x8d, 0x56, 0xc3, 0x67, 0x70, 0xb2, 0x5b, 0x1e, 0x5b, 0xc4, 0xd4, 0xea,
	0x97, 0x13, 0x68, 0x6d, 0x8f, 0xfd, 0x04, 0x70, 0xc6, 0xe3, 0x0c, 0xa7, 0x93, 0x1c, 0xe1, 0x08,
	0xd4, 0x87, 0x03, 0x55, 0x78, 0xba, 0xb5, 0x7f, 0x06, 0xca, 0xda, 0xec, 0x6b, 0x33, 0xfd, 0x42,
	0xbd, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x81, 0x18, 0xbf, 0x1a, 0xb5, 0x04, 0x00, 0x00,
}
