// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/qr_code.proto

package pb

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

type TraceLocationType int32

const (
	TraceLocationType_LOCATION_TYPE_UNSPECIFIED                       TraceLocationType = 0
	TraceLocationType_LOCATION_TYPE_PERMANENT_OTHER                   TraceLocationType = 1
	TraceLocationType_LOCATION_TYPE_TEMPORARY_OTHER                   TraceLocationType = 2
	TraceLocationType_LOCATION_TYPE_PERMANENT_RETAIL                  TraceLocationType = 3
	TraceLocationType_LOCATION_TYPE_PERMANENT_FOOD_SERVICE            TraceLocationType = 4
	TraceLocationType_LOCATION_TYPE_PERMANENT_CRAFT                   TraceLocationType = 5
	TraceLocationType_LOCATION_TYPE_PERMANENT_WORKPLACE               TraceLocationType = 6
	TraceLocationType_LOCATION_TYPE_PERMANENT_EDUCATIONAL_INSTITUTION TraceLocationType = 7
	TraceLocationType_LOCATION_TYPE_PERMANENT_PUBLIC_BUILDING         TraceLocationType = 8
	TraceLocationType_LOCATION_TYPE_TEMPORARY_CULTURAL_EVENT          TraceLocationType = 9
	TraceLocationType_LOCATION_TYPE_TEMPORARY_CLUB_ACTIVITY           TraceLocationType = 10
	TraceLocationType_LOCATION_TYPE_TEMPORARY_PRIVATE_EVENT           TraceLocationType = 11
	TraceLocationType_LOCATION_TYPE_TEMPORARY_WORSHIP_SERVICE         TraceLocationType = 12
)

// Enum value maps for TraceLocationType.
var (
	TraceLocationType_name = map[int32]string{
		0:  "LOCATION_TYPE_UNSPECIFIED",
		1:  "LOCATION_TYPE_PERMANENT_OTHER",
		2:  "LOCATION_TYPE_TEMPORARY_OTHER",
		3:  "LOCATION_TYPE_PERMANENT_RETAIL",
		4:  "LOCATION_TYPE_PERMANENT_FOOD_SERVICE",
		5:  "LOCATION_TYPE_PERMANENT_CRAFT",
		6:  "LOCATION_TYPE_PERMANENT_WORKPLACE",
		7:  "LOCATION_TYPE_PERMANENT_EDUCATIONAL_INSTITUTION",
		8:  "LOCATION_TYPE_PERMANENT_PUBLIC_BUILDING",
		9:  "LOCATION_TYPE_TEMPORARY_CULTURAL_EVENT",
		10: "LOCATION_TYPE_TEMPORARY_CLUB_ACTIVITY",
		11: "LOCATION_TYPE_TEMPORARY_PRIVATE_EVENT",
		12: "LOCATION_TYPE_TEMPORARY_WORSHIP_SERVICE",
	}
	TraceLocationType_value = map[string]int32{
		"LOCATION_TYPE_UNSPECIFIED":                       0,
		"LOCATION_TYPE_PERMANENT_OTHER":                   1,
		"LOCATION_TYPE_TEMPORARY_OTHER":                   2,
		"LOCATION_TYPE_PERMANENT_RETAIL":                  3,
		"LOCATION_TYPE_PERMANENT_FOOD_SERVICE":            4,
		"LOCATION_TYPE_PERMANENT_CRAFT":                   5,
		"LOCATION_TYPE_PERMANENT_WORKPLACE":               6,
		"LOCATION_TYPE_PERMANENT_EDUCATIONAL_INSTITUTION": 7,
		"LOCATION_TYPE_PERMANENT_PUBLIC_BUILDING":         8,
		"LOCATION_TYPE_TEMPORARY_CULTURAL_EVENT":          9,
		"LOCATION_TYPE_TEMPORARY_CLUB_ACTIVITY":           10,
		"LOCATION_TYPE_TEMPORARY_PRIVATE_EVENT":           11,
		"LOCATION_TYPE_TEMPORARY_WORSHIP_SERVICE":         12,
	}
)

