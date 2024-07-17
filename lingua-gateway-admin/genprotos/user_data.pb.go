// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.1
// source: lingua-protos/user_data.proto

package genprotos

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

type UserDataGRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level             string  `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	NativeLang        string  `protobuf:"bytes,2,opt,name=native_lang,json=nativeLang,proto3" json:"native_lang,omitempty"`
	Xp                int64   `protobuf:"varint,3,opt,name=xp,proto3" json:"xp,omitempty"`
	DailyStreak       int32   `protobuf:"varint,4,opt,name=daily_streak,json=dailyStreak,proto3" json:"daily_streak,omitempty"`
	PlayedGamesCount  int64   `protobuf:"varint,5,opt,name=played_games_count,json=playedGamesCount,proto3" json:"played_games_count,omitempty"`
	WinningPercentage float32 `protobuf:"fixed32,6,opt,name=winning_percentage,json=winningPercentage,proto3" json:"winning_percentage,omitempty"`
}

func (x *UserDataGRes) Reset() {
	*x = UserDataGRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lingua_protos_user_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDataGRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDataGRes) ProtoMessage() {}

func (x *UserDataGRes) ProtoReflect() protoreflect.Message {
	mi := &file_lingua_protos_user_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDataGRes.ProtoReflect.Descriptor instead.
func (*UserDataGRes) Descriptor() ([]byte, []int) {
	return file_lingua_protos_user_data_proto_rawDescGZIP(), []int{0}
}

func (x *UserDataGRes) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *UserDataGRes) GetNativeLang() string {
	if x != nil {
		return x.NativeLang
	}
	return ""
}

func (x *UserDataGRes) GetXp() int64 {
	if x != nil {
		return x.Xp
	}
	return 0
}

func (x *UserDataGRes) GetDailyStreak() int32 {
	if x != nil {
		return x.DailyStreak
	}
	return 0
}

func (x *UserDataGRes) GetPlayedGamesCount() int64 {
	if x != nil {
		return x.PlayedGamesCount
	}
	return 0
}

func (x *UserDataGRes) GetWinningPercentage() float32 {
	if x != nil {
		return x.WinningPercentage
	}
	return 0
}

type XPUReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Xp     int64  `protobuf:"varint,2,opt,name=xp,proto3" json:"xp,omitempty"`
}

func (x *XPUReq) Reset() {
	*x = XPUReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lingua_protos_user_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XPUReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XPUReq) ProtoMessage() {}

func (x *XPUReq) ProtoReflect() protoreflect.Message {
	mi := &file_lingua_protos_user_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XPUReq.ProtoReflect.Descriptor instead.
func (*XPUReq) Descriptor() ([]byte, []int) {
	return file_lingua_protos_user_data_proto_rawDescGZIP(), []int{1}
}

func (x *XPUReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *XPUReq) GetXp() int64 {
	if x != nil {
		return x.Xp
	}
	return 0
}

type StreakUReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DailyStreak int64  `protobuf:"varint,2,opt,name=daily_streak,json=dailyStreak,proto3" json:"daily_streak,omitempty"`
}

func (x *StreakUReq) Reset() {
	*x = StreakUReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lingua_protos_user_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreakUReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreakUReq) ProtoMessage() {}

func (x *StreakUReq) ProtoReflect() protoreflect.Message {
	mi := &file_lingua_protos_user_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreakUReq.ProtoReflect.Descriptor instead.
func (*StreakUReq) Descriptor() ([]byte, []int) {
	return file_lingua_protos_user_data_proto_rawDescGZIP(), []int{2}
}

func (x *StreakUReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *StreakUReq) GetDailyStreak() int64 {
	if x != nil {
		return x.DailyStreak
	}
	return 0
}

type PlayedGamesCountUReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PlayedGamesCount int64  `protobuf:"varint,2,opt,name=played_games_count,json=playedGamesCount,proto3" json:"played_games_count,omitempty"`
}

func (x *PlayedGamesCountUReq) Reset() {
	*x = PlayedGamesCountUReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lingua_protos_user_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayedGamesCountUReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayedGamesCountUReq) ProtoMessage() {}

func (x *PlayedGamesCountUReq) ProtoReflect() protoreflect.Message {
	mi := &file_lingua_protos_user_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayedGamesCountUReq.ProtoReflect.Descriptor instead.
func (*PlayedGamesCountUReq) Descriptor() ([]byte, []int) {
	return file_lingua_protos_user_data_proto_rawDescGZIP(), []int{3}
}

func (x *PlayedGamesCountUReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PlayedGamesCountUReq) GetPlayedGamesCount() int64 {
	if x != nil {
		return x.PlayedGamesCount
	}
	return 0
}

type WinningPercentageUReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Percentage float32 `protobuf:"fixed32,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *WinningPercentageUReq) Reset() {
	*x = WinningPercentageUReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lingua_protos_user_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinningPercentageUReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinningPercentageUReq) ProtoMessage() {}

