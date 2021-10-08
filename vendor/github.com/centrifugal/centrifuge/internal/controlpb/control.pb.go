// Install protoc compiler https://github.com/google/protobuf/releases
// Install gogofaster program:
// go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
// protoc --proto_path=../../vendor:. --gogofaster_out=. control.proto
// Note that we use vendored gogoprotobuf path in example above.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: control.proto

package controlpb

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

type Command_MethodType int32

const (
	Command_NODE            Command_MethodType = 0
	Command_UNSUBSCRIBE     Command_MethodType = 1
	Command_DISCONNECT      Command_MethodType = 2
	Command_SHUTDOWN        Command_MethodType = 3
	Command_SURVEY_REQUEST  Command_MethodType = 4
	Command_SURVEY_RESPONSE Command_MethodType = 5
	Command_SUBSCRIBE       Command_MethodType = 6
	Command_NOTIFICATION    Command_MethodType = 7
	Command_REFRESH         Command_MethodType = 8
)

// Enum value maps for Command_MethodType.
var (
	Command_MethodType_name = map[int32]string{
		0: "NODE",
		1: "UNSUBSCRIBE",
		2: "DISCONNECT",
		3: "SHUTDOWN",
		4: "SURVEY_REQUEST",
		5: "SURVEY_RESPONSE",
		6: "SUBSCRIBE",
		7: "NOTIFICATION",
		8: "REFRESH",
	}
	Command_MethodType_value = map[string]int32{
		"NODE":            0,
		"UNSUBSCRIBE":     1,
		"DISCONNECT":      2,
		"SHUTDOWN":        3,
		"SURVEY_REQUEST":  4,
		"SURVEY_RESPONSE": 5,
		"SUBSCRIBE":       6,
		"NOTIFICATION":    7,
		"REFRESH":         8,
	}
)

func (x Command_MethodType) Enum() *Command_MethodType {
	p := new(Command_MethodType)
	*p = x
	return p
}

func (x Command_MethodType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Command_MethodType) Descriptor() protoreflect.EnumDescriptor {
	return file_control_proto_enumTypes[0].Descriptor()
}

func (Command_MethodType) Type() protoreflect.EnumType {
	return &file_control_proto_enumTypes[0]
}

func (x Command_MethodType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Command_MethodType.Descriptor instead.
func (Command_MethodType) EnumDescriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{0, 0}
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string             `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Method Command_MethodType `protobuf:"varint,2,opt,name=method,proto3,enum=controlpb.Command_MethodType" json:"method,omitempty"`
	Params []byte             `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Command) GetMethod() Command_MethodType {
	if x != nil {
		return x.Method
	}
	return Command_NODE
}

func (x *Command) GetParams() []byte {
	if x != nil {
		return x.Params
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version     string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	NumClients  uint32   `protobuf:"varint,4,opt,name=num_clients,json=numClients,proto3" json:"num_clients,omitempty"`
	NumUsers    uint32   `protobuf:"varint,5,opt,name=num_users,json=numUsers,proto3" json:"num_users,omitempty"`
	NumChannels uint32   `protobuf:"varint,6,opt,name=num_channels,json=numChannels,proto3" json:"num_channels,omitempty"`
	Uptime      uint32   `protobuf:"varint,7,opt,name=uptime,proto3" json:"uptime,omitempty"`
	Metrics     *Metrics `protobuf:"bytes,8,opt,name=metrics,proto3" json:"metrics,omitempty"`
	Data        []byte   `protobuf:"bytes,9,opt,name=data,proto3" json:"data,omitempty"`
	NumSubs     uint32   `protobuf:"varint,10,opt,name=num_subs,json=numSubs,proto3" json:"num_subs,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Node) GetNumClients() uint32 {
	if x != nil {
		return x.NumClients
	}
	return 0
}

func (x *Node) GetNumUsers() uint32 {
	if x != nil {
		return x.NumUsers
	}
	return 0
}

func (x *Node) GetNumChannels() uint32 {
	if x != nil {
		return x.NumChannels
	}
	return 0
}

func (x *Node) GetUptime() uint32 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *Node) GetMetrics() *Metrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *Node) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Node) GetNumSubs() uint32 {
	if x != nil {
		return x.NumSubs
	}
	return 0
}

type Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval float64            `protobuf:"fixed64,1,opt,name=interval,proto3" json:"interval,omitempty"`
	Items    map[string]float64 `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *Metrics) Reset() {
	*x = Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics) ProtoMessage() {}

func (x *Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics.ProtoReflect.Descriptor instead.
func (*Metrics) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{2}
}

func (x *Metrics) GetInterval() float64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *Metrics) GetItems() map[string]float64 {
	if x != nil {
		return x.Items
	}
	return nil
}

