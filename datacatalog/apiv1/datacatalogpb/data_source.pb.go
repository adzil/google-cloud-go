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
// source: google/cloud/datacatalog/v1/data_source.proto

package datacatalogpb

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

// Name of a service that stores the data.
type DataSource_Service int32

const (
	// Default unknown service.
	DataSource_SERVICE_UNSPECIFIED DataSource_Service = 0
	// Google Cloud Storage service.
	DataSource_CLOUD_STORAGE DataSource_Service = 1
	// BigQuery service.
	DataSource_BIGQUERY DataSource_Service = 2
)

// Enum value maps for DataSource_Service.
var (
	DataSource_Service_name = map[int32]string{
		0: "SERVICE_UNSPECIFIED",
		1: "CLOUD_STORAGE",
		2: "BIGQUERY",
	}
	DataSource_Service_value = map[string]int32{
		"SERVICE_UNSPECIFIED": 0,
		"CLOUD_STORAGE":       1,
		"BIGQUERY":            2,
	}
)

func (x DataSource_Service) Enum() *DataSource_Service {
	p := new(DataSource_Service)
	*p = x
	return p
}

func (x DataSource_Service) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataSource_Service) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_datacatalog_v1_data_source_proto_enumTypes[0].Descriptor()
}

func (DataSource_Service) Type() protoreflect.EnumType {
	return &file_google_cloud_datacatalog_v1_data_source_proto_enumTypes[0]
}

func (x DataSource_Service) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataSource_Service.Descriptor instead.
func (DataSource_Service) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_data_source_proto_rawDescGZIP(), []int{0, 0}
}

