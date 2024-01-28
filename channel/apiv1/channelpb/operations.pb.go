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
// source: google/cloud/channel/v1/operations.proto

package channelpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RPCs that return a Long Running Operation.
type OperationMetadata_OperationType int32

const (
	// Not used.
	OperationMetadata_OPERATION_TYPE_UNSPECIFIED OperationMetadata_OperationType = 0
	// Long Running Operation was triggered by CreateEntitlement.
	OperationMetadata_CREATE_ENTITLEMENT OperationMetadata_OperationType = 1
	// Long Running Operation was triggered by ChangeRenewalSettings.
	OperationMetadata_CHANGE_RENEWAL_SETTINGS OperationMetadata_OperationType = 3
	// Long Running Operation was triggered by StartPaidService.
	OperationMetadata_START_PAID_SERVICE OperationMetadata_OperationType = 5
	// Long Running Operation was triggered by ActivateEntitlement.
	OperationMetadata_ACTIVATE_ENTITLEMENT OperationMetadata_OperationType = 7
	// Long Running Operation was triggered by SuspendEntitlement.
	OperationMetadata_SUSPEND_ENTITLEMENT OperationMetadata_OperationType = 8
	// Long Running Operation was triggered by CancelEntitlement.
	OperationMetadata_CANCEL_ENTITLEMENT OperationMetadata_OperationType = 9
	// Long Running Operation was triggered by TransferEntitlements.
	OperationMetadata_TRANSFER_ENTITLEMENTS OperationMetadata_OperationType = 10
	// Long Running Operation was triggered by TransferEntitlementsToGoogle.
	OperationMetadata_TRANSFER_ENTITLEMENTS_TO_GOOGLE OperationMetadata_OperationType = 11
	// Long Running Operation was triggered by ChangeOffer.
	OperationMetadata_CHANGE_OFFER OperationMetadata_OperationType = 14
	// Long Running Operation was triggered by ChangeParameters.
	OperationMetadata_CHANGE_PARAMETERS OperationMetadata_OperationType = 15
	// Long Running Operation was triggered by ProvisionCloudIdentity.
	OperationMetadata_PROVISION_CLOUD_IDENTITY OperationMetadata_OperationType = 16
)

// Enum value maps for OperationMetadata_OperationType.
var (
	OperationMetadata_OperationType_name = map[int32]string{
		0:  "OPERATION_TYPE_UNSPECIFIED",
		1:  "CREATE_ENTITLEMENT",
		3:  "CHANGE_RENEWAL_SETTINGS",
		5:  "START_PAID_SERVICE",
		7:  "ACTIVATE_ENTITLEMENT",
		8:  "SUSPEND_ENTITLEMENT",
		9:  "CANCEL_ENTITLEMENT",
		10: "TRANSFER_ENTITLEMENTS",
		11: "TRANSFER_ENTITLEMENTS_TO_GOOGLE",
		14: "CHANGE_OFFER",
		15: "CHANGE_PARAMETERS",
		16: "PROVISION_CLOUD_IDENTITY",
	}
	OperationMetadata_OperationType_value = map[string]int32{
		"OPERATION_TYPE_UNSPECIFIED":      0,
		"CREATE_ENTITLEMENT":              1,
		"CHANGE_RENEWAL_SETTINGS":         3,
		"START_PAID_SERVICE":              5,
		"ACTIVATE_ENTITLEMENT":            7,
		"SUSPEND_ENTITLEMENT":             8,
		"CANCEL_ENTITLEMENT":              9,
		"TRANSFER_ENTITLEMENTS":           10,
		"TRANSFER_ENTITLEMENTS_TO_GOOGLE": 11,
		"CHANGE_OFFER":                    14,
		"CHANGE_PARAMETERS":               15,
		"PROVISION_CLOUD_IDENTITY":        16,
	}
)

func (x OperationMetadata_OperationType) Enum() *OperationMetadata_OperationType {
	p := new(OperationMetadata_OperationType)
	*p = x
	return p
}

func (x OperationMetadata_OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationMetadata_OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_channel_v1_operations_proto_enumTypes[0].Descriptor()
}

func (OperationMetadata_OperationType) Type() protoreflect.EnumType {
	return &file_google_cloud_channel_v1_operations_proto_enumTypes[0]
}