func (x *WinningPercentageUReq) ProtoReflect() protoreflect.Message {
	mi := &file_lingua_protos_user_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinningPercentageUReq.ProtoReflect.Descriptor instead.
func (*WinningPercentageUReq) Descriptor() ([]byte, []int) {
	return file_lingua_protos_user_data_proto_rawDescGZIP(), []int{4}
}

func (x *WinningPercentageUReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *WinningPercentageUReq) GetPercentage() float32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

var File_lingua_protos_user_data_proto protoreflect.FileDescriptor

var file_lingua_protos_user_data_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6c, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6c, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x1a, 0x18, 0x6c, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x47, 0x52,
	0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x78, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x78, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x69,
	0x6c, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x12, 0x2c, 0x0a, 0x12,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64,
	0x47, 0x61, 0x6d, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x77, 0x69,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x77, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x06, 0x58, 0x50, 0x55,
	0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x78, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x78, 0x70, 0x22, 0x48, 0x0a, 0x0a,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x55, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x61, 0x69, 0x6c, 0x79,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x22, 0x5d, 0x0a, 0x14, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64,
	0x47, 0x61, 0x6d, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x64, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x15, 0x57, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x55, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x32, 0xbd, 0x02, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0c, 0x2e, 0x6c, 0x69, 0x6e,
	0x67, 0x75, 0x61, 0x2e, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x14, 0x2e, 0x6c, 0x69, 0x6e, 0x67, 0x75,
	0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x47, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x2a, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x58, 0x50, 0x12, 0x0e, 0x2e, 0x6c,
	0x69, 0x6e, 0x67, 0x75, 0x61, 0x2e, 0x58, 0x50, 0x55, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x6c,
	0x69, 0x6e, 0x67, 0x75, 0x61, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6b, 0x12, 0x12, 0x2e, 0x6c, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6b, 0x55, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x6c, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1c, 0x2e, 0x6c, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x47,
	0x61, 0x6d, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e,
	0x6c, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x6c, 0x69, 0x6e, 0x67, 0x75,
	0x61, 0x2e, 0x57, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x55, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x6c, 0x69, 0x6e, 0x67, 0x75, 0x61,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lingua_protos_user_data_proto_rawDescOnce sync.Once
	file_lingua_protos_user_data_proto_rawDescData = file_lingua_protos_user_data_proto_rawDesc
)

func file_lingua_protos_user_data_proto_rawDescGZIP() []byte {
	file_lingua_protos_user_data_proto_rawDescOnce.Do(func() {
		file_lingua_protos_user_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_lingua_protos_user_data_proto_rawDescData)
	})
	return file_lingua_protos_user_data_proto_rawDescData
}

var file_lingua_protos_user_data_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_lingua_protos_user_data_proto_goTypes = []any{
	(*UserDataGRes)(nil),          // 0: lingua.UserDataGRes
	(*XPUReq)(nil),                // 1: lingua.XPUReq
	(*StreakUReq)(nil),            // 2: lingua.StreakUReq
	(*PlayedGamesCountUReq)(nil),  // 3: lingua.PlayedGamesCountUReq
	(*WinningPercentageUReq)(nil), // 4: lingua.WinningPercentageUReq
	(*ByID)(nil),                  // 5: lingua.ByID
	(*Void)(nil),                  // 6: lingua.Void
}
var file_lingua_protos_user_data_proto_depIdxs = []int32{
	5, // 0: lingua.UserDataService.GetUserData:input_type -> lingua.ByID
	1, // 1: lingua.UserDataService.UpdateXP:input_type -> lingua.XPUReq
	2, // 2: lingua.UserDataService.UpdateDailyStreak:input_type -> lingua.StreakUReq
	3, // 3: lingua.UserDataService.UpdatePlayedGamesCount:input_type -> lingua.PlayedGamesCountUReq
	4, // 4: lingua.UserDataService.UpdateWinningPercentage:input_type -> lingua.WinningPercentageUReq
	0, // 5: lingua.UserDataService.GetUserData:output_type -> lingua.UserDataGRes
	6, // 6: lingua.UserDataService.UpdateXP:output_type -> lingua.Void
	6, // 7: lingua.UserDataService.UpdateDailyStreak:output_type -> lingua.Void
	6, // 8: lingua.UserDataService.UpdatePlayedGamesCount:output_type -> lingua.Void
	6, // 9: lingua.UserDataService.UpdateWinningPercentage:output_type -> lingua.Void
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lingua_protos_user_data_proto_init() }
func file_lingua_protos_user_data_proto_init() {
	if File_lingua_protos_user_data_proto != nil {
		return
	}
	file_lingua_protos_void_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lingua_protos_user_data_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserDataGRes); i {
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
		file_lingua_protos_user_data_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*XPUReq); i {
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
		file_lingua_protos_user_data_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StreakUReq); i {
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
		file_lingua_protos_user_data_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PlayedGamesCountUReq); i {
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
		file_lingua_protos_user_data_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*WinningPercentageUReq); i {
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
			RawDescriptor: file_lingua_protos_user_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lingua_protos_user_data_proto_goTypes,
		DependencyIndexes: file_lingua_protos_user_data_proto_depIdxs,
		MessageInfos:      file_lingua_protos_user_data_proto_msgTypes,
	}.Build()
	File_lingua_protos_user_data_proto = out.File
	file_lingua_protos_user_data_proto_rawDesc = nil
	file_lingua_protos_user_data_proto_goTypes = nil
	file_lingua_protos_user_data_proto_depIdxs = nil
}