// Physical location of an entry.
type DataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Service that physically stores the data.
	Service DataSource_Service `protobuf:"varint,1,opt,name=service,proto3,enum=google.cloud.datacatalog.v1.DataSource_Service" json:"service,omitempty"`
	// Full name of a resource as defined by the service. For example:
	//
	// `//bigquery.googleapis.com/projects/{PROJECT_ID}/locations/{LOCATION}/datasets/{DATASET_ID}/tables/{TABLE_ID}`
	Resource string `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	// Output only. Data Catalog entry name, if applicable.
	SourceEntry string `protobuf:"bytes,3,opt,name=source_entry,json=sourceEntry,proto3" json:"source_entry,omitempty"`
	// Types that are assignable to Properties:
	//
	//	*DataSource_StorageProperties
	Properties isDataSource_Properties `protobuf_oneof:"properties"`
}

func (x *DataSource) Reset() {
	*x = DataSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_data_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSource) ProtoMessage() {}

func (x *DataSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_data_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSource.ProtoReflect.Descriptor instead.
func (*DataSource) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_data_source_proto_rawDescGZIP(), []int{0}
}

func (x *DataSource) GetService() DataSource_Service {
	if x != nil {
		return x.Service
	}
	return DataSource_SERVICE_UNSPECIFIED
}

func (x *DataSource) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *DataSource) GetSourceEntry() string {
	if x != nil {
		return x.SourceEntry
	}
	return ""
}

func (m *DataSource) GetProperties() isDataSource_Properties {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (x *DataSource) GetStorageProperties() *StorageProperties {
	if x, ok := x.GetProperties().(*DataSource_StorageProperties); ok {
		return x.StorageProperties
	}
	return nil
}

type isDataSource_Properties interface {
	isDataSource_Properties()
}

type DataSource_StorageProperties struct {
	// Detailed properties of the underlying storage.
	StorageProperties *StorageProperties `protobuf:"bytes,4,opt,name=storage_properties,json=storageProperties,proto3,oneof"`
}

func (*DataSource_StorageProperties) isDataSource_Properties() {}

// Details the properties of the underlying storage.
type StorageProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Patterns to identify a set of files for this fileset.
	//
	// Examples of a valid `file_pattern`:
	//
	//   - `gs://bucket_name/dir/*`: matches all files in the `bucket_name/dir`
	//     directory
	//   - `gs://bucket_name/dir/**`: matches all files in the `bucket_name/dir`
	//     and all subdirectories recursively
	//   - `gs://bucket_name/file*`: matches files prefixed by `file` in
	//     `bucket_name`
	//   - `gs://bucket_name/??.txt`: matches files with two characters followed by
	//     `.txt` in `bucket_name`
	//   - `gs://bucket_name/[aeiou].txt`: matches files that contain a single
	//     vowel character followed by `.txt` in
	//     `bucket_name`
	//   - `gs://bucket_name/[a-m].txt`: matches files that contain `a`, `b`, ...
	//     or `m` followed by `.txt` in `bucket_name`
	//   - `gs://bucket_name/a/*/b`: matches all files in `bucket_name` that match
	//     the `a/*/b` pattern, such as `a/c/b`, `a/d/b`
	//   - `gs://another_bucket/a.txt`: matches `gs://another_bucket/a.txt`
	FilePattern []string `protobuf:"bytes,1,rep,name=file_pattern,json=filePattern,proto3" json:"file_pattern,omitempty"`
	// File type in MIME format, for example, `text/plain`.
	FileType string `protobuf:"bytes,2,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
}

func (x *StorageProperties) Reset() {
	*x = StorageProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_data_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageProperties) ProtoMessage() {}

func (x *StorageProperties) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_data_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageProperties.ProtoReflect.Descriptor instead.
func (*StorageProperties) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_data_source_proto_rawDescGZIP(), []int{1}
}

func (x *StorageProperties) GetFilePattern() []string {
	if x != nil {
		return x.FilePattern
	}
	return nil
}

func (x *StorageProperties) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

var File_google_cloud_datacatalog_v1_data_source_proto protoreflect.FileDescriptor

var file_google_cloud_datacatalog_v1_data_source_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x02,
	0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x5f, 0x0a, 0x12, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x48, 0x00, 0x52, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x07,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47,
	0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x49, 0x47, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10,
	0x02, 0x42, 0x0c, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22,
	0x53, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x42, 0xd7, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x70, 0x62,
	0x3b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0xf8, 0x01,
	0x01, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_datacatalog_v1_data_source_proto_rawDescOnce sync.Once
	file_google_cloud_datacatalog_v1_data_source_proto_rawDescData = file_google_cloud_datacatalog_v1_data_source_proto_rawDesc
)

func file_google_cloud_datacatalog_v1_data_source_proto_rawDescGZIP() []byte {
	file_google_cloud_datacatalog_v1_data_source_proto_rawDescOnce.Do(func() {
		file_google_cloud_datacatalog_v1_data_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_datacatalog_v1_data_source_proto_rawDescData)
	})
	return file_google_cloud_datacatalog_v1_data_source_proto_rawDescData
}

var file_google_cloud_datacatalog_v1_data_source_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_datacatalog_v1_data_source_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_datacatalog_v1_data_source_proto_goTypes = []interface{}{
	(DataSource_Service)(0),   // 0: google.cloud.datacatalog.v1.DataSource.Service
	(*DataSource)(nil),        // 1: google.cloud.datacatalog.v1.DataSource
	(*StorageProperties)(nil), // 2: google.cloud.datacatalog.v1.StorageProperties
}
var file_google_cloud_datacatalog_v1_data_source_proto_depIdxs = []int32{
	0, // 0: google.cloud.datacatalog.v1.DataSource.service:type_name -> google.cloud.datacatalog.v1.DataSource.Service
	2, // 1: google.cloud.datacatalog.v1.DataSource.storage_properties:type_name -> google.cloud.datacatalog.v1.StorageProperties
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_datacatalog_v1_data_source_proto_init() }
func file_google_cloud_datacatalog_v1_data_source_proto_init() {
	if File_google_cloud_datacatalog_v1_data_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_datacatalog_v1_data_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSource); i {
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
		file_google_cloud_datacatalog_v1_data_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageProperties); i {
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
	file_google_cloud_datacatalog_v1_data_source_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DataSource_StorageProperties)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_datacatalog_v1_data_source_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_datacatalog_v1_data_source_proto_goTypes,
		DependencyIndexes: file_google_cloud_datacatalog_v1_data_source_proto_depIdxs,
		EnumInfos:         file_google_cloud_datacatalog_v1_data_source_proto_enumTypes,
		MessageInfos:      file_google_cloud_datacatalog_v1_data_source_proto_msgTypes,
	}.Build()
	File_google_cloud_datacatalog_v1_data_source_proto = out.File
	file_google_cloud_datacatalog_v1_data_source_proto_rawDesc = nil
	file_google_cloud_datacatalog_v1_data_source_proto_goTypes = nil
	file_google_cloud_datacatalog_v1_data_source_proto_depIdxs = nil
}
