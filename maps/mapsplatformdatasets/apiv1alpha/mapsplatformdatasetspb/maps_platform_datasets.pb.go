// Copyright 2022 Google LLC
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
// source: google/maps/mapsplatformdatasets/v1alpha/maps_platform_datasets.proto

package mapsplatformdatasetspb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request to create a maps dataset.
type CreateDatasetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Parent project that will own the dataset.
	// Format: projects/{$project_number}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The dataset version to create.
	Dataset *Dataset `protobuf:"bytes,2,opt,name=dataset,proto3" json:"dataset,omitempty"`
}

func (x *CreateDatasetRequest) Reset() {
	*x = CreateDatasetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDatasetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatasetRequest) ProtoMessage() {}

func (x *CreateDatasetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatasetRequest.ProtoReflect.Descriptor instead.
func (*CreateDatasetRequest) Descriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDatasetRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateDatasetRequest) GetDataset() *Dataset {
	if x != nil {
		return x.Dataset
	}
	return nil
}

// Request to update the metadata fields of the dataset.
type UpdateDatasetMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The dataset to update. The dataset's name is used to identify the dataset
	// to be updated. The name has the format:
	// projects/{project}/datasets/{dataset_id}
	Dataset *Dataset `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
	// The list of fields to be updated. Support the value "*" for full
	// replacement.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateDatasetMetadataRequest) Reset() {
	*x = UpdateDatasetMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDatasetMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDatasetMetadataRequest) ProtoMessage() {}

func (x *UpdateDatasetMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDatasetMetadataRequest.ProtoReflect.Descriptor instead.
func (*UpdateDatasetMetadataRequest) Descriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateDatasetMetadataRequest) GetDataset() *Dataset {
	if x != nil {
		return x.Dataset
	}
	return nil
}

func (x *UpdateDatasetMetadataRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Request to get the specified dataset.
type GetDatasetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name. Can also fetch a specified version
	// projects/{project}/datasets/{dataset_id}
	// projects/{project}/datasets/{dataset_id}@{version-id}
	//
	// In order to retrieve a previous version of the dataset, also provide
	// the version ID.
	// Example: projects/123/datasets/assisted-driving-preferences@c7cfa2a8
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// If specified, will fetch the dataset details of the version published for
	// the specified use case rather than the latest, if one exists. If a
	// published version does not exist, will default to getting the dataset
	// details of the latest version.
	PublishedUsage Usage `protobuf:"varint,2,opt,name=published_usage,json=publishedUsage,proto3,enum=google.maps.mapsplatformdatasets.v1alpha.Usage" json:"published_usage,omitempty"`
}

func (x *GetDatasetRequest) Reset() {
	*x = GetDatasetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDatasetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatasetRequest) ProtoMessage() {}

func (x *GetDatasetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatasetRequest.ProtoReflect.Descriptor instead.
func (*GetDatasetRequest) Descriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescGZIP(), []int{2}
}

func (x *GetDatasetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetDatasetRequest) GetPublishedUsage() Usage {
	if x != nil {
		return x.PublishedUsage
	}
	return Usage_USAGE_UNSPECIFIED
}

// Request to list of all versions of the dataset.
type ListDatasetVersionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the dataset to list all the versions for.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The maximum number of versions to return per page.
	// If unspecified (or zero), at most 1000 versions will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The page token, received from a previous GetDatasetVersions call.
	// Provide this to retrieve the subsequent page.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListDatasetVersionsRequest) Reset() {
	*x = ListDatasetVersionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatasetVersionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatasetVersionsRequest) ProtoMessage() {}

func (x *ListDatasetVersionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatasetVersionsRequest.ProtoReflect.Descriptor instead.
func (*ListDatasetVersionsRequest) Descriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescGZIP(), []int{3}
}

func (x *ListDatasetVersionsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListDatasetVersionsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDatasetVersionsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response with list of all versions of the dataset.
type ListDatasetVersionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All the versions of the dataset.
	Datasets []*Dataset `protobuf:"bytes,1,rep,name=datasets,proto3" json:"datasets,omitempty"`
	// A token that can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListDatasetVersionsResponse) Reset() {
	*x = ListDatasetVersionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatasetVersionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatasetVersionsResponse) ProtoMessage() {}

func (x *ListDatasetVersionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatasetVersionsResponse.ProtoReflect.Descriptor instead.
func (*ListDatasetVersionsResponse) Descriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescGZIP(), []int{4}
}

func (x *ListDatasetVersionsResponse) GetDatasets() []*Dataset {
	if x != nil {
		return x.Datasets
	}
	return nil
}

func (x *ListDatasetVersionsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request to list datasets for the project.
type ListDatasetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the project to list all the datasets for.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of versions to return per page.
	// If unspecified (or zero), at most 1000 datasets will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The page token, received from a previous GetDatasetVersions call.
	// Provide this to retrieve the subsequent page.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListDatasetsRequest) Reset() {
	*x = ListDatasetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatasetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatasetsRequest) ProtoMessage() {}

func (x *ListDatasetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatasetsRequest.ProtoReflect.Descriptor instead.
func (*ListDatasetsRequest) Descriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescGZIP(), []int{5}
}

func (x *ListDatasetsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListDatasetsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDatasetsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response to list datasets for the project.
type ListDatasetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All the datasets for the project.
	Datasets []*Dataset `protobuf:"bytes,1,rep,name=datasets,proto3" json:"datasets,omitempty"`
	// A token that can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListDatasetsResponse) Reset() {
	*x = ListDatasetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatasetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatasetsResponse) ProtoMessage() {}

func (x *ListDatasetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatasetsResponse.ProtoReflect.Descriptor instead.
func (*ListDatasetsResponse) Descriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescGZIP(), []int{6}
}

func (x *ListDatasetsResponse) GetDatasets() []*Dataset {
	if x != nil {
		return x.Datasets
	}
	return nil
}

func (x *ListDatasetsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request to delete a dataset.
//
// The dataset to be deleted.
type DeleteDatasetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Format: projects/${project}/datasets/{dataset_id}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// If set to true, any dataset version for this dataset will also be deleted.
	// (Otherwise, the request will only work if the dataset has no versions.)
	Force bool `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *DeleteDatasetRequest) Reset() {
	*x = DeleteDatasetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDatasetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatasetRequest) ProtoMessage() {}

