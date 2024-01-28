// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.2
// source: google/cloud/support/v2/comment.proto

package supportpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A comment associated with a support case.
type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name for the comment.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The time when this comment was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The user or Google Support agent created this comment.
	Creator *Actor `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	// The full comment body. Maximum of 12800 characters. This can contain rich
	// text syntax.
	Body string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	// Output only. DEPRECATED. An automatically generated plain text version of
	// body with all rich text syntax stripped.
	PlainTextBody string `protobuf:"bytes,5,opt,name=plain_text_body,json=plainTextBody,proto3" json:"plain_text_body,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_support_v2_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_support_v2_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_google_cloud_support_v2_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Comment) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Comment) GetCreator() *Actor {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *Comment) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Comment) GetPlainTextBody() string {
	if x != nil {
		return x.PlainTextBody
	}
	return ""
}

var File_google_cloud_support_v2_comment_proto protoreflect.FileDescriptor

var file_google_cloud_support_v2_comment_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x82, 0x03, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x2b, 0x0a, 0x0f,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x54, 0x65, 0x78, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x3a, 0x9b, 0x01, 0xea, 0x41, 0x97, 0x01,
	0x0a, 0x23, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x73, 0x65, 0x7d,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x7d, 0x12, 0x32, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63,
	0x61, 0x73, 0x65, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x42, 0xb5, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32, 0x42, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x70, 0x62, 0x3b, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0xaa, 0x02,
	0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5c,
	0x56, 0x32, 0xea, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x3a, 0x56, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_support_v2_comment_proto_rawDescOnce sync.Once
	file_google_cloud_support_v2_comment_proto_rawDescData = file_google_cloud_support_v2_comment_proto_rawDesc
)

func file_google_cloud_support_v2_comment_proto_rawDescGZIP() []byte {
	file_google_cloud_support_v2_comment_proto_rawDescOnce.Do(func() {
		file_google_cloud_support_v2_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_support_v2_comment_proto_rawDescData)
	})
	return file_google_cloud_support_v2_comment_proto_rawDescData
}

var file_google_cloud_support_v2_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_support_v2_comment_proto_goTypes = []interface{}{
	(*Comment)(nil),               // 0: google.cloud.support.v2.Comment
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*Actor)(nil),                 // 2: google.cloud.support.v2.Actor
}
var file_google_cloud_support_v2_comment_proto_depIdxs = []int32{
	1, // 0: google.cloud.support.v2.Comment.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: google.cloud.support.v2.Comment.creator:type_name -> google.cloud.support.v2.Actor
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_support_v2_comment_proto_init() }
func file_google_cloud_support_v2_comment_proto_init() {
	if File_google_cloud_support_v2_comment_proto != nil {
		return
	}
	file_google_cloud_support_v2_actor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_support_v2_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
			RawDescriptor: file_google_cloud_support_v2_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_support_v2_comment_proto_goTypes,
		DependencyIndexes: file_google_cloud_support_v2_comment_proto_depIdxs,
		MessageInfos:      file_google_cloud_support_v2_comment_proto_msgTypes,
	}.Build()
	File_google_cloud_support_v2_comment_proto = out.File
	file_google_cloud_support_v2_comment_proto_rawDesc = nil
	file_google_cloud_support_v2_comment_proto_goTypes = nil
	file_google_cloud_support_v2_comment_proto_depIdxs = nil
}
