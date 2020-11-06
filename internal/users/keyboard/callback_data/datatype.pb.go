// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/users/keyboard/callback_data/datatype.proto

package callback_data

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

type KeyboardType int32

const (
	KeyboardType_NULL          KeyboardType = 0
	KeyboardType_CONNECTION    KeyboardType = 1
	KeyboardType_SUBSCRIPTIONS KeyboardType = 2
	KeyboardType_PUBLISH       KeyboardType = 3
	KeyboardType_BUTTONS       KeyboardType = 4
	KeyboardType_COMMAND       KeyboardType = 5
	KeyboardType_CHART         KeyboardType = 6
)

var KeyboardType_name = map[int32]string{
	0: "NULL",
	1: "CONNECTION",
	2: "SUBSCRIPTIONS",
	3: "PUBLISH",
	4: "BUTTONS",
	5: "COMMAND",
	6: "CHART",
}

var KeyboardType_value = map[string]int32{
	"NULL":          0,
	"CONNECTION":    1,
	"SUBSCRIPTIONS": 2,
	"PUBLISH":       3,
	"BUTTONS":       4,
	"COMMAND":       5,
	"CHART":         6,
}

func (x KeyboardType) String() string {
	return proto.EnumName(KeyboardType_name, int32(x))
}

func (KeyboardType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_13504e582ceba832, []int{0}
}

type ActionType int32

const (
	ActionType_EMPTY                               ActionType = 0
	ActionType_EDIT                                ActionType = 1
	ActionType_DELETE                              ActionType = 2
	ActionType_SWITCH_QOS                          ActionType = 3
	ActionType_SWITCH_RETAINED                     ActionType = 4
	ActionType_SWITCH_SUB_DATA_TYPE                ActionType = 5
	ActionType_BEFORE_VALUE_TEXT                   ActionType = 6
	ActionType_AFTER_VALUE_TEXT                    ActionType = 7
	ActionType_BACK_TO_MENU                        ActionType = 8
	ActionType_BACK_TO_LIST                        ActionType = 9
	ActionType_SWITCH_BUTTON_TYPE                  ActionType = 10
	ActionType_ADD_BUTTON                          ActionType = 11
	ActionType_ADD_SUBSCRIPTION                    ActionType = 12
	ActionType_ADD_COMMAND                         ActionType = 13
	ActionType_SWITCH_SUBSCRIPTION_TYPE            ActionType = 14
	ActionType_SWITCH_ADDED_SUBSCRIPTION_TYPE      ActionType = 15
	ActionType_SWITCH_ADDED_SUBSCRIPTION_QOS       ActionType = 16
	ActionType_SWITCH_ADDED_SUBSCRIPTION_DATA_TYPE ActionType = 17
	ActionType_EDIT_COMMAND_NAME                   ActionType = 18
	ActionType_EDIT_COMMAND_TOPIC                  ActionType = 19
	ActionType_EDIT_COMMAND_VALUE                  ActionType = 20
	ActionType_DELETE_COMMAND                      ActionType = 21
	ActionType_EDIT_COMMAND                        ActionType = 22
	ActionType_SWITCH_SUBSCRIPTION                 ActionType = 23
	ActionType_DELETE_SUB_CHART                    ActionType = 24
	ActionType_ADD_SUB_CHART                       ActionType = 25
)

var ActionType_name = map[int32]string{
	0:  "EMPTY",
	1:  "EDIT",
	2:  "DELETE",
	3:  "SWITCH_QOS",
	4:  "SWITCH_RETAINED",
	5:  "SWITCH_SUB_DATA_TYPE",
	6:  "BEFORE_VALUE_TEXT",
	7:  "AFTER_VALUE_TEXT",
	8:  "BACK_TO_MENU",
	9:  "BACK_TO_LIST",
	10: "SWITCH_BUTTON_TYPE",
	11: "ADD_BUTTON",
	12: "ADD_SUBSCRIPTION",
	13: "ADD_COMMAND",
	14: "SWITCH_SUBSCRIPTION_TYPE",
	15: "SWITCH_ADDED_SUBSCRIPTION_TYPE",
	16: "SWITCH_ADDED_SUBSCRIPTION_QOS",
	17: "SWITCH_ADDED_SUBSCRIPTION_DATA_TYPE",
	18: "EDIT_COMMAND_NAME",
	19: "EDIT_COMMAND_TOPIC",
	20: "EDIT_COMMAND_VALUE",
	21: "DELETE_COMMAND",
	22: "EDIT_COMMAND",
	23: "SWITCH_SUBSCRIPTION",
	24: "DELETE_SUB_CHART",
	25: "ADD_SUB_CHART",
}