func (x TraceLocationType) Enum() *TraceLocationType {
	p := new(TraceLocationType)
	*p = x
	return p
}

func (x TraceLocationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TraceLocationType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_qr_code_proto_enumTypes[0].Descriptor()
}

func (TraceLocationType) Type() protoreflect.EnumType {
	return &file_proto_qr_code_proto_enumTypes[0]
}

func (x TraceLocationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TraceLocationType.Descriptor instead.
func (TraceLocationType) EnumDescriptor() ([]byte, []int) {
	return file_proto_qr_code_proto_rawDescGZIP(), []int{0}
}

type QRCodePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version           uint32             `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	LocationData      *TraceLocation     `protobuf:"bytes,2,opt,name=locationData,proto3" json:"locationData,omitempty"`
	CrowdNotifierData *CrowdNotifierData `protobuf:"bytes,3,opt,name=crowdNotifierData,proto3" json:"crowdNotifierData,omitempty"`
	// byte sequence of CWALocationData
	VendorData []byte `protobuf:"bytes,4,opt,name=vendorData,proto3" json:"vendorData,omitempty"`
}

func (x *QRCodePayload) Reset() {
	*x = QRCodePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_qr_code_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QRCodePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QRCodePayload) ProtoMessage() {}

func (x *QRCodePayload) ProtoReflect() protoreflect.Message {
	mi := &file_proto_qr_code_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QRCodePayload.ProtoReflect.Descriptor instead.
func (*QRCodePayload) Descriptor() ([]byte, []int) {
	return file_proto_qr_code_proto_rawDescGZIP(), []int{0}
}

func (x *QRCodePayload) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *QRCodePayload) GetLocationData() *TraceLocation {
	if x != nil {
		return x.LocationData
	}
	return nil
}

func (x *QRCodePayload) GetCrowdNotifierData() *CrowdNotifierData {
	if x != nil {
		return x.CrowdNotifierData
	}
	return nil
}

func (x *QRCodePayload) GetVendorData() []byte {
	if x != nil {
		return x.VendorData
	}
	return nil
}

type TraceLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// max. 100 characters
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// max. 100 characters
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// UNIX timestamp (in seconds)
	StartTimestamp uint64 `protobuf:"varint,5,opt,name=startTimestamp,proto3" json:"startTimestamp,omitempty"`
	// UNIX timestamp (in seconds)
	EndTimestamp uint64 `protobuf:"varint,6,opt,name=endTimestamp,proto3" json:"endTimestamp,omitempty"`
}

func (x *TraceLocation) Reset() {
	*x = TraceLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_qr_code_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceLocation) ProtoMessage() {}

func (x *TraceLocation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_qr_code_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceLocation.ProtoReflect.Descriptor instead.
func (*TraceLocation) Descriptor() ([]byte, []int) {
	return file_proto_qr_code_proto_rawDescGZIP(), []int{1}
}

func (x *TraceLocation) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *TraceLocation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TraceLocation) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TraceLocation) GetStartTimestamp() uint64 {
	if x != nil {
		return x.StartTimestamp
	}
	return 0
}

func (x *TraceLocation) GetEndTimestamp() uint64 {
	if x != nil {
		return x.EndTimestamp
	}
	return 0
}

type CrowdNotifierData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version           uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	PublicKey         []byte `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	CryptographicSeed []byte `protobuf:"bytes,3,opt,name=cryptographicSeed,proto3" json:"cryptographicSeed,omitempty"`
}

func (x *CrowdNotifierData) Reset() {
	*x = CrowdNotifierData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_qr_code_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrowdNotifierData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrowdNotifierData) ProtoMessage() {}

