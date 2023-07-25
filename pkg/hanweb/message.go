package hanweb

import "github.com/golang/protobuf/proto"

type Data struct {
	Event         int32       `protobuf:"varint,1,opt,name=event,proto3" json:"event,omitempty"`
	ParticipantId string      `protobuf:"bytes,2,opt,name=participant_id,proto3" json:"participant_id,omitempty"`
	RoomId        string      `protobuf:"bytes,3,opt,name=room_id,proto3" json:"room_id,omitempty"`
	DetailInfo    []*KeyValue `protobuf:"bytes,4,rep,name=detail_info,json=detailInfo,proto3" json:"detail_info,omitempty"`
}

type KeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

// Reset 新增 Data 结构体的 Reset、String 和 ProtoMessage 方法
func (d *Data) Reset() {
	*d = Data{}
}

func (d *Data) String() string {
	return proto.CompactTextString(d)
}

func (*Data) ProtoMessage() {}

func (d *KeyValue) Reset() {
	*d = KeyValue{}
}

func (d *KeyValue) String() string {
	return proto.CompactTextString(d)
}

// ProtoMessage Implement ProtoMessage for KeyValue
func (*KeyValue) ProtoMessage() {}
