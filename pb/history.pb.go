// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: history.proto

package pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type History struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAgent      string               `protobuf:"bytes,1,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	UserId         string               `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClientIp       string               `protobuf:"bytes,3,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	MangadexId     string               `protobuf:"bytes,4,opt,name=mangadex_id,json=mangadexId,proto3" json:"mangadex_id,omitempty"`
	AlId           string               `protobuf:"bytes,5,opt,name=al_id,json=alId,proto3" json:"al_id,omitempty"`
	Path           string               `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	CoverImage     string               `protobuf:"bytes,7,opt,name=cover_image,json=coverImage,proto3" json:"cover_image,omitempty"`
	Title          string               `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
	ReadingChapter string               `protobuf:"bytes,9,opt,name=reading_chapter,json=readingChapter,proto3" json:"reading_chapter,omitempty"`
	LastReadAt     *timestamp.Timestamp `protobuf:"bytes,10,opt,name=last_read_at,json=lastReadAt,proto3" json:"last_read_at,omitempty"`
	CreatedAt      *timestamp.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *History) Reset() {
	*x = History{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *History) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*History) ProtoMessage() {}

func (x *History) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use History.ProtoReflect.Descriptor instead.
func (*History) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{0}
}

func (x *History) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *History) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *History) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

func (x *History) GetMangadexId() string {
	if x != nil {
		return x.MangadexId
	}
	return ""
}

func (x *History) GetAlId() string {
	if x != nil {
		return x.AlId
	}
	return ""
}

func (x *History) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *History) GetCoverImage() string {
	if x != nil {
		return x.CoverImage
	}
	return ""
}

func (x *History) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *History) GetReadingChapter() string {
	if x != nil {
		return x.ReadingChapter
	}
	return ""
}

func (x *History) GetLastReadAt() *timestamp.Timestamp {
	if x != nil {
		return x.LastReadAt
	}
	return nil
}

func (x *History) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAgent      string `protobuf:"bytes,1,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	MangadexId     string `protobuf:"bytes,2,opt,name=mangadex_id,json=mangadexId,proto3" json:"mangadex_id,omitempty"`
	CoverImage     string `protobuf:"bytes,3,opt,name=cover_image,json=coverImage,proto3" json:"cover_image,omitempty"`
	Title          string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	ReadingChapter string `protobuf:"bytes,5,opt,name=reading_chapter,json=readingChapter,proto3" json:"reading_chapter,omitempty"`
	AlId           string `protobuf:"bytes,6,opt,name=al_id,json=alId,proto3" json:"al_id,omitempty"`
	Path           string `protobuf:"bytes,7,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *CreateHistoryRequest) Reset() {
	*x = CreateHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHistoryRequest) ProtoMessage() {}

func (x *CreateHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHistoryRequest.ProtoReflect.Descriptor instead.
func (*CreateHistoryRequest) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{1}
}

func (x *CreateHistoryRequest) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *CreateHistoryRequest) GetMangadexId() string {
	if x != nil {
		return x.MangadexId
	}
	return ""
}

func (x *CreateHistoryRequest) GetCoverImage() string {
	if x != nil {
		return x.CoverImage
	}
	return ""
}

func (x *CreateHistoryRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateHistoryRequest) GetReadingChapter() string {
	if x != nil {
		return x.ReadingChapter
	}
	return ""
}

func (x *CreateHistoryRequest) GetAlId() string {
	if x != nil {
		return x.AlId
	}
	return ""
}

func (x *CreateHistoryRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type CreateHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok      bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateHistoryResponse) Reset() {
	*x = CreateHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHistoryResponse) ProtoMessage() {}

func (x *CreateHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHistoryResponse.ProtoReflect.Descriptor instead.
func (*CreateHistoryResponse) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{2}
}

func (x *CreateHistoryResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *CreateHistoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetHistoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAgent *string `protobuf:"bytes,1,opt,name=user_agent,json=userAgent,proto3,oneof" json:"user_agent,omitempty"`
	Limit     int32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset    int32   `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetHistoriesRequest) Reset() {
	*x = GetHistoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoriesRequest) ProtoMessage() {}

func (x *GetHistoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoriesRequest.ProtoReflect.Descriptor instead.
func (*GetHistoriesRequest) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{3}
}

func (x *GetHistoriesRequest) GetUserAgent() string {
	if x != nil && x.UserAgent != nil {
		return *x.UserAgent
	}
	return ""
}

func (x *GetHistoriesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetHistoriesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type HistoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Histories []*History `protobuf:"bytes,1,rep,name=histories,proto3" json:"histories,omitempty"`
}