func (x *CrowdNotifierData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_qr_code_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrowdNotifierData.ProtoReflect.Descriptor instead.
func (*CrowdNotifierData) Descriptor() ([]byte, []int) {
	return file_proto_qr_code_proto_rawDescGZIP(), []int{2}
}

func (x *CrowdNotifierData) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CrowdNotifierData) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *CrowdNotifierData) GetCryptographicSeed() []byte {
	if x != nil {
		return x.CryptographicSeed
	}
	return nil
}

type CWALocationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version                       uint32            `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Type                          TraceLocationType `protobuf:"varint,2,opt,name=type,proto3,enum=lukasmalkmus.cwa.v1.TraceLocationType" json:"type,omitempty"`
	DefaultCheckInLengthInMinutes uint32            `protobuf:"varint,3,opt,name=defaultCheckInLengthInMinutes,proto3" json:"defaultCheckInLengthInMinutes,omitempty"`
}

func (x *CWALocationData) Reset() {
	*x = CWALocationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_qr_code_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CWALocationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CWALocationData) ProtoMessage() {}

func (x *CWALocationData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_qr_code_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CWALocationData.ProtoReflect.Descriptor instead.
func (*CWALocationData) Descriptor() ([]byte, []int) {
	return file_proto_qr_code_proto_rawDescGZIP(), []int{3}
}

func (x *CWALocationData) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CWALocationData) GetType() TraceLocationType {
	if x != nil {
		return x.Type
	}
	return TraceLocationType_LOCATION_TYPE_UNSPECIFIED
}

func (x *CWALocationData) GetDefaultCheckInLengthInMinutes() uint32 {
	if x != nil {
		return x.DefaultCheckInLengthInMinutes
	}
	return 0
}

var File_proto_qr_code_proto protoreflect.FileDescriptor

var file_proto_qr_code_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6c, 0x75, 0x6b, 0x61, 0x73, 0x6d, 0x61, 0x6c, 0x6b,
	0x6d, 0x75, 0x73, 0x2e, 0x63, 0x77, 0x61, 0x2e, 0x76, 0x31, 0x22, 0xe7, 0x01, 0x0a, 0x0d, 0x51,
	0x52, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c,
	0x75, 0x6b, 0x61, 0x73, 0x6d, 0x61, 0x6c, 0x6b, 0x6d, 0x75, 0x73, 0x2e, 0x63, 0x77, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x54,
	0x0a, 0x11, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6c, 0x75, 0x6b, 0x61,
	0x73, 0x6d, 0x61, 0x6c, 0x6b, 0x6d, 0x75, 0x73, 0x2e, 0x63, 0x77, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x6f, 0x77, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x11, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x22, 0xb1, 0x01, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x79, 0x0a, 0x11, 0x43, 0x72, 0x6f, 0x77,
	0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x53, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x11, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x53,
	0x65, 0x65, 0x64, 0x22, 0xad, 0x01, 0x0a, 0x0f, 0x43, 0x57, 0x41, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x3a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x26, 0x2e, 0x6c, 0x75, 0x6b, 0x61, 0x73, 0x6d, 0x61, 0x6c, 0x6b, 0x6d, 0x75, 0x73, 0x2e, 0x63,
	0x77, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x44, 0x0a,
	0x1d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x49, 0x6e, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x1d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x49, 0x6e, 0x4d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x73, 0x2a, 0xa1, 0x04, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x63, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x4c, 0x4f, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x4c, 0x4f, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x41, 0x4e,
	0x45, 0x4e, 0x54, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x4c,
	0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x4d,
	0x50, 0x4f, 0x52, 0x41, 0x52, 0x59, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x02, 0x12, 0x22,
	0x0a, 0x1e, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x50, 0x45, 0x52, 0x4d, 0x41, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x54, 0x41, 0x49, 0x4c,
	0x10, 0x03, 0x12, 0x28, 0x0a, 0x24, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x41, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x4f,
	0x4f, 0x44, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x04, 0x12, 0x21, 0x0a, 0x1d,
	0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x45,
	0x52, 0x4d, 0x41, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x52, 0x41, 0x46, 0x54, 0x10, 0x05, 0x12,
	0x25, 0x0a, 0x21, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x50, 0x45, 0x52, 0x4d, 0x41, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x10, 0x06, 0x12, 0x33, 0x0a, 0x2f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x41, 0x4e, 0x45, 0x4e,
	0x54, 0x5f, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x49, 0x4e,
	0x53, 0x54, 0x49, 0x54, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x2b, 0x0a, 0x27, 0x4c,
	0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x45, 0x52,
	0x4d, 0x41, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x5f, 0x42, 0x55,
	0x49, 0x4c, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x08, 0x12, 0x2a, 0x0a, 0x26, 0x4c, 0x4f, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4f, 0x52,
	0x41, 0x52, 0x59, 0x5f, 0x43, 0x55, 0x4c, 0x54, 0x55, 0x52, 0x41, 0x4c, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x10, 0x09, 0x12, 0x29, 0x0a, 0x25, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4f, 0x52, 0x41, 0x52, 0x59, 0x5f,
	0x43, 0x4c, 0x55, 0x42, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x10, 0x0a, 0x12,
	0x29, 0x0a, 0x25, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4f, 0x52, 0x41, 0x52, 0x59, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41,
	0x54, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x0b, 0x12, 0x2b, 0x0a, 0x27, 0x4c, 0x4f,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x4d, 0x50,
	0x4f, 0x52, 0x41, 0x52, 0x59, 0x5f, 0x57, 0x4f, 0x52, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x0c, 0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_qr_code_proto_rawDescOnce sync.Once
	file_proto_qr_code_proto_rawDescData = file_proto_qr_code_proto_rawDesc
)

func file_proto_qr_code_proto_rawDescGZIP() []byte {
	file_proto_qr_code_proto_rawDescOnce.Do(func() {
		file_proto_qr_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_qr_code_proto_rawDescData)
	})
	return file_proto_qr_code_proto_rawDescData
}

var file_proto_qr_code_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_qr_code_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_qr_code_proto_goTypes = []interface{}{
	(TraceLocationType)(0),    // 0: lukasmalkmus.cwa.v1.TraceLocationType
	(*QRCodePayload)(nil),     // 1: lukasmalkmus.cwa.v1.QRCodePayload
	(*TraceLocation)(nil),     // 2: lukasmalkmus.cwa.v1.TraceLocation
	(*CrowdNotifierData)(nil), // 3: lukasmalkmus.cwa.v1.CrowdNotifierData
	(*CWALocationData)(nil),   // 4: lukasmalkmus.cwa.v1.CWALocationData
}
var file_proto_qr_code_proto_depIdxs = []int32{
	2, // 0: lukasmalkmus.cwa.v1.QRCodePayload.locationData:type_name -> lukasmalkmus.cwa.v1.TraceLocation
	3, // 1: lukasmalkmus.cwa.v1.QRCodePayload.crowdNotifierData:type_name -> lukasmalkmus.cwa.v1.CrowdNotifierData
	0, // 2: lukasmalkmus.cwa.v1.CWALocationData.type:type_name -> lukasmalkmus.cwa.v1.TraceLocationType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_qr_code_proto_init() }
func file_proto_qr_code_proto_init() {
	if File_proto_qr_code_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_qr_code_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QRCodePayload); i {
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
		file_proto_qr_code_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceLocation); i {
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
		file_proto_qr_code_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrowdNotifierData); i {
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
		file_proto_qr_code_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CWALocationData); i {
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
			RawDescriptor: file_proto_qr_code_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_qr_code_proto_goTypes,
		DependencyIndexes: file_proto_qr_code_proto_depIdxs,
		EnumInfos:         file_proto_qr_code_proto_enumTypes,
		MessageInfos:      file_proto_qr_code_proto_msgTypes,
	}.Build()
	File_proto_qr_code_proto = out.File
	file_proto_qr_code_proto_rawDesc = nil
	file_proto_qr_code_proto_goTypes = nil
	file_proto_qr_code_proto_depIdxs = nil
}