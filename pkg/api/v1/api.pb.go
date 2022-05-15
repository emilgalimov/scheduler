// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/v1/api.proto

package api

import (
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{1}
}

type TaskID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *TaskID) Reset() {
	*x = TaskID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskID) ProtoMessage() {}

func (x *TaskID) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskID.ProtoReflect.Descriptor instead.
func (*TaskID) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *TaskID) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type CreateTaskStageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID    uint64     `protobuf:"varint,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	TaskStage *TaskStage `protobuf:"bytes,2,opt,name=taskStage,proto3" json:"taskStage,omitempty"`
}

func (x *CreateTaskStageRequest) Reset() {
	*x = CreateTaskStageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskStageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskStageRequest) ProtoMessage() {}

func (x *CreateTaskStageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskStageRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskStageRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTaskStageRequest) GetTaskID() uint64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *CreateTaskStageRequest) GetTaskStage() *TaskStage {
	if x != nil {
		return x.TaskStage
	}
	return nil
}

type TaskStageID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *TaskStageID) Reset() {
	*x = TaskStageID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskStageID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStageID) ProtoMessage() {}

func (x *TaskStageID) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStageID.ProtoReflect.Descriptor instead.
func (*TaskStageID) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{4}
}

func (x *TaskStageID) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type SubscribeUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	TaskID uint64 `protobuf:"varint,2,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
}

func (x *SubscribeUserRequest) Reset() {
	*x = SubscribeUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeUserRequest) ProtoMessage() {}

func (x *SubscribeUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeUserRequest.ProtoReflect.Descriptor instead.
func (*SubscribeUserRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{5}
}

func (x *SubscribeUserRequest) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *SubscribeUserRequest) GetTaskID() uint64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

type SubscribeUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSubscribed bool `protobuf:"varint,1,opt,name=isSubscribed,proto3" json:"isSubscribed,omitempty"`
}

func (x *SubscribeUserResponse) Reset() {
	*x = SubscribeUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeUserResponse) ProtoMessage() {}

func (x *SubscribeUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeUserResponse.ProtoReflect.Descriptor instead.
func (*SubscribeUserResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{6}
}

func (x *SubscribeUserResponse) GetIsSubscribed() bool {
	if x != nil {
		return x.IsSubscribed
	}
	return false
}

type UnsubscribeUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	TaskID uint64 `protobuf:"varint,2,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
}

func (x *UnsubscribeUserRequest) Reset() {
	*x = UnsubscribeUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeUserRequest) ProtoMessage() {}

func (x *UnsubscribeUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeUserRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeUserRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{7}
}

func (x *UnsubscribeUserRequest) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UnsubscribeUserRequest) GetTaskID() uint64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

type UnsubscribeUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsUnsubscribed bool `protobuf:"varint,1,opt,name=isUnsubscribed,proto3" json:"isUnsubscribed,omitempty"`
}

func (x *UnsubscribeUserResponse) Reset() {
	*x = UnsubscribeUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeUserResponse) ProtoMessage() {}

func (x *UnsubscribeUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeUserResponse.ProtoReflect.Descriptor instead.
func (*UnsubscribeUserResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{8}
}

func (x *UnsubscribeUserResponse) GetIsUnsubscribed() bool {
	if x != nil {
		return x.IsUnsubscribed
	}
	return false
}

type GetTasksListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTasksListRequest) Reset() {
	*x = GetTasksListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTasksListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksListRequest) ProtoMessage() {}

func (x *GetTasksListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksListRequest.ProtoReflect.Descriptor instead.
func (*GetTasksListRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{9}
}

type TasksList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TasksList) Reset() {
	*x = TasksList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TasksList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TasksList) ProtoMessage() {}

func (x *TasksList) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TasksList.ProtoReflect.Descriptor instead.
func (*TasksList) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{10}
}

func (x *TasksList) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint64       `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	TaskStages  []*TaskStage `protobuf:"bytes,4,rep,name=taskStages,proto3" json:"taskStages,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{11}
}

func (x *Task) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetTaskStages() []*TaskStage {
	if x != nil {
		return x.TaskStages
	}
	return nil
}

type TaskStage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID               uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description      string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	MinutesFromStart uint64 `protobuf:"varint,4,opt,name=minutesFromStart,proto3" json:"minutesFromStart,omitempty"`
	DurationMinutes  uint64 `protobuf:"varint,5,opt,name=durationMinutes,proto3" json:"durationMinutes,omitempty"`
}

func (x *TaskStage) Reset() {
	*x = TaskStage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskStage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStage) ProtoMessage() {}

func (x *TaskStage) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStage.ProtoReflect.Descriptor instead.
func (*TaskStage) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{12}
}

func (x *TaskStage) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *TaskStage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskStage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TaskStage) GetMinutesFromStart() uint64 {
	if x != nil {
		return x.MinutesFromStart
	}
	return 0
}

func (x *TaskStage) GetDurationMinutes() uint64 {
	if x != nil {
		return x.DurationMinutes
	}
	return 0
}

type UserTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Task      *Task                  `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	StartTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=startTime,proto3" json:"startTime,omitempty"`
}

