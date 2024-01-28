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
// source: google/cloud/talent/v4beta1/histogram.proto

package talentpb

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

// The histogram request.
type HistogramQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An expression specifies a histogram request against matching resources
	// (for example, jobs, profiles) for searches.
	//
	// See
	// [SearchJobsRequest.histogram_queries][google.cloud.talent.v4beta1.SearchJobsRequest.histogram_queries]
	// and
	// [SearchProfilesRequest.histogram_queries][google.cloud.talent.v4beta1.SearchProfilesRequest.histogram_queries]
	// for details about syntax.
	HistogramQuery string `protobuf:"bytes,1,opt,name=histogram_query,json=histogramQuery,proto3" json:"histogram_query,omitempty"`
}

func (x *HistogramQuery) Reset() {
	*x = HistogramQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_talent_v4beta1_histogram_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistogramQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistogramQuery) ProtoMessage() {}

func (x *HistogramQuery) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4beta1_histogram_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistogramQuery.ProtoReflect.Descriptor instead.
func (*HistogramQuery) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4beta1_histogram_proto_rawDescGZIP(), []int{0}
}

func (x *HistogramQuery) GetHistogramQuery() string {
	if x != nil {
		return x.HistogramQuery
	}
	return ""
}

// Histogram result that matches
// [HistogramQuery][google.cloud.talent.v4beta1.HistogramQuery] specified in
// searches.
type HistogramQueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requested histogram expression.
	HistogramQuery string `protobuf:"bytes,1,opt,name=histogram_query,json=histogramQuery,proto3" json:"histogram_query,omitempty"`
	// A map from the values of the facet associated with distinct values to the
	// number of matching entries with corresponding value.
	//
	// The key format is:
	//
	//   - (for string histogram) string values stored in the field.
	//   - (for named numeric bucket) name specified in `bucket()` function, like
	//     for `bucket(0, MAX, "non-negative")`, the key will be `non-negative`.
	//   - (for anonymous numeric bucket) range formatted as `<low>-<high>`, for
	//     example, `0-1000`, `MIN-0`, and `0-MAX`.
	Histogram map[string]int64 `protobuf:"bytes,2,rep,name=histogram,proto3" json:"histogram,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *HistogramQueryResult) Reset() {
	*x = HistogramQueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_talent_v4beta1_histogram_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistogramQueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistogramQueryResult) ProtoMessage() {}

func (x *HistogramQueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4beta1_histogram_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistogramQueryResult.ProtoReflect.Descriptor instead.
func (*HistogramQueryResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4beta1_histogram_proto_rawDescGZIP(), []int{1}
}

func (x *HistogramQueryResult) GetHistogramQuery() string {
	if x != nil {
		return x.HistogramQuery
	}
	return ""
}

func (x *HistogramQueryResult) GetHistogram() map[string]int64 {
	if x != nil {
		return x.Histogram
	}
	return nil
}

var File_google_cloud_talent_v4beta1_histogram_proto protoreflect.FileDescriptor

var file_google_cloud_talent_v4beta1_histogram_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74,
	0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0x39, 0x0a, 0x0e, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x0f,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0xdd, 0x01, 0x0a, 0x14, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x5e, 0x0a, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x1a, 0x3c, 0x0a, 0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x72, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x34, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x74, 0x61, 0x6c, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x43, 0x54, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_talent_v4beta1_histogram_proto_rawDescOnce sync.Once
	file_google_cloud_talent_v4beta1_histogram_proto_rawDescData = file_google_cloud_talent_v4beta1_histogram_proto_rawDesc
)

func file_google_cloud_talent_v4beta1_histogram_proto_rawDescGZIP() []byte {
	file_google_cloud_talent_v4beta1_histogram_proto_rawDescOnce.Do(func() {
		file_google_cloud_talent_v4beta1_histogram_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_talent_v4beta1_histogram_proto_rawDescData)
	})
	return file_google_cloud_talent_v4beta1_histogram_proto_rawDescData
}

var file_google_cloud_talent_v4beta1_histogram_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_talent_v4beta1_histogram_proto_goTypes = []interface{}{
	(*HistogramQuery)(nil),       // 0: google.cloud.talent.v4beta1.HistogramQuery
	(*HistogramQueryResult)(nil), // 1: google.cloud.talent.v4beta1.HistogramQueryResult
	nil,                          // 2: google.cloud.talent.v4beta1.HistogramQueryResult.HistogramEntry
}
var file_google_cloud_talent_v4beta1_histogram_proto_depIdxs = []int32{
	2, // 0: google.cloud.talent.v4beta1.HistogramQueryResult.histogram:type_name -> google.cloud.talent.v4beta1.HistogramQueryResult.HistogramEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_talent_v4beta1_histogram_proto_init() }
func file_google_cloud_talent_v4beta1_histogram_proto_init() {
	if File_google_cloud_talent_v4beta1_histogram_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_talent_v4beta1_histogram_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistogramQuery); i {
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
		file_google_cloud_talent_v4beta1_histogram_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistogramQueryResult); i {
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
			RawDescriptor: file_google_cloud_talent_v4beta1_histogram_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_talent_v4beta1_histogram_proto_goTypes,
		DependencyIndexes: file_google_cloud_talent_v4beta1_histogram_proto_depIdxs,
		MessageInfos:      file_google_cloud_talent_v4beta1_histogram_proto_msgTypes,
	}.Build()
	File_google_cloud_talent_v4beta1_histogram_proto = out.File
	file_google_cloud_talent_v4beta1_histogram_proto_rawDesc = nil
	file_google_cloud_talent_v4beta1_histogram_proto_goTypes = nil
	file_google_cloud_talent_v4beta1_histogram_proto_depIdxs = nil
}