func (x *DeleteDatasetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatasetRequest.ProtoReflect.Descriptor instead.
func (*DeleteDatasetRequest) Descriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteDatasetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteDatasetRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

// Request to delete a version of a dataset.
type DeleteDatasetVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Format: projects/${project}/datasets/{dataset_id}@{version-id}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteDatasetVersionRequest) Reset() {
	*x = DeleteDatasetVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDatasetVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatasetVersionRequest) ProtoMessage() {}

func (x *DeleteDatasetVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatasetVersionRequest.ProtoReflect.Descriptor instead.
func (*DeleteDatasetVersionRequest) Descriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteDatasetVersionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto protoreflect.FileDescriptor

var file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x6d, 0x61,
	0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x5f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x4b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x50, 0x0a,
	0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70,
	0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x22,
	0xad, 0x01, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x50, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22,
	0xb6, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x6d, 0x61, 0x70,
	0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x58,
	0x0a, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x6d,
	0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x94, 0x01, 0x0a,
	0x1b, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70,
	0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x9e, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8d, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61,
	0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x75, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa,
	0x41, 0x2d, 0x0a, 0x2b, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x22, 0x66, 0x0a, 0x1b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d,
	0x0a, 0x2b, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x89, 0x02, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x42, 0x19, 0x4d, 0x61, 0x70, 0x73, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x66, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x61, 0x70, 0x73,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x70, 0x62, 0x3b, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x70, 0x62, 0xaa, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x56, 0x31, 0x41,
	0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61,
	0x70, 0x73, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescOnce sync.Once
	file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescData = file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDesc
)

