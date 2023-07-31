package hanweb

import "github.com/golang/protobuf/proto"

type DataOption func(*Data)

type Data struct {
	Event         CustomHookEvent `protobuf:"varint,1,opt,name=event,proto3" json:"event,omitempty"`
	ParticipantId string          `protobuf:"bytes,2,opt,name=participant_id,proto3" json:"participant_id,omitempty"`
	RoomId        string          `protobuf:"bytes,3,opt,name=room_id,proto3" json:"room_id,omitempty"`
	DetailInfo    []*KeyValue     `protobuf:"bytes,4,rep,name=detail_info,json=detailInfo,proto3" json:"detail_info,omitempty"`
	Logger        bool
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

// NewData creates a new instance of Data with options.
func NewData(options ...DataOption) *Data {
	data := &Data{}
	for _, option := range options {
		option(data)
	}
	return data
}

// WithEvent sets the Event field of Data.
func WithEvent(event CustomHookEvent) DataOption {
	return func(d *Data) {
		d.Event = event
	}
}

// WithParticipantID sets the ParticipantId field of Data.
func WithParticipantID(participantID string) DataOption {
	return func(d *Data) {
		d.ParticipantId = participantID
	}
}

// WithRoomID sets the RoomId field of Data.
func WithRoomID(roomID string) DataOption {
	return func(d *Data) {
		d.RoomId = roomID
	}
}

// WithDetailInfo sets the DetailInfo field of Data.
func WithDetailInfo(detailInfo []*KeyValue) DataOption {
	return func(d *Data) {
		d.DetailInfo = detailInfo
	}
}

// WithLogger sets the Logger field of Data.
func WithLogger(logger bool) DataOption {
	return func(d *Data) {
		d.Logger = logger
	}
}
