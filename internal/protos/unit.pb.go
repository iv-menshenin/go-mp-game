// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: internal/protos/unit.proto

package protos

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

type Action int32

const (
	Action_idle Action = 0
	Action_run  Action = 1
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "idle",
		1: "run",
	}
	Action_value = map[string]int32{
		"idle": 0,
		"run":  1,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_protos_unit_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_internal_protos_unit_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_internal_protos_unit_proto_rawDescGZIP(), []int{0}
}

type Skin int32

const (
	Skin_big_demon  Skin = 0
	Skin_big_zombie Skin = 1
	Skin_elf_f      Skin = 2
)

// Enum value maps for Skin.
var (
	Skin_name = map[int32]string{
		0: "big_demon",
		1: "big_zombie",
		2: "elf_f",
	}
	Skin_value = map[string]int32{
		"big_demon":  0,
		"big_zombie": 1,
		"elf_f":      2,
	}
)

func (x Skin) Enum() *Skin {
	p := new(Skin)
	*p = x
	return p
}

func (x Skin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Skin) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_protos_unit_proto_enumTypes[1].Descriptor()
}

func (Skin) Type() protoreflect.EnumType {
	return &file_internal_protos_unit_proto_enumTypes[1]
}

func (x Skin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Skin.Descriptor instead.
func (Skin) EnumDescriptor() ([]byte, []int) {
	return file_internal_protos_unit_proto_rawDescGZIP(), []int{1}
}

type Unit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Position *Position `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Frame    int32     `protobuf:"varint,3,opt,name=frame,proto3" json:"frame,omitempty"`
	Skin     Skin      `protobuf:"varint,4,opt,name=skin,proto3,enum=protos.Skin" json:"skin,omitempty"`
	Action   Action    `protobuf:"varint,5,opt,name=action,proto3,enum=protos.Action" json:"action,omitempty"`
	Velocity *Velocity `protobuf:"bytes,6,opt,name=velocity,proto3" json:"velocity,omitempty"`
	Side     Direction `protobuf:"varint,7,opt,name=side,proto3,enum=protos.Direction" json:"side,omitempty"`
	Hp       uint32    `protobuf:"varint,8,opt,name=hp,proto3" json:"hp,omitempty"`
}

func (x *Unit) Reset() {
	*x = Unit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_unit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit) ProtoMessage() {}

func (x *Unit) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_unit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit.ProtoReflect.Descriptor instead.
func (*Unit) Descriptor() ([]byte, []int) {
	return file_internal_protos_unit_proto_rawDescGZIP(), []int{0}
}

func (x *Unit) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Unit) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Unit) GetFrame() int32 {
	if x != nil {
		return x.Frame
	}
	return 0
}

func (x *Unit) GetSkin() Skin {
	if x != nil {
		return x.Skin
	}
	return Skin_big_demon
}

func (x *Unit) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_idle
}

func (x *Unit) GetVelocity() *Velocity {
	if x != nil {
		return x.Velocity
	}
	return nil
}

func (x *Unit) GetSide() Direction {
	if x != nil {
		return x.Side
	}
	return Direction_left
}

func (x *Unit) GetHp() uint32 {
	if x != nil {
		return x.Hp
	}
	return 0
}

type PatchUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Position *Position  `protobuf:"bytes,2,opt,name=position,proto3,oneof" json:"position,omitempty"`
	Action   *Action    `protobuf:"varint,5,opt,name=action,proto3,enum=protos.Action,oneof" json:"action,omitempty"`
	Velocity *Velocity  `protobuf:"bytes,6,opt,name=velocity,proto3,oneof" json:"velocity,omitempty"`
	Side     *Direction `protobuf:"varint,7,opt,name=side,proto3,enum=protos.Direction,oneof" json:"side,omitempty"`
	Hp       *uint32    `protobuf:"varint,8,opt,name=hp,proto3,oneof" json:"hp,omitempty"`
}

func (x *PatchUnit) Reset() {
	*x = PatchUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_unit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchUnit) ProtoMessage() {}

func (x *PatchUnit) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_unit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchUnit.ProtoReflect.Descriptor instead.
func (*PatchUnit) Descriptor() ([]byte, []int) {
	return file_internal_protos_unit_proto_rawDescGZIP(), []int{1}
}

func (x *PatchUnit) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PatchUnit) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *PatchUnit) GetAction() Action {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return Action_idle
}

func (x *PatchUnit) GetVelocity() *Velocity {
	if x != nil {
		return x.Velocity
	}
	return nil
}

func (x *PatchUnit) GetSide() Direction {
	if x != nil && x.Side != nil {
		return *x.Side
	}
	return Direction_left
}

func (x *PatchUnit) GetHp() uint32 {
	if x != nil && x.Hp != nil {
		return *x.Hp
	}
	return 0
}

var File_internal_protos_unit_proto protoreflect.FileDescriptor

var file_internal_protos_unit_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x89, 0x02, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x04, 0x73, 0x6b, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x6b, 0x69, 0x6e, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x6e,
	0x12, 0x26, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x08, 0x76, 0x65, 0x6c, 0x6f,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x52, 0x08, 0x76, 0x65,
	0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x68, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x68, 0x70, 0x22, 0xa4, 0x02,
	0x0a, 0x09, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2b,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x08, 0x76,
	0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x48,
	0x02, 0x52, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x2a,
	0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x68, 0x70,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x04, 0x52, 0x02, 0x68, 0x70, 0x88, 0x01, 0x01, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x76, 0x65, 0x6c, 0x6f,
	0x63, 0x69, 0x74, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x42, 0x05, 0x0a,
	0x03, 0x5f, 0x68, 0x70, 0x2a, 0x1b, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08,
	0x0a, 0x04, 0x69, 0x64, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x72, 0x75, 0x6e, 0x10,
	0x01, 0x2a, 0x30, 0x0a, 0x04, 0x53, 0x6b, 0x69, 0x6e, 0x12, 0x0d, 0x0a, 0x09, 0x62, 0x69, 0x67,
	0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x62, 0x69, 0x67, 0x5f,
	0x7a, 0x6f, 0x6d, 0x62, 0x69, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x65, 0x6c, 0x66, 0x5f,
	0x66, 0x10, 0x02, 0x42, 0x11, 0x5a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_protos_unit_proto_rawDescOnce sync.Once
	file_internal_protos_unit_proto_rawDescData = file_internal_protos_unit_proto_rawDesc
)

func file_internal_protos_unit_proto_rawDescGZIP() []byte {
	file_internal_protos_unit_proto_rawDescOnce.Do(func() {
		file_internal_protos_unit_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_protos_unit_proto_rawDescData)
	})
	return file_internal_protos_unit_proto_rawDescData
}

var file_internal_protos_unit_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_internal_protos_unit_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_protos_unit_proto_goTypes = []any{
	(Action)(0),       // 0: protos.Action
	(Skin)(0),         // 1: protos.Skin
	(*Unit)(nil),      // 2: protos.Unit
	(*PatchUnit)(nil), // 3: protos.PatchUnit
	(*Position)(nil),  // 4: protos.Position
	(*Velocity)(nil),  // 5: protos.Velocity
	(Direction)(0),    // 6: protos.Direction
}
var file_internal_protos_unit_proto_depIdxs = []int32{
	4, // 0: protos.Unit.position:type_name -> protos.Position
	1, // 1: protos.Unit.skin:type_name -> protos.Skin
	0, // 2: protos.Unit.action:type_name -> protos.Action
	5, // 3: protos.Unit.velocity:type_name -> protos.Velocity
	6, // 4: protos.Unit.side:type_name -> protos.Direction
	4, // 5: protos.PatchUnit.position:type_name -> protos.Position
	0, // 6: protos.PatchUnit.action:type_name -> protos.Action
	5, // 7: protos.PatchUnit.velocity:type_name -> protos.Velocity
	6, // 8: protos.PatchUnit.side:type_name -> protos.Direction
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_internal_protos_unit_proto_init() }
func file_internal_protos_unit_proto_init() {
	if File_internal_protos_unit_proto != nil {
		return
	}
	file_internal_protos_utils_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_protos_unit_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Unit); i {
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
		file_internal_protos_unit_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PatchUnit); i {
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
	file_internal_protos_unit_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_protos_unit_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_protos_unit_proto_goTypes,
		DependencyIndexes: file_internal_protos_unit_proto_depIdxs,
		EnumInfos:         file_internal_protos_unit_proto_enumTypes,
		MessageInfos:      file_internal_protos_unit_proto_msgTypes,
	}.Build()
	File_internal_protos_unit_proto = out.File
	file_internal_protos_unit_proto_rawDesc = nil
	file_internal_protos_unit_proto_goTypes = nil
	file_internal_protos_unit_proto_depIdxs = nil
}