func file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescGZIP() []byte {
	file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescOnce.Do(func() {
		file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescData)
	})
	return file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDescData
}

var file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_goTypes = []interface{}{
	(*CreateDatasetRequest)(nil),         // 0: google.maps.mapsplatformdatasets.v1alpha.CreateDatasetRequest
	(*UpdateDatasetMetadataRequest)(nil), // 1: google.maps.mapsplatformdatasets.v1alpha.UpdateDatasetMetadataRequest
	(*GetDatasetRequest)(nil),            // 2: google.maps.mapsplatformdatasets.v1alpha.GetDatasetRequest
	(*ListDatasetVersionsRequest)(nil),   // 3: google.maps.mapsplatformdatasets.v1alpha.ListDatasetVersionsRequest
	(*ListDatasetVersionsResponse)(nil),  // 4: google.maps.mapsplatformdatasets.v1alpha.ListDatasetVersionsResponse
	(*ListDatasetsRequest)(nil),          // 5: google.maps.mapsplatformdatasets.v1alpha.ListDatasetsRequest
	(*ListDatasetsResponse)(nil),         // 6: google.maps.mapsplatformdatasets.v1alpha.ListDatasetsResponse
	(*DeleteDatasetRequest)(nil),         // 7: google.maps.mapsplatformdatasets.v1alpha.DeleteDatasetRequest
	(*DeleteDatasetVersionRequest)(nil),  // 8: google.maps.mapsplatformdatasets.v1alpha.DeleteDatasetVersionRequest
	(*Dataset)(nil),                      // 9: google.maps.mapsplatformdatasets.v1alpha.Dataset
	(*fieldmaskpb.FieldMask)(nil),        // 10: google.protobuf.FieldMask
	(Usage)(0),                           // 11: google.maps.mapsplatformdatasets.v1alpha.Usage
}
var file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_depIdxs = []int32{
	9,  // 0: google.maps.mapsplatformdatasets.v1alpha.CreateDatasetRequest.dataset:type_name -> google.maps.mapsplatformdatasets.v1alpha.Dataset
	9,  // 1: google.maps.mapsplatformdatasets.v1alpha.UpdateDatasetMetadataRequest.dataset:type_name -> google.maps.mapsplatformdatasets.v1alpha.Dataset
	10, // 2: google.maps.mapsplatformdatasets.v1alpha.UpdateDatasetMetadataRequest.update_mask:type_name -> google.protobuf.FieldMask
	11, // 3: google.maps.mapsplatformdatasets.v1alpha.GetDatasetRequest.published_usage:type_name -> google.maps.mapsplatformdatasets.v1alpha.Usage
	9,  // 4: google.maps.mapsplatformdatasets.v1alpha.ListDatasetVersionsResponse.datasets:type_name -> google.maps.mapsplatformdatasets.v1alpha.Dataset
	9,  // 5: google.maps.mapsplatformdatasets.v1alpha.ListDatasetsResponse.datasets:type_name -> google.maps.mapsplatformdatasets.v1alpha.Dataset
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_init() }
func file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_init() {
	if File_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto != nil {
		return
	}
	file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDatasetRequest); i {
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
		file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDatasetMetadataRequest); i {
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
		file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDatasetRequest); i {
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
		file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatasetVersionsRequest); i {
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
		file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatasetVersionsResponse); i {
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
		file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatasetsRequest); i {
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
		file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatasetsResponse); i {
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
		file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDatasetRequest); i {
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
		file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDatasetVersionRequest); i {
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
			RawDescriptor: file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_goTypes,
		DependencyIndexes: file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_depIdxs,
		MessageInfos:      file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_msgTypes,
	}.Build()
	File_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto = out.File
	file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_rawDesc = nil
	file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_goTypes = nil
	file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_depIdxs = nil
}
