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
// source: google/cloud/policytroubleshooter/v1/checker.proto

package policytroubleshooterpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request for
// [TroubleshootIamPolicy][google.cloud.policytroubleshooter.v1.IamChecker.TroubleshootIamPolicy].
type TroubleshootIamPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The information to use for checking whether a principal has a permission
	// for a resource.
	AccessTuple *AccessTuple `protobuf:"bytes,1,opt,name=access_tuple,json=accessTuple,proto3" json:"access_tuple,omitempty"`
}

func (x *TroubleshootIamPolicyRequest) Reset() {
	*x = TroubleshootIamPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_policytroubleshooter_v1_checker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TroubleshootIamPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TroubleshootIamPolicyRequest) ProtoMessage() {}

func (x *TroubleshootIamPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_policytroubleshooter_v1_checker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TroubleshootIamPolicyRequest.ProtoReflect.Descriptor instead.
func (*TroubleshootIamPolicyRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_policytroubleshooter_v1_checker_proto_rawDescGZIP(), []int{0}
}

func (x *TroubleshootIamPolicyRequest) GetAccessTuple() *AccessTuple {
	if x != nil {
		return x.AccessTuple
	}
	return nil
}

// Response for
// [TroubleshootIamPolicy][google.cloud.policytroubleshooter.v1.IamChecker.TroubleshootIamPolicy].
type TroubleshootIamPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates whether the principal has the specified permission for the
	// specified resource, based on evaluating all of the applicable IAM policies.
	Access AccessState `protobuf:"varint,1,opt,name=access,proto3,enum=google.cloud.policytroubleshooter.v1.AccessState" json:"access,omitempty"`
	// List of IAM policies that were evaluated to check the principal's
	// permissions, with annotations to indicate how each policy contributed to
	// the final result.
	//
	// The list of policies can include the policy for the resource itself. It can
	// also include policies that are inherited from higher levels of the resource
	// hierarchy, including the organization, the folder, and the project.
	//
	// To learn more about the resource hierarchy, see
	// https://cloud.google.com/iam/help/resource-hierarchy.
	ExplainedPolicies []*ExplainedPolicy `protobuf:"bytes,2,rep,name=explained_policies,json=explainedPolicies,proto3" json:"explained_policies,omitempty"`
	// The general errors contained in the troubleshooting response.
	Errors []*status.Status `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *TroubleshootIamPolicyResponse) Reset() {
	*x = TroubleshootIamPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_policytroubleshooter_v1_checker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TroubleshootIamPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TroubleshootIamPolicyResponse) ProtoMessage() {}

func (x *TroubleshootIamPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_policytroubleshooter_v1_checker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TroubleshootIamPolicyResponse.ProtoReflect.Descriptor instead.
func (*TroubleshootIamPolicyResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_policytroubleshooter_v1_checker_proto_rawDescGZIP(), []int{1}
}

func (x *TroubleshootIamPolicyResponse) GetAccess() AccessState {
	if x != nil {
		return x.Access
	}
	return AccessState_ACCESS_STATE_UNSPECIFIED
}

func (x *TroubleshootIamPolicyResponse) GetExplainedPolicies() []*ExplainedPolicy {
	if x != nil {
		return x.ExplainedPolicies
	}
	return nil
}

func (x *TroubleshootIamPolicyResponse) GetErrors() []*status.Status {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_google_cloud_policytroubleshooter_v1_checker_proto protoreflect.FileDescriptor

var file_google_cloud_policytroubleshooter_v1_checker_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x74, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f,
	0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x74, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x73, 0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x74,
	0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x1c, 0x54, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68,
	0x6f, 0x6f, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x75,
	0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x74,
	0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x22, 0xfc, 0x01, 0x0a, 0x1d, 0x54, 0x72,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x74, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x64, 0x0a, 0x12, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x74, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73,
	0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x65, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x11, 0x65, 0x78, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x65, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x32, 0xa9, 0x02, 0x0a, 0x0a, 0x49, 0x61, 0x6d,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x12, 0xc1, 0x01, 0x0a, 0x15, 0x54, 0x72, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x74, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68,
	0x6f, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x73, 0x68, 0x6f, 0x6f, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x74, 0x72, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x3a, 0x74,
	0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x1a, 0x57, 0xca, 0x41, 0x23,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x74, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f,
	0x6f, 0x74, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77,
	0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x42, 0x96, 0x02, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x74, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x0f, 0x49, 0x41, 0x4d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x74, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x74, 0x72, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x74, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72,
	0x70, 0x62, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x72, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x24, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x54, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72,
	0x5c, 0x56, 0x31, 0xea, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x72, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x50, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_policytroubleshooter_v1_checker_proto_rawDescOnce sync.Once
	file_google_cloud_policytroubleshooter_v1_checker_proto_rawDescData = file_google_cloud_policytroubleshooter_v1_checker_proto_rawDesc
)

func file_google_cloud_policytroubleshooter_v1_checker_proto_rawDescGZIP() []byte {
	file_google_cloud_policytroubleshooter_v1_checker_proto_rawDescOnce.Do(func() {
		file_google_cloud_policytroubleshooter_v1_checker_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_policytroubleshooter_v1_checker_proto_rawDescData)
	})
	return file_google_cloud_policytroubleshooter_v1_checker_proto_rawDescData
}

var file_google_cloud_policytroubleshooter_v1_checker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_policytroubleshooter_v1_checker_proto_goTypes = []interface{}{
	(*TroubleshootIamPolicyRequest)(nil),  // 0: google.cloud.policytroubleshooter.v1.TroubleshootIamPolicyRequest
	(*TroubleshootIamPolicyResponse)(nil), // 1: google.cloud.policytroubleshooter.v1.TroubleshootIamPolicyResponse
	(*AccessTuple)(nil),                   // 2: google.cloud.policytroubleshooter.v1.AccessTuple
	(AccessState)(0),                      // 3: google.cloud.policytroubleshooter.v1.AccessState
	(*ExplainedPolicy)(nil),               // 4: google.cloud.policytroubleshooter.v1.ExplainedPolicy
	(*status.Status)(nil),                 // 5: google.rpc.Status
}
var file_google_cloud_policytroubleshooter_v1_checker_proto_depIdxs = []int32{
	2, // 0: google.cloud.policytroubleshooter.v1.TroubleshootIamPolicyRequest.access_tuple:type_name -> google.cloud.policytroubleshooter.v1.AccessTuple
	3, // 1: google.cloud.policytroubleshooter.v1.TroubleshootIamPolicyResponse.access:type_name -> google.cloud.policytroubleshooter.v1.AccessState
	4, // 2: google.cloud.policytroubleshooter.v1.TroubleshootIamPolicyResponse.explained_policies:type_name -> google.cloud.policytroubleshooter.v1.ExplainedPolicy
	5, // 3: google.cloud.policytroubleshooter.v1.TroubleshootIamPolicyResponse.errors:type_name -> google.rpc.Status
	0, // 4: google.cloud.policytroubleshooter.v1.IamChecker.TroubleshootIamPolicy:input_type -> google.cloud.policytroubleshooter.v1.TroubleshootIamPolicyRequest
	1, // 5: google.cloud.policytroubleshooter.v1.IamChecker.TroubleshootIamPolicy:output_type -> google.cloud.policytroubleshooter.v1.TroubleshootIamPolicyResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_policytroubleshooter_v1_checker_proto_init() }
func file_google_cloud_policytroubleshooter_v1_checker_proto_init() {
	if File_google_cloud_policytroubleshooter_v1_checker_proto != nil {
		return
	}
	file_google_cloud_policytroubleshooter_v1_explanations_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_policytroubleshooter_v1_checker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TroubleshootIamPolicyRequest); i {
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
		file_google_cloud_policytroubleshooter_v1_checker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TroubleshootIamPolicyResponse); i {
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
			RawDescriptor: file_google_cloud_policytroubleshooter_v1_checker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_policytroubleshooter_v1_checker_proto_goTypes,
		DependencyIndexes: file_google_cloud_policytroubleshooter_v1_checker_proto_depIdxs,
		MessageInfos:      file_google_cloud_policytroubleshooter_v1_checker_proto_msgTypes,
	}.Build()
	File_google_cloud_policytroubleshooter_v1_checker_proto = out.File
	file_google_cloud_policytroubleshooter_v1_checker_proto_rawDesc = nil
	file_google_cloud_policytroubleshooter_v1_checker_proto_goTypes = nil
	file_google_cloud_policytroubleshooter_v1_checker_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IamCheckerClient is the client API for IamChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IamCheckerClient interface {
	// Checks whether a principal has a specific permission for a specific
	// resource, and explains why the principal does or does not have that
	// permission.
	TroubleshootIamPolicy(ctx context.Context, in *TroubleshootIamPolicyRequest, opts ...grpc.CallOption) (*TroubleshootIamPolicyResponse, error)
}

type iamCheckerClient struct {
	cc grpc.ClientConnInterface
}

func NewIamCheckerClient(cc grpc.ClientConnInterface) IamCheckerClient {
	return &iamCheckerClient{cc}
}

func (c *iamCheckerClient) TroubleshootIamPolicy(ctx context.Context, in *TroubleshootIamPolicyRequest, opts ...grpc.CallOption) (*TroubleshootIamPolicyResponse, error) {
	out := new(TroubleshootIamPolicyResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.policytroubleshooter.v1.IamChecker/TroubleshootIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IamCheckerServer is the server API for IamChecker service.
type IamCheckerServer interface {
	// Checks whether a principal has a specific permission for a specific
	// resource, and explains why the principal does or does not have that
	// permission.
	TroubleshootIamPolicy(context.Context, *TroubleshootIamPolicyRequest) (*TroubleshootIamPolicyResponse, error)
}

// UnimplementedIamCheckerServer can be embedded to have forward compatible implementations.
type UnimplementedIamCheckerServer struct {
}

func (*UnimplementedIamCheckerServer) TroubleshootIamPolicy(context.Context, *TroubleshootIamPolicyRequest) (*TroubleshootIamPolicyResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method TroubleshootIamPolicy not implemented")
}

func RegisterIamCheckerServer(s *grpc.Server, srv IamCheckerServer) {
	s.RegisterService(&_IamChecker_serviceDesc, srv)
}

func _IamChecker_TroubleshootIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TroubleshootIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamCheckerServer).TroubleshootIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.policytroubleshooter.v1.IamChecker/TroubleshootIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamCheckerServer).TroubleshootIamPolicy(ctx, req.(*TroubleshootIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IamChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.policytroubleshooter.v1.IamChecker",
	HandlerType: (*IamCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TroubleshootIamPolicy",
			Handler:    _IamChecker_TroubleshootIamPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/policytroubleshooter/v1/checker.proto",
}
