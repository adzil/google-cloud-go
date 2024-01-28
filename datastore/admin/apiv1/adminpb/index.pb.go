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
// source: google/datastore/admin/v1/index.proto

package adminpb

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

// For an ordered index, specifies whether each of the entity's ancestors
// will be included.
type Index_AncestorMode int32

const (
	// The ancestor mode is unspecified.
	Index_ANCESTOR_MODE_UNSPECIFIED Index_AncestorMode = 0
	// Do not include the entity's ancestors in the index.
	Index_NONE Index_AncestorMode = 1
	// Include all the entity's ancestors in the index.
	Index_ALL_ANCESTORS Index_AncestorMode = 2
)

// Enum value maps for Index_AncestorMode.
var (
	Index_AncestorMode_name = map[int32]string{
		0: "ANCESTOR_MODE_UNSPECIFIED",
		1: "NONE",
		2: "ALL_ANCESTORS",
	}
	Index_AncestorMode_value = map[string]int32{
		"ANCESTOR_MODE_UNSPECIFIED": 0,
		"NONE":                      1,
		"ALL_ANCESTORS":             2,
	}
)

func (x Index_AncestorMode) Enum() *Index_AncestorMode {
	p := new(Index_AncestorMode)
	*p = x
	return p
}

func (x Index_AncestorMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Index_AncestorMode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_datastore_admin_v1_index_proto_enumTypes[0].Descriptor()
}

func (Index_AncestorMode) Type() protoreflect.EnumType {
	return &file_google_datastore_admin_v1_index_proto_enumTypes[0]
}

func (x Index_AncestorMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Index_AncestorMode.Descriptor instead.
func (Index_AncestorMode) EnumDescriptor() ([]byte, []int) {
	return file_google_datastore_admin_v1_index_proto_rawDescGZIP(), []int{0, 0}
}

// The direction determines how a property is indexed.
type Index_Direction int32

const (
	// The direction is unspecified.
	Index_DIRECTION_UNSPECIFIED Index_Direction = 0
	// The property's values are indexed so as to support sequencing in
	// ascending order and also query by <, >, <=, >=, and =.
	Index_ASCENDING Index_Direction = 1
	// The property's values are indexed so as to support sequencing in
	// descending order and also query by <, >, <=, >=, and =.
	Index_DESCENDING Index_Direction = 2
)

// Enum value maps for Index_Direction.
var (
	Index_Direction_name = map[int32]string{
		0: "DIRECTION_UNSPECIFIED",
		1: "ASCENDING",
		2: "DESCENDING",
	}
	Index_Direction_value = map[string]int32{
		"DIRECTION_UNSPECIFIED": 0,
		"ASCENDING":             1,
		"DESCENDING":            2,
	}
)

func (x Index_Direction) Enum() *Index_Direction {
	p := new(Index_Direction)
	*p = x
	return p
}

func (x Index_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Index_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_google_datastore_admin_v1_index_proto_enumTypes[1].Descriptor()
}

func (Index_Direction) Type() protoreflect.EnumType {
	return &file_google_datastore_admin_v1_index_proto_enumTypes[1]
}

func (x Index_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Index_Direction.Descriptor instead.
func (Index_Direction) EnumDescriptor() ([]byte, []int) {
	return file_google_datastore_admin_v1_index_proto_rawDescGZIP(), []int{0, 1}
}

// The possible set of states of an index.
type Index_State int32

const (
	// The state is unspecified.
	Index_STATE_UNSPECIFIED Index_State = 0
	// The index is being created, and cannot be used by queries.
	// There is an active long-running operation for the index.
	// The index is updated when writing an entity.
	// Some index data may exist.
	Index_CREATING Index_State = 1
	// The index is ready to be used.
	// The index is updated when writing an entity.
	// The index is fully populated from all stored entities it applies to.
	Index_READY Index_State = 2
	// The index is being deleted, and cannot be used by queries.
	// There is an active long-running operation for the index.
	// The index is not updated when writing an entity.
	// Some index data may exist.
	Index_DELETING Index_State = 3
	// The index was being created or deleted, but something went wrong.
	// The index cannot by used by queries.
	// There is no active long-running operation for the index,
	// and the most recently finished long-running operation failed.
	// The index is not updated when writing an entity.
	// Some index data may exist.
	Index_ERROR Index_State = 4
)

// Enum value maps for Index_State.
var (
	Index_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "DELETING",
		4: "ERROR",
	}
	Index_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"READY":             2,
		"DELETING":          3,
		"ERROR":             4,
	}
)

func (x Index_State) Enum() *Index_State {
	p := new(Index_State)
	*p = x
	return p
}

func (x Index_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Index_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_datastore_admin_v1_index_proto_enumTypes[2].Descriptor()
}

func (Index_State) Type() protoreflect.EnumType {
	return &file_google_datastore_admin_v1_index_proto_enumTypes[2]
}