func (x *UserTask) Reset() {
	*x = UserTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTask) ProtoMessage() {}

func (x *UserTask) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTask.ProtoReflect.Descriptor instead.
func (*UserTask) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{13}
}

func (x *UserTask) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserTask) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *UserTask) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

type GetUserTasksByTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint64                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TimeFrom *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timeFrom,proto3" json:"timeFrom,omitempty"`
	TimeTo   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timeTo,proto3" json:"timeTo,omitempty"`
}

func (x *GetUserTasksByTimeRequest) Reset() {
	*x = GetUserTasksByTimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTasksByTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTasksByTimeRequest) ProtoMessage() {}

func (x *GetUserTasksByTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTasksByTimeRequest.ProtoReflect.Descriptor instead.
func (*GetUserTasksByTimeRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{14}
}

func (x *GetUserTasksByTimeRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUserTasksByTimeRequest) GetTimeFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeFrom
	}
	return nil
}

func (x *GetUserTasksByTimeRequest) GetTimeTo() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeTo
	}
	return nil
}

var File_api_v1_api_proto protoreflect.FileDescriptor

var file_api_v1_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x13, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x18, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x5a, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x09,
	0x74, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x09, 0x74, 0x61, 0x73,
	0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x46, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x22, 0x3b, 0x0a,
	0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x22, 0x48, 0x0a, 0x16, 0x55, 0x6e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x44, 0x22, 0x41, 0x0a, 0x17, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x69, 0x73, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x55, 0x6e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28,
	0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x78, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65,
	0x73, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x10, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x69,
	0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x22, 0x7a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x38,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x42, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x36, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x74,
	0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x32, 0x0a, 0x06, 0x74, 0x69, 0x6d, 0x65, 0x54,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x06, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x32, 0x90, 0x03, 0x0a, 0x09,
	0x53, 0x6f, 0x6d, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x12, 0x29, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x05, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x07, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x44, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x44, 0x22, 0x00,
	0x12, 0x32, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x55, 0x6e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x42, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x42, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0a, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65,
	0x76, 0x2f, 0x65, 0x6d, 0x69, 0x6c, 0x67, 0x61, 0x6c, 0x69, 0x6d, 0x6f, 0x76, 0x2f, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_api_proto_rawDescOnce sync.Once
	file_api_v1_api_proto_rawDescData = file_api_v1_api_proto_rawDesc
)

func file_api_v1_api_proto_rawDescGZIP() []byte {
	file_api_v1_api_proto_rawDescOnce.Do(func() {
		file_api_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_api_proto_rawDescData)
	})
	return file_api_v1_api_proto_rawDescData
}