func (x OperationMetadata_OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationMetadata_OperationType.Descriptor instead.
func (OperationMetadata_OperationType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_operations_proto_rawDescGZIP(), []int{0, 0}
}

// Provides contextual information about a
// [google.longrunning.Operation][google.longrunning.Operation].
type OperationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The RPC that initiated this Long Running Operation.
	OperationType OperationMetadata_OperationType `protobuf:"varint,1,opt,name=operation_type,json=operationType,proto3,enum=google.cloud.channel.v1.OperationMetadata_OperationType" json:"operation_type,omitempty"`
}

func (x *OperationMetadata) Reset() {
	*x = OperationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_channel_v1_operations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationMetadata) ProtoMessage() {}

func (x *OperationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_channel_v1_operations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationMetadata.ProtoReflect.Descriptor instead.
func (*OperationMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_operations_proto_rawDescGZIP(), []int{0}
}

func (x *OperationMetadata) GetOperationType() OperationMetadata_OperationType {
	if x != nil {
		return x.OperationType
	}
	return OperationMetadata_OPERATION_TYPE_UNSPECIFIED
}

var File_google_cloud_channel_v1_operations_proto protoreflect.FileDescriptor

var file_google_cloud_channel_v1_operations_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x22, 0xc5, 0x03, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5f, 0x0a, 0x0e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0xce, 0x02, 0x0a, 0x0d, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x4d, 0x45,
	0x4e, 0x54, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x52,
	0x45, 0x4e, 0x45, 0x57, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x53, 0x10,
	0x03, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x50, 0x41, 0x49, 0x44, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x5f, 0x45,
	0x4e, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12,
	0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x4d, 0x45,
	0x4e, 0x54, 0x10, 0x09, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52,
	0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x0a, 0x12,
	0x23, 0x0a, 0x1f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x45, 0x4e, 0x54, 0x49,
	0x54, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x54, 0x4f, 0x5f, 0x47, 0x4f, 0x4f, 0x47,
	0x4c, 0x45, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4f,
	0x46, 0x46, 0x45, 0x52, 0x10, 0x0e, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x53, 0x10, 0x0f, 0x12, 0x1c, 0x0a,
	0x18, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44,
	0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x10, 0x10, 0x42, 0x67, 0x0a, 0x1b, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x70, 0x62, 0x3b, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_channel_v1_operations_proto_rawDescOnce sync.Once
	file_google_cloud_channel_v1_operations_proto_rawDescData = file_google_cloud_channel_v1_operations_proto_rawDesc
)

func file_google_cloud_channel_v1_operations_proto_rawDescGZIP() []byte {
	file_google_cloud_channel_v1_operations_proto_rawDescOnce.Do(func() {
		file_google_cloud_channel_v1_operations_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_channel_v1_operations_proto_rawDescData)
	})
	return file_google_cloud_channel_v1_operations_proto_rawDescData
}

var file_google_cloud_channel_v1_operations_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_channel_v1_operations_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_channel_v1_operations_proto_goTypes = []interface{}{
	(OperationMetadata_OperationType)(0), // 0: google.cloud.channel.v1.OperationMetadata.OperationType
	(*OperationMetadata)(nil),            // 1: google.cloud.channel.v1.OperationMetadata
}
var file_google_cloud_channel_v1_operations_proto_depIdxs = []int32{
	0, // 0: google.cloud.channel.v1.OperationMetadata.operation_type:type_name -> google.cloud.channel.v1.OperationMetadata.OperationType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_channel_v1_operations_proto_init() }
func file_google_cloud_channel_v1_operations_proto_init() {
	if File_google_cloud_channel_v1_operations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_channel_v1_operations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationMetadata); i {
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
			RawDescriptor: file_google_cloud_channel_v1_operations_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_channel_v1_operations_proto_goTypes,
		DependencyIndexes: file_google_cloud_channel_v1_operations_proto_depIdxs,
		EnumInfos:         file_google_cloud_channel_v1_operations_proto_enumTypes,
		MessageInfos:      file_google_cloud_channel_v1_operations_proto_msgTypes,
	}.Build()
	File_google_cloud_channel_v1_operations_proto = out.File
	file_google_cloud_channel_v1_operations_proto_rawDesc = nil
	file_google_cloud_channel_v1_operations_proto_goTypes = nil
	file_google_cloud_channel_v1_operations_proto_depIdxs = nil
}
