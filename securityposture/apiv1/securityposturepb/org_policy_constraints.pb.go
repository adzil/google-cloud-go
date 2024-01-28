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
// source: google/cloud/securityposture/v1/org_policy_constraints.proto

package securityposturepb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Message for Org Policy Canned Constraint.
type OrgPolicyConstraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Org Policy Canned Constraint id.
	CannedConstraintId string `protobuf:"bytes,1,opt,name=canned_constraint_id,json=cannedConstraintId,proto3" json:"canned_constraint_id,omitempty"`
	// Required. Org PolicySpec rules.
	PolicyRules []*PolicyRule `protobuf:"bytes,2,rep,name=policy_rules,json=policyRules,proto3" json:"policy_rules,omitempty"`
}

func (x *OrgPolicyConstraint) Reset() {
	*x = OrgPolicyConstraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securityposture_v1_org_policy_constraints_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgPolicyConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgPolicyConstraint) ProtoMessage() {}

func (x *OrgPolicyConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securityposture_v1_org_policy_constraints_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgPolicyConstraint.ProtoReflect.Descriptor instead.
func (*OrgPolicyConstraint) Descriptor() ([]byte, []int) {
	return file_google_cloud_securityposture_v1_org_policy_constraints_proto_rawDescGZIP(), []int{0}
}

func (x *OrgPolicyConstraint) GetCannedConstraintId() string {
	if x != nil {
		return x.CannedConstraintId
	}
	return ""
}

func (x *OrgPolicyConstraint) GetPolicyRules() []*PolicyRule {
	if x != nil {
		return x.PolicyRules
	}
	return nil
}

// Message for Org Policy Custom Constraint.
type OrgPolicyConstraintCustom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Org Policy Custom Constraint.
	CustomConstraint *CustomConstraint `protobuf:"bytes,1,opt,name=custom_constraint,json=customConstraint,proto3" json:"custom_constraint,omitempty"`
	// Required. Org Policyspec rules.
	PolicyRules []*PolicyRule `protobuf:"bytes,2,rep,name=policy_rules,json=policyRules,proto3" json:"policy_rules,omitempty"`
}

func (x *OrgPolicyConstraintCustom) Reset() {
	*x = OrgPolicyConstraintCustom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securityposture_v1_org_policy_constraints_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgPolicyConstraintCustom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgPolicyConstraintCustom) ProtoMessage() {}

func (x *OrgPolicyConstraintCustom) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securityposture_v1_org_policy_constraints_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgPolicyConstraintCustom.ProtoReflect.Descriptor instead.
func (*OrgPolicyConstraintCustom) Descriptor() ([]byte, []int) {
	return file_google_cloud_securityposture_v1_org_policy_constraints_proto_rawDescGZIP(), []int{1}
}

func (x *OrgPolicyConstraintCustom) GetCustomConstraint() *CustomConstraint {
	if x != nil {
		return x.CustomConstraint
	}
	return nil
}

func (x *OrgPolicyConstraintCustom) GetPolicyRules() []*PolicyRule {
	if x != nil {
		return x.PolicyRules
	}
	return nil
}

var File_google_cloud_securityposture_v1_org_policy_constraints_proto protoreflect.FileDescriptor

var file_google_cloud_securityposture_v1_org_policy_constraints_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x13, 0x4f, 0x72,
	0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x74, 0x12, 0x35, 0x0a, 0x14, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x12, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x53, 0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x22, 0xd5, 0x01,
	0x0a, 0x19, 0x4f, 0x72, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x63, 0x0a, 0x11, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f,
	0x73, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x10,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x12, 0x53, 0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f,
	0x73, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x75, 0x6c, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x42, 0x91, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x19, 0x4f,
	0x72, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f,
	0x73, 0x74, 0x75, 0x72, 0x65, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_securityposture_v1_org_policy_constraints_proto_rawDescOnce sync.Once
	file_google_cloud_securityposture_v1_org_policy_constraints_proto_rawDescData = file_google_cloud_securityposture_v1_org_policy_constraints_proto_rawDesc
)

func file_google_cloud_securityposture_v1_org_policy_constraints_proto_rawDescGZIP() []byte {
	file_google_cloud_securityposture_v1_org_policy_constraints_proto_rawDescOnce.Do(func() {
		file_google_cloud_securityposture_v1_org_policy_constraints_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securityposture_v1_org_policy_constraints_proto_rawDescData)
	})
	return file_google_cloud_securityposture_v1_org_policy_constraints_proto_rawDescData
}

var file_google_cloud_securityposture_v1_org_policy_constraints_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_securityposture_v1_org_policy_constraints_proto_goTypes = []interface{}{
	(*OrgPolicyConstraint)(nil),       // 0: google.cloud.securityposture.v1.OrgPolicyConstraint
	(*OrgPolicyConstraintCustom)(nil), // 1: google.cloud.securityposture.v1.OrgPolicyConstraintCustom
	(*PolicyRule)(nil),                // 2: google.cloud.securityposture.v1.PolicyRule
	(*CustomConstraint)(nil),          // 3: google.cloud.securityposture.v1.CustomConstraint
}
var file_google_cloud_securityposture_v1_org_policy_constraints_proto_depIdxs = []int32{
	2, // 0: google.cloud.securityposture.v1.OrgPolicyConstraint.policy_rules:type_name -> google.cloud.securityposture.v1.PolicyRule
	3, // 1: google.cloud.securityposture.v1.OrgPolicyConstraintCustom.custom_constraint:type_name -> google.cloud.securityposture.v1.CustomConstraint
	2, // 2: google.cloud.securityposture.v1.OrgPolicyConstraintCustom.policy_rules:type_name -> google.cloud.securityposture.v1.PolicyRule
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_securityposture_v1_org_policy_constraints_proto_init() }
func file_google_cloud_securityposture_v1_org_policy_constraints_proto_init() {
	if File_google_cloud_securityposture_v1_org_policy_constraints_proto != nil {
		return
	}
	file_google_cloud_securityposture_v1_org_policy_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securityposture_v1_org_policy_constraints_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgPolicyConstraint); i {
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
		file_google_cloud_securityposture_v1_org_policy_constraints_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgPolicyConstraintCustom); i {
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
			RawDescriptor: file_google_cloud_securityposture_v1_org_policy_constraints_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securityposture_v1_org_policy_constraints_proto_goTypes,
		DependencyIndexes: file_google_cloud_securityposture_v1_org_policy_constraints_proto_depIdxs,
		MessageInfos:      file_google_cloud_securityposture_v1_org_policy_constraints_proto_msgTypes,
	}.Build()
	File_google_cloud_securityposture_v1_org_policy_constraints_proto = out.File
	file_google_cloud_securityposture_v1_org_policy_constraints_proto_rawDesc = nil
	file_google_cloud_securityposture_v1_org_policy_constraints_proto_goTypes = nil
	file_google_cloud_securityposture_v1_org_policy_constraints_proto_depIdxs = nil
}
