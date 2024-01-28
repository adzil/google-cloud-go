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
// source: google/cloud/netapp/v1/snapshot.proto

package netapppb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The Snapshot States
type Snapshot_State int32

const (
	// Unspecified Snapshot State
	Snapshot_STATE_UNSPECIFIED Snapshot_State = 0
	// Snapshot State is Ready
	Snapshot_READY Snapshot_State = 1
	// Snapshot State is Creating
	Snapshot_CREATING Snapshot_State = 2
	// Snapshot State is Deleting
	Snapshot_DELETING Snapshot_State = 3
	// Snapshot State is Updating
	Snapshot_UPDATING Snapshot_State = 4
	// Snapshot State is Disabled
	Snapshot_DISABLED Snapshot_State = 5
	// Snapshot State is Error
	Snapshot_ERROR Snapshot_State = 6
)

// Enum value maps for Snapshot_State.
var (
	Snapshot_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "READY",
		2: "CREATING",
		3: "DELETING",
		4: "UPDATING",
		5: "DISABLED",
		6: "ERROR",
	}
	Snapshot_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"READY":             1,
		"CREATING":          2,
		"DELETING":          3,
		"UPDATING":          4,
		"DISABLED":          5,
		"ERROR":             6,
	}
)

func (x Snapshot_State) Enum() *Snapshot_State {
	p := new(Snapshot_State)
	*p = x
	return p
}

func (x Snapshot_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Snapshot_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_netapp_v1_snapshot_proto_enumTypes[0].Descriptor()
}

func (Snapshot_State) Type() protoreflect.EnumType {
	return &file_google_cloud_netapp_v1_snapshot_proto_enumTypes[0]
}

func (x Snapshot_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Snapshot_State.Descriptor instead.
func (Snapshot_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_snapshot_proto_rawDescGZIP(), []int{6, 0}
}

// ListSnapshotsRequest lists snapshots.
type ListSnapshotsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The volume for which to retrieve snapshot information,
	// in the format
	// `projects/{project_id}/locations/{location}/volumes/{volume_id}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value to use if there are additional
	// results to retrieve for this list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Sort results. Supported values are "name", "name desc" or "" (unsorted).
	OrderBy string `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// List filter.
	Filter string `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListSnapshotsRequest) Reset() {
	*x = ListSnapshotsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_netapp_v1_snapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnapshotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnapshotsRequest) ProtoMessage() {}