var ActionType_value = map[string]int32{
	"EMPTY":                               0,
	"EDIT":                                1,
	"DELETE":                              2,
	"SWITCH_QOS":                          3,
	"SWITCH_RETAINED":                     4,
	"SWITCH_SUB_DATA_TYPE":                5,
	"BEFORE_VALUE_TEXT":                   6,
	"AFTER_VALUE_TEXT":                    7,
	"BACK_TO_MENU":                        8,
	"BACK_TO_LIST":                        9,
	"SWITCH_BUTTON_TYPE":                  10,
	"ADD_BUTTON":                          11,
	"ADD_SUBSCRIPTION":                    12,
	"ADD_COMMAND":                         13,
	"SWITCH_SUBSCRIPTION_TYPE":            14,
	"SWITCH_ADDED_SUBSCRIPTION_TYPE":      15,
	"SWITCH_ADDED_SUBSCRIPTION_QOS":       16,
	"SWITCH_ADDED_SUBSCRIPTION_DATA_TYPE": 17,
	"EDIT_COMMAND_NAME":                   18,
	"EDIT_COMMAND_TOPIC":                  19,
	"EDIT_COMMAND_VALUE":                  20,
	"DELETE_COMMAND":                      21,
	"EDIT_COMMAND":                        22,
	"SWITCH_SUBSCRIPTION":                 23,
	"DELETE_SUB_CHART":                    24,
	"ADD_SUB_CHART":                       25,
}

func (x ActionType) String() string {
	return proto.EnumName(ActionType_name, int32(x))
}

func (ActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_13504e582ceba832, []int{1}
}

