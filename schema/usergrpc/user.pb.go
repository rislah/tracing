// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: schema/user.proto

package usergrpc

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

type GetUserByNameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetUserByNameReq) Reset() {
	*x = GetUserByNameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByNameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByNameReq) ProtoMessage() {}

func (x *GetUserByNameReq) ProtoReflect() protoreflect.Message {
	mi := &file_schema_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByNameReq.ProtoReflect.Descriptor instead.
func (*GetUserByNameReq) Descriptor() ([]byte, []int) {
	return file_schema_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserByNameReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username     string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	PasswordHash string `protobuf:"bytes,3,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_schema_user_proto_msgTypes[1]
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
	return file_schema_user_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

type GetUserByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserByNameResponse) Reset() {
	*x = GetUserByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByNameResponse) ProtoMessage() {}

func (x *GetUserByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByNameResponse.ProtoReflect.Descriptor instead.
func (*GetUserByNameResponse) Descriptor() ([]byte, []int) {
	return file_schema_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserByNameResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_schema_user_proto protoreflect.FileDescriptor

var file_schema_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x74, 0x6f, 0x79, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x57, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x22, 0x36, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x74, 0x6f, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x32, 0x51, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x42, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x15, 0x2e, 0x74, 0x6f, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x74, 0x6f, 0x79, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_user_proto_rawDescOnce sync.Once
	file_schema_user_proto_rawDescData = file_schema_user_proto_rawDesc
)

func file_schema_user_proto_rawDescGZIP() []byte {
	file_schema_user_proto_rawDescOnce.Do(func() {
		file_schema_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_user_proto_rawDescData)
	})
	return file_schema_user_proto_rawDescData
}

var file_schema_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_schema_user_proto_goTypes = []interface{}{
	(*GetUserByNameReq)(nil),      // 0: toy.GetUserByNameReq
	(*User)(nil),                  // 1: toy.User
	(*GetUserByNameResponse)(nil), // 2: toy.GetUserByNameResponse
}
var file_schema_user_proto_depIdxs = []int32{
	1, // 0: toy.GetUserByNameResponse.user:type_name -> toy.User
	0, // 1: toy.UserService.GetUserByName:input_type -> toy.GetUserByNameReq
	2, // 2: toy.UserService.GetUserByName:output_type -> toy.GetUserByNameResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_schema_user_proto_init() }
func file_schema_user_proto_init() {
	if File_schema_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByNameReq); i {
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
		file_schema_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_schema_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByNameResponse); i {
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
			RawDescriptor: file_schema_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schema_user_proto_goTypes,
		DependencyIndexes: file_schema_user_proto_depIdxs,
		MessageInfos:      file_schema_user_proto_msgTypes,
	}.Build()
	File_schema_user_proto = out.File
	file_schema_user_proto_rawDesc = nil
	file_schema_user_proto_goTypes = nil
	file_schema_user_proto_depIdxs = nil
}