type Subscribe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User         string          `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Channel      string          `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Presence     bool            `protobuf:"varint,3,opt,name=presence,proto3" json:"presence,omitempty"`
	JoinLeave    bool            `protobuf:"varint,4,opt,name=join_leave,json=joinLeave,proto3" json:"join_leave,omitempty"`
	ExpireAt     int64           `protobuf:"varint,5,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
	Position     bool            `protobuf:"varint,6,opt,name=position,proto3" json:"position,omitempty"`
	Recover      bool            `protobuf:"varint,7,opt,name=recover,proto3" json:"recover,omitempty"`
	ChannelInfo  []byte          `protobuf:"bytes,8,opt,name=channel_info,json=channelInfo,proto3" json:"channel_info,omitempty"`
	Client       string          `protobuf:"bytes,9,opt,name=client,proto3" json:"client,omitempty"`
	Data         []byte          `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
	RecoverSince *StreamPosition `protobuf:"bytes,11,opt,name=recover_since,json=recoverSince,proto3" json:"recover_since,omitempty"`
}

func (x *Subscribe) Reset() {
	*x = Subscribe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscribe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscribe) ProtoMessage() {}

func (x *Subscribe) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscribe.ProtoReflect.Descriptor instead.
func (*Subscribe) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{3}
}

func (x *Subscribe) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Subscribe) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Subscribe) GetPresence() bool {
	if x != nil {
		return x.Presence
	}
	return false
}

func (x *Subscribe) GetJoinLeave() bool {
	if x != nil {
		return x.JoinLeave
	}
	return false
}

func (x *Subscribe) GetExpireAt() int64 {
	if x != nil {
		return x.ExpireAt
	}
	return 0
}

func (x *Subscribe) GetPosition() bool {
	if x != nil {
		return x.Position
	}
	return false
}

func (x *Subscribe) GetRecover() bool {
	if x != nil {
		return x.Recover
	}
	return false
}

func (x *Subscribe) GetChannelInfo() []byte {
	if x != nil {
		return x.ChannelInfo
	}
	return nil
}

func (x *Subscribe) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *Subscribe) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Subscribe) GetRecoverSince() *StreamPosition {
	if x != nil {
		return x.RecoverSince
	}
	return nil
}

type StreamPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset uint64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Epoch  string `protobuf:"bytes,2,opt,name=epoch,proto3" json:"epoch,omitempty"`
}

func (x *StreamPosition) Reset() {
	*x = StreamPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamPosition) ProtoMessage() {}

func (x *StreamPosition) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamPosition.ProtoReflect.Descriptor instead.
func (*StreamPosition) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{4}
}

func (x *StreamPosition) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *StreamPosition) GetEpoch() string {
	if x != nil {
		return x.Epoch
	}
	return ""
}

type Unsubscribe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	User    string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Client  string `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *Unsubscribe) Reset() {
	*x = Unsubscribe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unsubscribe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unsubscribe) ProtoMessage() {}

func (x *Unsubscribe) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unsubscribe.ProtoReflect.Descriptor instead.
func (*Unsubscribe) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{5}
}

func (x *Unsubscribe) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Unsubscribe) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Unsubscribe) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

type Disconnect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Whitelist []string `protobuf:"bytes,2,rep,name=whitelist,proto3" json:"whitelist,omitempty"`
	Code      uint32   `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Reason    string   `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	Reconnect bool     `protobuf:"varint,5,opt,name=reconnect,proto3" json:"reconnect,omitempty"`
	Client    string   `protobuf:"bytes,6,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *Disconnect) Reset() {
	*x = Disconnect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Disconnect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Disconnect) ProtoMessage() {}

func (x *Disconnect) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Disconnect.ProtoReflect.Descriptor instead.
func (*Disconnect) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{6}
}

func (x *Disconnect) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Disconnect) GetWhitelist() []string {
	if x != nil {
		return x.Whitelist
	}
	return nil
}

func (x *Disconnect) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Disconnect) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Disconnect) GetReconnect() bool {
	if x != nil {
		return x.Reconnect
	}
	return false
}

func (x *Disconnect) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

type SurveyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Op   string `protobuf:"bytes,2,opt,name=op,proto3" json:"op,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SurveyRequest) Reset() {
	*x = SurveyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyRequest) ProtoMessage() {}

func (x *SurveyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyRequest.ProtoReflect.Descriptor instead.
func (*SurveyRequest) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{7}
}

func (x *SurveyRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SurveyRequest) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *SurveyRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SurveyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SurveyResponse) Reset() {
	*x = SurveyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyResponse) ProtoMessage() {}

func (x *SurveyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyResponse.ProtoReflect.Descriptor instead.
func (*SurveyResponse) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{8}
}

func (x *SurveyResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SurveyResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SurveyResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op   string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{9}
}

func (x *Notification) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *Notification) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Refresh struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Client   string `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Expired  bool   `protobuf:"varint,3,opt,name=expired,proto3" json:"expired,omitempty"`
	ExpireAt int64  `protobuf:"varint,4,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
	Info     []byte `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Refresh) Reset() {
	*x = Refresh{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Refresh) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Refresh) ProtoMessage() {}

func (x *Refresh) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Refresh.ProtoReflect.Descriptor instead.
func (*Refresh) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{10}
}

func (x *Refresh) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Refresh) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *Refresh) GetExpired() bool {
	if x != nil {
		return x.Expired
	}
	return false
}

func (x *Refresh) GetExpireAt() int64 {
	if x != nil {
		return x.ExpireAt
	}
	return 0
}

func (x *Refresh) GetInfo() []byte {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_control_proto protoreflect.FileDescriptor

var file_control_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x62, 0x22, 0x89, 0x02, 0x0a, 0x07, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x48, 0x55, 0x54, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x03, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x55, 0x52, 0x56, 0x45, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x55, 0x52, 0x56, 0x45, 0x59, 0x5f, 0x52, 0x45,
	0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x42, 0x53,
	0x43, 0x52, 0x49, 0x42, 0x45, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x54, 0x49, 0x46,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x46,
	0x52, 0x45, 0x53, 0x48, 0x10, 0x08, 0x22, 0x9c, 0x02, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x07, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x75,
	0x6d, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x75,
	0x6d, 0x53, 0x75, 0x62, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x33, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd6, 0x02, 0x0a,
	0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x61,
	0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72,
	0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x53, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x70, 0x6f, 0x63, 0x68, 0x22, 0x53, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0xa0, 0x01, 0x0a, 0x0a, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x43, 0x0a,
	0x0d, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x48, 0x0a, 0x0e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x32, 0x0a, 0x0c,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x80, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_control_proto_rawDescOnce sync.Once
	file_control_proto_rawDescData = file_control_proto_rawDesc
)

func file_control_proto_rawDescGZIP() []byte {
	file_control_proto_rawDescOnce.Do(func() {
		file_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_control_proto_rawDescData)
	})
	return file_control_proto_rawDescData
}

var file_control_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_control_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_control_proto_goTypes = []interface{}{
	(Command_MethodType)(0), // 0: controlpb.Command.MethodType
	(*Command)(nil),         // 1: controlpb.Command
	(*Node)(nil),            // 2: controlpb.Node
	(*Metrics)(nil),         // 3: controlpb.Metrics
	(*Subscribe)(nil),       // 4: controlpb.Subscribe
	(*StreamPosition)(nil),  // 5: controlpb.StreamPosition
	(*Unsubscribe)(nil),     // 6: controlpb.Unsubscribe
	(*Disconnect)(nil),      // 7: controlpb.Disconnect
	(*SurveyRequest)(nil),   // 8: controlpb.SurveyRequest
	(*SurveyResponse)(nil),  // 9: controlpb.SurveyResponse
	(*Notification)(nil),    // 10: controlpb.Notification
	(*Refresh)(nil),         // 11: controlpb.Refresh
	nil,                     // 12: controlpb.Metrics.ItemsEntry
}
var file_control_proto_depIdxs = []int32{
	0,  // 0: controlpb.Command.method:type_name -> controlpb.Command.MethodType
	3,  // 1: controlpb.Node.metrics:type_name -> controlpb.Metrics
	12, // 2: controlpb.Metrics.items:type_name -> controlpb.Metrics.ItemsEntry
	5,  // 3: controlpb.Subscribe.recover_since:type_name -> controlpb.StreamPosition
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_control_proto_init() }
func file_control_proto_init() {
	if File_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_control_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_control_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics); i {
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
		file_control_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscribe); i {
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
		file_control_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamPosition); i {
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
		file_control_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unsubscribe); i {
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
		file_control_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Disconnect); i {
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
		file_control_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyRequest); i {
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
		file_control_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyResponse); i {
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
		file_control_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
		file_control_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Refresh); i {
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
			RawDescriptor: file_control_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_control_proto_goTypes,
		DependencyIndexes: file_control_proto_depIdxs,
		EnumInfos:         file_control_proto_enumTypes,
		MessageInfos:      file_control_proto_msgTypes,
	}.Build()
	File_control_proto = out.File
	file_control_proto_rawDesc = nil
	file_control_proto_goTypes = nil
	file_control_proto_depIdxs = nil
}