var file_api_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_api_v1_api_proto_goTypes = []interface{}{
	(*User)(nil),                      // 0: User
	(*CreateUserRequest)(nil),         // 1: CreateUserRequest
	(*TaskID)(nil),                    // 2: TaskID
	(*CreateTaskStageRequest)(nil),    // 3: CreateTaskStageRequest
	(*TaskStageID)(nil),               // 4: TaskStageID
	(*SubscribeUserRequest)(nil),      // 5: SubscribeUserRequest
	(*SubscribeUserResponse)(nil),     // 6: SubscribeUserResponse
	(*UnsubscribeUserRequest)(nil),    // 7: UnsubscribeUserRequest
	(*UnsubscribeUserResponse)(nil),   // 8: UnsubscribeUserResponse
	(*GetTasksListRequest)(nil),       // 9: GetTasksListRequest
	(*TasksList)(nil),                 // 10: TasksList
	(*Task)(nil),                      // 11: Task
	(*TaskStage)(nil),                 // 12: TaskStage
	(*UserTask)(nil),                  // 13: UserTask
	(*GetUserTasksByTimeRequest)(nil), // 14: GetUserTasksByTimeRequest
	(*timestamppb.Timestamp)(nil),     // 15: google.protobuf.Timestamp
}
var file_api_v1_api_proto_depIdxs = []int32{
	12, // 0: CreateTaskStageRequest.taskStage:type_name -> TaskStage
	11, // 1: TasksList.tasks:type_name -> Task
	12, // 2: Task.taskStages:type_name -> TaskStage
	0,  // 3: UserTask.user:type_name -> User
	11, // 4: UserTask.task:type_name -> Task
	15, // 5: UserTask.startTime:type_name -> google.protobuf.Timestamp
	15, // 6: GetUserTasksByTimeRequest.timeFrom:type_name -> google.protobuf.Timestamp
	15, // 7: GetUserTasksByTimeRequest.timeTo:type_name -> google.protobuf.Timestamp
	1,  // 8: SomeStuff.CreateUser:input_type -> CreateUserRequest
	11, // 9: SomeStuff.CreateTask:input_type -> Task
	3,  // 10: SomeStuff.CreateTaskStage:input_type -> CreateTaskStageRequest
	9,  // 11: SomeStuff.GetTasksList:input_type -> GetTasksListRequest
	5,  // 12: SomeStuff.SubscribeUser:input_type -> SubscribeUserRequest
	7,  // 13: SomeStuff.UnsubscribeUser:input_type -> UnsubscribeUserRequest
	14, // 14: SomeStuff.GetUserTasksByTime:input_type -> GetUserTasksByTimeRequest
	0,  // 15: SomeStuff.CreateUser:output_type -> User
	2,  // 16: SomeStuff.CreateTask:output_type -> TaskID
	4,  // 17: SomeStuff.CreateTaskStage:output_type -> TaskStageID
	10, // 18: SomeStuff.GetTasksList:output_type -> TasksList
	6,  // 19: SomeStuff.SubscribeUser:output_type -> SubscribeUserResponse
	8,  // 20: SomeStuff.UnsubscribeUser:output_type -> UnsubscribeUserResponse
	10, // 21: SomeStuff.GetUserTasksByTime:output_type -> TasksList
	15, // [15:22] is the sub-list for method output_type
	8,  // [8:15] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_v1_api_proto_init() }
func file_api_v1_api_proto_init() {
	if File_api_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_api_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_api_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskID); i {
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
		file_api_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskStageRequest); i {
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
		file_api_v1_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskStageID); i {
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
		file_api_v1_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeUserRequest); i {
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
		file_api_v1_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeUserResponse); i {
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
		file_api_v1_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeUserRequest); i {
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
		file_api_v1_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeUserResponse); i {
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
		file_api_v1_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTasksListRequest); i {
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
		file_api_v1_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TasksList); i {
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
		file_api_v1_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_api_v1_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskStage); i {
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
		file_api_v1_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTask); i {
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
		file_api_v1_api_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTasksByTimeRequest); i {
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
			RawDescriptor: file_api_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_api_proto_goTypes,
		DependencyIndexes: file_api_v1_api_proto_depIdxs,
		MessageInfos:      file_api_v1_api_proto_msgTypes,
	}.Build()
	File_api_v1_api_proto = out.File
	file_api_v1_api_proto_rawDesc = nil
	file_api_v1_api_proto_goTypes = nil
	file_api_v1_api_proto_depIdxs = nil
}