func (x *HistoriesResponse) Reset() {
	*x = HistoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoriesResponse) ProtoMessage() {}

func (x *HistoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoriesResponse.ProtoReflect.Descriptor instead.
func (*HistoriesResponse) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{4}
}

func (x *HistoriesResponse) GetHistories() []*History {
	if x != nil {
		return x.Histories
	}
	return nil
}

type UpdateHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MangadexId     string  `protobuf:"bytes,1,opt,name=mangadex_id,json=mangadexId,proto3" json:"mangadex_id,omitempty"`
	ReadingChapter string  `protobuf:"bytes,2,opt,name=reading_chapter,json=readingChapter,proto3" json:"reading_chapter,omitempty"`
	UserAgent      *string `protobuf:"bytes,3,opt,name=user_agent,json=userAgent,proto3,oneof" json:"user_agent,omitempty"`
	UserId         *string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
}

func (x *UpdateHistoryRequest) Reset() {
	*x = UpdateHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHistoryRequest) ProtoMessage() {}

func (x *UpdateHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHistoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateHistoryRequest) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateHistoryRequest) GetMangadexId() string {
	if x != nil {
		return x.MangadexId
	}
	return ""
}

func (x *UpdateHistoryRequest) GetReadingChapter() string {
	if x != nil {
		return x.ReadingChapter
	}
	return ""
}

func (x *UpdateHistoryRequest) GetUserAgent() string {
	if x != nil && x.UserAgent != nil {
		return *x.UserAgent
	}
	return ""
}

func (x *UpdateHistoryRequest) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

type UpdateHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok      bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateHistoryResponse) Reset() {
	*x = UpdateHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHistoryResponse) ProtoMessage() {}

func (x *UpdateHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHistoryResponse.ProtoReflect.Descriptor instead.
func (*UpdateHistoryResponse) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateHistoryResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *UpdateHistoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_history_proto protoreflect.FileDescriptor

var file_history_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x81, 0x03, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x64, 0x65, 0x78,
	0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xdf, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x61, 0x6e, 0x67, 0x61, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x64, 0x65, 0x78, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72,
	0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x13, 0x0a,
	0x05, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x6c,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x41, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x76, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x22, 0x3b, 0x0a, 0x11, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0xbd,
	0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x6e, 0x67, 0x61,
	0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61,
	0x6e, 0x67, 0x61, 0x64, 0x65, 0x78, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x12, 0x22, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x41,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x64, 0x65, 0x76, 0x39, 0x39, 0x2f, 0x64, 0x72, 0x6f, 0x70,
	0x62, 0x79, 0x74, 0x65, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_history_proto_rawDescOnce sync.Once
	file_history_proto_rawDescData = file_history_proto_rawDesc
)

func file_history_proto_rawDescGZIP() []byte {
	file_history_proto_rawDescOnce.Do(func() {
		file_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_history_proto_rawDescData)
	})
	return file_history_proto_rawDescData
}

var file_history_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_history_proto_goTypes = []interface{}{
	(*History)(nil),               // 0: History
	(*CreateHistoryRequest)(nil),  // 1: CreateHistoryRequest
	(*CreateHistoryResponse)(nil), // 2: CreateHistoryResponse
	(*GetHistoriesRequest)(nil),   // 3: GetHistoriesRequest
	(*HistoriesResponse)(nil),     // 4: HistoriesResponse
	(*UpdateHistoryRequest)(nil),  // 5: UpdateHistoryRequest
	(*UpdateHistoryResponse)(nil), // 6: UpdateHistoryResponse
	(*timestamp.Timestamp)(nil),   // 7: google.protobuf.Timestamp
}
var file_history_proto_depIdxs = []int32{
	7, // 0: History.last_read_at:type_name -> google.protobuf.Timestamp
	7, // 1: History.created_at:type_name -> google.protobuf.Timestamp
	0, // 2: HistoriesResponse.histories:type_name -> History
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_history_proto_init() }
func file_history_proto_init() {
	if File_history_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*History); i {
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
		file_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHistoryRequest); i {
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
		file_history_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHistoryResponse); i {
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
		file_history_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHistoriesRequest); i {
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
		file_history_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoriesResponse); i {
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
		file_history_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHistoryRequest); i {
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
		file_history_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHistoryResponse); i {
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
	file_history_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_history_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_history_proto_goTypes,
		DependencyIndexes: file_history_proto_depIdxs,
		MessageInfos:      file_history_proto_msgTypes,
	}.Build()
	File_history_proto = out.File
	file_history_proto_rawDesc = nil
	file_history_proto_goTypes = nil
	file_history_proto_depIdxs = nil
}