func (x Index_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Index_State.Descriptor instead.
func (Index_State) EnumDescriptor() ([]byte, []int) {
	return file_google_datastore_admin_v1_index_proto_rawDescGZIP(), []int{0, 2}
}

// Datastore composite index definition.
type Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Project ID.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Output only. The resource ID of the index.
	IndexId string `protobuf:"bytes,3,opt,name=index_id,json=indexId,proto3" json:"index_id,omitempty"`
	// Required. The entity kind to which this index applies.
	Kind string `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
	// Required. The index's ancestor mode.  Must not be
	// ANCESTOR_MODE_UNSPECIFIED.
	Ancestor Index_AncestorMode `protobuf:"varint,5,opt,name=ancestor,proto3,enum=google.datastore.admin.v1.Index_AncestorMode" json:"ancestor,omitempty"`
	// Required. An ordered sequence of property names and their index attributes.
	//
	// Requires:
	//
	// * A maximum of 100 properties.
	Properties []*Index_IndexedProperty `protobuf:"bytes,6,rep,name=properties,proto3" json:"properties,omitempty"`
	// Output only. The state of the index.
	State Index_State `protobuf:"varint,7,opt,name=state,proto3,enum=google.datastore.admin.v1.Index_State" json:"state,omitempty"`
}

func (x *Index) Reset() {
	*x = Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_admin_v1_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_admin_v1_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_google_datastore_admin_v1_index_proto_rawDescGZIP(), []int{0}
}

func (x *Index) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Index) GetIndexId() string {
	if x != nil {
		return x.IndexId
	}
	return ""
}

func (x *Index) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Index) GetAncestor() Index_AncestorMode {
	if x != nil {
		return x.Ancestor
	}
	return Index_ANCESTOR_MODE_UNSPECIFIED
}

func (x *Index) GetProperties() []*Index_IndexedProperty {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Index) GetState() Index_State {
	if x != nil {
		return x.State
	}
	return Index_STATE_UNSPECIFIED
}

// A property of an index.
type Index_IndexedProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The property name to index.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The indexed property's direction.  Must not be
	// DIRECTION_UNSPECIFIED.
	Direction Index_Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=google.datastore.admin.v1.Index_Direction" json:"direction,omitempty"`
}

func (x *Index_IndexedProperty) Reset() {
	*x = Index_IndexedProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_admin_v1_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index_IndexedProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index_IndexedProperty) ProtoMessage() {}

func (x *Index_IndexedProperty) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_admin_v1_index_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index_IndexedProperty.ProtoReflect.Descriptor instead.
func (*Index_IndexedProperty) Descriptor() ([]byte, []int) {
	return file_google_datastore_admin_v1_index_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Index_IndexedProperty) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Index_IndexedProperty) GetDirection() Index_Direction {
	if x != nil {
		return x.Direction
	}
	return Index_DIRECTION_UNSPECIFIED
}

var File_google_datastore_admin_v1_index_proto protoreflect.FileDescriptor

var file_google_datastore_admin_v1_index_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xae, 0x05, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x08, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x4e, 0x0a, 0x08, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x41,
	0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x08, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x12, 0x55, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x41, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x1a, 0x79, 0x0a, 0x0f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x4d, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x4a, 0x0a, 0x0c, 0x41, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x1d, 0x0a, 0x19, 0x41, 0x4e, 0x43, 0x45, 0x53, 0x54, 0x4f, 0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x4c, 0x4c, 0x5f,
	0x41, 0x4e, 0x43, 0x45, 0x53, 0x54, 0x4f, 0x52, 0x53, 0x10, 0x02, 0x22, 0x45, 0x0a, 0x09, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x49, 0x52, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x22, 0x50, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x04, 0x42, 0xd2, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0xaa,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x5c, 0x56, 0x31, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_datastore_admin_v1_index_proto_rawDescOnce sync.Once
	file_google_datastore_admin_v1_index_proto_rawDescData = file_google_datastore_admin_v1_index_proto_rawDesc
)

func file_google_datastore_admin_v1_index_proto_rawDescGZIP() []byte {
	file_google_datastore_admin_v1_index_proto_rawDescOnce.Do(func() {
		file_google_datastore_admin_v1_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_datastore_admin_v1_index_proto_rawDescData)
	})
	return file_google_datastore_admin_v1_index_proto_rawDescData
}

var file_google_datastore_admin_v1_index_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_google_datastore_admin_v1_index_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_datastore_admin_v1_index_proto_goTypes = []interface{}{
	(Index_AncestorMode)(0),       // 0: google.datastore.admin.v1.Index.AncestorMode
	(Index_Direction)(0),          // 1: google.datastore.admin.v1.Index.Direction
	(Index_State)(0),              // 2: google.datastore.admin.v1.Index.State
	(*Index)(nil),                 // 3: google.datastore.admin.v1.Index
	(*Index_IndexedProperty)(nil), // 4: google.datastore.admin.v1.Index.IndexedProperty
}
var file_google_datastore_admin_v1_index_proto_depIdxs = []int32{
	0, // 0: google.datastore.admin.v1.Index.ancestor:type_name -> google.datastore.admin.v1.Index.AncestorMode
	4, // 1: google.datastore.admin.v1.Index.properties:type_name -> google.datastore.admin.v1.Index.IndexedProperty
	2, // 2: google.datastore.admin.v1.Index.state:type_name -> google.datastore.admin.v1.Index.State
	1, // 3: google.datastore.admin.v1.Index.IndexedProperty.direction:type_name -> google.datastore.admin.v1.Index.Direction
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_datastore_admin_v1_index_proto_init() }
func file_google_datastore_admin_v1_index_proto_init() {
	if File_google_datastore_admin_v1_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_datastore_admin_v1_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index); i {
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
		file_google_datastore_admin_v1_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index_IndexedProperty); i {
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
			RawDescriptor: file_google_datastore_admin_v1_index_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_datastore_admin_v1_index_proto_goTypes,
		DependencyIndexes: file_google_datastore_admin_v1_index_proto_depIdxs,
		EnumInfos:         file_google_datastore_admin_v1_index_proto_enumTypes,
		MessageInfos:      file_google_datastore_admin_v1_index_proto_msgTypes,
	}.Build()
	File_google_datastore_admin_v1_index_proto = out.File
	file_google_datastore_admin_v1_index_proto_rawDesc = nil
	file_google_datastore_admin_v1_index_proto_goTypes = nil
	file_google_datastore_admin_v1_index_proto_depIdxs = nil
}
