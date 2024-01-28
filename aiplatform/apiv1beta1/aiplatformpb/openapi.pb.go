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
// source: google/cloud/aiplatform/v1beta1/openapi.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type contains the list of OpenAPI data types as defined by
// https://swagger.io/docs/specification/data-models/data-types/
type Type int32

const (
	// Not specified, should not be used.
	Type_TYPE_UNSPECIFIED Type = 0
	// OpenAPI string type
	Type_STRING Type = 1
	// OpenAPI number type
	Type_NUMBER Type = 2
	// OpenAPI integer type
	Type_INTEGER Type = 3
	// OpenAPI boolean type
	Type_BOOLEAN Type = 4
	// OpenAPI array type
	Type_ARRAY Type = 5
	// OpenAPI object type
	Type_OBJECT Type = 6
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "STRING",
		2: "NUMBER",
		3: "INTEGER",
		4: "BOOLEAN",
		5: "ARRAY",
		6: "OBJECT",
	}
	Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"STRING":           1,
		"NUMBER":           2,
		"INTEGER":          3,
		"BOOLEAN":          4,
		"ARRAY":            5,
		"OBJECT":           6,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1beta1_openapi_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1beta1_openapi_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_openapi_proto_rawDescGZIP(), []int{0}
}

// Schema is used to define the format of input/output data. Represents a select
// subset of an [OpenAPI 3.0 schema
// object](https://spec.openapis.org/oas/v3.0.3#schema). More fields may be
// added in the future as needed.
type Schema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The type of the data.
	Type Type `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.aiplatform.v1beta1.Type" json:"type,omitempty"`
	// Optional. The format of the data.
	// Supported formats:
	//
	//	for NUMBER type: float, double
	//	for INTEGER type: int32, int64
	Format string `protobuf:"bytes,7,opt,name=format,proto3" json:"format,omitempty"`
	// Optional. The description of the data.
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// Optional. Indicates if the value may be null.
	Nullable bool `protobuf:"varint,6,opt,name=nullable,proto3" json:"nullable,omitempty"`
	// Optional. Schema of the elements of Type.ARRAY.
	Items *Schema `protobuf:"bytes,2,opt,name=items,proto3" json:"items,omitempty"`
	// Optional. Possible values of the element of Type.STRING with enum format.
	// For example we can define an Enum Direction as :
	// {type:STRING, format:enum, enum:["EAST", NORTH", "SOUTH", "WEST"]}
	Enum []string `protobuf:"bytes,9,rep,name=enum,proto3" json:"enum,omitempty"`
	// Optional. Properties of Type.OBJECT.
	Properties map[string]*Schema `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. Required properties of Type.OBJECT.
	Required []string `protobuf:"bytes,5,rep,name=required,proto3" json:"required,omitempty"`
	// Optional. Example of the object. Will only populated when the object is the
	// root.
	Example *structpb.Value `protobuf:"bytes,4,opt,name=example,proto3" json:"example,omitempty"`
}

func (x *Schema) Reset() {
	*x = Schema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_openapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema) ProtoMessage() {}

func (x *Schema) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_openapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema.ProtoReflect.Descriptor instead.
func (*Schema) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_openapi_proto_rawDescGZIP(), []int{0}
}

func (x *Schema) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNSPECIFIED
}

func (x *Schema) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *Schema) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Schema) GetNullable() bool {
	if x != nil {
		return x.Nullable
	}
	return false
}

func (x *Schema) GetItems() *Schema {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Schema) GetEnum() []string {
	if x != nil {
		return x.Enum
	}
	return nil
}

func (x *Schema) GetProperties() map[string]*Schema {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Schema) GetRequired() []string {
	if x != nil {
		return x.Required
	}
	return nil
}

func (x *Schema) GetExample() *structpb.Value {
	if x != nil {
		return x.Example
	}
	return nil
}

var File_google_cloud_aiplatform_v1beta1_openapi_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_openapi_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa8, 0x04, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x3e, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x08, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x42, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x5c, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x07, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x1a, 0x66, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x65, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0b, 0x0a,
	0x07, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x52,
	0x52, 0x41, 0x59, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x10,
	0x06, 0x42, 0xe3, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x41,
	0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70,
	0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02,
	0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31,
	0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_openapi_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_openapi_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_openapi_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_openapi_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_openapi_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_openapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_openapi_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_openapi_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_openapi_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_aiplatform_v1beta1_openapi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_aiplatform_v1beta1_openapi_proto_goTypes = []interface{}{
	(Type)(0),              // 0: google.cloud.aiplatform.v1beta1.Type
	(*Schema)(nil),         // 1: google.cloud.aiplatform.v1beta1.Schema
	nil,                    // 2: google.cloud.aiplatform.v1beta1.Schema.PropertiesEntry
	(*structpb.Value)(nil), // 3: google.protobuf.Value
}
var file_google_cloud_aiplatform_v1beta1_openapi_proto_depIdxs = []int32{
	0, // 0: google.cloud.aiplatform.v1beta1.Schema.type:type_name -> google.cloud.aiplatform.v1beta1.Type
	1, // 1: google.cloud.aiplatform.v1beta1.Schema.items:type_name -> google.cloud.aiplatform.v1beta1.Schema
	2, // 2: google.cloud.aiplatform.v1beta1.Schema.properties:type_name -> google.cloud.aiplatform.v1beta1.Schema.PropertiesEntry
	3, // 3: google.cloud.aiplatform.v1beta1.Schema.example:type_name -> google.protobuf.Value
	1, // 4: google.cloud.aiplatform.v1beta1.Schema.PropertiesEntry.value:type_name -> google.cloud.aiplatform.v1beta1.Schema
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_openapi_proto_init() }
func file_google_cloud_aiplatform_v1beta1_openapi_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_openapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_openapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_openapi_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_openapi_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_openapi_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1beta1_openapi_proto_enumTypes,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_openapi_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_openapi_proto = out.File
	file_google_cloud_aiplatform_v1beta1_openapi_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_openapi_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_openapi_proto_depIdxs = nil
}