type QueryDataType struct {
	MessageId            int64        `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Keyboard             KeyboardType `protobuf:"varint,2,opt,name=keyboard,proto3,enum=callback_data.KeyboardType" json:"keyboard,omitempty"`
	Path                 []int32      `protobuf:"varint,3,rep,packed,name=path,proto3" json:"path,omitempty"`
	Action               ActionType   `protobuf:"varint,4,opt,name=action,proto3,enum=callback_data.ActionType" json:"action,omitempty"`
	IntValue             int32        `protobuf:"varint,6,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	BoolValue            bool         `protobuf:"varint,7,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	Index                int32        `protobuf:"varint,8,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *QueryDataType) Reset()         { *m = QueryDataType{} }
func (m *QueryDataType) String() string { return proto.CompactTextString(m) }
func (*QueryDataType) ProtoMessage()    {}
func (*QueryDataType) Descriptor() ([]byte, []int) {
	return fileDescriptor_13504e582ceba832, []int{0}
}

func (m *QueryDataType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryDataType.Unmarshal(m, b)
}
func (m *QueryDataType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryDataType.Marshal(b, m, deterministic)
}
func (m *QueryDataType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDataType.Merge(m, src)
}
func (m *QueryDataType) XXX_Size() int {
	return xxx_messageInfo_QueryDataType.Size(m)
}
func (m *QueryDataType) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDataType.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDataType proto.InternalMessageInfo

func (m *QueryDataType) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *QueryDataType) GetKeyboard() KeyboardType {
	if m != nil {
		return m.Keyboard
	}
	return KeyboardType_NULL
}

func (m *QueryDataType) GetPath() []int32 {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *QueryDataType) GetAction() ActionType {
	if m != nil {
		return m.Action
	}
	return ActionType_EMPTY
}

func (m *QueryDataType) GetIntValue() int32 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *QueryDataType) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *QueryDataType) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterEnum("callback_data.KeyboardType", KeyboardType_name, KeyboardType_value)
	proto.RegisterEnum("callback_data.ActionType", ActionType_name, ActionType_value)
	proto.RegisterType((*QueryDataType)(nil), "callback_data.QueryDataType")
}

func init() {
	proto.RegisterFile("internal/users/keyboard/callback_data/datatype.proto", fileDescriptor_13504e582ceba832)
}

var fileDescriptor_13504e582ceba832 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4f, 0x4f, 0xdb, 0x4e,
	0x10, 0xc5, 0xf9, 0x47, 0x32, 0x90, 0x30, 0x0c, 0x01, 0x8c, 0xf8, 0xf1, 0x53, 0x4a, 0x0f, 0x8d,
	0x38, 0x10, 0xf5, 0x8f, 0xd4, 0xb3, 0xe3, 0x5d, 0x84, 0x45, 0x62, 0x07, 0x7b, 0x4d, 0xcb, 0xc9,
	0xda, 0x10, 0xab, 0x8d, 0x08, 0x71, 0x94, 0x98, 0xaa, 0xb9, 0xf5, 0xeb, 0xf5, 0x5b, 0x55, 0x6b,
	0x2f, 0x10, 0x5a, 0xb8, 0x44, 0x99, 0xf7, 0xde, 0xce, 0xbc, 0x79, 0xbb, 0x86, 0x4f, 0xe3, 0x69,
	0x1a, 0xcf, 0xa7, 0x72, 0xd2, 0xb9, 0x5f, 0xc4, 0xf3, 0x45, 0xe7, 0x36, 0x5e, 0x0e, 0x13, 0x39,
	0x1f, 0x75, 0x6e, 0xe4, 0x64, 0x32, 0x94, 0x37, 0xb7, 0xd1, 0x48, 0xa6, 0xb2, 0xa3, 0x7e, 0xd2,
	0xe5, 0x2c, 0x3e, 0x9d, 0xcd, 0x93, 0x34, 0xa1, 0xfa, 0x33, 0xf6, 0xf8, 0x57, 0x01, 0xea, 0x97,
	0xf7, 0xf1, 0x7c, 0xc9, 0x64, 0x2a, 0xc5, 0x72, 0x16, 0xd3, 0x11, 0xc0, 0x5d, 0xbc, 0x58, 0xc8,
	0x6f, 0x71, 0x34, 0x1e, 0x99, 0x46, 0xcb, 0x68, 0x17, 0xfd, 0x9a, 0x46, 0x9c, 0x11, 0x7d, 0x86,
	0xea, 0xc3, 0x20, 0xb3, 0xd0, 0x32, 0xda, 0x8d, 0x0f, 0x87, 0xa7, 0xcf, 0x5a, 0x9e, 0x5e, 0x68,
	0x5a, 0x75, 0xf3, 0x1f, 0xc5, 0x44, 0x50, 0x9a, 0xc9, 0xf4, 0xbb, 0x59, 0x6c, 0x15, 0xdb, 0x65,
	0x3f, 0xfb, 0x4f, 0xef, 0xa1, 0x22, 0x6f, 0xd2, 0x71, 0x32, 0x35, 0x4b, 0x59, 0xab, 0x83, 0xbf,
	0x5a, 0x59, 0x19, 0x99, 0x35, 0xd2, 0x42, 0x3a, 0x84, 0xda, 0x78, 0x9a, 0x46, 0x3f, 0xe4, 0xe4,
	0x3e, 0x36, 0x2b, 0x2d, 0xa3, 0x5d, 0xf6, 0xab, 0xe3, 0x69, 0x7a, 0xa5, 0x6a, 0xe5, 0x7d, 0x98,
	0x24, 0x13, 0xcd, 0xae, 0xb7, 0x8c, 0x76, 0xd5, 0xaf, 0x29, 0x24, 0xa7, 0x9b, 0x50, 0x1e, 0x4f,
	0x47, 0xf1, 0x4f, 0xb3, 0x9a, 0x9d, 0xcb, 0x8b, 0x93, 0x3b, 0xd8, 0x5c, 0xb5, 0x4c, 0x55, 0x28,
	0xb9, 0x61, 0xaf, 0x87, 0x6b, 0xd4, 0x00, 0xb0, 0x3d, 0xd7, 0xe5, 0xb6, 0x70, 0x3c, 0x17, 0x0d,
	0xda, 0x86, 0x7a, 0x10, 0x76, 0x03, 0xdb, 0x77, 0x06, 0x0a, 0x09, 0xb0, 0x40, 0x1b, 0xb0, 0x3e,
	0x08, 0xbb, 0x3d, 0x27, 0x38, 0xc7, 0xa2, 0x2a, 0xba, 0xa1, 0x10, 0x8a, 0x29, 0xa9, 0xc2, 0xf6,
	0xfa, 0x7d, 0xcb, 0x65, 0x58, 0xa6, 0x1a, 0x94, 0xed, 0x73, 0xcb, 0x17, 0x58, 0x39, 0xf9, 0x5d,
	0x02, 0x78, 0xda, 0x4b, 0x31, 0xbc, 0x3f, 0x10, 0xd7, 0xb8, 0xa6, 0x06, 0x73, 0xe6, 0x08, 0x34,
	0x08, 0xa0, 0xc2, 0x78, 0x8f, 0x0b, 0x8e, 0x05, 0x65, 0x22, 0xf8, 0xe2, 0x08, 0xfb, 0x3c, 0xba,
	0xf4, 0x02, 0x2c, 0xd2, 0x0e, 0x6c, 0xe9, 0xda, 0xe7, 0xc2, 0x72, 0x5c, 0xce, 0xb0, 0x44, 0x26,
	0x34, 0x35, 0x18, 0x84, 0xdd, 0x88, 0x59, 0xc2, 0x8a, 0xc4, 0xf5, 0x80, 0x63, 0x99, 0x76, 0x61,
	0xbb, 0xcb, 0xcf, 0x3c, 0x9f, 0x47, 0x57, 0x56, 0x2f, 0xe4, 0x91, 0xe0, 0x5f, 0x05, 0x56, 0xa8,
	0x09, 0x68, 0x9d, 0x09, 0xee, 0xaf, 0xa2, 0xeb, 0x84, 0xb0, 0xd9, 0xb5, 0xec, 0x8b, 0x48, 0x78,
	0x51, 0x9f, 0xbb, 0x21, 0x56, 0x57, 0x91, 0x9e, 0x13, 0x08, 0xac, 0xd1, 0x1e, 0x90, 0x1e, 0x95,
	0xef, 0x9a, 0x0f, 0x02, 0xe5, 0xd3, 0x62, 0x4c, 0x83, 0xb8, 0x91, 0x4d, 0x60, 0x2c, 0x5a, 0x0d,
	0x0c, 0x37, 0x69, 0x0b, 0x36, 0x14, 0xfa, 0x90, 0x4c, 0x9d, 0xfe, 0x03, 0xf3, 0xc9, 0xf9, 0xa3,
	0x32, 0x6f, 0xda, 0xa0, 0x63, 0xf8, 0x5f, 0xb3, 0x16, 0x63, 0x9c, 0xbd, 0xa0, 0xd9, 0xa2, 0x37,
	0x70, 0xf4, 0xba, 0x46, 0x65, 0x86, 0xf4, 0x0e, 0xde, 0xbe, 0x2e, 0x79, 0x4a, 0x6b, 0x5b, 0xa5,
	0xa5, 0xae, 0xe0, 0xc1, 0x5f, 0xe4, 0x5a, 0x7d, 0x8e, 0xa4, 0x76, 0x7e, 0x06, 0x0b, 0x6f, 0xe0,
	0xd8, 0xb8, 0xf3, 0x0f, 0x9e, 0x85, 0x89, 0x4d, 0x22, 0x68, 0xe4, 0xf7, 0xf7, 0xb8, 0xe8, 0xae,
	0x4a, 0x72, 0x55, 0x8b, 0x7b, 0xb4, 0x0f, 0x3b, 0x2f, 0xac, 0x8e, 0xfb, 0x2a, 0x3a, 0x7d, 0x5c,
	0xdd, 0x66, 0xfe, 0x70, 0x4c, 0xf5, 0xfa, 0x74, 0xa0, 0x1a, 0x3a, 0x18, 0x56, 0xb2, 0x6f, 0xfa,
	0xe3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0xdc, 0x66, 0xf4, 0x0b, 0x04, 0x00, 0x00,
}