func (x *ListSnapshotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_netapp_v1_snapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnapshotsRequest.ProtoReflect.Descriptor instead.
func (*ListSnapshotsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *ListSnapshotsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListSnapshotsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListSnapshotsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListSnapshotsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListSnapshotsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// ListSnapshotsResponse is the result of ListSnapshotsRequest.
type ListSnapshotsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of snapshots in the project for the specified volume.
	Snapshots []*Snapshot `protobuf:"bytes,1,rep,name=snapshots,proto3" json:"snapshots,omitempty"`
	// The token you can use to retrieve the next page of results. Not returned
	// if there are no more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Locations that could not be reached.
	Unreachable []string `protobuf:"bytes,3,rep,name=unreachable,proto3" json:"unreachable,omitempty"`
}

func (x *ListSnapshotsResponse) Reset() {
	*x = ListSnapshotsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_netapp_v1_snapshot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnapshotsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnapshotsResponse) ProtoMessage() {}

func (x *ListSnapshotsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_netapp_v1_snapshot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnapshotsResponse.ProtoReflect.Descriptor instead.
func (*ListSnapshotsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_snapshot_proto_rawDescGZIP(), []int{1}
}

func (x *ListSnapshotsResponse) GetSnapshots() []*Snapshot {
	if x != nil {
		return x.Snapshots
	}
	return nil
}

func (x *ListSnapshotsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListSnapshotsResponse) GetUnreachable() []string {
	if x != nil {
		return x.Unreachable
	}
	return nil
}

// GetSnapshotRequest gets the state of a snapshot.
type GetSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The snapshot resource name, in the format
	// `projects/{project_id}/locations/{location}/volumes/{volume_id}/snapshots/{snapshot_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetSnapshotRequest) Reset() {
	*x = GetSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_netapp_v1_snapshot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnapshotRequest) ProtoMessage() {}

func (x *GetSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_netapp_v1_snapshot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnapshotRequest.ProtoReflect.Descriptor instead.
func (*GetSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_snapshot_proto_rawDescGZIP(), []int{2}
}

func (x *GetSnapshotRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// CreateSnapshotRequest creates a snapshot.
type CreateSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The NetApp volume to create the snapshots of, in the format
	// `projects/{project_id}/locations/{location}/volumes/{volume_id}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. A snapshot resource
	Snapshot *Snapshot `protobuf:"bytes,2,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	// Required. ID of the snapshot to create.
	// This value must start with a lowercase letter followed by up to 62
	// lowercase letters, numbers, or hyphens, and cannot end with a hyphen.
	SnapshotId string `protobuf:"bytes,3,opt,name=snapshot_id,json=snapshotId,proto3" json:"snapshot_id,omitempty"`
}

func (x *CreateSnapshotRequest) Reset() {
	*x = CreateSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_netapp_v1_snapshot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnapshotRequest) ProtoMessage() {}

func (x *CreateSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_netapp_v1_snapshot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnapshotRequest.ProtoReflect.Descriptor instead.
func (*CreateSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_snapshot_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSnapshotRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateSnapshotRequest) GetSnapshot() *Snapshot {
	if x != nil {
		return x.Snapshot
	}
	return nil
}

func (x *CreateSnapshotRequest) GetSnapshotId() string {
	if x != nil {
		return x.SnapshotId
	}
	return ""
}

// DeleteSnapshotRequest deletes a snapshot.
type DeleteSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The snapshot resource name, in the format
	// `projects/*/locations/*/volumes/*/snapshots/{snapshot_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteSnapshotRequest) Reset() {
	*x = DeleteSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_netapp_v1_snapshot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSnapshotRequest) ProtoMessage() {}

func (x *DeleteSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_netapp_v1_snapshot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSnapshotRequest.ProtoReflect.Descriptor instead.
func (*DeleteSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_snapshot_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSnapshotRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// UpdateSnapshotRequest updates description and/or labels for a snapshot.
type UpdateSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Mask of fields to update.  At least one path must be supplied in
	// this field.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,1,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Required. A snapshot resource
	Snapshot *Snapshot `protobuf:"bytes,2,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
}

func (x *UpdateSnapshotRequest) Reset() {
	*x = UpdateSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_netapp_v1_snapshot_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSnapshotRequest) ProtoMessage() {}

func (x *UpdateSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_netapp_v1_snapshot_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSnapshotRequest.ProtoReflect.Descriptor instead.
func (*UpdateSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_snapshot_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSnapshotRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateSnapshotRequest) GetSnapshot() *Snapshot {
	if x != nil {
		return x.Snapshot
	}
	return nil
}

// Snapshot is a point-in-time version of a Volume's content.
type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The resource name of the snapshot.
	// Format:
	// `projects/{project_id}/locations/{location}/volumes/{volume_id}/snapshots/{snapshot_id}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The snapshot state.
	State Snapshot_State `protobuf:"varint,2,opt,name=state,proto3,enum=google.cloud.netapp.v1.Snapshot_State" json:"state,omitempty"`
	// Output only. State details of the storage pool
	StateDetails string `protobuf:"bytes,3,opt,name=state_details,json=stateDetails,proto3" json:"state_details,omitempty"`
	// A description of the snapshot with 2048 characters or less.
	// Requests with longer descriptions will be rejected.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. Current storage usage for the snapshot in bytes.
	UsedBytes float64 `protobuf:"fixed64,5,opt,name=used_bytes,json=usedBytes,proto3" json:"used_bytes,omitempty"`
	// Output only. The time when the snapshot was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Resource labels to represent user provided metadata.
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_netapp_v1_snapshot_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_netapp_v1_snapshot_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_snapshot_proto_rawDescGZIP(), []int{6}
}

func (x *Snapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Snapshot) GetState() Snapshot_State {
	if x != nil {
		return x.State
	}
	return Snapshot_STATE_UNSPECIFIED
}

func (x *Snapshot) GetStateDetails() string {
	if x != nil {
		return x.StateDetails
	}
	return ""
}

func (x *Snapshot) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Snapshot) GetUsedBytes() float64 {
	if x != nil {
		return x.UsedBytes
	}
	return 0
}

func (x *Snapshot) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Snapshot) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_google_cloud_netapp_v1_snapshot_proto protoreflect.FileDescriptor

var file_google_cloud_netapp_v1_snapshot_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e,
	0x65, 0x74, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5,
	0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x20, 0x12,
	0x1e, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xa1, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3e, 0x0a, 0x09, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x09, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x72, 0x65,
	0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x75,
	0x6e, 0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x50, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26,
	0xe0, 0x41, 0x02, 0xfa, 0x41, 0x20, 0x0a, 0x1e, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc0, 0x01, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x20, 0x12, 0x1e,
	0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0a, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x64, 0x22,
	0x53, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x20, 0x0a, 0x1e,
	0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x41, 0x0a, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x22, 0x92, 0x05, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0d,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x09, 0x75, 0x73, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x44,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65,
	0x74, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x6c, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x50, 0x44, 0x41, 0x54, 0x49,
	0x4e, 0x47, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44,
	0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x3a, 0x88, 0x01,
	0xea, 0x41, 0x84, 0x01, 0x0a, 0x1e, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x12, 0x4d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2f, 0x7b, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x7d, 0x2f, 0x73,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x2f, 0x7b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x7d, 0x2a, 0x09, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x32, 0x08,
	0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x42, 0xaf, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65,
	0x74, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6e, 0x65,
	0x74, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x61, 0x70,
	0x70, 0x70, 0x62, 0x3b, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x70, 0x62, 0xaa, 0x02, 0x16, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4e, 0x65, 0x74, 0x41, 0x70, 0x70, 0x5c, 0x56, 0x31, 0xea, 0x02,
	0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x4e, 0x65, 0x74, 0x41, 0x70, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_netapp_v1_snapshot_proto_rawDescOnce sync.Once
	file_google_cloud_netapp_v1_snapshot_proto_rawDescData = file_google_cloud_netapp_v1_snapshot_proto_rawDesc
)

func file_google_cloud_netapp_v1_snapshot_proto_rawDescGZIP() []byte {
	file_google_cloud_netapp_v1_snapshot_proto_rawDescOnce.Do(func() {
		file_google_cloud_netapp_v1_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_netapp_v1_snapshot_proto_rawDescData)
	})
	return file_google_cloud_netapp_v1_snapshot_proto_rawDescData
}

var file_google_cloud_netapp_v1_snapshot_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_netapp_v1_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_google_cloud_netapp_v1_snapshot_proto_goTypes = []interface{}{
	(Snapshot_State)(0),           // 0: google.cloud.netapp.v1.Snapshot.State
	(*ListSnapshotsRequest)(nil),  // 1: google.cloud.netapp.v1.ListSnapshotsRequest
	(*ListSnapshotsResponse)(nil), // 2: google.cloud.netapp.v1.ListSnapshotsResponse
	(*GetSnapshotRequest)(nil),    // 3: google.cloud.netapp.v1.GetSnapshotRequest
	(*CreateSnapshotRequest)(nil), // 4: google.cloud.netapp.v1.CreateSnapshotRequest
	(*DeleteSnapshotRequest)(nil), // 5: google.cloud.netapp.v1.DeleteSnapshotRequest
	(*UpdateSnapshotRequest)(nil), // 6: google.cloud.netapp.v1.UpdateSnapshotRequest
	(*Snapshot)(nil),              // 7: google.cloud.netapp.v1.Snapshot
	nil,                           // 8: google.cloud.netapp.v1.Snapshot.LabelsEntry
	(*fieldmaskpb.FieldMask)(nil), // 9: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_google_cloud_netapp_v1_snapshot_proto_depIdxs = []int32{
	7,  // 0: google.cloud.netapp.v1.ListSnapshotsResponse.snapshots:type_name -> google.cloud.netapp.v1.Snapshot
	7,  // 1: google.cloud.netapp.v1.CreateSnapshotRequest.snapshot:type_name -> google.cloud.netapp.v1.Snapshot
	9,  // 2: google.cloud.netapp.v1.UpdateSnapshotRequest.update_mask:type_name -> google.protobuf.FieldMask
	7,  // 3: google.cloud.netapp.v1.UpdateSnapshotRequest.snapshot:type_name -> google.cloud.netapp.v1.Snapshot
	0,  // 4: google.cloud.netapp.v1.Snapshot.state:type_name -> google.cloud.netapp.v1.Snapshot.State
	10, // 5: google.cloud.netapp.v1.Snapshot.create_time:type_name -> google.protobuf.Timestamp
	8,  // 6: google.cloud.netapp.v1.Snapshot.labels:type_name -> google.cloud.netapp.v1.Snapshot.LabelsEntry
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_netapp_v1_snapshot_proto_init() }
func file_google_cloud_netapp_v1_snapshot_proto_init() {
	if File_google_cloud_netapp_v1_snapshot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_netapp_v1_snapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnapshotsRequest); i {
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
		file_google_cloud_netapp_v1_snapshot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnapshotsResponse); i {
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
		file_google_cloud_netapp_v1_snapshot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSnapshotRequest); i {
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
		file_google_cloud_netapp_v1_snapshot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSnapshotRequest); i {
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
		file_google_cloud_netapp_v1_snapshot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSnapshotRequest); i {
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
		file_google_cloud_netapp_v1_snapshot_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSnapshotRequest); i {
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
		file_google_cloud_netapp_v1_snapshot_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snapshot); i {
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
			RawDescriptor: file_google_cloud_netapp_v1_snapshot_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_netapp_v1_snapshot_proto_goTypes,
		DependencyIndexes: file_google_cloud_netapp_v1_snapshot_proto_depIdxs,
		EnumInfos:         file_google_cloud_netapp_v1_snapshot_proto_enumTypes,
		MessageInfos:      file_google_cloud_netapp_v1_snapshot_proto_msgTypes,
	}.Build()
	File_google_cloud_netapp_v1_snapshot_proto = out.File
	file_google_cloud_netapp_v1_snapshot_proto_rawDesc = nil
	file_google_cloud_netapp_v1_snapshot_proto_goTypes = nil
	file_google_cloud_netapp_v1_snapshot_proto_depIdxs = nil
}
